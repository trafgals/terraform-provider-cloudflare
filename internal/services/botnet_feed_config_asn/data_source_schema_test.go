// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed_config_asn_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/botnet_feed_config_asn"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestBotnetFeedConfigASNDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*botnet_feed_config_asn.BotnetFeedConfigASNDataSourceModel)(nil)
	schema := botnet_feed_config_asn.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
