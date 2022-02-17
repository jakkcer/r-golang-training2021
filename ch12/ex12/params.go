package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Check func(v interface{}) error

// Unpackは、req内のHTTPリクエストパラメータから
// ptrが指す構造体のフィールドに値を移しかえます。
func Unpack(req *http.Request, ptr interface{}, checks map[string]Check) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 実効的な名前をキーとするフィールドのマップを構築する。
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // 構造体変数
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		checkName := tag.Get("check")
		if checkName != "" {
			if check, ok := checks[checkName]; ok {
				if err := check(v.Field(i).Interface()); err != nil {
					return err
				}
			}
		}
		fields[name] = v.Field(i)
	}

	// リクエスト中の個々のパラメータに対する構造体のフィールドを更新
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 認識されなかったHTTPパラメータを無視
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
