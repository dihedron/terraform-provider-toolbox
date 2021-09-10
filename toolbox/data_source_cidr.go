package toolbox

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"inet.af/netaddr"
)

func dataSourceCIDR() *schema.Resource {
	return &schema.Resource{
		Description: "Perform operations on CIDRs.",

		Schema: map[string]*schema.Schema{
			"added": {
				Type:        schema.TypeSet,
				Description: "A list of prefixes (in CIDR format) to be added from the base IP Set.",
				Required:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateCIDR,
				},
				// Default:  []string{"10.0.0.0/8", "192.168.0.1/8", "172.0.0.2/8"}, // default is RFC1918
				MinItems: 1,
			},
			"subtracted": {
				Type:        schema.TypeSet,
				Description: "A list of prefixes (in CIDR format) to be subtracted from the base IP Set.",
				Required:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateCIDR,
				},
				MinItems: 1,
			},
			"prefixes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
		},
		ReadContext: dataSourceCIDRRead,
	}
}

func dataSourceCIDRRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	var base netaddr.IPSetBuilder

	prefixes := d.Get("added").(*schema.Set)
	for _, value := range prefixes.List() {
		log.Printf("[INFO] adding prefix to base: %q\n", value.(string))
		base.AddPrefix(netaddr.MustParseIPPrefix(value.(string)))
	}

	prefixes = d.Get("subtracted").(*schema.Set)
	for _, value := range prefixes.List() {
		log.Printf("[INFO] subtracting prefix from base: %q\n", value.(string))
		base.RemovePrefix(netaddr.MustParseIPPrefix(value.(string)))
	}

	//values := []string{"10.0.0.0/8", "192.168.0.1/8", "172.0.0.2/8"}
	set, err := base.IPSet()
	if err != nil {
		log.Printf("[ERROR] error parsing CIDRs: %v\n", err)
		return diag.FromErr(err)
	}

	values := []interface{}{}
	for _, prefix := range set.Prefixes() {
		log.Printf("[INFO] adding prefix to output set: %q\n", prefix)
		values = append(values, prefix.String())
	}

	if err := d.Set("prefixes", values); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

// 	return []string{"10.0.0.0/8", "192.168.0.1/8", "172.0.0.2/8"}
// }

func validateCIDR(value interface{}, key cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	_, err := netaddr.ParseIPPrefix(value.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
