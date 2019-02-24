package config

import (
	"os"
	"errors"
	"github.com/caarlos0/env"
	"strings"
)

type parserEnv struct{
}

func newEnv(key ...string) *parserEnv {
	return &parserEnv{}
}
func (p *parserEnv) Get(key string) (string,error) {
	if value,err:=os.LookupEnv(key);err{
		return value,nil
	}else{
		return value,errors.New("Not fount config key=["+key+"]")
	}
}
func (p *parserEnv) GetAll() (map[string]interface{},error) {
	m := make(map[string]interface{})
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			m[e[:i]] = e[i+1:]
		}
	}
	return m,nil
}
func (p *parserEnv) AssignTo(data interface{}) (error) {
	err := env.Parse(data)
	if err != nil {
		return err
	}
	return nil
}


func (p *parserEnv) GetDB() (*dbConfig) {
	db:=new(dbConfig)
	err := env.Parse(db)
	if err != nil {
		panic(err)
	}
	return db
}
func (p *parserEnv) GetString(key string) (string) {
	return getString(p,key)
}
func (p *parserEnv) GetDefault(key,def string) (string) {
	return getDefault(p,key,def)
}

func (p *parserEnv) GetBool(key string) (bool,error) {
	return getBool(p,key)
}

func (p *parserEnv) GetInt64(key string, property ...int) (int64,error) {
	return getInt64(p,key,property...)
}

func (p *parserEnv) GetFloat64(key string) (float64,error) {
	return getFloat64(p,key)
}

func (p *parserEnv) GetUint64(key string, property ...int) (uint64,error) {
	return getUint64(p,key,property...)
}
func (p *parserEnv) GetInterface(key string) (interface{}) {
	panic(errors.New("This metod not support ENV"))
}
