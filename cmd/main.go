package main

import (
	"github.com/senaxor/goproxyrotator/internal/config"
	"github.com/senaxor/goproxyrotator/internal/proxy"
	"github.com/senaxor/goproxyrotator/internal/requester"
)

func main() {
	cfg := config.Load()

	manager := proxy.NewManager(cfg.ProxyList)
	manager.FetchAndInit()

	requester.Run(cfg, manager)
}
