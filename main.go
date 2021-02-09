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
	c.AddFunc("0 0 0 1 1 *", func() {
		SendNotify("test")
	})
	c.Start()
	r := gin.Default()

	r.GET("/message/:data", func(c *gin.Context) {
		req := c.Param("data")
		str := SendNotify(req)
		c.JSON(200, gin.H{
			"message": str,
		})
	})
	r.Run(":8088")

}
func SendNotify(req string) (str string) {
	url := "https://notify-api.line.me/api/notify"
	queries := map[string]string{
		"message": req,
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
