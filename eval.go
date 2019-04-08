package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func EvalSnippet(s *Snippet) error {
	tmpfile, err := ioutil.TempFile("", "isitpython-*.py")
	if err != nil {
		return err
	}

	log.Println("Created new file", tmpfile.Name())

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(s.Body)); err != nil {
		return err
	}

	if err := tmpfile.Close(); err != nil {
		return err
	}
	log.Println("Executing this python", s.Body)
	cmd := exec.Command("python", tmpfile.Name())
	out, err := cmd.CombinedOutput()

	fmt.Println("this is the output", string(out))

	s.Output = string(out)
	s.Error = err
	s.IsValidPython = cmd.ProcessState.Success()

	return nil
}
