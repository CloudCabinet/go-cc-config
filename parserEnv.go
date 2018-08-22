package config

import (
	"os"
	"errors"
	"github.com/caarlos0/env"
	"strings"
)

type parserEnv struct{
	getters
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


