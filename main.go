package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-cloudforms/cloudforms"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//Call provider
		ProviderFunc: cloudforms.Provider})
}
