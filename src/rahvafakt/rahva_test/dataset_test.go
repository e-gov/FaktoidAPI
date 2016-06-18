package rahva_test

import(
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "rahvafakt"
	"EHAK"
)

var _ = Describe("Dataset test", func() {
	var dataF = "RV0241_utf.csv"
	var ehakF = "EHAK2015v1.txt"
	
	Describe("Basic dataset handling" , func(){
		Context("Loading", func(){
			It("Should get some lines", func(){
				e, err := EHAK.Load(ehakF)
				Expect(err).To(BeNil())

				r := LoadData(dataF, e)
				Expect(len(*r)).To(Not(Equal(0)))				
			})
		})
	})
})

