package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-invoice-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputItems []dpfm_api_input_reader.Items,
) *Header {
	pm := &Header{}

	var items []Items
	for _, dataItems := range inputItems {
		if headerData.InvoiceDocument == dataItems.InvoiceDocument {
			items = append(items, Items{
				InvoiceDocument:						dataItems.InvoiceDocument,
				InvoiceDocumentItem:					dataItems.InvoiceDocumentItem,
				Product:								dataItems.Product,
				SizeOrDimensionText:					dataItems.SizeOrDimensionText,
				InvoiceDocumentItemText:        		dataItems.InvoiceDocumentItemText,
				InvoiceQuantity:						dataItems.InvoiceQuantity,
				InvoiceQuantityUnit:					dataItems.InvoiceQuantityUnit,
				InvoiceQuantityInBaseUnit:				dataItems.InvoiceQuantityInBaseUnit,
				ActualGoodsReceiptDate:        			dataItems.ActualGoodsReceiptDate,
				ActualGoodsReceiptTime:        			dataItems.ActualGoodsReceiptTime,
				NetAmount:								dataItems.NetAmount,
				TaxAmount:								dataItems.TaxAmount,
				GrossAmount:							dataItems.GrossAmount,
			})
		}
	}

	pm = &Header{
				InvoiceDocument:   					headerData.InvoiceDocument,
				InvoiceDocumentDate:   				headerData.InvoiceDocumentDate,
				InvoicePeriodStartDate:   			headerData.InvoicePeriodStartDate,
				InvoicePeriodEndDate:   			headerData.InvoicePeriodEndDate,
				BillToParty: 						headerData.BillToParty,
				BillToPartyName: 					headerData.BillToPartyName,
				BillFromParty: 						headerData.BillFromParty,
				BillFromPartyName:              	headerData.BillFromPartyName,
				IsExportImport:              		headerData.IsExportImport,
				TotalNetAmount: 					headerData.TotalNetAmount,
				TotalTaxAmount: 					headerData.TotalTaxAmount,
				TotalGrossAmount: 					headerData.TotalGrossAmount,
				TransactionCurrency: 				headerData.TransactionCurrency,
				Incoterms: 							headerData.Incoterms,
				PaymentTerms: 						headerData.PaymentTerms,
				DueCalculationBaseDate: 			headerData.DueCalculationBaseDate,
				PaymentDueDate: 					headerData.PaymentDueDate,
				NetPaymentDays: 					headerData.NetPaymentDays,
				PaymentMethod: 						headerData.PaymentMethod,
				Contract: 							headerData.Contract,
				ContractItem: 						headerData.ContractItem,
				Items:     							items,
	}

	return pm
}
