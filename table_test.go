package telegrammarkdown_test

import (
	"testing"

	md "github.com/onuruluag/telegram-markdown-go"
)

func TestTableAddRow(t *testing.T) {
	table := md.Table{}
	table.AddColumns(md.Column{Width: 4})
	table.AddRow("a")

	got := table.String()
	want := "`   a`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableAddMoreCellsThanColumns(t *testing.T) {
	table := md.Table{}
	table.AddColumns(md.Column{Width: 4})
	table.AddRow("a", "b")

	got := table.String()
	want := "`   a`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableAddLessCellsThanColumns(t *testing.T) {
	table := md.Table{}
	table.AddColumns(
		md.Column{Width: 4},
		md.Column{Width: 4},
	)
	table.AddRow("a")

	got := table.String()
	want := "`   a`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableLeftAlign(t *testing.T) {
	table := md.Table{}
	table.AddColumns(
		md.Column{Width: 4, Align: md.Left},
		md.Column{Width: 4},
	)
	table.AddRow("a", "b")

	got := table.String()
	want := "`a   " + "   b`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableWithSeparator(t *testing.T) {
	table := md.Table{Separator: "|"}
	table.AddColumns(
		md.Column{Width: 4, Align: md.Left},
		md.Column{Width: 4},
	)
	table.AddRow("a", "b")

	got := table.String()
	want := "`a   |" + "   b`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableWithSeparatorAndMargin(t *testing.T) {
	table := md.Table{Separator: "|"}
	table.AddColumns(
		md.Column{Width: 4, Margin: 1},
		md.Column{Width: 4, Margin: 1, Align: md.Left},
	)
	table.AddRow("a", "b")

	got := table.String()
	want := "`    a | b    `"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableWithSeparatorAndMarginMoreRows(t *testing.T) {
	table := md.Table{Separator: "|"}
	table.AddColumns(
		md.Column{Width: 4, Margin: 1},
		md.Column{Width: 4, Margin: 1, Align: md.Left},
	)
	table.AddRow("a", "b")
	table.AddRow("aa", "bb")
	table.AddRow("aaa", "bbb")

	got := table.String()
	want := "`    a | b    `\n"
	want += "`   aa | bb   `\n"
	want += "`  aaa | bbb  `"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableEscapes(t *testing.T) {
	table := md.Table{}
	table.AddColumns(md.Column{Width: 4})
	table.AddRow("`a`")

	got := table.String()
	want := "` \\`a\\``"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestChangingSeparatorOnTheGo(t *testing.T) {
	table := md.Table{Separator: "|"}
	table.AddColumns(
		md.Column{Width: 4, Margin: 1},
		md.Column{Width: 4, Margin: 1, Align: md.Left},
	)
	table.AddRow("a", "b")
	table.AddRow("aa", "bb")
	table.AddRow("aaa", "bbb")

	table.Separator = "-"

	got := table.String()
	want := "`    a - b    `\n"
	want += "`   aa - bb   `\n"
	want += "`  aaa - bbb  `"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableWithHeader(t *testing.T) {
	table := md.Table{}
	table.AddColumns(
		md.Column{Width: 4, Align: md.Left},
		md.Column{Width: 4},
	)
	table.AddRow("a", "b")
	table.SetHeader("c1", "c2")

	got := table.String()
	want := "`c1  " + "  c2`\n"
	want += "`a   " + "   b`"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}

func TestTableCodeBlock(t *testing.T) {
	table := md.Table{CodeBlock: true}
	table.AddColumns(
		md.Column{Width: 4, Align: md.Left},
		md.Column{Width: 4},
	)
	table.AddRow("a", "b")
	table.SetHeader("c1", "c2")

	got := table.String()
	want := "```\n"
	want += "c1  " + "  c2\n"
	want += "a   " + "   b\n"
	want += "```"

	if got != want {
		t.Error(errorMessage(got, want))
	}
}
