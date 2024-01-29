package requests

type Header struct {
	InvoiceDocument                   int      `json:"InvoiceDocument"`
	InvoiceDocumentDate               string   `json:"InvoiceDocumentDate"`
	InvoicePeriodStartDate            string   `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              string   `json:"InvoicePeriodEndDate"`
	BillToParty                       int      `json:"BillToParty"`
	BillToPartyName                   string   `json:"BillToPartyName"`
	BillFromParty                     int      `json:"BillFromParty"`
	BillFromPartyName                 string   `json:"BillFromPartyName"`
	IsExportImport                    *bool    `json:"IsExportImport"`
	TotalNetAmount                    float32  `json:"TotalNetAmount"`
	TotalTaxAmount                    float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                  float32  `json:"TotalGrossAmount"`
	TransactionCurrency               string   `json:"TransactionCurrency"`
	Incoterms                         *string  `json:"Incoterms"`
	PaymentTerms                      string   `json:"PaymentTerms"`
	DueCalculationBaseDate            *string  `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string  `json:"PaymentDueDate"`
	NetPaymentDays                    *int     `json:"NetPaymentDays"`
	PaymentMethod                     string   `json:"PaymentMethod"`
	Contract		                  *int     `json:"Contract"`
	ContractItem	                  *int     `json:"ContractItem"`
	Items				  			  []Items  `json:"Items"`
}

type Items struct {
	InvoiceDocument                         int      `json:"InvoiceDocument"`
	InvoiceDocumentItem                     int      `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory             string   `json:"InvoiceDocumentItemCategory"`
	Product                                 string   `json:"Product"`
	SizeOrDimensionText                     *string  `json:"SizeOrDimensionText"`
	InvoiceDocumentItemText                 *string  `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer          string   `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller         string   `json:"InvoiceDocumentItemTextBySeller"`
	InvoiceQuantity                         float32  `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                     string   `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit               float32  `json:"InvoiceQuantityInBaseUnit"`
	ActualGoodsReceiptDate                  *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string  `json:"ActualGoodsReceiptTime"`
	NetAmount                               float32  `json:"NetAmount"`
	TaxAmount                               float32  `json:"TaxAmount"`
	GrossAmount                             float32  `json:"GrossAmount"`
}
