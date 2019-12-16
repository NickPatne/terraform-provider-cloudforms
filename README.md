# Terraform Cloudform Provider

This is the repository for the Terraform Cloudform Provider, which one can use
with Terraform to work with Cloudform.

For general information about Terraform, visit the [official website][3] and the
[GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform


# Using the Provider

The current version of this provider requires Terraform v0.12.9 or higher to
run.

Note that you need to run `terraform init` to fetch the provider before
deploying. Read about the provider split and other changes to TF v0.10.0 in the
official release announcement found [here][5].

[5]: https://www.hashicorp.com/blog/hashicorp-terraform-0-10/


## Full Provider Documentation

The provider is useful for Ordering services from Service catalog.

### Example
```hcl

# Configure the Cloudform Provider
provider "cloudforms" {
	ip = "${var.CF_SERVER_IP}"
	user_name = "${var.CF_USER_NAME}"
	password = "${var.CF_PASSWORD}"
}

# Data Source cloudforms_service
data  "cloudforms_service" "myservice"{
    name = "${var.SERVICE_NAME}"
}

# Data Source cloudforms_service_template
data "cloudforms_service_template" "mytemplate"{
	name = "${var.SERVICE_TEMPLATE_NAME}"
}


# Resource cloudforms_service_request
resource "cloudforms_service_request" "test" {	
	name = "${var.TEMPLATE_NAME}"
	template_href = "${data.cloudforms_service_template.mytemplate.href}"
	catalog_id ="${data.cloudforms_service_template.mytemplate.service_template_catalog_id}"
	input_file_name = "${var.INPUT_FILE_NAME}"
	time_out= 50
}	

output "Service_templates_href"{
	value = "${data.cloudforms_service_template.mytemplate.href}"
}


```

# Building The Provider

**NOTE:** Unless you are [developing][6] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][7]).

[6]: #developing-the-provider
[7]: #using-the-provider


## Cloning the Project

First, you will want to clone the repository to
`$GOPATH/src/github.com/terraform-providers/terraform-provider-cloudforms`:

```sh
mkdir -p $GOPATH/src/github.com/terraform-providers
cd $GOPATH/src/github.com/terraform-providers
git clone git@github.com:terraform-providers/terraform-provider-cloudforms
```

## Running the Build

After the clone has been completed, you can enter the provider directory and
build the provider.

```sh
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-cloudforms
make build
```

## Installing the Local Plugin

After the build is complete, copy the `terraform-provider-cloudforms` binary into
the same path as your `terraform` binary, and re-run `terraform init`.

After this, your project-local `.terraform/plugins/ARCH/lock.json` (where `ARCH`
matches the architecture of your machine) file should contain a SHA256 sum that
matches the local plugin. Run `shasum -a 256` on the binary to verify the values
match.

# Developing the Provider

If you wish to work on the provider, you'll first need [Go][8] installed on your
machine (version 1.13.1+ is **required**). You'll also need to correctly setup a
[GOPATH][9], as well as adding `$GOPATH/bin` to your `$PATH`.

[8]: https://golang.org/
[9]: http://golang.org/doc/code.html#GOPATH

See [Building the Provider][10] for details on building the provider.

[10]: #building-the-provider


## Checking the Logs
To persist logged output you can set TF_LOG_PATH in order to force the log to always be appended to a specific file when logging is enabled. Note that even when TF_LOG_PATH is set, TF_LOG must be set in order for any logging to be enabled.

To check logs use the following commands :
```sh
# Specify Log Level
export TF_LOG=DEBUG
# Specify Log File Path
export TF_LOG_PATH='. . .'
```

## Configuring Environment Variables

Most of the tests in this provider require a comprehensive list of environment
variables to run. See the individual `*_test.go` files in the
[`cloudforms/`](cloudforms/) directory for more details. The next section also
describes how you can manage a configuration file of the test environment
variables.

## Running the Acceptance Tests
In order to perform acceptance tests of cloudforms, first set in your environment variables required for the connection (`CF_SERVER_IP`,`CF_USER_NAME`,`CF_PASSWORD`,`CF_SERVICE_NAME`,`CF_INPUT_FILE_NAME`,`CF_TEMPLATE_NAME`).
After this is done, you can run the acceptance tests by running:

```sh
$ make testacc
```


# Building The Provider

**NOTE:** Unless you are [developing][7] or require a pre-release bugfix or feature,
you will want to use the officially released version of the provider (see [the
section above][8]).

[7]: #developing-the-provider
[8]: #using-the-provider

