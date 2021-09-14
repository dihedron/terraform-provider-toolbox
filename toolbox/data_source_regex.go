package toolbox

import (
	"context"
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
			"matches": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
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
	config := m.(*configuration)

	pattern := d.Get("pattern").(string)
	config.Logger.Debug("applying pattern", "pattern", pattern)

	input := d.Get("input").(string)
	config.Logger.Debug("using input value", "input", input)

	r, err := regexp.Compile(pattern)
	if err != nil {
		config.Logger.Error("error compiling regular expression", "error", err)
		return diag.FromErr(err)
	}

	groups := r.FindAllStringSubmatch(input, -1)
	config.Logger.Debug("regular expression applied", "result", groups)

	matched := groups != nil

	if err := d.Set("matched", matched); err != nil {
		config.Logger.Error("error setting 'matched' value in state", "error", err)
		return diag.FromErr(err)
	}

	matches := []interface{}{}

	for _, group := range groups {
		config.Logger.Debug("adding group to output set", "group", group)
		submatches := []interface{}{}
		for _, match := range group {
			submatches = append(submatches, match)
		}
		matches = append(matches, submatches)
	}

	if err := d.Set("matches", matches); err != nil {
		config.Logger.Error("error setting 'matches' value in state", "error", err)
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
