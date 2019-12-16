package cloudforms

import "time"

//===================================================================================
// Structure for cloudfroms resource
//===================================================================================

type requestJsonstruct struct {
	Results []Results `json:"results"`
}

// Dialog : this is an interface for dialog parameters present in input template.
type Dialog interface {
}

// WorkflowSettings :
type WorkflowSettings struct {
	ResourceActionID string `json:"resource_action_id"`
	DialogID         string `json:"dialog_id"`
}

//RequestOptions :
type RequestOptions struct {
	SubmitWorkflow bool `json:"submit_workflow"`
	InitDefaults   bool `json:"init_defaults"`
}

// Options :
type Options struct {
	Dialog           Dialog           `json:"dialog"`
	WorkflowSettings WorkflowSettings `json:"workflow_settings"`
	Initiator        interface{}      `json:"initiator"`
	SrcID            string           `json:"src_id"`
	RequestOptions   RequestOptions   `json:"request_options"`
	CartState        string           `json:"cart_state"`
	RequesterGroup   string           `json:"requester_group"`
}

// Results : Struct for result obtain after ordering service.
type Results struct {
	Href              string      `json:"href"`
	ID                string      `json:"id"`
	Description       string      `json:"description"`
	ApprovalState     string      `json:"approval_state"`
	Type              string      `json:"type"`
	CreatedOn         time.Time   `json:"created_on"`
	UpdatedOn         time.Time   `json:"updated_on"`
	FulfilledOn       interface{} `json:"fulfilled_on"`
	RequesterID       string      `json:"requester_id"`
	RequesterName     string      `json:"requester_name"`
	RequestType       string      `json:"request_type"`
	RequestState      string      `json:"request_state"`
	Message           string      `json:"message"`
	Status            string      `json:"status"`
	Options           Options     `json:"options"`
	Userid            string      `json:"userid"`
	SourceID          string      `json:"source_id"`
	SourceType        string      `json:"source_type"`
	DestinationID     interface{} `json:"destination_id"`
	DestinationType   interface{} `json:"destination_type"`
	TenantID          string      `json:"tenant_id"`
	ServiceOrderID    string      `json:"service_order_id"`
	Process           bool        `json:"process"`
	CancelationStatus interface{} `json:"cancelation_status"`
}

//Struct for input json template.

type template struct {
	Action   string   `json:"action"`
	Resource Resource `json:"resource"`
}

// Resource : struct to store attribute of service
type Resource interface {
}

//===================================================================================
// Structure for cloudforms data source
//===================================================================================

// ServiceDetails : Struct for ServiceDetails expand = service_templates
type ServiceDetails struct {
	Href             string             `json:"href"`
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	TenantID         string             `json:"tenant_id"`
	ServiceTemplates []ServiceTemplates `json:"service_templates"`
}

// Options1 :
type Options1 struct {
}

// Resources1 : Resources under Service_templates
type Resources1 struct {
	Href                     string      `json:"href"`
	ID                       string      `json:"id"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	GUID                     string      `json:"guid"`
	Type                     string      `json:"type"`
	ServiceTemplateID        interface{} `json:"service_template_id"`
	Options                  Options1    `json:"options"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Display                  bool        `json:"display"`
	EvmOwnerID               interface{} `json:"evm_owner_id"`
	MiqGroupID               string      `json:"miq_group_id"`
	ServiceType              string      `json:"service_type"`
	ProvType                 string      `json:"prov_type"`
	ProvisionCost            interface{} `json:"provision_cost"`
	ServiceTemplateCatalogID string      `json:"service_template_catalog_id"`
	LongDescription          string      `json:"long_description"`
	TenantID                 string      `json:"tenant_id"`
	GenericSubtype           interface{} `json:"generic_subtype"`
	DeletedOn                interface{} `json:"deleted_on"`
	Internal                 bool        `json:"internal"`
}

// ServiceTemplates : From expand = Service_templates
type ServiceTemplates struct {
	Href                     string      `json:"href"`
	ID                       string      `json:"id"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	GUID                     string      `json:"guid"`
	Type                     string      `json:"type"`
	ServiceTemplateID        interface{} `json:"service_template_id"`
	Options                  Options1    `json:"options"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Display                  bool        `json:"display"`
	EvmOwnerID               interface{} `json:"evm_owner_id"`
	MiqGroupID               string      `json:"miq_group_id"`
	ServiceType              string      `json:"service_type"`
	ProvType                 string      `json:"prov_type"`
	ProvisionCost            interface{} `json:"provision_cost"`
	ServiceTemplateCatalogID string      `json:"service_template_catalog_id"`
	LongDescription          string      `json:"long_description"`
	TenantID                 string      `json:"tenant_id"`
	GenericSubtype           interface{} `json:"generic_subtype"`
	DeletedOn                interface{} `json:"deleted_on"`
	Internal                 bool        `json:"internal"`
}

// ServiceCatalogs : This is struct for service catalogs
type ServiceCatalogs struct {
	Name      string      `json:"name"`
	Count     int         `json:"count"`
	Subcount  int         `json:"subcount"`
	Pages     int         `json:"pages"`
	Resources []Resources `json:"resources"`
	Actions   []Actions   `json:"actions"`
	Links     Links       `json:"links"`
}

// Resources11 ... under ServiceTemplates
type Resources11 struct {
	Href string `json:"href"`
}

// Actions : Generic Actions
type Actions struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	Href   string `json:"href"`
}

// Links : Generic Links
type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// ServiceTemplatesUndercatalogs : From expand = resources
type ServiceTemplatesUndercatalogs struct {
	Count     int           `json:"count"`
	Pages     int           `json:"pages"`
	Resources []Resources11 `json:"resources"`
	Actions   []Actions     `json:"actions"`
	Links     Links         `json:"links"`
}

// Resources : parent resource which contains hidden service template
type Resources struct {
	Href             string                        `json:"href"`
	ID               string                        `json:"id"`
	Name             string                        `json:"name"`
	Description      string                        `json:"description"`
	TenantID         string                        `json:"tenant_id"`
	ServiceTemplates ServiceTemplatesUndercatalogs `json:"service_templates"`
}

// TemplateDetail :
type TemplateDetail struct {
	Href                     string      `json:"href"`
	ID                       string      `json:"id"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	GUID                     string      `json:"guid"`
	Type                     string      `json:"type"`
	ServiceTemplateID        interface{} `json:"service_template_id"`
	Options                  Options1    `json:"options"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Display                  bool        `json:"display"`
	EvmOwnerID               interface{} `json:"evm_owner_id"`
	MiqGroupID               string      `json:"miq_group_id"`
	ServiceType              string      `json:"service_type"`
	ProvType                 string      `json:"prov_type"`
	ProvisionCost            interface{} `json:"provision_cost"`
	ServiceTemplateCatalogID string      `json:"service_template_catalog_id"`
	LongDescription          string      `json:"long_description"`
	TenantID                 string      `json:"tenant_id"`
	GenericSubtype           interface{} `json:"generic_subtype"`
	DeletedOn                interface{} `json:"deleted_on"`
	Internal                 bool        `json:"internal"`
	ConfigInfo               ConfigInfo  `json:"config_info"`
	Actions                  []Actions   `json:"actions"`
}

// Provision :
type Provision struct {
	DialogID string `json:"dialog_id"`
	Fqname   string `json:"fqname"`
}

// ConfigInfo :
type ConfigInfo struct {
	ConfigurationScriptID string    `json:"configuration_script_id"`
	Provision             Provision `json:"provision"`
}

// ResponseError : Structure to handle response error
type ResponseError struct {
	Error Error `json:"error"`
}

// Error : Contains kind, message,klass
type Error struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
	Klass   string `json:"klass"`
}

// -----------------------------------------

// TemplateQuery : Struct for filtered response of template
type TemplateQuery struct {
	Name          string     `json:"name"`
	Count         int        `json:"count"`
	Subcount      int        `json:"subcount"`
	SubqueryCount int        `json:"subquery_count"`
	Pages         int        `json:"pages"`
	Resources     []Template `json:"resources"`
	Actions       []Actions  `json:"actions"`
	Links         Links      `json:"links"`
}

// Template :
type Template struct {
	Href                     string      `json:"href"`
	ID                       string      `json:"id"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	GUID                     string      `json:"guid"`
	Type                     string      `json:"type"`
	ServiceTemplateID        interface{} `json:"service_template_id"`
	Options                  Options1    `json:"options"`
	CreatedAt                time.Time   `json:"created_at"`
	UpdatedAt                time.Time   `json:"updated_at"`
	Display                  bool        `json:"display"`
	EvmOwnerID               interface{} `json:"evm_owner_id"`
	MiqGroupID               string      `json:"miq_group_id"`
	ServiceType              string      `json:"service_type"`
	ProvType                 string      `json:"prov_type"`
	ProvisionCost            interface{} `json:"provision_cost"`
	ServiceTemplateCatalogID string      `json:"service_template_catalog_id"`
	LongDescription          string      `json:"long_description"`
	TenantID                 string      `json:"tenant_id"`
	GenericSubtype           interface{} `json:"generic_subtype"`
	DeletedOn                interface{} `json:"deleted_on"`
	Internal                 bool        `json:"internal"`
	Actions                  []Actions   `json:"actions"`
}
