---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cs_google_hosted_marketplace_connection Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_google_hosted_marketplace_connection (Resource)

Connect a marketplace provider to your Google cloud environment. For more information, see [Cloud Connections in the PacketFabric documentation](https://docs.packetfabric.com/cloud/) and [Marketplace to Cloud Connections](https://docs.packetfabric.com/eco/marketplace_cloud/).

## Example Usage

```terraform
resource "packetfabric_cs_google_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider                    = packetfabric
  description                 = var.description
  account_uuid                = var.pf_account_uuid
  routing_id                  = var.routing_id
  market                      = var.market
  speed                       = var.pf_cs_speed
  google_pairing_key          = var.google_pairing_key
  google_vlan_attachment_name = var.google_vlan_attachment_name
  pop                         = var.pf_cs_pop

}

output "packetfabric_cs_google_hosted_marketplace_connection" {
  value     = packetfabric_cs_google_hosted_marketplace_connection.cs_marketplace_conn1
  sensitive = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed. This is your billing account, not the marketplace provider's.
- `google_pairing_key` (String) The Google pairing key to use for this connection. This is provided when you create the VLAN attachment from the Google Cloud console.
- `google_vlan_attachment_name` (String) The name you used for your VLAN attachment in Google.
- `market` (String) The market code (e.g. "ATL" or "DAL") in which you would like the marketplace provider to provision their side of the connection.

	If the marketplace provider has services published in the marketplace, you can use the PacketFabric portal to see which POPs they are in. Simply remove the number from the POP to get the market code (e.g. if they offer services in "DAL5", enter "DAL" for the market).
- `pop` (String) The POP in which the hosted connection should be provisioned (the cloud on-ramp).
- `routing_id` (String) The routing ID of the marketplace provider that will be receiving this request.

	Example: TR-1RI-OQ85
- `speed` (String) The speed of the new connection.

	Enum: ["50Mbps", "100Mbps", "200Mbps", "300Mbps", "400Mbps", "500Mbps", "1Gbps", "2Gbps", "5Gbps", "10Gbps"]

### Optional

- `description` (String) A brief description of this connection.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


