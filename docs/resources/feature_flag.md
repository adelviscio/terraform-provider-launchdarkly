---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "launchdarkly_feature_flag Resource - launchdarkly"
subcategory: ""
description: |-
  Provides a LaunchDarkly feature flag resource.
  This resource allows you to create and manage feature flags within your LaunchDarkly organization.
---

# launchdarkly_feature_flag (Resource)

Provides a LaunchDarkly feature flag resource.

This resource allows you to create and manage feature flags within your LaunchDarkly organization.

## Example Usage

```terraform
resource "launchdarkly_feature_flag" "building_materials" {
  project_key = launchdarkly_project.example.key
  key         = "building-materials"
  name        = "Building materials"
  description = "this is a multivariate flag with string variations."

  variation_type = "string"
  variations {
    value       = "straw"
    name        = "Straw"
    description = "Watch out for wind."
  }
  variations {
    value       = "sticks"
    name        = "Sticks"
    description = "Sturdier than straw"
  }
  variations {
    value       = "bricks"
    name        = "Bricks"
    description = "The strongest variation"
  }

  client_side_availability {
    using_environment_id = false
    using_mobile_key     = true
  }

  defaults {
    on_variation  = 2
    off_variation = 0
  }

  tags = [
    "example",
    "terraform",
    "multivariate",
    "building-materials",
  ]
}

resource "launchdarkly_feature_flag" "json_example" {
  project_key = "example-project"
  key         = "json-example"
  name        = "JSON example flag"

  variation_type = "json"
  variations {
    name  = "Single foo"
    value = jsonencode({ "foo" : "bar" })
  }
  variations {
    name  = "Multiple foos"
    value = jsonencode({ "foos" : ["bar1", "bar2"] })
  }

  defaults {
    on_variation  = 1
    off_variation = 0
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `key` (String) The unique feature flag key that references the flag in your application code. A change in this field will force the destruction of the existing resource and the creation of a new one.
- `name` (String) The human-readable name of the feature flag.
- `project_key` (String) The feature flag's project key. A change in this field will force the destruction of the existing resource and the creation of a new one.
- `variation_type` (String) The feature flag's variation type: `boolean`, `string`, `number` or `json`.

### Optional

- `archived` (Boolean) Specifies whether the flag is archived or not. Note that you cannot create a new flag that is archived, but can update a flag to be archived.
- `client_side_availability` (Block List) (see [below for nested schema](#nestedblock--client_side_availability))
- `custom_properties` (Block Set, Max: 64) List of nested blocks describing the feature flag's [custom properties](https://docs.launchdarkly.com/home/connecting/custom-properties) (see [below for nested schema](#nestedblock--custom_properties))
- `defaults` (Block List, Max: 1) A block containing the indices of the variations to be used as the default on and off variations in all new environments. Flag configurations in existing environments will not be changed nor updated if the configuration block is removed. (see [below for nested schema](#nestedblock--defaults))
- `description` (String) The feature flag's description.
- `include_in_snippet` (Boolean, Deprecated) Specifies whether this flag should be made available to the client-side JavaScript SDK using the client-side Id. This value gets its default from your project configuration if not set. `include_in_snippet` is now deprecated. Please migrate to `client_side_availability.using_environment_id` to maintain future compatibility.
- `maintainer_id` (String) The feature flag maintainer's 24 character alphanumeric team member ID. If not set, it will automatically be or stay set to the member ID associated with the API key used by your LaunchDarkly Terraform provider or the most recently-set maintainer.
- `tags` (Set of String) Tags associated with your resource.
- `temporary` (Boolean) Specifies whether the flag is a temporary flag.
- `variations` (Block List) An array of possible variations for the flag (see [below for nested schema](#nestedblock--variations))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--client_side_availability"></a>
### Nested Schema for `client_side_availability`

Optional:

- `using_environment_id` (Boolean) Whether this flag is available to SDKs using the client-side ID.
- `using_mobile_key` (Boolean) Whether this flag is available to SDKs using a mobile key.


<a id="nestedblock--custom_properties"></a>
### Nested Schema for `custom_properties`

Required:

- `key` (String) The unique custom property key.
- `name` (String) The name of the custom property.
- `value` (List of String) The list of custom property value strings.


<a id="nestedblock--defaults"></a>
### Nested Schema for `defaults`

Required:

- `off_variation` (Number) The index of the variation the flag will default to in all new environments when off.
- `on_variation` (Number) The index of the variation the flag will default to in all new environments when on.


<a id="nestedblock--variations"></a>
### Nested Schema for `variations`

Required:

- `value` (String) The variation value. The value's type must correspond to the `variation_type` argument. For example: `variation_type = "boolean"` accepts only `true` or `false`. The `number` variation type accepts both floats and ints, but please note that any trailing zeroes on floats will be trimmed (i.e. `1.1` and `1.100` will both be converted to `1.1`).

If you wish to define an empty string variation, you must still define the value field on the variations block like so:

```terraform
variations {
  value = ""
}
```

Optional:

- `description` (String) The variation's description.
- `name` (String) The name of the variation.

## Import

Import is supported using the following syntax:

```shell
# Import a feature flag using the feature flag's ID in the format `project_key/flag_key`.
terraform import launchdarkly_feature_flag.building_materials example-project/building-materials
```