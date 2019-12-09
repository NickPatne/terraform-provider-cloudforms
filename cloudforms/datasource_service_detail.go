package cloudforms

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceServiceDetail() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServiceDetailRead,

		Schema: map[string]*schema.Schema{
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
			"name": {
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
			// Aggregate Type
			"service_templates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"href": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"guid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"miq_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Structure to store service catalogs
var serviceCatalogStruct ServiceCatalogs

// Structure to store service details
var tempalteListStruct ServiceDetails

// Get index of service
var index int

// Get ServiceID to fetch Service_templates associated with it
var serviceID string

func dataSourceServiceDetailRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	serviceName := d.Get("name").(string)

	log.Println("[DEBUG] Reading Service Catalog...")
	response, err := GetServiceCatalog(config)
	if err != nil {
		log.Printf("[ERROR] Error while getting response %s", err)
		return fmt.Errorf("Error while getting response ->\n %s", err)
	}

	//store service catalog
	if json.Unmarshal(response, &serviceCatalogStruct); err != nil {
		log.Printf("[Error] Error while unmarshal request json %s ", err)
		return fmt.Errorf("Error while unmarshal request json -> %s ", err)
	}

	if serviceCatalogStruct.Subcount == 0 {
		return fmt.Errorf("Service Catalog is empty")
	}

	// Checking whether this service is availabe
	for i := 0; i < serviceCatalogStruct.Subcount; i++ {
		if serviceCatalogStruct.Resources[i].Name == serviceName {
			index = i
			serviceID = serviceCatalogStruct.Resources[i].ID
			log.Printf("[DEBUG] Service called `%s` found in catalogs ", serviceName)
			break
		}
	}

	if serviceID == "" {
		return fmt.Errorf("Service called `%s` Not found in catalog", serviceName)
	}

	// Set values into schema
	d.Set("href", serviceCatalogStruct.Resources[index].Href)
	d.Set("id", serviceCatalogStruct.Resources[index].ID)
	d.Set("name", serviceCatalogStruct.Resources[index].Name)
	d.Set("description", serviceCatalogStruct.Resources[index].Description)
	d.Set("tenant_id", serviceCatalogStruct.Resources[index].TenantID)

	// Get list of service_templates
	templates, err := GetTemplateList(config, serviceID)
	if json.Unmarshal(templates, &tempalteListStruct); err != nil {
		log.Printf("[Error] Error while unmarshal request json %s ", err)
		return fmt.Errorf("Error while unmarshal request json -> \n %s ", err)
	}
	// calling helper function to settle aggregate type of schema
	d.Set("service_templates", FlattenServiceTemplate(tempalteListStruct.ServiceTemplates))

	//	Calling SetId on our schema.ResourceData using a value suitable for your resource.
	//	This ensures whatever resource state we set on schema.ResourceData will be persisted in local state.
	// 	If we neglect to SetId, no resource state will be persisted.
	d.SetId(fmt.Sprintf("%s", serviceID))

	return nil
}
