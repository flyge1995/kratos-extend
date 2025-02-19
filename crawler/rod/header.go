package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:133.0) Gecko/20100101 Firefox/133.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/119.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.1 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/106.0.0.0",
}

func NewHTTPService() *HTTPService {
	return &HTTPService{}
}

type HTTPService struct {
	next atomic.Int64
}

func (s *HTTPService) CreateHeaders(_url *url.URL) http.Header {
	return http.Header{
		"User-Agent":      []string{userAgents[int(s.next.Add(1))%len(userAgents)]},
		"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"},
		"Accept-Encoding": []string{"gzip, deflate"},
		"Accept-Language": []string{"zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3"},
		"Connection":      []string{"keep-alive"},
		"Referer":         []string{fmt.Sprintf("%s://%s", _url.Scheme, _url.Host)},
	}
}
