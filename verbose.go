package verbose

import (
  "fmt"
  "runtime"
  "path/filepath"
  "strings"
  "regexp"
  "os"
  "sync"
)

var mainPath string
var mutex = &sync.Mutex{}

func init () {
  _,mainPath,_,_ = runtime.Caller(2)
  mainPath = filepath.Dir(mainPath)
  // fmt.Printf("---- %s\n", mainPath) // a kind of problem here with go test
}

type RuntimeVerbose struct {
  MainFile string
  IsWindows bool
  Gopath string
  VerboseEnv string
  VerboseRegs []*regexp.Regexp
  hasInit bool
}

func (r *RuntimeVerbose) InitFromMain() {
  mutex.Lock()
  if r.hasInit==false {
    r.Init(mainPath)
  }
  mutex.Unlock()
}

func (r *RuntimeVerbose) Init(from string) {
  if r.hasInit {
    return
  }
  r.hasInit     = true
  r.IsWindows   = strings.Index(from, "\\") > -1
  r.Gopath      = r.determineGoPath(from)
  r.MainFile    = from
  r.VerboseEnv  = os.Getenv("VERBOSE")
  r.initVerboseRegexp()
}

func (r *RuntimeVerbose) initVerboseRegexp() {
  items := strings.Split(r.VerboseEnv, ",")

  for _, v := range items {
    if v!="" {
      v := regexp.QuoteMeta(v)
      v = strings.Replace(v, "\\*", ".+", -1)
      r.VerboseRegs = append(r.VerboseRegs, regexp.MustCompile(v))
    }
  }
}

func (r *RuntimeVerbose) determineGoPath(from string) string {
  pathItems := strings.Split(from, "/")
  if r.IsWindows {
    pathItems = strings.Split(from, "\\")
  }

  var index int
  for i, v := range pathItems {
    if v=="src" {
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

func (r *RuntimeVerbose) DeterminePackage(from string) string {
  var pkg string
  if r.isVendored(from) {
    pkg = r.pathFromVendor(from)
  } else {
    pkg = r.pathFromGopath(from)
  }
  return pkg
}

func (r *RuntimeVerbose) pathFromVendor (from string) string {
  pathItems := strings.Split(from, "/")
  if r.IsWindows {
    pathItems = strings.Split(from, "\\")
  }

  index := -1
  for i, v := range pathItems {
    if v=="vendor" {
      index = i
    }
  }
  strlen := 0
  for i, v := range pathItems {
    if i<=index {
      strlen+=len(v)+1
    }
  }

  return filepath.Dir(from[strlen:])
}

func (r *RuntimeVerbose) pathFromGopath (path string) string {
  return filepath.Dir(path[len(r.Gopath)+len("/src/"):])
}

func (r *RuntimeVerbose) isVendored (from string) bool {
  pathItems := strings.Split(from, "/")
  if r.IsWindows {
    pathItems = strings.Split(from, "\\")
  }

  index := -1
  for i, v := range pathItems {
    if v=="vendor" {
      index = i
    }
  }

  return index>-1
}

func (r *RuntimeVerbose) isEnabled (from string) bool {
  isEnabled := false
  for _, reg := range r.VerboseRegs {
    if reg.MatchString(from) {
      isEnabled = true
      break
    }
  }
  return isEnabled
}

var runtimeV *RuntimeVerbose

var loggerMutex = &sync.Mutex{}

func Auto () *Logger {
  loggerMutex.Lock()
  if runtimeV==nil {
    runtimeV = &RuntimeVerbose{}
  }
  runtimeV.InitFromMain()
  _,file,_,_ := runtime.Caller(1)
  name := runtimeV.DeterminePackage(file)
  loggerMutex.Unlock()
  return From(name)
}
func From (name string) *Logger {
  loggerMutex.Lock()
  if runtimeV==nil {
    runtimeV = &RuntimeVerbose{}
  }
  runtimeV.InitFromMain()
  loggerMutex.Unlock()
  return &Logger{
    name:     name,
    enabled:  runtimeV.isEnabled(name),
  }
}

type Logger struct {
  name string
  enabled bool
}

func (l *Logger) Println (msg string) {
  if l.enabled {
    fmt.Println(msg)
  }
}
