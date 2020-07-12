package main

import (
	cli "github.com/rancher/wrangler-cli"
	"github.com/rancher/wrangler-cli/example/pkg/app"
)

func main() {
	cli.Main(app.New())
}
