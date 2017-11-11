package dbg

import (
	"fmt"
	"testing"
)

func TestTrace(t *testing.T) {
	if err := TraceStart("test.trace"); err != nil {
		fmt.Println("failed to start tracing")
		t.Error(err)
		return
	}

	n := 0
	for i := 0; i < 100000; i++ {
		n++
	}
	fmt.Println(n)

	if err := TraceOpenBrowser(); err != nil {
		t.Error(err)
	}
}
