package requester

import (
	"fmt"
	"sync"
	"time"

	"github.com/senaxor/goproxyrotator/internal/config"
	"github.com/senaxor/goproxyrotator/internal/proxy"
)

func Run(cfg config.Config, pm *proxy.ProxyManager) {
	var wg sync.WaitGroup
	targets := readTargets(cfg.TargetList)

	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for _, url := range targets {
				proxyAddr := pm.GetProxy()
				if proxyAddr == "" {
					time.Sleep(time.Second)
					continue
				}

				client := createHttpClient(proxyAddr)
				resp, err := client.Get(url)
				if err != nil {
					fmt.Println("Request failed:", err)
					pm.Cooldown(proxyAddr, cfg.Cooldown)
					continue
				}

				if resp.StatusCode == 429 {
					fmt.Println("Rate limited. Cooling down proxy", proxyAddr)
					pm.Cooldown(proxyAddr, cfg.Cooldown)
				} else {
					fmt.Println("Worker", id, "->", url, "status:", resp.StatusCode)
				}
				resp.Body.Close()
			}
		}(i)
	}

	wg.Wait()
}
