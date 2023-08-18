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
	
		curr_stc_end_index := strings.Index(src, curr_stc) + len(curr_stc)
		next_stc_start_index := strings.Index(src, next_stc)
	
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

func recursive_rewrite(content string) {
	
	var iter int = 0
	var fake_percentage float64

	t_start := time.Now()

	for { 

		resp, err_check := zerogpt.Check(content)
		if err_check != nil {
			fmt.Printf("Error: %s\n", err_check)
			os.Exit(1)
		}
	
		// get the array of highlighted sentences
		var sentences []string = (*resp).Data.Sentences
		fake_percentage = (*resp).Data.FakePercentage
	
		prose, num := group(content, sentences)
		
		if num == 0 {
			break	
		}
					
		// rewrite the given sentence using openai
		rewrt,err_rewrite_sentence := openai.Rewrite(prose)
		if err_rewrite_sentence != nil {
			fmt.Printf("Error: %s\n", err_rewrite_sentence)	
			os.Exit(1)
		}
		
		// replace the initial sentence with the current sentence
		content = strings.Replace(content, prose, rewrt, 1)
	
	
		fmt.Printf("#%d hs -> %d. Rewrite %d sentences\n", iter, len(sentences), num)
	
		iter += 1
	}
	
	elapsed := time.Since(t_start)
	fmt.Printf("\n------ rewritten content (%f seconds) (fake: %f) ------\n\n%s\n", elapsed.Seconds(),fake_percentage, content)
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


