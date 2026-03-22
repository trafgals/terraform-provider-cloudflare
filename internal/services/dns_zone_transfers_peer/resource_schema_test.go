// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_peer_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/dns_zone_transfers_peer"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestDNSZoneTransfersPeerModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_zone_transfers_peer.DNSZoneTransfersPeerModel)(nil)
	schema := dns_zone_transfers_peer.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
