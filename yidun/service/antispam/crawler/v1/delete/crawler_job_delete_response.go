package delete

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// CrawlerJobDeleteV1Response 网站任务检测批量查询接口响应对象v1.0
type CrawlerJobDeleteV1Response struct {
	*types.CommonResponse
	Result *CrawlerJobDeleteV1Result `json:"result,omitempty"`
}

// CrawlerJobDeleteV1Result
type CrawlerJobDeleteV1Result struct {
	// 本次查询的任务id
	DeletedJobIds *[]int64 `json:"deletedJobIds,omitempty"`
}
