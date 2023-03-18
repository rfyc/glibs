package flags

import (
	"flag"
	"fmt"
	"github.com/rfyc/glibs/conv"
	"github.com/rfyc/glibs/php"
	"github.com/rfyc/glibs/structs"
	"reflect"
	"strings"
)

var (
	ptrs       []interface{}
	maps       = map[string]interface{}{}
	Format     = `flag:"{fname1},{fname2},{fname3};{default};{fusage}"`
	allowTypes = []reflect.Kind{
		reflect.Bool,
		reflect.String,
		reflect.Int,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint64,
		reflect.Float64,
	}
)
var nstr *string

func flagVar(fs []string, ftype reflect.Kind) {

	var val, usage string

	if len(fs) == 2 {
		val = fs[1]
	} else if len(fs) > 2 {
		usage = fs[2]
	}

	for _, name := range strings.Split(fs[0], ",") {
		switch ftype {
		case reflect.Bool:
			maps[name] = flag.Bool(name, conv.Bool(val), usage)
		case reflect.String:
			nstr = flag.String(name, val, usage)
			maps[name] = nstr
		case reflect.Int:
			maps[name] = flag.Int(name, conv.Int(val), usage)
		case reflect.Int64:
			maps[name] = flag.Int64(name, conv.Int64(val), usage)
		case reflect.Uint:
			maps[name] = flag.Uint(name, conv.Uint(val), usage)
		case reflect.Uint64:
			maps[name] = flag.Uint64(name, conv.Uint64(val), usage)
		case reflect.Float64:
			maps[name] = flag.Float64(name, conv.Float64(val), usage)
		}
	}
}

func StructVar(ptr interface{}) {

	if !structs.IsPtr(ptr) {
		panic("flags struct var not ptr")
	}

	var (
		exist bool
		valOf = structs.ValueOf(ptr)
	)

	for k := 0; k < valOf.NumField(); k++ {
		if field := valOf.Type().Field(k); field.Tag.Get("flag") != "" {
			if fs := strings.Split(field.Tag.Get("flag"), ";"); len(fs[0]) > 0 {
				if !php.InArray(field.Type.Kind(), allowTypes) {
					panic("field typie must bool,string,int,int64,uint,uint64,float64")
				}
				exist = true
				flagVar(fs, field.Type.Kind())
			}
		}
	}

	if exist {
		ptrs = append(ptrs, ptr)
	}
}

func parseVar(field reflect.Value, val interface{}) {
	switch field.Kind() {
	case reflect.Bool:
		if v, ok := val.(*bool); ok {
			field.SetBool(*v)
		}
	case reflect.String:
		if v, ok := val.(*string); ok {
			field.SetString(*v)
		}
	case reflect.Int:
		if v, ok := val.(*int); ok {
			field.SetInt(int64(*v))
		}
	case reflect.Int64:
		if v, ok := val.(*int64); ok {
			field.SetInt(*v)
		}
	case reflect.Uint:
		if v, ok := val.(*uint); ok {
			field.SetUint(uint64(*v))
		}
	case reflect.Uint64:
		if v, ok := val.(*uint64); ok {
			field.SetUint(*v)
		}
	case reflect.Float64:
		if v, ok := val.(*float64); ok {
			field.SetFloat(*v)
		}
	}
}

func Parse() {

	flag.Parse()
	fmt.Println("nstr:", *nstr)
	for _, ptr := range ptrs {
		var valOf = structs.ValueOf(ptr)
		for k := 0; k < valOf.NumField(); k++ {
			if field := valOf.Type().Field(k); field.Tag.Get("flag") != "" && field.Anonymous {
				if fs := strings.Split(field.Tag.Get("flag"), ";"); len(fs[0]) > 0 {
					if php.InArray(field.Type.Kind(), allowTypes) {
						for _, name := range strings.Split(fs[0], ",") {
							if val, ok := maps[name]; ok {
								parseVar(valOf.Field(k), val)
							}
						}
					}
				}
			}
		}
	}
}
