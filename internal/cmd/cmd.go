package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"

	"go-websocket-test/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			s.BindHandler("/socket", func(r *ghttp.Request) {
				var ctx = r.Context()
				ws, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}
				for {
					msgType, msg, err := ws.ReadMessage()
					if err != nil {
						return
					}
					if err = ws.WriteMessage(msgType, msg); err != nil {
						return
					}
				}
			})
			s.SetServerRoot("./resource/public/html")
			s.Run()
			return nil
		},
	}
)
