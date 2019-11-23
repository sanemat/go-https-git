package httpsgit

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	outStreem := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	err := Run([]string{}, outStreem, errStream)
	if err != nil {
		t.Errorf("error should be nil, but: %s", err)
	}
}
