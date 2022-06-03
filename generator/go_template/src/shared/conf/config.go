package conf

import (
	"log"
	"os"
	"sync"

	"github.com/gustavohenrique/coolconf"
)

const (
	KEY_DB            = "db"
	DB_NAME_SQLITE    = "sqlite"
	DB_NAME_DGRAPH    = "dgraph"
	EVENT_ACTION_DONE = "action.done"
)

type Config struct {
	Debug bool `env:"DEBUG" `
	Auth  struct {
		Disabled        bool   `env:"AUTH_DISABLED" yaml:"disabled"`
		TokenExpiration string `env:"AUTH_TOKEN_EXPIRATION" yaml:"token_expiration" default:"900"`
		SkipRouters     []string
	}
	Log struct {
		Level  string `env:"LOG_LEVEL" yaml:"level" default:"info"`
		Output string `env:"LOG_OUTPUT" yaml:"output" default:"stdout"`
		Format string `env:"LOG_FORMAT" yaml:"format" default:"text"`
	}
	Grpc struct {
		Address           string   `env:"GRPC_ADDRESS" yaml:"address" default:"0.0.0.0"`
		Port              int      `env:"GRPC_PORT" yaml:"port" default:"8002"`
		SkipRouters       []string `default:"/SignIn,/SignUp"`
		MaxReceiveMsgSize int      `env:"GRPC_MAX_RECEIVE_MSG_SIZE" yaml:"max_receive_msg_size" default:"5368709120"` // 5MB
		MaxSendMsgSize    int      `env:"GRPC_MAX_SEND_MSG_SIZE" yaml:"max_send_msg_size" default:"5368709120"`       // 5MB
		TLS               struct {
			Enabled bool   `env:"GRPC_TLS_ENABLED" yaml:"enabled"`
			Key     string `env:"GRPC_TLS_KEY" yaml:"key"`
			Cert    string `env:"GRPC_TLS_CERT" yaml:"cert"`
		}
	}
	Http struct {
		Address string `env:"HTTP_ADDRESS" yaml:"address" default:"0.0.0.0"`
		Port    int    `env:"HTTP_PORT" yaml:"port" default:"8001"`
		TLS     struct {
			Enabled bool   `env:"HTTP_TLS_ENABLED" yaml:"enabled"`
			Key     string `env:"HTTP_TLS_KEY" yaml:"key"`
			Cert    string `env:"HTTP_TLS_CERT" yaml:"cert"`
		}
	}
	Store struct {
		Postgres struct {
			URL             string `env:"STORE_POSTGRES_URL" yaml:"url" default:"postgres://root:root@127.0.0.1/{{ .ProjectName }}?sslmode=disable"`
			MaxConns        int32  `env:"STORE_POSTGRES_MAX_CONN" yaml:"max_conns" default:"500"`
			MaxConnLifetime int    `env:"STORE_POSTGRES_MAX_CONN_LIFETIME" yaml:"max_conn_lifetime" default:"480"`
			MaxConnIdleTime int    `env:"STORE_POSTGRES_MAX_CON_IDLETIME" yaml:"max_conn_iddle_time" default:"380"`
		}
		Sqlite struct {
			Address string `env:"STORE_SQLITE_ADDRESS" yaml:"address" default:":memory:"`
			Schema  string `env:"STORE_SQLITE_SCHEMA" yaml:"schema"`
		}
		Dgraph struct {
			Http string `env:"STORE_DGRAPH_URL" yaml:"http" default:"http://localhost:8080"`
			Grpc string `env:"STORE_DGRAPH_GRPC" yaml:"grpc" default:"localhost:9080"`
		}
	}
}

var globalConfig *Config

func Parse(configFile string) *Config {
	coolconf.New(coolconf.Settings{
		Source:                  configFile,
		ShouldCreateDefaultYaml: true,
	})
	if err := coolconf.Load(&globalConfig); err != nil {
		log.Fatalln("Parser:", err)
	}
	globalConfig.Auth.SkipRouters = []string{"/SignIn", "/FindOneByEmail", "/CreatePasswordFor"}
	singleton(globalConfig)
	return globalConfig
}

func Get() *Config {
	if globalConfig != nil {
		return globalConfig
	}
	key := "CONFIG_FILE"
	var configFile string = os.Getenv(key)
	if configFile == "" {
		log.Fatalln("[INFO]", key, "is empty. Cannot read the configuration file")
	}
	return Parse(configFile)
}

func singleton(c *Config) {
	var once sync.Once
	once.Do(func() {
		globalConfig = c
	})
}