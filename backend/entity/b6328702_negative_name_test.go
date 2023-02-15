package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name       string `valid:"required~Name Not Null"`
	Email      string `valid:"email"`
	CustomerID string `valid:"matches(^[H|L|M]\\d{7}$)~aa"`
}

func Test_NameNotNull(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	cu := Customer{
		Name:       "",
		Email:      "sabaithip@gmail.com",
		CustomerID: "H1234567",
	}

	ok, err := govalidator.ValidateStruct(cu)
	g.Expect(ok).NotTo(gomega.BeTrue())
	g.Expect(err.Error()).To(gomega.Equal("Name Not Null"))
}

// func Test_Email(t *testing.T) {
// 	g := gomega.NewGomegaWithT(t)

// 	cu := Customer{
// 		Name:       "Sabaithip",
// 		Email:      "sabaithip",
// 		CustomerId: "H1234567",
// 	}

// 	ok, err := govalidator.ValidateStruct(cu)
// 	g.Expect(ok).NotTo(gomega.BeTrue())
// 	g.Expect(err.Error()).To(gomega.Equal("Email: sabaithip does not validate as email"))
// }
