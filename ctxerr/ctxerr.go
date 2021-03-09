package ctxerr

import (
	"fmt"
	"strings"
)

//Context represents error contextual information.
type Context struct {
	Name string
	Data interface{}
}

//Contexter can return Context.
type Contexter interface {
	Context() Context
}

//Error represents service error.
type Error interface {
	error
	Unwrap() error
	Text() string
	Contexter
}

//InvalData wraps error due to invalid data.
type InvalData struct {
	wrapped error
	invalid map[string]string
}

func (e InvalData) Error() string {
	msg := "Recived invalid data for fields: %v. Error: %s"
	return fmt.Sprintf(msg, e.invalid, e.wrapped.Error())
}

func (e InvalData) Unwrap() error {
	return e.wrapped
}

//Text returns human readable error text.
func (e InvalData) Text() string {
	var sb strings.Builder
	sb.WriteString("Recived invalid data: ")
	last := len(e.invalid) - 1
	i := 0
	for field, failed := range e.invalid {
		sb.WriteString(field)
		sb.WriteString(" > ")
		sb.WriteString(failed)
		if i != last {
			sb.WriteByte(',')
			sb.WriteByte(' ')
		}
		i++
	}
	return sb.String()
}

//Context returns error Context.
func (e InvalData) Context() Context {
	return Context{
		Name: "invalid-data",
		Data: e.invalid,
	}
}

//NewInvalData constructs InvalData error.
func NewInvalData(err error, invalid map[string]string) Error {
	return InvalData{err, invalid}
}

//NotAuthed wraps error due to Authentication.
type NotAuthed struct {
	wrapped error
}

func (e NotAuthed) Error() string {
	return fmt.Sprintf("Not Authenticated. Error: %s", e.wrapped.Error())
}

func (e NotAuthed) Unwrap() error {
	return e.wrapped
}

//Text returns human readable error text.
func (e NotAuthed) Text() string {
	return "Not Authenticated"
}

//Context returns error Context.
func (e NotAuthed) Context() Context {
	return Context{
		Name: "not-authenticated",
	}
}

//NewNotAuthed constructs NotAuthed error.
func NewNotAuthed(err error) Error {
	return NotAuthed{err}
}

//Internal wraps error signaling it is an internal one.
type Internal struct {
	wrapped error
}

func (e Internal) Error() string {
	return fmt.Sprintf("Internal error: %s", e.wrapped.Error())
}

func (e Internal) Unwrap() error {
	return e.wrapped
}

//Text returns human readable error text.
func (e Internal) Text() string {
	return "Internal error"
}

//Context returns error Context.
func (e Internal) Context() Context {
	return Context{
		Name: "bad-format",
	}
}

//NewInternal constructs Internal error.
func NewInternal(err error) Error {
	return Internal{err}
}

//BadFormat wraps error due to bad formating.
type BadFormat struct {
	wrapped error
}

func (e BadFormat) Error() string {
	return fmt.Sprintf("Bad formating. Error: %s", e.wrapped.Error())
}

func (e BadFormat) Unwrap() error {
	return e.wrapped
}

//Text returns human readable error text.
func (e BadFormat) Text() string {
	return "Bad data format"
}

//Context returns error Context.
func (e BadFormat) Context() Context {
	return Context{
		Name: "bad-format",
	}
}

//NewBadFormat constructs BadFormat error.
func NewBadFormat(err error) Error {
	return BadFormat{err}
}
