package discourse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"context"
	"terraform-provider-discourse/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BASE_URL", ""),
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_KEY", ""),
			},
			"api_username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("API_USERNAME", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"discourse_user": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"discourse_user": dataSourceUser(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData)  (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	base_url := d.Get("base_url").(string)
	api_key := d.Get("api_key").(string)
	api_username := d.Get("api_username").(string)
	return client.NewClient(base_url, api_key, api_username), diags
}