package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/wujiangweiphp/go-curl"
	"log"
)

func main() {
	fmt.Println("start crawler...")

	c := cron.New()
	c.AddFunc("*/10 * * * * *", func() {
		SendNotify()
	})
	c.Start()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.Run(":8088")

}
func SendNotify() (str string) {
	url := "https://notify-api.line.me/api/notify"
	queries := map[string]string{
		"message": "test",
	}
	post := map[string]interface{}{}
	str, err := HttpPost(url, queries, post)
	if err != nil {
		fmt.Println("failed")
	}
	return str

}

func HttpPost(url string, queries map[string]string, postData map[string]interface{}) (string, error) {
	headers := map[string]string{
		//"User-Agent":    "Sublime",
		"Authorization": "Bearer kX8qgUwx4SNAgzGM5CmYKjaheMyJayhbiMhxXBREpOG",
		"Content-Type":  "application/json",
	}
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).SetHeaders(headers).SetQueries(queries).SetPostData(postData).Post()

	if err != nil {
		return "", err
	} else {
		if resp.IsOk() {
			return resp.Body, nil
		} else {
			log.Fatalf("%s\n", resp.Raw)
			return "", errors.New("request Failed")
		}
	}
}
