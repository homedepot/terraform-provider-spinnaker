---
page_title: "spinnaker_application"
---

# spinnaker_application Resource

Manage spinnaker applications

## Example Usage

```hcl
resource "spinnaker_application" "terraformtest" {
    application = "terraformtest"
    email       = "user@example.com"
}
```

```hcl
resource "spinnaker_application" "terraformtest" {
    application = "terraformtest"
    email       = "user@example.com"
    cloud_providers = [
       "kubernetes",
       "appengine"
    ]
    permissions {
       read = [ "group1", "group2", "group3" ]
       write = [ "group1" ]
       execute = [ "group1", "group2" ]
    }
}
```

## Argument Reference

- `application` - (Required) Spinnaker application name.
- `email` - (Required) Application owner email.
- `description` - (Optional) Description. (Default: `""`)
- `platform_health_only` - (Optional) Consider only cloud provider health when executing tasks. (Default: `false`)
- `platform_health_only_show_override` - (Optional) Show health override option for each operation. (Default: `false`)
- `cloud_providers` - (Optional) Array list of cloud providers.
- `permissions` - (Optional) Array of permissions in key/value pairs. Valid permission keys are `read`, `write` and `execute`. Permission value is array list of groups.

## Attribute Reference
