// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector_rules_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/cloud_connector_rules"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestCloudConnectorRulesModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*cloud_connector_rules.CloudConnectorRulesModel)(nil)
	schema := cloud_connector_rules.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
