package main

import (
	"context"
	"flag"

	"git-pong-go/api"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)

}

type Args struct {
	Port int
}

func parseArgs() Args {
	var args Args

	flag.IntVar(&args.Port, "port", 1337, "Server port (1337))")

	flag.Parse()

	return args
}

func main() {
	ctx := context.Background()
	args := parseArgs()

	var api api.Api
	api.Run(ctx, args.Port)
}
