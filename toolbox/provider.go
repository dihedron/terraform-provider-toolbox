package toolbox

import (
	"context"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"loglevel": {
				Description:  "The logging level of the toolbox provider  ",
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "info",
				ValidateFunc: validation.StringInSlice([]string{"trace", "debug", "info", "warn", "error", "off"}, true),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"toolbox_regex":    dataSourceRegex(),
			"toolbox_cidr":     dataSourceCIDR(),
			"toolbox_map2json": dataSourceMap2JSON(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

type configuration struct {
	Logger hclog.Logger
	// TODO: add furter fields here if needed
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	logger := hclog.New(&hclog.LoggerOptions{
		// avoid adding the provider's name: it makes messages too verbose
		// Name:  "terraform-provider-toolbox",
		Level: hclog.LevelFromString(strings.ToLower(data.Get("loglevel").(string))),
	})
	configuration := &configuration{
		Logger: logger,
	}
	return configuration, nil
}
