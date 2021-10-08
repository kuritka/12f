/*
Copyright 2021 The k8gb Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Generated by GoLic, for more details see: https://github.com/AbsaOSS/golic
*/
package env

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

type field struct {
	env        env
	fieldName  string
	fieldType  *reflect.Type
	fieldValue *reflect.Value
	public     bool
}

// contains raw info about string tag field. e.g: default=hello,
type strTag struct {
	value  string
	exists bool
}

type env struct {
	value   string
	name    string
	tagName string
	def     strTag
	req     strTag
	present bool
}

type meta map[string]field

const (
	tagEnv = "env"
)

// Bind binds environment variables into structure
// ✅ repeated values
// ✅ Bind two fields by one envvar
// ✅ nested structres
// ✅ anonymous structures
// ✅ binding to private fields
// ✅ default values
// ✅ required values
// ✅ env prefixes
// ✅ slices
// ❎ maps
// ❎ datetimes
// ❌ keys 🔑
func Bind(s interface{}) (err error) {
	var meta meta
	if s == nil {
		return fmt.Errorf("invalid argument value (nil)")
	}
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s).Kind()
	if t != reflect.Ptr {
		return fmt.Errorf("argument must be pointer to structure")
	}
	if v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("argument must be pointer to structure")
	}
	meta, err = roll(v.Elem(), v.Elem().Type().Name(), "")
	if err != nil {
		return
	}
	err = bind(meta)
	return
}

// binds meta to structure pointer
func bind(m meta) (err error) {
	for k, v := range m {
		f := reflect.NewAt(v.fieldValue.Type(), unsafe.Pointer(v.fieldValue.UnsafeAddr())).Elem()
		switch v.fieldValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i, intdef int
			if v.env.def.exists {
				intdef, err = strconv.Atoi(v.env.def.value)
				if err != nil {
					err = fmt.Errorf("can't convert default value %s of '%s' to int", v.env.name, v.env.def.value)
					return
				}
			}
			i, err = GetEnvAsIntOrFallback(v.env.name, intdef)
			if err != nil {
				err = fmt.Errorf("can't read %s and parse value %s to int", v.env.name, v.env.def.value)
				return
			}
			f.SetInt(int64(i))
			continue

		case reflect.Bool:
			var b, booldef bool
			if v.env.def.exists {
				booldef, err = strconv.ParseBool(v.env.def.value)
				if err != nil {
					err = fmt.Errorf("can't convert default value %s of '%s' to bool", v.env.name, v.env.def.value)
					return
				}
			}
			b = GetEnvAsBoolOrFallback(v.env.name, booldef)
			f.SetBool(b)
			continue

		case reflect.Float32, reflect.Float64:
			var fl, floatdef float64
			if v.env.def.exists {
				floatdef, err = strconv.ParseFloat(v.env.def.value, 64)
				if err != nil {
					err = fmt.Errorf("can't convert default value %s of '%s' to float64", v.env.name, v.env.def.value)
					return
				}
			}
			fl, err = GetEnvAsFloat64OrFallback(v.env.name, floatdef)
			if err != nil {
				err = fmt.Errorf("can't read %s and parse value %s to float64", v.env.name, v.env.def.value)
				return
			}
			f.SetFloat(fl)
			continue

		case reflect.String:
			var s string
			s = GetEnvAsStringOrFallback(v.env.name, v.env.def.value)
			f.SetString(s)
			continue

		case reflect.Slice:
			var strvalues []string
			if v.env.def.exists {
				envdef := strings.TrimSuffix(strings.TrimPrefix(v.env.def.value, " "), " ")
				envdef = strings.TrimSuffix(strings.TrimPrefix(envdef, "["), "]")
				strvalues = strings.Split(envdef, ",")
				if strvalues[0] == "" {
					strvalues = []string{}
				}
			}
			switch f.Interface().(type) {
			case []string:
				var strslice []string
				strslice = GetEnvAsArrayOfStringsOrFallback(v.env.name, strvalues)
				f.Set(reflect.ValueOf(strslice))
				continue

			case []int, []int8, []int16, []int32, []int64:
				var intslice []int
				sintdef := []int{}
				for _, s := range strvalues {
					var i int
					i, err = strconv.Atoi(strings.Trim(s, " "))
					if err != nil {
						err = fmt.Errorf("can't convert default %s to slice of int", strvalues)
						return
					}
					sintdef = append(sintdef, i)
				}
				intslice, err = GetEnvAsArrayOfIntsOrFallback(v.env.name, sintdef)
				if err != nil {
					err = fmt.Errorf("can't parse %s as slice of int %s", v.env.name, v.env.value)
				}
				f.Set(reflect.ValueOf(intslice))
				continue

			case []float32, []float64:
				var floatslice []float64
				sfloatdef := []float64{}
				for _, s := range strvalues {
					var fl float64
					fl, err = strconv.ParseFloat(strings.Trim(s, " "), 64)
					if err != nil {
						err = fmt.Errorf("can't convert default %s to slice of float64", strvalues)
						return
					}
					sfloatdef = append(sfloatdef, fl)
				}
				floatslice, err = GetEnvAsArrayOfFloat64OrFallback(v.env.name, sfloatdef)
				if err != nil {
					err = fmt.Errorf("can't parse %s as slice of float64 %s", v.env.name, v.env.value)
				}
				f.Set(reflect.ValueOf(floatslice))
				continue

			case []bool:
				var boolslice []bool
				sbooldef := []bool{}
				for _, s := range strvalues {
					var b bool
					b, err = strconv.ParseBool(strings.Trim(s, " "))
					if err != nil {
						err = fmt.Errorf("can't convert default %s to slice of bool", strvalues)
						return
					}
					sbooldef = append(sbooldef, b)
				}
				boolslice, err = GetEnvAsArrayOfBoolOrFallback(v.env.name, sbooldef)
				if err != nil {
					err = fmt.Errorf("can't parse %s as array of ints %s", v.env.name, v.env.value)
				}
				f.Set(reflect.ValueOf(boolslice))
				continue

			default:
				err = fmt.Errorf("unsupported type %s: %s", k, v.fieldValue.Type().Name())
			}
		default:
			err = fmt.Errorf("unsupported type %s: %s", k, v.fieldValue.Type().Name())
			return
		}
	}
	return err
}

// recoursive function builds meta structure
func roll(value reflect.Value, n, prefix string) (m meta, err error) {
	m = meta{}
	for i := 0; i < value.NumField(); i++ {
		var e env
		vf := value.Field(i)
		tf := value.Type().Field(i)
		key := fmt.Sprintf("%s.%s", n, tf.Name)
		tag := tf.Tag.Get(tagEnv)
		if vf.Kind() == reflect.Struct {
			var sm meta
			prefix := strings.TrimPrefix(fmt.Sprintf("%s_%s", prefix, getTagName(tag)), "_")
			sm, err = roll(vf, key, prefix)
			if err != nil {
				return
			}
			for k, v := range sm {
				m[k] = v
			}
			continue
		}
		if tag == "" {
			continue
		}
		if e, err = parseTag(tag, prefix); err != nil {
			return
		}
		if !e.present && e.req.value == "true" {
			err = fmt.Errorf("%s is required", e.name)
			return
		}
		m[key] = field{
			env:        e,
			fieldName:  tf.Name,
			fieldType:  &tf.Type,
			fieldValue: &vf,
			public:     tf.IsExported(),
		}
	}
	return m, err
}

// parseTag, retrieves env info and metadata
func parseTag(tag, prefix string) (e env, err error) {
	var def, req strTag
	var tagName = getTagName(tag)
	req, err = getTagProperty(tag, "require")
	if err != nil {
		return
	}
	def, err = getTagProperty(tag, "default")
	if err != nil {
		return
	}
	envName := getEnvName(tagName, prefix)
	value, exists := os.LookupEnv(envName)
	e = env{
		name:    envName,
		tagName: tagName,
		value:   value,
		req:     req,
		def:     def,
		present: exists,
	}
	return
}

func getEnvName(envName, prefix string) string {
	if prefix != "" {
		return fmt.Sprintf("%s_%s", prefix, envName)
	}
	return envName
}

func getTagName(tag string) string {
	return regexp.MustCompile("[a-zA-Z_]+[a-zA-Z0-9_]*").FindString(tag)
}

// parses value from env tag and returns <tag value, tag value exists, error>
func getTagProperty(tag, t string) (r strTag, err error) {
	const arr = `\[\w*\s*\!*\@*\#*\$*\%*\^*\&*\**\(*\)*\_*\-*\+*\<*\>*\?*\~*\=*\,*\.*\/*\{*\}*\|*\;*\:*\/*\'*\"*\/*\\*`
	const scalar = `\[*\]*\w*\s*\!*\@*\#*\$*\%*\^*\&*\**\(*\)*\_*\-*\+*\<*\>*\?*\~*\=*\.*\/*\{*\}*\|*\;*\:*\/*\'*\"*\/*\\*`
	r = strTag{}
	var findRegex, removeRegex *regexp.Regexp
	//	findRegex, err = regexp.Compile(",\\s*" + t + "\\s*=((\\s*([\\[\\w*\\,*\\.*\\s*\\-*])*\\])|(\\s*\\w*\\.*\\-*)*)")
	findRegex, err = regexp.Compile(",\\s*" + t + "\\s*=((\\s*([" + arr + "])*\\])|(" + scalar + ")*)")
	if err != nil {
		err = fmt.Errorf("ivalid %s", t)
		return
	}
	removeRegex, err = regexp.Compile(",\\s*" + t + "\\s*=\\s*")
	if err != nil {
		err = fmt.Errorf("ivalid %s", t)
		return
	}
	match := findRegex.FindString(tag)
	if match == "" {
		return
	}
	remove := removeRegex.FindString(strings.ToLower(tag))
	r.value = strings.ReplaceAll(match, remove, "")
	r.exists = true
	return
}
