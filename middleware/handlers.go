package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/damilarelana/goSystemDesign/accountService/dbClients"
	"github.com/gorilla/mux"
)

// DBClient references the boltDB interface i.e. making it ready for  use
var DBClient dbClients.InterfaceBoltClient

// GetAccount defines the HandlerFunc required by Route struct
func GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID := mux.Vars(r)["accountID"]            // extract accountID from path
	account, err := DBClient.QueryAccount(accountID) // read the account from boltDB. This occurs only after `initializeBoltClient()` in main.go, thus openDB, seed() etc. to give us a DBClient to use
	if err != nil {                                  // return a 404 when we are unable to query the account
		w.WriteHeader(http.StatusNotFound)
		return
	}
	accountJSON, err := json.Marshal(account) // converts the type struct data `account` back to json for display in a browser. type struct is only useful for internal golang
	ErrMsgHandler("Unable to convert account's 'struct data' to 'json data'", err)

	// set the request headers metadata
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(accountJSON)))
	w.WriteHeader(http.StatusOK)
	w.Write(accountJSON)
}
