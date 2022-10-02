package analyticshub

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/analyticshub/v1"
)

const (
	defaultAPIVersion = "4.0"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"analyticshub_exchange":     resourceExchange(),
			"analyticshub_listing":      resourceListing(),
			"analyticshub_subscription": resourceSubscription(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	analyticshubService, err := analyticshub.NewService(ctx)
	if err != nil {
		println(err)
	}
	return analyticshubService, diag.Diagnostics{}
}
