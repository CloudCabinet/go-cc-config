package config

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type parserJson struct {
	storageByte    []byte
	storage map[string]interface{}
}
func newJson(key ...string) *parserJson {
	file, e := ioutil.ReadFile(cfg.ConfigStoragePrefix + key[0] + ".json")
	if (e != nil) {
		panic(e)
	}
	var jsontype map[string]interface{}
	json.Unmarshal(file, &jsontype)
	data  := &parserJson{
		storageByte: file,
		storage: jsontype,
	}
	return data
}
func (p *parserJson) Get(key string) (string, error) {
	if value, err := p.storage[key]; err {
		return value.(string), nil
	} else {
		return "", errors.New("Not found config key=[" + key + "]")
	}
}
func (p *parserJson) GetAll() (map[string]interface{}, error) {
	return p.storage, nil
}
func (p *parserJson) AssignTo(data interface{}) (error) {
	err := json.Unmarshal(p.storageByte, data)
	if err != nil {
		return err
	}
	return nil
}
func (p *parserJson) GetDB() (*dbConfig) {
	db:=new(dbConfig)
	err := json.Unmarshal(p.storageByte, db)
	if err != nil {
		panic(err)
	}
	return db
}
func (p *parserJson) GetString(key string) (string) {
	return getString(p,key)
}
func (p *parserJson) GetDefault(key,def string) (string) {
	return getDefault(p,key,def)
}

func (p *parserJson) GetBool(key string) (bool,error) {
	return getBool(p,key)
}

func (p *parserJson) GetInt64(key string, property ...int) (int64,error) {
	return getInt64(p,key,property...)
}

func (p *parserJson) GetFloat64(key string) (float64,error) {
	return getFloat64(p,key)
}

func (p *parserJson) GetUint64(key string, property ...int) (uint64,error) {
	return getUint64(p,key,property...)
}
func (p *parserJson) GetInterface(key string) (interface{}) {
	if value, err := p.storage[key]; err {
		return value
	} else {
		panic(errors.New("Not found config key=[" + key + "]"))
	}
}
