package toolbox

import (
	"context"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRegex() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRegexRead,
		Schema: map[string]*schema.Schema{
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: func(value interface{}, key cty.Path) diag.Diagnostics {
					var diags diag.Diagnostics

					_, err := regexp.Compile(value.(string))
					if err != nil {
						return diag.FromErr(err)
					}

					return diags
				},
			},
			"input": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"matched": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			// "matches": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"ingredient_id": &schema.Schema{
			// 				Type:     schema.TypeInt,
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func dataSourceRegexRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	pattern := d.Get("pattern").(string)
	log.Printf("[INFO] pattern is %q\n", pattern)

	input := d.Get("input").(string)
	log.Printf("[INFO] input is %q\n", input)

	r, err := regexp.Compile(pattern)
	if err != nil {
		return diag.FromErr(err)
	}

	result := r.MatchString("peach")

	if err := d.Set("matched", result); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
