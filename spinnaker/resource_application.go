package spinnaker

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-spinnaker/spinnaker/api"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"application": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateApplicationName,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"platform_health_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"platform_health_only_show_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"cloud_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"permissions": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read": {
							Type:	schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"execute": {
							Type:	schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"write": {
							Type:	schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
		Create: resourceApplicationCreateOrUpdate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationCreateOrUpdate,
		Delete: resourceApplicationDelete,
		Exists: resourceApplicationExists,
	}
}

type applicationRead struct {
	Name       string `json:"name"`
	Attributes struct {
		Email string `json:"email"`
	} `json:"attributes"`
}

func resourceApplicationCreateOrUpdate(data *schema.ResourceData, meta interface{}) error {
	clientConfig := meta.(gateConfig)
	client := clientConfig.client
	application := data.Get("application").(string)
	email := data.Get("email").(string)
	description := data.Get("description").(string)
	platform_health_only := data.Get("platform_health_only").(bool)
	platform_health_only_show_override := data.Get("platform_health_only_show_override").(bool)
	cloud_providers := data.Get("cloud_providers").([]interface{})
	permissions := data.Get("permissions").(*schema.Set)

	if err := api.CreateOrUpdateApplication(client, application, email, description, platform_health_only, platform_health_only_show_override, cloud_providers, permissions); err != nil {
		return err
	}

	return resourceApplicationRead(data, meta)
}

func resourceApplicationRead(data *schema.ResourceData, meta interface{}) error {
	clientConfig := meta.(gateConfig)
	client := clientConfig.client
	applicationName := data.Get("application").(string)
	var app applicationRead
	if err := api.GetApplication(client, applicationName, &app); err != nil {
		return err
	}

	return readApplication(data, app)
}

func resourceApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	clientConfig := meta.(gateConfig)
	client := clientConfig.client
	applicationName := data.Get("application").(string)

	return api.DeleteAppliation(client, applicationName)
}

func resourceApplicationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	clientConfig := meta.(gateConfig)
	client := clientConfig.client
	applicationName := data.Get("application").(string)

	var app applicationRead
	if err := api.GetApplication(client, applicationName, &app); err != nil {
		errmsg := err.Error()
		if strings.Contains(errmsg, "not found") {
			return false, nil
		}
		return false, err
	}

	if app.Name == "" {
		return false, nil
	}

	return true, nil
}

func readApplication(data *schema.ResourceData, application applicationRead) error {
	data.SetId(application.Name)
	return nil
}
