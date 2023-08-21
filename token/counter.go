package token

import (
	"os/exec"
	"strconv"
	"fmt"
	"io"
)

func Count(str string) (int64, error) {
	
	// execute node js and run ../counter/app.js
	node_cmd := exec.Command("node", "./counter/app.js")
	
	stdin_pipe, err_get_pipe := node_cmd.StdinPipe()
	if err_get_pipe != nil {
		return 0, err_get_pipe	
	}
	
	// write the text to the pipe
	io.WriteString(stdin_pipe, str)
	stdin_pipe.Close()
	
	output, err_run_command := node_cmd.CombinedOutput()
	if err_run_command != nil {
		return 0, err_run_command	
	}
	
	output_as_string := string(output)
	output_as_int,err_conv := strconv.ParseInt(output_as_string, 10, 64)
	if err_conv != nil {
		return 0, err_conv	
	}
	
	return output_as_int, nil
}
