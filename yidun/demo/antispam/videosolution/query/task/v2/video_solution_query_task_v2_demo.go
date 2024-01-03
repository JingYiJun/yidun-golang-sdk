package main

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/query/v2/request"
	"log"
	"os"
)

func main() {
	YourSecretId := os.Getenv("VS_SECRET_ID")
	YourSecretKey := os.Getenv("VS_SECRET_KEY")
	request := request3.NewVideoSolutionQueryTaskV2Request()

	// 实例化一个 Client，入参需要传入易盾内容安全分配的secretId，secretKey
	client := videosolution.NewVideoSolutionClientWithAccessKey(YourSecretId, YourSecretKey)
	request.SetTaskIds([]string{"xw9aez6luaaq7ochsvb1720h0200a4jj"})
	response, err := client.QueryTaskV2(request)
	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}

	if response.GetCode() == 200 {
		result, _ := json.Marshal(response.Result)
		fmt.Println("result: ", string(result))
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
