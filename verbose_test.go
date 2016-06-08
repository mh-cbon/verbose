package verbose

import (
  // "fmt"
  "os"
  "strings"
  "testing"
)

func TestRuntimeVerboseMainFile(t *testing.T) {
	r := &RuntimeVerbose{}
  r.InitFromMain()

  if len(r.MainFile)==0 {
    t.Errorf("MainFile must be set\n")
  }
  gopath := os.Getenv("GOPATH") + "/src/github.com/mh-cbon/verbose"
  if strings.Index(r.MainFile, gopath)!=1 {
    t.Errorf("MainFile should look like _" + gopath + ", got " + r.MainFile + "\n")
  }
}
