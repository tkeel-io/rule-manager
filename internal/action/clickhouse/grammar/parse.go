package grammar

import (
	"context"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/tkeel-io/rule-manager/internal/action/clickhouse/grammar/parser"
)

type ClickHouseListener struct {
	ctx context.Context
	*parser.BaseClickHouseParserListener
	stack           []string
	Fields          []*Field
	PartitionFields []string
	OrderByFields   []string
	errors          []error
}

//PARTITION BY action_date ORDER BY

type Field struct {
	Name string
	Type string
}

type TableInfo struct {
	Name            string
	Fields          []*Field
	PartitionFields []string
	OrderByFields   []string
}

func Parse(expr string) *ClickHouseListener {
	// Setup the input
	is := antlr.NewInputStream(expr)

	// Create the Lexer
	lexer := parser.NewClickHouseLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewClickHouseParser(stream)

	// Finally parseField the expression (by walking the tree)
	var listener ClickHouseListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Parse())

	return &listener
}

func (l *ClickHouseListener) ExitColumn_declaration(c *parser.Column_declarationContext) {
	//column_declaration_list
	column_name := c.Column_name().GetText()
	column_name = strings.ReplaceAll(column_name, "`", "")
	column_type := c.Column_type().GetText()
	l.Fields = append(l.Fields, &Field{column_name, column_type})
	return
}

func (l *ClickHouseListener) EnterOrder_by_expression_list(c *parser.Order_by_expression_listContext) {
	//column_declaration_list
	l.stack = make([]string, 0)
	//fmt.Println("##EnterOrder_by_expression_list", l.stack)
	return
}

func (l *ClickHouseListener) ExitOrder_by_expression_list(c *parser.Order_by_expression_listContext) {
	//column_declaration_list
	//fmt.Println("##EnterOrder_by_expression_list", l.stack)
	for _, e := range l.stack {
		l.OrderByFields = append(l.OrderByFields, e)
	}
	return
}

func (l *ClickHouseListener) EnterPartition_by_element(c *parser.Partition_by_elementContext) {
	l.stack = make([]string, 0)
	return
}

func (l *ClickHouseListener) ExitPartition_by_element(c *parser.Partition_by_elementContext) {
	//column_declaration_list
	for _, e := range l.stack {
		l.PartitionFields = append(l.PartitionFields, e)
	}
	return
}

func (l *ClickHouseListener) ExitCompound_identifier(c *parser.Compound_identifierContext) {
	//column_declaration_list
	l.stack = append(l.stack, c.GetText())
	return
}

func (l *ClickHouseListener) Table(name string) (*TableInfo, error) {
	table := &TableInfo{
		Name:            name,
		Fields:          []*Field{},
		PartitionFields: []string{},
		OrderByFields:   []string{},
	}
	//Fields
	for _, field := range l.Fields {
		table.Fields = append(table.Fields, &Field{
			Name: field.Name,
			Type: field.Type,
		})
	}
	//PartitionFields
	for _, field := range l.PartitionFields {
		table.PartitionFields = append(table.PartitionFields, field)
	}
	//OrderByFields
	for _, field := range l.OrderByFields {
		table.OrderByFields = append(table.OrderByFields, field)
	}
	var err error
	if le := len(l.errors); le > 0 {
		err = l.errors[le-1]
	}
	return table, err
}
