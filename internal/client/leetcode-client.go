package client

import (
	"io"
	"net/http"

	"github.com/mikekbnv/auto-leetcode/internal/utils"
)

type LeetcodeHttpClient struct {
	httpClient *http.Client
	csrf_token string
	jwt_token  string
}

func NewLeetcodeHttpClient(csrf_token, jwt_token string) *LeetcodeHttpClient {
	client := &LeetcodeHttpClient{
		http.DefaultClient,
		csrf_token,
		jwt_token,
	}

	return client
}

func (c *LeetcodeHttpClient) Post(url, referer, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Cookie", buildCookie(utils.Keyvalue{Key: "csrftoken", Value: c.csrf_token}, utils.Keyvalue{Key: "LEETCODE_SESSION", Value: c.jwt_token}))
	req.Header.Add("X-Csrftoken", c.csrf_token)
	req.Header.Add("referer", referer)

	return c.httpClient.Do(req)
}

func (c *LeetcodeHttpClient) Get(url,contentType string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("Cookie", buildCookie(utils.Keyvalue{Key: "csrftoken", Value: c.csrf_token}, utils.Keyvalue{Key: "LEETCODE_SESSION", Value: c.jwt_token}))
	req.Header.Add("X-Csrftoken", c.csrf_token)
	//req.Header.Add("referer", referer)

	return c.httpClient.Do(req)
}

func buildCookie(pairs ...utils.Keyvalue) string {
	cookie := ""

	for _, pair := range pairs {
		cookie += pair.Key + "=" + pair.Value + "; "
	}
	return cookie
}
