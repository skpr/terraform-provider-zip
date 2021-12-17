FROM hashicorp/terraform:1.1.1

RUN mkdir -p /root/.terraform.d/plugins

ADD terraform-provider-zip /root/.terraform.d/plugins/registry.terraform.io/skpr/zip/99.0.0/linux_amd64/terraform-provider-zip_v99.0.0
RUN chmod +x /root/.terraform.d/plugins/registry.terraform.io/*/*/*/linux_amd64/terraform-provider-*