package domain

import "encoding/xml"

type PartyIdentification struct {
	ID       string `xml:"id"`
	SchemeID string `xml:"schemeid,attr"`
}

type PartyName struct {
	Name string `xml:"name"`
}

type Country struct {
	IdentificationCode string `xml:"identificationcode"`
}

type PostalAddress struct {
	StreetName       string `xml:"streetname"`
	CityName         string `xml:"cityname"`
	PostalZone       string `xml:"postalzone"`
	CountrySubentity string `xml:"countrysubentity"`

	Country Country `xml:"country"`
}

type InvoiceParty struct {
	PartyIdentification PartyIdentification `xml:"partyidentification"`
	PartyName           PartyName           `xml:"partyname"`
	PostalAddress       PostalAddress       `xml:"postaladdress"`
	PartyTaxScheme      PartyTaxScheme      `xml:"partytaxscheme"`
}

type Price struct {
	PriceAmount Amount `xml:"priceamount"`
}

type Item struct {
	Name                  string      `xml:"name"`
	ClassifiedTaxCategory TaxCategory `xml:"classifiedtaxcategory"`
}

type Invoiceline struct {
	ID string `xml:"id"`

	InvoicedQuantity    Quantity `xml:"invoicedquantity"`
	LineExtensionAmount Amount   `xml:"lineextensionamount"`
	Item                Item     `xml:"item"`
	Price               Price    `xml:"price"`
}

type Quantity struct {
	UnitCode string `xml:"unitcode,attr"`
	Value    string `xml:",chardata"`
}

type Invoice struct {
	XMLName xml.Name `xml:"applicationresponse"`

	SchemaLocation       string `xml:"xsi:schemalocation,attr"`
	UBLVersionID         string `xml:"ublversionid"`
	CustomizationID      string `xml:"customizationid"`
	ID                   string `xml:"id"`
	IssueDate            string `xml:"issuedate"`
	IssueTime            string `xml:"issuetime"`
	InvoiceTypeCode      string `xml:"invoicetypecode"`
	DocumentCurrencyCode string `xml:"documentcurrencycode"`

	AccountingSupplierParty InvoiceParty       `xml:"accountingsupplierparty"`
	AccountingCustomerParty InvoiceParty       `xml:"accountingcustomerparty"`
	TaxTotal                Taxtotal           `xml:"taxtotal"`
	LegalMonetaryTotal      LegalMonetaryTotal `xml:"legalmonetarytotal"`

	InvoiceLine []Invoiceline `xml:"invoiceline"`
}
