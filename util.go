// Copyright (C) 2014 Constantin Schomburg <me@cschomburg.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gozelle

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func structToValues(i interface{}, values url.Values) {
	val := reflect.ValueOf(i).Elem()
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		var v string
		switch f.Interface().(type) {
		case int:
			if f.Int() != 0 {
				v = strconv.FormatInt(f.Int(), 10)
			}
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case bool:
			if f.Bool() {
				v = "1"
			}
		case string:
			v = f.String()
		}
		if v != "" {
			values.Set(strings.ToLower(typ.Field(i).Name), v)
		}
	}
}
