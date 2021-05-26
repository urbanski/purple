package purple

import "testing"

func Test_osExec(t *testing.T) {
	_, err := osExec("uname -a", false)
	if err != nil {
		t.Fail()
	}
}
