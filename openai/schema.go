package openai 

const (
	GPT3 = "gpt-3.5-turbo"
)

type OPEN_AI_API_MESSAGE struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OPEN_AI_API_REQUEST struct {
	Model    string                `json:"model"`
	Messages []OPEN_AI_API_MESSAGE `json:"messages"`
}

type OPEN_AI_API_CHOICE struct {
	Index        uint                `json:"index"`
	Message      OPEN_AI_API_MESSAGE `json:"message"`
	FinishReason string              `json:"finish_reason"`
}

type OPEN_AI_API_USAGE struct {
	PromptTokens     uint `json:"propt_tokens"`
	CompletionTokens uint `json:"completion_tokens"`
	TotalTokens      uint `json:"total_tokens"`
}

type OPEN_AI_API_RESPONSE struct {
	Id      string               `json:"id"`
	Object  string               `json:"object"`
	Created uint                 `json:"created"`
	Usage   OPEN_AI_API_USAGE    `json:"usage"`
	Choices []OPEN_AI_API_CHOICE `json:"choices"`
}
