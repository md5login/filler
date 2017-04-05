package filler

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName         = "fill"
	defaultsTagName = "defaults"
	ignoreTag       = "-"
	emptyTag        = ""
)

var fillers []Filler

// Filler instance
type Filler struct {
	// Tag is the prefix inside fill tag ie. "fill:mytag"
	Tag string
	// Fn function to call - helps us to fill the gaps
	Fn func(obj interface{}) (interface{}, error)
}

// RegFiller - register new filler into []fillers
func RegFiller(f Filler) {
	fillers = append(fillers, f)
}

// Fill - fill the object with all the current fillers
func Fill(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		log.Panic("panic at [md5login/filler] : obj kind passed to Fill should be Ptr")
	}
	v := reflect.TypeOf(obj).Elem()
	s := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		currentField := v.Field(i)
		tag := currentField.Tag.Get(tagName)
		if tag == emptyTag || tag == ignoreTag {
			continue
		}
		t, elm := parseTag(tag)
		for _, filter := range fillers {
			var elmValue interface{}
			if filter.Tag == t {
				if elm != "" {
					elmValue = s.FieldByName(elm).Interface()
				}
				res, err := filter.Fn(elmValue)
				if err != nil {
					continue
				}
				if s.FieldByName(currentField.Name).CanSet() {
					s.FieldByName(currentField.Name).Set(reflect.ValueOf(res))
				}
			}
		}
	}
}

// Defaults - fill the object with its default values
func Defaults(obj interface{}) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		log.Panic("panic at [yaronsumel/filler] : obj kind passed to Defaults should be Ptr")
	}
	v := reflect.TypeOf(obj).Elem()
	s := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		currentField := v.Field(i)
		tag := currentField.Tag.Get(defaultsTagName)
		if tag == emptyTag || tag == ignoreTag {
			continue
		}
		if s.FieldByName(currentField.Name).CanSet() {
			switch s.FieldByName(currentField.Name).Kind() {
			case reflect.Int:
				i, _ := strconv.Atoi(tag)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(i))
			case reflect.Int8:
				i, _ := strconv.ParseInt(tag, 10, 8)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(i))
			case reflect.Int16:
				i, _ := strconv.ParseInt(tag, 10, 16)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(i))
			case reflect.Int32:
				i, _ := strconv.ParseInt(tag, 10, 32)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(i))
			case reflect.Int64:
				i, _ := strconv.ParseInt(tag, 10, 64)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(i))
			case reflect.Uint:
				u, _ := strconv.ParseUint(tag, 10, 32)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(u))
			case reflect.Uint8:
				u, _ := strconv.ParseUint(tag, 10, 8)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(u))
			case reflect.Uint16:
				u, _ := strconv.ParseUint(tag, 10, 16)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(u))
			case reflect.Uint32:
				u, _ := strconv.ParseUint(tag, 10, 32)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(u))
			case reflect.Uint64:
				u, _ := strconv.ParseUint(tag, 10, 64)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(u))
			case reflect.String:
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(tag))
			case reflect.Bool:
				b, _ := strconv.ParseBool(tag)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(b))
			case reflect.Float32:
				f, _ := strconv.ParseFloat(tag, 32)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(f))
			case reflect.Float64:
				f, _ := strconv.ParseFloat(tag, 64)
				s.FieldByName(currentField.Name).Set(reflect.ValueOf(f))
			}
		}
	}
}

// parseTag split the string by ":" and return two strings
func parseTag(tag string) (string, string) {
	x := strings.Split(tag, ":")
	if len(x) != 2 {
		return x[0], ""
	}
	return x[0], x[1]
}
