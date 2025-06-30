package site

type Site struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Config string `json:"config"`
}

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

type FtpSite struct {
	Id     int64         `json:"id"`
	Name   string        `json:"name"`
	Config FtpSiteConfig `json:"config"`
}

type FtpSiteConfig struct {
	Url          string `json:"url"`
	Port         int    `json:"port"`
	AuthRequired bool   `json:"authenticationRequired"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}

type NewFtpSite struct {
	Name   string        `json:"name"`
	Config FtpSiteConfig `json:"config"`
}
