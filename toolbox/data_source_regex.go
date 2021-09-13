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
			// 	Elem:     &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// },
			"matches": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
					// Computed: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
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

	groups := r.FindAllStringSubmatch(input, -1)
	log.Printf("[INFO] result is %v (%T)\n", groups, groups)

	matched := groups != nil

	if err := d.Set("matched", matched); err != nil {
		return diag.FromErr(err)
	}

	matches := []interface{}{}

	for _, group := range groups {
		log.Printf("[INFO] adding group to output set: %v\n", group)
		submatches := []interface{}{}
		for _, match := range group {
			submatches = append(submatches, match)
		}
		matches = append(matches, submatches)
	}

	if err := d.Set("matches", matches); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
