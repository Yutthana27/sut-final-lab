package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBooks2(t *testing.T){
	g := NewGomegaWithT(t)

	book := Book{
		Title: "yutthana",
		Price: 5,
		Code: "BK123456",
	}

	ok, err := govalidator.ValidateStruct(book)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("Price must be between 50 and 5000"))

}