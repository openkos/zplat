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
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saltbo/gopkg/ginutil"

	"github.com/zplat-core/zplat/api/proxy"
	"github.com/zplat-core/zplat/assets"
)

// 这里做一个鉴权代理，默认挂载主服务的路由，通过配置可以挂载其他服务或静态文件
func proxyRun() {
	ge := gin.Default()
	ge.StaticFS("/zplat", assets.NewFS())
	ge.Any("/api/*action", proxy.ReverseProxyHandler("http://localhost:8218", http.Header{}))
	ge.NoRoute(func(c *gin.Context) {
	}, func(c *gin.Context) {
		// fetch routes
		// 匹配设置的路由
		// 判断是否需要鉴权，如果是则进行鉴权
		proxy.ReverseProxy("", map[string]string{})(c)
	})

	//jwtutil.Init()
	ginutil.SetupPing(ge)
	ginutil.Startup(ge, ":8216")
}
