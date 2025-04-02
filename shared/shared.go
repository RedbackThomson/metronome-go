// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/Metronome-Industries/metronome-go/internal/apijson"
	"github.com/Metronome-Industries/metronome-go/packages/param"
	"github.com/tidwall/gjson"
)

type BaseUsageFilter struct {
	GroupKey    string              `json:"group_key,required"`
	GroupValues []string            `json:"group_values,required"`
	StartingAt  time.Time           `json:"starting_at" format:"date-time"`
	JSON        baseUsageFilterJSON `json:"-"`
}

// baseUsageFilterJSON contains the JSON metadata for the struct [BaseUsageFilter]
type baseUsageFilterJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseUsageFilterJSON) RawJSON() string {
	return r.raw
}

type BaseUsageFilterParam struct {
	GroupKey    param.Field[string]    `json:"group_key,required"`
	GroupValues param.Field[[]string]  `json:"group_values,required"`
	StartingAt  param.Field[time.Time] `json:"starting_at" format:"date-time"`
}

func (r BaseUsageFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Commit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CommitProduct `json:"product,required"`
	Type    CommitType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits purposed with
	// this commit.
	AccessSchedule ScheduleDuration `json:"access_schedule"`
	// (DEPRECATED) Use access_schedule + invoice_schedule instead.
	Amount                float64  `json:"amount"`
	ApplicableContractIDs []string `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64           `json:"balance"`
	Contract     CommitContract    `json:"contract"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// The contract that this commit will be billed on.
	InvoiceContract CommitInvoiceContract `json:"invoice_contract"`
	// The schedule that the customer will be invoiced for this commit.
	InvoiceSchedule SchedulePointInTime `json:"invoice_schedule"`
	// A list of ordered events that impact the balance of a commit. For example, an
	// invoice deduction or a rollover.
	Ledger []CommitLedger `json:"ledger"`
	Name   string         `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority         float64              `json:"priority"`
	RateType         CommitRateType       `json:"rate_type"`
	RolledOverFrom   CommitRolledOverFrom `json:"rolled_over_from"`
	RolloverFraction float64              `json:"rollover_fraction"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string     `json:"uniqueness_key"`
	JSON          commitJSON `json:"-"`
}

// commitJSON contains the JSON metadata for the struct [Commit]
type commitJSON struct {
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

func (r *Commit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitJSON) RawJSON() string {
	return r.raw
}

func (r Commit) ImplementsV1ContractListBalancesResponseData() {}

type CommitProduct struct {
	ID   string            `json:"id,required" format:"uuid"`
	Name string            `json:"name,required"`
	JSON commitProductJSON `json:"-"`
}

// commitProductJSON contains the JSON metadata for the struct [CommitProduct]
type commitProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitProductJSON) RawJSON() string {
	return r.raw
}

type CommitType string

const (
	CommitTypePrepaid  CommitType = "PREPAID"
	CommitTypePostpaid CommitType = "POSTPAID"
)

func (r CommitType) IsKnown() bool {
	switch r {
	case CommitTypePrepaid, CommitTypePostpaid:
		return true
	}
	return false
}

type CommitContract struct {
	ID   string             `json:"id,required" format:"uuid"`
	JSON commitContractJSON `json:"-"`
}

// commitContractJSON contains the JSON metadata for the struct [CommitContract]
type commitContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitContractJSON) RawJSON() string {
	return r.raw
}

// The contract that this commit will be billed on.
type CommitInvoiceContract struct {
	ID   string                    `json:"id,required" format:"uuid"`
	JSON commitInvoiceContractJSON `json:"-"`
}

// commitInvoiceContractJSON contains the JSON metadata for the struct
// [CommitInvoiceContract]
type commitInvoiceContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitInvoiceContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitInvoiceContractJSON) RawJSON() string {
	return r.raw
}

type CommitLedger struct {
	Amount        float64          `json:"amount,required"`
	Timestamp     time.Time        `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerType `json:"type,required"`
	InvoiceID     string           `json:"invoice_id" format:"uuid"`
	NewContractID string           `json:"new_contract_id" format:"uuid"`
	Reason        string           `json:"reason"`
	SegmentID     string           `json:"segment_id" format:"uuid"`
	JSON          commitLedgerJSON `json:"-"`
	union         CommitLedgerUnion
}

// commitLedgerJSON contains the JSON metadata for the struct [CommitLedger]
type commitLedgerJSON struct {
	Amount        apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	InvoiceID     apijson.Field
	NewContractID apijson.Field
	Reason        apijson.Field
	SegmentID     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r commitLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *CommitLedger) UnmarshalJSON(data []byte) (err error) {
	*r = CommitLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CommitLedgerUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [shared.CommitLedgerPrepaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitExpirationLedgerEntry].
func (r CommitLedger) AsUnion() CommitLedgerUnion {
	return r.union
}

// Union satisfied by [shared.CommitLedgerPrepaidCommitSegmentStartLedgerEntry],
// [shared.CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPrepaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPrepaidCommitExpirationLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCanceledLedgerEntry],
// [shared.CommitLedgerPrepaidCommitCreditedLedgerEntry],
// [shared.CommitLedgerPostpaidCommitInitialBalanceLedgerEntry],
// [shared.CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry],
// [shared.CommitLedgerPostpaidCommitRolloverLedgerEntry],
// [shared.CommitLedgerPostpaidCommitTrueupLedgerEntry],
// [shared.CommitLedgerPrepaidCommitManualLedgerEntry],
// [shared.CommitLedgerPostpaidCommitManualLedgerEntry] or
// [shared.CommitLedgerPostpaidCommitExpirationLedgerEntry].
type CommitLedgerUnion interface {
	implementsCommitLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CommitLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitInitialBalanceLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitRolloverLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitTrueupLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPrepaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitManualLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CommitLedgerPostpaidCommitExpirationLedgerEntry{}),
		},
	)
}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntry struct {
	Amount    float64                                              `json:"amount,required"`
	SegmentID string                                               `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                            `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitSegmentStartLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitSegmentStartLedgerEntry]
type commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitSegmentStartLedgerEntryType string

const (
	CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart CommitLedgerPrepaidCommitSegmentStartLedgerEntryType = "PREPAID_COMMIT_SEGMENT_START"
)

func (r CommitLedgerPrepaidCommitSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitSegmentStartLedgerEntryTypePrepaidCommitSegmentStart:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                           `json:"amount,required"`
	InvoiceID string                                                            `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                            `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                         `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitAutomatedInvoiceDeductionLedgerEntryTypePrepaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitRolloverLedgerEntry struct {
	Amount        float64                                          `json:"amount,required"`
	NewContractID string                                           `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerPrepaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          commitLedgerPrepaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitRolloverLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitRolloverLedgerEntry]
type commitLedgerPrepaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitRolloverLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitRolloverLedgerEntryType string

const (
	CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover CommitLedgerPrepaidCommitRolloverLedgerEntryType = "PREPAID_COMMIT_ROLLOVER"
)

func (r CommitLedgerPrepaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitRolloverLedgerEntryTypePrepaidCommitRollover:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitExpirationLedgerEntry struct {
	Amount    float64                                            `json:"amount,required"`
	SegmentID string                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                          `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitExpirationLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPrepaidCommitExpirationLedgerEntry]
type commitLedgerPrepaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitExpirationLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitExpirationLedgerEntryType string

const (
	CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration CommitLedgerPrepaidCommitExpirationLedgerEntryType = "PREPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerPrepaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitExpirationLedgerEntryTypePrepaidCommitExpiration:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitCanceledLedgerEntry struct {
	Amount    float64                                          `json:"amount,required"`
	InvoiceID string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitCanceledLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitCanceledLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCanceledLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCanceledLedgerEntry]
type commitLedgerPrepaidCommitCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCanceledLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitCanceledLedgerEntryType string

const (
	CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled CommitLedgerPrepaidCommitCanceledLedgerEntryType = "PREPAID_COMMIT_CANCELED"
)

func (r CommitLedgerPrepaidCommitCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitCanceledLedgerEntryTypePrepaidCommitCanceled:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitCreditedLedgerEntry struct {
	Amount    float64                                          `json:"amount,required"`
	InvoiceID string                                           `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                           `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                        `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitCreditedLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitCreditedLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitCreditedLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitCreditedLedgerEntry]
type commitLedgerPrepaidCommitCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitCreditedLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitCreditedLedgerEntryType string

const (
	CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited CommitLedgerPrepaidCommitCreditedLedgerEntryType = "PREPAID_COMMIT_CREDITED"
)

func (r CommitLedgerPrepaidCommitCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitCreditedLedgerEntryTypePrepaidCommitCredited:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntry struct {
	Amount    float64                                                 `json:"amount,required"`
	Timestamp time.Time                                               `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON contains the JSON
// metadata for the struct [CommitLedgerPostpaidCommitInitialBalanceLedgerEntry]
type commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitInitialBalanceLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType string

const (
	CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType = "POSTPAID_COMMIT_INITIAL_BALANCE"
)

func (r CommitLedgerPostpaidCommitInitialBalanceLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitInitialBalanceLedgerEntryTypePostpaidCommitInitialBalance:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                            `json:"amount,required"`
	InvoiceID string                                                             `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                             `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                          `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON contains the
// JSON metadata for the struct
// [CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry]
type commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType string

const (
	CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitAutomatedInvoiceDeductionLedgerEntryTypePostpaidCommitAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitRolloverLedgerEntry struct {
	Amount        float64                                           `json:"amount,required"`
	NewContractID string                                            `json:"new_contract_id,required" format:"uuid"`
	SegmentID     string                                            `json:"segment_id,required" format:"uuid"`
	Timestamp     time.Time                                         `json:"timestamp,required" format:"date-time"`
	Type          CommitLedgerPostpaidCommitRolloverLedgerEntryType `json:"type,required"`
	JSON          commitLedgerPostpaidCommitRolloverLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitRolloverLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitRolloverLedgerEntry]
type commitLedgerPostpaidCommitRolloverLedgerEntryJSON struct {
	Amount        apijson.Field
	NewContractID apijson.Field
	SegmentID     apijson.Field
	Timestamp     apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitRolloverLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitRolloverLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitRolloverLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitRolloverLedgerEntryType string

const (
	CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover CommitLedgerPostpaidCommitRolloverLedgerEntryType = "POSTPAID_COMMIT_ROLLOVER"
)

func (r CommitLedgerPostpaidCommitRolloverLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitRolloverLedgerEntryTypePostpaidCommitRollover:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitTrueupLedgerEntry struct {
	Amount    float64                                         `json:"amount,required"`
	InvoiceID string                                          `json:"invoice_id,required" format:"uuid"`
	Timestamp time.Time                                       `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitTrueupLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitTrueupLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitTrueupLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitTrueupLedgerEntry]
type commitLedgerPostpaidCommitTrueupLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitTrueupLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitTrueupLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitTrueupLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitTrueupLedgerEntryType string

const (
	CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup CommitLedgerPostpaidCommitTrueupLedgerEntryType = "POSTPAID_COMMIT_TRUEUP"
)

func (r CommitLedgerPostpaidCommitTrueupLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitTrueupLedgerEntryTypePostpaidCommitTrueup:
		return true
	}
	return false
}

type CommitLedgerPrepaidCommitManualLedgerEntry struct {
	Amount    float64                                        `json:"amount,required"`
	Reason    string                                         `json:"reason,required"`
	Timestamp time.Time                                      `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPrepaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPrepaidCommitManualLedgerEntryJSON `json:"-"`
}

// commitLedgerPrepaidCommitManualLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPrepaidCommitManualLedgerEntry]
type commitLedgerPrepaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPrepaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPrepaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPrepaidCommitManualLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPrepaidCommitManualLedgerEntryType string

const (
	CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual CommitLedgerPrepaidCommitManualLedgerEntryType = "PREPAID_COMMIT_MANUAL"
)

func (r CommitLedgerPrepaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPrepaidCommitManualLedgerEntryTypePrepaidCommitManual:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitManualLedgerEntry struct {
	Amount    float64                                         `json:"amount,required"`
	Reason    string                                          `json:"reason,required"`
	Timestamp time.Time                                       `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitManualLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitManualLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitManualLedgerEntryJSON contains the JSON metadata for
// the struct [CommitLedgerPostpaidCommitManualLedgerEntry]
type commitLedgerPostpaidCommitManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitManualLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitManualLedgerEntryType string

const (
	CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual CommitLedgerPostpaidCommitManualLedgerEntryType = "POSTPAID_COMMIT_MANUAL"
)

func (r CommitLedgerPostpaidCommitManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitManualLedgerEntryTypePostpaidCommitManual:
		return true
	}
	return false
}

type CommitLedgerPostpaidCommitExpirationLedgerEntry struct {
	Amount    float64                                             `json:"amount,required"`
	Timestamp time.Time                                           `json:"timestamp,required" format:"date-time"`
	Type      CommitLedgerPostpaidCommitExpirationLedgerEntryType `json:"type,required"`
	JSON      commitLedgerPostpaidCommitExpirationLedgerEntryJSON `json:"-"`
}

// commitLedgerPostpaidCommitExpirationLedgerEntryJSON contains the JSON metadata
// for the struct [CommitLedgerPostpaidCommitExpirationLedgerEntry]
type commitLedgerPostpaidCommitExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitLedgerPostpaidCommitExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitLedgerPostpaidCommitExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CommitLedgerPostpaidCommitExpirationLedgerEntry) implementsCommitLedger() {}

type CommitLedgerPostpaidCommitExpirationLedgerEntryType string

const (
	CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration CommitLedgerPostpaidCommitExpirationLedgerEntryType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerPostpaidCommitExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CommitLedgerPostpaidCommitExpirationLedgerEntryTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type CommitLedgerType string

const (
	CommitLedgerTypePrepaidCommitSegmentStart               CommitLedgerType = "PREPAID_COMMIT_SEGMENT_START"
	CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction  CommitLedgerType = "PREPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	CommitLedgerTypePrepaidCommitRollover                   CommitLedgerType = "PREPAID_COMMIT_ROLLOVER"
	CommitLedgerTypePrepaidCommitExpiration                 CommitLedgerType = "PREPAID_COMMIT_EXPIRATION"
	CommitLedgerTypePrepaidCommitCanceled                   CommitLedgerType = "PREPAID_COMMIT_CANCELED"
	CommitLedgerTypePrepaidCommitCredited                   CommitLedgerType = "PREPAID_COMMIT_CREDITED"
	CommitLedgerTypePostpaidCommitInitialBalance            CommitLedgerType = "POSTPAID_COMMIT_INITIAL_BALANCE"
	CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction CommitLedgerType = "POSTPAID_COMMIT_AUTOMATED_INVOICE_DEDUCTION"
	CommitLedgerTypePostpaidCommitRollover                  CommitLedgerType = "POSTPAID_COMMIT_ROLLOVER"
	CommitLedgerTypePostpaidCommitTrueup                    CommitLedgerType = "POSTPAID_COMMIT_TRUEUP"
	CommitLedgerTypePrepaidCommitManual                     CommitLedgerType = "PREPAID_COMMIT_MANUAL"
	CommitLedgerTypePostpaidCommitManual                    CommitLedgerType = "POSTPAID_COMMIT_MANUAL"
	CommitLedgerTypePostpaidCommitExpiration                CommitLedgerType = "POSTPAID_COMMIT_EXPIRATION"
)

func (r CommitLedgerType) IsKnown() bool {
	switch r {
	case CommitLedgerTypePrepaidCommitSegmentStart, CommitLedgerTypePrepaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePrepaidCommitRollover, CommitLedgerTypePrepaidCommitExpiration, CommitLedgerTypePrepaidCommitCanceled, CommitLedgerTypePrepaidCommitCredited, CommitLedgerTypePostpaidCommitInitialBalance, CommitLedgerTypePostpaidCommitAutomatedInvoiceDeduction, CommitLedgerTypePostpaidCommitRollover, CommitLedgerTypePostpaidCommitTrueup, CommitLedgerTypePrepaidCommitManual, CommitLedgerTypePostpaidCommitManual, CommitLedgerTypePostpaidCommitExpiration:
		return true
	}
	return false
}

type CommitRateType string

const (
	CommitRateTypeCommitRate CommitRateType = "COMMIT_RATE"
	CommitRateTypeListRate   CommitRateType = "LIST_RATE"
)

func (r CommitRateType) IsKnown() bool {
	switch r {
	case CommitRateTypeCommitRate, CommitRateTypeListRate:
		return true
	}
	return false
}

type CommitRolledOverFrom struct {
	CommitID   string                   `json:"commit_id,required" format:"uuid"`
	ContractID string                   `json:"contract_id,required" format:"uuid"`
	JSON       commitRolledOverFromJSON `json:"-"`
}

// commitRolledOverFromJSON contains the JSON metadata for the struct
// [CommitRolledOverFrom]
type commitRolledOverFromJSON struct {
	CommitID    apijson.Field
	ContractID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommitRolledOverFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commitRolledOverFromJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendments struct {
	Commits                []Commit                                        `json:"commits,required"`
	CreatedAt              time.Time                                       `json:"created_at,required" format:"date-time"`
	CreatedBy              string                                          `json:"created_by,required"`
	Overrides              []Override                                      `json:"overrides,required"`
	ScheduledCharges       []ScheduledCharge                               `json:"scheduled_charges,required"`
	StartingAt             time.Time                                       `json:"starting_at,required" format:"date-time"`
	Transitions            []ContractWithoutAmendmentsTransition           `json:"transitions,required"`
	UsageStatementSchedule ContractWithoutAmendmentsUsageStatementSchedule `json:"usage_statement_schedule,required"`
	Credits                []Credit                                        `json:"credits"`
	// This field's availability is dependent on your client's configuration.
	Discounts           []Discount `json:"discounts"`
	EndingBefore        time.Time  `json:"ending_before" format:"date-time"`
	Name                string     `json:"name"`
	NetPaymentTermsDays float64    `json:"net_payment_terms_days"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// This field's availability is dependent on your client's configuration.
	ProfessionalServices []ProService                               `json:"professional_services"`
	RateCardID           string                                     `json:"rate_card_id" format:"uuid"`
	RecurringCommits     []ContractWithoutAmendmentsRecurringCommit `json:"recurring_commits"`
	RecurringCredits     []ContractWithoutAmendmentsRecurringCredit `json:"recurring_credits"`
	// This field's availability is dependent on your client's configuration.
	ResellerRoyalties []ContractWithoutAmendmentsResellerRoyalty `json:"reseller_royalties"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Determines which scheduled and commit charges to consolidate onto the Contract's
	// usage invoice. The charge's `timestamp` must match the usage invoice's
	// `ending_before` date for consolidation to occur. This field cannot be modified
	// after a Contract has been created. If this field is omitted, charges will appear
	// on a separate invoice from usage charges.
	ScheduledChargesOnUsageInvoices ContractWithoutAmendmentsScheduledChargesOnUsageInvoices `json:"scheduled_charges_on_usage_invoices"`
	ThresholdBillingConfiguration   ContractWithoutAmendmentsThresholdBillingConfiguration   `json:"threshold_billing_configuration"`
	// This field's availability is dependent on your client's configuration.
	TotalContractValue float64                              `json:"total_contract_value"`
	UsageFilter        ContractWithoutAmendmentsUsageFilter `json:"usage_filter"`
	JSON               contractWithoutAmendmentsJSON        `json:"-"`
}

// contractWithoutAmendmentsJSON contains the JSON metadata for the struct
// [ContractWithoutAmendments]
type contractWithoutAmendmentsJSON struct {
	Commits                         apijson.Field
	CreatedAt                       apijson.Field
	CreatedBy                       apijson.Field
	Overrides                       apijson.Field
	ScheduledCharges                apijson.Field
	StartingAt                      apijson.Field
	Transitions                     apijson.Field
	UsageStatementSchedule          apijson.Field
	Credits                         apijson.Field
	Discounts                       apijson.Field
	EndingBefore                    apijson.Field
	Name                            apijson.Field
	NetPaymentTermsDays             apijson.Field
	NetsuiteSalesOrderID            apijson.Field
	ProfessionalServices            apijson.Field
	RateCardID                      apijson.Field
	RecurringCommits                apijson.Field
	RecurringCredits                apijson.Field
	ResellerRoyalties               apijson.Field
	SalesforceOpportunityID         apijson.Field
	ScheduledChargesOnUsageInvoices apijson.Field
	ThresholdBillingConfiguration   apijson.Field
	TotalContractValue              apijson.Field
	UsageFilter                     apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ContractWithoutAmendments) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsTransition struct {
	FromContractID string                                   `json:"from_contract_id,required" format:"uuid"`
	ToContractID   string                                   `json:"to_contract_id,required" format:"uuid"`
	Type           ContractWithoutAmendmentsTransitionsType `json:"type,required"`
	JSON           contractWithoutAmendmentsTransitionJSON  `json:"-"`
}

// contractWithoutAmendmentsTransitionJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsTransition]
type contractWithoutAmendmentsTransitionJSON struct {
	FromContractID apijson.Field
	ToContractID   apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsTransitionJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsTransitionsType string

const (
	ContractWithoutAmendmentsTransitionsTypeSupersede ContractWithoutAmendmentsTransitionsType = "SUPERSEDE"
	ContractWithoutAmendmentsTransitionsTypeRenewal   ContractWithoutAmendmentsTransitionsType = "RENEWAL"
)

func (r ContractWithoutAmendmentsTransitionsType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsTransitionsTypeSupersede, ContractWithoutAmendmentsTransitionsTypeRenewal:
		return true
	}
	return false
}

type ContractWithoutAmendmentsUsageStatementSchedule struct {
	// Contract usage statements follow a selected cadence based on this date.
	BillingAnchorDate time.Time                                                `json:"billing_anchor_date,required" format:"date-time"`
	Frequency         ContractWithoutAmendmentsUsageStatementScheduleFrequency `json:"frequency,required"`
	JSON              contractWithoutAmendmentsUsageStatementScheduleJSON      `json:"-"`
}

// contractWithoutAmendmentsUsageStatementScheduleJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsUsageStatementSchedule]
type contractWithoutAmendmentsUsageStatementScheduleJSON struct {
	BillingAnchorDate apijson.Field
	Frequency         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageStatementSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageStatementScheduleJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageStatementScheduleFrequency string

const (
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly   ContractWithoutAmendmentsUsageStatementScheduleFrequency = "MONTHLY"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly ContractWithoutAmendmentsUsageStatementScheduleFrequency = "QUARTERLY"
	ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual    ContractWithoutAmendmentsUsageStatementScheduleFrequency = "ANNUAL"
)

func (r ContractWithoutAmendmentsUsageStatementScheduleFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsUsageStatementScheduleFrequencyMonthly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyQuarterly, ContractWithoutAmendmentsUsageStatementScheduleFrequencyAnnual:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCommitsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCommitsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCommitsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractWithoutAmendmentsRecurringCommitsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCommitsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// The amount the customer should be billed for the commit. Not required.
	InvoiceAmount ContractWithoutAmendmentsRecurringCommitsInvoiceAmount `json:"invoice_amount"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractWithoutAmendmentsRecurringCommitsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                      `json:"rollover_fraction"`
	JSON             contractWithoutAmendmentsRecurringCommitJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsRecurringCommit]
type contractWithoutAmendmentsRecurringCommitJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	InvoiceAmount         apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCommitsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                   `json:"quantity,required"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	JSON         contractWithoutAmendmentsRecurringCommitsAccessAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsAccessAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCommitsAccessAmount]
type contractWithoutAmendmentsRecurringCommitsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCommitsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit `json:"unit"`
	JSON  contractWithoutAmendmentsRecurringCommitsCommitDurationJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsCommitDurationJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsRecurringCommitsCommitDuration]
type contractWithoutAmendmentsRecurringCommitsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit string

const (
	ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit = "PERIODS"
)

func (r ContractWithoutAmendmentsRecurringCommitsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON contractWithoutAmendmentsRecurringCommitsProductJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsProductJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCommitsProduct]
type contractWithoutAmendmentsRecurringCommitsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractWithoutAmendmentsRecurringCommitsRateType string

const (
	ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate ContractWithoutAmendmentsRecurringCommitsRateType = "COMMIT_RATE"
	ContractWithoutAmendmentsRecurringCommitsRateTypeListRate   ContractWithoutAmendmentsRecurringCommitsRateType = "LIST_RATE"
)

func (r ContractWithoutAmendmentsRecurringCommitsRateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsRateTypeCommitRate, ContractWithoutAmendmentsRecurringCommitsRateTypeListRate:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCommitsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON contractWithoutAmendmentsRecurringCommitsContractJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsContractJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCommitsContract]
type contractWithoutAmendmentsRecurringCommitsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsContractJSON) RawJSON() string {
	return r.raw
}

// The amount the customer should be billed for the commit. Not required.
type ContractWithoutAmendmentsRecurringCommitsInvoiceAmount struct {
	CreditTypeID string                                                     `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                    `json:"quantity,required"`
	UnitPrice    float64                                                    `json:"unit_price,required"`
	JSON         contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCommitsInvoiceAmount]
type contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCommitsInvoiceAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCommitsInvoiceAmountJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractWithoutAmendmentsRecurringCommitsProration string

const (
	ContractWithoutAmendmentsRecurringCommitsProrationNone         ContractWithoutAmendmentsRecurringCommitsProration = "NONE"
	ContractWithoutAmendmentsRecurringCommitsProrationFirst        ContractWithoutAmendmentsRecurringCommitsProration = "FIRST"
	ContractWithoutAmendmentsRecurringCommitsProrationLast         ContractWithoutAmendmentsRecurringCommitsProration = "LAST"
	ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast ContractWithoutAmendmentsRecurringCommitsProration = "FIRST_AND_LAST"
)

func (r ContractWithoutAmendmentsRecurringCommitsProration) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsProrationNone, ContractWithoutAmendmentsRecurringCommitsProrationFirst, ContractWithoutAmendmentsRecurringCommitsProrationLast, ContractWithoutAmendmentsRecurringCommitsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency string

const (
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly   ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "MONTHLY"
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "QUARTERLY"
	ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual    ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency = "ANNUAL"
)

func (r ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyMonthly, ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyQuarterly, ContractWithoutAmendmentsRecurringCommitsRecurrenceFrequencyAnnual:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCredit struct {
	ID string `json:"id,required" format:"uuid"`
	// The amount of commit to grant.
	AccessAmount ContractWithoutAmendmentsRecurringCreditsAccessAmount `json:"access_amount,required"`
	// The amount of time the created commits will be valid for
	CommitDuration ContractWithoutAmendmentsRecurringCreditsCommitDuration `json:"commit_duration,required"`
	// Will be passed down to the individual commits
	Priority float64                                          `json:"priority,required"`
	Product  ContractWithoutAmendmentsRecurringCreditsProduct `json:"product,required"`
	// Whether the created commits will use the commit rate or list rate
	RateType ContractWithoutAmendmentsRecurringCreditsRateType `json:"rate_type,required"`
	// Determines the start time for the first commit
	StartingAt time.Time `json:"starting_at,required" format:"date-time"`
	// Will be passed down to the individual commits
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Will be passed down to the individual commits
	ApplicableProductTags []string                                          `json:"applicable_product_tags"`
	Contract              ContractWithoutAmendmentsRecurringCreditsContract `json:"contract"`
	// Will be passed down to the individual commits
	Description string `json:"description"`
	// Determines when the contract will stop creating recurring commits. Optional
	EndingBefore time.Time `json:"ending_before" format:"date-time"`
	// Displayed on invoices. Will be passed through to the individual commits
	Name string `json:"name"`
	// Will be passed down to the individual commits
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// Determines whether the first and last commit will be prorated. If not provided,
	// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
	Proration ContractWithoutAmendmentsRecurringCreditsProration `json:"proration"`
	// The frequency at which the recurring commits will be created. If not provided: -
	// The commits will be created on the usage invoice frequency. If provided: - The
	// period defined in the duration will correspond to this frequency. - Commits will
	// be created aligned with the recurring commit's start_date rather than the usage
	// invoice dates.
	RecurrenceFrequency ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency `json:"recurrence_frequency"`
	// Will be passed down to the individual commits. This controls how much of an
	// individual unexpired commit will roll over upon contract transition. Must be
	// between 0 and 1.
	RolloverFraction float64                                      `json:"rollover_fraction"`
	JSON             contractWithoutAmendmentsRecurringCreditJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsRecurringCredit]
type contractWithoutAmendmentsRecurringCreditJSON struct {
	ID                    apijson.Field
	AccessAmount          apijson.Field
	CommitDuration        apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	RateType              apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Contract              apijson.Field
	Description           apijson.Field
	EndingBefore          apijson.Field
	Name                  apijson.Field
	NetsuiteSalesOrderID  apijson.Field
	Proration             apijson.Field
	RecurrenceFrequency   apijson.Field
	RolloverFraction      apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCredit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditJSON) RawJSON() string {
	return r.raw
}

// The amount of commit to grant.
type ContractWithoutAmendmentsRecurringCreditsAccessAmount struct {
	CreditTypeID string                                                    `json:"credit_type_id,required" format:"uuid"`
	Quantity     float64                                                   `json:"quantity,required"`
	UnitPrice    float64                                                   `json:"unit_price,required"`
	JSON         contractWithoutAmendmentsRecurringCreditsAccessAmountJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsAccessAmountJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsRecurringCreditsAccessAmount]
type contractWithoutAmendmentsRecurringCreditsAccessAmountJSON struct {
	CreditTypeID apijson.Field
	Quantity     apijson.Field
	UnitPrice    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsAccessAmount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsAccessAmountJSON) RawJSON() string {
	return r.raw
}

// The amount of time the created commits will be valid for
type ContractWithoutAmendmentsRecurringCreditsCommitDuration struct {
	Value float64                                                     `json:"value,required"`
	Unit  ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit `json:"unit"`
	JSON  contractWithoutAmendmentsRecurringCreditsCommitDurationJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsCommitDurationJSON contains the JSON
// metadata for the struct
// [ContractWithoutAmendmentsRecurringCreditsCommitDuration]
type contractWithoutAmendmentsRecurringCreditsCommitDurationJSON struct {
	Value       apijson.Field
	Unit        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsCommitDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsCommitDurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit string

const (
	ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit = "PERIODS"
)

func (r ContractWithoutAmendmentsRecurringCreditsCommitDurationUnit) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsCommitDurationUnitPeriods:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsProduct struct {
	ID   string                                               `json:"id,required" format:"uuid"`
	Name string                                               `json:"name,required"`
	JSON contractWithoutAmendmentsRecurringCreditsProductJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsProductJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCreditsProduct]
type contractWithoutAmendmentsRecurringCreditsProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsProductJSON) RawJSON() string {
	return r.raw
}

// Whether the created commits will use the commit rate or list rate
type ContractWithoutAmendmentsRecurringCreditsRateType string

const (
	ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate ContractWithoutAmendmentsRecurringCreditsRateType = "COMMIT_RATE"
	ContractWithoutAmendmentsRecurringCreditsRateTypeListRate   ContractWithoutAmendmentsRecurringCreditsRateType = "LIST_RATE"
)

func (r ContractWithoutAmendmentsRecurringCreditsRateType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsRateTypeCommitRate, ContractWithoutAmendmentsRecurringCreditsRateTypeListRate:
		return true
	}
	return false
}

type ContractWithoutAmendmentsRecurringCreditsContract struct {
	ID   string                                                `json:"id,required" format:"uuid"`
	JSON contractWithoutAmendmentsRecurringCreditsContractJSON `json:"-"`
}

// contractWithoutAmendmentsRecurringCreditsContractJSON contains the JSON metadata
// for the struct [ContractWithoutAmendmentsRecurringCreditsContract]
type contractWithoutAmendmentsRecurringCreditsContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsRecurringCreditsContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsRecurringCreditsContractJSON) RawJSON() string {
	return r.raw
}

// Determines whether the first and last commit will be prorated. If not provided,
// the default is FIRST_AND_LAST (i.e. prorate both the first and last commits).
type ContractWithoutAmendmentsRecurringCreditsProration string

const (
	ContractWithoutAmendmentsRecurringCreditsProrationNone         ContractWithoutAmendmentsRecurringCreditsProration = "NONE"
	ContractWithoutAmendmentsRecurringCreditsProrationFirst        ContractWithoutAmendmentsRecurringCreditsProration = "FIRST"
	ContractWithoutAmendmentsRecurringCreditsProrationLast         ContractWithoutAmendmentsRecurringCreditsProration = "LAST"
	ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast ContractWithoutAmendmentsRecurringCreditsProration = "FIRST_AND_LAST"
)

func (r ContractWithoutAmendmentsRecurringCreditsProration) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsProrationNone, ContractWithoutAmendmentsRecurringCreditsProrationFirst, ContractWithoutAmendmentsRecurringCreditsProrationLast, ContractWithoutAmendmentsRecurringCreditsProrationFirstAndLast:
		return true
	}
	return false
}

// The frequency at which the recurring commits will be created. If not provided: -
// The commits will be created on the usage invoice frequency. If provided: - The
// period defined in the duration will correspond to this frequency. - Commits will
// be created aligned with the recurring commit's start_date rather than the usage
// invoice dates.
type ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency string

const (
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly   ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "MONTHLY"
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "QUARTERLY"
	ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual    ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency = "ANNUAL"
)

func (r ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequency) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyMonthly, ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyQuarterly, ContractWithoutAmendmentsRecurringCreditsRecurrenceFrequencyAnnual:
		return true
	}
	return false
}

type ContractWithoutAmendmentsResellerRoyalty struct {
	Fraction              float64                                                `json:"fraction,required"`
	NetsuiteResellerID    string                                                 `json:"netsuite_reseller_id,required"`
	ResellerType          ContractWithoutAmendmentsResellerRoyaltiesResellerType `json:"reseller_type,required"`
	StartingAt            time.Time                                              `json:"starting_at,required" format:"date-time"`
	ApplicableProductIDs  []string                                               `json:"applicable_product_ids"`
	ApplicableProductTags []string                                               `json:"applicable_product_tags"`
	AwsAccountNumber      string                                                 `json:"aws_account_number"`
	AwsOfferID            string                                                 `json:"aws_offer_id"`
	AwsPayerReferenceID   string                                                 `json:"aws_payer_reference_id"`
	EndingBefore          time.Time                                              `json:"ending_before" format:"date-time"`
	GcpAccountID          string                                                 `json:"gcp_account_id"`
	GcpOfferID            string                                                 `json:"gcp_offer_id"`
	ResellerContractValue float64                                                `json:"reseller_contract_value"`
	JSON                  contractWithoutAmendmentsResellerRoyaltyJSON           `json:"-"`
}

// contractWithoutAmendmentsResellerRoyaltyJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsResellerRoyalty]
type contractWithoutAmendmentsResellerRoyaltyJSON struct {
	Fraction              apijson.Field
	NetsuiteResellerID    apijson.Field
	ResellerType          apijson.Field
	StartingAt            apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	AwsAccountNumber      apijson.Field
	AwsOfferID            apijson.Field
	AwsPayerReferenceID   apijson.Field
	EndingBefore          apijson.Field
	GcpAccountID          apijson.Field
	GcpOfferID            apijson.Field
	ResellerContractValue apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsResellerRoyalty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsResellerRoyaltyJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsResellerRoyaltiesResellerType string

const (
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws           ContractWithoutAmendmentsResellerRoyaltiesResellerType = "AWS"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService ContractWithoutAmendmentsResellerRoyaltiesResellerType = "AWS_PRO_SERVICE"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp           ContractWithoutAmendmentsResellerRoyaltiesResellerType = "GCP"
	ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService ContractWithoutAmendmentsResellerRoyaltiesResellerType = "GCP_PRO_SERVICE"
)

func (r ContractWithoutAmendmentsResellerRoyaltiesResellerType) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAws, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeAwsProService, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcp, ContractWithoutAmendmentsResellerRoyaltiesResellerTypeGcpProService:
		return true
	}
	return false
}

// Determines which scheduled and commit charges to consolidate onto the Contract's
// usage invoice. The charge's `timestamp` must match the usage invoice's
// `ending_before` date for consolidation to occur. This field cannot be modified
// after a Contract has been created. If this field is omitted, charges will appear
// on a separate invoice from usage charges.
type ContractWithoutAmendmentsScheduledChargesOnUsageInvoices string

const (
	ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll ContractWithoutAmendmentsScheduledChargesOnUsageInvoices = "ALL"
)

func (r ContractWithoutAmendmentsScheduledChargesOnUsageInvoices) IsKnown() bool {
	switch r {
	case ContractWithoutAmendmentsScheduledChargesOnUsageInvoicesAll:
		return true
	}
	return false
}

type ContractWithoutAmendmentsThresholdBillingConfiguration struct {
	Commit ContractWithoutAmendmentsThresholdBillingConfigurationCommit `json:"commit,required"`
	// When set to false, the contract will not be evaluated against the
	// threshold_amount. Toggling to true will result an immediate evaluation,
	// regardless of prior state
	IsEnabled bool `json:"is_enabled,required"`
	// Specify the threshold amount for the contract. Each time the contract's usage
	// hits this amount, a threshold charge will be initiated.
	ThresholdAmount float64                                                    `json:"threshold_amount,required"`
	JSON            contractWithoutAmendmentsThresholdBillingConfigurationJSON `json:"-"`
}

// contractWithoutAmendmentsThresholdBillingConfigurationJSON contains the JSON
// metadata for the struct [ContractWithoutAmendmentsThresholdBillingConfiguration]
type contractWithoutAmendmentsThresholdBillingConfigurationJSON struct {
	Commit          apijson.Field
	IsEnabled       apijson.Field
	ThresholdAmount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsThresholdBillingConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsThresholdBillingConfigurationJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsThresholdBillingConfigurationCommit struct {
	ProductID string `json:"product_id,required"`
	// Which products the threshold commit applies to. If both applicable_product_ids
	// and applicable_product_tags are not provided, the commit applies to all
	// products.
	ApplicableProductIDs []string `json:"applicable_product_ids" format:"uuid"`
	// Which tags the threshold commit applies to. If both applicable_product_ids and
	// applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags"`
	Description           string   `json:"description"`
	// Specify the name of the line item for the threshold charge. If left blank, it
	// will default to the commit product name.
	Name string                                                           `json:"name"`
	JSON contractWithoutAmendmentsThresholdBillingConfigurationCommitJSON `json:"-"`
}

// contractWithoutAmendmentsThresholdBillingConfigurationCommitJSON contains the
// JSON metadata for the struct
// [ContractWithoutAmendmentsThresholdBillingConfigurationCommit]
type contractWithoutAmendmentsThresholdBillingConfigurationCommitJSON struct {
	ProductID             apijson.Field
	ApplicableProductIDs  apijson.Field
	ApplicableProductTags apijson.Field
	Description           apijson.Field
	Name                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsThresholdBillingConfigurationCommit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsThresholdBillingConfigurationCommitJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageFilter struct {
	Current BaseUsageFilter                              `json:"current,required,nullable"`
	Initial BaseUsageFilter                              `json:"initial,required"`
	Updates []ContractWithoutAmendmentsUsageFilterUpdate `json:"updates,required"`
	JSON    contractWithoutAmendmentsUsageFilterJSON     `json:"-"`
}

// contractWithoutAmendmentsUsageFilterJSON contains the JSON metadata for the
// struct [ContractWithoutAmendmentsUsageFilter]
type contractWithoutAmendmentsUsageFilterJSON struct {
	Current     apijson.Field
	Initial     apijson.Field
	Updates     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterJSON) RawJSON() string {
	return r.raw
}

type ContractWithoutAmendmentsUsageFilterUpdate struct {
	GroupKey    string                                         `json:"group_key,required"`
	GroupValues []string                                       `json:"group_values,required"`
	StartingAt  time.Time                                      `json:"starting_at,required" format:"date-time"`
	JSON        contractWithoutAmendmentsUsageFilterUpdateJSON `json:"-"`
}

// contractWithoutAmendmentsUsageFilterUpdateJSON contains the JSON metadata for
// the struct [ContractWithoutAmendmentsUsageFilterUpdate]
type contractWithoutAmendmentsUsageFilterUpdateJSON struct {
	GroupKey    apijson.Field
	GroupValues apijson.Field
	StartingAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContractWithoutAmendmentsUsageFilterUpdate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contractWithoutAmendmentsUsageFilterUpdateJSON) RawJSON() string {
	return r.raw
}

type Credit struct {
	ID      string        `json:"id,required" format:"uuid"`
	Product CreditProduct `json:"product,required"`
	Type    CreditType    `json:"type,required"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        ScheduleDuration `json:"access_schedule"`
	ApplicableContractIDs []string         `json:"applicable_contract_ids" format:"uuid"`
	ApplicableProductIDs  []string         `json:"applicable_product_ids" format:"uuid"`
	ApplicableProductTags []string         `json:"applicable_product_tags"`
	// The current balance of the credit or commit. This balance reflects the amount of
	// credit or commit that the customer has access to use at this moment - thus,
	// expired and upcoming credit or commit segments contribute 0 to the balance. The
	// balance will match the sum of all ledger entries with the exception of the case
	// where the sum of negative manual ledger entries exceeds the positive amount
	// remaining on the credit or commit - in that case, the balance will be 0. All
	// manual ledger entries associated with active credit or commit segments are
	// included in the balance, including future-dated manual ledger entries.
	Balance      float64           `json:"balance"`
	Contract     CreditContract    `json:"contract"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// A list of ordered events that impact the balance of a credit. For example, an
	// invoice deduction or an expiration.
	Ledger []CreditLedger `json:"ledger"`
	Name   string         `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string `json:"netsuite_sales_order_id"`
	// If multiple credits or commits are applicable, the one with the lower priority
	// will apply first.
	Priority float64        `json:"priority"`
	RateType CreditRateType `json:"rate_type"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID string `json:"salesforce_opportunity_id"`
	// Prevents the creation of duplicates. If a request to create a commit or credit
	// is made with a uniqueness key that was previously used to create a commit or
	// credit, a new record will not be created and the request will fail with a 409
	// error.
	UniquenessKey string     `json:"uniqueness_key"`
	JSON          creditJSON `json:"-"`
}

// creditJSON contains the JSON metadata for the struct [Credit]
type creditJSON struct {
	ID                      apijson.Field
	Product                 apijson.Field
	Type                    apijson.Field
	AccessSchedule          apijson.Field
	ApplicableContractIDs   apijson.Field
	ApplicableProductIDs    apijson.Field
	ApplicableProductTags   apijson.Field
	Balance                 apijson.Field
	Contract                apijson.Field
	CustomFields            apijson.Field
	Description             apijson.Field
	Ledger                  apijson.Field
	Name                    apijson.Field
	NetsuiteSalesOrderID    apijson.Field
	Priority                apijson.Field
	RateType                apijson.Field
	SalesforceOpportunityID apijson.Field
	UniquenessKey           apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Credit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditJSON) RawJSON() string {
	return r.raw
}

func (r Credit) ImplementsV1ContractListBalancesResponseData() {}

type CreditProduct struct {
	ID   string            `json:"id,required" format:"uuid"`
	Name string            `json:"name,required"`
	JSON creditProductJSON `json:"-"`
}

// creditProductJSON contains the JSON metadata for the struct [CreditProduct]
type creditProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditProductJSON) RawJSON() string {
	return r.raw
}

type CreditType string

const (
	CreditTypeCredit CreditType = "CREDIT"
)

func (r CreditType) IsKnown() bool {
	switch r {
	case CreditTypeCredit:
		return true
	}
	return false
}

type CreditContract struct {
	ID   string             `json:"id,required" format:"uuid"`
	JSON creditContractJSON `json:"-"`
}

// creditContractJSON contains the JSON metadata for the struct [CreditContract]
type creditContractJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditContract) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditContractJSON) RawJSON() string {
	return r.raw
}

type CreditLedger struct {
	Amount    float64          `json:"amount,required"`
	Timestamp time.Time        `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerType `json:"type,required"`
	InvoiceID string           `json:"invoice_id" format:"uuid"`
	Reason    string           `json:"reason"`
	SegmentID string           `json:"segment_id" format:"uuid"`
	JSON      creditLedgerJSON `json:"-"`
	union     CreditLedgerUnion
}

// creditLedgerJSON contains the JSON metadata for the struct [CreditLedger]
type creditLedgerJSON struct {
	Amount      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	InvoiceID   apijson.Field
	Reason      apijson.Field
	SegmentID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r creditLedgerJSON) RawJSON() string {
	return r.raw
}

func (r *CreditLedger) UnmarshalJSON(data []byte) (err error) {
	*r = CreditLedger{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CreditLedgerUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [shared.CreditLedgerCreditSegmentStartLedgerEntry],
// [shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [shared.CreditLedgerCreditExpirationLedgerEntry],
// [shared.CreditLedgerCreditCanceledLedgerEntry],
// [shared.CreditLedgerCreditCreditedLedgerEntry],
// [shared.CreditLedgerCreditManualLedgerEntry].
func (r CreditLedger) AsUnion() CreditLedgerUnion {
	return r.union
}

// Union satisfied by [shared.CreditLedgerCreditSegmentStartLedgerEntry],
// [shared.CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry],
// [shared.CreditLedgerCreditExpirationLedgerEntry],
// [shared.CreditLedgerCreditCanceledLedgerEntry],
// [shared.CreditLedgerCreditCreditedLedgerEntry] or
// [shared.CreditLedgerCreditManualLedgerEntry].
type CreditLedgerUnion interface {
	implementsCreditLedger()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CreditLedgerUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditSegmentStartLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditExpirationLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditCanceledLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditCreditedLedgerEntry{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CreditLedgerCreditManualLedgerEntry{}),
		},
	)
}

type CreditLedgerCreditSegmentStartLedgerEntry struct {
	Amount    float64                                       `json:"amount,required"`
	SegmentID string                                        `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                     `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditSegmentStartLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditSegmentStartLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditSegmentStartLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditSegmentStartLedgerEntry]
type creditLedgerCreditSegmentStartLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditSegmentStartLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditSegmentStartLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditSegmentStartLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditSegmentStartLedgerEntryType string

const (
	CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart CreditLedgerCreditSegmentStartLedgerEntryType = "CREDIT_SEGMENT_START"
)

func (r CreditLedgerCreditSegmentStartLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditSegmentStartLedgerEntryTypeCreditSegmentStart:
		return true
	}
	return false
}

type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry struct {
	Amount    float64                                                    `json:"amount,required"`
	InvoiceID string                                                     `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                                     `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                                  `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON contains the JSON
// metadata for the struct [CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry]
type creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditAutomatedInvoiceDeductionLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType string

const (
	CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (r CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditAutomatedInvoiceDeductionLedgerEntryTypeCreditAutomatedInvoiceDeduction:
		return true
	}
	return false
}

type CreditLedgerCreditExpirationLedgerEntry struct {
	Amount    float64                                     `json:"amount,required"`
	SegmentID string                                      `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                   `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditExpirationLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditExpirationLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditExpirationLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditExpirationLedgerEntry]
type creditLedgerCreditExpirationLedgerEntryJSON struct {
	Amount      apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditExpirationLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditExpirationLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditExpirationLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditExpirationLedgerEntryType string

const (
	CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration CreditLedgerCreditExpirationLedgerEntryType = "CREDIT_EXPIRATION"
)

func (r CreditLedgerCreditExpirationLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditExpirationLedgerEntryTypeCreditExpiration:
		return true
	}
	return false
}

type CreditLedgerCreditCanceledLedgerEntry struct {
	Amount    float64                                   `json:"amount,required"`
	InvoiceID string                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                 `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditCanceledLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditCanceledLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditCanceledLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditCanceledLedgerEntry]
type creditLedgerCreditCanceledLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditCanceledLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditCanceledLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditCanceledLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditCanceledLedgerEntryType string

const (
	CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled CreditLedgerCreditCanceledLedgerEntryType = "CREDIT_CANCELED"
)

func (r CreditLedgerCreditCanceledLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditCanceledLedgerEntryTypeCreditCanceled:
		return true
	}
	return false
}

type CreditLedgerCreditCreditedLedgerEntry struct {
	Amount    float64                                   `json:"amount,required"`
	InvoiceID string                                    `json:"invoice_id,required" format:"uuid"`
	SegmentID string                                    `json:"segment_id,required" format:"uuid"`
	Timestamp time.Time                                 `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditCreditedLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditCreditedLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditCreditedLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditCreditedLedgerEntry]
type creditLedgerCreditCreditedLedgerEntryJSON struct {
	Amount      apijson.Field
	InvoiceID   apijson.Field
	SegmentID   apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditCreditedLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditCreditedLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditCreditedLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditCreditedLedgerEntryType string

const (
	CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited CreditLedgerCreditCreditedLedgerEntryType = "CREDIT_CREDITED"
)

func (r CreditLedgerCreditCreditedLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditCreditedLedgerEntryTypeCreditCredited:
		return true
	}
	return false
}

type CreditLedgerCreditManualLedgerEntry struct {
	Amount    float64                                 `json:"amount,required"`
	Reason    string                                  `json:"reason,required"`
	Timestamp time.Time                               `json:"timestamp,required" format:"date-time"`
	Type      CreditLedgerCreditManualLedgerEntryType `json:"type,required"`
	JSON      creditLedgerCreditManualLedgerEntryJSON `json:"-"`
}

// creditLedgerCreditManualLedgerEntryJSON contains the JSON metadata for the
// struct [CreditLedgerCreditManualLedgerEntry]
type creditLedgerCreditManualLedgerEntryJSON struct {
	Amount      apijson.Field
	Reason      apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditLedgerCreditManualLedgerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditLedgerCreditManualLedgerEntryJSON) RawJSON() string {
	return r.raw
}

func (r CreditLedgerCreditManualLedgerEntry) implementsCreditLedger() {}

type CreditLedgerCreditManualLedgerEntryType string

const (
	CreditLedgerCreditManualLedgerEntryTypeCreditManual CreditLedgerCreditManualLedgerEntryType = "CREDIT_MANUAL"
)

func (r CreditLedgerCreditManualLedgerEntryType) IsKnown() bool {
	switch r {
	case CreditLedgerCreditManualLedgerEntryTypeCreditManual:
		return true
	}
	return false
}

type CreditLedgerType string

const (
	CreditLedgerTypeCreditSegmentStart              CreditLedgerType = "CREDIT_SEGMENT_START"
	CreditLedgerTypeCreditAutomatedInvoiceDeduction CreditLedgerType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
	CreditLedgerTypeCreditExpiration                CreditLedgerType = "CREDIT_EXPIRATION"
	CreditLedgerTypeCreditCanceled                  CreditLedgerType = "CREDIT_CANCELED"
	CreditLedgerTypeCreditCredited                  CreditLedgerType = "CREDIT_CREDITED"
	CreditLedgerTypeCreditManual                    CreditLedgerType = "CREDIT_MANUAL"
)

func (r CreditLedgerType) IsKnown() bool {
	switch r {
	case CreditLedgerTypeCreditSegmentStart, CreditLedgerTypeCreditAutomatedInvoiceDeduction, CreditLedgerTypeCreditExpiration, CreditLedgerTypeCreditCanceled, CreditLedgerTypeCreditCredited, CreditLedgerTypeCreditManual:
		return true
	}
	return false
}

type CreditRateType string

const (
	CreditRateTypeCommitRate CreditRateType = "COMMIT_RATE"
	CreditRateTypeListRate   CreditRateType = "LIST_RATE"
)

func (r CreditRateType) IsKnown() bool {
	switch r {
	case CreditRateTypeCommitRate, CreditRateTypeListRate:
		return true
	}
	return false
}

type CreditTypeData struct {
	ID   string             `json:"id,required" format:"uuid"`
	Name string             `json:"name,required"`
	JSON creditTypeDataJSON `json:"-"`
}

// creditTypeDataJSON contains the JSON metadata for the struct [CreditTypeData]
type creditTypeDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreditTypeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r creditTypeDataJSON) RawJSON() string {
	return r.raw
}

type Discount struct {
	ID           string              `json:"id,required" format:"uuid"`
	Product      DiscountProduct     `json:"product,required"`
	Schedule     SchedulePointInTime `json:"schedule,required"`
	CustomFields map[string]string   `json:"custom_fields"`
	Name         string              `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string       `json:"netsuite_sales_order_id"`
	JSON                 discountJSON `json:"-"`
}

// discountJSON contains the JSON metadata for the struct [Discount]
type discountJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	CustomFields         apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Discount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountJSON) RawJSON() string {
	return r.raw
}

type DiscountProduct struct {
	ID   string              `json:"id,required" format:"uuid"`
	Name string              `json:"name,required"`
	JSON discountProductJSON `json:"-"`
}

// discountProductJSON contains the JSON metadata for the struct [DiscountProduct]
type discountProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DiscountProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r discountProductJSON) RawJSON() string {
	return r.raw
}

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilter struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues []string `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues []string            `json:"not_in_values"`
	JSON        eventTypeFilterJSON `json:"-"`
}

// eventTypeFilterJSON contains the JSON metadata for the struct [EventTypeFilter]
type eventTypeFilterJSON struct {
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventTypeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventTypeFilterJSON) RawJSON() string {
	return r.raw
}

// An optional filtering rule to match the 'event_type' property of an event.
type EventTypeFilterParam struct {
	// A list of event types that are explicitly included in the billable metric. If
	// specified, only events of these types will match the billable metric. Must be
	// non-empty if present.
	InValues param.Field[[]string] `json:"in_values"`
	// A list of event types that are explicitly excluded from the billable metric. If
	// specified, events of these types will not match the billable metric. Must be
	// non-empty if present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r EventTypeFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ID struct {
	ID   string `json:"id,required" format:"uuid"`
	JSON idJSON `json:"-"`
}

// idJSON contains the JSON metadata for the struct [ID]
type idJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idJSON) RawJSON() string {
	return r.raw
}

type IDParam struct {
	ID param.Field[string] `json:"id,required" format:"uuid"`
}

func (r IDParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Override struct {
	ID                    string         `json:"id,required" format:"uuid"`
	StartingAt            time.Time      `json:"starting_at,required" format:"date-time"`
	ApplicableProductTags []string       `json:"applicable_product_tags"`
	CreditType            CreditTypeData `json:"credit_type"`
	EndingBefore          time.Time      `json:"ending_before" format:"date-time"`
	Entitled              bool           `json:"entitled"`
	IsCommitSpecific      bool           `json:"is_commit_specific"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated         bool                        `json:"is_prorated"`
	Multiplier         float64                     `json:"multiplier"`
	OverrideSpecifiers []OverrideOverrideSpecifier `json:"override_specifiers"`
	OverrideTiers      []OverrideOverrideTier      `json:"override_tiers"`
	OverwriteRate      OverrideOverwriteRate       `json:"overwrite_rate"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price    float64         `json:"price"`
	Priority float64         `json:"priority"`
	Product  OverrideProduct `json:"product"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64          `json:"quantity"`
	RateType OverrideRateType `json:"rate_type"`
	Target   OverrideTarget   `json:"target"`
	// Only set for TIERED rate_type.
	Tiers []Tier       `json:"tiers"`
	Type  OverrideType `json:"type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	Value map[string]interface{} `json:"value"`
	JSON  overrideJSON           `json:"-"`
}

// overrideJSON contains the JSON metadata for the struct [Override]
type overrideJSON struct {
	ID                    apijson.Field
	StartingAt            apijson.Field
	ApplicableProductTags apijson.Field
	CreditType            apijson.Field
	EndingBefore          apijson.Field
	Entitled              apijson.Field
	IsCommitSpecific      apijson.Field
	IsProrated            apijson.Field
	Multiplier            apijson.Field
	OverrideSpecifiers    apijson.Field
	OverrideTiers         apijson.Field
	OverwriteRate         apijson.Field
	Price                 apijson.Field
	Priority              apijson.Field
	Product               apijson.Field
	Quantity              apijson.Field
	RateType              apijson.Field
	Target                apijson.Field
	Tiers                 apijson.Field
	Type                  apijson.Field
	Value                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideJSON) RawJSON() string {
	return r.raw
}

type OverrideOverrideSpecifier struct {
	BillingFrequency        OverrideOverrideSpecifiersBillingFrequency `json:"billing_frequency"`
	CommitIDs               []string                                   `json:"commit_ids"`
	PresentationGroupValues map[string]string                          `json:"presentation_group_values"`
	PricingGroupValues      map[string]string                          `json:"pricing_group_values"`
	ProductID               string                                     `json:"product_id" format:"uuid"`
	ProductTags             []string                                   `json:"product_tags"`
	RecurringCommitIDs      []string                                   `json:"recurring_commit_ids"`
	RecurringCreditIDs      []string                                   `json:"recurring_credit_ids"`
	JSON                    overrideOverrideSpecifierJSON              `json:"-"`
}

// overrideOverrideSpecifierJSON contains the JSON metadata for the struct
// [OverrideOverrideSpecifier]
type overrideOverrideSpecifierJSON struct {
	BillingFrequency        apijson.Field
	CommitIDs               apijson.Field
	PresentationGroupValues apijson.Field
	PricingGroupValues      apijson.Field
	ProductID               apijson.Field
	ProductTags             apijson.Field
	RecurringCommitIDs      apijson.Field
	RecurringCreditIDs      apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OverrideOverrideSpecifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverrideSpecifierJSON) RawJSON() string {
	return r.raw
}

type OverrideOverrideSpecifiersBillingFrequency string

const (
	OverrideOverrideSpecifiersBillingFrequencyMonthly   OverrideOverrideSpecifiersBillingFrequency = "MONTHLY"
	OverrideOverrideSpecifiersBillingFrequencyQuarterly OverrideOverrideSpecifiersBillingFrequency = "QUARTERLY"
	OverrideOverrideSpecifiersBillingFrequencyAnnual    OverrideOverrideSpecifiersBillingFrequency = "ANNUAL"
)

func (r OverrideOverrideSpecifiersBillingFrequency) IsKnown() bool {
	switch r {
	case OverrideOverrideSpecifiersBillingFrequencyMonthly, OverrideOverrideSpecifiersBillingFrequencyQuarterly, OverrideOverrideSpecifiersBillingFrequencyAnnual:
		return true
	}
	return false
}

type OverrideOverrideTier struct {
	Multiplier float64                  `json:"multiplier,required"`
	Size       float64                  `json:"size"`
	JSON       overrideOverrideTierJSON `json:"-"`
}

// overrideOverrideTierJSON contains the JSON metadata for the struct
// [OverrideOverrideTier]
type overrideOverrideTierJSON struct {
	Multiplier  apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideOverrideTier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverrideTierJSON) RawJSON() string {
	return r.raw
}

type OverrideOverwriteRate struct {
	RateType   OverrideOverwriteRateRateType `json:"rate_type,required"`
	CreditType CreditTypeData                `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []Tier                    `json:"tiers"`
	JSON  overrideOverwriteRateJSON `json:"-"`
}

// overrideOverwriteRateJSON contains the JSON metadata for the struct
// [OverrideOverwriteRate]
type overrideOverwriteRateJSON struct {
	RateType    apijson.Field
	CreditType  apijson.Field
	CustomRate  apijson.Field
	IsProrated  apijson.Field
	Price       apijson.Field
	Quantity    apijson.Field
	Tiers       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideOverwriteRate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideOverwriteRateJSON) RawJSON() string {
	return r.raw
}

type OverrideOverwriteRateRateType string

const (
	OverrideOverwriteRateRateTypeFlat         OverrideOverwriteRateRateType = "FLAT"
	OverrideOverwriteRateRateTypePercentage   OverrideOverwriteRateRateType = "PERCENTAGE"
	OverrideOverwriteRateRateTypeSubscription OverrideOverwriteRateRateType = "SUBSCRIPTION"
	OverrideOverwriteRateRateTypeTiered       OverrideOverwriteRateRateType = "TIERED"
	OverrideOverwriteRateRateTypeCustom       OverrideOverwriteRateRateType = "CUSTOM"
)

func (r OverrideOverwriteRateRateType) IsKnown() bool {
	switch r {
	case OverrideOverwriteRateRateTypeFlat, OverrideOverwriteRateRateTypePercentage, OverrideOverwriteRateRateTypeSubscription, OverrideOverwriteRateRateTypeTiered, OverrideOverwriteRateRateTypeCustom:
		return true
	}
	return false
}

type OverrideProduct struct {
	ID   string              `json:"id,required" format:"uuid"`
	Name string              `json:"name,required"`
	JSON overrideProductJSON `json:"-"`
}

// overrideProductJSON contains the JSON metadata for the struct [OverrideProduct]
type overrideProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideProductJSON) RawJSON() string {
	return r.raw
}

type OverrideRateType string

const (
	OverrideRateTypeFlat         OverrideRateType = "FLAT"
	OverrideRateTypePercentage   OverrideRateType = "PERCENTAGE"
	OverrideRateTypeSubscription OverrideRateType = "SUBSCRIPTION"
	OverrideRateTypeTiered       OverrideRateType = "TIERED"
	OverrideRateTypeCustom       OverrideRateType = "CUSTOM"
)

func (r OverrideRateType) IsKnown() bool {
	switch r {
	case OverrideRateTypeFlat, OverrideRateTypePercentage, OverrideRateTypeSubscription, OverrideRateTypeTiered, OverrideRateTypeCustom:
		return true
	}
	return false
}

type OverrideTarget string

const (
	OverrideTargetCommitRate OverrideTarget = "COMMIT_RATE"
	OverrideTargetListRate   OverrideTarget = "LIST_RATE"
)

func (r OverrideTarget) IsKnown() bool {
	switch r {
	case OverrideTargetCommitRate, OverrideTargetListRate:
		return true
	}
	return false
}

type OverrideType string

const (
	OverrideTypeOverwrite  OverrideType = "OVERWRITE"
	OverrideTypeMultiplier OverrideType = "MULTIPLIER"
	OverrideTypeTiered     OverrideType = "TIERED"
)

func (r OverrideType) IsKnown() bool {
	switch r {
	case OverrideTypeOverwrite, OverrideTypeMultiplier, OverrideTypeTiered:
		return true
	}
	return false
}

type PropertyFilter struct {
	// The name of the event property.
	Name string `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists bool `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues []string `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues []string           `json:"not_in_values"`
	JSON        propertyFilterJSON `json:"-"`
}

// propertyFilterJSON contains the JSON metadata for the struct [PropertyFilter]
type propertyFilterJSON struct {
	Name        apijson.Field
	Exists      apijson.Field
	InValues    apijson.Field
	NotInValues apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PropertyFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r propertyFilterJSON) RawJSON() string {
	return r.raw
}

type PropertyFilterParam struct {
	// The name of the event property.
	Name param.Field[string] `json:"name,required"`
	// Determines whether the property must exist in the event. If true, only events
	// with this property will pass the filter. If false, only events without this
	// property will pass the filter. If null or omitted, the existence of the property
	// is optional.
	Exists param.Field[bool] `json:"exists"`
	// Specifies the allowed values for the property to match an event. An event will
	// pass the filter only if its property value is included in this list. If
	// undefined, all property values will pass the filter. Must be non-empty if
	// present.
	InValues param.Field[[]string] `json:"in_values"`
	// Specifies the values that prevent an event from matching the filter. An event
	// will not pass the filter if its property value is included in this list. If null
	// or empty, all property values will pass the filter. Must be non-empty if
	// present.
	NotInValues param.Field[[]string] `json:"not_in_values"`
}

func (r PropertyFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProService struct {
	ID string `json:"id,required" format:"uuid"`
	// Maximum amount for the term.
	MaxAmount float64 `json:"max_amount,required"`
	ProductID string  `json:"product_id,required" format:"uuid"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the
	// amount.
	Quantity float64 `json:"quantity,required"`
	// Unit price for the charge. Will be multiplied by quantity to determine the
	// amount and must be specified.
	UnitPrice    float64           `json:"unit_price,required"`
	CustomFields map[string]string `json:"custom_fields"`
	Description  string            `json:"description"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string         `json:"netsuite_sales_order_id"`
	JSON                 proServiceJSON `json:"-"`
}

// proServiceJSON contains the JSON metadata for the struct [ProService]
type proServiceJSON struct {
	ID                   apijson.Field
	MaxAmount            apijson.Field
	ProductID            apijson.Field
	Quantity             apijson.Field
	UnitPrice            apijson.Field
	CustomFields         apijson.Field
	Description          apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProService) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proServiceJSON) RawJSON() string {
	return r.raw
}

type Rate struct {
	RateType   RateRateType   `json:"rate_type,required"`
	CreditType CreditTypeData `json:"credit_type"`
	// Only set for CUSTOM rate_type. This field is interpreted by custom rate
	// processors.
	CustomRate map[string]interface{} `json:"custom_rate"`
	// Default proration configuration. Only valid for SUBSCRIPTION rate_type. Must be
	// set to true.
	IsProrated bool `json:"is_prorated"`
	// Default price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type,
	// this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price float64 `json:"price"`
	// if pricing groups are used, this will contain the values used to calculate the
	// price
	PricingGroupValues map[string]string `json:"pricing_group_values"`
	// Default quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity float64 `json:"quantity"`
	// Only set for TIERED rate_type.
	Tiers []Tier `json:"tiers"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed
	// using list prices rather than the standard rates for this product on the
	// contract.
	UseListPrices bool     `json:"use_list_prices"`
	JSON          rateJSON `json:"-"`
}

// rateJSON contains the JSON metadata for the struct [Rate]
type rateJSON struct {
	RateType           apijson.Field
	CreditType         apijson.Field
	CustomRate         apijson.Field
	IsProrated         apijson.Field
	Price              apijson.Field
	PricingGroupValues apijson.Field
	Quantity           apijson.Field
	Tiers              apijson.Field
	UseListPrices      apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Rate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateJSON) RawJSON() string {
	return r.raw
}

type RateRateType string

const (
	RateRateTypeFlat         RateRateType = "FLAT"
	RateRateTypePercentage   RateRateType = "PERCENTAGE"
	RateRateTypeSubscription RateRateType = "SUBSCRIPTION"
	RateRateTypeCustom       RateRateType = "CUSTOM"
	RateRateTypeTiered       RateRateType = "TIERED"
)

func (r RateRateType) IsKnown() bool {
	switch r {
	case RateRateTypeFlat, RateRateTypePercentage, RateRateTypeSubscription, RateRateTypeCustom, RateRateTypeTiered:
		return true
	}
	return false
}

type ScheduledCharge struct {
	ID           string                 `json:"id,required" format:"uuid"`
	Product      ScheduledChargeProduct `json:"product,required"`
	Schedule     SchedulePointInTime    `json:"schedule,required"`
	CustomFields map[string]string      `json:"custom_fields"`
	// displayed on invoices
	Name string `json:"name"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID string              `json:"netsuite_sales_order_id"`
	JSON                 scheduledChargeJSON `json:"-"`
}

// scheduledChargeJSON contains the JSON metadata for the struct [ScheduledCharge]
type scheduledChargeJSON struct {
	ID                   apijson.Field
	Product              apijson.Field
	Schedule             apijson.Field
	CustomFields         apijson.Field
	Name                 apijson.Field
	NetsuiteSalesOrderID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ScheduledCharge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledChargeJSON) RawJSON() string {
	return r.raw
}

type ScheduledChargeProduct struct {
	ID   string                     `json:"id,required" format:"uuid"`
	Name string                     `json:"name,required"`
	JSON scheduledChargeProductJSON `json:"-"`
}

// scheduledChargeProductJSON contains the JSON metadata for the struct
// [ScheduledChargeProduct]
type scheduledChargeProductJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduledChargeProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduledChargeProductJSON) RawJSON() string {
	return r.raw
}

type ScheduleDuration struct {
	ScheduleItems []ScheduleDurationScheduleItem `json:"schedule_items,required"`
	CreditType    CreditTypeData                 `json:"credit_type"`
	JSON          scheduleDurationJSON           `json:"-"`
}

// scheduleDurationJSON contains the JSON metadata for the struct
// [ScheduleDuration]
type scheduleDurationJSON struct {
	ScheduleItems apijson.Field
	CreditType    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ScheduleDuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDurationJSON) RawJSON() string {
	return r.raw
}

type ScheduleDurationScheduleItem struct {
	ID           string                           `json:"id,required" format:"uuid"`
	Amount       float64                          `json:"amount,required"`
	EndingBefore time.Time                        `json:"ending_before,required" format:"date-time"`
	StartingAt   time.Time                        `json:"starting_at,required" format:"date-time"`
	JSON         scheduleDurationScheduleItemJSON `json:"-"`
}

// scheduleDurationScheduleItemJSON contains the JSON metadata for the struct
// [ScheduleDurationScheduleItem]
type scheduleDurationScheduleItemJSON struct {
	ID           apijson.Field
	Amount       apijson.Field
	EndingBefore apijson.Field
	StartingAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ScheduleDurationScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleDurationScheduleItemJSON) RawJSON() string {
	return r.raw
}

type SchedulePointInTime struct {
	CreditType    CreditTypeData                    `json:"credit_type"`
	ScheduleItems []SchedulePointInTimeScheduleItem `json:"schedule_items"`
	JSON          schedulePointInTimeJSON           `json:"-"`
}

// schedulePointInTimeJSON contains the JSON metadata for the struct
// [SchedulePointInTime]
type schedulePointInTimeJSON struct {
	CreditType    apijson.Field
	ScheduleItems apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchedulePointInTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schedulePointInTimeJSON) RawJSON() string {
	return r.raw
}

type SchedulePointInTimeScheduleItem struct {
	ID        string                              `json:"id,required" format:"uuid"`
	Amount    float64                             `json:"amount,required"`
	InvoiceID string                              `json:"invoice_id,required" format:"uuid"`
	Quantity  float64                             `json:"quantity,required"`
	Timestamp time.Time                           `json:"timestamp,required" format:"date-time"`
	UnitPrice float64                             `json:"unit_price,required"`
	JSON      schedulePointInTimeScheduleItemJSON `json:"-"`
}

// schedulePointInTimeScheduleItemJSON contains the JSON metadata for the struct
// [SchedulePointInTimeScheduleItem]
type schedulePointInTimeScheduleItemJSON struct {
	ID          apijson.Field
	Amount      apijson.Field
	InvoiceID   apijson.Field
	Quantity    apijson.Field
	Timestamp   apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchedulePointInTimeScheduleItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schedulePointInTimeScheduleItemJSON) RawJSON() string {
	return r.raw
}

type Tier struct {
	Price float64  `json:"price,required"`
	Size  float64  `json:"size"`
	JSON  tierJSON `json:"-"`
}

// tierJSON contains the JSON metadata for the struct [Tier]
type tierJSON struct {
	Price       apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tierJSON) RawJSON() string {
	return r.raw
}

type TierParam struct {
	Price param.Field[float64] `json:"price,required"`
	Size  param.Field[float64] `json:"size"`
}

func (r TierParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
