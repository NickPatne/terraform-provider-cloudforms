---
layout: "cloudforms"
page_title: "Cloudforms: cloudforms_service_template"
sidebar_current: "docs-cloudforms-resource-inventory-folder"
description: |-
  Get details of available cloudforms service templates
---

# cloudforms\_service

Provides details of service templates available in service catalog using their name

## Example Usage

```hcl

# Data Source cloudforms_service_template
data "cloudforms_service_template" "mytemplate"{
    name = "${var.SERVICE_TEMPLATE_NAME}"
}

output "Service_templates_href"{
    value = "${data.cloudforms_service_template.mytemplate.href}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The distinguished name of the target service template .
