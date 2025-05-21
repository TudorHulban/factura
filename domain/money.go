package domain

type Amount struct {
	CurrencyID string `xml:"currencyid,attr"`
	Value      string `xml:",chardata"`
}

type TaxScheme struct {
	ID string `xml:"id"`
}

type TaxCategory struct {
	ID                     string `xml:"id"`
	Percent                string `xml:"percent"`
	TaxExemptionReasonCode string `xml:"taxexemptionreasoncode,omitempty"`

	TaxScheme TaxScheme `xml:"taxscheme"`
}

type Taxsubtotal struct {
	TaxableAmount Amount `xml:"taxableamount"`
	TaxAmount     Amount `xml:"taxamount"`

	TaxCategory TaxCategory `xml:"taxcategory"`
}

type Taxtotal struct {
	TaxAmount   Amount      `xml:"taxamount"`
	TaxSubtotal Taxsubtotal `xml:"taxsubtotal"`
}

type PartyTaxScheme struct {
	CompanyID string `xml:"companyid"`

	TaxScheme TaxScheme `xml:"taxscheme"`
}

type LegalMonetaryTotal struct {
	LineExtensionAmount Amount `xml:"lineextensionamount"`
	TaxExclusiveAmount  Amount `xml:"taxexclusiveamount"`
	TaxInclusiveAmount  Amount `xml:"taxinclusiveamount"`
	PayableAmount       Amount `xml:"payableamount"`
}
