package discourse

import (
	"context"	
	"terraform-provider-discourse/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"fmt"
	"time"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceUserRead,
		CreateContext: resourceUserCreate,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"admin" : &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"email" : &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"active" : &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default: true,
			},
		},
	}
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	time.Sleep(60*time.Second)
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	email := d.Id()
	user, err := apiClient.GetUser(email)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	d.Set("user_id", user.Id)
	d.Set("username", user.Username)
	d.Set("name", user.Name)
	d.Set("admin", user.Admin)
	d.Set("active", user.Active)
	d.Set("email", user.Email)
	return diags
}

func resourceUserCreate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	apiClient := m.(*client.Client)
	email := d.Get("email").(string)
	err := apiClient.NewUser(email)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	d.SetId(email)
	return diags
}

func resourceUserUpdate(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var _ diag.Diagnostics
	apiClient := m.(*client.Client)
	var diags diag.Diagnostics
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "User not allowed to change email",
			Detail:   "User not allowed to change email",
		})
		return diags
	}
	user := client.User{
		Email:   d.Get("email").(string),
		Username:  d.Get("username").(string),
		Name:      d.Get("name").(string),
	}
	if d.HasChange("active") && d.Get("active").(bool) == true {
		err  := apiClient.ActivateUser(d.Get("user_id").(int))
		if err != nil{
			log.Println("[ERROR]: ",err)
			return diag.FromErr(err)
		}
	}  else if d.HasChange("active") && d.Get("active").(bool) == false {
		err := apiClient.DeactivateUser(d.Get("user_id").(int))
		if err != nil{
			log.Println("[ERROR]: ",err)
			return diag.FromErr(err)
		}
	}
	err := apiClient.UpdateUser(&user)
	if err != nil {
		log.Println("[ERROR]: ",err)
		return diag.FromErr(err)
	}
	return diags
}

func resourceUserDelete(ctx context.Context,d *schema.ResourceData, m interface{}) diag.Diagnostics {
 	var diags diag.Diagnostics
 	apiClient := m.(*client.Client)
 	username := d.Get("username").(string)
 	err := apiClient.DeleteUser(username)
 	if err != nil {
		log.Println("[ERROR]: ",err)
 		return diag.FromErr(err)
 	}
 	d.SetId("")
 	return diags
}
