package discourse

import (
	"context"	
	"log"
	"terraform-provider-discourse/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"admin" : &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"active" : &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	email := d.Get("email").(string)
	user, err := apiClient.GetUser(email)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	d.SetId(user.Email)
	d.Set("user_id", user.Id)
	d.Set("username", user.Username)
	d.Set("name", user.Name)
	d.Set("admin", user.Admin)
	d.Set("active", user.Active)
	d.Set("email", user.Email)		
	return diags
}