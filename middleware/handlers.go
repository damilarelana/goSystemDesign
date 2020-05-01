package middleware

import (
	"github.com/damilarelana/goSystemDesign/accountService/dbClients"
)

// DBClient references the boltDB interface i.e. making it ready for  use
var DBClient dbClients.InterfaceBoltClient
