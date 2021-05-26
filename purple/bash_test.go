package purple

import "testing"

func Test_BashExec(t *testing.T) {
	test := TestExecutionSpec{Name: "bash", Command: "ls"}

	b := BashExecutor{Exec: test}
	err := b.Run()
	if err != nil {
		t.Fail()
	}
}
