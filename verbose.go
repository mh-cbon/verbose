// A library to display verbose messages per package.
package verbose

import (
	// "fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"

  "github.com/mh-cbon/verbose/printer"
  "github.com/mh-cbon/verbose/color"
)

var mainPath string

// find path of the entry point
// set default printer
func init() {
  SetPrinter(printer.LogPrinter{})
	_, mainPath, _, _ = runtime.Caller(1)
	mainPath = filepath.Dir(mainPath)
}

// stores information about runtime
type runtimeVerbose struct {
	MainFile    string
	IsWindows   bool
	Gopath      string
	VerboseEnv  string
	VerboseRegs []*regexp.Regexp
	hasInit     bool
}

var initRuntimeVMutex = &sync.Mutex{}
// initiliaze from the entry point determmined at runtime
func (r *runtimeVerbose) InitFromMain() {
	initRuntimeVMutex.Lock()
	if r.hasInit == false {
		r.Init(mainPath)
	}
	initRuntimeVMutex.Unlock()
}

// initialize from a pre defined entry point
func (r *runtimeVerbose) Init(from string) {
	if r.hasInit {
		return
	}
	r.hasInit = true
	r.IsWindows = strings.Index(from, "\\") > -1
	r.Gopath = r.determineGoPath(from)
	r.MainFile = from
	r.VerboseEnv = os.Getenv("VERBOSE")
	r.initVerboseRegexp()
}

// transforms the value of VERBOSE env var into a
// slice of regexp
// VERBOSE is a string of comma separated pattern package
// when a pattern package contains *, it is replaced by .+
// examples:
// VERBOSE=* => .+
// VERBOSE=this/pkg* => this/pkg.+
func (r *runtimeVerbose) initVerboseRegexp() {
	items := strings.Split(r.VerboseEnv, ",")

	for _, v := range items {
		if v != "" {
			v := regexp.QuoteMeta(v)
			v = strings.Replace(v, "\\*", ".+", -1)
			r.VerboseRegs = append(r.VerboseRegs, regexp.MustCompile(v))
		}
	}
}

// Given a string path return the part preceding the first src/
// example:
// from= /some/path/src/whatever => /some/path
func (r *runtimeVerbose) determineGoPath(from string) string {
	pathItems := strings.Split(from, "/")
	if r.IsWindows {
		pathItems = strings.Split(from, "\\")
	}

	var index int
	for i, v := range pathItems {
		if v == "src" {
			index = i
			break
		}
	}

	gopath := strings.Join(pathItems[0:index], "/")
	if r.IsWindows {
		gopath = strings.Join(pathItems[0:index], "\\")
	}

	return gopath
}

// Given a string path return the package name
// When it s a package hosted into GOPATH,
// it returns the part following the first src/
// When it s a package hosteed in vendor/
// it returns the part following the last vendor/
// example:
// from= /some/whatever/src/package/ => package
// from= /some/whatever/src/package/vendor/sub => sub
func (r *runtimeVerbose) DeterminePackage(from string) string {
	var pkg string
	if r.isVendored(from) {
		pkg = r.pathFromVendor(from)
	} else {
		pkg = r.pathFromGopath(from)
	}
	return pkg
}

// Given a string path to a vendored package,
// returns the package name
// example:
// from= /some/whatever/src/package/vendor/sub => sub
func (r *runtimeVerbose) pathFromVendor(from string) string {
	pathItems := strings.Split(from, "/")
	if r.IsWindows {
		pathItems = strings.Split(from, "\\")
	}

	index := -1
	for i, v := range pathItems {
		if v == "vendor" {
			index = i
		}
	}
	strlen := 0
	for i, v := range pathItems {
		if i <= index {
			strlen += len(v) + 1
		}
	}

	return filepath.Dir(from[strlen:])
}

// Given a string path to a GOPATH hosted package,
// returns the package name
// example:
// from= /some/whatever/src/package/ => package
func (r *runtimeVerbose) pathFromGopath(path string) string {
	return filepath.Dir(path[len(r.Gopath)+len("/src/"):])
}

func (r *runtimeVerbose) isVendored(from string) bool {
	pathItems := strings.Split(from, "/")
	if r.IsWindows {
		pathItems = strings.Split(from, "\\")
	}

	index := -1
	for i, v := range pathItems {
		if v == "vendor" {
			index = i
		}
	}

	return index > -1
}

// Given VERBOSE and a package name,
// tells if the log are enabled for this package
func (r *runtimeVerbose) isEnabled(from string) bool {
	isEnabled := false
	for _, reg := range r.VerboseRegs {
		if reg.MatchString(from) {
			isEnabled = true
			break
		}
	}
	return isEnabled
}

var runtimeV *runtimeVerbose
var runtimeVMutex = &sync.Mutex{}

func initRuntimeVerbose() {
	runtimeVMutex.Lock()
	if runtimeV == nil {
		runtimeV = &runtimeVerbose{}
	}
	runtimeV.InitFromMain()
	runtimeVMutex.Unlock()
}

// Configure current Printer
func SetPrinter(p printer.Printer) {
	printer.SetPrinter(p)
}

// Automatically determine package name
// return a new Logger instance
func Auto() *printer.Logger {
	initRuntimeVerbose()
	_, file, _, _ := runtime.Caller(1)
	name := runtimeV.DeterminePackage(file)
	return From(name)
}

// Return a new printer.Logger with given name
func From(name string) *printer.Logger {
	initRuntimeVerbose()
	return &printer.Logger{
		Name:     name,
		Enabled:  runtimeV.isEnabled(name),
    Color:    color.PickColor(),
	}
}
