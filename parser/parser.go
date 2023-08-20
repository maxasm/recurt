package parser

import (
	"strings"
	"unicode"
	"bytes"
)

	
type Token struct {
	Chars []rune
	IsSentence bool
}

func (tk Token) String() string {
	return string(tk.Chars)
}

func parseChar(src []rune, tokens *[]Token, index *int) {

	curr_char := src[(*index)]	

	tk := Token{}	
	tk.Chars = []rune{curr_char}
	tk.IsSentence = false
	
	(*tokens) = append((*tokens), tk)
	(*index) = (*index)+1
} 

func parseSentence(src []rune, tokens *[]Token, index *int) {

	var curr_index int = (*index) 
	var start_index int = curr_index 
	
	for {
	
		if curr_index+1 == len(src) {
			curr_index += 1	
			break
		}		
	
		// the current character in the array	
		var curr_char rune = src[curr_index]	
		var next_char rune = src[curr_index+1]

		// check if the next character is a new line
		if next_char == '\r' || next_char == '\n' {
			curr_index += 1
			break	
		}

		if curr_char == '.' {
			if next_char == ' ' {
				curr_index += 1
				break
			}
		} 	
	
		curr_index += 1
	}
	
	tk := Token{}
	tk.Chars = src[start_index:curr_index]
	
	// check if Chars just contains spaces
	chars_str := string(tk.Chars)
	if len(strings.Trim(chars_str, " \r\n\t.")) != 0 {
		tk.IsSentence = true	
	}

	(*tokens) = append((*tokens), tk)
	(*index) = curr_index
}

func Parse(src []rune) []Token {

	var curr_index int = 0	
	
	var tokens []Token = make([]Token, 0)

	for {
		
		// break if you get to the end of the input
		if curr_index == len(src) {
			break	
		}

		// the current character
		var curr_char rune = src[curr_index] 	
	
		// the sentence must start with a letter or number
		if unicode.IsLetter(curr_char) || unicode.IsNumber(curr_char) || curr_char == ' '{
			parseSentence(src, &tokens, &curr_index)	
			continue
		}
	
		// if curr_index is not a letter or a character, parser the character alone
		parseChar(src, &tokens, &curr_index)
	}
	
	return tokens
}

type Paragraph struct {
	Text string
	Editable bool
}

func parseSentenceParagraph(tokens []Token, paragraphs *[]Paragraph, index *int) {
	
	var start_index int = (*index)
	var end_index int = start_index 

	var next_token Token
	
	for {
		
		if end_index+1 == len(tokens) {
			end_index += 1	
			break
		}	

		next_token = tokens[end_index+1]
		if !next_token.IsSentence {
			end_index += 1		
			break
		}
		
		end_index += 1
	}
	
	paragraph := Paragraph{}
	paragraph.Editable = true
		
	// set the text of the paragraph
	var out bytes.Buffer = bytes.Buffer{}

	for _, tk := range tokens[start_index:end_index] {
		out.WriteString(tk.String())			
	}
	
	paragraph.Text = out.String()
	
	// append the paragraph
	(*paragraphs) = append((*paragraphs), paragraph)
	
	// update index
	(*index) = end_index
}

func parseCharParagraph(tokens []Token, paragraphs *[]Paragraph, index *int) {
	
	paragraph := Paragraph{}
	paragraph.Editable = false
	paragraph.Text = tokens[(*index)].String()	
	
	// append the new paragraph
	(*paragraphs) = append((*paragraphs), paragraph)
	
	// update index
	(*index) = (*index) + 1
}

func ParseParagraphs(tokens []Token) []Paragraph {

	var curr_index int = 0	

	var paragraphs []Paragraph = make([]Paragraph, 0)

	for {
		
		if curr_index == len(tokens) {
			break
		}	
		
		curr_token := tokens[curr_index]	
		if curr_token.IsSentence {
			parseSentenceParagraph(tokens, &paragraphs, &curr_index)	
			continue
		}
	
		parseCharParagraph(tokens, &paragraphs, &curr_index)
	}
	
	return paragraphs
}
