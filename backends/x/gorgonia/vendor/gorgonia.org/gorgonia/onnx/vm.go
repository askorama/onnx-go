package onnx

import (
	"log"

	"gorgonia.org/gorgonia/debugger"
	"gorgonia.org/gorgonia/internal/engine"
)

// WithLogger creates a VM with the supplied logger. If the logger is nil, a default logger, writing to os.stderr will be created.
func WithLogger(logger *log.Logger) engine.VMOpt {
	return engine.WithLogger(logger)
}

// WithWatchlist creates a VM with a watchlist. When the execution touches the things in the watchlist, the VM's logger will the log it.
// This allows for watching and finetuning of the algorithm. When nothing is passed in, then the VM will default to watching and logging every single
// execution object.
//
// The watchlist allows for different things to be watched, depending on VM type:
//		*lispMachine will ONLY take *Node
//		*tapeMachine will take int (for register IDs) or *Node.
func WithWatchlist(list ...interface{}) engine.VMOpt {
	return engine.WithWatchlist(list...)
}

// WithValueFmt defines how the logger will output the values. It defaults to "%3.3f"
func WithValueFmt(format string) engine.VMOpt {
	return engine.WithValueFmt(format)
}

// WithDebuggingChannel receives information at runtime
func WithDebuggingChannel(c chan debugger.DebugMsg) engine.VMOpt {
	return engine.WithDebuggingChannel(c)
}
