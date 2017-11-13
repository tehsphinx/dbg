package dbg

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strconv"
	"sync/atomic"
)

var (
	fileCPU            string
	executableFileName string
	fileNumber         uint64
)

type ProfileType string

const (
	ProfileWebChart ProfileType = "web"
	ProfileWebList              = "weblist=."
	ProfileList                 = "list=."
)

func SetExecutableName(executable string) {
	executableFileName = executable
}

func FlameGraph(seconds int, pprofURL ...string) error {
	url := "http://localhost:6060/debug/pprof/profile"
	if len(pprofURL) != 0 {
		url = pprofURL[0]
	}
	return urlFlameGraph(seconds, "torch.svg", url)
}

func CPUProfileStart() error {
	fileName := getFileName("profile_%d.cpu")
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

// CPUProfileFlameGraph creates a go-torch cpu flamegraph and opens it in the browser
// Needs gotorch to be installed with flamegraph scripts in PATH
// TODO: make open cmd plattform independent
func CPUProfileFlameGraph() error {
	pprof.StopCPUProfile()
	return fileFlameGraph(fileCPU, executableFileName)
}

func CPUProfile(profile ProfileType) error {
	pprof.StopCPUProfile()
	return createReport(getProfileString(profile), fileCPU, executableFileName)
}

func GoroutineProfile(profile ProfileType) error {
	fileName := getFileName("profile_%d.goroutine")
	if err := writeProfile("goroutine", fileName); err != nil {
		return err
	}
	return createReport(getProfileString(profile), fileName, executableFileName)
}

func HeapProfile(profile ProfileType) error {
	fileName := getFileName("profile_%d.heap")
	if err := writeProfile("heap", fileName); err != nil {
		return err
	}
	return createReport(getProfileString(profile), fileName, executableFileName)
}

func ThreadcreateProfile(profile ProfileType) error {
	fileName := getFileName("profile_%d.thread")
	if err := writeProfile("threadcreate", fileName); err != nil {
		return err
	}
	return createReport(getProfileString(profile), fileName, executableFileName)
}

func BlockProfileStart() {
	runtime.SetBlockProfileRate(1)
}

func BlockProfile(profile ProfileType) error {
	fileName := getFileName("profile_%d.block")
	if err := writeProfile("block", fileName); err != nil {
		return err
	}
	return createReport(getProfileString(profile), fileName, executableFileName)
}

func MutexProfileStart() {
	runtime.SetMutexProfileFraction(1)
}

func MutexProfile(profile ProfileType) error {
	fileName := getFileName("profile_%d.mutex")
	if err := writeProfile("mutex", fileName); err != nil {
		return err
	}
	return createReport(getProfileString(profile), fileName, executableFileName)
}

func writeProfile(profileName, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	runtime.GC()
	if err := pprof.Lookup(profileName).WriteTo(f, 0); err != nil {
		return fmt.Errorf("could not write %s profile: %s", profileName, err.Error())
	}
	return f.Close()
}

func urlFlameGraph(seconds int, fileName, url string) error {
	svgFile := fmt.Sprintf("%s.svg", fileName)
	cmd := exec.Command("go-torch", "--seconds", strconv.Itoa(seconds), "-f", svgFile, url)
	return openFlameGraph(cmd, svgFile)
}

func fileFlameGraph(fileName, executable string) error {
	svgFile := fmt.Sprintf("%s.svg", fileName)
	cmd := exec.Command("go-torch", "-f", svgFile, executable, fileName)
	return openFlameGraph(cmd, svgFile)
}

func openFlameGraph(cmd *exec.Cmd, svgFile string) error {
	if b, err := cmd.CombinedOutput(); err != nil {
		Debug(string(b))
		Debug("error running go-torch:", err)
		return err
	}
	cmd = exec.Command("open", svgFile)
	if b, err := cmd.CombinedOutput(); err != nil {
		Debug(string(b))
		Debug("error opening", svgFile, ":", err)
		return err
	}
	return nil
}

func createReport(report, fileName, executable string) error {
	paramReport := fmt.Sprintf("-%s", report)
	cmd := exec.Command("go", "tool", "pprof", paramReport, executable, fileName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		Debug("error running pprof:", err)
		return err
	}
	return nil
}

func getFileName(filePattern string) string {
	atomic.AddUint64(&fileNumber, 1)
	return fmt.Sprintf(filePattern, atomic.LoadUint64(&fileNumber))
}

func getProfileString(profile ProfileType) string {
	switch profile {
	default:
		return string(profile)
	}
}
