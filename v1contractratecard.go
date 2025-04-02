// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/packages/pagination"
	"github.com/Metronome-Industries/metronome-go/shared"
)

// V1ContractRateCardService contains methods and other services that help with
// interacting with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractRateCardService] method instead.
type V1ContractRateCardService struct {
	Options        []option.RequestOption
	ProductOrders  *V1ContractRateCardProductOrderService
	Rates          *V1ContractRateCardRateService
	NamedSchedules *V1ContractRateCardNamedScheduleService
}

// NewV1ContractRateCardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContractRateCardService(opts ...option.RequestOption) (r *V1ContractRateCardService) {
	r = &V1ContractRateCardService{}
	r.Options = opts
	r.ProductOrders = NewV1ContractRateCardProductOrderService(opts...)
	r.Rates = NewV1ContractRateCardRateService(opts...)
	r.NamedSchedules = NewV1ContractRateCardNamedScheduleService(opts...)
	return
}

// Create a new rate card
func (r *V1ContractRateCardService) New(ctx context.Context, body V1ContractRateCardNewParams, opts ...option.RequestOption) (res *V1ContractRateCardNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific rate card NOTE: Use `/contract-pricing/rate-cards/getRates` to
// retrieve rate card rates.
func (r *V1ContractRateCardService) Get(ctx context.Context, body V1ContractRateCardGetParams, opts ...option.RequestOption) (res *V1ContractRateCardGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a rate card
func (r *V1ContractRateCardService) Update(ctx context.Context, body V1ContractRateCardUpdateParams, opts ...option.RequestOption) (res *V1ContractRateCardUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List rate cards NOTE: Use `/contract-pricing/rate-cards/getRates` to retrieve
// rate card rates.
func (r *V1ContractRateCardService) List(ctx context.Context, params V1ContractRateCardListParams, opts ...option.RequestOption) (res *pagination.CursorPage[V1ContractRateCardListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/contract-pricing/rate-cards/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List rate cards NOTE: Use `/contract-pricing/rate-cards/getRates` to retrieve
// rate card rates.
func (r *V1ContractRateCardService) ListAutoPaging(ctx context.Context, params V1ContractRateCardListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[V1ContractRateCardListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, params, opts...))
}

// Get all rates for a rate card from starting_at (either in perpetuity or until
// ending_before, if provided)
func (r *V1ContractRateCardService) GetRateSchedule(ctx context.Context, params V1ContractRateCardGetRateScheduleParams, opts ...option.RequestOption) (res *V1ContractRateCardGetRateScheduleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contract-pricing/rate-cards/getRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type V1ContractRateCardNewResponse struct {
	Data shared.ID                         `json:"data,required"`
	JSON v1ContractRateCardNewResponseJSON `json:"-"`
}

// v1ContractRateCardNewResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardNewResponse]
type v1ContractRateCardNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetResponse struct {
	Data V1ContractRateCardGetResponseData `json:"data,required"`
	JSON v1ContractRateCardGetResponseJSON `json:"-"`
}

// v1ContractRateCardGetResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardGetResponse]
type v1ContractRateCardGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetResponseData struct {
	ID                    string                                                  `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                               `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                                  `json:"created_by,required"`
	Name                  string                                                  `json:"name,required"`
	Aliases               []V1ContractRateCardGetResponseDataAlias                `json:"aliases"`
	CreditTypeConversions []V1ContractRateCardGetResponseDataCreditTypeConversion `json:"credit_type_conversions"`
	CustomFields          map[string]string                                       `json:"custom_fields"`
	Description           string                                                  `json:"description"`
	FiatCreditType        shared.CreditTypeData                                   `json:"fiat_credit_type"`
	JSON                  v1ContractRateCardGetResponseDataJSON                   `json:"-"`
}

// v1ContractRateCardGetResponseDataJSON contains the JSON metadata for the struct
// [V1ContractRateCardGetResponseData]
type v1ContractRateCardGetResponseDataJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CreatedBy             apijson.Field
	Name                  apijson.Field
	Aliases               apijson.Field
	CreditTypeConversions apijson.Field
	CustomFields          apijson.Field
	Description           apijson.Field
	FiatCreditType        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1ContractRateCardGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetResponseDataAlias struct {
	Name         string                                     `json:"name,required"`
	EndingBefore time.Time                                  `json:"ending_before" format:"date-time"`
	StartingAt   time.Time                                  `json:"starting_at" format:"date-time"`
	JSON         v1ContractRateCardGetResponseDataAliasJSON `json:"-"`
}

// v1ContractRateCardGetResponseDataAliasJSON contains the JSON metadata for the
// struct [V1ContractRateCardGetResponseDataAlias]
type v1ContractRateCardGetResponseDataAliasJSON struct {
	Name         apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractRateCardGetResponseDataAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetResponseDataAliasJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetResponseDataCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData                                     `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                                                    `json:"fiat_per_custom_credit,required"`
	JSON                v1ContractRateCardGetResponseDataCreditTypeConversionJSON `json:"-"`
}

// v1ContractRateCardGetResponseDataCreditTypeConversionJSON contains the JSON
// metadata for the struct [V1ContractRateCardGetResponseDataCreditTypeConversion]
type v1ContractRateCardGetResponseDataCreditTypeConversionJSON struct {
	CustomCreditType    apijson.Field
	FiatPerCustomCredit apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1ContractRateCardGetResponseDataCreditTypeConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetResponseDataCreditTypeConversionJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardUpdateResponse struct {
	Data shared.ID                            `json:"data,required"`
	JSON v1ContractRateCardUpdateResponseJSON `json:"-"`
}

// v1ContractRateCardUpdateResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardUpdateResponse]
type v1ContractRateCardUpdateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardListResponse struct {
	ID                    string                                               `json:"id,required" format:"uuid"`
	CreatedAt             time.Time                                            `json:"created_at,required" format:"date-time"`
	CreatedBy             string                                               `json:"created_by,required"`
	Name                  string                                               `json:"name,required"`
	Aliases               []V1ContractRateCardListResponseAlias                `json:"aliases"`
	CreditTypeConversions []V1ContractRateCardListResponseCreditTypeConversion `json:"credit_type_conversions"`
	CustomFields          map[string]string                                    `json:"custom_fields"`
	Description           string                                               `json:"description"`
	FiatCreditType        shared.CreditTypeData                                `json:"fiat_credit_type"`
	JSON                  v1ContractRateCardListResponseJSON                   `json:"-"`
}

// v1ContractRateCardListResponseJSON contains the JSON metadata for the struct
// [V1ContractRateCardListResponse]
type v1ContractRateCardListResponseJSON struct {
	ID                    apijson.Field
	CreatedAt             apijson.Field
	CreatedBy             apijson.Field
	Name                  apijson.Field
	Aliases               apijson.Field
	CreditTypeConversions apijson.Field
	CustomFields          apijson.Field
	Description           apijson.Field
	FiatCreditType        apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1ContractRateCardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardListResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardListResponseAlias struct {
	Name         string                                  `json:"name,required"`
	EndingBefore time.Time                               `json:"ending_before" format:"date-time"`
	StartingAt   time.Time                               `json:"starting_at" format:"date-time"`
	JSON         v1ContractRateCardListResponseAliasJSON `json:"-"`
}

// v1ContractRateCardListResponseAliasJSON contains the JSON metadata for the
// struct [V1ContractRateCardListResponseAlias]
type v1ContractRateCardListResponseAliasJSON struct {
	Name         apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractRateCardListResponseAlias) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardListResponseAliasJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardListResponseCreditTypeConversion struct {
	CustomCreditType    shared.CreditTypeData                                  `json:"custom_credit_type,required"`
	FiatPerCustomCredit string                                                 `json:"fiat_per_custom_credit,required"`
	JSON                v1ContractRateCardListResponseCreditTypeConversionJSON `json:"-"`
}

// v1ContractRateCardListResponseCreditTypeConversionJSON contains the JSON
// metadata for the struct [V1ContractRateCardListResponseCreditTypeConversion]
type v1ContractRateCardListResponseCreditTypeConversionJSON struct {
	CustomCreditType    apijson.Field
	FiatPerCustomCredit apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1ContractRateCardListResponseCreditTypeConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardListResponseCreditTypeConversionJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetRateScheduleResponse struct {
	Data     []V1ContractRateCardGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                          `json:"next_page,nullable"`
	JSON     v1ContractRateCardGetRateScheduleResponseJSON   `json:"-"`
}

// v1ContractRateCardGetRateScheduleResponseJSON contains the JSON metadata for the
// struct [V1ContractRateCardGetRateScheduleResponse]
type v1ContractRateCardGetRateScheduleResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardGetRateScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetRateScheduleResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetRateScheduleResponseData struct {
	Entitled            bool                                                          `json:"entitled,required"`
	ProductCustomFields map[string]string                                             `json:"product_custom_fields,required"`
	ProductID           string                                                        `json:"product_id,required" format:"uuid"`
	ProductName         string                                                        `json:"product_name,required"`
	ProductTags         []string                                                      `json:"product_tags,required"`
	Rate                shared.Rate                                                   `json:"rate,required"`
	StartingAt          time.Time                                                     `json:"starting_at,required" format:"date-time"`
	BillingFrequency    V1ContractRateCardGetRateScheduleResponseDataBillingFrequency `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         V1ContractRateCardGetRateScheduleResponseDataCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                               `json:"ending_before" format:"date-time"`
	PricingGroupValues map[string]string                                       `json:"pricing_group_values"`
	JSON               v1ContractRateCardGetRateScheduleResponseDataJSON       `json:"-"`
}

// v1ContractRateCardGetRateScheduleResponseDataJSON contains the JSON metadata for
// the struct [V1ContractRateCardGetRateScheduleResponseData]
type v1ContractRateCardGetRateScheduleResponseDataJSON struct {
	Entitled            apijson.Field
	ProductCustomFields apijson.Field
	ProductID           apijson.Field
	ProductName         apijson.Field
	ProductTags         apijson.Field
	Rate                apijson.Field
	StartingAt          apijson.Field
	BillingFrequency    apijson.Field
	CommitRate          apijson.Field
	EndingBefore        apijson.Field
	PricingGroupValues  apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1ContractRateCardGetRateScheduleResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetRateScheduleResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetRateScheduleResponseDataBillingFrequency string

const (
	V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyMonthly   V1ContractRateCardGetRateScheduleResponseDataBillingFrequency = "MONTHLY"
	V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyQuarterly V1ContractRateCardGetRateScheduleResponseDataBillingFrequency = "QUARTERLY"
	V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyAnnual    V1ContractRateCardGetRateScheduleResponseDataBillingFrequency = "ANNUAL"
)

func (r V1ContractRateCardGetRateScheduleResponseDataBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyMonthly, V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyQuarterly, V1ContractRateCardGetRateScheduleResponseDataBillingFrequencyAnnual:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractRateCardGetRateScheduleResponseDataCommitRate struct {
	RateType V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                               `json:"tiers"`
	JSON  v1ContractRateCardGetRateScheduleResponseDataCommitRateJSON `json:"-"`
}

// v1ContractRateCardGetRateScheduleResponseDataCommitRateJSON contains the JSON
// metadata for the struct
// [V1ContractRateCardGetRateScheduleResponseDataCommitRate]
type v1ContractRateCardGetRateScheduleResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractRateCardGetRateScheduleResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractRateCardGetRateScheduleResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType string

const (
	V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeFlat         V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "FLAT"
	V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypePercentage   V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "PERCENTAGE"
	V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeSubscription V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "SUBSCRIPTION"
	V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeTiered       V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "TIERED"
	V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeCustom       V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType = "CUSTOM"
)

func (r V1ContractRateCardGetRateScheduleResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeFlat, V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypePercentage, V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeSubscription, V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeTiered, V1ContractRateCardGetRateScheduleResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
}

type V1ContractRateCardNewParams struct {
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Field[string] `json:"name,required"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases param.Field[[]V1ContractRateCardNewParamsAlias] `json:"aliases"`
	// Required when using custom pricing units in rates.
	CreditTypeConversions param.Field[[]V1ContractRateCardNewParamsCreditTypeConversion] `json:"credit_type_conversions"`
	CustomFields          param.Field[map[string]string]                                 `json:"custom_fields"`
	Description           param.Field[string]                                            `json:"description"`
	// The Metronome ID of the credit type to associate with the rate card, defaults to
	// USD (cents) if not passed.
	FiatCreditTypeID param.Field[string] `json:"fiat_credit_type_id" format:"uuid"`
}

func (r V1ContractRateCardNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardNewParamsAlias struct {
	Name         param.Field[string]    `json:"name,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1ContractRateCardNewParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardNewParamsCreditTypeConversion struct {
	CustomCreditTypeID  param.Field[string]  `json:"custom_credit_type_id,required" format:"uuid"`
	FiatPerCustomCredit param.Field[float64] `json:"fiat_per_custom_credit,required"`
}

func (r V1ContractRateCardNewParamsCreditTypeConversion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardGetParams struct {
	ID shared.IDParam `json:"id,required"`
}

func (r V1ContractRateCardGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ID)
}

type V1ContractRateCardUpdateParams struct {
	// ID of the rate card to update
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// Reference this alias when creating a contract. If the same alias is assigned to
	// multiple rate cards, it will reference the rate card to which it was most
	// recently assigned. It is not exposed to end customers.
	Aliases     param.Field[[]V1ContractRateCardUpdateParamsAlias] `json:"aliases"`
	Description param.Field[string]                                `json:"description"`
	// Used only in UI/API. It is not exposed to end customers.
	Name param.Field[string] `json:"name"`
}

func (r V1ContractRateCardUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardUpdateParamsAlias struct {
	Name         param.Field[string]    `json:"name,required"`
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	StartingAt   param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1ContractRateCardUpdateParamsAlias) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractRateCardListParams struct {
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	Body     interface{}         `json:"body"`
}

func (r V1ContractRateCardListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [V1ContractRateCardListParams]'s query parameters as
// `url.Values`.
func (r V1ContractRateCardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardGetRateScheduleParams struct {
	// ID of the rate card to get the schedule for
	RateCardID param.Field[string] `json:"rate_card_id,required" format:"uuid"`
	// inclusive starting point for the rates schedule
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// optional exclusive end date for the rates schedule. When not specified rates
	// will show all future schedule segments.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selector will be included in
	// the response Passing no selectors will result in all rates being returned.
	Selectors param.Field[[]V1ContractRateCardGetRateScheduleParamsSelector] `json:"selectors"`
}

func (r V1ContractRateCardGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1ContractRateCardGetRateScheduleParams]'s query parameters
// as `url.Values`.
func (r V1ContractRateCardGetRateScheduleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractRateCardGetRateScheduleParamsSelector struct {
	// Subscription rates matching the billing frequency will be included in the
	// response.
	BillingFrequency param.Field[V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency] `json:"billing_frequency"`
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues param.Field[map[string]string] `json:"partial_pricing_group_values"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Rates matching the product id will be included in the response.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
}

func (r V1ContractRateCardGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Subscription rates matching the billing frequency will be included in the
// response.
type V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency string

const (
	V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyMonthly   V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency = "MONTHLY"
	V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyQuarterly V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency = "QUARTERLY"
	V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyAnnual    V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency = "ANNUAL"
)

func (r V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyMonthly, V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyQuarterly, V1ContractRateCardGetRateScheduleParamsSelectorsBillingFrequencyAnnual:
		return true
	}
	return false
}
