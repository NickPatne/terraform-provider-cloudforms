package cloudforms

import (
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider : Defines provider schema
// Contains registry of Data sources and Resources
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{

			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP of Cloudforms service",
				DefaultFunc: schema.EnvDefaultFunc("CF_SERVER_IP", nil),
			},

			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The UserName of ManageIQ service",
				DefaultFunc: schema.EnvDefaultFunc("CF_USER_NAME", nil),
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The Password of ManageIQ service",
				DefaultFunc: schema.EnvDefaultFunc("CF_PASSWORD", nil),
			},
		},

		//Supported Data Source by this provider
		DataSourcesMap: map[string]*schema.Resource{
			"cloudforms_service":          dataSourceServiceDetail(),
			"cloudforms_service_template": dataSourceServiceTemplate(),
		},
		//Supported Resources by this provider
		ResourcesMap: map[string]*schema.Resource{
			"cloudforms_service_request": resourceServiceRequest(),
		},
		ConfigureFunc: providerConfigure,
	}
}

// providerConfigure : This funtion will read provider module form '.tf' file store data into config structure
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config, err := CFConnect(d)
	if err != nil {
		log.Println("[ERROR] Failed to Establish Connection")
		os.Exit(1)
	}
	log.Println("[DEBUG] Connecting to Cloudforms...")
	return config, nil
}
