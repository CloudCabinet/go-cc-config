package config

import (
	"errors"
)

func GetStorage(key ...string) (ConfigInterface) {
	var (
		name_storage string
		type_storage string
	)
	if(len(key)==0){
		name_storage=cfg.ConfigStorageDefault;
		type_storage=cfg.ConfigType;
	}else if (len(key)==1){
		name_storage=key[0]
		type_storage=cfg.ConfigType;
	}else{
		name_storage=key[0]
		type_storage=key[1]
	}
	if value, err := configDataStorage[name_storage]; err {
		return value.configObject
	}else{
		switch type_storage {
		case "ENV":
			configDataStorage[name_storage]=&dataStorage{
				configObject:newEnv(key...)}
		case "JSON":
			if(len(cfg.ConfigStoragePrefix)==0){
				cfg.ConfigStoragePrefix="."
			}
			configDataStorage[name_storage]=&dataStorage{
				configObject:newJson(key...)}

		default:
			panic(errors.New("Not support CONFIG_TYPE [" +cfg.ConfigType+"]" ))
		}
	}
	return configDataStorage[name_storage].configObject;
}