// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
)

// V1ContractNamedScheduleService contains methods and other services that help
// with interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractNamedScheduleService] method instead.
type V1ContractNamedScheduleService struct {
	Options []option.RequestOption
}

// NewV1ContractNamedScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContractNamedScheduleService(opts ...option.RequestOption) (r *V1ContractNamedScheduleService) {
	r = &V1ContractNamedScheduleService{}
	r.Options = opts
	return
}

// Get a named schedule for the given rate card. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1ContractNamedScheduleService) Get(ctx context.Context, body V1ContractNamedScheduleGetParams, opts ...option.RequestOption) (res *V1ContractNamedScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/getNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a named schedule for the given rate card. This endpoint's availability is
// dependent on your client's configuration.
func (r *V1ContractNamedScheduleService) Update(ctx context.Context, body V1ContractNamedScheduleUpdateParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/contract-pricing/rate-cards/updateNamedSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1ContractNamedScheduleGetResponse struct {
	Data []V1ContractNamedScheduleGetResponseData `json:"data,required"`
	JSON v1ContractNamedScheduleGetResponseJSON   `json:"-"`
}

// v1ContractNamedScheduleGetResponseJSON contains the JSON metadata for the struct
// [V1ContractNamedScheduleGetResponse]
type v1ContractNamedScheduleGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractNamedScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractNamedScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractNamedScheduleGetResponseData struct {
	StartingAt   time.Time                                  `json:"starting_at,required" format:"date-time"`
	Value        interface{}                                `json:"value,required"`
	EndingBefore time.Time                                  `json:"ending_before" format:"date-time"`
	JSON         v1ContractNamedScheduleGetResponseDataJSON `json:"-"`
}

// v1ContractNamedScheduleGetResponseDataJSON contains the JSON metadata for the
// struct [V1ContractNamedScheduleGetResponseData]
type v1ContractNamedScheduleGetResponseDataJSON struct {
	StartingAt   apijson.Field
	Value        apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractNamedScheduleGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractNamedScheduleGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractNamedScheduleGetParams struct {
	// ID of the rate card whose named schedule is to be retrieved
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// The identifier for the schedule to be retrieved
	ScheduleName param.Field[string] `json:"schedule_name,required"`
	// If provided, at most one schedule segment will be returned (the one that covers
	// this date). If not provided, all segments will be returned.
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
}

func (r V1ContractNamedScheduleGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNamedScheduleUpdateParams struct {
	// ID of the rate card whose named schedule is to be updated
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// The identifier for the schedule to be updated
	ScheduleName param.Field[string]    `json:"schedule_name,required"`
	StartingAt   param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// The value to set for the named schedule. The structure of this object is
	// specific to the named schedule.
	Value        param.Field[interface{}] `json:"value,required"`
	EndingBefore param.Field[time.Time]   `json:"ending_before" format:"date-time"`
}

func (r V1ContractNamedScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
