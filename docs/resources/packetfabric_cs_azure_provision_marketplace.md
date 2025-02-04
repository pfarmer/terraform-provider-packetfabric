---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_cs_azure_provision_marketplace Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_azure_provision_marketplace (Resource)

From the marketplace provider's side, provision a requested Azure hosted connection. For more information, see [Cloud Connections in the PacketFabric documentation](https://docs.packetfabric.com/cloud/) and [Marketplace to Cloud Connections](https://docs.packetfabric.com/eco/marketplace_cloud/).


## Example Usage

```terraform
resource "packetfabric_cs_azure_hosted_marketplace_connection" "cs_marketplace_conn1" {
  provider          = packetfabric
  description       = var.description
  account_uuid      = var.pf_account_uuid
  azure_service_key = var.azure_service_key
  routing_id        = var.routing_id
  market            = var.market
  speed             = var.pf_cs_speed # will be deprecated
}

output "packetfabric_cs_azure_hosted_marketplace_connection" {
  sensitive = true
  value     = packetfabric_cs_azure_hosted_marketplace_connection.cs_marketplace_conn1
}

resource "packetfabric_cs_azure_provision_marketplace" "accept_request_azure" {
  provider        = packetfabric
  description     = var.description
  port_circuit_id = var.port_circuit_id_marketplace
  vc_request_uuid = packetfabric_cs_azure_hosted_marketplace_connection.cs_marketplace_conn1.id
  vlan_private    = var.pf_cs_vlan_private
  vlan_microsoft  = var.pf_cs_vlan_microsoft
}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `port_circuit_id` (String) The circuit ID of the port on which you want to provision the request. This starts with "PF-AP-".
- `vc_request_uuid` (String) UUID of the service request you received from the marketplace customer.

### Optional

- `description` (String) A brief description of this connection.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `vlan_microsoft` (Number) Valid VLAN range is from 4-4094, inclusive.
- `vlan_private` (Number) Valid VLAN range is from 4-4094, inclusive.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)





