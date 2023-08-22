package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/maxasm/recurt/token"
)

const URL = "https://api.openai.com/v1/chat/completions"

func chat(sys_prompt string, user_prompt string) (string, error) {

	client := &http.Client{}
	req, err_req := http.NewRequest("POST", URL, nil)
	if err_req != nil {
		return "", err_req
	}

	// Add the required Headers
	req.Header.Add("Authorization", "Bearer "+os.Getenv("OPEN_AI_API_ACCESS_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	// sample request
	api_req := OPEN_AI_API_REQUEST{
		Model: GPT3,
		Messages: []OPEN_AI_API_MESSAGE{
			{Role: "system", Content: sys_prompt},
			{Role: "user", Content: user_prompt},
		},
	}

	// encode request to json
	json_str, err_json_str := json.Marshal(api_req)
	if err_json_str != nil {
		return "", err_json_str
	}

	// req.Body -> io.ReadCloser
	str_read := bytes.NewBuffer(json_str)
	str_read_closer := io.NopCloser(str_read)

	req.Body = str_read_closer

	resp, err_resp := client.Do(req)
	if err_resp != nil {
		return "", err_resp
	}

	data, err_data := io.ReadAll(resp.Body)
	if err_data != nil {
		return "", err_data
	}

	// encode the responce
	api_resp := &OPEN_AI_API_RESPONSE{}
	err_unmarshal := json.Unmarshal(data, api_resp)
	if err_unmarshal != nil {
		return "", err_unmarshal
	}

	open_ai_resp := api_resp.Choices[0].Message.Content
	return open_ai_resp, nil
}

func read_file(fname string) (string, error) {
	f, err_open_file := os.Open(fname)
	if err_open_file != nil {
		return "", err_open_file
	}

	data, err_read_data := io.ReadAll(f)
	if err_read_data != nil {
		return "", err_read_data
	}

	return string(data), nil
}

func Rewrite(sentence string, gpt_tokens *int64) (string, error) {
	// the prompt to use when rewriting the sentence.
	rewrite_prompt := `Rewrite the following sentence(s) by:
- Write in the style of an advanced University graduate student. 
- Use academic synonyms, and high-level scholarly phrases and wording.
`

	// create user prompt
	user_prompt := fmt.Sprintf("Sentence(s):\n%s", sentence)

	resp,err_resp := chat(rewrite_prompt, user_prompt)
	
	if err_resp != nil {
		return "", err_resp	
	}
	
	full_text_tokens := user_prompt + rewrite_prompt + resp 
	n_tokens, err_count_tokens :=  token.Count(full_text_tokens)
	
	if err_count_tokens != nil {
		return "", err_count_tokens	
	}
	
	(*gpt_tokens) += n_tokens 

	return resp, nil
}
