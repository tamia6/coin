package middleware

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func CheckWhite(r *ghttp.Request) {
	appCfg := gmap.NewStrAnyMap()
	_ = gconv.Scan(g.Cfg().MustGet(r.GetCtx(), "app"), appCfg)
	whiteList := gset.NewStrSet()
	_ = gconv.Scan(appCfg.Get("whiteList"), &whiteList)
	if !whiteList.Contains(r.GetClientIp()) {
		r.Response.WriteStatusExit(403)
	}
	r.Middleware.Next()
}
