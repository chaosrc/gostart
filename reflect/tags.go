package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func search(res http.ResponseWriter, req *http.Request) {
	var data struct {
		Labels     []string `http:"1"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	data.MaxResults = 10
	if err := Unpack(req, &data); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(res, "Search: %+v\n", data)
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)

	v := reflect.ValueOf(ptr).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(field.Name)
		}
		fields[name] = v.Field(i)
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return err
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func populate(rv reflect.Value, v string) error {
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(v)
	case reflect.Int:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		rv.SetInt(int64(i))

	case reflect.Bool:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		rv.SetBool(b)
	default:
		return fmt.Errorf("Unsuport kind %s", rv.Type())
	}
	return nil
}

func tagServer() {
	http.HandleFunc("/", search)
	http.ListenAndServe(":8888", nil)
}

func printMethods(val interface{}) {
	ref := reflect.ValueOf(val)
	fmt.Printf("type %s\n", ref.Type())

	for i := 0; i < ref.NumMethod(); i++ {
		method := ref.Method(i)
		fmt.Printf("func (%s) %s%s\n", ref.Type(),
			ref.Type().Method(i).Name,
			strings.TrimPrefix(method.Type().String(), "func"))
	}
}
