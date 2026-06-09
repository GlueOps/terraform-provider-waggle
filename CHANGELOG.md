## 0.0.4 (June 09, 2026)

## 0.0.3 (June 09, 2026)

BUG FIXES:

* Remove the `$schema` metadata field from resource and data source models. It was generated with an invalid `tfsdk:"__schema"` tag, which caused `terraform apply` to fail with "invalid tfsdk tag, must ... start with a letter" when creating a resource.
## 0.0.2 (June 09, 2026)

## 0.0.1 (June 09, 2026)

FEATURES:

* Initial Provider ([#1](https://github.com/hashicorp/terraform-plugin-framework/issues/1))

