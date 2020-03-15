Terraform Provider for Databricks
==================

[![Build Status](https://travis-ci.com/innovationnorway/terraform-provider-databricks.svg?branch=master)](https://travis-ci.com/innovationnorway/terraform-provider-databricks)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x
-	[Go](https://golang.org/doc/install) 1.13

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/innovationnorway/terraform-provider-databricks`

```sh
$ mkdir -p $GOPATH/src/github.com/innovationnorway; cd $GOPATH/src/github.com/innovationnorway
$ git clone https://github.com/innovationnorway/terraform-provider-databricks.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/innovationnorway/terraform-provider-databricks
$ make build
```

Using the provider
----------------------

```hcl
provider "databricks" {
  host  = var.databricks_host
  token = var.databricks_token
}

data "databricks_cluster" "example" {
  cluster_id = "0308-153622-deity853"
}
```

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-databricks
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```