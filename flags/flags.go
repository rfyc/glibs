package flags

import (
	"flag"
	"github.com/rfyc/glibs/structs"
)

var maps = map[string]interface{}

func StructVar(ptr interface{}) {
	if !structs.IsPtr(ptr) {
		panic("flags struct var not ptr")
	}
	valOf := structs.ValueOf(ptr)
	for k := 0; k < valOf.NumField(); k++ {
		name := valOf.Type().Field(k).Tag.Get("flag")
		if name != "" {
			var val = flag.String(name, "tt", "")
		}
	}
}

func Parse() {

}
