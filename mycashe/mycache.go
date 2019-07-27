package mycashe

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"time"
)

func MyCache() (cache.Cache) {
	var redis cache.Cache
	mstr:=map[string]string{}
	mstr["key"]="myapp"
	mstr["conn"]="127.0.0.1:6379"
	mstr["dbNum"]="0"
	bytes, _ := json.Marshal(mstr)
	redis, err := cache.NewCache("redis", string(bytes))
	if err != nil {
		fmt.Println("Redis Connet Error",err)
	}
	return  redis
}

func SetCashe(key string,val interface{}, timeout time.Duration) *interface{} {
	jsonBytes, _ := json.Marshal(val)
	red :=MyCache()
	red.Put(key,string(jsonBytes),  timeout)
	return &val
}

func GetCashe(key string) []byte {
	red :=MyCache()
	val := red.Get(key)
	if val == nil {
		return nil
	}else{
		return  val.([]byte)
	}

}

func StrToBase64(str string) string {
	strbytes := []byte(str)

   return base64.StdEncoding.EncodeToString(strbytes)
}

func Base64ToStr(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Base64ToStr err",err)
	}
	return  string(decoded)
}

