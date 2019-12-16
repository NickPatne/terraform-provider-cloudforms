---
layout: "cloudforms"
page_title: "Active Directory: cloudforms_miq_reques"
sidebar_current: "docs-cloudforms-resource-inventory-folder"
description: |-
  Orders a service from catalog
---

# cloudforms\_miq\_request

Creates a group object in an Active Directory Organizational Unit.

## Example Usage

```hcl

# Resource cloudforms_miq_request
resource "cloudforms_miq_request" "test" {  
    name = "${var.TEMPLATE_NAME}"
    href = "${data.cloudforms_service_template.mytemplate.href}"
    catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
    input_file_name = "data.json"
    time_out= 50
}  

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the service template of service to order.
* `href` - (Required) The distinguished href of the service template of service to order.
* `catalog_id` - (Required) The distinguished id of service to which this template belongs.
* `input_file_name` - (Required) The input file which contains attributes of service.
* `time_out` - (Optional) Number of seconds to wait for timeout.

