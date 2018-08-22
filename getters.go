package config

import (
	"strconv"
	"errors"
)

func GetStorage(key ...string) (ConfigInterface) {
	if(len(key)==0){
		return configInterface.getStorage(cfg.ConfigStorageDefault);
	}
	return configInterface.getStorage(key[0]);
}

type getters struct{
}
func (p *getters) Get(key string) (string,error) {
	return "",errors.New("Not realized func by type source = "+cfg.ConfigType)
}
func (p *getters) GetAll() (map[string]interface{},error) {
	return make(map[string]interface{}),errors.New("Not realized func by type source = "+cfg.ConfigType)
}
func (p *getters) AssignTo(data interface{}) (error) {
	return errors.New("Not realized func by type source = "+cfg.ConfigType)
}
func (p *getters) getStorage(name_storage string) (ConfigInterface) {
	return configInterface
}

func (p *getters) GetDefault(key,def string) (string) {
	if value,err:=p.Get(key);err==nil{
		return value
	}else{
		return def
	}
}

func (p *getters) GetBool(key string) (bool,error) {
	if value,err:=p.Get(key);err==nil{
		return strconv.ParseBool(value)
	}else{
		return false,err
	}
}

func (p *getters) GetInt64(key string, property ...int) (int64,error) {
	if(len(property)!=2){
		property[0]=10
		property[1]=64
	}
	if value,err:=p.Get(key);err==nil{
		return strconv.ParseInt(value,property[0],property[1])
	}else{
		return 0,err
	}
}

func (p *getters) GetFloat64(key string) (float64,error) {
	if value,err:=p.Get(key);err==nil{
		return strconv.ParseFloat(value,64)
	}else{
		return 0,err
	}
}

func (p *getters) GetUint64(key string, property ...int) (uint64,error) {
	if(len(property)!=2){
		property[0]=10
		property[1]=64
	}
	if value,err:=p.Get(key);err==nil{
		return strconv.ParseUint(value,property[0],property[1])
	}else{
		return 0,err
	}
}
