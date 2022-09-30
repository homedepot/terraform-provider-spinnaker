package spinnaker

import (
   "os"
   "io/ioutil"

   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

   "github.com/spinnaker/spin/cmd/output"
	gate "github.com/spinnaker/spin/cmd/gateclient"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL for Gate",
				DefaultFunc: schema.EnvDefaultFunc("GATE_URL", nil),
			},
			"config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path to Gate config file",
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_CONFIG_PATH", nil),
			},
         "upsert_strategy": {
            Type:        schema.TypeBool,
            Optional:    true,
            Description: "When creating pipelines, update pipeline if it already exists.",
            Default:     true,
         },
         "https_proxy": {
            Type:        schema.TypeString,
            Optional:    true,
            Description: "HTTPS proxy",
            Default: "",
         },
			"ignore_cert_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Ignore certificate errors from Gate",
				Default:     false,
			},
			"default_headers": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Headers to be passed to the gate endpoint by the client on each request",
				Default:     "",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"spinnaker_application":              resourceApplication(),
			"spinnaker_pipeline":                 resourcePipeline(),
			"spinnaker_pipeline_template":        resourcePipelineTemplate(),
			"spinnaker_pipeline_template_config": resourcePipelineTemplateConfig(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"spinnaker_pipeline": datasourcePipeline(),
		},
		ConfigureFunc: providerConfigureFunc,
	}
}

type gateConfig struct {
	server string
	client *gate.GatewayClient
   upSertStrategy bool
}

func providerConfigureFunc(data *schema.ResourceData) (interface{}, error) {
	server := data.Get("server").(string)
	config := data.Get("config").(string)
   upSertStrategy := data.Get("upsert_strategy").(bool)
   httpsProxy := data.Get("https_proxy").(string)
	ignoreCertErrors := data.Get("ignore_cert_errors").(bool)
	defaultHeaders := data.Get("default_headers").(string)

   if httpsProxy != "" {
      os.Setenv("HTTPS_PROXY", httpsProxy)
   }

   quiet := false
   noColor := true
   outputFormater, err := output.ParseOutputFormat("")
   if err != nil {
      return nil, err
   }
   ui := output.NewUI(quiet, noColor, outputFormater, ioutil.Discard, ioutil.Discard)

   client, err := gate.NewGateClient(ui, server, defaultHeaders, config, ignoreCertErrors)
	if err != nil {
		return nil, err
	}
	return gateConfig{
		server: data.Get("server").(string),
		client: client,
      upSertStrategy: upSertStrategy,
	}, nil
}
