package cloudforms

import "time"

// Service Catalog information:
type ServiceCatalogJsonstruct struct {
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Subcount  int    `json:"subcount"`
	Pages     int    `json:"pages"`
	Resources []struct {
		Href             string `json:"href"`
		ID               string `json:"id"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		TenantID         string `json:"tenant_id"`
		ServiceTemplates struct {
			Count     int `json:"count"`
			Pages     int `json:"pages"`
			Resources []struct {
				Href string `json:"href"`
			} `json:"resources"`
			Actions []struct {
				Name   string `json:"name"`
				Method string `json:"method"`
				Href   string `json:"href"`
			} `json:"actions"`
			Links struct {
				Self  string `json:"self"`
				First string `json:"first"`
				Last  string `json:"last"`
			} `json:"links"`
		} `json:"service_templates"`
	} `json:"resources"`
	Actions []struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		Href   string `json:"href"`
	} `json:"actions"`
	Links struct {
		Self  string `json:"self"`
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
}

type requestJsonstruct struct {
	Results []Results `json:"results"`
}

// Dialog is interface for dialog parameters present in input template.

type Dialog interface {
}

type WorkflowSettings struct {
	ResourceActionID string `json:"resource_action_id"`
	DialogID         string `json:"dialog_id"`
}
type RequestOptions struct {
	SubmitWorkflow bool `json:"submit_workflow"`
	InitDefaults   bool `json:"init_defaults"`
}
type Options struct {
	Dialog           Dialog           `json:"dialog"`
	WorkflowSettings WorkflowSettings `json:"workflow_settings"`
	Initiator        interface{}      `json:"initiator"`
	SrcID            string           `json:"src_id"`
	RequestOptions   RequestOptions   `json:"request_options"`
	CartState        string           `json:"cart_state"`
	RequesterGroup   string           `json:"requester_group"`
}

// Struct for result obtain after ordering service.

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
type Resource interface {
}
