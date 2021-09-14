package toolbox

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"inet.af/netaddr"
)

func dataSourceCIDR() *schema.Resource {
	return &schema.Resource{
		Description: "Perform prefix set operations (union, symmetric difference) on CIDRs.",
		Schema: map[string]*schema.Schema{
			"added": {
				Type:        schema.TypeSet,
				Description: "A list of prefixes (in CIDR format) to be added from the base IP Set.",
				Required:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateCIDR,
				},
				MinItems: 1,
			},
			"subtracted": {
				Type:        schema.TypeSet,
				Description: "A list of prefixes (in CIDR format) to be subtracted from the base IP Set.",
				Optional:    true,
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validateCIDR, // maybe validation.IsCIDR???
				},
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
	config := m.(*configuration)

	var base netaddr.IPSetBuilder

	prefixes := d.Get("added").(*schema.Set)
	for _, value := range prefixes.List() {
		config.Logger.Debug("adding prefix to base set", "prefix", value.(string))
		base.AddPrefix(netaddr.MustParseIPPrefix(value.(string)))
	}

	prefixes = d.Get("subtracted").(*schema.Set)
	for _, value := range prefixes.List() {
		config.Logger.Debug("subtracting prefix from base set", "prefix", value.(string))
		base.RemovePrefix(netaddr.MustParseIPPrefix(value.(string)))
	}

	set, err := base.IPSet()
	if err != nil {
		config.Logger.Error("error parsing CIDRs", "error", err)
		return diag.FromErr(err)
	}

	values := []interface{}{}
	for _, prefix := range set.Prefixes() {
		config.Logger.Debug("adding prefix to output set", "prefix", prefix)
		values = append(values, prefix.String())
	}

	if err := d.Set("prefixes", values); err != nil {
		config.Logger.Error("error setting prefixes into state", "error", err)
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func validateCIDR(value interface{}, key cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	_, err := netaddr.ParseIPPrefix(value.(string))
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
