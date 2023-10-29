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
	"log"
	"os"
)

// Determine total file upload size
func determineFileSize(fileName string) int64 {
	fileStat, err := os.Stat(fileName)
	if err != nil {
		log.Fatalf("ERROR: Failed to stat file:[%v]", err)
	}

	fileTotalSize := fileStat.Size()
	if fileTotalSize <= 0 {
		log.Fatalf("ERROR: Illegal file size:[%d]", fileTotalSize)
	}
	return fileTotalSize
}
