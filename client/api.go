// Package client provides a Stripe client for invoking APIs across all resources
package client

import (
	. "github.com/splice/stripe-go"
	"github.com/splice/stripe-go/account"
	"github.com/splice/stripe-go/balance"
	"github.com/splice/stripe-go/bankaccount"
	"github.com/splice/stripe-go/bitcoinreceiver"
	"github.com/splice/stripe-go/bitcointransaction"
	"github.com/splice/stripe-go/card"
	"github.com/splice/stripe-go/charge"
	"github.com/splice/stripe-go/countryspec"
	"github.com/splice/stripe-go/coupon"
	"github.com/splice/stripe-go/customer"
	"github.com/splice/stripe-go/discount"
	"github.com/splice/stripe-go/dispute"
	"github.com/splice/stripe-go/event"
	"github.com/splice/stripe-go/fee"
	"github.com/splice/stripe-go/feerefund"
	"github.com/splice/stripe-go/fileupload"
	"github.com/splice/stripe-go/invoice"
	"github.com/splice/stripe-go/invoiceitem"
	"github.com/splice/stripe-go/order"
	"github.com/splice/stripe-go/plan"
	"github.com/splice/stripe-go/product"
	"github.com/splice/stripe-go/recipient"
	"github.com/splice/stripe-go/refund"
	"github.com/splice/stripe-go/reversal"
	"github.com/splice/stripe-go/sku"
	"github.com/splice/stripe-go/sub"
	"github.com/splice/stripe-go/token"
	"github.com/splice/stripe-go/transfer"
)

// API is the Stripe client. It contains all the different resources available.
type API struct {
	// Charges is the client used to invoke /charges APIs.
	// For more details see https://stripe.com/docs/api#charges.
	Charges *charge.Client
	// Customers is the client used to invoke /customers APIs.
	// For more details see https://stripe.com/docs/api#customers.
	Customers *customer.Client
	// Cards is the client used to invoke /cards APIs.
	// For more details see https://stripe.com/docs/api#cards.
	Cards *card.Client
	// Subs is the client used to invoke /subscriptions APIs.
	// For more details see https://stripe.com/docs/api#subscriptions.
	Subs *sub.Client
	// Plans is the client used to invoke /plans APIs.
	// For more details see https://stripe.com/docs/api#plans.
	Plans *plan.Client
	// Coupons is the client used to invoke /coupons APIs.
	// For more details see https://stripe.com/docs/api#coupons.
	Coupons *coupon.Client
	// Discounts is the client used to invoke discount-related APIs.
	// For mode details see https://stripe.com/docs/api#discounts.
	Discounts *discount.Client
	// Invoices is the client used to invoke /invoices APIs.
	// For more details see https://stripe.com/docs/api#invoices.
	Invoices *invoice.Client
	// InvoiceItems is the client used to invoke /invoiceitems APIs.
	// For more details see https://stripe.com/docs/api#invoiceitems.
	InvoiceItems *invoiceitem.Client
	// Disputes is the client used to invoke dispute-related APIs.
	// For more details see https://stripe.com/docs/api#disputes.
	Disputes *dispute.Client
	// Transfers is the client used to invoke /transfers APIs.
	// For more details see https://stripe.com/docs/api#transfers.
	Transfers *transfer.Client
	// Recipients is the client used to invoke /recipients APIs.
	// For more details see https://stripe.com/docs/api#recipients.
	Recipients *recipient.Client
	// Refunds is the client used to invoke /refunds APIs.
	// For more details see https://stripe.com/docs/api#refunds.
	Refunds *refund.Client
	// Fees is the client used to invoke /application_fees APIs.
	// For more details see https://stripe.com/docs/api#application_fees.
	Fees *fee.Client
	// FeeRefunds is the client used to invoke /application_fees/refunds APIs.
	// For more details see https://stripe.com/docs/api#fee_refundss.
	FeeRefunds *feerefund.Client
	// Account is the client used to invoke /account APIs.
	// For more details see https://stripe.com/docs/api#account.
	Account *account.Client
	// CountrySpec is the client used to invoke /country_specs APIs.
	// For more details see https://stripe.com/docs/api#country_specs.
	CountrySpec *countryspec.Client
	// Balance is the client used to invoke /balance and transaction-related APIs.
	// For more details see https://stripe.com/docs/api#balance.
	Balance *balance.Client
	// Events is the client used to invoke /events APIs.
	// For more details see https://stripe.com/docs/api#events.
	Events *event.Client
	// Tokens is the client used to invoke /tokens APIs.
	// For more details see https://stripe.com/docs/api#tokens.
	Tokens *token.Client
	// FileUploads is the client used to invoke the uploads /files APIs.
	// For more details see https://stripe.com/docs/api#file_uploads.
	FileUploads *fileupload.Client
	// BitcoinReceivers is the client used to invoke /bitcoin/receivers APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinReceivers *bitcoinreceiver.Client
	// BitcoinTransactions is the client used to invoke /bitcoin/transactions APIs.
	// For more details see https://stripe.com/docs/api#bitcoin_receivers.
	BitcoinTransactions *bitcointransaction.Client
	// Reversals is the client used to invoke /transfers/reversals APIs.
	Reversals *reversal.Client
	// BankAccounts is the client used to invoke /accounts/bank_accounts APIs.
	BankAccounts *bankaccount.Client
	// Products is the client used to invoke /products APIs.
	// For more details see https://stripe.com/docs/api#products.
	Products *product.Client
	// Orders is the client used to invoke /orders APIs.
	// For more details see https://stripe.com/docs/api#orders.
	Orders *order.Client
	// Skus is the client used to invoke /skus APIs.
	// For more details see https://stripe.com/docs/api#skus.
	Skus *sku.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *Backends) {
	if backends == nil {
		backends = &Backends{GetBackend(APIBackend), GetBackend(UploadsBackend)}
	}

	a.Charges = &charge.Client{B: backends.API, Key: key}
	a.Customers = &customer.Client{B: backends.API, Key: key}
	a.Cards = &card.Client{B: backends.API, Key: key}
	a.Subs = &sub.Client{B: backends.API, Key: key}
	a.Plans = &plan.Client{B: backends.API, Key: key}
	a.Coupons = &coupon.Client{B: backends.API, Key: key}
	a.Discounts = &discount.Client{B: backends.API, Key: key}
	a.Invoices = &invoice.Client{B: backends.API, Key: key}
	a.InvoiceItems = &invoiceitem.Client{B: backends.API, Key: key}
	a.Disputes = &dispute.Client{B: backends.API, Key: key}
	a.Transfers = &transfer.Client{B: backends.API, Key: key}
	a.Recipients = &recipient.Client{B: backends.API, Key: key}
	a.Refunds = &refund.Client{B: backends.API, Key: key}
	a.Fees = &fee.Client{B: backends.API, Key: key}
	a.FeeRefunds = &feerefund.Client{B: backends.API, Key: key}
	a.Account = &account.Client{B: backends.API, Key: key}
	a.CountrySpec = &countryspec.Client{B: backends.API, Key: key}
	a.Balance = &balance.Client{B: backends.API, Key: key}
	a.Events = &event.Client{B: backends.API, Key: key}
	a.Tokens = &token.Client{B: backends.API, Key: key}
	a.FileUploads = &fileupload.Client{B: backends.Uploads, Key: key}
	a.BitcoinReceivers = &bitcoinreceiver.Client{B: backends.API, Key: key}
	a.BitcoinTransactions = &bitcointransaction.Client{B: backends.API, Key: key}
	a.Reversals = &reversal.Client{B: backends.API, Key: key}
	a.BankAccounts = &bankaccount.Client{B: backends.API, Key: key}
	a.Products = &product.Client{B: backends.API, Key: key}
	a.Orders = &order.Client{B: backends.API, Key: key}
	a.Skus = &sku.Client{B: backends.API, Key: key}
}

// New creates a new Stripe client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}
