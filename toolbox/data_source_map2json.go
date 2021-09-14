package toolbox

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMap2JSON() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRegexRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
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
			"json": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceMap2JSONRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	items := d.Get("items").(map[string]interface{})
	for k, v := range items {
		log.Printf("[INFO] item is %v => %v\n", k, v)
	}

	result, err := json.Marshal(m)
	if err != nil {
		log.Printf("[ERROR] error marshalling output: %v", err)
		return diag.FromErr(err)
	}

	if err := d.Set("json", string(result)); err != nil {
		log.Printf("[ERROR] error setting JSON value: %v", err)
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
