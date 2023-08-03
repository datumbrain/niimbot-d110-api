package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

func runPythonScript(cmdAndArgs ...string) error {
	cmd := exec.Command("python3", cmdAndArgs...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	return cmd.Wait()
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
