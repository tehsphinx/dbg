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
		file = fileName
	}
	return err
}

func TraceOpenBrowser() error {
	trace.Stop()
	cmd := exec.Command("go", "tool", "trace", file)

	if b, err := cmd.CombinedOutput(); err != nil {
		Debug(string(b))
		Debug("error running go tool trace:", err)
		return err
	}
	return nil
}
