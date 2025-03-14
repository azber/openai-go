// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package openai

import (
	"github.com/azber/openai-go/option"
)

// ChatService contains methods and other services that help with interacting with
// the openai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options     []option.RequestOption
	Completions *ChatCompletionService
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r *ChatService) {
	r = &ChatService{}
	r.Options = opts
	r.Completions = NewChatCompletionService(opts...)
	return
}

type ChatModel = string

const (
	ChatModelO3Mini                          ChatModel = "o3-mini"
	ChatModelO3Mini2025_01_31                ChatModel = "o3-mini-2025-01-31"
	ChatModelO1                              ChatModel = "o1"
	ChatModelO1_2024_12_17                   ChatModel = "o1-2024-12-17"
	ChatModelO1Preview                       ChatModel = "o1-preview"
	ChatModelO1Preview2024_09_12             ChatModel = "o1-preview-2024-09-12"
	ChatModelO1Mini                          ChatModel = "o1-mini"
	ChatModelO1Mini2024_09_12                ChatModel = "o1-mini-2024-09-12"
	ChatModelGPT4o                           ChatModel = "gpt-4o"
	ChatModelGPT4o2024_11_20                 ChatModel = "gpt-4o-2024-11-20"
	ChatModelGPT4o2024_08_06                 ChatModel = "gpt-4o-2024-08-06"
	ChatModelGPT4o2024_05_13                 ChatModel = "gpt-4o-2024-05-13"
	ChatModelGPT4oAudioPreview               ChatModel = "gpt-4o-audio-preview"
	ChatModelGPT4oAudioPreview2024_10_01     ChatModel = "gpt-4o-audio-preview-2024-10-01"
	ChatModelGPT4oAudioPreview2024_12_17     ChatModel = "gpt-4o-audio-preview-2024-12-17"
	ChatModelGPT4oMiniAudioPreview           ChatModel = "gpt-4o-mini-audio-preview"
	ChatModelGPT4oMiniAudioPreview2024_12_17 ChatModel = "gpt-4o-mini-audio-preview-2024-12-17"
	ChatModelChatgpt4oLatest                 ChatModel = "chatgpt-4o-latest"
	ChatModelGPT4oMini                       ChatModel = "gpt-4o-mini"
	ChatModelGPT4oMini2024_07_18             ChatModel = "gpt-4o-mini-2024-07-18"
	ChatModelGPT4Turbo                       ChatModel = "gpt-4-turbo"
	ChatModelGPT4Turbo2024_04_09             ChatModel = "gpt-4-turbo-2024-04-09"
	ChatModelGPT4_0125Preview                ChatModel = "gpt-4-0125-preview"
	ChatModelGPT4TurboPreview                ChatModel = "gpt-4-turbo-preview"
	ChatModelGPT4_1106Preview                ChatModel = "gpt-4-1106-preview"
	ChatModelGPT4VisionPreview               ChatModel = "gpt-4-vision-preview"
	ChatModelGPT4                            ChatModel = "gpt-4"
	ChatModelGPT4_0314                       ChatModel = "gpt-4-0314"
	ChatModelGPT4_0613                       ChatModel = "gpt-4-0613"
	ChatModelGPT4_32k                        ChatModel = "gpt-4-32k"
	ChatModelGPT4_32k0314                    ChatModel = "gpt-4-32k-0314"
	ChatModelGPT4_32k0613                    ChatModel = "gpt-4-32k-0613"
	ChatModelGPT3_5Turbo                     ChatModel = "gpt-3.5-turbo"
	ChatModelGPT3_5Turbo16k                  ChatModel = "gpt-3.5-turbo-16k"
	ChatModelGPT3_5Turbo0301                 ChatModel = "gpt-3.5-turbo-0301"
	ChatModelGPT3_5Turbo0613                 ChatModel = "gpt-3.5-turbo-0613"
	ChatModelGPT3_5Turbo1106                 ChatModel = "gpt-3.5-turbo-1106"
	ChatModelGPT3_5Turbo0125                 ChatModel = "gpt-3.5-turbo-0125"
	ChatModelGPT3_5Turbo16k0613              ChatModel = "gpt-3.5-turbo-16k-0613"
)
