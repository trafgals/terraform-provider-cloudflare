// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_event_test

import (
	"context"
	"testing"

	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/services/waiting_room_event"
	"github.com/trafgals/terraform-provider-cloudflare-trafgals/internal/test_helpers"
)

func TestWaitingRoomEventsDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*waiting_room_event.WaitingRoomEventsDataSourceModel)(nil)
	schema := waiting_room_event.ListDataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
