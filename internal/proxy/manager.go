package proxy

import (
	"sync"
	"time"
)

type ProxyManager struct {
	proxies   []string
	cooldowns map[string]time.Time
	mu        sync.Mutex
}

func NewManager(proxies []string) *ProxyManager {
	return &ProxyManager{
		proxies:   proxies,
		cooldowns: make(map[string]time.Time),
	}
}

func (pm *ProxyManager) GetProxy() string {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	now := time.Now()
	for _, p := range pm.proxies {
		if cd, ok := pm.cooldowns[p]; !ok || cd.Before(now) {
			return p
		}
	}
	return ""
}

func (pm *ProxyManager) Cooldown(proxy string, seconds int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.cooldowns[proxy] = time.Now().Add(time.Duration(seconds) * time.Second)
}
