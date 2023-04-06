package fcm

import (
	"fmt"
	"testing"
)

var apiKey = ""

var pushTokens = []string{"aa", "bb", "cc", "dd", "eee", "ff", "gg", "hh", "ii", "jj", "qq", "ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo", "zz", "xx", "cc", "vv", "bb", "123"}

func init() {

}

func TestFCM(t *testing.T) {
	pushClient, err := NewFromAPIKey(apiKey)
	if err != nil {
		t.Error(err)
	}

	pushClient.pushTokens = pushTokens
	pushClient.fcmMessage.Notification.Title = "title"
	pushClient.fcmMessage.Notification.Body = "body"

	resp := pushClient.Send()
	fmt.Printf("Success: %d\n", resp.Success)
	fmt.Printf("Failure: %d\n", resp.Failure)
}

func TestTokensSlice(t *testing.T) {

	tokens := []string{"aa", "bb", "cc", "dd", "eee", "ff", "gg", "hh", "ii", "jj", "qq", "ww", "ee", "rr", "tt", "yy", "uu", "ii", "oo", "zz", "xx", "cc", "vv", "PA91bFpYvVHnbsZXl"}

	count := 5
	var j int
	for i := 0; i < len(tokens); i += count {
		j += count
		if j > len(tokens) {
			j = len(tokens)
		}
		fmt.Printf("i : %d j : %d\n", i, j)
		fmt.Println(tokens[i:j])
	}

}
