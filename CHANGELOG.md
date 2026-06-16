## 0.1.13 (June 16, 2026)

## 0.1.12 (June 16, 2026)

## 0.1.11 (June 16, 2026)

## 0.0.7 (June 09, 2026)

## 0.0.6 (June 09, 2026)

BUG FIXES:

* Send the configured `api_key` as a Bearer token. The client emitted a raw `Authorization: <key>` header, which the API rejected with 401 "missing bearer token".
* Omit read-only, server-assigned fields (`id`, `created_at`, `updated_at`, ...) from create and update request bodies. They were sent as empty values and the API rejected them with 422 "unexpected property".
## 0.0.5 (June 09, 2026)

## 0.0.4 (June 09, 2026)

## 0.0.3 (June 09, 2026)

BUG FIXES:

* Remove the `$schema` metadata field from resource and data source models. It was generated with an invalid `tfsdk:"__schema"` tag, which caused `terraform apply` to fail with "invalid tfsdk tag, must ... start with a letter" when creating a resource.
## 0.0.2 (June 09, 2026)

## 0.0.1 (June 09, 2026)

FEATURES:

* Initial Provider ([#1](https://github.com/hashicorp/terraform-plugin-framework/issues/1))

