package main

import (
	"fmt"
	"io"
	"strings"
	"os"
	"github.com/maxasm/recurt/zerogpt"
	"github.com/maxasm/recurt/openai"
)

func recursive_rewrite(content string) {
	
	var iter int = 0
	
	for { 
		
		resp, err_check := zerogpt.Check(content)
		if err_check != nil {
			fmt.Printf("Error: %s\n", err_check)
			os.Exit(1)
		}
	
		// get the array of highlighted sentences
		var sentences []string = (*resp).Data.Sentences
		
		// check if there are any sentences detected as AI	
		if len(sentences) == 0 {
			break
		}	
	
		var curr_sentence string = sentences[0]
		
		// rewrite the given sentence using openai
		rewrt,err_rewrite_sentence := openai.Rewrite(curr_sentence)
		if err_rewrite_sentence != nil {
			fmt.Printf("Error: %s\n", err_rewrite_sentence)	
			os.Exit(1)
		}
		
		// replace the initial sentence with the current sentence
		content = strings.Replace(content, curr_sentence, rewrt, 1)
	
		fmt.Printf("#%d hs -> %d\n", iter, len(sentences))
	
		iter += 1
	}
	
	fmt.Printf("\n------ rewritten content ------\n\n%s\n", content)
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
	
	return string(data), nil
}

func main() {
	text, err_read_file := read_file("input.txt")
	if err_read_file != nil {
		fmt.Printf("Error: %s\n", err_read_file)	
		os.Exit(1)
	} 

	recursive_rewrite(text)
}


