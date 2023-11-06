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
	"fmt"
)

const (

	//  NOTE: YOU MUST SET YOUR SNAPWONDERS API KEY BEFORE START (visit https://snapwonders.com for details)
	SNAPWONDERS_API_KEY = "<put-your-api-key-here>"

	//  Full path/filename to your media (this can be an image or video)
	MEDIA_PATH_FILENAME = "./Data/Images/me-mountain.jpg"

	//  Uploads are done via resumable uploading in chunks (this must be less than 5MB)
	DATA_CHUNK_SIZE = 400000
)

func main() {
	fmt.Println("snapWONDERS Client OpenAPI v3 Go Example!")
	fmt.Println("You must set your API key and media path/filename")

	//  Create an analyse job and display results
	analyseJob(MEDIA_PATH_FILENAME)
}
