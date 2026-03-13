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

func (this *ErrorHandler) ErrorWithLoc(loc fmt.Stringer, msg string) {
	this.Error(loc.String() + ": " + msg)
}

func (this *ErrorHandler) Error(msg string) {
	fmt.Fprintf(this.writer, "%s: error: %s\n", this.programID, msg)
	this.nError++
}

func (this *ErrorHandler) WarnWithLoc(loc fmt.Stringer, msg string) {
	this.Warn(loc.String() + ": " + msg)
}

func (this *ErrorHandler) Warn(msg string) {
	fmt.Fprintf(this.writer, "%s: warning: %s\n", this.programID, msg)
	this.nWarning++
}

func (this *ErrorHandler) ErrorOccured() bool {
	return this.nError > 0
}

func (this *ErrorHandler) IssueOccured() bool {
	return this.nError > 0 || this.nWarning > 0
}

func (this *ErrorHandler) DumpTotal() {
	fmt.Fprintf(this.writer, "total %d errors, %d warnings\n", this.nError, this.nWarning)
}
