package nerosql

import "errors"

var (
	ErrTableDoesNotExist         = errors.New("table does not exist")
	ErrTableAlreadyExists        = errors.New("table already exists")
	ErrIndexAlreadyExists        = errors.New("index already exists")
	ErrViolatesUniqueConstraint  = errors.New("duplicate key value violates unique constraint")
	ErrViolatesNotNullConstraint = errors.New("value violates not null constraint")
	ErrColumnDoesNotExist        = errors.New("column does not exist")
	ErrInvalidSelectItem         = errors.New("select item is not valid")
	ErrInvalidDatatype           = errors.New("invalid datatype")
	ErrMissingValues             = errors.New("missing values")
	ErrInvalidCell               = errors.New("cell is invalid")
	ErrInvalidOperands           = errors.New("operands are invalid")
	ErrPrimaryKeyAlreadyExists   = errors.New("primary key already exists")
	ErrModelDoesNotExist         = errors.New("model does not exist")
	ErrModelArgsException        = errors.New("model args raise an exception")
)
