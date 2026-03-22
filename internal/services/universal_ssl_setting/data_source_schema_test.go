// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package universal_ssl_setting_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/universal_ssl_setting"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestUniversalSSLSettingDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*universal_ssl_setting.UniversalSSLSettingDataSourceModel)(nil)
	schema := universal_ssl_setting.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
