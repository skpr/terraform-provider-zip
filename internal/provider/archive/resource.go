package archive

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

const (
	// FieldFiles which are added to the archive.
	FieldFiles = "files"
	// FieldOutputPath provides a path to the file which was generated.
	FieldOutputPath = "output_path"
	// FieldOutputMD5 provides an MD5 hash of the file which was generated.
	FieldOutputMD5 = "output_md5"
)

// Resource returns this packages resource.
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			FieldFiles: {
				Type:        schema.TypeMap,
				Description: "Files which are added to the archive",
				Required:    true,
				ForceNew:    true,
			},
			FieldOutputPath: {
				Type:        schema.TypeString,
				Description: "Path to the file which was generated",
				Computed:    true,
			},
			FieldOutputMD5: {
				Type:        schema.TypeString,
				Description: "MD5 hash of the file which was generated",
				Computed:    true,
			},
		},
		CreateContext: CreateAndRead,
		ReadContext:   CreateAndRead,
		DeleteContext: Delete,
	}
}
