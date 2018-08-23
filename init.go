package config

import (
	"github.com/caarlos0/env"
)
type config struct {
	ConfigType    		 string  `env:"CONFIG_TYPE" envDefault:"ENV"`
	ConfigStorageDefault string  `env:"CONFIG_SOURCE_DEFAULT"`
	ConfigStoragePrefix  string  `env:"CONFIG_SOURCE_PREFIX"`
}
var cfg = config{}
type ConfigInterface interface {
	Get(key string) (string,error)
	GetString(key string) (string)
	GetAll() (map[string]interface{},error)
	GetBool(key string) (bool,error)
	GetInt64(key string, property ...int) (int64,error)
	GetFloat64(key string) (float64,error)
	GetUint64(key string, property ...int) (uint64,error)
	GetDefault(key,def string) (string)
	AssignTo(data interface{}) (error)
}
type dataStorage struct {
	configObject ConfigInterface
}
var configDataStorage map[string]*dataStorage

func init(){
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	configDataStorage = make(map[string]*dataStorage)
}
