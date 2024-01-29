package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
	InvoicePdf	        string      `json:"invoice_pdf"`
}

type Message struct {
	Header     []Header   `json:"Header"`
}

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
