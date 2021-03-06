package web_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance/check"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type AppServiceSourceControlResource struct {
}

func TestAccAppServiceSourceControlToken(t *testing.T) {
	data := acceptance.BuildTestData(t, "azurerm_app_service_source_control_token", "test")
	r := AppServiceSourceControlResource{}
	token := strings.ToLower(acctest.RandString(41))
	tokenSecret := strings.ToLower(acctest.RandString(41))

	data.ResourceTest(t, r, []resource.TestStep{
		{
			Config: testAccAppServiceSourceControlToken(token, tokenSecret),
			Check: resource.ComposeTestCheckFunc(
				check.That(data.ResourceName).Key("type").HasValue("GitHub"),
				check.That(data.ResourceName).Key("token").HasValue(token),
				check.That(data.ResourceName).Key("token_secret").HasValue(tokenSecret),
			),
		},
		data.ImportStep(),
	})
}

func (r AppServiceSourceControlResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
	resp, err := client.Web.BaseClient.GetSourceControl(ctx, state.ID)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return utils.Bool(false), nil
		}
		return nil, fmt.Errorf("retrieving %s: %+v", state.ID, err)
	}

	return utils.Bool(resp.SourceControlProperties != nil), nil
}

func testAccAppServiceSourceControlToken(token, tokenSecret string) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_app_service_source_control_token" "test" {
  type         = "GitHub"
  token        = "%s"
  token_secret = "%s"
}
`, token, tokenSecret)
}
