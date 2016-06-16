package rahva_test

import (
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
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
		
		InitFakt(new(PopulationFakt))
	})

	Describe("Basic faktoid fetch", func() {
		Context("Get a random fact", func() {

			It("Should return 200", func() {
				var f *Faktoid
		
				request, _ = http.NewRequest("GET", "/faktoid", nil)

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
				request, _ = http.NewRequest("GET", "/faktoid/0143", nil)
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
	
	Describe("Get the dataset", func(){
		
		Context("Get the entire dataset", func(){
			It("Should return 200", func(){
				var f map[string]interface{}

				request, _ = http.NewRequest("GET", "/andmed", nil)
				router.ServeHTTP(recorder, request)
				Expect(recorder.Code).To(Equal(200))				

				body, err := ioutil.ReadAll(io.LimitReader(recorder.Body, bufferLength))
				Expect(err).To(BeNil())
			
				// There should be some arbitrary json available
				err = json.Unmarshal(body, &f)
				Expect(err).To(BeNil())	
			})
		})
	})
})

