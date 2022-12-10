package telegrammarkdown_test

import (
	"testing"

	md "github.com/onuruluag/telegram-markdown-go"
)

func TestSimpleText(t *testing.T) {
	got := md.Text("hello")
	want := `hello`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestTextWithSpecialCharacters(t *testing.T) {
	got := md.Text("hello * World")
	want := `hello \* World`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleBold(t *testing.T) {
	got := md.Bold(md.Text("bold text"))
	want := `*bold text*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleBoldAndBoldText(t *testing.T) {
	got := md.Bold(md.Text("bold text"))
	want := md.BoldText("bold text")

	if got != want {
		t.Errorf(errorMessage(got.String(), want.String()))
	}
}

func TestBoldWithInlineAsterix(t *testing.T) {
	got := md.Bold(md.Text("bold *text"))
	want := `*bold \*text*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleItalic(t *testing.T) {
	got := md.Italic(md.Text("italic text"))
	want := `_italic text_`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleItalicAndItalicText(t *testing.T) {
	got := md.Italic(md.Text("italic text"))
	want := md.ItalicText("italic text")

	if got != want {
		t.Errorf(errorMessage(got.String(), want.String()))
	}
}

func TestItalicWithInlineAsterix(t *testing.T) {
	got := md.Italic(md.Text("italic *text"))
	want := `_italic \*text_`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestBoldAndItalic(t *testing.T) {
	got := md.Bold(md.Italic(md.Text("italic text")))
	want := `*_italic text_*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestBoldAndItalicWithInlineAsterix(t *testing.T) {
	got := md.Bold(md.Italic(md.Text("italic *text")))
	want := `*_italic \*text_*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleUnderline(t *testing.T) {
	got := md.Bold(md.Text("underlined text"))
	want := `*underlined text*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleUnderlineAndUnderlineText(t *testing.T) {
	got := md.Underline(md.Text("underlined text"))
	want := md.UnderlineText("underlined text")

	if got != want {
		t.Errorf(errorMessage(got.String(), want.String()))
	}
}

func TestUnderlineBoldItalic(t *testing.T) {
	got := md.Underline(md.Bold(md.Italic(md.Text("text"))))
	want := `__*_text_*__`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestUnderlineItalic(t *testing.T) {
	got := md.Underline(md.Italic(md.Text("text")))
	want := "___text___"

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestItalicUnderline(t *testing.T) {
	got := md.Italic(md.Underline(md.Text("text")))
	want := "___text___"

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleStrikethrough(t *testing.T) {
	got := md.Strikethrough(md.Text("strikethrough"))
	want := `~strikethrough~`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleStrikethroughAndStrikethroughText(t *testing.T) {
	got := md.Strikethrough(md.Text("strikethrough"))
	want := md.StrikethroughText("strikethrough")

	if got != want {
		t.Errorf(errorMessage(got.String(), want.String()))
	}
}

func TestSimpleSpoiler(t *testing.T) {
	got := md.Spoiler(md.Text("spoiler"))
	want := `||spoiler||`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleSpoilerAndSpoilerText(t *testing.T) {
	got := md.Spoiler(md.Text("spoiler text"))
	want := md.SpoilerText("spoiler text")

	if got != want {
		t.Errorf(errorMessage(got.String(), want.String()))
	}
}

func TestNested(t *testing.T) {
	got := md.Bold(
		md.Text("bold "),
		md.Italic(md.Text("italic bold "),
			md.Strikethrough(md.Text("italic bold strikethrough "),
				md.Spoiler(md.Text("italic bold strikethrough spoiler")),
			),
			md.Text(" "),
			md.Underline(md.Text("underline italic bold")),
		),
		md.Text(" bold"),
	)
	want := `*bold _italic bold ~italic bold strikethrough ||italic bold strikethrough spoiler||~ __underline italic bold___ bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleInlineUrl(t *testing.T) {
	got := md.InlineURL("telegram", "api.telegram.org")
	want := `[telegram](api.telegram.org)`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}
func TestSimpleInlineFixWidth(t *testing.T) {
	got := md.InlineFixWidth("telegram")
	want := "`telegram`"

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleMentionUser(t *testing.T) {
	got := md.MentionUser("an_awsome_user")
	want := `@an_awsome_user`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSimpleInlineMentionUser(t *testing.T) {
	got := md.InlineMentionUser("an awsome user", "123456789")
	want := `[an awsome user](tg://user?id=123456789)`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestHashtag(t *testing.T) {
	got := md.Hashtag("hashtag")
	want := `\#hashtag`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestHashtagWithSeveralWordsSeparatedBySpace(t *testing.T) {
	got := md.Hashtag("serveral words")
	want := `\#serveral\_words`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestHashtagWithDash(t *testing.T) {
	got := md.Hashtag("serveral-words")
	want := `\#serveral\_words`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestCombine(t *testing.T) {
	got := md.Combine(
		md.Hashtag("serveral words"),
		md.InlineURL("link", "golang.org"),
		md.BoldText("bold"),
	)
	want := `\#serveral\_words[link](golang.org)*bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestCombineSpace(t *testing.T) {
	got := md.CombineWithSpace(
		md.Hashtag("serveral words"),
		md.InlineURL("link", "golang.org"),
		md.BoldText("bold"),
	)
	want := `\#serveral\_words [link](golang.org) *bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestCombineWithNewLine(t *testing.T) {
	got := md.CombineWithNewLine(
		md.Hashtag("serveral words"),
		md.InlineURL("link", "golang.org"),
		md.BoldText("bold"),
	)
	want := `\#serveral\_words\\n[link](golang.org)\\n*bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestSpace(t *testing.T) {
	got := md.Combine(
		md.Hashtag("serveral words"),
		md.Space(),
		md.InlineURL("link", "golang.org"),
		md.Space(),
		md.BoldText("bold"),
	)
	want := `\#serveral\_words [link](golang.org) *bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestNewLine(t *testing.T) {
	got := md.Combine(
		md.Hashtag("serveral words"),
		md.NewLine(),
		md.InlineURL("link", "golang.org"),
		md.NewLine(),
		md.BoldText("bold"),
	)
	want := `\#serveral\_words\\n[link](golang.org)\\n*bold*`

	if !got.Equals(want) {
		t.Errorf(errorMessage(got.String(), want))
	}
}

func TestGetString(t *testing.T) {
	got := md.BoldText("bold")
	want := `*bold*`

	if got.String() != want {
		t.Errorf(errorMessage(got.String(), want))
	}
}
