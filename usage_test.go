// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai_test

import (
	"context"
	"os"
	"testing"

	"github.com/azber/openai-go"
	"github.com/azber/openai-go/internal/testutil"
	"github.com/azber/openai-go/option"
)

func TestUsage(t *testing.T) {
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
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{openai.ChatCompletionUserMessageParam{
			Role:    openai.F(openai.ChatCompletionUserMessageParamRoleUser),
			Content: openai.F([]openai.ChatCompletionContentPartUnionParam{openai.ChatCompletionContentPartTextParam{Text: openai.F("text"), Type: openai.F(openai.ChatCompletionContentPartTextTypeText)}}),
		}}),
		Model: openai.F(openai.ChatModelO3Mini),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", chatCompletion)
}
