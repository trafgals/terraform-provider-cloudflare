// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/api_token"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestAPITokensDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*api_token.APITokensDataSourceModel)(nil)
	schema := api_token.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
