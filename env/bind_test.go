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
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testee public test structure
type Testee struct {
	ID    int    `env:"TESTEE_ID"`
	Value string `env:"TESTEE_VALUE"`
}

func TestDefaultAndRequireinOppositeOrder(t *testing.T) {
	defer cleanup()
	// arrange
	_ = os.Setenv(tokenID, "")
	_ = os.Setenv(tokenValue, "")
	type token struct {
		ID    int    `env:"TOKEN_ID, require=true, default = 066"`
		Value string `env:"TOKEN_VALUE, default = AAAA, require=true"`
	}
	// act
	tok := &token{ID: 5}
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, 66, tok.ID)
	assert.Equal(t, "AAAA", tok.Value)
}

func TestParseNil(t *testing.T) {
	defer cleanup()
	err := Bind(nil)
	assert.Error(t, err)
}

func TestParseNotPointer(t *testing.T) {
	defer cleanup()
	err := Bind(Testee{})
	assert.Error(t, err)
}

func TestInvalidEnvVar(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID int `env:"GG%%^"`
	}
	// act
	tok := &token{ID: 5}
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, 0, tok.ID)
}

func TestUnsupportedDataType(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID reflect.Type `env:"TOKEN_ID"`
	}
	// act
	tok := &token{}
	err := Bind(tok)
	// assert
	assert.Error(t, err)
}

func TestENVVARIsRequiredError(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID    int    `env:"TOKEN_ID, require=true"`
		Value string `env:"TOKEN_VALUE"`
	}
	// act
	tok := &token{ID: 5}
	err := Bind(tok)
	// assert
	assert.Error(t, err)
	assert.Equal(t, 5, tok.ID)
}

func TestENVVARIsRequiredPass(t *testing.T) {
	defer cleanup()
	// arrange
	const val = "4BMKKDsdfsf5f7="
	_ = os.Setenv(tokenValue, val)
	type token struct {
		Value string `env:"TOKEN_VALUE, require=true"`
	}
	// act
	tok := &token{Value: "AAAA"}
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, val, tok.Value)
}

func TestENVVARIsNotRequiredPass(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID string `env:"TOKEN_ID, require=false"`
	}
	tok := &token{ID: "AAAA"}
	// act
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, "", tok.ID)
}

func TestSetDefaultEmpty(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID          int       `env:"TOKEN_ID, default=0"`
		Value       string    `env:"TOKEN_VALUE, default="`
		Ratio       float64   `env:"TOKEN_RATIO, default=0"`
		Readonly    bool      `env:"TOKEN_READONLY, default=false"`
		Hours       []int     `env:"TOKEN_HOURS, default[]"`
		URLs        []string  `env:"TOKEN_URLS, default=[]"`
		Enabled     []bool    `env:"TOKEN_BOOLS, default=[]"`
		Coordinates []float64 `env:"TOKEN_COORDINATES, default=[]"`
	}
	tok := &token{}
	// act
	err := Bind(tok)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 0, tok.ID)
	assert.Equal(t, "", tok.Value)
	assert.Equal(t, 0., tok.Ratio)
	assert.Equal(t, false, tok.Readonly)
	assert.Equal(t, []int{}, tok.Hours)
	assert.Equal(t, []string{}, tok.URLs)
	assert.Equal(t, []bool{}, tok.Enabled)
	assert.Equal(t, []float64{}, tok.Coordinates)
}

func TestInvalidValue(t *testing.T) {
	defer cleanup()
	_ = os.Setenv(envString, "invalid")
	_ = os.Setenv(envBool, "invalid")
	_ = os.Setenv(envFloat64, "invalid")
	_ = os.Setenv(envInt, "invalid")
	_ = os.Setenv(envStringSlice, "[]")
	_ = os.Setenv(envBoolSlice, "invalid")
	_ = os.Setenv(envFloat64Slice, "invalid")
	_ = os.Setenv(envIntSlice, "invalid")
	// arrange
	type token1 struct {
		envString string `env:"ENV_STRING, default=test"`
	}
	type token2 struct {
		envInt int `env:"ENV_INT, default=22"`
	}
	type token3 struct {
		envBool bool `env:"ENV_BOOL, default=true"`
	}
	type token4 struct {
		envFloat float32 `env:"ENV_FLOAT64, default=22.0"`
	}
	type token5 struct {
		envStringSlice []string `env:"ENV_STRING_SLICE, default=[test,test]"`
	}
	type token6 struct {
		envIntSlice []int `env:"ENV_INT_SLICE, default=[22]"`
	}
	type token7 struct {
		envBoolSlice []bool `env:"ENV_BOOL_SLICE, default=[T]"`
	}
	type token8 struct {
		envFloatSlice []float32 `env:"ENV_FLOAT64_SLICE, default=[22.0]"`
	}

	// act
	// assert
	err := Bind(&token1{})
	assert.NoError(t, err)
	err = Bind(&token2{})
	assert.Error(t, err)
	err = Bind(&token3{})
	assert.Error(t, err)
	err = Bind(&token4{})
	assert.Error(t, err)
	err = Bind(&token5{})
	assert.NoError(t, err)
	err = Bind(&token6{})
	assert.Error(t, err)
	err = Bind(&token7{})
	assert.Error(t, err)
	err = Bind(&token8{})
	assert.Error(t, err)
}

func TestEmptyValue(t *testing.T) {
	defer cleanup()
	_ = os.Setenv(envString, "")
	_ = os.Setenv(envInt, "")
	_ = os.Setenv(envBool, "")
	_ = os.Setenv(envFloat64, "")
	_ = os.Setenv(envStringSlice, "")
	_ = os.Setenv(envIntSlice, "")
	_ = os.Setenv(envBoolSlice, "")
	_ = os.Setenv(envFloat64Slice, "")
	// arrange
	type token struct {
		envInt         int       `env:"ENV_INT, default=22"`
		envString      string    `env:"ENV_STRING, default=hello"`
		envBool        bool      `env:"ENV_BOOL, default=T"`
		envFloat       float64   `env:"ENV_FLOAT, default=1.0"`
		envIntSlice    []int     `env:"ENV_INT_SLICE, default=[10]"`
		envStringSlice []string  `env:"ENV_STRING_SLICE, default=[10]"`
		envBoolSlice   []bool    `env:"ENV_BOOL_SLICE, default=[1]"`
		envFloatSlice  []float64 `env:"ENV_FLOAT_SLICE, default=[10]"`
	}
	tok := &token{}
	// act
	err := Bind(tok)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, tok.envInt, 22)
	assert.Equal(t, tok.envString, "hello")
	assert.Equal(t, tok.envBool, true)
	assert.Equal(t, tok.envFloat, 1.0)
	assert.Equal(t, tok.envIntSlice, []int{10})
	assert.Equal(t, tok.envStringSlice, []string{"10"})
	assert.Equal(t, tok.envBoolSlice, []bool{true})
	assert.Equal(t, tok.envFloatSlice, []float64{10})
}

func TestProtected(t *testing.T) {
	defer cleanup()
	_ = os.Setenv(envString, "200")
	_ = os.Setenv(envInt, "200")
	_ = os.Setenv(envBool, "false")
	_ = os.Setenv(envFloat64, "200")
	_ = os.Setenv(envStringSlice, "2,0,0")
	_ = os.Setenv(envIntSlice, "2,0,0")
	_ = os.Setenv(envBoolSlice, "1,0,0")
	_ = os.Setenv(envFloat64Slice, "2,0,0")
	// arrange
	type token struct {
		envInt         int       `env:"ENV_INT, protected=true, default=100"`
		envString      string    `env:"ENV_STRING, default=100, protected=true"`
		envBool        bool      `env:"ENV_BOOL, default=F, protected=true"`
		envFloat       float64   `env:"ENV_FLOAT64, default=100, protected =true"`
		envIntSlice    []int     `env:"ENV_INT_SLICE, default=[1,0,0], protected=true"`
		envStringSlice []string  `env:"ENV_STRING_SLICE, protected=true, default=[1,0,0]"`
		envBoolSlice   []bool    `env:"ENV_BOOL_SLICE, protected=true, default=[1,0,0]"`
		envFloatSlice  []float64 `env:"ENV_FLOAT64_SLICE, protected=true, default=[100,100,100]"`
	}
	type token2 struct {
		envInt         int       `env:"ENV_INT, protected=false, default=100"`
		envString      string    `env:"ENV_STRING, default=100, protected=false"`
		envBool        bool      `env:"ENV_BOOL, default=F, protected=false"`
		envFloat       float64   `env:"ENV_FLOAT64, default=100"`
		envIntSlice    []int     `env:"ENV_INT_SLICE, default=[1,0,0]"`
		envStringSlice []string  `env:"ENV_STRING_SLICE, default=[1,0,0]"`
		envBoolSlice   []bool    `env:"ENV_BOOL_SLICE, protected=false, default=[1,0,0]"`
		envFloatSlice  []float64 `env:"ENV_FLOAT64_SLICE, protected=false, default=[100,100,100]"`
	}
	unprotected := &token2{
		envInt:         300,
		envString:      "300",
		envFloat:       300.0,
		envBool:        true,
		envIntSlice:    []int{3, 0, 0},
		envBoolSlice:   []bool{true, true, true},
		envStringSlice: []string{"300"},
		envFloatSlice:  []float64{3, 0, 0},
	}
	unprotectedEmpty := &token2{}

	filled := &token{
		envInt:         300,
		envString:      "300",
		envFloat:       300.0,
		envBool:        true,
		envIntSlice:    []int{3, 0, 0},
		envBoolSlice:   []bool{true, true, true},
		envStringSlice: []string{"300"},
		envFloatSlice:  []float64{3, 0, 0},
	}
	empty := &token{}
	// act
	err1 := Bind(filled)
	err2 := Bind(empty)
	err3 := Bind(unprotected)
	err4 := Bind(unprotectedEmpty)

	// assert
	assert.NoError(t, err1)
	assert.Equal(t, filled.envInt, 300)
	assert.Equal(t, filled.envString, "300")
	assert.Equal(t, filled.envBool, true)
	assert.Equal(t, filled.envFloat, 300.0)
	assert.Equal(t, filled.envIntSlice, []int{3, 0, 0})
	assert.Equal(t, filled.envStringSlice, []string{"300"})
	assert.Equal(t, filled.envBoolSlice, []bool{true, true, true})
	assert.Equal(t, filled.envFloatSlice, []float64{3, 0, 0})
	assert.NoError(t, err2)
	assert.Equal(t, empty.envInt, 200)
	assert.Equal(t, empty.envString, "200")
	assert.Equal(t, empty.envBool, false)
	assert.Equal(t, empty.envFloat, 200.0)
	assert.Equal(t, empty.envIntSlice, []int{2, 0, 0})
	assert.Equal(t, empty.envStringSlice, []string{"2", "0", "0"})
	assert.Equal(t, empty.envBoolSlice, []bool{true, false, false})
	assert.Equal(t, empty.envFloatSlice, []float64{2, 0, 0})

	assert.NoError(t, err3)
	assert.Equal(t, unprotected.envInt, 200)
	assert.Equal(t, unprotected.envString, "200")
	assert.Equal(t, unprotected.envBool, false)
	assert.Equal(t, unprotected.envFloat, 200.0)
	assert.Equal(t, unprotected.envIntSlice, []int{2, 0, 0})
	assert.Equal(t, unprotected.envStringSlice, []string{"2", "0", "0"})
	assert.Equal(t, unprotected.envBoolSlice, []bool{true, false, false})
	assert.Equal(t, unprotected.envFloatSlice, []float64{2, 0, 0})
	assert.NoError(t, err4)
	assert.Equal(t, unprotectedEmpty.envInt, 200)
	assert.Equal(t, unprotectedEmpty.envString, "200")
	assert.Equal(t, unprotectedEmpty.envBool, false)
	assert.Equal(t, unprotectedEmpty.envFloat, 200.0)
	assert.Equal(t, unprotectedEmpty.envIntSlice, []int{2, 0, 0})
	assert.Equal(t, unprotectedEmpty.envStringSlice, []string{"2", "0", "0"})
	assert.Equal(t, unprotectedEmpty.envBoolSlice, []bool{true, false, false})
	assert.Equal(t, unprotectedEmpty.envFloatSlice, []float64{2, 0, 0})
}

func TestEnvVarDoesntExists(t *testing.T) {
	defer cleanup()
	// arrange
	type token struct {
		ID int `env:"TOKEN_IDX, default=22"`
	}
	tok := &token{ID: 50}
	// act
	err := Bind(tok)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 22, tok.ID)
}

func TestSetFieldWhenVariableDoesntExists(t *testing.T) {
	defer cleanup()

	// arrange
	type token struct {
		ID          int       `env:"TOKEN_ID, default=22"`
		Value       string    `env:"TOKEN_VALUE, default=fc3"`
		Ratio       float64   `env:"TOKEN_RATIO, default=-0.000123"`
		Readonly    bool      `env:"TOKEN_READONLY, default=1"`
		Hours       []int     `env:"TOKEN_HOURS, default=[2,5,10]"`
		URLs        []string  `env:"TOKEN_URLS, default=[http://server.local:8080,https://server.exposed.com:80]"`
		Enabled     []bool    `env:"TOKEN_BOOLS, default=[true, false, true]"`
		Coordinates []float64 `env:"TOKEN_COORDINATES, default=[0.000123,-12.2250]"`
		EmptySlice  []bool    `env:"EMPTY_SLICE, default=[]"`
	}
	tok := &token{
		ID:          11,
		Value:       "AAAA",
		Ratio:       70.0,
		Readonly:    false,
		Hours:       []int{0},
		URLs:        nil,
		Enabled:     []bool{},
		Coordinates: nil,
		EmptySlice:  []bool{true, true},
	}

	// act
	err := Bind(tok)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 22, tok.ID)
	assert.Equal(t, "fc3", tok.Value)
	assert.Equal(t, -0.000123, tok.Ratio)
	assert.Equal(t, true, tok.Readonly)
	assert.Equal(t, []int{2, 5, 10}, tok.Hours)
	assert.Equal(t, []string{"http://server.local:8080", "https://server.exposed.com:80"}, tok.URLs)
	assert.Equal(t, []bool{true, false, true}, tok.Enabled)
	assert.Equal(t, []float64{0.000123, -12.2250}, tok.Coordinates)
	assert.Equal(t, []bool{}, tok.EmptySlice)
}

func TestReadingPrivateFields(t *testing.T) {
	defer cleanup()
	// arrange
	const val = "4BMKKDsdfsf5f7="
	_ = os.Setenv(tokenValue, val)
	type token struct {
		id       int    `env:"TOKEN_ID, default=22"`
		readonly bool   `env:"TOKEN_READONLY, require=0"`
		value    string `env:"TOKEN_VALUE, default=fc3, require=1"`
	}

	// act
	tok := &token{id: 11, value: ""}

	// act
	err := Bind(tok)
	assert.NoError(t, err)
	assert.Equal(t, 22, tok.id)
	assert.Equal(t, false, tok.readonly)
	assert.Equal(t, val, tok.value)
}

func TestReadingAnonymousFields(t *testing.T) {
	defer cleanup()
	defer cleanup()
	// arrange
	_ = os.Setenv(tokenReadOnly, "false, false")
	type token struct {
		private struct {
			readonly []bool `env:"TOKEN_READONLY, require=0"`
		}
		Exported struct {
			Value string `env:"TOKEN_VALUE, default=#fc3, require=1"`
		}
	}

	// act
	tok := &token{}

	// act
	err := Bind(tok)
	assert.NoError(t, err)
	assert.Equal(t, []bool{false, false}, tok.private.readonly)
	assert.Equal(t, "#fc3", tok.Exported.Value)
}

func TestReadingStructuresFields(t *testing.T) {
	defer cleanup()
	// arrange
	_ = os.Setenv(tokenReadOnly, "false, false")
	type private struct {
		readonly []bool `env:"TOKEN_READONLY, require=0"`
	}
	type Exported struct {
		Value string `env:"TOKEN_VALUE, default=#fc3, require=1"`
	}
	type token struct {
		ID       int
		private  private
		Exported Exported
	}

	// act
	tok := &token{ID: 50}

	// act
	err := Bind(tok)
	assert.NoError(t, err)
	assert.Equal(t, []bool{false, false}, tok.private.readonly)
	assert.Equal(t, "#fc3", tok.Exported.Value)
}

func TestNoEnv(t *testing.T) {
	defer cleanup()
	// arrange
	type private struct {
		readonly []bool
	}
	type Exported struct {
		Value string
	}
	type token struct {
		ID        int
		private   private
		Exported  Exported
		Anonymous struct {
			A string
			b string
		}
	}

	// act
	tok := &token{
		ID: 22,
		private: private{
			readonly: []bool{true, true},
		},
		Exported: Exported{
			Value: "exported"},
	}
	tok.Anonymous.A = "B"
	tok.Anonymous.b = "a"

	// assert
	err := Bind(tok)
	assert.NoError(t, err)
	assert.Equal(t, 22, tok.ID)
	assert.Equal(t, []bool{true, true}, tok.private.readonly)
	assert.Equal(t, "exported", tok.Exported.Value)
	assert.Equal(t, "B", tok.Anonymous.A)
	assert.Equal(t, "a", tok.Anonymous.b)
}

func TestSpecialSymbols(t *testing.T) {
	defer cleanup()
	// arrange
	_ = os.Setenv(tokenValue, `----~<>/?.;:/!@#$%^&*()_+_=---\-\`)
	type token struct {
		ID    string   `env:"TOKEN_ID, default=----~<>/?.;:/!@#$%^&*()_+_=---\\-"`
		Value string   `env:"TOKEN_VALUE, require=true"`
		Slice []string `env:"TOKEN_SLICE, default=[--- -~<>/?.;:/!@#$%^&*()_+_=----,  ----~<>/?.;:/!@#$%^&*()_+_=----]"`
	}
	tok := &token{}
	// act
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, `----~<>/?.;:/!@#$%^&*()_+_=---\-`, tok.ID)
	assert.Equal(t, "----~<>/?.;:/!@#$%^&*()_+_=---\\-\\", tok.Value)
	assert.Equal(t, []string{"--- -~<>/?.;:/!@#$%^&*()_+_=----", "  ----~<>/?.;:/!@#$%^&*()_+_=----"}, tok.Slice)
}

func TestSetFieldWhenVariableExists(t *testing.T) {
	defer cleanup()
	// arrange
	const val = "4BMKKDsdfsf5f7="
	_ = os.Setenv(tokenValue, val)
	type token struct {
		ID    int    `env:"TOKEN_ID"`
		Value string `env:"TOKEN_VALUE"`
	}
	tok := &token{Value: "AAAA", ID: 5}
	// act
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, val, tok.Value)
	assert.Equal(t, 0, tok.ID)
}

func TestArray(t *testing.T) {
	defer cleanup()
	// arrange
	_ = os.Setenv(tokenValue, "1,2,3")
	_ = os.Setenv(tokenID, "1, 2, 3")
	_ = os.Setenv(tokenBools, "1, 0, false, true, F,T")
	_ = os.Setenv(tokenCoordinates, "1.01, -0.05")
	type token struct {
		IDs    []int     `env:"TOKEN_VALUE"`
		Value  []string  `env:"TOKEN_VALUE"`
		Bools  []bool    `env:"TOKEN_SWITCH"`
		Floats []float64 `env:"TOKEN_COORDINATES"`
	}
	tok := &token{}
	// act
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, tok.IDs)
	assert.Equal(t, []string{"1", "2", "3"}, tok.Value)
	assert.Equal(t, []bool{true, false, false, true, false, true}, tok.Bools)
	assert.Equal(t, []float64{1.01, -0.05}, tok.Floats)
}

func TestSetMultipleFieldsByOneVariable(t *testing.T) {
	// arrange
	defer cleanup()
	const val = "4BMKKDsdfsf5f7="
	_ = os.Setenv(tokenValue, val)
	type token struct {
		ID1    int    `env:"TOKEN_ID, default=50"`
		ID2    int    `env:"TOKEN_ID, default=10"`
		Value1 string `env:"TOKEN_VALUE, require=true"`
		Value2 string `env:"TOKEN_VALUE, require=true"`
	}
	// act
	tok := &token{}
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, tok.ID1, 50)
	assert.Equal(t, tok.ID2, 10)
	assert.Equal(t, tok.Value1, val)
	assert.Equal(t, tok.Value2, val)
}

func TestLowercase(t *testing.T) {
	// arrange
	val := "private field"
	_ = os.Unsetenv(strings.ToLower(strings.ToLower(privateTokenValue)))
	_ = os.Setenv(strings.ToLower(privateTokenValue), val)
	// act
	type token struct {
		valLowerCase string `env:"private_token_value"`
		valUpperCase string `env:"PRIVATE_TOKEN_VALUE"`
	}
	tok := &token{}

	// assert
	err := Bind(tok)
	assert.NoError(t, err)
	assert.Equal(t, val, tok.valLowerCase)
	assert.Equal(t, "", tok.valUpperCase)
}

func TestPrefixSuccessfully(t *testing.T) {
	defer cleanup()
	// arrange
	_ = os.Setenv(privateTokenValue, "private field")
	_ = os.Setenv(exportedTokenValue, "exported field")
	_ = os.Setenv(privateExportedTokenValue, "private in exported field")
	_ = os.Setenv("TEMP_TOKEN_ID", "20")
	_ = os.Setenv("TEMP_FIELD_HIDDEN", "very hidden")
	type private struct {
		Value string `env:"TOKEN_VALUE, require=0"`
	}
	type Exported struct {
		Value string `env:"TOKEN_VALUE, require=1"`
	}
	type token struct {
		private  private  `env:"PRIVATE"`
		Exported Exported `env:"EXPORTED"`
		Temp     struct {
			ID    int `env:"TOKEN_ID"`
			field struct {
				hidden string `env:"HIDDEN"`
			} `env:"FIELD"`
		} `env:"TEMP"`
		ID int `env:"TOKEN_ID, default=-1"`
	}

	// act
	tok := &token{ID: 50}
	err := Bind(tok)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, -1, tok.ID)
	assert.Equal(t, tok.private.Value, "private field")
	assert.Equal(t, tok.Exported.Value, "exported field")
	assert.Equal(t, tok.Temp.ID, 20)
	assert.Equal(t, tok.Temp.field.hidden, "very hidden")
}

func TestParseSimple(t *testing.T) {
	// arrange
	defer cleanup()
	env := make(map[string]string)
	env[envString] = "foo"
	env[envBool] = "true"
	setEnv(env)

	type token struct {
		ID    int    `env:"TOKEN_ID"`
		Value string `env:"TOKEN_VALUE"`
	}
	type s struct {
		UID       int       `env:"ENV_INT,default=55"`
		name      string    `env:"ENV_STRING,require=true, default=michal"`
		B         bool      `env:"ENV_BOOL"`
		Bs        []bool    `env:"ENV_BOOL_SLICE,default=[true, false, true]"`
		nums      []int     `env:"ENV_INT_SLICE,default=[10,0,5, 3, 4, 5]"`
		F         float64   `env:"ENV_FLOAT64, default=0.000000002121"`
		Fs        []float64 `env:"ENV_FLOAT64_SLICE,default=[0.0021,0.002,1.13,2.15]"`
		surname   string
		Token     token
		Anonymous struct {
			secret      string `env:"ENV_STRING_PRIVATE"`
			TopSecret   bool
			TokenBase64 string   `env:"ENV_STRING_EXPORTED"`
			arr         []string `env:"ANONYMOUS_ARR, default= [abc, xyz, 123] "`
		}
	}
	testee := &s{name: "Michal"}

	// act
	err := Bind(testee)

	// assert
	assert.Equal(t, testee.name, "foo")
	assert.Equal(t, testee.UID, 55)
	assert.Equal(t, testee.nums, []int{10, 0, 5, 3, 4, 5})
	assert.Equal(t, testee.Bs, []bool{true, false, true})
	assert.True(t, reflect.DeepEqual(testee.Anonymous.arr, []string{"abc", " xyz", " 123"}))
	assert.Equal(t, testee.F, 0.000000002121)
	assert.Equal(t, []float64{0.0021, 0.002, 1.13, 2.15}, testee.Fs)
	assert.False(t, testee.Anonymous.TopSecret)
	assert.NoError(t, err)
}

func TestService(t *testing.T) {
	type Endpoint struct {
		URL string `env:"ENDPOINT_URL, require=true"`
	}
	type Config struct {
		// reading string value from NAME
		Name string `env:"NAME"`
		// reading int with 8080 as default value
		DefaultPot int `env:"PORT, default=8080"`
		// reading slice of strings with default values
		Regions []string `env:"REGIONS, default=[us-east-1,us-east-2,us-west-1]"`
		// inline structure
		Credentials struct {
			// binding required value
			KeyID string `env:"ACCESS_KEY_ID, require=true"`
			// binding to private field
			secretKey string `env:"SECRET_ACCESS_KEY, require=true"`
		}
		// expected PRIMARY_ prefix in nested environment variables
		Endpoint1 Endpoint `env:"PRIMARY"`
		// expected FAILOVER_ prefix in nested environment variables
		Endpoint2 Endpoint `env:"FAILOVER"`
		// reuse an already bound env variable NAME
		Description string `env:"NAME"`
		// the field does not have a bind tag set, so it will be ignored during bind
		Args []string
	}

	_ = os.Setenv(primaryEndpointURL, "https://ep1.cloud.example.com")
	_ = os.Setenv(failoverEndpointURL, "https://ep2.cloud.example.com")
	_ = os.Setenv(name, "Hello from 12-factor")
	_ = os.Setenv(defaultPort, "9000")
	_ = os.Setenv(accessKeyID, "AKIAIOSFODNN7EXAMPLE")
	_ = os.Setenv(secretAccessKey, `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`)

	// act
	c := Config{Description: "Hello from os.LookupEnv()", Args: []string{"debug=true"}}
	err := Bind(&c)
	// assert
	assert.NoError(t, err)
	assert.Equal(t, "https://ep1.cloud.example.com", c.Endpoint1.URL)
	assert.Equal(t, "https://ep2.cloud.example.com", c.Endpoint2.URL)
	assert.Equal(t, "Hello from 12-factor", c.Name)
	assert.Equal(t, 9000, c.DefaultPot)
	assert.Equal(t, "AKIAIOSFODNN7EXAMPLE", c.Credentials.KeyID)
	assert.Equal(t, `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`, c.Credentials.secretKey)
}

func setEnv(m map[string]string) {
	for k, v := range m {
		_ = os.Setenv(k, v)
	}
}

const (
	tokenID          = "TOKEN_ID"
	tokenValue       = "TOKEN_VALUE"
	tokenRatio       = "TOKEN_RATIO"
	tokenReadOnly    = "TOKEN_READONLY"
	tokenHours       = "TOKEN_HOURS"
	tokenURLs        = "TOKEN_URLS"
	tokenBools       = "TOKEN_SWITCH"
	tokenCoordinates = "TOKEN_COORDINATES"

	privateTokenValue = "PRIVATE_TOKEN_VALUE"
	// #nosec G101;
	exportedTokenValue        = "EXPORTED_TOKEN_VALUE"
	privateExportedTokenValue = "PRIVATE_EXPORTED_TOKEN_VALUE"

	primaryEndpointURL  = "PRIMARY_ENDPOINT_URL"
	failoverEndpointURL = "FAILOVER_ENDPOINT_URL"
	name                = "NAME"
	defaultPort         = "PORT"
	accessKeyID         = "ACCESS_KEY_ID"
	secretAccessKey     = "SECRET_ACCESS_KEY"

	envInt            = "ENV_INT"
	envString         = "ENV_STRING"
	envStringSlice    = "ENV_STRING_SLICE"
	envBool           = "ENV_BOOL"
	envBoolSlice      = "ENV_BOOL_SLICE"
	envFloat64        = "ENV_FLOAT64"
	envFloat64Slice   = "ENV_FLOAT64_SLICE"
	envStringPrivate  = "ENV_STRING_PRIVATE"
	envStringExported = "ENV_STRING_EXPORTED"
	envIntSlice       = "ENV_INT_SLICE"
)

func cleanup() {
	_ = os.Unsetenv(tokenID)
	_ = os.Unsetenv(tokenValue)
	_ = os.Unsetenv(tokenRatio)
	_ = os.Unsetenv(tokenReadOnly)
	_ = os.Unsetenv(tokenHours)
	_ = os.Unsetenv(tokenURLs)
	_ = os.Unsetenv(tokenBools)
	_ = os.Unsetenv(tokenCoordinates)

	_ = os.Unsetenv(privateTokenValue)
	_ = os.Unsetenv(privateExportedTokenValue)
	_ = os.Unsetenv(exportedTokenValue)

	_ = os.Unsetenv(primaryEndpointURL)
	_ = os.Unsetenv(failoverEndpointURL)
	_ = os.Unsetenv(name)
	_ = os.Unsetenv(defaultPort)
	_ = os.Unsetenv(accessKeyID)
	_ = os.Unsetenv(secretAccessKey)

	_ = os.Unsetenv(envInt)
	_ = os.Unsetenv(envString)
	_ = os.Unsetenv(envBool)
	_ = os.Unsetenv(envBoolSlice)
	_ = os.Unsetenv(envStringPrivate)
	_ = os.Unsetenv(envStringExported)
	_ = os.Unsetenv(envIntSlice)
	_ = os.Unsetenv(envFloat64)
	_ = os.Unsetenv(envFloat64Slice)
}
