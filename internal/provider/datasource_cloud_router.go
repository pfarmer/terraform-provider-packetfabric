package provider

import (
	"context"

	"github.com/PacketFabric/terraform-provider-packetfabric/internal/packetfabric"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudRouter() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudRouterRead,
		Schema: map[string]*schema.Schema{
			"cloud_routers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Whether the cloud router is private or public. Deprecated.",
						},
						"asn": {
							Type:        schema.TypeInt,
							Computed:    true,
							Optional:    true,
							Description: "The configured ASN of the instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The name of this particular CloudRouter.",
						},
						"regions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Optional:    true,
							Description: "List of PacketFabric Reigions for the Cloud Router.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Common Name of POP location.",
									},
									"code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "PacketFabric POP code.",
									},
								},
							},
						},
						"capacity": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "The total configured capacity of this particular Cloud Router.",
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudRouterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*packetfabric.PFClient)
	c.Ctx = ctx
	var diags diag.Diagnostics
	routers, err := c.ListCloudRouters()
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("cloud_routers", flattenCloudRouters(&routers))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(uuid.New().String())
	return diags
}

func flattenCloudRouters(routers *[]packetfabric.CloudRouterResponse) []interface{} {
	if routers != nil {
		flattens := make([]interface{}, len(*routers), len(*routers))

		for i, router := range *routers {
			flatten := make(map[string]interface{})
			flatten["scope"] = router.Scope
			flatten["asn"] = router.Asn
			flatten["name"] = router.Name
			flatten["capacity"] = router.Capacity
			flatten["regions"] = flattenRegions(&router.Regions)
			flatten["time_created"] = router.TimeCreated
			flatten["time_updated"] = router.TimeUpdated
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}

func flattenRegions(regions *[]packetfabric.Region) []interface{} {
	if regions != nil {
		flattens := make([]interface{}, len(*regions), len(*regions))

		for i, region := range *regions {
			flatten := make(map[string]interface{})
			flatten["name"] = region.Name
			flatten["code"] = region.Code
			flattens[i] = flatten
		}
		return flattens
	}
	return make([]interface{}, 0)
}
