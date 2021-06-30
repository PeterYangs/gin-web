package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"gin-web/queue/task/email"
	"gin-web/queue/taskName"
	"gin-web/redis"
	"github.com/spf13/cast"
	"log"
	"runtime/debug"
)

type tasks struct {
	task     taskName.Task     //处理器
	taskName taskName.TaskName //任务名称
}

type data struct {
	TaskName   string
	Parameters map[string]string `json:"parameters"`
}

func Run() {

	defer func() {
		if r := recover(); r != nil {

			fmt.Println(r)

			fmt.Println(string(debug.Stack()))

		}
	}()

	//t:=taskName.Task(email.TaskEmail{})

	handles := map[string]tasks{
		"email": tasks{
			task:     &email.TaskEmail{Parameters: &email.Parameter{}},
			taskName: email.Name,
		},
	}

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
		h, ok := handles[jsons["TaskName"].(string)]

		if !ok {

			fmt.Println("获取task失败")

			continue
		}

		h.task.BindParameters(cast.ToStringMapString(jsons["Parameters"]))

		////执行任务
		h.task.Run()

	}

}

func Dispatch(task interface{}) {

	t, _ := json.Marshal(task)

	redis.GetClient().LPush(context.TODO(), "queue:default", t)

}
