// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package shared_test

import (
	"github.com/gardener/gardener/extensions/pkg/terraformer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	"github.com/gardener/gardener-extension-provider-alicloud/pkg/controller/infrastructure/infraflow/shared"
)

var _ = Describe("TerraformState", func() {
	It("should unmarshall terraformer state", func() {
		var ()

		rawState := &terraformer.RawState{
			Data:     tfstate,
			Encoding: "none",
		}

		tf, err := shared.UnmarshalTerraformStateFromTerraformer(rawState)
		Expect(err).NotTo(HaveOccurred())

		Expect(tf.Outputs["vpc_id"]).To(Equal(shared.TFOutput{Type: "string", Value: "vpc-0123456"}))

		tables := tf.FindManagedResourcesByType("fake_routetable")
		Expect(tables).To(HaveLen(2))

		Expect(tf.GetManagedResourceInstanceID("fake_routetable", "routetable_private_utility_z0")).
			To(Equal(ptr.To("rtb-77777")))

		Expect(tf.GetManagedResourceInstanceName("fake_resource_a", "nodes")).
			To(Equal(ptr.To("shoot--foo--bar-nodes")))

		Expect(tf.GetManagedResourceInstanceAttribute("fake_natgateway", "natgw_z0", "private_ip")).
			To(Equal(ptr.To("10.100.201.202")))

		Expect(tf.GetManagedResourceInstanceAttribute("fake_natgateway", "natgw_z0", "foobar")).
			To(BeNil())

		Expect(tf.GetManagedResourceInstances("fake_routetable")).
			To(Equal(map[string]string{
				"routetable_main":               "rtb-66666",
				"routetable_private_utility_z0": "rtb-77777",
			}))
	})
})

const tfstate = `{
  "version": 4,
  "terraform_version": "0.15.5",
  "serial": 83,
  "lineage": "674a5a9a-d0e5-eee1-ce57-d820c4313bf0",
  "outputs": {
    "vpc_id": {
      "value": "vpc-0123456",
      "type": "string"
    }
  },
  "resources": [
    {
      "mode": "managed",
      "type": "fake_resource_a",
      "name": "nodes",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "shoot--foo--bar-nodes-id",
            "max_session_duration": 3600,
            "name": "shoot--foo--bar-nodes"
          }
        }
      ]
    },
    {
      "mode": "managed",
      "type": "fake_natgateway",
      "name": "natgw_z0",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "allocation_id": "eipalloc-07aaaaa",
            "connectivity_type": "public",
            "id": "nat-22222",
            "network_interface_id": "eni-33333",
            "private_ip": "10.100.201.202",
            "public_ip": "1.2.3.4",
            "subnet_id": "subnet-44444",
            "tags": {
              "Name": "shoot--foo--bar-natgw-z0",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "tags_all": {
              "Name": "shoot--foo--bar-natgw-z0",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            }
          },
          "sensitive_attributes": [],
          "private": "xxx",
          "dependencies": [
            "fake_eip.eip_natgw_z0",
            "fake_subnet.public_utility_z0",
            "fake_vpc.vpc"
          ]
        }
      ]
    },
    {
      "mode": "managed",
      "type": "fake_routetable",
      "name": "routetable_main",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "rtb-66666",
            "owner_id": "999999",
            "propagating_vgws": [],
            "route": [
              {
                "cidr_block": "0.0.0.0/0"
              }
            ],
            "tags": {
              "Name": "shoot--foo--bar",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "tags_all": {
              "Name": "shoot--foo--bar",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "timeouts": {
              "create": "5m",
              "delete": null,
              "update": null
            },
            "vpc_id": "vpc-0123456"
          },
          "sensitive_attributes": [],
          "private": "xxx",
          "dependencies": [
            "fake_vpc.vpc"
          ]
        }
      ]
    },
    {
      "mode": "managed",
      "type": "fake_routetable",
      "name": "routetable_private_utility_z0",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "id": "rtb-77777",
            "owner_id": "999999",
            "route": [
              {
                "cidr_block": "0.0.0.0/0",
                "nat_gateway_id": "nat-22222"
              }
            ],
            "tags": {
              "Name": "shoot--foo--bar-private-eu-west-1a",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "tags_all": {
              "Name": "shoot--foo--bar-private-eu-west-1a",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "timeouts": {
              "create": "5m",
              "delete": null,
              "update": null
            },
            "vpc_id": "vpc-0123456"
          },
          "sensitive_attributes": [],
          "private": "xxx",
          "dependencies": [
            "fake_vpc.vpc"
          ]
        }
      ]
    },
    {
      "mode": "managed",
      "type": "fake_vpc",
      "name": "vpc",
      "provider": "provider[\"registry.terraform.io/hashicorp/aws\"]",
      "instances": [
        {
          "schema_version": 1,
          "attributes": {
            "arn": "arn:aws:ec2:eu-west-1:999999:vpc/vpc-0123456",
            "assign_generated_ipv6_cidr_block": false,
            "cidr_block": "10.180.0.0/16",
            "default_security_group_id": "sg-11111",
            "enable_classiclink": false,
            "enable_classiclink_dns_support": false,
            "enable_dns_hostnames": true,
            "enable_dns_support": true,
            "id": "vpc-0123456",
            "instance_tenancy": "default",
            "ipv6_association_id": "",
            "ipv6_cidr_block": "",
            "main_route_table_id": "rtb-77888",
            "owner_id": "999999",
            "tags": {
              "Name": "shoot--foo--bar",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            },
            "tags_all": {
              "Name": "shoot--foo--bar",
              "kubernetes.io/cluster/shoot--foo--bar": "1"
            }
          },
          "sensitive_attributes": [],
          "private": "eyJzY2hlbWFfdmVyc2lvbiI6IjEifQ=="
        }
      ]
    }
  ]
}
`
