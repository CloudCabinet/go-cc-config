package config

import (
	"github.com/caarlos0/env"
)
type config struct {
	ConfigType    		 string  `env:"CONFIG_TYPE" envDefault:"ENV"`
	ConfigStorageDefault string  `env:"CONFIG_SOURCE_DEFAULT"`
	ConfigStoragePrefix  string  `env:"CONFIG_SOURCE_PREFIX"`
}
type dbConfig struct{
	Id string `json:"id" env:"DB_ID"`
	Host string `json:"host" env:"DB_HOST"`
	Port uint16 `json:"port,string" env:"DB_PORT"`
	Service string `json:"service" env:"DB_SERVICE"`
	User string `json:"user" env:"DB_USER"`
	MaxConnections int `json:"max_connections,string" env:"DB_MAX_CONNECTION"`
	LifeConnections int64 `json:"life_connections,string" env:"DB_LIFE_CONNECTIONS"`
	DataBase string `json:"database" env:"DB_DATA_BASE"`
	Shem string `json:"shem" env:"DB_SHEM"`
	Password string `json:"password" env:"DB_PASSWORD"`
}
var cfg = config{}
type ConfigInterface interface {
	Get(key string) (string,error)
	GetDB() (*dbConfig)
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
