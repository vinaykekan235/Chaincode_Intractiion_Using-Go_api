package web

import (
	"fmt"
	"net/http"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/rs/cors"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
}

// Serve starts http web server.
func Serve(setups OrgSetup) {
	// Create a new CORS middleware
	c := cors.Default()

	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/invoke", setups.Invoke)

	// Use the CORS middleware to handle CORS requests
	handler := c.Handler(http.DefaultServeMux)

	fmt.Println("Listening (http://localhost:3000/)...")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		fmt.Println(err)
	}
}
