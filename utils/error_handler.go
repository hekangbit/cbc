package utils

import (
	"fmt"
	"io"
	"os"
)

type ErrorHandler struct {
	programID string
	writer    io.Writer
	nError    int64
	nWarning  int64
}

func NewErrorHandler(progid string) *ErrorHandler {
	return &ErrorHandler{
		programID: progid,
		writer:    os.Stderr,
	}
}

func NewErrorHandlerWithWriter(progid string, w io.Writer) *ErrorHandler {
	return &ErrorHandler{
		programID: progid,
		writer:    w,
	}
}

func (this *ErrorHandler) Error(loc string, msg string) {
	this.ErrorMsg(loc + ": " + msg)
}

func (this *ErrorHandler) ErrorMsg(msg string) {
	fmt.Fprintf(this.writer, "%s: error: %s\n", this.programID, msg)
	this.nError++
}

func (this *ErrorHandler) Warn(loc string, msg string) {
	this.WarnMsg(loc + ": " + msg)
}

func (this *ErrorHandler) WarnMsg(msg string) {
	fmt.Fprintf(this.writer, "%s: warning: %s\n", this.programID, msg)
	this.nWarning++
}

func (this *ErrorHandler) ErrorOccurred() bool {
	return this.nError > 0
}
