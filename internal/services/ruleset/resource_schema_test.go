// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/ruleset"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestRulesetModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*ruleset.RulesetModel)(nil)
	schema := ruleset.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
