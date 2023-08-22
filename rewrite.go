package main

import (
	"fmt"
	"io"
	"strings"
	"os"
	"bytes"
	"time"
	"github.com/maxasm/recurt/zerogpt"
	"github.com/maxasm/recurt/openai"
	"github.com/maxasm/recurt/parser"
	"github.com/maxasm/recurt/token"
)

func parseStr(str string) bool {
	
	var curr_index int = 0
	
	for {
		
		if curr_index == len(str) {
			break	
		}

		curr_char := str[curr_index]	
		if (curr_char == ' ' || curr_char == '\t') {
			curr_index += 1
			continue
		}
	
		return false
	}

	return true
}

func group(src string, stc []string) (string, int) {

	var curr_stc_index int = 0 

	var curr_stc string
	var next_stc string 
	
	var buffer bytes.Buffer = bytes.Buffer{}
	
	if len(stc) == 0 {
		return "", 0	
	}	
	
	for {
	
		curr_stc = stc[curr_stc_index]
		
		if curr_stc_index+1 == len(stc) {
			buffer.WriteString(curr_stc)				
			break
		}	

		next_stc = stc[curr_stc_index+1]
	
		// get the index of the last character of the current sentence
		curr_stc_end_index := strings.Index(src, curr_stc) + len(curr_stc)
		// get the index of the next sentence. start looking after the last senetence and not from the begining
		next_stc_start_index := strings.Index(src[curr_stc_end_index:], next_stc) + curr_stc_end_index
	
		str_between := src[curr_stc_end_index:next_stc_start_index] 

		// check if the chars in between only contain spaces, tabs or nothing.
		if ok := parseStr(str_between); ok {
			// write the current sentence and the chars in between.
			buffer.WriteString(curr_stc)
			buffer.WriteString(str_between)

			curr_stc_index += 1
			continue
		} 
	
		buffer.WriteString(curr_stc)
		break
	}	
	
	return buffer.String(), curr_stc_index+1 
}

func recursive_rewrite(content string, gpt_tokens *int64) (string,float64) {
	
	var iter int = 0
	var is_human float64

	t_start := time.Now()

	for { 

		resp, err_check := zerogpt.Check(content)
		if err_check != nil {
			panic(err_check)
		}
	
		// get the array of highlighted sentences
		var sentences []string = (*resp).Data.Sentences
		is_human = (*resp).Data.IsHuman
	
		prose, num := group(content, sentences)
		
		if num == 0 {
			break	
		}
					
		// rewrite the given sentence using openai
		rewrt,err_rewrite_sentence := openai.Rewrite(prose, gpt_tokens)
		if err_rewrite_sentence != nil {
			panic(err_rewrite_sentence)	
		}
		
		// replace the initial sentence with the current sentence
		content = strings.Replace(content, prose, rewrt, 1)

		fmt.Printf("#%d hs -> %d. Rewrite %d sentences\n", iter, len(sentences), num)
	
		iter += 1
	}
	
	elapsed := time.Since(t_start)
	fmt.Printf("\n\033[40m------\033[49m rewritten content (%.2f seconds) (human: %.2f) \033[40m------\033[49m\n\n%s\n", elapsed.Seconds(),is_human, content)
    
    return content, is_human
}


func read_file(fname string) (string, error) {
	f, err_open := os.Open(fname)
	if err_open != nil {
		return "", err_open	
	}
	
	data, err_read_data := io.ReadAll(f)
	if err_read_data != nil {
		return "", err_read_data	
	}
	
	// remove the eof char \0
	data = data[:(len(data)-1)]

	return string(data), nil
}


func rewrite_paragraph(paragraph parser.Paragraph, iter int, gpt_tokens *int64) (string, error) {
	var res string = paragraph.Text
	
	for a := 0; a < iter; a++ {
		rewrt, err_rewrt := openai.Rewrite(res, gpt_tokens) 	
		if err_rewrt != nil {
			return "", err_rewrt	
		}
		res = rewrt
	}
	
	return res, nil	
}

func rewrite_paragraphs(paragraphs []parser.Paragraph, iter int, gpt_tokens *int64) (string, error) {
	
	var out bytes.Buffer = bytes.Buffer{}
	
	for a, pr := range paragraphs {
		
		if !pr.Editable {
			out.WriteString(pr.Text)	
			continue
		} 
		
		fmt.Printf("rewriting block %d of %d ...\n", a+1, len(paragraphs))	
		rewrt, err_rewrt := rewrite_paragraph(pr, iter, gpt_tokens)	
		if err_rewrt != nil {
			return "", err_rewrt 
		}
	
		out.WriteString(rewrt)
	}
	
	return out.String(), nil	
}


type RewriteResponse struct {
    Text string `json:"text"`
    Tokens int64 `json:"tokens"`
    Cost float64 `json:"cost"`
    Human float64 `json:"human"`
}

// run takes a string and returns a response containing the ID of the text being processed.
func run(content string) RewriteResponse {

	// recover from any error that may occur due to text parsing and more ...
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("There was an error, retring again ...\n%s\n", err)	
			run(content)
		}	
	}()
	
    // parse sentence and char tokens
	tokens := parser.Parse([]rune(content))	
		
	// counter for the number of gpt tokens used
	var gpt_tokens int64 = 0

    // takes the tokens and groups paragraphs together
	prs := parser.ParseParagraphs(tokens)
	
	resp, err_resp := rewrite_paragraphs(prs, 2, &gpt_tokens) 
	
	if err_resp != nil {
		panic(err_resp)	
	}
	
	fmt.Printf("\n ---- recursive rewrite ----\n")
	resp_content, human_p := recursive_rewrite(resp, &gpt_tokens)
    
	token_cost := token.Cost(gpt_tokens)
	fmt.Printf("\n ---- done rewriting. used %d tokens costing ($%.3f) ----\n", gpt_tokens, token_cost)
    
    // create the response object
    rp := RewriteResponse {
        Text: resp_content,
        Human: human_p, 
        Cost: token_cost,
        Tokens: gpt_tokens,
    }
    
    return rp
}


