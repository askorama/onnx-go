package debugger

// MsgType represents the type of the message sent on the wire
type MsgType byte

const (
	// ExecOp ...
	ExecOp MsgType = iota
	// InputsMsg ...
	InputsMsg
	// ResultMsg ...
	ResultMsg
	// StopMsg ...
	StopMsg
)

// DebugMsg ...
type DebugMsg struct {
	MsgType MsgType
	Payload []byte
}

// Instruction holds the informations about the current instruction
type Instruction struct {
	NodeID       int64
	Op           string
	Readfrom     []int
	Writeto      int
	CallsExtern  bool
	UseUnsafe    bool
	PreAllocated bool
}

// UnmarshalBinary modifies the receiver so it must take a pointer receiver.
/*
func (v *Instruction) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.Op, &v.Readfrom, &v.Writeto, &v.CallsExtern, &v.UseUnsafe, &v.PreAllocated)
	return err
}
*/
