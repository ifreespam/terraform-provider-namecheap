package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/namecheap/terraform-provider-namecheap/namecheap"
)

func main() {
	// I did something and it works 123
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: namecheap.Provider,
	})
}
