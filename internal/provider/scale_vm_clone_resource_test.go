// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccScaleVMCloneResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccScaleVMCloneResourceConfig("tf-vm", "tf-src"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scale_vm_clone.test", "description", ""), // TODO "tf-vm-description"
					resource.TestCheckResourceAttr("scale_vm_clone.test", "memory", "0"),     // TODO 1 GB
					// resource.TestCheckResourceAttr("scale_vm_clone.test", "nics", []),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "group", ""), // TODO "testtf"
					resource.TestCheckResourceAttr("scale_vm_clone.test", "meta_data", ""),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "user_data", ""),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "source_vm_name", "tf-src"),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "name", "tf-vm"),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "vcpu", "0"),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "disk_size", "0"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "scale_vm_clone.test",
				ImportState:       true,
				ImportStateVerify: true,
				// This is not normally necessary, but is here because this
				// example code does not have an actual upstream service.
				// Once the Read method is able to refresh information from
				// the upstream service, this can be removed.
				ImportStateVerifyIgnore: []string{
					"vm_list",
					// TODO do not ignore below attributes
					"name",
					"description",
					"group",
					"vcpu",
					"memory",
					"disk_size",
					"source_vm_name",
					"user_data",
					"meta_data",
				},
			},
			// Update and Read testing
			{
				Config: testAccScaleVMCloneResourceConfig("tf-vm", "tf-src"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scale_vm_clone.test", "name", "tf-vm"),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "description", ""),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "group", ""),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "vcpu", "0"),
					resource.TestCheckResourceAttr("scale_vm_clone.test", "memory", "0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccScaleVMCloneResourceConfig(vm_name string, source_vm_name string) string {
	return fmt.Sprintf(`
resource "scale_vm_clone" "test" {
  name = %[1]q
  source_vm_name = %[2]q
  group = ""
  vcpu = 0
  memory = 0
  user_data = ""
  meta_data = ""
  description = ""
  nics = []
  disk_size = 0
}
`, vm_name, source_vm_name)
}
