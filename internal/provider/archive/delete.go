package archive

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Delete the zip file.
// This is a noop given the file we generate is temporary.
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	path := d.Get(FieldOutputPath).(string)

	if err := os.Remove(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return diag.FromErr(err)
	}

	return nil
}
