package toolbox

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCURL() *schema.Resource {
	return &schema.Resource{
		Description: "Perform an HTTP(s) request to a remote endpoint.",
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Description:  "The URL to use for the request.",
				Required:     true,
				ValidateFunc: validation.IsURLWithScheme([]string{"http", "https"}),
			},
			"method": {
				Type:         schema.TypeString,
				Description:  "The HTTP verb to use for the request.",
				Optional:     true,
				Default:      "GET",
				ValidateFunc: validation.StringInSlice([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}, true),
			},
			"header": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Description:  "The name of the header on the wire (e.g. 'Proxy-Authenticate').",
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"values": {
							Type:        schema.TypeList,
							Description: "The values associated with the current header.",
							Required:    true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringIsNotEmpty,
							},
							MinItems: 1,
						},
					},
				},
			},
			"body": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ReadContext: dataSourceCURLRead,
	}
}

func dataSourceCURLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	config := m.(*configuration)

	url := d.Get("url").(string)
	config.Logger.Debug("sending request to URL", "url", url)

	set := d.Get("header").(*schema.Set)
	config.Logger.Debug("type of list", "type", fmt.Sprintf("%T", set.List()))
	for _, value := range set.List() {
		config.Logger.Debug("type of list element", "type", fmt.Sprintf("%T", value))
		for k, v := range value.(map[string]interface{}) {
			config.Logger.Debug("type of list element value", "key", k, "value", v, "type", fmt.Sprintf("%T", v))
		}
	}

	// input := d.Get("headers").(string)
	// config.Logger.Debug("using input value", "input", input)

	// r, err := regexp.Compile(pattern)
	// if err != nil {
	// 	config.Logger.Error("error compiling regular expression", "error", err)
	// 	return diag.FromErr(err)
	// }

	// groups := r.FindAllStringSubmatch(input, -1)
	// config.Logger.Debug("regular expression applied", "result", groups)

	// matched := groups != nil

	// if err := d.Set("matched", matched); err != nil {
	// 	config.Logger.Error("error setting 'matched' value in state", "error", err)
	// 	return diag.FromErr(err)
	// }

	// matches := []interface{}{}

	// for _, group := range groups {
	// 	config.Logger.Debug("adding group to output set", "group", group)
	// 	submatches := []interface{}{}
	// 	for _, match := range group {
	// 		submatches = append(submatches, match)
	// 	}
	// 	matches = append(matches, submatches)
	// }

	// if err := d.Set("matches", matches); err != nil {
	// 	config.Logger.Error("error setting 'matches' value in state", "error", err)
	// 	return diag.FromErr(err)
	// }

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
