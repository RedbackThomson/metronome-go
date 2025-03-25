// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package metronome_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/internal/testutil"
	"github.com/Metronome-Industries/metronome-go/option"
	"github.com/Metronome-Industries/metronome-go/shared"
)

func TestV2ContractGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.Get(context.TODO(), metronome.V2ContractGetParams{
		ContractID:     metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID:     metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		AsOfDate:       metronome.F(time.Now()),
		IncludeBalance: metronome.F(true),
		IncludeLedgers: metronome.F(true),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.List(context.TODO(), metronome.V2ContractListParams{
		CustomerID:      metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		CoveringDate:    metronome.F(time.Now()),
		IncludeArchived: metronome.F(true),
		IncludeBalance:  metronome.F(true),
		IncludeLedgers:  metronome.F(true),
		StartingAt:      metronome.F(time.Now()),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractEditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.Edit(context.TODO(), metronome.V2ContractEditParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
		AddCommits: metronome.F([]metronome.V2ContractEditParamsAddCommit{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Type:      metronome.F(metronome.V2ContractEditParamsAddCommitsTypePrepaid),
			AccessSchedule: metronome.F(metronome.V2ContractEditParamsAddCommitsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V2ContractEditParamsAddCommitsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			Amount:                metronome.F(0.000000),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description: metronome.F("description"),
			InvoiceSchedule: metronome.F(metronome.V2ContractEditParamsAddCommitsInvoiceSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V2ContractEditParamsAddCommitsInvoiceScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V2ContractEditParamsAddCommitsInvoiceScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V2ContractEditParamsAddCommitsInvoiceScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.V2ContractEditParamsAddCommitsRateTypeCommitRate),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		AddCredits: metronome.F([]metronome.V2ContractEditParamsAddCredit{{
			AccessSchedule: metronome.F(metronome.V2ContractEditParamsAddCreditsAccessSchedule{
				ScheduleItems: metronome.F([]metronome.V2ContractEditParamsAddCreditsAccessScheduleScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Priority:             metronome.F(0.000000),
			RateType:             metronome.F(metronome.V2ContractEditParamsAddCreditsRateTypeCommitRate),
		}}),
		AddDiscounts: metronome.F([]metronome.V2ContractEditParamsAddDiscount{{
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Schedule: metronome.F(metronome.V2ContractEditParamsAddDiscountsSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V2ContractEditParamsAddDiscountsScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V2ContractEditParamsAddDiscountsScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V2ContractEditParamsAddDiscountsScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V2ContractEditParamsAddDiscountsScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		AddOverrides: metronome.F([]metronome.V2ContractEditParamsAddOverride{{
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductTags: metronome.F([]string{"string"}),
			EndingBefore:          metronome.F(time.Now()),
			Entitled:              metronome.F(true),
			IsCommitSpecific:      metronome.F(true),
			Multiplier:            metronome.F(2.000000),
			OverrideSpecifiers: metronome.F([]metronome.V2ContractEditParamsAddOverridesOverrideSpecifier{{
				CommitIDs: metronome.F([]string{"string"}),
				PresentationGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				PricingGroupValues: metronome.F(map[string]string{
					"foo": "string",
				}),
				ProductID:          metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				ProductTags:        metronome.F([]string{"string"}),
				RecurringCommitIDs: metronome.F([]string{"string"}),
				RecurringCreditIDs: metronome.F([]string{"string"}),
			}}),
			OverwriteRate: metronome.F(metronome.V2ContractEditParamsAddOverridesOverwriteRate{
				RateType:     metronome.F(metronome.V2ContractEditParamsAddOverridesOverwriteRateRateTypeFlat),
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				CustomRate: metronome.F(map[string]interface{}{
					"foo": "bar",
				}),
				IsProrated: metronome.F(true),
				Price:      metronome.F(0.000000),
				Quantity:   metronome.F(0.000000),
				Tiers: metronome.F([]shared.TierParam{{
					Price: metronome.F(0.000000),
					Size:  metronome.F(0.000000),
				}}),
			}),
			Priority:  metronome.F(100.000000),
			ProductID: metronome.F("d4fc086c-d8e5-4091-a235-fbba5da4ec14"),
			Target:    metronome.F(metronome.V2ContractEditParamsAddOverridesTargetCommitRate),
			Tiers: metronome.F([]metronome.V2ContractEditParamsAddOverridesTier{{
				Multiplier: metronome.F(0.000000),
				Size:       metronome.F(0.000000),
			}}),
			Type: metronome.F(metronome.V2ContractEditParamsAddOverridesTypeOverwrite),
		}}),
		AddProfessionalServices: metronome.F([]metronome.V2ContractEditParamsAddProfessionalService{{
			MaxAmount: metronome.F(0.000000),
			ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Quantity:  metronome.F(0.000000),
			UnitPrice: metronome.F(0.000000),
			CustomFields: metronome.F(map[string]string{
				"foo": "string",
			}),
			Description:          metronome.F("description"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		AddRecurringCommits: metronome.F([]metronome.V2ContractEditParamsAddRecurringCommit{{
			AccessAmount: metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsCommitDuration{
				Unit:  metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsCommitDurationUnitPeriods),
				Value: metronome.F(0.000000),
			}),
			Priority:              metronome.F(0.000000),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			Description:           metronome.F("description"),
			EndingBefore:          metronome.F(time.Now()),
			InvoiceAmount: metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsInvoiceAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			Proration:            metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsProrationNone),
			RateType:             metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsRateTypeCommitRate),
			RecurrenceFrequency:  metronome.F(metronome.V2ContractEditParamsAddRecurringCommitsRecurrenceFrequencyMonthly),
			RolloverFraction:     metronome.F(0.000000),
			TemporaryID:          metronome.F("temporary_id"),
		}}),
		AddRecurringCredits: metronome.F([]metronome.V2ContractEditParamsAddRecurringCredit{{
			AccessAmount: metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsAccessAmount{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Quantity:     metronome.F(0.000000),
				UnitPrice:    metronome.F(0.000000),
			}),
			CommitDuration: metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsCommitDuration{
				Unit:  metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsCommitDurationUnitPeriods),
				Value: metronome.F(0.000000),
			}),
			Priority:              metronome.F(0.000000),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			StartingAt:            metronome.F(time.Now()),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			Description:           metronome.F("description"),
			EndingBefore:          metronome.F(time.Now()),
			Name:                  metronome.F("x"),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			Proration:             metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsProrationNone),
			RateType:              metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsRateTypeCommitRate),
			RecurrenceFrequency:   metronome.F(metronome.V2ContractEditParamsAddRecurringCreditsRecurrenceFrequencyMonthly),
			RolloverFraction:      metronome.F(0.000000),
			TemporaryID:           metronome.F("temporary_id"),
		}}),
		AddResellerRoyalties: metronome.F([]metronome.V2ContractEditParamsAddResellerRoyalty{{
			ResellerType:          metronome.F(metronome.V2ContractEditParamsAddResellerRoyaltiesResellerTypeAws),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			AwsOptions: metronome.F(metronome.V2ContractEditParamsAddResellerRoyaltiesAwsOptions{
				AwsAccountNumber:    metronome.F("aws_account_number"),
				AwsOfferID:          metronome.F("aws_offer_id"),
				AwsPayerReferenceID: metronome.F("aws_payer_reference_id"),
			}),
			EndingBefore: metronome.F(time.Now()),
			Fraction:     metronome.F(0.000000),
			GcpOptions: metronome.F(metronome.V2ContractEditParamsAddResellerRoyaltiesGcpOptions{
				GcpAccountID: metronome.F("gcp_account_id"),
				GcpOfferID:   metronome.F("gcp_offer_id"),
			}),
			NetsuiteResellerID:    metronome.F("netsuite_reseller_id"),
			ResellerContractValue: metronome.F(0.000000),
			StartingAt:            metronome.F(time.Now()),
		}}),
		AddScheduledCharges: metronome.F([]metronome.V2ContractEditParamsAddScheduledCharge{{
			ProductID: metronome.F("2e30f074-d04c-412e-a134-851ebfa5ceb2"),
			Schedule: metronome.F(metronome.V2ContractEditParamsAddScheduledChargesSchedule{
				CreditTypeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				RecurringSchedule: metronome.F(metronome.V2ContractEditParamsAddScheduledChargesScheduleRecurringSchedule{
					AmountDistribution: metronome.F(metronome.V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleAmountDistributionDivided),
					EndingBefore:       metronome.F(time.Now()),
					Frequency:          metronome.F(metronome.V2ContractEditParamsAddScheduledChargesScheduleRecurringScheduleFrequencyMonthly),
					StartingAt:         metronome.F(time.Now()),
					Amount:             metronome.F(0.000000),
					Quantity:           metronome.F(0.000000),
					UnitPrice:          metronome.F(0.000000),
				}),
				ScheduleItems: metronome.F([]metronome.V2ContractEditParamsAddScheduledChargesScheduleScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(1.000000),
					UnitPrice: metronome.F(1000000.000000),
				}}),
			}),
			Name:                 metronome.F("x"),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
		UpdateCommits: metronome.F([]metronome.V2ContractEditParamsUpdateCommit{{
			CommitID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.V2ContractEditParamsUpdateCommitsAccessSchedule{
				AddScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsAccessScheduleAddScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				RemoveScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsAccessScheduleRemoveScheduleItem{{
					ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				UpdateScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsAccessScheduleUpdateScheduleItem{{
					ID:           metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
			}),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			InvoiceSchedule: metronome.F(metronome.V2ContractEditParamsUpdateCommitsInvoiceSchedule{
				AddScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsInvoiceScheduleAddScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
				RemoveScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsInvoiceScheduleRemoveScheduleItem{{
					ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				UpdateScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCommitsInvoiceScheduleUpdateScheduleItem{{
					ID:        metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
			ProductID:            metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			RolloverFraction:     metronome.F(0.000000),
		}}),
		UpdateCredits: metronome.F([]metronome.V2ContractEditParamsUpdateCredit{{
			CreditID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			AccessSchedule: metronome.F(metronome.V2ContractEditParamsUpdateCreditsAccessSchedule{
				AddScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCreditsAccessScheduleAddScheduleItem{{
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
				RemoveScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCreditsAccessScheduleRemoveScheduleItem{{
					ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				UpdateScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateCreditsAccessScheduleUpdateScheduleItem{{
					ID:           metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Amount:       metronome.F(0.000000),
					EndingBefore: metronome.F(time.Now()),
					StartingAt:   metronome.F(time.Now()),
				}}),
			}),
			ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
			ApplicableProductTags: metronome.F([]string{"string"}),
			NetsuiteSalesOrderID:  metronome.F("netsuite_sales_order_id"),
			ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		}}),
		UpdateScheduledCharges: metronome.F([]metronome.V2ContractEditParamsUpdateScheduledCharge{{
			ScheduledChargeID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			InvoiceSchedule: metronome.F(metronome.V2ContractEditParamsUpdateScheduledChargesInvoiceSchedule{
				AddScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleAddScheduleItem{{
					Timestamp: metronome.F(time.Now()),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					UnitPrice: metronome.F(0.000000),
				}}),
				RemoveScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleRemoveScheduleItem{{
					ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				}}),
				UpdateScheduleItems: metronome.F([]metronome.V2ContractEditParamsUpdateScheduledChargesInvoiceScheduleUpdateScheduleItem{{
					ID:        metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
					Amount:    metronome.F(0.000000),
					Quantity:  metronome.F(0.000000),
					Timestamp: metronome.F(time.Now()),
					UnitPrice: metronome.F(0.000000),
				}}),
			}),
			NetsuiteSalesOrderID: metronome.F("netsuite_sales_order_id"),
		}}),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractEditCommitWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.EditCommit(context.TODO(), metronome.V2ContractEditCommitParams{
		CommitID:   metronome.F("5e7e82cf-ccb7-428c-a96f-a8e4f67af822"),
		CustomerID: metronome.F("4c91c473-fc12-445a-9c38-40421d47023f"),
		AccessSchedule: metronome.F(metronome.V2ContractEditCommitParamsAccessSchedule{
			AddScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsAccessScheduleAddScheduleItem{{
				Amount:       metronome.F(0.000000),
				EndingBefore: metronome.F(time.Now()),
				StartingAt:   metronome.F(time.Now()),
			}}),
			RemoveScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsAccessScheduleRemoveScheduleItem{{
				ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}}),
			UpdateScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsAccessScheduleUpdateScheduleItem{{
				ID:           metronome.F("d5edbd32-c744-48cb-9475-a9bca0e6fa39"),
				Amount:       metronome.F(0.000000),
				EndingBefore: metronome.F(time.Now()),
				StartingAt:   metronome.F(time.Now()),
			}}),
		}),
		ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ApplicableProductTags: metronome.F([]string{"string"}),
		InvoiceContractID:     metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		InvoiceSchedule: metronome.F(metronome.V2ContractEditCommitParamsInvoiceSchedule{
			AddScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsInvoiceScheduleAddScheduleItem{{
				Timestamp: metronome.F(time.Now()),
				Amount:    metronome.F(0.000000),
				Quantity:  metronome.F(0.000000),
				UnitPrice: metronome.F(0.000000),
			}}),
			RemoveScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsInvoiceScheduleRemoveScheduleItem{{
				ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}}),
			UpdateScheduleItems: metronome.F([]metronome.V2ContractEditCommitParamsInvoiceScheduleUpdateScheduleItem{{
				ID:        metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Amount:    metronome.F(0.000000),
				Quantity:  metronome.F(0.000000),
				Timestamp: metronome.F(time.Now()),
				UnitPrice: metronome.F(0.000000),
			}}),
		}),
		ProductID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractEditCreditWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.EditCredit(context.TODO(), metronome.V2ContractEditCreditParams{
		CreditID:   metronome.F("5e7e82cf-ccb7-428c-a96f-a8e4f67af822"),
		CustomerID: metronome.F("4c91c473-fc12-445a-9c38-40421d47023f"),
		AccessSchedule: metronome.F(metronome.V2ContractEditCreditParamsAccessSchedule{
			AddScheduleItems: metronome.F([]metronome.V2ContractEditCreditParamsAccessScheduleAddScheduleItem{{
				Amount:       metronome.F(0.000000),
				EndingBefore: metronome.F(time.Now()),
				StartingAt:   metronome.F(time.Now()),
			}}),
			RemoveScheduleItems: metronome.F([]metronome.V2ContractEditCreditParamsAccessScheduleRemoveScheduleItem{{
				ID: metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}}),
			UpdateScheduleItems: metronome.F([]metronome.V2ContractEditCreditParamsAccessScheduleUpdateScheduleItem{{
				ID:           metronome.F("d5edbd32-c744-48cb-9475-a9bca0e6fa39"),
				Amount:       metronome.F(0.000000),
				EndingBefore: metronome.F(time.Now()),
				StartingAt:   metronome.F(time.Now()),
			}}),
		}),
		ApplicableProductIDs:  metronome.F([]string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"}),
		ApplicableProductTags: metronome.F([]string{"string"}),
		ProductID:             metronome.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV2ContractGetEditHistory(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := metronome.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.V2.Contracts.GetEditHistory(context.TODO(), metronome.V2ContractGetEditHistoryParams{
		ContractID: metronome.F("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		CustomerID: metronome.F("13117714-3f05-48e5-a6e9-a66093f13b4d"),
	})
	if err != nil {
		var apierr *metronome.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
