package main

import (
	"fmt"

	"github.com/antchfx/xmlquery"
	// _ "github.com/dgrijalva/jwt-go"
	// _ "github.com/gogo/protobuf/proto"

	// // - "github.com/gogs/gogs"
	// _ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

func main() {
	fmt.Println("AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE")
	fmt.Println("AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")
	fmt.Println("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEz")
	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())
}
