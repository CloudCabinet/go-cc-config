package config

import (
	"strconv"

)
func getString(configObject ConfigInterface,key string) (string) {
	if value,err:=configObject.Get(key);err==nil{
		return value
	}else{
		panic("Not fount config key=["+key+"]");
	}
}
func getDefault(configObject ConfigInterface,key,def string) (string) {
	if value,err:=configObject.Get(key);err==nil{
		return value
	}else{
		return def
	}
}

func getBool(configObject ConfigInterface,key string) (bool,error) {
	if value,err:=configObject.Get(key);err==nil{
		return strconv.ParseBool(value)
	}else{
		return false,err
	}
}

func getInt64(configObject ConfigInterface,key string, property ...int) (int64,error) {
	var (
		base,bitSize int
	)
	if(len(property)==2){
		base = property[0]
		bitSize = property[1]
	}else{
		base = 10
		bitSize = 64
	}
	if value,err:=configObject.Get(key);err==nil{
		return strconv.ParseInt(value,base,bitSize)
	}else{
		return 0,err
	}
}

func getFloat64(configObject ConfigInterface,key string) (float64,error) {
	if value,err:=configObject.Get(key);err==nil{
		return strconv.ParseFloat(value,64)
	}else{
		return 0,err
	}
}

func getUint64(configObject ConfigInterface,key string, property ...int) (uint64,error) {
	var (
		base,bitSize int
	)
	if(len(property)==2){
		base = property[0]
		bitSize = property[1]
	}else{
		base = 10
		bitSize = 64
	}
	if value,err:=configObject.Get(key);err==nil{
		return strconv.ParseUint(value,base,bitSize)
	}else{
		return 0,err
	}
}
