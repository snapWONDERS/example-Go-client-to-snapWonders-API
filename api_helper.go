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
	"io"
	"log"
	"net/http"
)

const (

	//  The API URLs
	URL_SNAPWONDERS_API         = "https://api.snapwonders.com/v1/"
	URL_UPLOAD_CREATE_MEDIA_URL = "upload/create-media-url"
	URL_JOB_CREATE_ANALYSE      = "job/analyse"

	//  Content types
	HTTP_CONTENT_TYPE_JSON = "application/json"

	//  Job status
	JOB_STATUS_WAITING    = "WAITING"
	JOB_STATUS_PROCESSING = "PROCESSING"
	JOB_STATUS_COMPLETED  = "COMPLETED"
)

// Dumps API error and exit
func dumpApiError(what string, res *http.Response) {
	resStatus := res.Status
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR: Reading response body failed:[%v], status:[%s]", err, resStatus)
	}

	var jsonResBody map[string]interface{}
	json.Unmarshal([]byte(resBody), &jsonResBody)
	resMessage, ok := jsonResBody["message"]
	if !ok {
		resMessage = resBody
	}

	log.Fatalf("ERROR: %s:[%s], status:[%s]", what, resMessage, resStatus)
}

// Adds the headers
func addApiHeaders(req *http.Request, contentType string) {
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Api_key", SNAPWONDERS_API_KEY)
}
