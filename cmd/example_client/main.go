package main

import "github.com/hyperledger/aries-framework-go/pkg/client/issuecredential"

func main() {

	ctx, err := createAriesAgent(parameters)
	if err != nil {
		return err
	}
	issueCredentialClient, issueCredentialClientError := issuecredential.New()
}
