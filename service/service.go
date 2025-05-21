package service

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/TudorHulban/factura/domain"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RenderInvoiceDataTo(invoice *domain.Invoice, w io.Writer) (int, error) {
	xmloutput, errMarshal := xml.MarshalIndent(invoice, "", "  ")
	if errMarshal != nil {
		return 0,
			fmt.Errorf(
				"error marshaling xml: %v",
				errMarshal,
			)
	}

	return w.Write(
		[]byte(
			xml.Header + string(xmloutput),
		),
	)
}
