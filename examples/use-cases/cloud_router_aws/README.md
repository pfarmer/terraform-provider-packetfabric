# Use Case: PacketFabric Cloud Router with AWS

This use case shows an example on how to use the PacketFabric & AWS Terraform providers 
to automate the connection setup between 2 AWS regions using PacketFabric Cloud Router.

![Deployment Diagram](./images/diagram.png)

## Useful links

- [PacketFabric Terraform Docs](https://docs.packetfabric.com/api/terraform/)
- [PacketFabric Cloud Router Docs](https://docs.packetfabric.com/cr/)
- [PacketFabric Terraform Provider](https://registry.terraform.io/providers/PacketFabric/packetfabric)
- [HashiCorp AWS Terraform Provider](https://registry.terraform.io/providers/hashicorp/aws)
- [HashiCorp Random Terraform Provider](https://registry.terraform.io/providers/hashicorp/random)
- [Best practices for using Terraform](https://cloud.google.com/docs/terraform/best-practices-for-terraform)
- [Automate Terraform with GitHub Actions](https://learn.hashicorp.com/tutorials/terraform/github-actions?in=terraform/automation)

## Terraform resources deployed

- resource **"random_pet"**: Get a random pet name (use to name objects created)
- resource **"aws_vpc"**: Create VPC in 2 AWS regions
- resource **"aws_subnet"**: Create subnet in VPCs
- resource **"aws_internet_gateway"**: Create internet gateway (used to access future EC2 instances)
- resource **"aws_vpn_gateway"**: Create Virtual Private Gateway (or Private VIF - Virtual Interface)
- resource **"aws_route_table"**: Create route table for the VPCs
- resource **"aws_route_table_association"**: Associate Route Table to the VPCs subnets
- resource **"aws_security_group"**: Create Security groups for future EC2 instances
- resource **"aws_network_interface"**: Create NICs for future EC2 instances
- resource **"aws_key_pair"**: Create SSH key Pair for future EC2 instances
- resource **"aws_instance"**: Create demo EC2 instances with [iperf3](https://github.com/esnet/iperf) and [locust](https://locust.io/)
- resource **"aws_eip"**: Associate a Public IP to the EC2 instances (so you can access it)
- resource **"packetfabric_cloud_router"**: Create the Cloud Router in PacketFabric NaaS
- resource & data source **"packetfabric_aws_cloud_router_connection"**: Create Cloud Router Connection to the 2 AWS regions (and PacketFabric Dedicated Port in the future)
- resource **"time_sleep" "wait_60_seconds"**: Wait few seconds for the Connections to appear on AWS side
- data source **"aws_dx_connection"**: Retrieve Direct Connect Connection details
- resource **"aws_dx_connection_confirmation"**: Accept the connections coming from PacketFabric
- resource **"aws_dx_gateway"**: Create Direct Connect Gateways
- resource **"aws_dx_private_virtual_interface"**: Create Direct Connect Private Virtual interfaces
- resource **"aws_dx_gateway_association"**: Associates a Direct Connect Gateway with a Virtual Private Gateways (VPG) 
- resource **"packetfabric_cloud_router_bgp_session"**: Create BGP sessions in PacketFabric
- resource **"packetfabric_cloud_router_bgp_prefixes"**: Add BGP Prefixes to the BGP sessions in PacketFabric

**Estimated time:** ~15 min for AWS & PacketFabric resources + ~10-15 min for AWS Direct Connect Gateway association with AWS Virtual Private Gateways

**Warning**: Make sure you set the correct AWS region based on the PacketFabric pop selected (find details on location [here](https://packetfabric.com/locations/cloud-on-ramps) and [here](https://aws.amazon.com/directconnect/locations/). Essentially, select the PacketFabric pop the closest to the AWS region you want to connect to. Example: AWS region ``us-west-1`` is the closest to PacketFabric pop ``LAX1``.

## Before You Begin

- Before you begin we recommend you read about the [Terraform basics](https://www.terraform.io/intro)
- Don't have a PacketFabric Account? [Get Started](https://docs.packetfabric.com/intro/)
- Don't have an AWS Account? [Get Started](https://aws.amazon.com/free/)
    - Permissions required: VPC, EC2, Direct Connect

## Prerequisites

Make sure you have installed all of the following prerequisites on your machine:

- [Git](https://git-scm.com/downloads)
- [Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli)

Make sure you have the following items available:

- [AWS Account ID](https://docs.aws.amazon.com/IAM/latest/UserGuide/console_account-alias.html)
- [AWS Access and Secret Keys](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)
- [Packet Fabric Billing Account](https://docs.packetfabric.com/api/examples/account_uuid/)
- [PacketFabric API key](https://docs.packetfabric.com/admin/my_account/keys/)
- [SSH Public Key](https://www.ssh.com/academy/ssh/keygen)

## Quick Start

1. Create the file ``secret.tfvars`` and update each variables.

```sh
cp secret.tfvars.sample secret.tfvars
```

2. Initialize Terraform, create an execution plan and execute the plan.

```sh
terraform init
terraform plan -var-file="secret.tfvars"
```

Apply the plan:

```sh
terraform apply -var-file="secret.tfvars"
```

**Note 1:** You can edit the ``variables.tf`` file and change some of the default variables (default AWS regions used are ``us-west-1`` and ``us-east-1``, PacketFabric pops ``LAX1`` and ``NYC1``).

**Note 2:** Default login/password for Locust is ``demo:packetfabric`` edit ``user-data-ubuntu.sh`` script to change it.

3. Either use and [locust](https://locust.io/) or [iperf3](https://github.com/esnet/iperf) to simulate traffic between the 2 EC2 instances in the 2 AWS regions.

- iperf3: (replace ``<ec2_private_ip_2>`` with the EC2 private IP from instance 2)

    - Server side (on instance 1): ``iperf3 -s -p 5001``
    - Client side (on instance 2): ``iperf3 -c <ec2_private_ip_1> -p 5001``

- locust: (replace ``<ec2_public_ip_1>`` with the EC2 public IP)

In a browser, on instance 1), open ``http://<ec2_public_ip_1>:8089/``, then update the host with the correct IP (using `` <ec2_private_ip_2>`` from instance 2), set the number of users to ``500`` and spawn rate to ``50``.

If you want to use iperf3, open a ssh session using the user ``ubuntu`` and the ssh private key linked to the public key you specified in the ``secret.tfvars`` file.

4. Destroy all remote objects managed by the Terraform configuration.

```sh
terraform destroy -var-file="secret.tfvars"
```

## Troubleshooting

In case you get the following error:

```
╷
│ Error: error waiting for Direct Connection Connection (dxcon-fgq3o1ff) confirm: timeout while waiting for state to become 'available' (last state: 'pending', timeout: 10m0s)
│ 
│   with aws_dx_connection_confirmation.confirmation_2,
│   on main.tf line 451, in resource "aws_dx_connection_confirmation" "confirmation_2":
│  451: resource "aws_dx_connection_confirmation" "confirmation_2" {
│ 
```

You are hitting a timeout issue in AWS [aws_dx_connection_confirmation](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/dx_connection_confirmation) resource. Please [vote](https://github.com/hashicorp/terraform-provider-aws/issues/26335) for this issue on GitHub.

As a workaround, edit the `main.tf` and comment out the following resource:

```
# resource "aws_dx_connection_confirmation" "confirmation_2" {
#   provider      = aws.region2
#   connection_id = data.aws_dx_connection.current_2.id
# }
```

And comment out the dependency with `confirmation_2` in `packetfabric_aws_cloud_router_connection` data source: 

```
data "packetfabric_aws_cloud_router_connection" "current" {
  provider   = packetfabric
  circuit_id = packetfabric_cloud_router.cr.id

  depends_on = [
    aws_dx_connection_confirmation.confirmation_1,
    # aws_dx_connection_confirmation.confirmation_2,
  ]
}
```

Then remove the `confirmation_2` state and re-apply the terraform plan:
```
terraform state rm aws_dx_connection_confirmation.confirmation_2
terraform apply -var-file="secret.tfvars"
```

## Screenshots

Example AWS Direct Connect Connections:

![AWS Direct Connect Connections](./images/aws_direct_connect_connections.png)

Example AWS Direct Connect gateway:

![AWS Direct Connect gateway](./images/aws_direct_connect_gateway.png)

Example Direct Connect Private Virtual interfaces:

![AWS Direct Connect Private Virtual interfaces](./images/aws_direct_connect_private_virtual_interfaces.png)

Example PacketFabric Cloud Router AWS Connections:

![PacketFabric Cloud Router AWS Connections](./images/packetfabric_cloud_router_connections_aws.png)

Example AWS VPC Routing Table:

![AWS VPC Routing Table](./images/aws_vpc_routing_table.png)

Traffic Generator using Locust: *Response Time ~77ms*

![Demo Locust AWS PacketFabric](./images/demo_aws_traffic_direct_connect_through_PacketFabric_500_users_locust.png)
