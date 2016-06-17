package ehak_test

import (
	. "EHAK"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ehak_test", func() {
	var ehakF = "EHAK2015v1.txt"

	Describe("EHAK loading", func() {
		Context("Load a EHAK file", func() {
			It("Should return a meaningful array", func() {
				e, err := Load(ehakF)
				Expect(err).To(BeNil())
				Expect(*e).To(Not(BeEmpty()))
			})

			It("Should fail, no file", func() {
				_, err := Load("not there at all")
				Expect(err).To(Not(BeNil()))
			})
		})
	})

	Describe("EHAK parsing", func() {
		Context("Parse the entire EHAK dataset", func() {
			It("As many lines out as goes in", func() {
				_, err := Load(ehakF)
				Expect(err).To(BeNil())
			})
		})

		Context("Find a particular unit", func() {
			var ehak *[]string
			var err error

			BeforeEach(func() {
				ehak, err = Load(ehakF)
				Expect(err).To(BeNil())
			})
			It("Get a unit by name", func() {
				u := GetUnitByName("Tallinn", ehak)
				Expect(u).To(Not(BeNil()))
			})

			It("Get a unit by code", func() {
				u := GetUnitByCode("0129", ehak)
				Expect(u).To(Not(BeNil()))
			})
		})

		Context("Find by array", func() {
			var ehak *[]string
			var err error

			BeforeEach(func() {
				ehak, err = Load(ehakF)
				Expect(err).To(BeNil())
			})

			It("Test exact match", func() {
				u := GetUnitByArray([]string{"Tallinn"}, ehak)
				Expect(u.Name).To(Equal("Tallinn"))

				u = GetUnitByArray([]string{"Helsingi"}, ehak)
				Expect(u).To(BeNil())
			})

			It("Test multiple array items", func() {
				u := GetUnitByArray([]string{"Tamme küla"}, ehak)
				Expect(u.Name).To(Equal("Tamme küla"))

				u = GetUnitByArray([]string{"Orava vald", "Tamme küla"}, ehak)
				Expect(u.Code).To(Equal("7326"))

				u = GetUnitByArray([]string{"Helsingi", "Tamme küla"}, ehak)
				Expect(u.Name).To(Equal("Tamme küla"))

				u = GetUnitByArray([]string{"Pärnu maakond", "Koonga vald", "Tamme küla"}, ehak)
				Expect(u.Code).To(Equal("8092"))

				u = GetUnitByArray([]string{"Pärnu maakond", "Helsingi", "Tamme küla"}, ehak)
				Expect(u.Name).To(Equal("Tamme küla"))

			})

		})
	})
})
