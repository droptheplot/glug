package glug

import (
	"github.com/fatih/color"
	"log"
	"reflect"
	"runtime"
)

func logRequest(method string, path string) {
	color.Set(color.FgCyan)
	log.Printf("GET %s", path)
	color.Unset()
}

func logPlug(plug Plug) {
	log.Printf("Plug: %s", runtime.FuncForPC(reflect.ValueOf(plug).Pointer()).Name())
}

func logHalt(plug Plug) {
	color.Set(color.FgYellow)
	log.Printf("Halt: %s", runtime.FuncForPC(reflect.ValueOf(plug).Pointer()).Name())
	color.Unset()
}
