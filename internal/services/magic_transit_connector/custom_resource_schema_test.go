package magic_transit_connector_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/magic_transit_connector"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestCustomMagicTransitConnectorModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*magic_transit_connector.CustomMagicTransitConnectorModel)(nil)
	schema := magic_transit_connector.CustomResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
