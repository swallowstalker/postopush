package forwarder_test

import (
	"testing"

	"github.com/swallowstalker/postopush/go/forwarder"
)

type TestCase struct {
	Token          string
	Message        string
	ChatID         int64
	ExpectedResult error
}

func TestForward_PartialEnv(t *testing.T) {

	testCases := []TestCase{
		TestCase{
			Token:          "",
			Message:        "A message",
			ChatID:         0,
			ExpectedResult: forwarder.ErrIncompleteEnv,
		},
		TestCase{
			Token:          "token",
			Message:        "",
			ChatID:         0,
			ExpectedResult: forwarder.ErrIncompleteEnv,
		},
	}

	for _, testCase := range testCases {
		err := forwarder.Forward(testCase.Token, testCase.Message, testCase.ChatID)
		if err != testCase.ExpectedResult {
			t.Errorf("partial env test failed: %v; token: %s, message: %s, chat_id: %d", err, testCase.Token, testCase.Message, testCase.ChatID)
		}
	}
}
