package archive

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// CreateAndRead the zip archive.
func CreateAndRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	files := d.Get(FieldFiles).(map[string]interface{})

	path, err := writeZip(files)
	if err != nil {
		return diag.FromErr(err)
	}

	md5, err := getMD5(path)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(md5)

	d.Set(FieldOutputPath, path)
	d.Set(FieldOutputMD5, md5)

	return nil
}
