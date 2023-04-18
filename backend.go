package nerosql

import "errors"

type ColumnType uint

const (
	TextType ColumnType = iota
	IntType
	BoolType
)

func (c ColumnType) String() string {
	switch c {
	case TextType:
		return "TextType"
	case IntType:
		return "IntType"
	case BoolType:
		return "BoolType"
	default:
		return "Error"
	}
}

type Cell interface {
	AsText() *string
	AsInt() *int32
	AsBool() *bool
}

type Results struct {
	Columns []ResultColumn
	Rows    [][]Cell
}

type ResultColumn struct {
	Type    ColumnType
	Name    string
	NotNull bool
}

type Index struct {
	Name       string
	Exp        string
	Type       string
	Unique     bool
	PrimaryKey bool
}

type TableMetadata struct {
	Name    string
	Columns []ResultColumn
	Indexes []Index
}

type PredResult string
type TrainResult string

type Backend interface {
	CreateTable(*CreateTableStatement) error
	DropTable(*DropTableStatement) error
	CreateIndex(*CreateIndexStatement) error
	Insert(*InsertStatement) error
	Select(*SelectStatement) (*Results, error)
	CreateModel(*CreateModelStatement) error
	TrainModel(*TrainModelStatement) (TrainResult, error)
	Predict(*PredictStatement) (PredResult, error)
	GetTables() []TableMetadata
}

// Useful to embed when prototyping new backends
type EmptyBackend struct{}

func (eb EmptyBackend) CreateTable(_ *CreateTableStatement) error {
	return errors.New("create not supported")
}

func (eb EmptyBackend) DropTable(_ *DropTableStatement) error {
	return errors.New("drop not supported")
}

func (eb EmptyBackend) CreateIndex(_ *CreateIndexStatement) error {
	return errors.New("create index not supported")
}

func (eb EmptyBackend) Insert(_ *InsertStatement) error {
	return errors.New("insert not supported")
}

func (eb EmptyBackend) Select(_ *SelectStatement) (*Results, error) {
	return nil, errors.New("select not supported")
}

func (eb EmptyBackend) CreateModel(_ *CreateModelStatement) error {
	return errors.New("create model not supported")
}

func (eb EmptyBackend) TrainModel(_ *TrainModelStatement) (TrainResult, error) {
	return "", errors.New("train model not supported")
}

func (eb EmptyBackend) Predict(_ *PredictStatement) (PredResult, error) {
	return "", errors.New("predict not supported")
}

func (eb EmptyBackend) GetTables() []TableMetadata {
	return nil
}
