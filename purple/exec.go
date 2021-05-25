package purple

import (
	"bytes"
	"github.com/google/shlex"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExecTest(inputFile string) {

	spec,err := readTestSpec(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Infof("Running test suite '%s'", spec.DisplayName)

	// loop over the tests
	for _, test := range spec.Tests {
		log.Infof("Test: %s", test.Name)
		switch (test.Execution.Name) {
			case "bash":
				bashexec := BashExecutor{Exec: test.Execution}
				err := bashexec.Run()
				if err != nil {
					log.Error(err)
				}
			default:
				log.Errorf("Executor '%s' not defined.", test.Execution.Name)
		}

	}


	return
}

// readInputFile reads a test spec for parsing
func readInputFile(inputFile string) ([]byte, error) {
	// ensure the file exists on the filesystem
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return nil, err
	}

	return ioutil.ReadFile(inputFile)
}

func readTestSpec(inputFile string) (*TestSuiteSpec, error) {
	log.Debug("Reading ", inputFile)

	filedata, err := readInputFile(inputFile)
	if err != nil {
		return nil, err
	}

	spec, err := ParseART(string(filedata))
	if err != nil {
		return nil, err
	}

	return spec, nil
}

func osExec(command string, useStdout bool) (*exec.Cmd, error) {
	cmdTuple, err := shlex.Split(command)
	if err != nil {
		return nil, err
	}

	args := cmdTuple[1:]

	cmd := exec.Command(cmdTuple[0], args...)
	if useStdout {
		var errBuf, outBuf bytes.Buffer
		cmd.Stderr = io.MultiWriter(os.Stderr, &errBuf)
		cmd.Stdout = io.MultiWriter(os.Stdout, &outBuf)
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	return cmd, err
}