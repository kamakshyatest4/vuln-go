package main

import (
	"fmt"

	"github.com/antchfx/xmlquery"
	// _ "github.com/dgrijalva/jwt-go"
	// _ "github.com/gogo/protobuf/proto"

	// // - "github.com/gogs/gogs"
	// _ "github.com/hashicorp/golang-lru"
	// _ "github.com/owncast/owncast/logging"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE")
	fmt.Println("AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

	fmt.Println("AIRTABLE_API_KEY=keyv2AbCdEfGhIjKlMnOpQrStUvWxYz123")
	fmt.Println("ANTHROPIC_API_KEY=sk-ant-api-0123456789abcdefABCDEF12345678")
	fmt.Println("ASANA_CLIENT_ID=1200123456789012")
	fmt.Println("ASANA_CLIENT_SECRET=0f9a3a5d2b8e4f3a95a2ef1234567890")
	fmt.Println("CLOUDFLARE_API_KEY=cf_api_key_1234567890abcdef1234567890abcdef")
	fmt.Println("CURSOR_API_KEY=cursor_sk_abcdefghijklmnopqrstuvwxyz123456")

	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())
}
