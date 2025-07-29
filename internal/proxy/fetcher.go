package proxy

func FetchFromPublicSources() []string {
	// Later: use HTTP client to fetch from websites
	return []string{
		"185.104.252.10:8080",
		"209.141.54.142:3128",
	}
}
