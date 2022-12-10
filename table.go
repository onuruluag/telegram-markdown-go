package telegrammarkdown

import (
	"fmt"
	"strings"
)

type Alignment int

const (
	Right Alignment = iota
	Center
	Left
)

type Column struct {
	Header string
	Align  Alignment
	Width  int
	Margin uint
}

func (c Column) getCellFormat() string {
	width := c.Width
	if c.Align == Left {
		width = -width
	}
	margin := strings.Repeat(" ", int(c.Margin))
	format := fmt.Sprintf("%s%%%ds%s", margin, width, margin)
	return format
}

type Table struct {
	Separator string
	CodeBlock bool
	columns   []Column
	rows      [][]string
	header    []string
}

func (t *Table) AddColumns(columns ...Column) {
	t.columns = append(t.columns, columns...)
}

func (t *Table) AddRow(cells ...string) {
	row := make([]string, 0, len(t.columns))
	for i, col := range t.columns {
		if i >= len(cells) {
			break
		}
		row = append(row, fmt.Sprintf(col.getCellFormat(), cells[i]))
	}
	t.rows = append(t.rows, row)
}

func (t *Table) SetHeader(cells ...string) {
	t.header = make([]string, 0, len(t.columns))
	for i, col := range t.columns {
		if i >= len(cells) {
			break
		}
		t.header = append(t.header, fmt.Sprintf(col.getCellFormat(), cells[i]))
	}
}

func (t *Table) String() string {
	if t.CodeBlock {
		return t.blockEscapedString()
	}
	return t.perColumnEscapedString()
}

func (t *Table) perColumnEscapedString() string {
	var data string
	add := func(row []string) {
		joined := strings.Join(row, t.Separator)
		escaped := escape(joined, "`\\")
		data += encloseText(escaped, "`", "`") + "\n"
	}
	if len(t.header) > 0 {
		add(t.header)
	}
	for _, row := range t.rows {
		add(row)
	}
	return strings.TrimSuffix(data, "\n")
}

func (t *Table) blockEscapedString() string {
	var data string
	add := func(row []string) {
		joined := strings.Join(row, t.Separator)
		data += escape(joined, "`\\") + "\n"
	}
	if len(t.header) > 0 {
		add(t.header)
	}
	for _, row := range t.rows {
		add(row)
	}
	data = encloseText(data, "```\n", "```")
	return data
}
