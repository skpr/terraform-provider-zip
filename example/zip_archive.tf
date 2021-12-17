resource "zip_archive" "test" {
  files = {
    "foo" = "bar",
  }
}