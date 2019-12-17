package cloudforms

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceServiceTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceTemplateRead,

		Schema: map[string]*schema.Schema{

			// required values
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			// computed values
			"href": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_template_catalog_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// get catalogID
var serviceTemplateCatalogID string

// Get templateID to fetch Service_templates associated with it
var tmpID string

// structure to store template detail
var templateDetailstruct TemplateQuery

// dataSourceServiceTemplateRead performs service_template lookup
func dataSourceServiceTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	// Get index of template
	var index int

	templateName := d.Get("name").(string)
	if templateName == "" {
		return fmt.Errorf("name of template must be set for data source service_template")
	}

	log.Println("[DEBUG] Reading Service templates...")
	response, err := GetServiceTemplate(config, templateName)
	if err != nil {
		log.Printf("[ERROR] Error while getting response: %s", err)
		return fmt.Errorf("Error while getting response: %s", err)
	}

	// store service template
	if json.Unmarshal(response, &templateDetailstruct); err != nil {
		log.Printf("[Error] Error while unmarshalling json: %s", err)
		return fmt.Errorf("Error while unmarshalling json: %s", err)
	}

	// subcount is nothing but number of successful results
	if templateDetailstruct.Subcount == 0 {
		log.Printf("[DEBUG] Template called `%s` Not found in List ", templateName)
		return fmt.Errorf("Template called `%s` Not found in List", templateName)
	}

	// index of template from result
	for i := 0; i < templateDetailstruct.Subcount; i++ {
		if templateDetailstruct.Resources[i].Name == templateName {
			index = i
			tmpID = templateDetailstruct.Resources[i].ID
			log.Printf("[DEBUG] Template called `%s` found in List ", templateName)
			break
		}
	}

	// Set values into schema
	d.Set("href", templateDetailstruct.Resources[index].Href)
	d.Set("id", templateDetailstruct.Resources[index].ID)
	d.Set("name", templateDetailstruct.Resources[index].Name)
	d.Set("description", templateDetailstruct.Resources[index].Description)
	d.Set("tenant_id", templateDetailstruct.Resources[index].TenantID)
	d.Set("type", templateDetailstruct.Resources[index].Type)
	d.Set("service_template_catalog_id", templateDetailstruct.Resources[index].ServiceTemplateCatalogID)

	//	Calling SetId on our schema.ResourceData using a value suitable for resource.
	//	This ensures whatever resource state we set on schema.ResourceData will be persisted in local state.
	// 	If we neglect to SetId, no resource state will be persisted.
	d.SetId(fmt.Sprintf("%s", tmpID))

	return nil
}
