package provider

import (
	"context"
	"crypto/sha1"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dotenv "github.com/direnv/go-dotenv"
)

func dataSourceDotEnv() *schema.Resource {
	return &schema.Resource{
		Description: "A data source for parsing .env files",
		ReadContext: dataSourceDotEnvRead,
		Schema: map[string]*schema.Schema{
			"filename": {
				Description:  "Path to a dotenv file.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"filename", "string"},
			},
			"string": {
				Description:  "A string containing .env notation.",
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"filename", "string"},
			},
			"env": {
				Description: "A map of envronment variables read from the file.",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDotEnvRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var contents string
	if filename, ok := d.Get("filename").(string); ok && filename != "" {
		b, err := os.ReadFile(filename)
		if err != nil {
			return diag.Errorf("Could not open file %q: %v", filename, err)
		}
		contents = string(b)
	} else if str, ok := d.Get("string").(string); ok && str != "" {
		contents = str
	}

	m, err := dotenv.Parse(string(contents))
	if err != nil {
		return diag.Errorf("Could not parse %q: %v", contents, err)
	}
	d.Set("env", m)

	// generate ID as sha of the contents
	h := sha1.New()
	h.Write([]byte(contents))
	d.SetId(string(h.Sum(nil)))
	return nil
}
