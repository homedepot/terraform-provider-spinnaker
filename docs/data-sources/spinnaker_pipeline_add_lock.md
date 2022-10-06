---
page_title: "spinnaker_pipeline_add_lock"
---

# spinnaker_pipeline_add_lock Data Source

Add lock to spinnaker pipeline resource

## Example Usage

```
provider "spinnaker" {
    server = "http://spinnaker-gate.myorg.io"
}

data "spinnaker_pipeline_add_lock" "example_pipeline" {
    pipeline = file ("/path/to/pipeline/pipeline.json")
}
```

## Argument Reference

- `pipeline` - (Required) Pipeline json.
- `ui` - (Optional) Unknown behavior. (Default: `true`)
- `allow_unlock_ui` - (Optional) If set to true means pipelibe can be unlocked from inside the Spinnaker UI (deck). If set to false, then all changes to the pipeline must be done thru API. (Default: `true`)
- `description` - (Optional) Reason shown in Spinnaker UI (deck) for pipeline to be locked (Default: No default unless environment variables are set, then `"Maintained in $GITHUB_SERVER_URL/$GITHUB_REPOSITORY"` or `"Maintained in repo: $GITHUB_REPOSITORY"`).

## Attribute Reference

In addition to the above, the following attributes are exported:

- `rendered` - Pipeline JSON after lock is added.

