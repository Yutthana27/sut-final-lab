package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBooks1(t *testing.T){
	g := NewGomegaWithT(t)

	book := Book{
		Title: "yutthana",
		Price: 550,
		Code: "BK123456",
	}

	ok, err := govalidator.ValidateStruct(book)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}