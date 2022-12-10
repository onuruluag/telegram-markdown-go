package telegrammarkdown_test

import "fmt"

func errorMessage(got, want string) string {
	return fmt.Sprintf("\ngot: %s\nwant: %s\n", got, want)
}
