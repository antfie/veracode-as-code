package main

// https://sfxpt.wordpress.com/2015/06/30/go-yaml-parsing-example/
type Application struct {
	Name   string `yaml:"name" json:"name"`
	Policy string `yaml:"policy" json:"policy"`
}

type Config struct {
	Application          Application `yaml:"application" json:"application"`
	Sandbox              string      `yaml:"sandbox" json:"sandbox"`
	Name                 string      `yaml:"name" json:"name"`
	AutoScanAfterPrescan bool        `yaml:"autoScanAfterPrescan" json:"autoScanAfterPrescan"`
	Uploads              []string    `yaml:"uploads" json:"uploads"`
}
