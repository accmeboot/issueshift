package config

import (
	"flag"
	"os"
)

type Content struct {
	DSN  *string
	ADDR *string
}

func NewConfig() *Content {
	var dsn string
	var addr string

	readFlags(&dsn, &addr)

	return &Content{
		DSN:  &dsn,
		ADDR: &addr,
	}
}

func readFlags(dsn *string, addr *string) {
	defaultDsn := os.Getenv("DSN")
	defaultAddr := os.Getenv("ADDR")

	flag.StringVar(dsn, "dsn", defaultDsn, "Connection string to postgresql")
	flag.StringVar(addr, "addr", defaultAddr, "Api address (e.g localhost:4000)")
	flag.Parse()
}
