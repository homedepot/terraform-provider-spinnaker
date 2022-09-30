package main

import (
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
   "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

   "github.com/guido9j/terraform-provider-spinnaker/spinnaker"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return spinnaker.Provider()
		},
	})
}
