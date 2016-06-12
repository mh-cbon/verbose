package verbose

import (
	// "fmt"
	"os"
	// "strings"
	"testing"
)

func TestRuntimeVerboseMainFile(t *testing.T) {
	somePackagePath := "whatever/src/your/pkg/main.go"

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if len(r.MainFile) == 0 {
		t.Errorf("MainFile must be set\n")
	}
	if r.hasInit != true {
		t.Errorf("Must have init\n")
	}
	if r.IsWindows != false {
		t.Errorf("Expected IsWindows=false, got IsWindows=true\n")
	}
	if r.Gopath != "whatever" {
		t.Errorf("Expected %q got %q\n", "whatever", r.Gopath)
	}
	if r.VerboseEnv != "" {
		t.Errorf("Expected %q got %q\n", "", r.VerboseEnv)
	}
	if r.MainFile != somePackagePath {
		t.Errorf("Expected %q got %q\n", somePackagePath, r.MainFile)
	}
}

func TestEnvVarReading(t *testing.T) {

	verboseEnv := "a,b"
	somePackagePath := "whatever/src/your/pkg/main.go"
	os.Setenv("VERBOSE", verboseEnv)

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if r.VerboseEnv != "a,b" {
		t.Errorf("Expected %q got %q\n", verboseEnv, r.VerboseEnv)
	}

	if len(r.VerboseRegs) != 2 {
		t.Errorf("Expected %q got %q\n", 2, len(r.VerboseRegs))
	} else {
		if r.VerboseRegs[0].String() != "a" {
			t.Errorf("Expected %q got %q\n", "a", r.VerboseRegs[0].String())
		}
	}
}

func TestRegexpQuoting(t *testing.T) {

	verboseEnv := "a?c"
	somePackagePath := "whatever/src/your/pkg/main.go"
	os.Setenv("VERBOSE", verboseEnv)

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if len(r.VerboseRegs) != 1 {
		t.Errorf("Expected %q got %q\n", 1, len(r.VerboseRegs))
	} else {
		if r.VerboseRegs[0].String() != "a\\?c" {
			t.Errorf("Expected %q got %q\n", "a\\?c", r.VerboseRegs[0].String())
		}
	}
}

func TestRegexpReplacement(t *testing.T) {

	verboseEnv := "a*"
	somePackagePath := "whatever/src/your/pkg/main.go"
	os.Setenv("VERBOSE", verboseEnv)

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if len(r.VerboseRegs) != 1 {
		t.Errorf("Expected %q got %q\n", 1, len(r.VerboseRegs))
	} else {
		if r.VerboseRegs[0].String() != "a.+" {
			t.Errorf("Expected %q got %q\n", "a.+", r.VerboseRegs[0].String())
		}
	}
}

func TestWindowsDetection(t *testing.T) {

	somePackagePath := "c:\\whatever\\src\\your\\pkg\\main.go"

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if r.IsWindows != true {
		t.Errorf("Expected IsWindows=true, got IsWindows=false\n")
	}
	if r.Gopath != "c:\\whatever" {
		t.Errorf("Expected %q got %q\n", "c:\\whatever", r.Gopath)
	}
	if r.MainFile != somePackagePath {
		t.Errorf("Expected %q got %q\n", somePackagePath, r.MainFile)
	}
}

func TestIsVendored(t *testing.T) {

	gopathHosted := "/what/ever/src/package/main.go"
	vendoredHosted := "/what/ever/src/package/vendor/some/lib.go"

	r := &runtimeVerbose{}

	if r.isVendored(gopathHosted) != false {
		t.Errorf("Expected isVendored=false, got isVendored=true\n")
	}

	if r.isVendored(vendoredHosted) != true {
		t.Errorf("Expected isVendored=true, got isVendored=false\n")
	}
}

func TestPathFromVendor(t *testing.T) {
	vendoredHosted := "/what/ever/src/package/vendor/some/lib.go"

	r := &runtimeVerbose{}

	p := r.pathFromVendor(vendoredHosted)
	if p != "some" {
		t.Errorf("Expected %q, got %q\n", "some", p)
	}
}

func TestPathFromGopath(t *testing.T) {
	vendoredHosted := "/what/ever/src/package/some/main.go"

	r := &runtimeVerbose{}
	r.Init("/what/ever/src")

	p := r.pathFromGopath(vendoredHosted)
	if p != "package/some" {
		t.Errorf("Expected %q, got %q\n", "package/some", p)
	}
}

func TestIsEnabled(t *testing.T) {
	verboseEnv := "a*"
	somePackagePath := "whatever/src/your/pkg/main.go"
	os.Setenv("VERBOSE", verboseEnv)

	r := &runtimeVerbose{}
	r.Init(somePackagePath)

	if r.isEnabled("asomething") != true {
		t.Errorf("Expected %t got %t\n", true, false)
	}

	if r.isEnabled("b") != false {
		t.Errorf("Expected %t got %t\n", false, true)
	}
}

func TestLoggerAuto(t *testing.T) {
	os.Setenv("VERBOSE", "")
	logger := Auto()
	if logger.Enabled {
		t.Errorf("Expected %t got %t\n", false, true)
	}
	if logger.Name != "github.com/mh-cbon/verbose" {
		t.Errorf("Expected %q got %q\n", "github.com/mh-cbon/verbose", logger.Name)
	}
}

func TestLoggerFrom(t *testing.T) {
	os.Setenv("VERBOSE", "")
	logger := From("whatever")
	if logger.Enabled {
		t.Errorf("Expected %t got %t\n", false, true)
	}
	if logger.Name != "whatever" {
		t.Errorf("Expected %q got %q\n", "whatever", logger.Name)
	}
}
