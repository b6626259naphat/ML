	package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega" 
)

func TestProductValidation(t *testing.T) {
	g := NewGomegaWithT(t) 

	t.Run("Check Valid Product", func(t *testing.T) {
		p := Product{
			Name:  "Coke",
			Price: "20",
			SKU:   "P12345",
		}
		
		ok, err := govalidator.ValidateStruct(p)
		
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run("Check Invalid SKU", func(t *testing.T) {
		p := Product{
			Name:  "Pepsi",
			Price: "20",
			SKU:   "X99999", 
		}

		ok, err := govalidator.ValidateStruct(p)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid SKU format"))
	})
}
