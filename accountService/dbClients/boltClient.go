package dbClients

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	m "github.com/damilarelana/goSystemDesign/accountService/model"
	bolt "github.com/etcd-io/bbolt"
)

// ErrMsgHandler defines the error message handler
func ErrMsgHandler(msg string, err error) {
	if err != nil {
		log.Println(msg, err.Error())
		os.Exit(1)
	}
}

// InterfaceBoltClient defines an interface applicable to connecting to boltDB
type InterfaceBoltClient interface {
	OpenBoltDB()                                      // opens a BoltDB connection
	QueryAccount(accountID string) (m.Account, error) // returns the composite variable Account
	Seed()
}

// BoltClient defines structure of a typical boltDB instance via a pointer to a bolt.DB instance
// - pointer is used because we do not want a copy of the bolt.DB instance
// - we want the actual instance
// - so as to be sure that changes are persisted in that actual DB
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDB is a method that implements interfaceBoltClient
// - implements actual connection to the boltDB Instance
func (bc *BoltClient) OpenBoltDB() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	ErrMsgHandler("Unable to open a BoltDB connection", err)
}

// Seed is a method that implements interfaceBoltClient
// - handles the populating the database with test Accounts
func (bc *BoltClient) Seed() {
	bc.initializeBucket() // call the initializeBucket() method
	bc.seedAccounts()     // call the seedAccounts() method
	// ErrMsgHandler("Unable to update boltDB with account json data", err)
}

// QueryAccount is a method that implements interfaceBoltClient
// - takes in accountID as argument
// - retrieves matching test Accounts
// - returns matched account struct and error value
func (bc *BoltClient) QueryAccount(accountID string) (m.Account, error) {
	account := m.Account{}                          // initialize the model Account struct
	err := bc.boltDB.View(func(tx *bolt.Tx) error { // read the 'AccountBucket' object
		b := tx.Bucket([]byte("AccountBucket"))
		accountBytes := b.Get([]byte(accountID)) // search and get the specific account with key value `accountID`
		if accountBytes == nil {
			return fmt.Errorf("No account found with accountID: " + accountID)
		}
		json.Unmarshal(accountBytes, &account) // unmarshal the matched account json (into a struct format) into the memory address for account struct
		return nil                             // the anonymous function returns a nil error value
	})
	if err != nil {
		return m.Account{}, err // return the empty struct and the err value
	}
	return account, nil // return the matched account and nil as err
}

// initializeBucket creates an "AccountBucket" in BoltDB
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		ErrMsgHandler("Unable to create Account Bucket.", err)
		return nil
	})
}

// seedAccounts "n" number of fake accounts
func (bc *BoltClient) seedAccounts() {

	numOfAccounts := 100
	for indexCounter := 0; indexCounter < numOfAccounts; indexCounter++ {
		key := strconv.Itoa(10000 + indexCounter) // generate an index string value

		account := m.Account{ // create an account by instantiating with the key
			ID:   key,
			Name: "Individual_" + strconv.Itoa(indexCounter),
		}

		jsonBytes, err := json.Marshal(account) // convert the struct data to json format in prep for update to the BoltDB
		ErrMsgHandler("Unable to serialize account struct.", err)

		bc.boltDB.Update(func(tx *bolt.Tx) error { // connect to boltDB and update the account bucket with new account using the function
			b := tx.Bucket([]byte("AccountBucket"))
			err = b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Successfully created %v fake accounts (in BoltDB) ... \n", numOfAccounts)
}
