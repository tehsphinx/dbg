package dbg

import (
	"os"
	"os/exec"
	"runtime/trace"
)

var (
	file string
)

func TraceStart(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = trace.Start(f)
	if err == nil {
		Blue(fileName)
		file = fileName
	}
	return err
}

func TraceOpenBrowser() error {
	trace.Stop()
	cmd := exec.Command("go", "tool", "trace", file)
	return cmd.Run()
}
