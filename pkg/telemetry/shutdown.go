package telemetry

import (
	ctx "context"
	"errors"
	"go-template/pkg/logger"
)

type shutdownHandler struct {
	funcs []func(ctx.Context) error
}

// NewShutdownHandler creates a new shutdown handler
func newShutdownHandler() *shutdownHandler {
	return &shutdownHandler{
		funcs: make([]func(ctx.Context) error, 0),
	}
}

// AddFunction adds a shutdown function to the handler
func (handle *shutdownHandler) addFunction(function func(ctx.Context) error) {
	handle.funcs = append(handle.funcs, function)
}

// Shutdown executes all registered shutdown functions
func (handle *shutdownHandler) shutdown(ctx ctx.Context) error {
	var err error
	for _, functions := range handle.funcs {
		if ferr := functions(ctx); ferr != nil {
			err = errors.Join(err, ferr)
			logger.Errorf("shutdown function failed: %v", ferr)
		}
	}

	handle.funcs = nil
	return err
}
