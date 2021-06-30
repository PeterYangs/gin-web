package queue

import (
	"gin-web/contextPlus"
	"gin-web/queue"
	"gin-web/queue/task/email"
	"github.com/gin-gonic/gin"
)

func Task(c *contextPlus.Context) interface{} {

	//fmt.Println(limiter.GlobalLimiters)

	//task,_:=json.Marshal(map[string]interface{}{"name":"email","data":"data"})

	//redis.GetClient().LPush(context.TODO(),"queue:default",task)

	queue.Dispatch(email.NewTask("904801074@qq.com", "标题", "内容"))

	return gin.H{"code": 1, "msg": "go"}
}
