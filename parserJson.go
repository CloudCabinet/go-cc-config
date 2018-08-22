package config

import (
	"io/ioutil"
	"encoding/json"
	"errors"
)

type parserJson struct {
	getters
	file    []byte
	storage map[string]interface{}
}

var storageJson = make(map[string]parserJson)

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
	err := json.Unmarshal(p.file, data)
	if err != nil {
		return err
	}
	return nil
}
func (p *parserJson) getStorage(name_storage string) (ConfigInterface) {
	if storage, err := storageJson[name_storage]; err {
		return &storage
	} else {
		file, e := ioutil.ReadFile(cfg.ConfigStoragePrefix + name_storage + ".json")
		if (e != nil) {
			panic(e)
		}
		var jsontype map[string]interface{}
		json.Unmarshal(file, &jsontype)
		_parserJson := parserJson{
			file: file,
			storage: jsontype,
		}
		storageJson[name_storage] = _parserJson

		return &_parserJson
	}

}