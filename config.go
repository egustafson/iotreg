package main

type Config struct {
	SN  string            `json:"sn"`
	Ver uint              `json:"ver"`
	Cfg map[string]string `json:"cfg"`
}

type Registration struct {
	Config
	ID string `json:"id"`
	TS string `json:"ts"`
}
