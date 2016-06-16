package rahva_test

import(
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "rahvafakt"
)

var _ = Describe("UtilTest", func() {
	Describe("Stack", func(){
		var v *Stack
		BeforeEach(func(){
			v = new(Stack)
		})
		
		Context("Standard operations", func(){
			It("Should work",func(){
				s := "foo"
				v.Push(s)
				Expect(len(*v)).To(Equal(1))
				Expect(len(*v.Content())).To(Equal(1))
				Expect(v.Pop()).To(Equal(s))
				Expect(v.Pop()).To(Equal(""))
				
			})
		})
	})
	
	Describe("Count dots", func(){
		Context("Standard ops", func(){
			It("Should work", func(){
				Expect(CountDots("")).To(Equal(0))
				Expect(CountDots("A")).To(Equal(0))
				Expect(CountDots(".")).To(Equal(1))
				Expect(CountDots("..")).To(Equal(2))
				Expect(CountDots(".A")).To(Equal(1))
				Expect(CountDots("A.")).To(Equal(0))
			})
		})
	})	
})