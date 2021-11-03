---
page_title: "dotenv Provider"
subcategory: ""
description: |-
  
---

# dotenv Provider

This is a convenience provider for reading files containing `.env` notation. 

## Example Usage

Below is an example that illustrates how to use the data source to populate a `kubernetes_config_map`.

```terraform
provider kubernetes {
  config_path = "~/.kube/config"
}

data dotenv dev_config {
  # NOTE there must be a file called `dev.env` in the same directory as the .tf file
  filename = "dev.env"
}

resource kubernetes_config_map cm {
  metadata {
    name = "example"
  }

  data = data.dotenv.dev_config.env
}

```