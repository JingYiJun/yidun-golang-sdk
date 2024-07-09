package v5

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
	gohttp "net/http"
	"net/url"
)

type TextClient struct {
	Client client.ExecuteClient
}

// NewTextClient creates a new client for Text.
func NewTextClient(profile *client.ClientProfile) *TextClient {
	return &TextClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewTextClientWithAccessKey creates a new client for Text
func NewTextClientWithAccessKey(accessKeyId string, accessKeySecret string) *TextClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewTextClient(client.NewClientProfile(credential))
}

// NewTextClientWithAccessKeyWithProxy creates a new client for Text with proxy
func NewTextClientWithAccessKeyWithProxy(accessKeyId string, accessKeySecret string, proxy func(*gohttp.Request) (*url.URL, error)) *TextClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	profile := client.NewClientProfile(credential)
	profile.HttpClientConfig.Proxy = proxy
	return NewTextClient(profile)
}
