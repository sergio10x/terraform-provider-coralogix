// Copyright 2024 Coralogix Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package coralogix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var rulesGroupDataSourceName = "data." + rulesGroupResourceName

func TestAccCoralogixDataSourceRuleGroup_basic(t *testing.T) {
	r := getRandomRuleGroup()
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { TestAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCoralogixDataSourceRuleGroup_basic(r) +
					testAccCoralogixDataSourceRuleGroup_read(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(rulesGroupDataSourceName, "name", r.name),
					resource.TestCheckResourceAttr(rulesGroupDataSourceName, "rule_subgroups.0.rules.0.extract.0.name", r.ruleParams.name),
				),
			},
		},
	})
}

func testAccCoralogixDataSourceRuleGroup_basic(r *ruleGroupParams) string {
	return fmt.Sprintf(`resource "coralogix_rules_group" "test" {
  name         = "%s"
  description  = "%s"
  creator      = "%s"
  rule_subgroups {
    rules {
     extract {
      name               = "%s"
      description        = "%s"
      source_field       = "text"
      regular_expression = "(?P<remote_addr>\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3})\\s*-\\s*(?P<user>[^ ]+)\\s*\\[(?P<timestemp>\\d{4}-\\d{2}\\-\\d{2}T\\d{2}\\:\\d{2}\\:\\d{2}\\.\\d{1,6}Z)\\]\\s*\\\\\\\"(?P<method>[A-z]+)\\s[\\/\\\\]+(?P<request>[^\\s]+)\\s*(?P<protocol>[A-z0-9\\/\\.]+)\\\\\\\"\\s*(?P<status>\\d+)\\s*(?P<body_bytes_sent>\\d+)?\\s*?\\\\\\\"(?P<http_referer>[^\"]+)\\\"\\s*\\\\\\\"(?P<http_user_agent>[^\"]+)\\\"\\s(?P<request_time>\\d{1,6})\\s*(?P<response_time>\\d{1,6})"
    }
   }
  }
 }
`, r.name, r.description, r.creator, r.ruleParams.name, r.ruleParams.description)
}

func testAccCoralogixDataSourceRuleGroup_read() string {
	return `data "coralogix_rules_group" "test" {
	id = coralogix_rules_group.test.id
}
`
}
