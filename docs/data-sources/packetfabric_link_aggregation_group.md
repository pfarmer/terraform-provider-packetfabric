---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_link_aggregation_group Data Source - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_link_aggregation_group (Data Source)



## Example Usage

```terraform
data "packetfabric_link_aggregation_group" "lag_1" {
  provider            = packetfabric
  lag_port_circuit_id = packetfabric_link_aggregation_group.lag_1.id
}

output "packetfabric_link_aggregation_group" {
  value = data.packetfabric_link_aggregation_group.lag_1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `interfaces` (List of Object) (see [below for nested schema](#nestedatt--interfaces))

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- `account_uuid` (String)
- `admin_status` (String)
- `autoneg` (Boolean)
- `customer_name` (String)
- `customer_uuid` (String)
- `description` (String)
- `disabled` (Boolean)
- `is_cloud` (Boolean)
- `is_lag` (Boolean)
- `is_lag_member` (Boolean)
- `is_nni` (Boolean)
- `is_ptp` (Boolean)
- `lag_interval` (String)
- `market` (String)
- `market_description` (String)
- `media` (String)
- `member_count` (Number)
- `mtu` (Number)
- `operational_status` (String)
- `parent_lag_circuit_id` (String)
- `pop` (String)
- `port_circuit_id` (String)
- `region` (String)
- `site` (String)
- `site_code` (String)
- `speed` (String)
- `state` (String)
- `status` (String)
- `subscription_term` (Number)
- `time_created` (String)
- `time_updated` (String)
- `vc_mode` (String)
- `zone` (String)


