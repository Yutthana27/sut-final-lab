package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBooks3(t *testing.T){
	g := NewGomegaWithT(t)

	book := Book{
		Title: "yutthana",
		Price: 550,
		Code: "12345689",
	}

	ok, err := govalidator.ValidateStruct(book)

	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("Code must start with BK followed by 6 digits (0-9)"))

}