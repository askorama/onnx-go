// +build !cuda

package engine

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/pkg/errors"
	"gorgonia.org/gorgonia/debugger"
	"gorgonia.org/gorgonia/internal/execution"
	"gorgonia.org/gorgonia/internal/value"
	"gorgonia.org/gorgonia/ops"
	"gorgonia.org/tensor"
)

func finalizeTapeMachine(m *tapeMachine) {}

// UseCudaFor is an option for *tapeMachine. This function is NO-OP unless the program is built with the `cuda` tag.
func UseCudaFor(ops ...string) VMOpt {
	return func(m VM) {}
}

func (m *tapeMachine) getEngine(dev execution.Device) tensor.Engine { return m.Engine }

func (instr *execOp) exec(m *tapeMachine) (err error) {
	var enc *gob.Encoder
	var network bytes.Buffer
	if m.c != nil {
		enc = gob.NewEncoder(&network)

		defer func(c chan debugger.DebugMsg) {
			c <- debugger.DebugMsg{
				MsgType: debugger.ExecOp,
				Payload: network.Bytes(),
			}
		}(m.c)
		readfrom := make([]int, len(instr.readFrom))
		for i := 0; i < len(instr.readFrom); i++ {
			readfrom[i] = instr.readFrom[i].id
		}
		enc.Encode(debugger.Instruction{
			NodeID:       instr.id,
			Op:           instr.op.String(),
			Readfrom:     readfrom,
			Writeto:      instr.writeTo.id,
			CallsExtern:  instr.op.CallsExtern(),
			UseUnsafe:    instr.useUnsafe,
			PreAllocated: instr.preAllocated,
		})
	}

	//m.logf("Executing %v. Node is: %x", instr, instr.id)
	//m.enterLogScope()
	//defer m.leaveLogScope()

	// Read
	//m.watchedLogf("Inputs:")
	//m.enterLogScope()
	var inputs []value.Value
	for _, reg := range instr.readFrom {
		v := m.cpumem[reg.id]
		inputs = append(inputs, v)
		if m.c != nil {
			err := enc.Encode(v)
			if err != nil {
				log.Println(err)
			}
		}
		//	m.watchedLogf(m.valueFmt, v)
	}
	//m.leaveLogScope()

	// Execute
	var v value.Value
	switch {
	case instr.preAllocated:
		if pd, ok := instr.op.(ops.UsePreallocDoer); ok {
			p := m.cpumem[instr.writeTo.id]
			if v, err = pd.UsePreallocDo(p, inputs...); err != nil {
				return errors.Wrapf(err, "Happened while attempting to execute %v. Node is %x. Register was: %v ", instr, instr.id, instr.writeTo.id)
			}
		} else {
			// TODO: maybe warn?
			if v, err = instr.op.Do(inputs...); err != nil {
				return errors.Wrap(err, opDoFail)
			}
		}
	case instr.useUnsafe:
		if ud, ok := instr.op.(ops.UnsafeDoer); ok {
			if v, err = ud.UnsafeDo(inputs...); err != nil {
				return errors.Wrap(err, "Failed to carry UnsafeDo()")
			}
		} else {
			// TODO: warn?
			if v, err = instr.op.Do(inputs...); err != nil {
				return errors.Wrap(err, opDoFail)
			}
		}
	default:
		if v, err = instr.op.Do(inputs...); err != nil {
			return errors.Wrap(err, opDoFail)
		}
	}
	if m.c != nil {
		enc.Encode(v)
	}
	//m.watchedLogf("Result:")
	//m.enterLogScope()
	//m.watchedLogf(m.valueFmt, v)
	//m.leaveLogScope()
	// TODO: type and shape checks

	// Write
	value.SetEngine(v, m.Engine)
	dest := instr.writeTo.id
	m.cpumem[dest] = v
	node := m.p.g.Node(instr.id).(*Node)

	if m.trace() && (len(m.watchNodes) == 0 || m.watchNodes.Contains(node)) {
		if err = node.bindCopy(v); err != nil {
			return errors.Wrapf(err, "TraceExec failed to bind copy")
		}
	} else {
		node.bind(v)
	}

	// this is a gradient node then, we should also bind the value to the node's value.DualValue
	if m.bindDV() && node.derivOf != nil {
		for _, src := range node.derivOf {
			if len(m.bindNodesDV) > 0 && !m.bindNodesDV.Contains(src) {
				continue
			}

			if src.boundTo != nil {
				dv := dvUnit(src.boundTo)

				add := newEBOByType(addOpType, value.TypeOf(dv.D), value.TypeOf(v))

				if d, err := add.UnsafeDo(dv.D, v); err == nil {
					dv.SetDeriv(d)
					src.bind(dv)
				} else {
					return err
				}
			}
		}

	}

	m.watchedLogf("Written To: %v", instr.writeTo)
	m.enterLogScope()
	m.watchedLogf(m.valueFmt, v)
	m.leaveLogScope()
	return nil
}

func (instr deviceTransport) exec(m *tapeMachine) error {
	return nil
}
