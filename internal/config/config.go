package config

type Config struct {
	ProxyList    string
	TargetList   string
	Concurrency  int
	Cooldown     int // in seconds
}

func Load() Config {
	// Read flags or env vars
	// You can use the "flag" or "cobra" package
	return Config{
		ProxyList:    "proxy_list.txt",
		TargetList:   "urls.txt",
		Concurrency:  50,
		Cooldown:     3,
	}
}
