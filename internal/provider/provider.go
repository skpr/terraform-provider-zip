package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/skpr/terraform-provider-zip/internal/provider/archive"
)

const (
	// ResourceArchive which is built on demand.
	ResourceArchive = "zip_archive"
)

// Schema for this Terraform provider.
func Schema() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			ResourceArchive: archive.Resource(),
		},
	}
}
