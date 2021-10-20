# ENV binder
The env-binder package is used to easily bind values to GO structures. Env-binder is designed to 
be usable in the widest possible range of scenarios.Among other things, it supports variable 
prefixes and bindings to unexported arrays. Take a look at the following usage example:
```golang
import "github.com/kuritka/12f/env"

type Endpoint struct {
	URL string `env:"ENDPOINT_URL, require=true"`
}

type Config struct {

	// reading string value from NAME
	Name string `env:"NAME"`

	// reuse an already bound env variable NAME
	Description string `env:"NAME"`

	// reuse an already bound variable NAME, but replace only when name was not set before
	AlternativeName string `env:"NAME, protected=true"`

	// reading int with 8080 as default value
	DefaultPort int16 `env:"PORT, default=8080"`

	// reading slice of strings with default values
	Regions []string `env:"REGIONS, default=[us-east-1,us-east-2,us-west-1]"`

	// reading slice of strings from env var
	Subnets []string `env:"SUBNETS, default=[10.0.0.0/24,192.168.1.0/24]"`

	// nested structure
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

	// the field does not have a bind tag set, so it will be ignored during bind
	Args []string
}


func main() {
	defer clean()
	os.Setenv("PRIMARY_ENDPOINT_URL", "https://ep1.cloud.example.com")
	os.Setenv("FAILOVER_ENDPOINT_URL", "https://ep2.cloud.example.com")
	os.Setenv("ACCESS_KEY_ID", "AKIAIOSFODNN7EXAMPLE")
	os.Setenv("SECRET_ACCESS_KEY", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	os.Setenv("NAME", "Hello from 12-factor")
	os.Setenv("PORT", "9000")
	os.Setenv("SUBNETS", "10.0.0.0/24,10.0.1.0/24, 10.1.0.0/24,  10.1.1.0/24")

	c := &Config{}
	c.AlternativeName = "protected name"
	c.Description = "hello from env-binder"
	if err := env.Bind(c); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(JSONize(c))
}
```

```json
{
  "Name": "Hello from 12-factor",
  "Description": "Hello from 12-factor",
  "AlternativeName": "protected name",
  "DefaultPort": 9000,
  "Regions": [
    "us-east-1",
    "us-east-2",
    "us-west-1"
  ],
  "Subnets": [
    "10.0.0.0/24",
    "10.0.1.0/24",
    "10.1.0.0/24",
    "10.1.1.0/24"
  ],
  "Credentials": {
    "KeyID": "AKIAIOSFODNN7EXAMPLE"
  },
  "Endpoint1": {
    "URL": "https://ep1.cloud.example.com"
  },
  "Endpoint2": {
    "URL": "https://ep2.cloud.example.com"
  },
  "Args": null
}
```

## supported types
env-binder supports all types listed in the following table.  In addition, it should be noted that in the case 
of slices, env-binder creates an instance of an empty slice if the value of the environment variable is 
declared and its value is `"""`. In this case env-binder returns an empty slice instead of the vulnerable nil.  
| primitive types | slices |
|---|---|
| `int`,`int8`,`int16`,`int32`,`int64` | `[]int`,`[]int8`,`[]int16`,`[]int32`,`[]int64` |
| `float32`,`float64` | `[]float32`,`[]float64` |
| `uint`,`uint8`,`uint16`,`uint32`,`uint64` | `[]uint`,`[]uint8`,`[]uint16`,`[]uint32`,`[]uint64` |
| `bool` | `[]bool` |
| `string` | `[]string` |