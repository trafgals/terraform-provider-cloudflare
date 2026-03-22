// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/list"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestListsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*list.ListsDataSourceModel)(nil)
	schema := list.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
