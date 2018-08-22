package config

import (
	"github.com/caarlos0/env"
	"errors"
)
type config struct {
	ConfigType    		 string  `env:"CONFIG_TYPE" envDefault:"ENV"`
	ConfigStorageDefault string  `env:"CONFIG_SOURCE_DEFAULT"`
	ConfigStoragePrefix  string  `env:"CONFIG_SOURCE_PREFIX"`
}
var cfg = config{}
type ConfigInterface interface {
	Get(key string) (string,error)
	GetAll() (map[string]interface{},error)
	GetBool(key string) (bool,error)
	GetInt64(key string, property ...int) (int64,error)
	GetFloat64(key string) (float64,error)
	GetUint64(key string, property ...int) (uint64,error)
	GetDefault(key,def string) (string)
	AssignTo(data interface{}) (error)
	getStorage(name_storage string) (ConfigInterface)
}

var configInterface ConfigInterface

func init(){
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	switch cfg.ConfigType {
	case "ENV":
		configInterface=new(parserEnv)
	case "JSON":
		if(len(cfg.ConfigStoragePrefix)==0){
			cfg.ConfigStoragePrefix="."
		}
		configInterface=new(parserJson)
	default:
		panic(errors.New("Not support CONFIG_TYPE [" +cfg.ConfigType+"]" ))
	}
}
