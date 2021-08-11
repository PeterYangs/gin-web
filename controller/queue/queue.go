package queue

import (
	"gin-web/contextPlus"
	"gin-web/queue"
	"gin-web/queue/task/email"
	"gin-web/queue/task/sms"
	"gin-web/response"
)

func Task(c *contextPlus.Context) *response.Response {

	for i := 0; i < 100; i++ {

		queue.Dispatch(email.NewTask("904801074@qq.com", "标题", "内容"))
	}

	return response.Resp().String("gg")
}

func Task2(c *contextPlus.Context) *response.Response {

	for i := 0; i < 100; i++ {

		queue.Dispatch(sms.NewTask("110", "123"))
	}

	return response.Resp().String("gg")
}
