package main

import (
	_ "go-websocket-test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-websocket-test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
