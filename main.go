package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/GSLabDev/terraform-provider-cloudforms/cloudforms"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		//Call provider
		ProviderFunc: cloudforms.Provider})
}
