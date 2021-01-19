/*
Copyright © 2020 Ambor <saltbo@foxmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"
	"github.com/zplat-core/apiserver/client"

	"github.com/zplat-core/zplat/api/proxy"
	"github.com/zplat-core/zplat/assets"
)

var subsystems = make(map[string]client.ModelSubsystem)

func init() {
	go func() {
		for range time.Tick(time.Second * 2) {
			cfg := client.NewConfiguration()
			cfg.BasePath = "http://localhost:8218/api"
			cli := client.NewAPIClient(cfg)
			ret, _, err := cli.V1SubsystemApi.V1SubsystemsGet(context.Background())
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, subsystem := range ret.Data.List {
				subsystems[subsystem.Name] = subsystem
			}
			fmt.Println(subsystems)
		}
	}()
}

// 这里做一个鉴权代理，默认挂载主服务的路由，通过配置可以挂载其他服务或静态文件
func proxyRun() {
	ge := gin.Default()
	ge.StaticFS("/zplat", assets.NewFS())
	ge.Any("/api/*action", matchSubsystemAPI(), proxy.ReverseProxyHandler("http://localhost:8218", http.Header{}))
	ge.NoRoute(noRouteHandler)

	//jwtutil.Init()
	ginutil.SetupPing(ge)
	ginutil.Startup(ge, ":8216")
}

func matchSubsystemAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		ssName := c.GetHeader("X-Zplat-Subsystem")
		if ssName != "" {
			proxy.ReverseProxyHandler(subsystems[ssName].Address, http.Header{})(c)
			c.Abort()
			return
		}
	}
}

func noRouteHandler(c *gin.Context) {
	defaultSubsystemName := "zpan"
	subsystemName := findSubsystemName(c)
	// 匹配设置的路由
	subsystem, ok := subsystems[subsystemName]
	if !ok {
		subsystem = subsystems[defaultSubsystemName]
	}

	// todo 判断是否需要鉴权，如果是则进行鉴权

	proxy.ReverseProxy(subsystem.Address, map[string]string{})(c)
}

func findSubsystemName(c *gin.Context) string {
	ssName := c.GetHeader("X-Zplat-Subsystem")
	if ssName != "" {
		return ssName
	}

	reqPath := strings.TrimPrefix(c.Request.URL.Path, "/")
	secondSlashIdx := strings.Index(reqPath, "/")
	if secondSlashIdx > -1 {
		return reqPath[:strings.Index(reqPath, "/")]
	}

	return ""
}
