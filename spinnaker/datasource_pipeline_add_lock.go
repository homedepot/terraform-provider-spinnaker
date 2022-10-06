package spinnaker

import (
   "crypto/sha256"
   "encoding/json"
   "encoding/hex"
   "os"

   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourcePipelineAddLock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pipeline": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validatePipelineJson,
            Description:  "Pipeline json",
			},
			"ui": {
				Type:         schema.TypeBool,
				Optional:     true,
            Default:      true,
            Description:  "Unknown behavior.",
			},
			"allow_unlock_ui": {
				Type:         schema.TypeBool,
				Optional:     true,
            Default:      true,
            Description:  "If set to true means pipelibe can be unlocked from inside the Spinnaker UI (deck). If set to false, then all changes to the pipeline must be done thru API.",
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
            DefaultFunc:  getLockedDescription,
            Description:  "Reason shown in Spinnaker UI (deck) for pipeline to be locked.",
			},
			"rendered": {
				Type:         schema.TypeString,
				Computed:     true,
            Description:  "Pipeline JSON after lock is added.",
			},
		},
		Read: datasourcePipelineAddLockRead,
	}
}

func datasourcePipelineAddLockRead(data *schema.ResourceData, meta interface{}) error {
   pipelineBlob := []byte( data.Get("pipeline").(string) )
   ui := data.Get("ui").(bool)
   allowUnlockUi := data.Get("allow_unlock_ui").(bool)
   description := data.Get("description").(string)

   var pipeline_map map[string]interface{}

   err := json.Unmarshal(pipelineBlob, &pipeline_map)
   if err != nil {
      return err
   }

   // remove "locked" if exists
   delete (pipeline_map, "locked")

   // add "locked"
   if len(description) > 0 {
      pipeline_map["locked"] = map[string]interface{}{ "ui":ui, "allowUnlockUi": allowUnlockUi, "description": description }
   } else {
      pipeline_map["locked"] = map[string]interface{}{ "ui":ui, "allowUnlockUi": allowUnlockUi }
   }

   b, err := json.Marshal(pipeline_map)
   if err != nil {
      return err
   }
   err = data.Set("rendered", string(b))
   if err != nil {
      return err
   }
   sha := sha256.Sum256(b)
   data.SetId(hex.EncodeToString(sha[:]))
   
   return nil
}

func getLockedDescription() (v interface{}, err error) {
   repo := os.Getenv("GITHUB_REPOSITORY")
   baseUrl := os.Getenv("GITHUB_SERVER_URL")
   if len(repo) > 0 {
      if len(baseUrl) > 0 {
         v = "Maintained in "+baseUrl+"/"+repo
      } else {
         v = "Maintained in repo: "+repo
      }
   }
   return
}
