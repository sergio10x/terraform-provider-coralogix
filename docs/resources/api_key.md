---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "coralogix_api_key Resource - terraform-provider-coralogix"
subcategory: ""
description: |-
  Coralogix Api keys.
---

# coralogix_api_key (Resource)

Coralogix Api keys.

## Example Usage

```terraform
terraform {
  required_providers {
    coralogix = {
      version = "~> 2.0"
      source  = "coralogix/coralogix"
    }
  }
}

provider "coralogix" {
  #api_key = "<add your api key here or add env variable CORALOGIX_API_KEY>"
  #env = "<add the environment you want to work at or add env variable CORALOGIX_ENV>"
}

resource "coralogix_api_key" "example" {
  name  = "My APM KEY"
  owner = {
    team_id : "4013254"
  }
  active = true
  presets = ["APM"]
  permissions = ["livetail:Read"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Api Key name.
- `owner` (Attributes) Api Key Owner. It can either be a team_id, organisation_id, or a user_id (see [below for nested schema](#nestedatt--owner))
- `permissions` (Set of String) Api Key Permissions
- `presets` (Set of String) Api Key Presets

### Optional

- `active` (Boolean) Api Key Is Active.

### Read-Only

- `hashed` (Boolean) Api Key Is Hashed.
- `id` (String) ApiKey ID.
- `value` (String, Sensitive) Api Key value.

<a id="nestedatt--owner"></a>
### Nested Schema for `owner`

Optional:

- `organisation_id` (String)
- `team_id` (String)
- `user_id` (String)
