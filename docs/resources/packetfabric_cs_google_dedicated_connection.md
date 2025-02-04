---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_port Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_cs_google_dedicated_connection (Resource)

This is a port located in a Google cloud on-ramp facility, which will be connected to the Google network via cross connect. For more information, see [Cloud Connections in the PacketFabric documentation](https://docs.packetfabric.com/cloud/).

## Example Usage

```terraform
resource "packetfabric_cs_google_dedicated_connection" "pf_cs_conn1_dedicated_google" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  description       = var.description
  zone              = var.pf_cs_zone
  pop               = var.pf_cs_pop
  subscription_term = var.pf_cs_subterm
  service_class     = var.pf_cs_srvclass
  autoneg           = var.pf_cs_autoneg
  speed             = var.pf_cs_speed4
}

output "packetfabric_cs_google_dedicated_connection" {
  value = data.packetfabric_cs_google_dedicated_connection.pf_cs_conn1_dedicated_google
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed.
- `description` (String) A brief description of this connection.
- `pop` (String) The POP in which the dedicated port should be provisioned (the cloud on-ramp).
- `service_class` (String) The service class for the given port, either long haul or metro. Specify metro if the cloud on-ramp (the `pop`) is in the same market as the source ports (the ports to which you will be building out virtual circuits).

	Enum: ["longhaul" "metro"]
- `speed` (String) The desired capacity of the port.

	Enum: ["10Gps", "100Gbps"]
- `subscription_term` (Number) The billing term, in months, for this connection.

	Enum: ["1", "12", "24", "36"]

### Optional

- `autoneg` (Boolean) Whether the port auto-negotiates or not. This is currently only possible with 1Gbps ports and the request will fail if specified with 10Gbps.
- `loa` (String) A base64 encoded string of a PDF of the LOA that Google provided.
- `published_quote_line_uuid` (String) UUID of the published quote line with which this connection should be associated.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `zone` (String) The desired zone of the new connection.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


