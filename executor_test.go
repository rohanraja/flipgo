package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestExecutingCommand(t *testing.T) {

	out, _ := exec.Command("./flipgo", "-queues='catcrawl'").Output()
	fmt.Println(string(out))

}
