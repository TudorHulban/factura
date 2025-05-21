package service

import (
	"os"
	"testing"
	"time"

	"github.com/TudorHulban/factura/domain"
)

func TestRender(t *testing.T) {
	s := NewService()

	now := time.Now()

	inv := domain.Invoice{
		SchemaLocation:       "urn:oasis:names:specification:ubl:schema:xsd:ApplicationResponse-2 file://ro_cius-ApplicationResponse-1.0.xsd",
		UBLVersionID:         "2.1",
		CustomizationID:      "urn:cen.eu:en16931:2017#compliant#urn:ro-efactura.mfinante.gov.ro:ro-cius:1.0",
		ID:                   "RO1234567",
		IssueDate:            now.Format("2006-01-02"),
		IssueTime:            now.Format("15:04:05Z"),
		InvoiceTypeCode:      "380", // commercial invoice
		DocumentCurrencyCode: "RON",

		AccountingSupplierParty: domain.InvoiceParty{
			PartyIdentification: domain.PartyIdentification{
				ID:       "RO12345678",
				SchemeID: "VAT",
			},

			PartyName: domain.PartyName{
				Name: "Supplier Company SRL",
			},

			PostalAddress: domain.PostalAddress{
				StreetName:       "Str. Example",
				CityName:         "Bucharest",
				PostalZone:       "010001",
				CountrySubentity: "Bucuresti",

				Country: domain.Country{
					IdentificationCode: "RO",
				},
			},

			PartyTaxScheme: domain.PartyTaxScheme{
				CompanyID: "12345678",
				TaxScheme: domain.TaxScheme{
					ID: "VAT",
				},
			},
		},

		AccountingCustomerParty: domain.InvoiceParty{
			PartyIdentification: domain.PartyIdentification{
				ID:       "RO87654321",
				SchemeID: "VAT",
			},

			PartyName: domain.PartyName{
				Name: "Customer Company SA",
			},

			PostalAddress: domain.PostalAddress{
				StreetName:       "Blvd. Test",
				CityName:         "Cluj-Napoca",
				PostalZone:       "400001",
				CountrySubentity: "Cluj",

				Country: domain.Country{
					IdentificationCode: "RO",
				},
			},

			PartyTaxScheme: domain.PartyTaxScheme{
				CompanyID: "87654321",
				TaxScheme: domain.TaxScheme{
					ID: "VAT",
				},
			},
		},

		TaxTotal: domain.Taxtotal{
			TaxAmount: domain.Amount{
				CurrencyID: "RON",
				Value:      "19.00",
			},

			TaxSubtotal: domain.Taxsubtotal{
				TaxableAmount: domain.Amount{
					CurrencyID: "RON",
					Value:      "100.00",
				},
				TaxAmount: domain.Amount{
					CurrencyID: "RON",
					Value:      "19.00",
				},
				TaxCategory: domain.TaxCategory{
					ID:      "S",
					Percent: "19",

					TaxScheme: domain.TaxScheme{
						ID: "VAT",
					},
				},
			},
		},

		LegalMonetaryTotal: domain.LegalMonetaryTotal{
			LineExtensionAmount: domain.Amount{
				CurrencyID: "RON",
				Value:      "100.00",
			},
			TaxExclusiveAmount: domain.Amount{
				CurrencyID: "RON",
				Value:      "100.00",
			},
			TaxInclusiveAmount: domain.Amount{
				CurrencyID: "RON",
				Value:      "119.00",
			},
			PayableAmount: domain.Amount{
				CurrencyID: "RON",
				Value:      "119.00",
			},
		},

		InvoiceLine: []domain.Invoiceline{
			{
				ID: "1",
				InvoicedQuantity: domain.Quantity{
					UnitCode: "C62",
					Value:    "1",
				},
				LineExtensionAmount: domain.Amount{
					CurrencyID: "RON",
					Value:      "100.00",
				},
				Item: domain.Item{
					Name: "Product A",
					ClassifiedTaxCategory: domain.TaxCategory{
						ID:      "S",
						Percent: "19",
						TaxScheme: domain.TaxScheme{
							ID: "VAT",
						},
					},
				},
				Price: domain.Price{
					PriceAmount: domain.Amount{
						CurrencyID: "RON",
						Value:      "100.00",
					},
				},
			},
		},
	}

	s.RenderInvoiceDataTo(
		&inv,
		os.Stdout,
	)
}
