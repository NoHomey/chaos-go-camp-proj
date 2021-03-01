package errors

import (
	"fmt"
	"strings"
)

//QueryErr represents error from executing a mysql query.
type QueryErr struct {
	wrapped error
}

func (err QueryErr) Unwrap() error {
	return err.wrapped
}

func (err QueryErr) Error() string {
	return fmt.Sprintf("MySQL error: %s", err.wrapped.Error())
}

//EmptyResultErr is error wich signals that sql query returned empty result set.
type EmptyResultErr interface {
	Query() *Query
	Error() string
}

//NotFoundErr reports that single row query returned empty result.
type NotFoundErr struct {
	query *Query
}

func (err NotFoundErr) Error() string {
	return fmt.Sprintf("Single row query: %s did not found it's result", err.query.Text())
}

//Query implements EmptyResultErr.
func (err NotFoundErr) Query() *Query {
	return err.query
}

//Query represents SQL query.
type Query struct {
	SQL  string
	Args []interface{}
}

//Text returns valid SQL query text with bound paramiters.
func (q *Query) Text() string {
	format := strings.Replace(q.SQL, "?", "%v", len(q.Args))
	return fmt.Sprintf(format, q.Args...)
}
