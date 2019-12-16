---
layout: "cloudforms"
page_title: "Active Directory: cloudforms_service_request"
sidebar_current: "docs-cloudforms-resource-inventory-folder"
description: |-
  Orders a service from service catalog
---

# cloudforms\_service\_request

Creates a group object in an Active Directory Organizational Unit.

## Example Usage

```hcl

# Data Source cloudforms_service_template
data "cloudforms_service_template" "mytemplate"{
	name = "${var.CF_TEMPLATE_NAME}"
}

# Resource cloudforms_service_request
resource "cloudforms_service_request" "test" {  
    name = "${var.CF_TEMPLATE_NAME}"
    template_href = "${data.cloudforms_service_template.mytemplate.href}"
    catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
    input_file_name = "${var.CF_INPUT_FILE_NAME}"
    time_out= 50
}  

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the service template of service to order.
* `template_href` - (Required) The distinguished href of the service template of service to order.
* `catalog_id` - (Required) The distinguished id of service to which this template belongs.
* `input_file_name` - (Required) The input file which contains attributes of service.
* `time_out` - (Optional) Number of seconds to wait for timeout.

