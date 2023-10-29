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

// The job status and result-url
type JobStatus struct {
	// The key contained within the *job-url*
	Key string `json:"key"`
	// Displays the job status
	Status string `json:"status"`
	// Displays message information regarding the job status (if any)
	Message string `json:"message"`
	// When the job is successfully completed, this will contain a URL to the results in JSON form
	ResultUrl string `json:"resultUrl"`
}
