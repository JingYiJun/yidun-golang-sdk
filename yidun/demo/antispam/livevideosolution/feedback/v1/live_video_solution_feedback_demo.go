package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/feedback/v1/request"
	"log"
	"os"
)

func main() {
	// Create a AntispamRequester instance, the parameters need to pass in the secretId, secretKey distributed by Antispam.

	YourSecretId := os.Getenv("SECRET_ID")
	YourSecretKey := os.Getenv("SECRET_KEY")

	request := request2.NewLiveWallSolutionFeedbackV1Req()
	// 设置查询开始时间和结束时间
	feedbacks := make([]*request2.LiveWallSolutionFeedback, 1)
	feedback := &request2.LiveWallSolutionFeedback{}
	feedback.SetTaskId("YourTaskId")
	feedback.SetViewCount(100)
	feedbacks[0] = feedback
	request.SetFeedbacks(feedbacks)

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := livevideosolution.NewLiveVideoSolutionClientWithAccessKey(YourSecretId, YourSecretKey)

	response, err := client.Feedback(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		data := response.Result
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			fmt.Println("转换为JSON时出错:", err)
			return
		}

		jsonStr := string(jsonBytes)
		fmt.Println(jsonStr)

	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
