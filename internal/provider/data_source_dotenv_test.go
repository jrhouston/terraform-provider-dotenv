package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const testAccDataSourceDotEnvFilename = `
data "dotenv" "test" {
  filename = "./testdata/test.env"
}
`

func TestAccDataSourceDotEnvFilename(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDotEnvFilename,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.dotenv.test", "env.PGUSER", "username"),
					resource.TestCheckResourceAttr("data.dotenv.test", "env.PGHOST", "localhost"),
					resource.TestCheckResourceAttr("data.dotenv.test", "env.PGDATABASE", "main"),
					resource.TestCheckResourceAttr("data.dotenv.test", "env.PGPASSWORD", "p@ssw0rd"),
				),
			},
		},
	})
}

const testAccDataSourceDotEnvString = `
data "dotenv" "test" {
  string = "USERNAME=test\nPASSWORD=test"
}
`

func TestAccDataSourceDotEnvString(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceDotEnvString,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.dotenv.test", "env.USERNAME", "test"),
					resource.TestCheckResourceAttr("data.dotenv.test", "env.PASSWORD", "test"),
				),
			},
		},
	})
}
