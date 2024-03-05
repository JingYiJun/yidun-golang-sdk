package delete

import (
	"encoding/json"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// WeiboDeleteV1Request 网站任务检测批量查询接口v1.0
type WeiboDeleteV1Request struct {
	*types.BizPostFormRequest
	// 需要查询的任务id列表
	JobIds *[]int64 `json:"jobIds,omitempty"`
}

// For ImageV5CheckRequest  构造请求
func NewWeiboDeleteV1Request() *WeiboDeleteV1Request {
	var request = &WeiboDeleteV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/weibo-job/delete")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (j *WeiboDeleteV1Request) SetJobIds(jobIds []int64) {
	j.JobIds = &jobIds
}

// 获取参与签名的参数
func (c *WeiboDeleteV1Request) GetBusinessCustomSignParams() map[string]string {
	result := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.JobIds != nil {
		jobIdsJson, _ := json.Marshal(c.JobIds)
		result["jobIds"] = strings.Trim(string(jobIdsJson), "[]")
	}

	return result
}

// 参数校验
func (c *WeiboDeleteV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "WeiboDeleteV1Request"}
	if c.JobIds != nil {
		if len(*c.JobIds) == 0 {
			invalidParams.Add(validation.NewErrParamRequired("任务Ids不能为空"))
		}
		if len(*c.JobIds) > 100 {
			invalidParams.Add(validation.NewErrParamRequired("任务Ids个数不能超过100"))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
