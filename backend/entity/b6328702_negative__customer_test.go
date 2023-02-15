package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func Test_(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	cu := Customer{
		Name:       "sabaithip",
		Email:      "sabaithip@gmail.com",
		CustomerID: "C1234567",
	}

	ok, err := govalidator.ValidateStruct(cu)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err.Error()).To(gomega.Equal("aa"))
}
