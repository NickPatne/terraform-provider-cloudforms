
provider "cloudforms" {
	ip = "${var.SERVER_IP}"
	user_name = "${var.USER_NAME}"
	password = "${var.PASSWORD}"
}

# datasource : data.cloudforms_service.myservice
 data  "cloudforms_service" "myservice"{
    name = "${var.SERVICE_NAME}"
 }

# datasource : data.cloudforms_service_template.mytemplate
data "cloudforms_service_template" "mytemplate"{
	name = "${var.TEMPLATE_NAME}"
}


# resource : cloudforms_service_request.test
resource "cloudforms_service_request" "test" {	
 	name = "${var.TEMPLATE_NAME}"
 	template_href = "${data.cloudforms_service_template.mytemplate.href}"
 	catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
 	input_file_name = "data.json"
 	time_out= 50
}	

# Output variables
output "Service_Name"{
	value = "${data.cloudforms_service.myservice.name}"
}

output "Service_catalogID"{
	value = "${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
}

output "Service_templates_href"{
	value = "${data.cloudforms_service_template.mytemplate.href}"
}
