resource "packetfabric_cloud_router" "cr1" {
  provider     = packetfabric
  asn          = var.pf_cr_asn
  name         = var.pf_cr_name
  account_uuid = var.pf_account_uuid
  capacity     = var.pf_cr_capacity
  regions      = var.pf_cr_regions
}

resource "packetfabric_aws_cloud_router_connection" "crc1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  account_uuid   = var.pf_account_uuid
  aws_account_id = var.pf_aws_account_id
  maybe_nat      = var.pf_crc_maybe_nat
  description    = var.pf_crc_description
  pop            = var.pf_crc_pop
  zone           = var.pf_crc_zone
  is_public      = var.pf_crc_is_public
  speed          = var.pf_crc_speed
}

resource "packetfabric_cloud_router_bgp_session" "cr_bgp1" {
  provider       = packetfabric
  circuit_id     = packetfabric_cloud_router.cr1.id
  connection_id  = packetfabric_aws_cloud_router_connection.crc1.id
  address_family = var.pf_crbs_af
  multihop_ttl   = var.pf_crbs_mhttl
  remote_asn     = var.pf_crbs_rasn
  orlonger       = var.pf_crbs_orlonger
  remote_address = var.pf_crbs_remoteaddr
  l3_address     = var.pf_crbs_l3addr
  md5            = var.pf_crbs_md5
}

output "packetfabric_cloud_router_bgp_session" {
  value = packetfabric_cloud_router_bgp_session.cr_bgp1
}
