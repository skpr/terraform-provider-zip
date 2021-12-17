output "ouput_file" {
  value = zip_archive.test.output_path
}

output "ouput_md5" {
  value = zip_archive.test.output_md5
}