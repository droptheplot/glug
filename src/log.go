package glug

import (
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"net/url"
	"reflect"
	"runtime"
)

func logRequest(method string, path string) {
	color.Set(color.FgCyan, color.Bold)
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

func logParams(params url.Values) {
	color.Set(color.FgCyan)
	pretty_params, _ := json.Marshal(params)
	log.Println(string(pretty_params))
	color.Unset()
}
