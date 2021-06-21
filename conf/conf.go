package conf

import (
	"os"
	"sync"
)

var _conf map[string]interface{}

var lock sync.Mutex

func Load() {

	_conf = map[string]interface{}{

		"cookie_name":  os.Getenv("APP_NAME") + "_session", //浏览器cookie名称
		"cookie_key":   "cookie_key",                       //context中cookie的值的name
		"redis_prefix": "",
	}

}

func Get(key string) interface{} {

	lock.Lock()

	defer lock.Unlock()

	return _conf[key]

}
