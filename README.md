Terraform Provider: Zip
-----------------------

A Terraform provider for creating a dynamic ZIP archive.

## Example

```hcl
resource "zip_archive" "example" {
  files = {
    "example.yml" = "foo: bar",
  }
}
```
