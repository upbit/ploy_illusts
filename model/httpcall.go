package model

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
)

// GetPixivImage 获取 Pixiv 的图片
func GetPixivImage(ctx context.Context, url string) ([]byte, string, error) {
	response, err := HTTPCall(ctx, "GET", url, map[string]string{"Referer": "https://app-api.pixiv.net/"}, nil)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, "", err
	}

	return buf, response.Header.Get("content-type"), nil
}

// HTTPCall HTTP请求封装
func HTTPCall(ctx context.Context, method string, url string, headers map[string]string,
	body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
