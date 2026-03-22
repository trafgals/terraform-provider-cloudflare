// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_dns_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/email_routing_dns"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestEmailRoutingDNSDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*email_routing_dns.EmailRoutingDNSDataSourceModel)(nil)
	schema := email_routing_dns.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
