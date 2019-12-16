---
layout: "Cloudforms"
page_title: "Provider: Cloudforms"
sidebar_current: "docs-cloudforms-index"
description: |-
  The Cloudforms provider is used to interact with the resources supported by
  Redhat Cloudforms. The provider needs to be configured with the proper credentials
  before it can be used.
---

# Cloudforms Provider

The Cloudforms provider is used to interact with the resources supported by Redhat Cloudforms. 
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

~> **NOTE:** The Cloudforms Provider currently represents _initial support_
and therefore may undergo significant changes as the community improves it. This
provider at this time only supports adding Computer Resource

## Example Usage

```hcl

# Configure the Cloudform Provider
provider "cloudforms" {
	ip = "${var.CF_SERVER_IP}"
	user_name = "${var.CF_USER_NAME}"
	password = "${var.CF_PASSWORD}"
}

# Data Source cloudforms_service_template
data "cloudforms_service_template" "mytemplate"{
	name = "${var.CF_TEMPLATE_NAME}"
}

# Resource cloudforms_miq_request
resource "cloudforms_service_request" "test" {	
 	name = "${var.CF_TEMPLATE_NAME}"
 	template_href = "${data.cloudforms_service_template.mytemplate.href}"
 	catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
 	input_file_name = "${var.INPUT_FILE_NAME}"
 	time_out= 50
}
output "Service_templates_href"{
	value = "${data.cloudforms_service_template.mytemplate.href}"
}

```

## Argument Reference

The following arguments are used to configure the Cloudforms Provider:

* `user_name` - (Required) This is the username for ManageIQ user. Can also
  be specified with the `CF_USER_NAME` environment variable.
* `password` - (Required) This is the password for ManageIQ user. Can
  also be specified with the `CF_PASSWORD` environment variable.
* `ip` - (Required) This is the ManageIQ server ip. Can also be specified with the `CF_SERVER_IP` environment
  variable.

## Acceptance Tests

The Cloudforms provider's acceptance tests require the above provider
configuration fields to be set using the documented environment variables.

In addition, the following environment variables are used in tests, and must be
set to valid values for your Active Directory environment:

 * CF\_SERVICE\_NAME
 * CF\_TEMPLATE\_NAME
 * CF\_INPUT\_FILE\_NAME
 
 

Once all these variables are in place, the tests can be run like this:

```
make testacc

```
