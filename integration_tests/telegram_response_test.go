// +build integration

package telegrammarkdown_test

type telegramResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Result      result `json:"result"`
}
type from struct {
	ID        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}
type chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
type entities struct {
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Type     string `json:"type"`
	Language string `json:"language"`
}
type result struct {
	MessageID int        `json:"message_id"`
	From      from       `json:"from"`
	Chat      chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []entities `json:"entities"`
}
