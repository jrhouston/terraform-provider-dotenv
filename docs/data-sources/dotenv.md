---
page_title: "dotenv Data Source - terraform-provider-dotenv"
subcategory: ""
description: |-
  A data source for reading .env files.
---

# Data Source `dotenv`

A data source for reading files containing `.env` notation.

## Example Usage

```terraform
data "dotenv" "dev" {
  filename = "dev.env"
}
```

## Argument Reference

### Optional

- **filename** (String, Optional) A path to a file containing .env notation. One of of `filename` or `string` must be set.
- **string** (String, Optional) A string containing .env notation. One of of `filename` or `string` must be set.

## Attributes Reference


- **env** A map containing the parsed .env file.




