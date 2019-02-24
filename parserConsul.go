package config

import (
	"encoding/json"
	consulapi "github.com/hashicorp/consul/api"
	"errors"
"strings"

)

type parserConsul struct {
	storageByte    []byte
	storage map[string]interface{}
}
var consul *consulapi.Client

func newConsul(key ...string) *parserConsul {
	var e error
	if(consul==nil){
		sd := strings.Split(getOS("SERVER_DISCOVERY", "http://localhost:8500"),"://")
		config := consulapi.DefaultConfig()
		config.Address = sd[1]
		config.Scheme = sd[0]
		consul, e = consulapi.NewClient(config)
		if (e != nil) {
			panic(e)
		}
	}
	kv := consul.KV()
	pair, _, e := kv.Get(cfg.ConfigStoragePrefix+key[0], nil)
	if (e != nil) {
		panic(e)
	}
	if(pair==nil){
		panic( errors.New("Not found config storage=[" + cfg.ConfigStoragePrefix+key[0] + "]"))
	}

	var jsontype map[string]interface{}
	json.Unmarshal(pair.Value, &jsontype)
	data  := &parserConsul{
		storageByte: pair.Value,
		storage: jsontype,
	}
	return data
}
func (p *parserConsul) Get(key string) (string, error) {
	if value, err := p.storage[key]; err {
		return value.(string), nil
	} else {
		return "", errors.New("Not found config key=[" + key + "]")
	}
}
func (p *parserConsul) GetAll() (map[string]interface{}, error) {
	return p.storage, nil
}
func (p *parserConsul) AssignTo(data interface{}) (error) {
	err := json.Unmarshal(p.storageByte, data)
	if err != nil {
		return err
	}
	return nil
}
func (p *parserConsul) GetDB() (*dbConfig) {
	db:=new(dbConfig)
	err := json.Unmarshal(p.storageByte, db)
	if err != nil {
		panic(err)
	}
	return db
}
func (p *parserConsul) GetString(key string) (string) {
	return getString(p,key)
}
func (p *parserConsul) GetDefault(key,def string) (string) {
	return getDefault(p,key,def)
}

func (p *parserConsul) GetBool(key string) (bool,error) {
	return getBool(p,key)
}

func (p *parserConsul) GetInt64(key string, property ...int) (int64,error) {
	return getInt64(p,key,property...)
}

func (p *parserConsul) GetFloat64(key string) (float64,error) {
	return getFloat64(p,key)
}

func (p *parserConsul) GetUint64(key string, property ...int) (uint64,error) {
	return getUint64(p,key,property...)
}
func (p *parserConsul) GetInterface(key string) (interface{}) {
	if value, err := p.storage[key]; err {
		return value
	} else {
		panic(errors.New("Not found config key=[" + key + "]"))
	}
}