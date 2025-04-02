// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/internal/apiquery"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/Metronome-Industries/metronome-go/internal/requestconfig"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
	"github.com/tidwall/gjson"
)

// V1ContractService contains methods and other services that help with interacting
// with the metronome API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContractService] method instead.
type V1ContractService struct {
	Options        []option.RequestOption
	Products       *V1ContractProductService
	RateCards      *V1ContractRateCardService
	NamedSchedules *V1ContractNamedScheduleService
}

// NewV1ContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContractService(opts ...option.RequestOption) (r *V1ContractService) {
	r = &V1ContractService{}
	r.Options = opts
	r.Products = NewV1ContractProductService(opts...)
	r.RateCards = NewV1ContractRateCardService(opts...)
	r.NamedSchedules = NewV1ContractNamedScheduleService(opts...)
	return
}

// Create a new contract
func (r *V1ContractService) New(ctx context.Context, body V1ContractNewParams, opts ...option.RequestOption) (res *V1ContractNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific contract
func (r *V1ContractService) Get(ctx context.Context, body V1ContractGetParams, opts ...option.RequestOption) (res *V1ContractGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/get"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all contracts for a customer
func (r *V1ContractService) List(ctx context.Context, body V1ContractListParams, opts ...option.RequestOption) (res *V1ContractListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add a manual balance entry
func (r *V1ContractService) AddManualBalanceEntry(ctx context.Context, body V1ContractAddManualBalanceEntryParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/contracts/addManualBalanceLedgerEntry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Amend a contract
func (r *V1ContractService) Amend(ctx context.Context, body V1ContractAmendParams, opts ...option.RequestOption) (res *V1ContractAmendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/amend"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Archive a contract
func (r *V1ContractService) Archive(ctx context.Context, body V1ContractArchiveParams, opts ...option.RequestOption) (res *V1ContractArchiveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates historical usage invoices for a contract
func (r *V1ContractService) NewHistoricalInvoices(ctx context.Context, body V1ContractNewHistoricalInvoicesParams, opts ...option.RequestOption) (res *V1ContractNewHistoricalInvoicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/createHistoricalInvoices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List balances (commits and credits).
func (r *V1ContractService) ListBalances(ctx context.Context, body V1ContractListBalancesParams, opts ...option.RequestOption) (res *V1ContractListBalancesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/customerBalances/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the rate schedule for the rate card on a given contract.
func (r *V1ContractService) GetRateSchedule(ctx context.Context, params V1ContractGetRateScheduleParams, opts ...option.RequestOption) (res *V1ContractGetRateScheduleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/getContractRateSchedule"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a new scheduled invoice for Professional Services terms on a contract.
// This endpoint's availability is dependent on your client's configuration.
func (r *V1ContractService) ScheduleProServicesInvoice(ctx context.Context, body V1ContractScheduleProServicesInvoiceParams, opts ...option.RequestOption) (res *V1ContractScheduleProServicesInvoiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/scheduleProServicesInvoice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Set usage filter for a contract
func (r *V1ContractService) SetUsageFilter(ctx context.Context, body V1ContractSetUsageFilterParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/contracts/setUsageFilter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Update the end date of a contract
func (r *V1ContractService) UpdateEndDate(ctx context.Context, body V1ContractUpdateEndDateParams, opts ...option.RequestOption) (res *V1ContractUpdateEndDateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/contracts/updateEndDate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContractNewResponse struct {
	Data shared.ID                 `json:"data,required"`
	JSON v1ContractNewResponseJSON `json:"-"`
}

// v1ContractNewResponseJSON contains the JSON metadata for the struct
// [V1ContractNewResponse]
type v1ContractNewResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractNewResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponse struct {
	Data V1ContractGetResponseData `json:"data,required"`
	JSON v1ContractGetResponseJSON `json:"-"`
}

// v1ContractGetResponseJSON contains the JSON metadata for the struct
// [V1ContractGetResponse]
type v1ContractGetResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseData struct {
	ID         string                               `json:"id,required" format:"uuid"`
	Amendments []V1ContractGetResponseDataAmendment `json:"amendments,required"`
	Current    shared.ContractWithoutAmendments     `json:"current,required"`
	CustomerID string                               `json:"customer_id,required" format:"uuid"`
	Initial    shared.ContractWithoutAmendments     `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// The billing provider configuration associated with a contract.
	CustomerBillingProviderConfiguration V1ContractGetResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices V1ContractGetResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	Subscriptions                   []V1ContractGetResponseDataSubscription                  `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                        `json:"uniqueness_key"`
	JSON          v1ContractGetResponseDataJSON `json:"-"`
}

// v1ContractGetResponseDataJSON contains the JSON metadata for the struct
// [V1ContractGetResponseData]
type v1ContractGetResponseDataJSON struct {
	ID                                   apijson.Field
	Amendments                           apijson.Field
	Current                              apijson.Field
	CustomerID                           apijson.Field
	Initial                              apijson.Field
	ArchivedAt                           apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	Subscriptions                        apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *V1ContractGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataAmendment struct {
	ID               string                   `json:"id,required" format:"uuid"`
	Commits          []shared.Commit          `json:"commits,required"`
	CreatedAt        time.Time                `json:"created_at,required" format:"date-time"`
	CreatedBy        string                   `json:"created_by,required"`
	Overrides        []shared.Override        `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge `json:"scheduled_charges,required"`
	StartingAt       time.Time                `json:"starting_at,required" format:"date-time"`
	Credits          []shared.Credit          `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V1ContractGetResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                 `json:"salesforce_opportunity_id"`
	JSON                    v1ContractGetResponseDataAmendmentJSON `json:"-"`
}

// v1ContractGetResponseDataAmendmentJSON contains the JSON metadata for the struct
// [V1ContractGetResponseDataAmendment]
type v1ContractGetResponseDataAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1ContractGetResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataAmendmentJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                           `json:"aws_account_number"`
	AwsOfferID            string                                                           `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                           `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                        `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                          `json:"fraction"`
	GcpAccountID          string                                                           `json:"gcp_account_id"`
	GcpOfferID            string                                                           `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                           `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                          `json:"reseller_contract_value"`
	StartingAt            time.Time                                                        `json:"starting_at" format:"date-time"`
	JSON                  v1ContractGetResponseDataAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// v1ContractGetResponseDataAmendmentsResellerRoyaltyJSON contains the JSON
// metadata for the struct [V1ContractGetResponseDataAmendmentsResellerRoyalty]
type v1ContractGetResponseDataAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1ContractGetResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAws           V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp           V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
	V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAws, V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService, V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp, V1ContractGetResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// The billing provider configuration associated with a contract.
type V1ContractGetResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	ID              string                                                                       `json:"id" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]interface{}                                            `json:"configuration"`
	JSON          v1ContractGetResponseDataCustomerBillingProviderConfigurationJSON `json:"-"`
}

// v1ContractGetResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [V1ContractGetResponseDataCustomerBillingProviderConfiguration]
type v1ContractGetResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	ID              apijson.Field
	Configuration   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1ContractGetResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, V1ContractGetResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, V1ContractGetResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1ContractGetResponseDataScheduledChargesOnUsageInvoices string

const (
	V1ContractGetResponseDataScheduledChargesOnUsageInvoicesAll V1ContractGetResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r V1ContractGetResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type V1ContractGetResponseDataSubscription struct {
	CollectionSchedule V1ContractGetResponseDataSubscriptionsCollectionSchedule `json:"collection_schedule,required"`
	Proration          V1ContractGetResponseDataSubscriptionsProration          `json:"proration,required"`
	QuantitySchedule   []V1ContractGetResponseDataSubscriptionsQuantitySchedule `json:"quantity_schedule,required"`
	StartingAt         time.Time                                                `json:"starting_at,required" format:"date-time"`
	SubscriptionRate   V1ContractGetResponseDataSubscriptionsSubscriptionRate   `json:"subscription_rate,required"`
	Description        string                                                   `json:"description"`
	EndingBefore       time.Time                                                `json:"ending_before" format:"date-time"`
	Name               string                                                   `json:"name"`
	JSON               v1ContractGetResponseDataSubscriptionJSON                `json:"-"`
}

// v1ContractGetResponseDataSubscriptionJSON contains the JSON metadata for the
// struct [V1ContractGetResponseDataSubscription]
type v1ContractGetResponseDataSubscriptionJSON struct {
	CollectionSchedule apijson.Field
	Proration          apijson.Field
	QuantitySchedule   apijson.Field
	StartingAt         apijson.Field
	SubscriptionRate   apijson.Field
	Description        apijson.Field
	EndingBefore       apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1ContractGetResponseDataSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataSubscriptionJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataSubscriptionsCollectionSchedule string

const (
	V1ContractGetResponseDataSubscriptionsCollectionScheduleAdvance V1ContractGetResponseDataSubscriptionsCollectionSchedule = "ADVANCE"
	V1ContractGetResponseDataSubscriptionsCollectionScheduleArrears V1ContractGetResponseDataSubscriptionsCollectionSchedule = "ARREARS"
)

func (r V1ContractGetResponseDataSubscriptionsCollectionSchedule) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataSubscriptionsCollectionScheduleAdvance, V1ContractGetResponseDataSubscriptionsCollectionScheduleArrears:
		return true
	}
	return false
}

type V1ContractGetResponseDataSubscriptionsProration struct {
	InvoiceBehavior V1ContractGetResponseDataSubscriptionsProrationInvoiceBehavior `json:"invoice_behavior,required"`
	IsProrated      bool                                                           `json:"is_prorated,required"`
	JSON            v1ContractGetResponseDataSubscriptionsProrationJSON            `json:"-"`
}

// v1ContractGetResponseDataSubscriptionsProrationJSON contains the JSON metadata
// for the struct [V1ContractGetResponseDataSubscriptionsProration]
type v1ContractGetResponseDataSubscriptionsProrationJSON struct {
	InvoiceBehavior apijson.Field
	IsProrated      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1ContractGetResponseDataSubscriptionsProration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataSubscriptionsProrationJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataSubscriptionsProrationInvoiceBehavior string

const (
	V1ContractGetResponseDataSubscriptionsProrationInvoiceBehaviorBillImmediately          V1ContractGetResponseDataSubscriptionsProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	V1ContractGetResponseDataSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate V1ContractGetResponseDataSubscriptionsProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r V1ContractGetResponseDataSubscriptionsProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataSubscriptionsProrationInvoiceBehaviorBillImmediately, V1ContractGetResponseDataSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type V1ContractGetResponseDataSubscriptionsQuantitySchedule struct {
	Quantity     float64                                                    `json:"quantity,required"`
	StartingAt   time.Time                                                  `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                  `json:"ending_before" format:"date-time"`
	JSON         v1ContractGetResponseDataSubscriptionsQuantityScheduleJSON `json:"-"`
}

// v1ContractGetResponseDataSubscriptionsQuantityScheduleJSON contains the JSON
// metadata for the struct [V1ContractGetResponseDataSubscriptionsQuantitySchedule]
type v1ContractGetResponseDataSubscriptionsQuantityScheduleJSON struct {
	Quantity     apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractGetResponseDataSubscriptionsQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataSubscriptionsQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataSubscriptionsSubscriptionRate struct {
	BillingFrequency V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency `json:"billing_frequency,required"`
	Product          V1ContractGetResponseDataSubscriptionsSubscriptionRateProduct          `json:"product,required"`
	JSON             v1ContractGetResponseDataSubscriptionsSubscriptionRateJSON             `json:"-"`
}

// v1ContractGetResponseDataSubscriptionsSubscriptionRateJSON contains the JSON
// metadata for the struct [V1ContractGetResponseDataSubscriptionsSubscriptionRate]
type v1ContractGetResponseDataSubscriptionsSubscriptionRateJSON struct {
	BillingFrequency apijson.Field
	Product          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractGetResponseDataSubscriptionsSubscriptionRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataSubscriptionsSubscriptionRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency string

const (
	V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyMonthly   V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency = "MONTHLY"
	V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyQuarterly V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency = "QUARTERLY"
	V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyAnnual    V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency = "ANNUAL"
)

func (r V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyMonthly, V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyQuarterly, V1ContractGetResponseDataSubscriptionsSubscriptionRateBillingFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractGetResponseDataSubscriptionsSubscriptionRateProduct struct {
	ID   string                                                            `json:"id,required" format:"uuid"`
	Name string                                                            `json:"name,required"`
	JSON v1ContractGetResponseDataSubscriptionsSubscriptionRateProductJSON `json:"-"`
}

// v1ContractGetResponseDataSubscriptionsSubscriptionRateProductJSON contains the
// JSON metadata for the struct
// [V1ContractGetResponseDataSubscriptionsSubscriptionRateProduct]
type v1ContractGetResponseDataSubscriptionsSubscriptionRateProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractGetResponseDataSubscriptionsSubscriptionRateProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetResponseDataSubscriptionsSubscriptionRateProductJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponse struct {
	Data []V1ContractListResponseData `json:"data,required"`
	JSON v1ContractListResponseJSON   `json:"-"`
}

// v1ContractListResponseJSON contains the JSON metadata for the struct
// [V1ContractListResponse]
type v1ContractListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseData struct {
	ID         string                                `json:"id,required" format:"uuid"`
	Amendments []V1ContractListResponseDataAmendment `json:"amendments,required"`
	Current    shared.ContractWithoutAmendments      `json:"current,required"`
	CustomerID string                                `json:"customer_id,required" format:"uuid"`
	Initial    shared.ContractWithoutAmendments      `json:"initial,required"`
	// RFC 3339 timestamp indicating when the contract was archived. If not returned,
	// the contract is not archived.
	ArchivedAt   time.Time         `json:"archived_at" format:"date-time"`
	CustomFields map[string]string `json:"custom_fields"`
	// The billing provider configuration associated with a contract.
	CustomerBillingProviderConfiguration V1ContractListResponseDataCustomerBillingProviderConfiguration `json:"customer_billing_provider_configuration"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices V1ContractListResponseDataScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	Subscriptions                   []V1ContractListResponseDataSubscription                  `json:"subscriptions"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey string                         `json:"uniqueness_key"`
	JSON          v1ContractListResponseDataJSON `json:"-"`
}

// v1ContractListResponseDataJSON contains the JSON metadata for the struct
// [V1ContractListResponseData]
type v1ContractListResponseDataJSON struct {
	ID                                   apijson.Field
	Amendments                           apijson.Field
	Current                              apijson.Field
	CustomerID                           apijson.Field
	Initial                              apijson.Field
	ArchivedAt                           apijson.Field
	CustomFields                         apijson.Field
	CustomerBillingProviderConfiguration apijson.Field
	ScheduledChargesOnUsageInvoices      apijson.Field
	Subscriptions                        apijson.Field
	UniquenessKey                        apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *V1ContractListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataAmendment struct {
	ID               string                   `json:"id,required" format:"uuid"`
	Commits          []shared.Commit          `json:"commits,required"`
	CreatedAt        time.Time                `json:"created_at,required" format:"date-time"`
	CreatedBy        string                   `json:"created_by,required"`
	Overrides        []shared.Override        `json:"overrides,required"`
	ScheduledCharges []shared.ScheduledCharge `json:"scheduled_charges,required"`
	StartingAt       time.Time                `json:"starting_at,required" format:"date-time"`
	Credits          []shared.Credit          `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts []shared.Discount `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []shared.ProService `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []V1ContractListResponseDataAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string                                  `json:"salesforce_opportunity_id"`
	JSON                    v1ContractListResponseDataAmendmentJSON `json:"-"`
}

// v1ContractListResponseDataAmendmentJSON contains the JSON metadata for the
// struct [V1ContractListResponseDataAmendment]
type v1ContractListResponseDataAmendmentJSON struct {
	ID                      apijson.Field
	Commits                 apijson.Field
	CreatedAt               apijson.Field
	CreatedBy               apijson.Field
	Overrides               apijson.Field
	ScheduledCharges        apijson.Field
	StartingAt              apijson.Field
	Credits                 apijson.Field
	Discounts               apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	ProfessionalServices    apijson.Field
	ResellerRoyalties       apijson.Field
	SalesforceOpportunityID apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *V1ContractListResponseDataAmendment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataAmendmentJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataAmendmentsResellerRoyalty struct {
	ResellerType          V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	AwsAccountNumber      string                                                            `json:"aws_account_number"`
	AwsOfferID            string                                                            `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                            `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                                         `json:"ending_before,nullable" format:"date-time"`
	Fraction              float64                                                           `json:"fraction"`
	GcpAccountID          string                                                            `json:"gcp_account_id"`
	GcpOfferID            string                                                            `json:"gcp_offer_id"`
	NetsuiteResellerID    string                                                            `json:"netsuite_reseller_id"`
	ResellerContractValue float64                                                           `json:"reseller_contract_value"`
	StartingAt            time.Time                                                         `json:"starting_at" format:"date-time"`
	JSON                  v1ContractListResponseDataAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// v1ContractListResponseDataAmendmentsResellerRoyaltyJSON contains the JSON
// metadata for the struct [V1ContractListResponseDataAmendmentsResellerRoyalty]
type v1ContractListResponseDataAmendmentsResellerRoyaltyJSON struct {
	ResellerType          apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	Fraction              apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerContractValue apijson.Field
	StartingAt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *V1ContractListResponseDataAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType string

const (
	V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAws           V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS"
	V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp           V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP"
	V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAws, V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeAwsProService, V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcp, V1ContractListResponseDataAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// The billing provider configuration associated with a contract.
type V1ContractListResponseDataCustomerBillingProviderConfiguration struct {
	BillingProvider V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider `json:"billing_provider,required"`
	DeliveryMethod  V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod  `json:"delivery_method,required"`
	ID              string                                                                        `json:"id" format:"uuid"`
	// Configuration for the billing provider. The structure of this object is specific
	// to the billing provider.
	Configuration map[string]interface{}                                             `json:"configuration"`
	JSON          v1ContractListResponseDataCustomerBillingProviderConfigurationJSON `json:"-"`
}

// v1ContractListResponseDataCustomerBillingProviderConfigurationJSON contains the
// JSON metadata for the struct
// [V1ContractListResponseDataCustomerBillingProviderConfiguration]
type v1ContractListResponseDataCustomerBillingProviderConfigurationJSON struct {
	BillingProvider apijson.Field
	DeliveryMethod  apijson.Field
	ID              apijson.Field
	Configuration   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1ContractListResponseDataCustomerBillingProviderConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataCustomerBillingProviderConfigurationJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider string

const (
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace   V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "aws_marketplace"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe           V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "stripe"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite         V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "netsuite"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom           V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "custom"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "azure_marketplace"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "quickbooks_online"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday          V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "workday"
	V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace   V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider = "gcp_marketplace"
)

func (r V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAwsMarketplace, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderStripe, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderNetsuite, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderCustom, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderAzureMarketplace, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderQuickbooksOnline, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderWorkday, V1ContractListResponseDataCustomerBillingProviderConfigurationBillingProviderGcpMarketplace:
		return true
	}
	return false
}

type V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod string

const (
	V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs                  V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle                  V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "tackle"
	V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns                  V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSqs, V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodTackle, V1ContractListResponseDataCustomerBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1ContractListResponseDataScheduledChargesOnUsageInvoices string

const (
	V1ContractListResponseDataScheduledChargesOnUsageInvoicesAll V1ContractListResponseDataScheduledChargesOnUsageInvoices = "ALL"
)

func (r V1ContractListResponseDataScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type V1ContractListResponseDataSubscription struct {
	CollectionSchedule V1ContractListResponseDataSubscriptionsCollectionSchedule `json:"collection_schedule,required"`
	Proration          V1ContractListResponseDataSubscriptionsProration          `json:"proration,required"`
	QuantitySchedule   []V1ContractListResponseDataSubscriptionsQuantitySchedule `json:"quantity_schedule,required"`
	StartingAt         time.Time                                                 `json:"starting_at,required" format:"date-time"`
	SubscriptionRate   V1ContractListResponseDataSubscriptionsSubscriptionRate   `json:"subscription_rate,required"`
	Description        string                                                    `json:"description"`
	EndingBefore       time.Time                                                 `json:"ending_before" format:"date-time"`
	Name               string                                                    `json:"name"`
	JSON               v1ContractListResponseDataSubscriptionJSON                `json:"-"`
}

// v1ContractListResponseDataSubscriptionJSON contains the JSON metadata for the
// struct [V1ContractListResponseDataSubscription]
type v1ContractListResponseDataSubscriptionJSON struct {
	CollectionSchedule apijson.Field
	Proration          apijson.Field
	QuantitySchedule   apijson.Field
	StartingAt         apijson.Field
	SubscriptionRate   apijson.Field
	Description        apijson.Field
	EndingBefore       apijson.Field
	Name               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *V1ContractListResponseDataSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataSubscriptionJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataSubscriptionsCollectionSchedule string

const (
	V1ContractListResponseDataSubscriptionsCollectionScheduleAdvance V1ContractListResponseDataSubscriptionsCollectionSchedule = "ADVANCE"
	V1ContractListResponseDataSubscriptionsCollectionScheduleArrears V1ContractListResponseDataSubscriptionsCollectionSchedule = "ARREARS"
)

func (r V1ContractListResponseDataSubscriptionsCollectionSchedule) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataSubscriptionsCollectionScheduleAdvance, V1ContractListResponseDataSubscriptionsCollectionScheduleArrears:
		return true
	}
	return false
}

type V1ContractListResponseDataSubscriptionsProration struct {
	InvoiceBehavior V1ContractListResponseDataSubscriptionsProrationInvoiceBehavior `json:"invoice_behavior,required"`
	IsProrated      bool                                                            `json:"is_prorated,required"`
	JSON            v1ContractListResponseDataSubscriptionsProrationJSON            `json:"-"`
}

// v1ContractListResponseDataSubscriptionsProrationJSON contains the JSON metadata
// for the struct [V1ContractListResponseDataSubscriptionsProration]
type v1ContractListResponseDataSubscriptionsProrationJSON struct {
	InvoiceBehavior apijson.Field
	IsProrated      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *V1ContractListResponseDataSubscriptionsProration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataSubscriptionsProrationJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataSubscriptionsProrationInvoiceBehavior string

const (
	V1ContractListResponseDataSubscriptionsProrationInvoiceBehaviorBillImmediately          V1ContractListResponseDataSubscriptionsProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	V1ContractListResponseDataSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate V1ContractListResponseDataSubscriptionsProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r V1ContractListResponseDataSubscriptionsProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataSubscriptionsProrationInvoiceBehaviorBillImmediately, V1ContractListResponseDataSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type V1ContractListResponseDataSubscriptionsQuantitySchedule struct {
	Quantity     float64                                                     `json:"quantity,required"`
	StartingAt   time.Time                                                   `json:"starting_at,required" format:"date-time"`
	EndingBefore time.Time                                                   `json:"ending_before" format:"date-time"`
	JSON         v1ContractListResponseDataSubscriptionsQuantityScheduleJSON `json:"-"`
}

// v1ContractListResponseDataSubscriptionsQuantityScheduleJSON contains the JSON
// metadata for the struct
// [V1ContractListResponseDataSubscriptionsQuantitySchedule]
type v1ContractListResponseDataSubscriptionsQuantityScheduleJSON struct {
	Quantity     apijson.Field
	StartingAt   apijson.Field
	EndingBefore apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V1ContractListResponseDataSubscriptionsQuantitySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataSubscriptionsQuantityScheduleJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataSubscriptionsSubscriptionRate struct {
	BillingFrequency V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency `json:"billing_frequency,required"`
	Product          V1ContractListResponseDataSubscriptionsSubscriptionRateProduct          `json:"product,required"`
	JSON             v1ContractListResponseDataSubscriptionsSubscriptionRateJSON             `json:"-"`
}

// v1ContractListResponseDataSubscriptionsSubscriptionRateJSON contains the JSON
// metadata for the struct
// [V1ContractListResponseDataSubscriptionsSubscriptionRate]
type v1ContractListResponseDataSubscriptionsSubscriptionRateJSON struct {
	BillingFrequency apijson.Field
	Product          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V1ContractListResponseDataSubscriptionsSubscriptionRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataSubscriptionsSubscriptionRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency string

const (
	V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyMonthly   V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency = "MONTHLY"
	V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyQuarterly V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency = "QUARTERLY"
	V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyAnnual    V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency = "ANNUAL"
)

func (r V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyMonthly, V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyQuarterly, V1ContractListResponseDataSubscriptionsSubscriptionRateBillingFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractListResponseDataSubscriptionsSubscriptionRateProduct struct {
	ID   string                                                             `json:"id,required" format:"uuid"`
	Name string                                                             `json:"name,required"`
	JSON v1ContractListResponseDataSubscriptionsSubscriptionRateProductJSON `json:"-"`
}

// v1ContractListResponseDataSubscriptionsSubscriptionRateProductJSON contains the
// JSON metadata for the struct
// [V1ContractListResponseDataSubscriptionsSubscriptionRateProduct]
type v1ContractListResponseDataSubscriptionsSubscriptionRateProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractListResponseDataSubscriptionsSubscriptionRateProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListResponseDataSubscriptionsSubscriptionRateProductJSON) RawJSON() string {
	return r.raw
}

type V1ContractAmendResponse struct {
	Data shared.ID                   `json:"data,required"`
	JSON v1ContractAmendResponseJSON `json:"-"`
}

// v1ContractAmendResponseJSON contains the JSON metadata for the struct
// [V1ContractAmendResponse]
type v1ContractAmendResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractAmendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractAmendResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractArchiveResponse struct {
	Data shared.ID                     `json:"data,required"`
	JSON v1ContractArchiveResponseJSON `json:"-"`
}

// v1ContractArchiveResponseJSON contains the JSON metadata for the struct
// [V1ContractArchiveResponse]
type v1ContractArchiveResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractArchiveResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractNewHistoricalInvoicesResponse struct {
	Data []Invoice                                   `json:"data,required"`
	JSON v1ContractNewHistoricalInvoicesResponseJSON `json:"-"`
}

// v1ContractNewHistoricalInvoicesResponseJSON contains the JSON metadata for the
// struct [V1ContractNewHistoricalInvoicesResponse]
type v1ContractNewHistoricalInvoicesResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractNewHistoricalInvoicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractNewHistoricalInvoicesResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractListBalancesResponse struct {
	Data     []V1ContractListBalancesResponseData `json:"data,required"`
	NextPage string                               `json:"next_page,required,nullable"`
	JSON     v1ContractListBalancesResponseJSON   `json:"-"`
}

// v1ContractListBalancesResponseJSON contains the JSON metadata for the struct
// [V1ContractListBalancesResponse]
type v1ContractListBalancesResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractListBalancesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractListBalancesResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractListBalancesResponseData struct {
	ID string `json:"id,required" format:"uuid"`
	// This field can have the runtime type of [shared.CommitProduct],
	// [shared.CreditProduct].
	Product interface{}                            `json:"product,required"`
	Type    V1ContractListBalancesResponseDataType `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule shared.ScheduleDuration `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount float64 `json:"amount"`
	// This field can have the runtime type of [[]string].
	ApplicableContractIDs interface{} `json:"applicable_contract_ids"`
	// This field can have the runtime type of [[]string].
	ApplicableProductIDs interface{} `json:"applicable_product_ids"`
	// This field can have the runtime type of [[]string].
	ApplicableProductTags interface{} `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance float64 `json:"balance"`
	// This field can have the runtime type of [shared.CommitContract],
	// [shared.CreditContract].
	Contract interface{} `json:"contract"`
	// This field can have the runtime type of [map[string]string].
	CustomFields interface{} `json:"custom_fields"`
	Description  string      `json:"description"`
	// This field can have the runtime type of [shared.CommitInvoiceContract].
	InvoiceContract interface{} `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule shared.SchedulePointInTime `json:"invoice_schedule"`
	// This field can have the runtime type of [[]shared.CommitLedger],
	// [[]shared.CreditLedger].
	Ledger interface{} `json:"ledger"`
	Name   string      `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64                                    `json:"priority"`
	RateType V1ContractListBalancesResponseDataRateType `json:"rate_type"`
	// This field can have the runtime type of [shared.CommitRolledOverFrom].
	RolledOverFrom   interface{} `json:"rolled_over_from"`
	RolloverFraction float64     `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string                                 `json:"uniqueness_key"`
	JSON          v1ContractListBalancesResponseDataJSON `json:"-"`
	union         V1ContractListBalancesResponseDataUnion
}

// v1ContractListBalancesResponseDataJSON contains the JSON metadata for the struct
// [V1ContractListBalancesResponseData]
type v1ContractListBalancesResponseDataJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	Amount                  apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	InvoiceContract         apijson.Field
	InvoiceSchedule         apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	RolledOverFrom          apijson.Field
	RolloverFraction        apijson.Field
	SalesforceOpportunityID apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r v1ContractListBalancesResponseDataJSON) RawJSON() string {
	return r.raw
}

func (r *V1ContractListBalancesResponseData) UnmarshalJSON(data []byte) (err error) {
	*r = V1ContractListBalancesResponseData{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [V1ContractListBalancesResponseDataUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.Commit], [shared.Credit].
func (r V1ContractListBalancesResponseData) AsUnion() V1ContractListBalancesResponseDataUnion {
	return r.union
}

// Union satisfied by [shared.Commit] or [shared.Credit].
type V1ContractListBalancesResponseDataUnion interface {
	ImplementsV1ContractListBalancesResponseData()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1ContractListBalancesResponseDataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Commit{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.Credit{}),
		},
	)
}

type V1ContractListBalancesResponseDataType string

const (
	V1ContractListBalancesResponseDataTypePrepaid  V1ContractListBalancesResponseDataType = "PREPAID"
	V1ContractListBalancesResponseDataTypePostpaid V1ContractListBalancesResponseDataType = "POSTPAID"
	V1ContractListBalancesResponseDataTypeCredit   V1ContractListBalancesResponseDataType = "CREDIT"
)

func (r V1ContractListBalancesResponseDataType) IsKnown() bool {
	switch r {
	case V1ContractListBalancesResponseDataTypePrepaid, V1ContractListBalancesResponseDataTypePostpaid, V1ContractListBalancesResponseDataTypeCredit:
		return true
	}
	return false
}

type V1ContractListBalancesResponseDataRateType string

const (
	V1ContractListBalancesResponseDataRateTypeCommitRate V1ContractListBalancesResponseDataRateType = "COMMIT_RATE"
	V1ContractListBalancesResponseDataRateTypeListRate   V1ContractListBalancesResponseDataRateType = "LIST_RATE"
)

func (r V1ContractListBalancesResponseDataRateType) IsKnown() bool {
	switch r {
	case V1ContractListBalancesResponseDataRateTypeCommitRate, V1ContractListBalancesResponseDataRateTypeListRate:
		return true
	}
	return false
}

type V1ContractGetRateScheduleResponse struct {
	Data     []V1ContractGetRateScheduleResponseData `json:"data,required"`
	NextPage string                                  `json:"next_page,nullable"`
	JSON     v1ContractGetRateScheduleResponseJSON   `json:"-"`
}

// v1ContractGetRateScheduleResponseJSON contains the JSON metadata for the struct
// [V1ContractGetRateScheduleResponse]
type v1ContractGetRateScheduleResponseJSON struct {
	Data        apijson.Field
	NextPage    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractGetRateScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetRateScheduleResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetRateScheduleResponseData struct {
	Entitled            bool                                                  `json:"entitled,required"`
	ListRate            shared.Rate                                           `json:"list_rate,required"`
	ProductCustomFields map[string]string                                     `json:"product_custom_fields,required"`
	ProductID           string                                                `json:"product_id,required" format:"uuid"`
	ProductName         string                                                `json:"product_name,required"`
	ProductTags         []string                                              `json:"product_tags,required"`
	RateCardID          string                                                `json:"rate_card_id,required" format:"uuid"`
	StartingAt          time.Time                                             `json:"starting_at,required" format:"date-time"`
	BillingFrequency    V1ContractGetRateScheduleResponseDataBillingFrequency `json:"billing_frequency"`
	// A distinct rate on the rate card. You can choose to use this rate rather than
	// list rate when consuming a credit or commit.
	CommitRate         V1ContractGetRateScheduleResponseDataCommitRate `json:"commit_rate"`
	EndingBefore       time.Time                                       `json:"ending_before" format:"date-time"`
	OverrideRate       shared.Rate                                     `json:"override_rate"`
	PricingGroupValues map[string]string                               `json:"pricing_group_values"`
	JSON               v1ContractGetRateScheduleResponseDataJSON       `json:"-"`
}

// v1ContractGetRateScheduleResponseDataJSON contains the JSON metadata for the
// struct [V1ContractGetRateScheduleResponseData]
type v1ContractGetRateScheduleResponseDataJSON struct {
	Entitled            apijson.Field
	ListRate            apijson.Field
	ProductCustomFields apijson.Field
	ProductID           apijson.Field
	ProductName         apijson.Field
	ProductTags         apijson.Field
	RateCardID          apijson.Field
	StartingAt          apijson.Field
	BillingFrequency    apijson.Field
	CommitRate          apijson.Field
	EndingBefore        apijson.Field
	OverrideRate        apijson.Field
	PricingGroupValues  apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *V1ContractGetRateScheduleResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetRateScheduleResponseDataJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetRateScheduleResponseDataBillingFrequency string

const (
	V1ContractGetRateScheduleResponseDataBillingFrequencyMonthly   V1ContractGetRateScheduleResponseDataBillingFrequency = "MONTHLY"
	V1ContractGetRateScheduleResponseDataBillingFrequencyQuarterly V1ContractGetRateScheduleResponseDataBillingFrequency = "QUARTERLY"
	V1ContractGetRateScheduleResponseDataBillingFrequencyAnnual    V1ContractGetRateScheduleResponseDataBillingFrequency = "ANNUAL"
)

func (r V1ContractGetRateScheduleResponseDataBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractGetRateScheduleResponseDataBillingFrequencyMonthly, V1ContractGetRateScheduleResponseDataBillingFrequencyQuarterly, V1ContractGetRateScheduleResponseDataBillingFrequencyAnnual:
		return true
	}
	return false
}

// A distinct rate on the rate card. You can choose to use this rate rather than
// list rate when consuming a credit or commit.
type V1ContractGetRateScheduleResponseDataCommitRate struct {
	RateType V1ContractGetRateScheduleResponseDataCommitRateRateType `json:"rate_type,required"`
	// Commit rate price. For FLAT rate_type, this must be >=0.
	Price float64 `json:"price"`
	// Only set for TIERED rate_type.
	Tiers []shared.Tier                                       `json:"tiers"`
	JSON  v1ContractGetRateScheduleResponseDataCommitRateJSON `json:"-"`
}

// v1ContractGetRateScheduleResponseDataCommitRateJSON contains the JSON metadata
// for the struct [V1ContractGetRateScheduleResponseDataCommitRate]
type v1ContractGetRateScheduleResponseDataCommitRateJSON struct {
	RateType    apijson.Field
	Price       apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractGetRateScheduleResponseDataCommitRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractGetRateScheduleResponseDataCommitRateJSON) RawJSON() string {
	return r.raw
}

type V1ContractGetRateScheduleResponseDataCommitRateRateType string

const (
	V1ContractGetRateScheduleResponseDataCommitRateRateTypeFlat         V1ContractGetRateScheduleResponseDataCommitRateRateType = "FLAT"
	V1ContractGetRateScheduleResponseDataCommitRateRateTypePercentage   V1ContractGetRateScheduleResponseDataCommitRateRateType = "PERCENTAGE"
	V1ContractGetRateScheduleResponseDataCommitRateRateTypeSubscription V1ContractGetRateScheduleResponseDataCommitRateRateType = "SUBSCRIPTION"
	V1ContractGetRateScheduleResponseDataCommitRateRateTypeTiered       V1ContractGetRateScheduleResponseDataCommitRateRateType = "TIERED"
	V1ContractGetRateScheduleResponseDataCommitRateRateTypeCustom       V1ContractGetRateScheduleResponseDataCommitRateRateType = "CUSTOM"
)

func (r V1ContractGetRateScheduleResponseDataCommitRateRateType) IsKnown() bool {
	switch r {
	case V1ContractGetRateScheduleResponseDataCommitRateRateTypeFlat, V1ContractGetRateScheduleResponseDataCommitRateRateTypePercentage, V1ContractGetRateScheduleResponseDataCommitRateRateTypeSubscription, V1ContractGetRateScheduleResponseDataCommitRateRateTypeTiered, V1ContractGetRateScheduleResponseDataCommitRateRateTypeCustom:
		return true
	}
	return false
}

type V1ContractScheduleProServicesInvoiceResponse struct {
	Data []Invoice                                        `json:"data,required"`
	JSON v1ContractScheduleProServicesInvoiceResponseJSON `json:"-"`
}

// v1ContractScheduleProServicesInvoiceResponseJSON contains the JSON metadata for
// the struct [V1ContractScheduleProServicesInvoiceResponse]
type v1ContractScheduleProServicesInvoiceResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractScheduleProServicesInvoiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractScheduleProServicesInvoiceResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractUpdateEndDateResponse struct {
	Data shared.ID                           `json:"data,required"`
	JSON v1ContractUpdateEndDateResponseJSON `json:"-"`
}

// v1ContractUpdateEndDateResponseJSON contains the JSON metadata for the struct
// [V1ContractUpdateEndDateResponse]
type v1ContractUpdateEndDateResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ContractUpdateEndDateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ContractUpdateEndDateResponseJSON) RawJSON() string {
	return r.raw
}

type V1ContractNewParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive contract start time
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// The billing provider configuration associated with a contract.
	BillingProviderConfiguration param.Field[V1ContractNewParamsBillingProviderConfiguration] `json:"billing_provider_configuration"`
	Commits                      param.Field[[]V1ContractNewParamsCommit]                     `json:"commits"`
	Credits                      param.Field[[]V1ContractNewParamsCredit]                     `json:"credits"`
	CustomFields                 param.Field[map[string]string]                               `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	Discounts param.Field[[]V1ContractNewParamsDiscount] `json:"discounts"`
	// exclusive contract end time
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
	// prices automatically. EXPLICIT prioritization requires specifying priorities for
	// each multiplier; the one with the lowest priority value will be prioritized
	// first. If tiered overrides are used, prioritization must be explicit.
	MultiplierOverridePrioritization param.Field[V1ContractNewParamsMultiplierOverridePrioritization] `json:"multiplier_override_prioritization"`
	Name                             param.Field[string]                                              `json:"name"`
	NetPaymentTermsDays              param.Field[float64]                                             `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                        `json:"netsuite_sales_order_id"`
	Overrides            param.Field[[]V1ContractNewParamsOverride] `json:"overrides"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices param.Field[[]V1ContractNewParamsProfessionalService] `json:"professional_services"`
	// Selects the rate card linked to the specified alias as of the contract's start
	// date.
	RateCardAlias    param.Field[string]                               `json:"rate_card_alias"`
	RateCardID       param.Field[string]                               `json:"rate_card_id" format:"uuid"`
	RecurringCommits param.Field[[]V1ContractNewParamsRecurringCommit] `json:"recurring_commits"`
	RecurringCredits param.Field[[]V1ContractNewParamsRecurringCredit] `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties param.Field[[]V1ContractNewParamsResellerRoyalty] `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string]                               `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]V1ContractNewParamsScheduledCharge] `json:"scheduled_charges"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices param.Field[V1ContractNewParamsScheduledChargesOnUsageInvoices] `json:"scheduled_charges_on_usage_invoices"`
	Subscriptions                   param.Field[[]V1ContractNewParamsSubscription]                  `json:"subscriptions"`
	ThresholdBillingConfiguration   param.Field[V1ContractNewParamsThresholdBillingConfiguration]   `json:"threshold_billing_configuration"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Field[float64]                       `json:"total_contract_value"`
	Transition         param.Field[V1ContractNewParamsTransition] `json:"transition"`
	// Prevents the creation of duplicates. If a request to create a record is made
	// with a previously used uniqueness key, a new record will not be created and the
	// request will fail with a 409 error.
	UniquenessKey          param.Field[string]                                    `json:"uniqueness_key"`
	UsageFilter            param.Field[shared.BaseUsageFilterParam]               `json:"usage_filter"`
	UsageStatementSchedule param.Field[V1ContractNewParamsUsageStatementSchedule] `json:"usage_statement_schedule"`
}

func (r V1ContractNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The billing provider configuration associated with a contract.
type V1ContractNewParamsBillingProviderConfiguration struct {
	BillingProvider param.Field[V1ContractNewParamsBillingProviderConfigurationBillingProvider] `json:"billing_provider"`
	// The Metronome ID of the billing provider configuration
	BillingProviderConfigurationID param.Field[string]                                                        `json:"billing_provider_configuration_id" format:"uuid"`
	DeliveryMethod                 param.Field[V1ContractNewParamsBillingProviderConfigurationDeliveryMethod] `json:"delivery_method"`
}

func (r V1ContractNewParamsBillingProviderConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsBillingProviderConfigurationBillingProvider string

const (
	V1ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace   V1ContractNewParamsBillingProviderConfigurationBillingProvider = "aws_marketplace"
	V1ContractNewParamsBillingProviderConfigurationBillingProviderAzureMarketplace V1ContractNewParamsBillingProviderConfigurationBillingProvider = "azure_marketplace"
	V1ContractNewParamsBillingProviderConfigurationBillingProviderGcpMarketplace   V1ContractNewParamsBillingProviderConfigurationBillingProvider = "gcp_marketplace"
	V1ContractNewParamsBillingProviderConfigurationBillingProviderStripe           V1ContractNewParamsBillingProviderConfigurationBillingProvider = "stripe"
	V1ContractNewParamsBillingProviderConfigurationBillingProviderNetsuite         V1ContractNewParamsBillingProviderConfigurationBillingProvider = "netsuite"
)

func (r V1ContractNewParamsBillingProviderConfigurationBillingProvider) IsKnown() bool {
	switch r {
	case V1ContractNewParamsBillingProviderConfigurationBillingProviderAwsMarketplace, V1ContractNewParamsBillingProviderConfigurationBillingProviderAzureMarketplace, V1ContractNewParamsBillingProviderConfigurationBillingProviderGcpMarketplace, V1ContractNewParamsBillingProviderConfigurationBillingProviderStripe, V1ContractNewParamsBillingProviderConfigurationBillingProviderNetsuite:
		return true
	}
	return false
}

type V1ContractNewParamsBillingProviderConfigurationDeliveryMethod string

const (
	V1ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider V1ContractNewParamsBillingProviderConfigurationDeliveryMethod = "direct_to_billing_provider"
	V1ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSqs                  V1ContractNewParamsBillingProviderConfigurationDeliveryMethod = "aws_sqs"
	V1ContractNewParamsBillingProviderConfigurationDeliveryMethodTackle                  V1ContractNewParamsBillingProviderConfigurationDeliveryMethod = "tackle"
	V1ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSns                  V1ContractNewParamsBillingProviderConfigurationDeliveryMethod = "aws_sns"
)

func (r V1ContractNewParamsBillingProviderConfigurationDeliveryMethod) IsKnown() bool {
	switch r {
	case V1ContractNewParamsBillingProviderConfigurationDeliveryMethodDirectToBillingProvider, V1ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSqs, V1ContractNewParamsBillingProviderConfigurationDeliveryMethodTackle, V1ContractNewParamsBillingProviderConfigurationDeliveryMethodAwsSns:
		return true
	}
	return false
}

type V1ContractNewParamsCommit struct {
	ProductID param.Field[string]                         `json:"product_id,required" format:"uuid"`
	Type      param.Field[V1ContractNewParamsCommitsType] `json:"type,required"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule param.Field[V1ContractNewParamsCommitsAccessSchedule] `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Field[float64] `json:"amount"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[V1ContractNewParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                            `json:"priority"`
	RateType param.Field[V1ContractNewParamsCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V1ContractNewParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCommitsType string

const (
	V1ContractNewParamsCommitsTypePrepaid  V1ContractNewParamsCommitsType = "PREPAID"
	V1ContractNewParamsCommitsTypePostpaid V1ContractNewParamsCommitsType = "POSTPAID"
)

func (r V1ContractNewParamsCommitsType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsCommitsTypePrepaid, V1ContractNewParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type V1ContractNewParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]V1ContractNewParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1ContractNewParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1ContractNewParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V1ContractNewParamsCommitsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractNewParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractNewParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractNewParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractNewParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                           `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, V1ContractNewParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCommitsRateType string

const (
	V1ContractNewParamsCommitsRateTypeCommitRate V1ContractNewParamsCommitsRateType = "COMMIT_RATE"
	V1ContractNewParamsCommitsRateTypeListRate   V1ContractNewParamsCommitsRateType = "LIST_RATE"
)

func (r V1ContractNewParamsCommitsRateType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsCommitsRateTypeCommitRate, V1ContractNewParamsCommitsRateTypeListRate:
		return true
	}
	return false
}

type V1ContractNewParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[V1ContractNewParamsCreditsAccessSchedule] `json:"access_schedule,required"`
	ProductID      param.Field[string]                                   `json:"product_id,required" format:"uuid"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                            `json:"priority"`
	RateType param.Field[V1ContractNewParamsCreditsRateType] `json:"rate_type"`
}

func (r V1ContractNewParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type V1ContractNewParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]V1ContractNewParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1ContractNewParamsCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCreditsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1ContractNewParamsCreditsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsCreditsRateType string

const (
	V1ContractNewParamsCreditsRateTypeCommitRate V1ContractNewParamsCreditsRateType = "COMMIT_RATE"
	V1ContractNewParamsCreditsRateTypeListRate   V1ContractNewParamsCreditsRateType = "LIST_RATE"
)

func (r V1ContractNewParamsCreditsRateType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsCreditsRateTypeCommitRate, V1ContractNewParamsCreditsRateTypeListRate:
		return true
	}
	return false
}

type V1ContractNewParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule     param.Field[V1ContractNewParamsDiscountsSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                    `json:"custom_fields"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractNewParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractNewParamsDiscountsSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractNewParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractNewParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractNewParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractNewParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                      `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractNewParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, V1ContractNewParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defaults to LOWEST_MULTIPLIER, which applies the greatest discount to list
// prices automatically. EXPLICIT prioritization requires specifying priorities for
// each multiplier; the one with the lowest priority value will be prioritized
// first. If tiered overrides are used, prioritization must be explicit.
type V1ContractNewParamsMultiplierOverridePrioritization string

const (
	V1ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier V1ContractNewParamsMultiplierOverridePrioritization = "LOWEST_MULTIPLIER"
	V1ContractNewParamsMultiplierOverridePrioritizationExplicit         V1ContractNewParamsMultiplierOverridePrioritization = "EXPLICIT"
)

func (r V1ContractNewParamsMultiplierOverridePrioritization) IsKnown() bool {
	switch r {
	case V1ContractNewParamsMultiplierOverridePrioritizationLowestMultiplier, V1ContractNewParamsMultiplierOverridePrioritizationExplicit:
		return true
	}
	return false
}

type V1ContractNewParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Field[bool] `json:"is_commit_specific"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers param.Field[[]V1ContractNewParamsOverridesOverrideSpecifier] `json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[V1ContractNewParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Field[float64] `json:"priority"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	Target param.Field[V1ContractNewParamsOverridesTarget] `json:"target"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]V1ContractNewParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[V1ContractNewParamsOverridesType] `json:"type"`
}

func (r V1ContractNewParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsOverridesOverrideSpecifier struct {
	BillingFrequency param.Field[V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency] `json:"billing_frequency"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs param.Field[[]string] `json:"commit_ids"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values.
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs param.Field[[]string] `json:"recurring_commit_ids"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs param.Field[[]string] `json:"recurring_credit_ids"`
}

func (r V1ContractNewParamsOverridesOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency string

const (
	V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyMonthly   V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyQuarterly V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyAnnual    V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
)

func (r V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyMonthly, V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyQuarterly, V1ContractNewParamsOverridesOverrideSpecifiersBillingFrequencyAnnual:
		return true
	}
	return false
}

// Required for OVERWRITE type.
type V1ContractNewParamsOverridesOverwriteRate struct {
	RateType     param.Field[V1ContractNewParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	CreditTypeID param.Field[string]                                            `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r V1ContractNewParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsOverridesOverwriteRateRateType string

const (
	V1ContractNewParamsOverridesOverwriteRateRateTypeFlat         V1ContractNewParamsOverridesOverwriteRateRateType = "FLAT"
	V1ContractNewParamsOverridesOverwriteRateRateTypePercentage   V1ContractNewParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	V1ContractNewParamsOverridesOverwriteRateRateTypeSubscription V1ContractNewParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V1ContractNewParamsOverridesOverwriteRateRateTypeTiered       V1ContractNewParamsOverridesOverwriteRateRateType = "TIERED"
	V1ContractNewParamsOverridesOverwriteRateRateTypeCustom       V1ContractNewParamsOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V1ContractNewParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsOverridesOverwriteRateRateTypeFlat, V1ContractNewParamsOverridesOverwriteRateRateTypePercentage, V1ContractNewParamsOverridesOverwriteRateRateTypeSubscription, V1ContractNewParamsOverridesOverwriteRateRateTypeTiered, V1ContractNewParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

// Indicates whether the override applies to commit rates or list rates. Can only
// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
// `"LIST_RATE"`.
type V1ContractNewParamsOverridesTarget string

const (
	V1ContractNewParamsOverridesTargetCommitRate V1ContractNewParamsOverridesTarget = "COMMIT_RATE"
	V1ContractNewParamsOverridesTargetListRate   V1ContractNewParamsOverridesTarget = "LIST_RATE"
)

func (r V1ContractNewParamsOverridesTarget) IsKnown() bool {
	switch r {
	case V1ContractNewParamsOverridesTargetCommitRate, V1ContractNewParamsOverridesTargetListRate:
		return true
	}
	return false
}

type V1ContractNewParamsOverridesTier struct {
	Multiplier param.Field[float64] `json:"multiplier,required"`
	Size       param.Field[float64] `json:"size"`
}

func (r V1ContractNewParamsOverridesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Overwrites are prioritized over multipliers and tiered overrides.
type V1ContractNewParamsOverridesType string

const (
	V1ContractNewParamsOverridesTypeOverwrite  V1ContractNewParamsOverridesType = "OVERWRITE"
	V1ContractNewParamsOverridesTypeMultiplier V1ContractNewParamsOverridesType = "MULTIPLIER"
	V1ContractNewParamsOverridesTypeTiered     V1ContractNewParamsOverridesType = "TIERED"
)

func (r V1ContractNewParamsOverridesType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsOverridesTypeOverwrite, V1ContractNewParamsOverridesTypeMultiplier, V1ContractNewParamsOverridesTypeTiered:
		return true
	}
	return false
}

type V1ContractNewParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    param.Field[float64]           `json:"unit_price,required"`
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	Description  param.Field[string]            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractNewParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsRecurringCommit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[V1ContractNewParamsRecurringCommitsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[V1ContractNewParamsRecurringCommitsCommitDuration] `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Will be passed down to the individual commits
	Description param.Field[string] `json:"description"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount param.Field[V1ContractNewParamsRecurringCommitsInvoiceAmount] `json:"invoice_amount"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration param.Field[V1ContractNewParamsRecurringCommitsProration] `json:"proration"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[V1ContractNewParamsRecurringCommitsRateType] `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V1ContractNewParamsRecurringCommitsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V1ContractNewParamsRecurringCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type V1ContractNewParamsRecurringCommitsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V1ContractNewParamsRecurringCommitsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type V1ContractNewParamsRecurringCommitsCommitDuration struct {
	Unit  param.Field[V1ContractNewParamsRecurringCommitsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                               `json:"value,required"`
}

func (r V1ContractNewParamsRecurringCommitsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsRecurringCommitsCommitDurationUnit string

const (
	V1ContractNewParamsRecurringCommitsCommitDurationUnitPeriods V1ContractNewParamsRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r V1ContractNewParamsRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// The amount the customer should be billed for the commit. Not required.
type V1ContractNewParamsRecurringCommitsInvoiceAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V1ContractNewParamsRecurringCommitsInvoiceAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V1ContractNewParamsRecurringCommitsProration string

const (
	V1ContractNewParamsRecurringCommitsProrationNone         V1ContractNewParamsRecurringCommitsProration = "NONE"
	V1ContractNewParamsRecurringCommitsProrationFirst        V1ContractNewParamsRecurringCommitsProration = "FIRST"
	V1ContractNewParamsRecurringCommitsProrationLast         V1ContractNewParamsRecurringCommitsProration = "LAST"
	V1ContractNewParamsRecurringCommitsProrationFirstAndLast V1ContractNewParamsRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r V1ContractNewParamsRecurringCommitsProration) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCommitsProrationNone, V1ContractNewParamsRecurringCommitsProrationFirst, V1ContractNewParamsRecurringCommitsProrationLast, V1ContractNewParamsRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// Whether the created commits will use the commit rate or list rate
type V1ContractNewParamsRecurringCommitsRateType string

const (
	V1ContractNewParamsRecurringCommitsRateTypeCommitRate V1ContractNewParamsRecurringCommitsRateType = "COMMIT_RATE"
	V1ContractNewParamsRecurringCommitsRateTypeListRate   V1ContractNewParamsRecurringCommitsRateType = "LIST_RATE"
)

func (r V1ContractNewParamsRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCommitsRateTypeCommitRate, V1ContractNewParamsRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V1ContractNewParamsRecurringCommitsRecurrenceFrequency string

const (
	V1ContractNewParamsRecurringCommitsRecurrenceFrequencyMonthly   V1ContractNewParamsRecurringCommitsRecurrenceFrequency = "MONTHLY"
	V1ContractNewParamsRecurringCommitsRecurrenceFrequencyQuarterly V1ContractNewParamsRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	V1ContractNewParamsRecurringCommitsRecurrenceFrequencyAnnual    V1ContractNewParamsRecurringCommitsRecurrenceFrequency = "ANNUAL"
)

func (r V1ContractNewParamsRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCommitsRecurrenceFrequencyMonthly, V1ContractNewParamsRecurringCommitsRecurrenceFrequencyQuarterly, V1ContractNewParamsRecurringCommitsRecurrenceFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsRecurringCredit struct {
	// The amount of commit to grant.
	AccessAmount param.Field[V1ContractNewParamsRecurringCreditsAccessAmount] `json:"access_amount,required"`
	// The amount of time the created commits will be valid for.
	CommitDuration param.Field[V1ContractNewParamsRecurringCreditsCommitDuration] `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority  param.Field[float64] `json:"priority,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// determines the start time for the first commit
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// Will be passed down to the individual commits
	Description param.Field[string] `json:"description"`
	// Determines when the contract will stop creating recurring commits. optional
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	// displayed on invoices. will be passed through to the individual commits
	Name param.Field[string] `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration param.Field[V1ContractNewParamsRecurringCreditsProration] `json:"proration"`
	// Whether the created commits will use the commit rate or list rate
	RateType param.Field[V1ContractNewParamsRecurringCreditsRateType] `json:"rate_type"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency param.Field[V1ContractNewParamsRecurringCreditsRecurrenceFrequency] `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID that can be used to reference the recurring commit for commit
	// specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V1ContractNewParamsRecurringCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of commit to grant.
type V1ContractNewParamsRecurringCreditsAccessAmount struct {
	CreditTypeID param.Field[string]  `json:"credit_type_id,required" format:"uuid"`
	Quantity     param.Field[float64] `json:"quantity,required"`
	UnitPrice    param.Field[float64] `json:"unit_price,required"`
}

func (r V1ContractNewParamsRecurringCreditsAccessAmount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The amount of time the created commits will be valid for.
type V1ContractNewParamsRecurringCreditsCommitDuration struct {
	Unit  param.Field[V1ContractNewParamsRecurringCreditsCommitDurationUnit] `json:"unit,required"`
	Value param.Field[float64]                                               `json:"value,required"`
}

func (r V1ContractNewParamsRecurringCreditsCommitDuration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsRecurringCreditsCommitDurationUnit string

const (
	V1ContractNewParamsRecurringCreditsCommitDurationUnitPeriods V1ContractNewParamsRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r V1ContractNewParamsRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type V1ContractNewParamsRecurringCreditsProration string

const (
	V1ContractNewParamsRecurringCreditsProrationNone         V1ContractNewParamsRecurringCreditsProration = "NONE"
	V1ContractNewParamsRecurringCreditsProrationFirst        V1ContractNewParamsRecurringCreditsProration = "FIRST"
	V1ContractNewParamsRecurringCreditsProrationLast         V1ContractNewParamsRecurringCreditsProration = "LAST"
	V1ContractNewParamsRecurringCreditsProrationFirstAndLast V1ContractNewParamsRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r V1ContractNewParamsRecurringCreditsProration) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCreditsProrationNone, V1ContractNewParamsRecurringCreditsProrationFirst, V1ContractNewParamsRecurringCreditsProrationLast, V1ContractNewParamsRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// Whether the created commits will use the commit rate or list rate
type V1ContractNewParamsRecurringCreditsRateType string

const (
	V1ContractNewParamsRecurringCreditsRateTypeCommitRate V1ContractNewParamsRecurringCreditsRateType = "COMMIT_RATE"
	V1ContractNewParamsRecurringCreditsRateTypeListRate   V1ContractNewParamsRecurringCreditsRateType = "LIST_RATE"
)

func (r V1ContractNewParamsRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCreditsRateTypeCommitRate, V1ContractNewParamsRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type V1ContractNewParamsRecurringCreditsRecurrenceFrequency string

const (
	V1ContractNewParamsRecurringCreditsRecurrenceFrequencyMonthly   V1ContractNewParamsRecurringCreditsRecurrenceFrequency = "MONTHLY"
	V1ContractNewParamsRecurringCreditsRecurrenceFrequencyQuarterly V1ContractNewParamsRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	V1ContractNewParamsRecurringCreditsRecurrenceFrequencyAnnual    V1ContractNewParamsRecurringCreditsRecurrenceFrequency = "ANNUAL"
)

func (r V1ContractNewParamsRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsRecurringCreditsRecurrenceFrequencyMonthly, V1ContractNewParamsRecurringCreditsRecurrenceFrequencyQuarterly, V1ContractNewParamsRecurringCreditsRecurrenceFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsResellerRoyalty struct {
	Fraction           param.Field[float64]                                          `json:"fraction,required"`
	NetsuiteResellerID param.Field[string]                                           `json:"netsuite_reseller_id,required"`
	ResellerType       param.Field[V1ContractNewParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	StartingAt         param.Field[time.Time]                                        `json:"starting_at,required" format:"date-time"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags param.Field[[]string]                                       `json:"applicable_product_tags"`
	AwsOptions            param.Field[V1ContractNewParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
	EndingBefore          param.Field[time.Time]                                      `json:"ending_before" format:"date-time"`
	GcpOptions            param.Field[V1ContractNewParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	ResellerContractValue param.Field[float64]                                        `json:"reseller_contract_value"`
}

func (r V1ContractNewParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsResellerRoyaltiesResellerType string

const (
	V1ContractNewParamsResellerRoyaltiesResellerTypeAws           V1ContractNewParamsResellerRoyaltiesResellerType = "AWS"
	V1ContractNewParamsResellerRoyaltiesResellerTypeAwsProService V1ContractNewParamsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V1ContractNewParamsResellerRoyaltiesResellerTypeGcp           V1ContractNewParamsResellerRoyaltiesResellerType = "GCP"
	V1ContractNewParamsResellerRoyaltiesResellerTypeGcpProService V1ContractNewParamsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V1ContractNewParamsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsResellerRoyaltiesResellerTypeAws, V1ContractNewParamsResellerRoyaltiesResellerTypeAwsProService, V1ContractNewParamsResellerRoyaltiesResellerTypeGcp, V1ContractNewParamsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V1ContractNewParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r V1ContractNewParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r V1ContractNewParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[V1ContractNewParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractNewParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractNewParamsScheduledChargesSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractNewParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractNewParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractNewParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractNewParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                             `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractNewParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, V1ContractNewParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractNewParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type V1ContractNewParamsScheduledChargesOnUsageInvoices string

const (
	V1ContractNewParamsScheduledChargesOnUsageInvoicesAll V1ContractNewParamsScheduledChargesOnUsageInvoices = "ALL"
)

func (r V1ContractNewParamsScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case V1ContractNewParamsScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type V1ContractNewParamsSubscription struct {
	CollectionSchedule param.Field[V1ContractNewParamsSubscriptionsCollectionSchedule] `json:"collection_schedule,required"`
	InitialQuantity    param.Field[float64]                                            `json:"initial_quantity,required"`
	Proration          param.Field[V1ContractNewParamsSubscriptionsProration]          `json:"proration,required"`
	SubscriptionRate   param.Field[V1ContractNewParamsSubscriptionsSubscriptionRate]   `json:"subscription_rate,required"`
	Description        param.Field[string]                                             `json:"description"`
	// Exclusive end time for the subscription. If not provided, subscription inherits
	// contract end date.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Name         param.Field[string]    `json:"name"`
	// Inclusive start time for the subscription. If not provided, defaults to contract
	// start date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1ContractNewParamsSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsSubscriptionsCollectionSchedule string

const (
	V1ContractNewParamsSubscriptionsCollectionScheduleAdvance V1ContractNewParamsSubscriptionsCollectionSchedule = "ADVANCE"
	V1ContractNewParamsSubscriptionsCollectionScheduleArrears V1ContractNewParamsSubscriptionsCollectionSchedule = "ARREARS"
)

func (r V1ContractNewParamsSubscriptionsCollectionSchedule) IsKnown() bool {
	switch r {
	case V1ContractNewParamsSubscriptionsCollectionScheduleAdvance, V1ContractNewParamsSubscriptionsCollectionScheduleArrears:
		return true
	}
	return false
}

type V1ContractNewParamsSubscriptionsProration struct {
	// Indicates how mid-period quantity adjustments are invoiced. If BILL_IMMEDIATELY
	// is selected, the quantity increase will be billed on the scheduled date. If
	// BILL_ON_NEXT_COLLECTION_DATE is selected, the quantity increase will be billed
	// for in-arrears at the end of the period.
	InvoiceBehavior param.Field[V1ContractNewParamsSubscriptionsProrationInvoiceBehavior] `json:"invoice_behavior"`
	// Indicates if the partial period will be prorated or charged a full amount.
	IsProrated param.Field[bool] `json:"is_prorated"`
}

func (r V1ContractNewParamsSubscriptionsProration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates how mid-period quantity adjustments are invoiced. If BILL_IMMEDIATELY
// is selected, the quantity increase will be billed on the scheduled date. If
// BILL_ON_NEXT_COLLECTION_DATE is selected, the quantity increase will be billed
// for in-arrears at the end of the period.
type V1ContractNewParamsSubscriptionsProrationInvoiceBehavior string

const (
	V1ContractNewParamsSubscriptionsProrationInvoiceBehaviorBillImmediately          V1ContractNewParamsSubscriptionsProrationInvoiceBehavior = "BILL_IMMEDIATELY"
	V1ContractNewParamsSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate V1ContractNewParamsSubscriptionsProrationInvoiceBehavior = "BILL_ON_NEXT_COLLECTION_DATE"
)

func (r V1ContractNewParamsSubscriptionsProrationInvoiceBehavior) IsKnown() bool {
	switch r {
	case V1ContractNewParamsSubscriptionsProrationInvoiceBehaviorBillImmediately, V1ContractNewParamsSubscriptionsProrationInvoiceBehaviorBillOnNextCollectionDate:
		return true
	}
	return false
}

type V1ContractNewParamsSubscriptionsSubscriptionRate struct {
	// Frequency to bill subscription with. Together with product_id, must match
	// existing rate on the rate card.
	BillingFrequency param.Field[V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency] `json:"billing_frequency,required"`
	// Must be subscription type product
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
}

func (r V1ContractNewParamsSubscriptionsSubscriptionRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Frequency to bill subscription with. Together with product_id, must match
// existing rate on the rate card.
type V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency string

const (
	V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyMonthly   V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency = "MONTHLY"
	V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyQuarterly V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency = "QUARTERLY"
	V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyAnnual    V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency = "ANNUAL"
)

func (r V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyMonthly, V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyQuarterly, V1ContractNewParamsSubscriptionsSubscriptionRateBillingFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractNewParamsThresholdBillingConfiguration struct {
	Commit param.Field[V1ContractNewParamsThresholdBillingConfigurationCommit] `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state
	IsEnabled param.Field[bool] `json:"is_enabled,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount param.Field[float64] `json:"threshold_amount,required"`
}

func (r V1ContractNewParamsThresholdBillingConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsThresholdBillingConfigurationCommit struct {
	ProductID param.Field[string] `json:"product_id,required"`
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	Description           param.Field[string]   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name param.Field[string] `json:"name"`
}

func (r V1ContractNewParamsThresholdBillingConfigurationCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsTransition struct {
	FromContractID param.Field[string] `json:"from_contract_id,required" format:"uuid"`
	// This field's available values may vary based on your client's configuration.
	Type                  param.Field[V1ContractNewParamsTransitionType]                  `json:"type,required"`
	FutureInvoiceBehavior param.Field[V1ContractNewParamsTransitionFutureInvoiceBehavior] `json:"future_invoice_behavior"`
}

func (r V1ContractNewParamsTransition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field's available values may vary based on your client's configuration.
type V1ContractNewParamsTransitionType string

const (
	V1ContractNewParamsTransitionTypeSupersede V1ContractNewParamsTransitionType = "SUPERSEDE"
	V1ContractNewParamsTransitionTypeRenewal   V1ContractNewParamsTransitionType = "RENEWAL"
)

func (r V1ContractNewParamsTransitionType) IsKnown() bool {
	switch r {
	case V1ContractNewParamsTransitionTypeSupersede, V1ContractNewParamsTransitionTypeRenewal:
		return true
	}
	return false
}

type V1ContractNewParamsTransitionFutureInvoiceBehavior struct {
	// Controls whether future trueup invoices are billed or removed. Default behavior
	// is AS_IS if not specified.
	Trueup param.Field[V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueup] `json:"trueup"`
}

func (r V1ContractNewParamsTransitionFutureInvoiceBehavior) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls whether future trueup invoices are billed or removed. Default behavior
// is AS_IS if not specified.
type V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueup string

const (
	V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "REMOVE"
	V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs   V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueup = "AS_IS"
)

func (r V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueup) IsKnown() bool {
	switch r {
	case V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueupRemove, V1ContractNewParamsTransitionFutureInvoiceBehaviorTrueupAsIs:
		return true
	}
	return false
}

type V1ContractNewParamsUsageStatementSchedule struct {
	Frequency param.Field[V1ContractNewParamsUsageStatementScheduleFrequency] `json:"frequency,required"`
	// Required when using CUSTOM_DATE. This option lets you set a historical billing
	// anchor date, aligning future billing cycles with a chosen cadence. For example,
	// if a contract starts on 2024-09-15 and you set the anchor date to 2024-09-10
	// with a MONTHLY frequency, the first usage statement will cover 09-15 to 10-10.
	// Subsequent statements will follow the 10th of each month.
	BillingAnchorDate param.Field[time.Time] `json:"billing_anchor_date" format:"date-time"`
	// If not provided, defaults to the first day of the month.
	Day param.Field[V1ContractNewParamsUsageStatementScheduleDay] `json:"day"`
	// The date Metronome should start generating usage invoices. If unspecified,
	// contract start date will be used. This is useful to set if you want to import
	// historical invoices via our 'Create Historical Invoices' API rather than having
	// Metronome automatically generate them.
	InvoiceGenerationStartingAt param.Field[time.Time] `json:"invoice_generation_starting_at" format:"date-time"`
}

func (r V1ContractNewParamsUsageStatementSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewParamsUsageStatementScheduleFrequency string

const (
	V1ContractNewParamsUsageStatementScheduleFrequencyMonthly   V1ContractNewParamsUsageStatementScheduleFrequency = "MONTHLY"
	V1ContractNewParamsUsageStatementScheduleFrequencyQuarterly V1ContractNewParamsUsageStatementScheduleFrequency = "QUARTERLY"
	V1ContractNewParamsUsageStatementScheduleFrequencyAnnual    V1ContractNewParamsUsageStatementScheduleFrequency = "ANNUAL"
)

func (r V1ContractNewParamsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractNewParamsUsageStatementScheduleFrequencyMonthly, V1ContractNewParamsUsageStatementScheduleFrequencyQuarterly, V1ContractNewParamsUsageStatementScheduleFrequencyAnnual:
		return true
	}
	return false
}

// If not provided, defaults to the first day of the month.
type V1ContractNewParamsUsageStatementScheduleDay string

const (
	V1ContractNewParamsUsageStatementScheduleDayFirstOfMonth  V1ContractNewParamsUsageStatementScheduleDay = "FIRST_OF_MONTH"
	V1ContractNewParamsUsageStatementScheduleDayContractStart V1ContractNewParamsUsageStatementScheduleDay = "CONTRACT_START"
	V1ContractNewParamsUsageStatementScheduleDayCustomDate    V1ContractNewParamsUsageStatementScheduleDay = "CUSTOM_DATE"
)

func (r V1ContractNewParamsUsageStatementScheduleDay) IsKnown() bool {
	switch r {
	case V1ContractNewParamsUsageStatementScheduleDayFirstOfMonth, V1ContractNewParamsUsageStatementScheduleDayContractStart, V1ContractNewParamsUsageStatementScheduleDayCustomDate:
		return true
	}
	return false
}

type V1ContractGetParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
}

func (r V1ContractGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractListParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts effective on the provided date. This cannot be provided if the
	// starting_at filter is provided.
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include archived contracts in the response
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include commit ledgers in the response. Setting this flag may cause the query to
	// be slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// Optional RFC 3339 timestamp. If provided, the response will include only
	// contracts where effective_at is on or after the provided date. This cannot be
	// provided if the covering_date filter is provided.
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1ContractListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAddManualBalanceEntryParams struct {
	// ID of the balance (commit or credit) to update.
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// Amount to add to the segment. A negative number will draw down from the balance.
	Amount param.Field[float64] `json:"amount,required"`
	// ID of the customer whose balance is to be updated.
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Reason for the manual adjustment. This will be displayed in the ledger.
	Reason param.Field[string] `json:"reason,required"`
	// ID of the segment to update.
	SegmentID param.Field[string] `json:"segment_id,required" format:"uuid"`
	// ID of the contract to update. Leave blank to update a customer level balance.
	ContractID param.Field[string] `json:"contract_id" format:"uuid"`
	// RFC 3339 timestamp indicating when the manual adjustment takes place. If not
	// provided, it will default to the start of the segment.
	Timestamp param.Field[time.Time] `json:"timestamp" format:"date-time"`
}

func (r V1ContractAddManualBalanceEntryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParams struct {
	// ID of the contract to amend
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be amended
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// inclusive start time for the amendment
	StartingAt   param.Field[time.Time]                     `json:"starting_at,required" format:"date-time"`
	Commits      param.Field[[]V1ContractAmendParamsCommit] `json:"commits"`
	Credits      param.Field[[]V1ContractAmendParamsCredit] `json:"credits"`
	CustomFields param.Field[map[string]string]             `json:"custom_fields"`
	// This field's availability is dependent on your client's configuration.
	Discounts param.Field[[]V1ContractAmendParamsDiscount] `json:"discounts"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string]                          `json:"netsuite_sales_order_id"`
	Overrides            param.Field[[]V1ContractAmendParamsOverride] `json:"overrides"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices param.Field[[]V1ContractAmendParamsProfessionalService] `json:"professional_services"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties param.Field[[]V1ContractAmendParamsResellerRoyalty] `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID param.Field[string]                                 `json:"salesforce_opportunity_id"`
	ScheduledCharges        param.Field[[]V1ContractAmendParamsScheduledCharge] `json:"scheduled_charges"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue param.Field[float64] `json:"total_contract_value"`
}

func (r V1ContractAmendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCommit struct {
	ProductID param.Field[string]                           `json:"product_id,required" format:"uuid"`
	Type      param.Field[V1ContractAmendParamsCommitsType] `json:"type,required"`
	// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
	// commits only one schedule item is allowed and amount must match invoice_schedule
	// total.
	AccessSchedule param.Field[V1ContractAmendParamsCommitsAccessSchedule] `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule and invoice_schedule instead.
	Amount param.Field[float64] `json:"amount"`
	// Which products the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this
	// time and only one schedule item is allowed; the total must match access_schedule
	// amount. Optional for "PREPAID" commits: if not provided, this will be a
	// "complimentary" commit with no invoice.
	InvoiceSchedule param.Field[V1ContractAmendParamsCommitsInvoiceSchedule] `json:"invoice_schedule"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple commits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                              `json:"priority"`
	RateType param.Field[V1ContractAmendParamsCommitsRateType] `json:"rate_type"`
	// Fraction of unused segments that will be rolled over. Must be between 0 and 1.
	RolloverFraction param.Field[float64] `json:"rollover_fraction"`
	// A temporary ID for the commit that can be used to reference the commit for
	// commit specific overrides.
	TemporaryID param.Field[string] `json:"temporary_id"`
}

func (r V1ContractAmendParamsCommit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCommitsType string

const (
	V1ContractAmendParamsCommitsTypePrepaid  V1ContractAmendParamsCommitsType = "PREPAID"
	V1ContractAmendParamsCommitsTypePostpaid V1ContractAmendParamsCommitsType = "POSTPAID"
)

func (r V1ContractAmendParamsCommitsType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsCommitsTypePrepaid, V1ContractAmendParamsCommitsTypePostpaid:
		return true
	}
	return false
}

// Required: Schedule for distributing the commit to the customer. For "POSTPAID"
// commits only one schedule item is allowed and amount must match invoice_schedule
// total.
type V1ContractAmendParamsCommitsAccessSchedule struct {
	ScheduleItems param.Field[[]V1ContractAmendParamsCommitsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1ContractAmendParamsCommitsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCommitsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1ContractAmendParamsCommitsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Required for "POSTPAID" commits: the true up invoice will be generated at this
// time and only one schedule item is allowed; the total must match access_schedule
// amount. Optional for "PREPAID" commits: if not provided, this will be a
// "complimentary" commit with no invoice.
type V1ContractAmendParamsCommitsInvoiceSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractAmendParamsCommitsInvoiceScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractAmendParamsCommitsInvoiceSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                             `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsCommitsInvoiceScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided        V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach           V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided, V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency string

const (
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly    V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly  V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual     V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly, V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyQuarterly, V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencySemiAnnual, V1ContractAmendParamsCommitsInvoiceScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractAmendParamsCommitsInvoiceScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsCommitsInvoiceScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCommitsRateType string

const (
	V1ContractAmendParamsCommitsRateTypeCommitRate V1ContractAmendParamsCommitsRateType = "COMMIT_RATE"
	V1ContractAmendParamsCommitsRateTypeListRate   V1ContractAmendParamsCommitsRateType = "LIST_RATE"
)

func (r V1ContractAmendParamsCommitsRateType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsCommitsRateTypeCommitRate, V1ContractAmendParamsCommitsRateTypeListRate:
		return true
	}
	return false
}

type V1ContractAmendParamsCredit struct {
	// Schedule for distributing the credit to the customer.
	AccessSchedule param.Field[V1ContractAmendParamsCreditsAccessSchedule] `json:"access_schedule,required"`
	ProductID      param.Field[string]                                     `json:"product_id,required" format:"uuid"`
	// Which products the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Which tags the credit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the credit applies to all products.
	ApplicableProductTags param.Field[[]string]          `json:"applicable_product_tags"`
	CustomFields          param.Field[map[string]string] `json:"custom_fields"`
	// Used only in UI/API. It is not exposed to end customers.
	Description param.Field[string] `json:"description"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
	// If multiple credits are applicable, the one with the lower priority will apply
	// first.
	Priority param.Field[float64]                              `json:"priority"`
	RateType param.Field[V1ContractAmendParamsCreditsRateType] `json:"rate_type"`
}

func (r V1ContractAmendParamsCredit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Schedule for distributing the credit to the customer.
type V1ContractAmendParamsCreditsAccessSchedule struct {
	ScheduleItems param.Field[[]V1ContractAmendParamsCreditsAccessScheduleScheduleItem] `json:"schedule_items,required"`
	// Defaults to USD (cents) if not passed
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
}

func (r V1ContractAmendParamsCreditsAccessSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCreditsAccessScheduleScheduleItem struct {
	Amount param.Field[float64] `json:"amount,required"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before,required" format:"date-time"`
	// RFC 3339 timestamp (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1ContractAmendParamsCreditsAccessScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsCreditsRateType string

const (
	V1ContractAmendParamsCreditsRateTypeCommitRate V1ContractAmendParamsCreditsRateType = "COMMIT_RATE"
	V1ContractAmendParamsCreditsRateTypeListRate   V1ContractAmendParamsCreditsRateType = "LIST_RATE"
)

func (r V1ContractAmendParamsCreditsRateType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsCreditsRateTypeCommitRate, V1ContractAmendParamsCreditsRateTypeListRate:
		return true
	}
	return false
}

type V1ContractAmendParamsDiscount struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule     param.Field[V1ContractAmendParamsDiscountsSchedule] `json:"schedule,required"`
	CustomFields param.Field[map[string]string]                      `json:"custom_fields"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractAmendParamsDiscount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractAmendParamsDiscountsSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractAmendParamsDiscountsScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractAmendParamsDiscountsScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractAmendParamsDiscountsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractAmendParamsDiscountsScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                        `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsDiscountsScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided        V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach           V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDivided, V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractAmendParamsDiscountsScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency string

const (
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly    V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly  V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual     V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyMonthly, V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyQuarterly, V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencySemiAnnual, V1ContractAmendParamsDiscountsScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractAmendParamsDiscountsScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsDiscountsScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsOverride struct {
	// RFC 3339 timestamp indicating when the override will start applying (inclusive)
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// tags identifying products whose rates are being overridden. Cannot be used in
	// conjunction with override_specifiers.
	ApplicableProductTags param.Field[[]string] `json:"applicable_product_tags"`
	// RFC 3339 timestamp indicating when the override will stop applying (exclusive)
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
	Entitled     param.Field[bool]      `json:"entitled"`
	// Indicates whether the override should only apply to commits. Defaults to
	// `false`. If `true`, you can specify relevant commits in `override_specifiers` by
	// passing `commit_ids`. if you do not specify `commit_ids`, then the override will
	// apply when consuming any prepaid or postpaid commit.
	IsCommitSpecific param.Field[bool] `json:"is_commit_specific"`
	// Required for MULTIPLIER type. Must be >=0.
	Multiplier param.Field[float64] `json:"multiplier"`
	// Cannot be used in conjunction with product_id or applicable_product_tags. If
	// provided, the override will apply to all products with the specified specifiers.
	OverrideSpecifiers param.Field[[]V1ContractAmendParamsOverridesOverrideSpecifier] `json:"override_specifiers"`
	// Required for OVERWRITE type.
	OverwriteRate param.Field[V1ContractAmendParamsOverridesOverwriteRate] `json:"overwrite_rate"`
	// Required for EXPLICIT multiplier prioritization scheme and all TIERED overrides.
	// Under EXPLICIT prioritization, overwrites are prioritized first, and then tiered
	// and multiplier overrides are prioritized by their priority value (lowest first).
	// Must be > 0.
	Priority param.Field[float64] `json:"priority"`
	// ID of the product whose rate is being overridden. Cannot be used in conjunction
	// with override_specifiers.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// Indicates whether the override applies to commit rates or list rates. Can only
	// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
	// `"LIST_RATE"`.
	Target param.Field[V1ContractAmendParamsOverridesTarget] `json:"target"`
	// Required for TIERED type. Must have at least one tier.
	Tiers param.Field[[]V1ContractAmendParamsOverridesTier] `json:"tiers"`
	// Overwrites are prioritized over multipliers and tiered overrides.
	Type param.Field[V1ContractAmendParamsOverridesType] `json:"type"`
}

func (r V1ContractAmendParamsOverride) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsOverridesOverrideSpecifier struct {
	BillingFrequency param.Field[V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency] `json:"billing_frequency"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to the
	// specified commits. If not provided, the override will apply to all commits.
	CommitIDs param.Field[[]string] `json:"commit_ids"`
	// A map of group names to values. The override will only apply to line items with
	// the specified presentation group values.
	PresentationGroupValues param.Field[map[string]string] `json:"presentation_group_values"`
	// A map of pricing group names to values. The override will only apply to products
	// with the specified pricing group values.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// If provided, the override will only apply to the product with the specified ID.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// If provided, the override will only apply to products with all the specified
	// tags.
	ProductTags param.Field[[]string] `json:"product_tags"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to commits
	// created by the specified recurring commit ids.
	RecurringCommitIDs param.Field[[]string] `json:"recurring_commit_ids"`
	// Can only be used for commit specific overrides. Must be used in conjunction with
	// one of product_id, product_tags, pricing_group_values, or
	// presentation_group_values. If provided, the override will only apply to credits
	// created by the specified recurring credit ids.
	RecurringCreditIDs param.Field[[]string] `json:"recurring_credit_ids"`
}

func (r V1ContractAmendParamsOverridesOverrideSpecifier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency string

const (
	V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyMonthly   V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency = "MONTHLY"
	V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyQuarterly V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency = "QUARTERLY"
	V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyAnnual    V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency = "ANNUAL"
)

func (r V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyMonthly, V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyQuarterly, V1ContractAmendParamsOverridesOverrideSpecifiersBillingFrequencyAnnual:
		return true
	}
	return false
}

// Required for OVERWRITE type.
type V1ContractAmendParamsOverridesOverwriteRate struct {
	RateType     param.Field[V1ContractAmendParamsOverridesOverwriteRateRateType] `json:"rate_type,required"`
	CreditTypeID param.Field[string]                                              `json:"credit_type_id" format:"uuid"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate param.Field[map[string]interface{}] `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated param.Field[bool] `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price param.Field[float64] `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity param.Field[float64] `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers param.Field[[]shared.TierParam] `json:"tiers"`
}

func (r V1ContractAmendParamsOverridesOverwriteRate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsOverridesOverwriteRateRateType string

const (
	V1ContractAmendParamsOverridesOverwriteRateRateTypeFlat         V1ContractAmendParamsOverridesOverwriteRateRateType = "FLAT"
	V1ContractAmendParamsOverridesOverwriteRateRateTypePercentage   V1ContractAmendParamsOverridesOverwriteRateRateType = "PERCENTAGE"
	V1ContractAmendParamsOverridesOverwriteRateRateTypeSubscription V1ContractAmendParamsOverridesOverwriteRateRateType = "SUBSCRIPTION"
	V1ContractAmendParamsOverridesOverwriteRateRateTypeTiered       V1ContractAmendParamsOverridesOverwriteRateRateType = "TIERED"
	V1ContractAmendParamsOverridesOverwriteRateRateTypeCustom       V1ContractAmendParamsOverridesOverwriteRateRateType = "CUSTOM"
)

func (r V1ContractAmendParamsOverridesOverwriteRateRateType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsOverridesOverwriteRateRateTypeFlat, V1ContractAmendParamsOverridesOverwriteRateRateTypePercentage, V1ContractAmendParamsOverridesOverwriteRateRateTypeSubscription, V1ContractAmendParamsOverridesOverwriteRateRateTypeTiered, V1ContractAmendParamsOverridesOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

// Indicates whether the override applies to commit rates or list rates. Can only
// be used for overrides that have `is_commit_specific` set to `true`. Defaults to
// `"LIST_RATE"`.
type V1ContractAmendParamsOverridesTarget string

const (
	V1ContractAmendParamsOverridesTargetCommitRate V1ContractAmendParamsOverridesTarget = "COMMIT_RATE"
	V1ContractAmendParamsOverridesTargetListRate   V1ContractAmendParamsOverridesTarget = "LIST_RATE"
)

func (r V1ContractAmendParamsOverridesTarget) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsOverridesTargetCommitRate, V1ContractAmendParamsOverridesTargetListRate:
		return true
	}
	return false
}

type V1ContractAmendParamsOverridesTier struct {
	Multiplier param.Field[float64] `json:"multiplier,required"`
	Size       param.Field[float64] `json:"size"`
}

func (r V1ContractAmendParamsOverridesTier) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Overwrites are prioritized over multipliers and tiered overrides.
type V1ContractAmendParamsOverridesType string

const (
	V1ContractAmendParamsOverridesTypeOverwrite  V1ContractAmendParamsOverridesType = "OVERWRITE"
	V1ContractAmendParamsOverridesTypeMultiplier V1ContractAmendParamsOverridesType = "MULTIPLIER"
	V1ContractAmendParamsOverridesTypeTiered     V1ContractAmendParamsOverridesType = "TIERED"
)

func (r V1ContractAmendParamsOverridesType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsOverridesTypeOverwrite, V1ContractAmendParamsOverridesTypeMultiplier, V1ContractAmendParamsOverridesTypeTiered:
		return true
	}
	return false
}

type V1ContractAmendParamsProfessionalService struct {
	// Maximum amount for the term.
	MaxAmount param.Field[float64] `json:"max_amount,required"`
	ProductID param.Field[string]  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    param.Field[float64]           `json:"unit_price,required"`
	CustomFields param.Field[map[string]string] `json:"custom_fields"`
	Description  param.Field[string]            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractAmendParamsProfessionalService) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsResellerRoyalty struct {
	ResellerType param.Field[V1ContractAmendParamsResellerRoyaltiesResellerType] `json:"reseller_type,required"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductIDs param.Field[[]string] `json:"applicable_product_ids" format:"uuid"`
	// Must provide at least one of applicable_product_ids or applicable_product_tags.
	ApplicableProductTags param.Field[[]string]                                         `json:"applicable_product_tags"`
	AwsOptions            param.Field[V1ContractAmendParamsResellerRoyaltiesAwsOptions] `json:"aws_options"`
	// Use null to indicate that the existing end timestamp should be removed.
	EndingBefore          param.Field[time.Time]                                        `json:"ending_before" format:"date-time"`
	Fraction              param.Field[float64]                                          `json:"fraction"`
	GcpOptions            param.Field[V1ContractAmendParamsResellerRoyaltiesGcpOptions] `json:"gcp_options"`
	NetsuiteResellerID    param.Field[string]                                           `json:"netsuite_reseller_id"`
	ResellerContractValue param.Field[float64]                                          `json:"reseller_contract_value"`
	StartingAt            param.Field[time.Time]                                        `json:"starting_at" format:"date-time"`
}

func (r V1ContractAmendParamsResellerRoyalty) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsResellerRoyaltiesResellerType string

const (
	V1ContractAmendParamsResellerRoyaltiesResellerTypeAws           V1ContractAmendParamsResellerRoyaltiesResellerType = "AWS"
	V1ContractAmendParamsResellerRoyaltiesResellerTypeAwsProService V1ContractAmendParamsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	V1ContractAmendParamsResellerRoyaltiesResellerTypeGcp           V1ContractAmendParamsResellerRoyaltiesResellerType = "GCP"
	V1ContractAmendParamsResellerRoyaltiesResellerTypeGcpProService V1ContractAmendParamsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r V1ContractAmendParamsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsResellerRoyaltiesResellerTypeAws, V1ContractAmendParamsResellerRoyaltiesResellerTypeAwsProService, V1ContractAmendParamsResellerRoyaltiesResellerTypeGcp, V1ContractAmendParamsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

type V1ContractAmendParamsResellerRoyaltiesAwsOptions struct {
	AwsAccountNumber    param.Field[string] `json:"aws_account_number"`
	AwsOfferID          param.Field[string] `json:"aws_offer_id"`
	AwsPayerReferenceID param.Field[string] `json:"aws_payer_reference_id"`
}

func (r V1ContractAmendParamsResellerRoyaltiesAwsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsResellerRoyaltiesGcpOptions struct {
	GcpAccountID param.Field[string] `json:"gcp_account_id"`
	GcpOfferID   param.Field[string] `json:"gcp_offer_id"`
}

func (r V1ContractAmendParamsResellerRoyaltiesGcpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsScheduledCharge struct {
	ProductID param.Field[string] `json:"product_id,required" format:"uuid"`
	// Must provide either schedule_items or recurring_schedule.
	Schedule param.Field[V1ContractAmendParamsScheduledChargesSchedule] `json:"schedule,required"`
	// displayed on invoices
	Name param.Field[string] `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID param.Field[string] `json:"netsuite_sales_order_id"`
}

func (r V1ContractAmendParamsScheduledCharge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Must provide either schedule_items or recurring_schedule.
type V1ContractAmendParamsScheduledChargesSchedule struct {
	// Defaults to USD (cents) if not passed.
	CreditTypeID param.Field[string] `json:"credit_type_id" format:"uuid"`
	// Enter the unit price and quantity for the charge or instead only send the
	// amount. If amount is sent, the unit price is assumed to be the amount and
	// quantity is inferred to be 1.
	RecurringSchedule param.Field[V1ContractAmendParamsScheduledChargesScheduleRecurringSchedule] `json:"recurring_schedule"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems param.Field[[]V1ContractAmendParamsScheduledChargesScheduleScheduleItem] `json:"schedule_items"`
}

func (r V1ContractAmendParamsScheduledChargesSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enter the unit price and quantity for the charge or instead only send the
// amount. If amount is sent, the unit price is assumed to be the amount and
// quantity is inferred to be 1.
type V1ContractAmendParamsScheduledChargesScheduleRecurringSchedule struct {
	AmountDistribution param.Field[V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution] `json:"amount_distribution,required"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore param.Field[time.Time]                                                               `json:"ending_before,required" format:"date-time"`
	Frequency    param.Field[V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency] `json:"frequency,required"`
	// RFC 3339 timestamp (inclusive).
	StartingAt param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsScheduledChargesScheduleRecurringSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution string

const (
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided        V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED"
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "DIVIDED_ROUNDED"
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach           V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution = "EACH"
)

func (r V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistribution) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDivided, V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionDividedRounded, V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleAmountDistributionEach:
		return true
	}
	return false
}

type V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency string

const (
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly    V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "MONTHLY"
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly  V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "QUARTERLY"
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "SEMI_ANNUAL"
	V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual     V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency = "ANNUAL"
)

func (r V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequency) IsKnown() bool {
	switch r {
	case V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyMonthly, V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyQuarterly, V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencySemiAnnual, V1ContractAmendParamsScheduledChargesScheduleRecurringScheduleFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractAmendParamsScheduledChargesScheduleScheduleItem struct {
	// timestamp of the scheduled event
	Timestamp param.Field[time.Time] `json:"timestamp,required" format:"date-time"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If
	// amount is sent, the unit_price is assumed to be the amount and quantity is
	// inferred to be 1.
	Amount param.Field[float64] `json:"amount"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount and must be specified with unit_price. If specified amount cannot be
	// provided.
	Quantity param.Field[float64] `json:"quantity"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified with quantity. If specified amount cannot be
	// provided.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractAmendParamsScheduledChargesScheduleScheduleItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractArchiveParams struct {
	// ID of the contract to archive
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be archived
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// If false, the existing finalized invoices will remain after the contract is
	// archived.
	VoidInvoices param.Field[bool] `json:"void_invoices,required"`
}

func (r V1ContractArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewHistoricalInvoicesParams struct {
	Invoices param.Field[[]V1ContractNewHistoricalInvoicesParamsInvoice] `json:"invoices,required"`
	Preview  param.Field[bool]                                           `json:"preview,required"`
}

func (r V1ContractNewHistoricalInvoicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewHistoricalInvoicesParamsInvoice struct {
	ContractID         param.Field[string]                                                       `json:"contract_id,required" format:"uuid"`
	CreditTypeID       param.Field[string]                                                       `json:"credit_type_id,required" format:"uuid"`
	CustomerID         param.Field[string]                                                       `json:"customer_id,required" format:"uuid"`
	ExclusiveEndDate   param.Field[time.Time]                                                    `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate param.Field[time.Time]                                                    `json:"inclusive_start_date,required" format:"date-time"`
	IssueDate          param.Field[time.Time]                                                    `json:"issue_date,required" format:"date-time"`
	UsageLineItems     param.Field[[]V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem] `json:"usage_line_items,required"`
	// This field's availability is dependent on your client's configuration.
	BillableStatus       param.Field[V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatus]       `json:"billable_status"`
	BreakdownGranularity param.Field[V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity] `json:"breakdown_granularity"`
	CustomFields         param.Field[map[string]string]                                                 `json:"custom_fields"`
}

func (r V1ContractNewHistoricalInvoicesParamsInvoice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem struct {
	ExclusiveEndDate        param.Field[time.Time]                                                                          `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate      param.Field[time.Time]                                                                          `json:"inclusive_start_date,required" format:"date-time"`
	ProductID               param.Field[string]                                                                             `json:"product_id,required" format:"uuid"`
	PresentationGroupValues param.Field[map[string]string]                                                                  `json:"presentation_group_values"`
	PricingGroupValues      param.Field[map[string]string]                                                                  `json:"pricing_group_values"`
	Quantity                param.Field[float64]                                                                            `json:"quantity"`
	SubtotalsWithQuantity   param.Field[[]V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity] `json:"subtotals_with_quantity"`
}

func (r V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity struct {
	ExclusiveEndDate   param.Field[time.Time] `json:"exclusive_end_date,required" format:"date-time"`
	InclusiveStartDate param.Field[time.Time] `json:"inclusive_start_date,required" format:"date-time"`
	Quantity           param.Field[float64]   `json:"quantity,required"`
}

func (r V1ContractNewHistoricalInvoicesParamsInvoicesUsageLineItemsSubtotalsWithQuantity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field's availability is dependent on your client's configuration.
type V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatus string

const (
	V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable   V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatus = "billable"
	V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatusUnbillable V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatus = "unbillable"
)

func (r V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatus) IsKnown() bool {
	switch r {
	case V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatusBillable, V1ContractNewHistoricalInvoicesParamsInvoicesBillableStatusUnbillable:
		return true
	}
	return false
}

type V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity string

const (
	V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity = "HOUR"
	V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityDay  V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity = "DAY"
)

func (r V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularity) IsKnown() bool {
	switch r {
	case V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityHour, V1ContractNewHistoricalInvoicesParamsInvoicesBreakdownGranularityDay:
		return true
	}
	return false
}

type V1ContractListBalancesParams struct {
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	ID         param.Field[string] `json:"id" format:"uuid"`
	// Return only balances that have access schedules that "cover" the provided date
	CoveringDate param.Field[time.Time] `json:"covering_date" format:"date-time"`
	// Include only balances that have any access before the provided date (exclusive)
	EffectiveBefore param.Field[time.Time] `json:"effective_before" format:"date-time"`
	// Include credits from archived contracts.
	IncludeArchived param.Field[bool] `json:"include_archived"`
	// Include the balance of credits and commits in the response. Setting this flag
	// may cause the query to be slower.
	IncludeBalance param.Field[bool] `json:"include_balance"`
	// Include balances on the contract level.
	IncludeContractBalances param.Field[bool] `json:"include_contract_balances"`
	// Include ledgers in the response. Setting this flag may cause the query to be
	// slower.
	IncludeLedgers param.Field[bool] `json:"include_ledgers"`
	// The next page token from a previous response.
	NextPage param.Field[string] `json:"next_page"`
	// Include only balances that have any access on or after the provided date
	StartingAt param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r V1ContractListBalancesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractGetRateScheduleParams struct {
	// ID of the contract to get the rate schedule for.
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer for whose contract to get the rate schedule for.
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// Max number of results that should be returned
	Limit param.Field[int64] `query:"limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage param.Field[string] `query:"next_page"`
	// optional timestamp which overlaps with the returned rate schedule segments. When
	// not specified, the current timestamp will be used.
	At param.Field[time.Time] `json:"at" format:"date-time"`
	// List of rate selectors, rates matching ANY of the selectors will be included in
	// the response. Passing no selectors will result in all rates being returned.
	Selectors param.Field[[]V1ContractGetRateScheduleParamsSelector] `json:"selectors"`
}

func (r V1ContractGetRateScheduleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [V1ContractGetRateScheduleParams]'s query parameters as
// `url.Values`.
func (r V1ContractGetRateScheduleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContractGetRateScheduleParamsSelector struct {
	// Subscription rates matching the billing frequency will be included in the
	// response.
	BillingFrequency param.Field[V1ContractGetRateScheduleParamsSelectorsBillingFrequency] `json:"billing_frequency"`
	// List of pricing group key value pairs, rates containing the matching key / value
	// pairs will be included in the response.
	PartialPricingGroupValues param.Field[map[string]string] `json:"partial_pricing_group_values"`
	// List of pricing group key value pairs, rates matching all of the key / value
	// pairs will be included in the response.
	PricingGroupValues param.Field[map[string]string] `json:"pricing_group_values"`
	// Rates matching the product id will be included in the response.
	ProductID param.Field[string] `json:"product_id" format:"uuid"`
	// List of product tags, rates matching any of the tags will be included in the
	// response.
	ProductTags param.Field[[]string] `json:"product_tags"`
}

func (r V1ContractGetRateScheduleParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Subscription rates matching the billing frequency will be included in the
// response.
type V1ContractGetRateScheduleParamsSelectorsBillingFrequency string

const (
	V1ContractGetRateScheduleParamsSelectorsBillingFrequencyMonthly   V1ContractGetRateScheduleParamsSelectorsBillingFrequency = "MONTHLY"
	V1ContractGetRateScheduleParamsSelectorsBillingFrequencyQuarterly V1ContractGetRateScheduleParamsSelectorsBillingFrequency = "QUARTERLY"
	V1ContractGetRateScheduleParamsSelectorsBillingFrequencyAnnual    V1ContractGetRateScheduleParamsSelectorsBillingFrequency = "ANNUAL"
)

func (r V1ContractGetRateScheduleParamsSelectorsBillingFrequency) IsKnown() bool {
	switch r {
	case V1ContractGetRateScheduleParamsSelectorsBillingFrequencyMonthly, V1ContractGetRateScheduleParamsSelectorsBillingFrequencyQuarterly, V1ContractGetRateScheduleParamsSelectorsBillingFrequencyAnnual:
		return true
	}
	return false
}

type V1ContractScheduleProServicesInvoiceParams struct {
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// The date the invoice is issued
	IssuedAt param.Field[time.Time] `json:"issued_at,required" format:"date-time"`
	// Each line requires an amount or both unit_price and quantity.
	LineItems param.Field[[]V1ContractScheduleProServicesInvoiceParamsLineItem] `json:"line_items,required"`
	// The end date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderEnd param.Field[time.Time] `json:"netsuite_invoice_header_end" format:"date-time"`
	// The start date of the invoice header in Netsuite
	NetsuiteInvoiceHeaderStart param.Field[time.Time] `json:"netsuite_invoice_header_start" format:"date-time"`
}

func (r V1ContractScheduleProServicesInvoiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Describes the line item for a professional service charge on an invoice.
type V1ContractScheduleProServicesInvoiceParamsLineItem struct {
	ProfessionalServiceID param.Field[string] `json:"professional_service_id,required" format:"uuid"`
	// If the professional_service_id was added on an amendment, this is required.
	AmendmentID param.Field[string] `json:"amendment_id" format:"uuid"`
	// Amount for the term on the new invoice.
	Amount param.Field[float64] `json:"amount"`
	// For client use.
	Metadata param.Field[string] `json:"metadata"`
	// The end date for the billing period on the invoice.
	NetsuiteInvoiceBillingEnd param.Field[time.Time] `json:"netsuite_invoice_billing_end" format:"date-time"`
	// The start date for the billing period on the invoice.
	NetsuiteInvoiceBillingStart param.Field[time.Time] `json:"netsuite_invoice_billing_start" format:"date-time"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity param.Field[float64] `json:"quantity"`
	// If specified, this overrides the unit price on the pro service term. Must also
	// provide quantity (but not amount) if providing unit_price.
	UnitPrice param.Field[float64] `json:"unit_price"`
}

func (r V1ContractScheduleProServicesInvoiceParamsLineItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractSetUsageFilterParams struct {
	ContractID  param.Field[string]    `json:"contract_id,required" format:"uuid"`
	CustomerID  param.Field[string]    `json:"customer_id,required" format:"uuid"`
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at,required" format:"date-time"`
}

func (r V1ContractSetUsageFilterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ContractUpdateEndDateParams struct {
	// ID of the contract to update
	ContractID param.Field[string] `json:"contract_id,required" format:"uuid"`
	// ID of the customer whose contract is to be updated
	CustomerID param.Field[string] `json:"customer_id,required" format:"uuid"`
	// If true, allows setting the contract end date earlier than the end_timestamp of
	// existing finalized invoices. Finalized invoices will be unchanged; if you want
	// to incorporate the new end date, you can void and regenerate finalized usage
	// invoices. Defaults to true.
	AllowEndingBeforeFinalizedInvoice param.Field[bool] `json:"allow_ending_before_finalized_invoice"`
	// RFC 3339 timestamp indicating when the contract will end (exclusive). If not
	// provided, the contract will be updated to be open-ended.
	EndingBefore param.Field[time.Time] `json:"ending_before" format:"date-time"`
}

func (r V1ContractUpdateEndDateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
