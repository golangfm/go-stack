package http_client

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestFastHttpV1(t *testing.T) {
	url := "https://www.baidu.com"
	var req fasthttp.Request
	req.SetRequestURI(url)
	var res fasthttp.Response
	var c fasthttp.Client

	err := c.Do(&req, &res)
	if err != nil {
		t.Fatal("send error:", err)
	}
	if res.StatusCode() != fasthttp.StatusOK {
		t.Fatal("code:", res.StatusCode())
	}
	t.Fatal(string(res.Body()))
}

func BenchmarkFastHttp(b *testing.B) {
	url := "https://www.baidu.com"
	var req fasthttp.Request
	req.SetRequestURI(url)
	var res fasthttp.Response
	var c fasthttp.Client
	cc := &http.Client{}

	b.Run("NGet", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			res, err := cc.Get(url)
			if err != nil {
				b.Fatal("send error:", err)
			}
			if res.StatusCode != http.StatusOK {
				b.Fatal("code:", res.StatusCode)
			}
			data, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				b.Fatal("read error.")
			}
			_ = data
		}
	})

	b.Run("Do", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := c.Do(&req, &res)
			if err != nil {
				b.Fatal("send error:", err)
			}
			if res.StatusCode() != fasthttp.StatusOK {
				b.Fatal("code:", res.StatusCode())
			}
			_ = res.Body()
		}
	})
	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			statusCode, body, err := c.Get(nil, url)
			if err != nil {
				b.Fatal("send error:", err)
			}
			if statusCode != fasthttp.StatusOK {
				b.Fatal("code:", statusCode)
			}
			_ = body
		}
	})
}
