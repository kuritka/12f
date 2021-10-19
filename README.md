# ENV binder

```golang
// reusable structure
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

        // reading slice of strings from env var
        Subnets []string `env:"SUBNETS, default=[10.0.0.0/24,192.168.1.0/24]"`
        
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

    // PRIMARY_ENDPOINT_URL  = https://ep1.cloud.example.com
    // FAILOVER_ENDPOINT_URL = https://ep2.cloud.example.com
    // ACCESS_KEY_ID         = AKIAIOSFODNN7EXAMPLE
    // SECRET_ACCESS_KEY     = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
    // NAME                  = Hello from 12-factor
    // PORT                  = 9000
    // SUBNETS               = 10.0.0.0/24,10.0.1.0/24, 10,10.1.0.0/24,  10.1.1.0/24


    c := Config{Description: "Hello from os.LookupEnv()", Args: []string{"debug=true"}}
    err := Bind(&c)
    if err == nil {
        fmt.Println(Format(c))
    }
```

```json
{
	"Name": "Hello from 12-factor",
	"DefaultPot": 9000,
	"Regions": [
		"us-east-1",
		"us-east-2",
		"us-west-1"
	],
	"Subnets": [
		"10.0.0.0/24",
		"10.0.1.0/24",
		"10,10.1.0.0/24",
		"10.1.1.0/24"
	],
	"Credentials": {
		"KeyID": "AKIAIOSFODNN7EXAMPLE"
	},
	"Endpoint1": {
		"Url": "https://ep1.cloud.example.com"
	},
	"Endpoint2": {
		"Url": "https://ep2.cloud.example.com"
	},
	"Description": "Hello from 12-factor",
	"Args": [
		"debug=true"
	]
}

```


