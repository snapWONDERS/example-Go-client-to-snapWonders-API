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

// The Analyse preferences
type JobAnalyse struct {
	// The key contained within the *media-url*
	Key string `json:"key"`
	// Enable extra tips where available
	EnableTips bool `json:"enableTips"`
	// Enable extra deep anlaysis where available
	EnableExtraAnalysis bool `json:"enableExtraAnalysis"`
}
