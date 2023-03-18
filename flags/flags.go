package flags

import (
	"flag"
	"github.com/rfyc/glibs/conv"
	"github.com/rfyc/glibs/php"
	"github.com/rfyc/glibs/structs"
	"reflect"
	"strings"
)

var (
	ptrs       []interface{}
	maps       = map[string]interface{}{}
	format     = `flag:"{fname1},{fname2},{fname3};{default};{fusage}"`
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

func flagVar(fs []string, ftype reflect.Kind) {

	var val, usage string

	if len(fs) == 2 {
		val = fs[1]
	} else if len(fs) > 2 {
		usage = fs[2]
	}
	fnames := strings.Split(fs[0], ",")
	switch ftype {
	case reflect.Bool:
		var v = new(bool)
		for _, name := range fnames {
			flag.BoolVar(v, name, conv.Bool(val), usage)
			maps[name] = v
		}
	case reflect.String:
		var v = new(string)
		for _, name := range fnames {
			flag.StringVar(v, name, val, usage)
			maps[name] = v
		}
	case reflect.Int, reflect.Int64:
		var v = new(int64)
		for _, name := range fnames {
			flag.Int64Var(v, name, conv.Int64(val), usage)
			maps[name] = v
		}
	case reflect.Uint, reflect.Uint64:
		var v = new(uint64)
		for _, name := range fnames {
			flag.Uint64Var(v, name, conv.Uint64(val), usage)
			maps[name] = v
		}
	case reflect.Float64:
		var v = new(float64)
		for _, name := range fnames {
			flag.Float64Var(v, name, conv.Float64(val), usage)
			maps[name] = v
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
	case reflect.Int, reflect.Int64:
		if v, ok := val.(*int64); ok {
			field.SetInt(*v)
		}
	case reflect.Uint, reflect.Uint64:
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

	for _, ptr := range ptrs {
		var valOf = structs.ValueOf(ptr)
		for k := 0; k < valOf.NumField(); k++ {
			if field := valOf.Type().Field(k); field.Tag.Get("flag") != "" {
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
