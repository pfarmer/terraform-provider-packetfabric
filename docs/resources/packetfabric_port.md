---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "packetfabric_port Resource - terraform-provider-packetfabric"
subcategory: ""
description: |-
  
---

# packetfabric_port (Resource)

A port on the PacketFabric network. For more information, see [Ports in the PacketFabric documentation](https://docs.packetfabric.com/ports/).

## Example Usage

```terraform
resource "packetfabric_port" "port_1" {
  provider          = packetfabric
  account_uuid      = var.pf_account_uuid
  autoneg           = var.pf_port_autoneg
  description       = var.description
  media             = var.pf_port_media
  nni               = var.pf_port_nni
  pop               = var.pf_port_pop1
  speed             = var.pf_port_speed
  subscription_term = var.pf_port_subterm
  zone              = var.pf_port_avzone1
}
output "packetfabric_port_1" {
  value = packetfabric_port.port_1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_uuid` (String) The UUID for the billing account that should be billed.
- `autoneg` (Boolean) Only applicable to 1Gbps ports. Controls whether auto negotiation is on (true) or off (false). The request will fail if specified with 10Gbps.
- `description` (String) A brief description of the port.
- `media` (String) Optic media type.

	Enum: ["LX" "EX" "ZX" "LR" "ER" "ER DWDM" "ZR" "ZR DWDM" "LR4" "ER4" "CWDM4" "LR4" "ER4 Lite"]
- `nni` (Boolean) Set this to true to provision an ENNI port. ENNI ports will use a nni_svlan_tpid value of 0x8100.

	By default, ENNI ports are not available to all users. If you are provisioning your first ENNI port and are unsure if you have permission, contact support@packetfabric.com.
- `pop` (String) Point of presence in which the port should be located.
- `speed` (String) Speed of the port.

	Enum: ["1Gbps" "10Gbps" "40Gbps" "100Gbps"]
- `subscription_term` (Number) Duration of the subscription in months

	Enum ["1" "12" "24" "36"]
- `zone` (String) Availability zone of the port.

### Optional

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


