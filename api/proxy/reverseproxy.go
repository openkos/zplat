package proxy

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/httputil"
)

func ReverseProxy(address string, headers map[string]string) gin.HandlerFunc {
	header := http.Header{}
	for k, v := range headers {
		header.Set(k, v)
	}

	return ReverseProxyHandler(address, header)
}

func ReverseProxyHandler(addr string, header http.Header) gin.HandlerFunc {
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("[upstream] invalid address: %s", err)
	}

	upstream := httputil.NewReverseProxy(u, header)
	return func(c *gin.Context) {
		upstream.ServeHTTP(c.Writer, c.Request)
	}
}
