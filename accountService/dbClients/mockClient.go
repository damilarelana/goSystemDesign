package dbClients

import (
	m "github.com/damilarelana/goSystemDesign/accountService/model"
	"github.com/stretchr/testify/mock"
)

/*
MockBoltClient is a mock implementation of the InterfaceBoltClient for testing purposes
- instead of using BoltClient (via *bolt.DB) which implements the database interface
- we simply use a generic mock object from stretchr/testify
*/
type MockBoltClient struct {
	mock.Mock
}

// OpenBoltDB is a method that implements interfaceBoltClient
// - implements actual connection to the boltDB Instance
func (mbc *MockBoltClient) OpenBoltDB() {
	// ErrMsgHandler("Unable to open a Mock connection", err)
}

// QueryAccount is a method that implements interfaceBoltClient
// - takes in accountID as argument
// - retrieves matching test Accounts
// - returns matched account struct and error value
func (mbc *MockBoltClient) QueryAccount(accountID string) (m.Account, error) {
	args := mbc.Mock.Called(accountID)
	return args.Get(0).(m.Account), args.Error(1)
}

// Seed ...
func (mbc *MockBoltClient) Seed() {

}
