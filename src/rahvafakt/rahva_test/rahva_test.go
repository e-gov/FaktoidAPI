package rahva_test

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
	. "rahvafakt"
	. "faktoid"
	"io/ioutil"
	"io"
	"encoding/json"
)

var bufferLength int64 = 1048576

var _ = Describe("RahvaService", func() {
	var router *mux.Router
	var recorder *httptest.ResponseRecorder
	var request *http.Request

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		router = NewRouter()
		
		u, _ := url.Parse("/faktoid")
		request, _ = http.NewRequest("GET", u.RequestURI(), nil)
		InitFakt(new(PopulationFakt))
	})

	Describe("Basic faktoid fetch", func() {
		Context("Get a random fact", func() {

			It("Should return 200", func() {
				var f *Faktoid
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))

				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())
							
				f = new(Faktoid)
				err = json.Unmarshal(body, f)
				Expect(err).To(BeNil())

			})
		})
	})
	
	Describe("Filtered fakotid fetch", func(){
		
		Context("Get a random fact using EHAK code", func(){
			It("Should return 200", func(){
				var f *Faktoid
				request, _ = http.NewRequest("GET", "/faktoid/0037", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))

				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())
				
				f = new(Faktoid)
				err = json.Unmarshal(body, f)
				Expect(err).To(BeNil())
			})
		})
		
	})
	
})

