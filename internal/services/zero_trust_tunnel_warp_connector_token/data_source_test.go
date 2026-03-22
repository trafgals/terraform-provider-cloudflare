package zero_trust_tunnel_warp_connector_token_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/acctest"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWARPConnectorTokenDatasource_Basic(t *testing.T) {
	accID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	rnd := utils.GenerateRandomResourceName()
	name := fmt.Sprintf("data.cloudflare_zero_trust_tunnel_warp_connector_token.%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.TestAccPreCheck(t)
		},
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckWARPConnectorTokenDatasourceBasic(accID, rnd),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "token"),
				),
			},
		},
	})
}

func testAccCheckWARPConnectorTokenDatasourceBasic(accID, name string) string {
	return acctest.LoadTestCase("basic.tf", accID, name)
}
