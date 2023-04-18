package main

import (
	"fmt"
	"strings"
)

type ExpressionKind uint

const (
	LiteralKind ExpressionKind = iota
	BinaryKind
)

type BinaryExpression struct {
	A  Expression
	B  Expression
	Op Token
}

func (be BinaryExpression) GenerateCode() string {
	return fmt.Sprintf("(%s %s %s)", be.A.GenerateCode(), be.Op.Value, be.B.GenerateCode())
}

type Expression struct {
	Literal *Token
	Binary  *BinaryExpression
	Kind    ExpressionKind
}

func (e Expression) GenerateCode() string {
	switch e.Kind {
	case LiteralKind:
		switch e.Literal.Kind {
		case IdentifierKind:
			return fmt.Sprintf("\"%s\"", e.Literal.Value)
		case StringKind:
			return fmt.Sprintf("'%s'", e.Literal.Value)
		default:
			return fmt.Sprintf(e.Literal.Value)
		}

	case BinaryKind:
		return e.Binary.GenerateCode()
	}

	return ""
}

type SelectItem struct {
	Exp      *Expression
	Asterisk bool // for *
	As       *Token
}

type SelectStatement struct {
	Item   *[]*SelectItem
	From   *Token
	Where  *Expression
	Limit  *Expression
	Offset *Expression
}

func (ss SelectStatement) GenerateCode() string {
	item := []string{}
	for _, i := range *ss.Item {
		s := "\t*"
		if !i.Asterisk {
			s = "\t" + i.Exp.GenerateCode()

			if i.As != nil {
				s = fmt.Sprintf("\t%s AS \"%s\"", s, i.As.Value)
			}
		}
		item = append(item, s)
	}

	code := "SELECT\n" + strings.Join(item, ",\n")
	if ss.From != nil {
		code += fmt.Sprintf("\nFROM\n\t\"%s\"", ss.From.Value)
	}

	if ss.Where != nil {
		code += "\nWHERE\n\t" + ss.Where.GenerateCode()
	}

	if ss.Limit != nil {
		code += "\nLIMIT\n\t" + ss.Limit.GenerateCode()
	}

	if ss.Offset != nil {
		code += "\nOFFSET\n\t" + ss.Limit.GenerateCode()
	}

	return code + ";"
}

type ColumnDefinition struct {
	Name       Token
	Datatype   Token
	PrimaryKey bool
}

type CreateTableStatement struct {
	Name Token
	Cols *[]*ColumnDefinition
}

func (cts CreateTableStatement) GenerateCode() string {
	cols := []string{}
	for _, col := range *cts.Cols {
		modifiers := ""
		if col.PrimaryKey {
			modifiers += " " + "PRIMARY KEY"
		}
		spec := fmt.Sprintf("\t\"%s\" %s%s", col.Name.Value, strings.ToUpper(col.Datatype.Value), modifiers)
		cols = append(cols, spec)
	}
	return fmt.Sprintf("CREATE TABLE \"%s\" (\n%s\n);", cts.Name.Value, strings.Join(cols, ",\n"))
}

type CreateIndexStatement struct {
	Name       Token
	Unique     bool
	PrimaryKey bool
	Table      Token
	Exp        Expression
}

type DeleteStatement struct {
	TableName string
	Filter    string
}

func (cis CreateIndexStatement) GenerateCode() string {
	unique := ""
	if cis.Unique {
		unique = " UNIQUE"
	}
	return fmt.Sprintf("CREATE%s INDEX \"%s\" ON \"%s\" (%s);", unique, cis.Name.Value, cis.Table.Value, cis.Exp.GenerateCode())
}

type DropTableStatement struct {
	Name Token
}

func (dts DropTableStatement) GenerateCode() string {
	return fmt.Sprintf("DROP TABLE \"%s\";", dts.Name.Value)
}

type InsertStatement struct {
	Table  Token
	Values *[]*Expression
}

func (is InsertStatement) GenerateCode() string {
	values := []string{}
	for _, exp := range *is.Values {
		values = append(values, exp.GenerateCode())
	}
	return fmt.Sprintf("INSERT INTO \"%s\" VALUES (%s);", is.Table.Value, strings.Join(values, ", "))
}

func (cms CreateModelStatement) GenerateCode() string {
	args := []string{}
	for _, arg := range cms.Args {
		args = append(args, arg.Value)
	}
	return fmt.Sprintf("CREATE MODEL \"%s\" AS %s(%s) AS \"%s\" FROM \"%s\";", cms.Name.Value, strings.ToUpper(cms.ModelType.Value), strings.Join(args, ", "), cms.PredOutput.Value, cms.Table.Value)
}

func (tms TrainModelStatement) GenerateCode() string {
	args := []string{}
	for _, arg := range tms.Args {
		args = append(args, arg.Value)
	}
	return fmt.Sprintf("TRAIN MODEL \"%s\" WITH (%s);", tms.Name.Value, strings.Join(args, ", "))
}

func (prds PredictStatement) GenerateCode() string {
	inputs := []string{}
	for _, in := range prds.Input {
		inputs = append(inputs, in.Value)
	}
	return fmt.Sprintf("PREDICT \"%s\" (%s) FROM \"%s\";", prds.Name.Value, strings.Join(inputs, ", "), prds.Table.Value)
}

type AstKind uint

const (
	SelectKind AstKind = iota
	CreateTableKind
	CreateIndexKind
	DropTableKind
	InsertKind
	CreateModelKind
	TrainModelKind
	PredictKind
)

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	CreateIndexStatement *CreateIndexStatement
	DropTableStatement   *DropTableStatement
	InsertStatement      *InsertStatement
	CreateModelStatement *CreateModelStatement
	TrainModelStatement  *TrainModelStatement
	PredictStatement     *PredictStatement
	Kind                 AstKind
}

type CreateModelStatement struct {
	Name       Token
	ModelType  Token
	Args       []*Token
	PredOutput Token
	Table      Token
}

type TrainModelStatement struct {
	Name Token
	Args []*Token
}

type PredictStatement struct {
	Name  Token
	Input []*Token
	Table Token
}

func (s Statement) GenerateCode() string {
	switch s.Kind {
	case SelectKind:
		return s.SelectStatement.GenerateCode()
	case CreateTableKind:
		return s.CreateTableStatement.GenerateCode()
	case CreateIndexKind:
		return s.CreateIndexStatement.GenerateCode()
	case DropTableKind:
		return s.DropTableStatement.GenerateCode()
	case InsertKind:
		return s.InsertStatement.GenerateCode()
	case CreateModelKind:
		return s.CreateModelStatement.GenerateCode()
	case TrainModelKind:
		return s.TrainModelStatement.GenerateCode()
	case PredictKind:
		return s.PredictStatement.GenerateCode()
	}

	return "?unknown?"
}

type Ast struct {
	Statements []*Statement
}
