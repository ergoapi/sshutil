package sshutil

import "time"

type SSH struct {
	User string
	Pass string
	PkFile string
	PkPass string
	Timeout *time.Duration
	Debug bool
}