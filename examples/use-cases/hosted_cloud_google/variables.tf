## General VARs
variable "tag_name" {
  default = "demo-pf-google"
}

## PacketFabic VARs
variable "pf_api_key" {
  type        = string
  description = "PacketFabric platform API access key"
  sensitive   = true
}
variable "pf_account_uuid" {
  type = string
}
variable "pf_api_server" {
  type        = string
  default     = "https://api.packetfabric.com"
  description = "PacketFabric API endpoint URL"
}
# GCP Hosted Connection
variable "pf_port_circuit_id" {
  type    = string
  default = "PF-AP-WDC1-1726464"
}
variable "pf_cs_pop1" {
  type    = string
  default = "SFO1"
}
variable "pf_cs_speed" {
  type    = string
  default = "50Mbps"
}
variable "pf_cs_vlan1" {
  type    = number
  default = 105
}

# GCP VARs
variable "gcp_project_id" {
  type        = string
  sensitive   = true
  description = "Google Cloud project ID"
}

variable "gcp_credentials" {
  type        = string
  description = "Google Cloud service account credentials"
}
# https://cloud.google.com/compute/docs/regions-zones
variable "gcp_region1" {
  type        = string
  default     = "us-west1"
  description = "Google Cloud region"
}
variable "gcp_zone1" {
  type        = string
  default     = "us-west1-a"
  description = "Google Cloud zone"
}
variable "subnet_cidr1" {
  type        = string
  description = "CIDR for the subnet"
  default     = "10.5.1.0/24"
}

# BGP peering
variable "peer_asn" {
  type    = number
  default = 16550
}