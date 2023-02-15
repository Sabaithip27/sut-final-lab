package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func Test_Pass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	cu := Customer{
		Name:       "sabaithip",
		Email:      "sabaithip@gmail.com",
		CustomerID: "H1234567",
	}

	ok, err := govalidator.ValidateStruct(cu)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())
}
