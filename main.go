package main

import (
	"context"
	"flag"
	"log"

	"github.com/change-engine/terraform-provider-slack-app/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := tfsdk.ServeOpts{
		Debug: debug,

		Name: "registry.terraform.io/change-engine/slack-app",
	}

	err := tfsdk.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}