package errors

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/NoHomey/chaos-go-camp-proj/mysql/errcode"
	"github.com/go-sql-driver/mysql"
)

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

//QueryErr represents error from executing a mysql qry.
type QueryErr interface {
	Query() *Query
	Error() string
	Unwrap() error
}

//Wrap wraps the given error.
func Wrap(err error, sqlText string, args []interface{}) QueryErr {
	if err == nil {
		panic(errors.New("Expecting non nil error"))
	}
	query := &Query{sqlText, args}
	if err == sql.ErrNoRows {
		return ErrNotFound{query}
	}
	if driverErr, ok := err.(*mysql.MySQLError); ok {
		switch driverErr.Number {
		case errcode.DuplicateEntry:
			return ErrExists{err, query}
		case errcode.ConstraintViolated:
			return ErrConstraint{err, query}
		}
	}
	return UnknownErr{err, query}
}

//UnknownErr represents MySQL error wich is not recognized.
type UnknownErr struct {
	err error
	qry *Query
}

func (err UnknownErr) Unwrap() error {
	return err.err
}

func (err UnknownErr) Error() string {
	return fmt.Sprintf("MySQL error: %s", err.err.Error())
}

//Query implements EmptyResultErr.
func (err UnknownErr) Query() *Query {
	return err.qry
}

//ErrExists is used to report that a record alredy exists.
type ErrExists struct {
	err error
	qry *Query
}

func (err ErrExists) Error() string {
	return "MySQL error: Record already exists"
}

//Query implements EmptyResultErr.
func (err ErrExists) Query() *Query {
	return err.qry
}

//Unwrap implements EmptyResultErr.
func (err ErrExists) Unwrap() error {
	return err.err
}

//ErrNotFound reports that query returned empty result.
type ErrNotFound struct {
	qry *Query
}

func (err ErrNotFound) Error() string {
	return fmt.Sprintf("MySQL error: Result not found for MySQL query")
}

//Query implements EmptyResultErr.
func (err ErrNotFound) Query() *Query {
	return err.qry
}

//Unwrap implements EmptyResultErr.
func (err ErrNotFound) Unwrap() error {
	return sql.ErrNoRows
}

//ErrConstraint is used to report that a check constaint was violated.
type ErrConstraint struct {
	err error
	qry *Query
}

func (err ErrConstraint) Error() string {
	return "MySQL error: Check constrained violated"
}

//Query implements EmptyResultErr.
func (err ErrConstraint) Query() *Query {
	return err.qry
}

//Unwrap implements EmptyResultErr.
func (err ErrConstraint) Unwrap() error {
	return err.err
}
