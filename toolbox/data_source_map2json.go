package toolbox

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMap2JSON() *schema.Resource {
	return &schema.Resource{
		Description: "Maps a dictionary into a JSON object.",
		Schema: map[string]*schema.Schema{
			"items": {
				Type:     schema.TypeMap,
				Required: true,
				// ValidateDiagFunc: func(value interface{}, key cty.Path) diag.Diagnostics {
				// 	var diags diag.Diagnostics

				// 	_, err := regexp.Compile(value.(string))
				// 	if err != nil {
				// 		return diag.FromErr(err)
				// 	}

				// 	return diags
				// },
			},
			"json": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		ReadContext: dataSourceMap2JSONRead,
	}
}

func dataSourceMap2JSONRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	config := m.(*configuration)

	items := d.Get("items").(map[string]interface{})
	for k, v := range items {
		config.Logger.Debug("new item found", "key", k, "value", v)
	}

	result, err := json.Marshal(items)
	if err != nil {
		config.Logger.Error("error marshalling output", "error", err)
		return diag.FromErr(err)
	}

	config.Logger.Debug("input marshalled to JSON", "value", string(result))

	if err := d.Set("json", string(result)); err != nil {
		config.Logger.Error("error setting JSON value in state", "error", err)
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
