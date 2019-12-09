variable "MIQ_IP" {}
variable "USER_NAME" {}
variable "PASSWORD" {}
variable "SERVICE_NAME" {}

# Configure the Cloudform Provider
provider "cloudforms" {
	ip = "${var.MIQ_IP}"
	user_name = "${var.USER_NAME}"
	password = "${var.PASSWORD}"
}

# Data Source cloudforms_services
data  "cloudforms_services" "myservice"{
    name = "${var.SERVICE_NAME}"

}

# Resource cloudforms_miq_request
resource "cloudforms_miq_request" "test" {
 name = "${var.SERVICE_NAME}"
 input_file_name = "data.json"
 time_out= 50
}


output "Service_Name"{
	value = "${data.cloudforms_services.myservice.name}"
}

output "service_templates_href"{
	value = "${data.cloudforms_services.myservice.service_templates.*.href}"
}
