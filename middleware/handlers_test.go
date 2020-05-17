package middleware

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/damilarelana/goSystemDesign/accountService/dbClients"
	m "github.com/damilarelana/goSystemDesign/accountService/model"
	. "github.com/smartystreets/goconvey/convey"
)

/*
This asserts for ensuring we get an HTTP 404 given a request of an unknown path for a given Router
 - for wrong path "/invalid/123"
*/
func TestGetAccountWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

/*
This asserts for ensuring we get an HTTP 200 given a request for a known path for a given Router
*/
func TestGetAccount(t *testing.T) {
	// create mock data
	mockDB := &dbClients.MockBoltClient{} // create mock instance of the InterfaceBoltClient

	mockDB.On("QueryAccount", "123").Return(m.Account{ // create mock data retrieval behaviour for an existing accountID "123"
		ID:   "123",
		Name: "Person_123",
	}, nil)

	mockDB.On("QueryAccount", "419").Return(m.Account{ // create mock data retrieval behaviour for a non-existing accountID "419"
	}, fmt.Errorf("Some error"))

	// assign the mocked DB interface "mockRepo" to the DBClient (previously declared in handler.go)
	DBClient = mockDB

	// test with the mock data
	Convey("Given a HTTP request for /accounts/123", t, func() { // test correct account ID
		req := httptest.NewRequest("GET", "/accounts/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200) // test to see that response was correct and appropriate

				account := m.Account{} // start using the mock data
				json.Unmarshal(resp.Body.Bytes(), &account)
				So(account.ID, ShouldEqual, "123")
				So(account.Name, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /accounts/419", t, func() { // test wrong account ID
		req := httptest.NewRequest("GET", "/accounts/419", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
