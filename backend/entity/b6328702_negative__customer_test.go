package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func Test_customer(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	cu := Customer{
		Name:       "sabaithip",
		Email:      "sabaithip@gmail.com",
		CustomerID: "C123",
	}

	ok, err := govalidator.ValidateStruct(cu)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err).ToNot(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("กรุณาใส่  CustomerID ให้ถูกต้อง"))
}
