package site

type HttpSite struct {
	Id     int64          `json:"id"`
	Name   string         `json:"name"`
	Config HttpSiteConfig `json:"config"`
}

type HttpSiteConfig struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type NewHttpSite struct {
	Name   string         `json:"name"`
	Config HttpSiteConfig `json:"config"`
}
