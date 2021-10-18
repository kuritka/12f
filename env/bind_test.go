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
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	envUID          = "ENV_UID"
	envName         = "ENV_NAME"
	envB            = "ENV_B"
	envBs           = "ENV_Bs"
	anonymousSecret = "ANONYMOUS_SECRET"
	anonymousToken  = "ANONYMOUS_TOKEN"
	nums            = "ENV_NUMS"
	envF            = "ENV_F"
	envFs           = "ENV_Fs"
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

	_ = os.Unsetenv(envUID)
	_ = os.Unsetenv(envName)
	_ = os.Unsetenv(envB)
	_ = os.Unsetenv(envBs)
	_ = os.Unsetenv(anonymousSecret)
	_ = os.Unsetenv(anonymousToken)
	_ = os.Unsetenv(nums)
	_ = os.Unsetenv(envF)
	_ = os.Unsetenv(envFs)
}

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
	os.Setenv(tokenID, "blah")
	// arrange
	type token struct {
		ID int `env:"TOKEN_ID, default=22"`
	}
	tok := &token{}
	// act
	err := Bind(tok)

	// assert
	assert.Error(t, err)
}

func TestEmptyValue(t *testing.T) {
	defer cleanup()
	os.Setenv(tokenID, "")
	os.Setenv(tokenValue, "")
	// arrange
	type token struct {
		ID    int    `env:"TOKEN_ID, default=22"`
		Value string `env:"TOKEN_VALUE, default=hello"`
	}
	tok := &token{}
	// act
	err := Bind(tok)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, tok.ID, 22)
	assert.Equal(t, tok.Value, "hello")
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
	os.Setenv(tokenValue, `----~<>/?.;:/!@#$%^&*()_+_=---\-\`)
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
	env[envName] = "foo"
	env[envB] = "true"
	setEnv(env)

	type token struct {
		ID    int    `env:"TOKEN_ID"`
		Value string `env:"TOKEN_VALUE"`
	}
	type s struct {
		UID       int       `env:"ENV_UID,default=55"`
		name      string    `env:"ENV_NAME,require=true, default=michal"`
		B         bool      `env:"ENV_B"`
		Bs        []bool    `env:"ENV_BS,default=[true, false, true]"`
		nums      []int     `env:"ENV_NUMS,default=[10,0,5, 3, 4, 5]"`
		F         float64   `env:"ENV_F, default=0.000000002121"`
		Fs        []float64 `env:"ENV_Fs,default=[0.0021,0.002,1.13,2.15]"`
		surname   string
		Token     token
		Anonymous struct {
			secret      string `env:"ANONYMOUS_SECRET"`
			TopSecret   bool
			TokenBase64 string   `env:"ANONYMOUS_TOKEN"`
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
