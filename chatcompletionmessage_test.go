// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/azber/openai-go"
	"github.com/azber/openai-go/internal/testutil"
	"github.com/azber/openai-go/option"
)

func TestChatCompletionMessageListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Chat.Completions.Messages.List(
		context.TODO(),
		"completion_id",
		openai.ChatCompletionMessageListParams{
			After: openai.F("after"),
			Limit: openai.F(int64(0)),
			Order: openai.F(openai.ChatCompletionMessageListParamsOrderAsc),
		},
	)
	if err != nil {
		var apierr *openai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
