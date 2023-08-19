package zerogpt 

import (
	"net/http"	
	"io"
	"bytes"
	"encoding/json"
)

const access_token = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE0MCIsInJvbGUiOiIzIiwic2FsYXRhX2VuZ2luZSI6IjIuNyIsImNvc3RfcGVyX3Rob3VzYW5kIjoiMC4wMzQiLCJudW1iZXJfb2ZfY2hhcmFjdGVycyI6IjUwMDAwIiwibnVtYmVyX29mX2ZpbGVzIjoiNDAuMCIsImV4cCI6MTcyMTk5NDE3N30.AolsrSLz8xF4EkQfSGh5fCh7XHPnhWCksBmEDemB-vlRjbKJw5th7iefrRbwXlPBUYuZzByzQK-ka4NAFyzLM5YOrPy-k9xWC3FsvLea6G5uK_IADnqOBkni0Yywj7ofPqaQFf4qNl-8D1t0KlfLBV467XnXstyjpQkxLrtVxsA"
const api_end_point = "https://api.zerogpt.com/api/detect/detectText"

func upload(txt []byte) ([]byte, error) {
	// set up the data
	upload_data := bytes.NewBuffer(txt)
	// create a new http request
	req, err_create_req := http.NewRequest("POST", api_end_point, upload_data)
	if err_create_req != nil {
		return []byte{}, err_create_req 	
	}
	
	// set the required headers
	req.Header.Set("Authorization", "Bearer "+ access_token)
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://www.zerogpt.com/")
	req.Header.Set("Origin", "https://www.zerogpt.com")

	req.Header.Set("Content-Length", string(upload_data.Len()))

	req.Header.Set("Sec-Ch-Ua", `"Not/A)Brand";v="99", "Google Chrome";v="115", "Chromium";v="115"`)
	req.Header.Set("Sec-Ch-Ua-Platform", "Linux")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	resp, err_resp := http.DefaultClient.Do(req)	
	if err_resp != nil {
		return []byte{}, err_resp 	
	}
	
	/**
	fmt.Printf("\n ---- headers ----\n\n")
	
	for k, v := range resp.Header {
		fmt.Printf("%s:%s\n", k,v)	
	}
	
	fmt.Printf("\n")
	**/
	
	defer resp.Body.Close()

	data, err_read_data := io.ReadAll(resp.Body)
	if err_read_data != nil {
		return nil, err_read_data	
	}
	
	return data, nil	
}

type PayLoad struct {
	InputText string `json:"input_text"`	
}

type APIData struct {
	Sentences []string `json:"h"`
	FakePercentage float64 `json:"fakePercentage"`
	IsHuman float64 `json:"isHuman"`
}

type APIResponse struct {
	Data APIData `json:"data"` 
}

func Check(text string) (*APIResponse,error) {

	test_data := PayLoad{InputText: text}	
	test_as_json, err_marshal := json.Marshal(test_data)
 	if err_marshal != nil {
		return nil, err_marshal
	}	

	resp, err_upload := upload(test_as_json)
	if err_upload != nil {
		return nil, err_upload
	}
	
	api_resp := APIResponse{}
	
	err_unmarshal := json.Unmarshal(resp, &api_resp)
	if err_unmarshal != nil {
		return nil, err_unmarshal	
	}

	/**
	resp_json, err_resp_json := json.MarshalIndent(api_resp, "", " ")	
	if err_resp_json != nil {
		fmt.Printf("Error: %s\n", err_resp_json)	
		os.Exit(1)
	}
	**/	
	
	return &api_resp, nil
	
}
