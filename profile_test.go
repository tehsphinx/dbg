package dbg

import (
	"fmt"
	"sync"
	"testing"
)

func TestCPUProfile(t *testing.T) {
	if err := CPUProfileStart("test.cpu"); err != nil {
		fmt.Println("failed to start profiling")
		t.Error(err)
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go count(wg)
	}
	wg.Wait()

	if err := CPUProfileOpenBrowser(); err != nil {
		t.Error(err)
	}
}

func count(wg *sync.WaitGroup) {
	n := 0
	for i := 0; i < 100000000; i++ {
		n++
	}
	wg.Done()
}
