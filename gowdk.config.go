package config

import "github.com/cssbruno/gowdk"

var Config = gowdk.Config{
	AppName: "GOWDK Simple Login",
	Source: gowdk.SourceConfig{
		Include: []string{"src/**/*.gwdk"},
	},
	Build: gowdk.BuildConfig{
		Targets: []gowdk.BuildTargetConfig{{
			Name:   "gowdk-simple-login",
			App:    ".gowdk/gowdk-simple-login",
			Binary: "bin/gowdk-simple-login",
		}},
	},
	CSS: gowdk.CSSConfig{
		Include: []string{"styles/**/*.css"},
		Output: gowdk.CSSOutputConfig{
			Dir:        ".",
			HrefPrefix: "/",
		},
	},
}
