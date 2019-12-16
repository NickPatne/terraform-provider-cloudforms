---
layout: "cloudforms"
page_title: "Cloudforms: cloudforms_service"
sidebar_current: "docs-cloudforms-resource-inventory-folder"
description: |-
  Get details of available cloudforms services
---

# cloudforms\_service\_template

Provides details of services available in service catalog using their name

## Example Usage

```hcl

# Data Source cloudforms_service
data  "cloudforms_service" "myservice"{
    name = "${var.SERVICE_NAME}"
}


output "Service_templates_href"{
    value = "${data.cloudforms_service.myservice.*.href}"
}


```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the target service .
