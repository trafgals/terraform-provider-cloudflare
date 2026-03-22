// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippet_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/snippet"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestSnippetsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*snippet.SnippetsDataSourceModel)(nil)
	schema := snippet.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
