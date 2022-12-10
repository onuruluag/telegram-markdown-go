//go:build integration
// +build integration

// go test -v -tags=integration ./...

// for vs code, in settings.json
// "go.testFlags": [
//   "-v",
//   "-tags=integration"
// ],

package telegrammarkdown_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	md "github.com/onuruluag/telegram-markdown-go"
)

var (
	token  string
	chatId string
)

func init() {
	if env, ok := os.LookupEnv("TOKEN"); ok {
		token = env
	}
	if env, ok := os.LookupEnv("CHAT_ID"); ok {
		chatId = env
	}
}

func sendMessage(msg string) (*telegramResponse, bool) {
	if token == "" || chatId == "" {
		return nil, false
	}

	//TODO: add rate limiter
	apiEndpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	params := url.Values{}
	params.Set("chat_id", chatId)
	params.Set("text", msg)
	params.Set("parse_mode", "MarkdownV2")
	params.Set("disable_notification", "true")

	req, err := http.NewRequest(http.MethodPost, apiEndpoint, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, false
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()

	var result *telegramResponse
	json.NewDecoder(resp.Body).Decode(&result)

	if !result.Ok {
		log.Printf("Error: %d\n%s\n", result.ErrorCode, result.Description)
	}

	return result, resp.StatusCode == http.StatusOK
}

func TestSendNested(t *testing.T) {
	msg := md.Bold(
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

	result, ok := sendMessage(msg.String())
	if !ok || len(result.Result.Entities) != 13 {
		t.Error("failed")
	}
}

func TestSendHashtag(t *testing.T) {
	result, ok := sendMessage(md.Hashtag("tash tag").String())

	if !ok || result.Result.Entities[0].Type != "hashtag" {
		t.Error("failed")
	}
}

func TestSendHashtagWithDash(t *testing.T) {
	result, ok := sendMessage(md.Hashtag("tash-tag").String())

	if !ok || result.Result.Entities[0].Type != "hashtag" {
		t.Error("failed")
	}
}

func TestSendPython(t *testing.T) {
	code := `import math
		print('text')
	`
	result, ok := sendMessage(md.CodeBlock("python", code).String())

	if !ok || result.Result.Entities[0].Language != "python" {
		t.Error("failed")
	}
}

func TestSendGolang(t *testing.T) {
	code := `func main() {
		fmt.Printf("%s", "text")
	}`

	result, ok := sendMessage(md.CodeBlock("go", code).String())

	if !ok || result.Result.Entities[0].Language != "go" {
		t.Error("failed")
	}
}

func TestSendCpp(t *testing.T) {
	code := `int main() {
		std::cout<<"text"<<std::endl;
	}`

	result, ok := sendMessage(md.CodeBlock("c++", code).String())

	if !ok || result.Result.Entities[0].Language != "c++" {
		t.Error("failed")
	}
}

func TestSendBlockTable(t *testing.T) {
	table := md.Table{CodeBlock: true, Separator: "|"}
	table.AddColumns(
		md.Column{Width: 40, Align: md.Left},
		md.Column{Width: 40},
	)
	table.AddRow("a", "b")
	table.SetHeader("c1", "c2")

	result, ok := sendMessage(table.String())

	if !ok || result.Result.Entities[0].Type != "pre" {
		t.Error("failed")
	}
}
