package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-web/queue/taskName"
	"gin-web/redis"
	"github.com/spf13/cast"
	"log"
	"runtime/debug"
	"sync"
)

var handles = sync.Map{}

func Run() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println(r)

			fmt.Println(string(debug.Stack()))

		}
	}()

	for {

		s, err := redis.GetClient().BRPop(context.TODO(), 0, "queue:default").Result()

		if err != nil {

			log.Println(err)

			fmt.Println("队列退出")

			break
		}

		var jsons map[string]interface{}

		//fmt.Println(s[1])

		err = json.Unmarshal([]byte(s[1]), &jsons)

		if err != nil {

			fmt.Println(err)

			continue
		}

		////获取task
		//h, ok := handles[jsons["TaskName"].(string)]
		hh, ok := handles.Load(jsons["TaskName"].(string))

		h := hh.(taskName.Task)

		if !ok {

			fmt.Println("获取task失败")

			continue
		}

		//绑定参数
		h.BindParameters(cast.ToStringMapString(jsons["Parameters"]))

		//执行任务
		h.Run()

	}

}

func Dispatch(task taskName.Task) {

	handles.Store(task.GetName(), task)

	t, _ := json.Marshal(task)

	redis.GetClient().LPush(context.TODO(), "queue:default", t)

}
