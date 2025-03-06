package main

import (
	"ccarbon-coin/internal/cmd"
	_ "ccarbon-coin/internal/logic"
	_ "ccarbon-coin/internal/packed"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
