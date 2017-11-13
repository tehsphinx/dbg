package dbg

import (
	"fmt"
	"sync"
	"testing"
)

func TestCPUProfileFlameGraph(t *testing.T) {
	if err := CPUProfileStart(); err != nil {
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

	if err := CPUProfileFlameGraph(); err != nil {
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

func TestCPUProfile(t *testing.T) {
	var (
		mu    sync.Mutex
		items = make(map[int]struct{})
		wg    = &sync.WaitGroup{}
	)

	SetExecutableName("server")
	if err := CPUProfileStart(); err != nil {
		fmt.Println("failed to start profiling")
		t.Error(err)
		return
	}

	wg.Add(20 * 1000)
	for i := 0; i < 20*1000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	if err := CPUProfile(ProfileList); err != nil {
		t.Error(err)
	}
}

func TestMutexProfile(t *testing.T) {
	var (
		mu    sync.Mutex
		items = make(map[int]struct{})
		wg    = &sync.WaitGroup{}
	)

	SetExecutableName("server")
	MutexProfileStart()

	wg.Add(20 * 1000)
	for i := 0; i < 20*1000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	MutexProfile(ProfileList)
}

func TestBlockProfile(t *testing.T) {
	var (
		mu    sync.Mutex
		items = make(map[int]struct{})
		wg    = &sync.WaitGroup{}
	)

	SetExecutableName("server")
	BlockProfileStart()

	wg.Add(20 * 1000)
	for i := 0; i < 20*1000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	BlockProfile(ProfileList)
}

func TestHeapProfile(t *testing.T) {
	var (
		mu    sync.Mutex
		items = make(map[int]struct{})
		wg    = &sync.WaitGroup{}
	)

	SetExecutableName("server")

	wg.Add(20 * 1000)
	for i := 0; i < 20*1000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	HeapProfile(ProfileList)
}

func TestThreadcreateProfile(t *testing.T) {
	var (
		wg = &sync.WaitGroup{}
	)

	SetExecutableName("server")

	wg.Add(20 * 1000)
	for i := 0; i < 20*1000; i++ {
		go func(i int, wg *sync.WaitGroup) {
			wg.Done()
		}(i, wg)
	}
	wg.Wait()

	ThreadcreateProfile(ProfileList)
}
