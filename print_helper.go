/*
 * snapWONDERS OpenAPI Specification
 * API version: 1.0
 *
 * Copyright (c) snapWONDERS.com, All rights reserved 2023
 *
 * Author: Kenneth Springer (https://kennethbspringer.au)
 *
 * All the snapWONDERS API services is available over the Clearnet / **Web** and Dark Web **Tor** and **I2P**
 * Read details: https://snapwonders.com/snapwonders-openapi-specification
 *
 */
package main

import (
	"encoding/json"
	"fmt"
)

// Pretty print helper for JSON Map
func printPretty(obj interface{}) (err error) {
	elem, err := json.MarshalIndent(obj, "", "  ")
	if err == nil {
		fmt.Println(string(elem))
	}
	return
}
