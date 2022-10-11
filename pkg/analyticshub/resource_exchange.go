package analyticshub

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"google.golang.org/api/analyticshub/v1beta1"
)

func resourceExchange() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceExchangeCreate,
		ReadContext:   resourceExchangeRead,
		UpdateContext: resourceExchangeUpdate,
		DeleteContext: resourceExchangeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceExchangeImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				/*ForceNew:    true,*/
				Description: "This is a return only property. Any values placed here will not be used by the resource",
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},*/
				ValidateFunc: validation.StringDoesNotContainAny(" "),
			},
			"project": {
				Type:     schema.TypeString,
				Required: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},*/
				ValidateFunc: validation.StringDoesNotContainAny(" "),
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},*/
				ValidateFunc: validation.StringDoesNotContainAny(" "),
			},
			"data_exchange_id": {
				Type:     schema.TypeString,
				Required: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},*/
				ValidateFunc: validation.StringDoesNotContainAny(" "),
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "),
				*/
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "), // Max 2000 bytes
				*/
			},
			"primary_contact": {
				Type:     schema.TypeString,
				Optional: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "), //Max len 2000 bytes
				*/
			},
			"documentation": {
				Type:     schema.TypeString,
				Optional: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "),

				*/
			},
			"listing_count": {
				Type:     schema.TypeInt,
				Optional: true,
				/*ForceNew:    true,*/
				Description: "This is a return only property. Any values here will not be used by the resource",
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "),*/
			},
			"icon": {
				Type:     schema.TypeString,
				Optional: true,
				/*ForceNew: true,*/
				/*DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return strings.EqualFold(old, new) // case-insensive comparing
				},
				ValidateFunc: validation.StringDoesNotContainAny(" "),
				*/
			},
		},
	}
}

func resourceExchangeCreate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	// analyticshubService, err := analyticshub.NewService(ctx)
	svc := i.(*analyticshub.Service)
	client := analyticshub.NewProjectsLocationsDataExchangesService(svc)

	exchange, err := expandExchange(data)

	if err != nil {
		return diag.FromErr(err)
	}

	// Need Parent
	parent, id := getIds(data)

	createSvc := client.Create(parent, exchange)
	createSvc.DataExchangeId(id)
	_, err = createSvc.Do()

	if err != nil {
		return diag.FromErr(err)
	}

	exId := fmt.Sprintf("%s/dataExchanges/%s", parent, id)
	data.SetId(exId)
	return resourceExchangeRead(ctx, data, i)
}

func resourceExchangeUpdate(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	svc := i.(*analyticshub.Service)
	client := analyticshub.NewProjectsLocationsDataExchangesService(svc)

	exchange, err := expandExchange(data)

	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.Patch(exchange.DisplayName, exchange).Do()

	return resourceExchangeRead(ctx, data, i)
}

func resourceExchangeDelete(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	svc := i.(*analyticshub.Service)
	client := analyticshub.NewProjectsLocationsDataExchangesService(svc)

	_, err := client.Delete(data.Id()).Do()

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceExchangeImport(ctx context.Context, data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	if err := resourceExchangeRead(ctx, data, i); err != nil {
		return nil, fmt.Errorf("failed to read connection: %v", err)
	}
	return []*schema.ResourceData{data}, nil
}

func resourceExchangeRead(ctx context.Context, data *schema.ResourceData, i interface{}) diag.Diagnostics {
	svc := i.(*analyticshub.Service)
	client := analyticshub.NewProjectsLocationsDataExchangesService(svc)

	exchange, err := client.Get(data.Id()).Do()

	if err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("name", exchange.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("display_name", exchange.DisplayName); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("description", exchange.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("primary_contact", exchange.PrimaryContact); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("documentation", exchange.Documentation); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("listing_count", exchange.ListingCount); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("icon", exchange.Icon); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func expandExchange(data *schema.ResourceData) (*analyticshub.DataExchange, error) {
	displayName := data.Get("display_name").(string)

	exchange := &analyticshub.DataExchange{
		DisplayName: displayName,
	}

	if v, ok := data.GetOk("description"); ok {
		description := v.(string)
		exchange.Description = description
	}

	if v, ok := data.GetOk("primary_contact"); ok {
		primaryContact := v.(string)
		exchange.PrimaryContact = primaryContact
	}

	if v, ok := data.GetOk("documentation"); ok {
		documentation := v.(string)
		exchange.Documentation = documentation
	}

	if v, ok := data.GetOk("listing_count"); ok {
		listingCount := v.(int64)
		exchange.ListingCount = listingCount
	}

	if v, ok := data.GetOk("icon"); ok {
		icon := v.(string)
		exchange.Icon = icon
	}

	return exchange, nil
}

func getIds(data *schema.ResourceData) (string, string) {
	dataExchangeId := data.Get("data_exchange_id").(string)
	project := data.Get("project").(string)
	region := data.Get("region").(string)
	parent := fmt.Sprintf("projects/%s/locations/%s", project, region)
	return parent, dataExchangeId
}
