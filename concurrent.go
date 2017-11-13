package dbg

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGID returns the id of the current goroutine. Only use for debugging purpose!
func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
