package dbg

import (
	"os"
	"os/exec"
	"runtime/pprof"
)

var (
	fileCPU string
)

func CPUProfileStart(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	err = pprof.StartCPUProfile(f)
	if err == nil {
		fileCPU = fileName
	}
	return err
}

// CPUProfileOpenBrowser creates a go-torch cpu flamegraph and opens it in the browser
// Needs gotorch to be installed with flamegraph scripts in PATH
// TODO: make open cmd plattform independent
func CPUProfileOpenBrowser() error {
	pprof.StopCPUProfile()
	cmd := exec.Command("go-torch", fileCPU)
	if b, err := cmd.CombinedOutput(); err != nil {
		Debug(string(b))
		Debug("error running go-torch:", err)
		return err
	}
	cmd = exec.Command("open", "torch.svg")
	if b, err := cmd.CombinedOutput(); err != nil {
		Debug(string(b))
		Debug("error opening torch.svg:", err)
		return err
	}
	return nil
}
