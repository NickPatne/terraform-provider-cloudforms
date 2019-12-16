
provider "cloudforms" {
	ip = "${var.CF_SERVER_IP}"
	user_name = "${var.CF_USER_NAME}"
	password = "${var.CF_PASSWORD}"
}

# Data source to display Service details
data  "cloudforms_service" "myservice"{
    name = "${var.SERVICE_NAME}"
}

# Data source to display Service template details
data "cloudforms_service_template" "mytemplate"{
	name = "${var.TEMPLATE_NAME}"
}

# Resource to order service from catalog
resource "cloudforms_miq_request" "test" {	
	name = "${var.TEMPLATE_NAME}"
	href = "${data.cloudforms_service_template.mytemplate.href}"
	catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
	input_file_name = "data.json"
	time_out= 50
}	


output "Service_Name"{
	value = "${data.cloudforms_service.myservice.name}"
}

output "Service_catalogID"{
	value = "${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
}

output "Service_templates_href"{
	value = "${data.cloudforms_service_template.mytemplate.href}"
}