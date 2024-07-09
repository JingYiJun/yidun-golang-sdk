package image

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
	gohttp "net/http"
	"net/url"
)

type ImageClient struct {
	Client client.ExecuteClient
}

func NewImageClient(profile *client.ClientProfile) *ImageClient {
	return &ImageClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewImageClientWithAccessKey(accessKeyId string, accessKeySecret string) *ImageClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewImageClient(client.NewClientProfile(credential))
}

func NewImageClientWithAccessKeyWithProxy(accessKeyId string, accessKeySecret string, proxy func(*gohttp.Request) (*url.URL, error)) *ImageClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	profile := client.NewClientProfile(credential)
	profile.HttpClientConfig.Proxy = proxy
	return NewImageClient(profile)
}
