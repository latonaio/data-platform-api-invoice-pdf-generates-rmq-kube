package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey       	string         `json:"connection_key"`
	Result              	bool           `json:"result"`
	RedisKey            	string         `json:"redis_key"`
	Filepath            	string         `json:"filepath"`
	APIStatusCode       	int            `json:"api_status_code"`
	RuntimeSessionID    	string         `json:"runtime_session_id"`
	BusinessPartnerID   	*int           `json:"business_partner"`
	ServiceLabel        	string         `json:"service_label"`
	Header		        	[]Header 	   `json:"InvoiceDocumentHeaderWithItem"`
	Items		    		[]Items		   `json:"InvoiceDocumentItem"`
	APISchema           	string         `json:"api_schema"`
	Accepter            	[]string       `json:"accepter"`
	Deleted             	bool           `json:"deleted"`
	DocData             	string         `json:"doc_data"`
	APIProcessingResult 	*bool          `json:"api_processing_result"`
	APIProcessingError  	string         `json:"api_processing_error"`
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
