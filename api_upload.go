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
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Create an upload media URL
func createUploadMediaUrl(mediaPathFileName string) string {
	fmt.Println("CALL: createUploadMediaUrl()")

	//  Build up Json File Metadata
	fileMetadata := &FileMetadata{
		Name: filepath.Base(mediaPathFileName),
		Size: determineFileSize(mediaPathFileName),
	}
	jsonFileMetadata, _ := json.Marshal(fileMetadata)

	//  Call API to create the media url for uploading
	req, err := http.NewRequest(
		http.MethodPost,
		URL_SNAPWONDERS_API+URL_UPLOAD_CREATE_MEDIA_URL,
		bytes.NewBuffer(jsonFileMetadata),
	)
	if err != nil {
		log.Fatalf("ERROR: Create 'http.NewRequest' failed:[%v]", err)
	}
	addApiHeaders(req, HTTP_CONTENT_TYPE_JSON)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("ERROR: Send POST request failed:[%v]", err)
	}

	//  Check POST status for errors
	if res.StatusCode != http.StatusCreated {
		dumpApiError("Create media url failed", res)
	}

	//  Success - Extract the media url
	urlCreateMedia := res.Header.Get("Location")
	fmt.Printf("SUCCESS: Created resumable uploading media url:[%s]\n", urlCreateMedia)
	return urlCreateMedia
}

// Uploads a data chunk
func uploadDataChunk(urlUploadMedia string, offset int, dataChunk []byte) int {

	//  Build the multipart form data for uploading
	var formData bytes.Buffer
	multipartData := multipart.NewWriter(&formData)

	//  Offset
	partDataOffset, err := multipartData.CreateFormField("offset")
	if err != nil {
		log.Fatalf("ERROR: Creating multipart 'offset' key failed:[%v]", err)
	}
	_, err = io.Copy(partDataOffset, strings.NewReader(strconv.Itoa(offset)))
	if err != nil {
		log.Fatalf("ERROR: Creating multipart 'offset' data failed:[%v]", err)
	}

	//  Chunked data
	partDataChunk, err := multipartData.CreateFormField("file")
	if err != nil {
		log.Fatalf("ERROR: Creating multipart 'file' key failed:[%v]", err)
	}
	_, err = io.Copy(partDataChunk, strings.NewReader(string(dataChunk)))
	if err != nil {
		log.Fatalf("ERROR: Creating multipart 'offset' data failed:[%v]", err)
	}
	multipartData.Close()

	//  Patch the data chunk for uploading to given media url
	req, err := http.NewRequest(
		http.MethodPatch,
		urlUploadMedia,
		&formData,
	)
	if err != nil {
		log.Fatalf("ERROR: Create 'http.NewRequest' failed:[%v]", err)
	}
	addApiHeaders(req, multipartData.FormDataContentType())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("ERROR: Send PATCH request failed:[%v]", err)
	}

	//  Check PATCH status for errors
	if res.StatusCode != http.StatusOK {
		dumpApiError("Upload data chunk failed", res)
	}

	//  Check for upload errors.
	//  Note: If an upload failed, you can retry uploading from the last offset. Call the HEAD request to determine
	//  the last offset position if you are not sure what that is. Uploading is resumable and can be continued
	//  at a later time (which is useful if there is a network outage or connectivity issue)
	//  snapWONDERS uploading follows the Tus.io protocol
	uploadOffset, err := strconv.Atoi(res.Header.Get("Upload-Offset"))
	if err != nil {
		log.Fatalf("ERROR: New `offset` extraction failed:[%v]", err)

	} else if uploadOffset != (offset + len(dataChunk)) {
		log.Fatalf(`ERROR: Uploading data chunk failed! TODO: You can retry uploading the last data chunk or 
			resume uploading at a later point in time`)
	}

	//  Success - Uploaded the data chunk
	fmt.Printf("INFO: Uploaded data chunk starting at offset:[%d], newOffset:[%d]\n", offset, uploadOffset)
	return uploadOffset
}

// Uploads file to given media url
func uploadMedia(urlUploadMedia string, mediaPathFileName string) {
	fmt.Println("CALL: uploadMedia()")

	//  Open file
	file, err := os.Open(mediaPathFileName)
	if err != nil {
		log.Fatalf("ERROR: Open file failed:[%v]", err)
	}
	defer file.Close()

	//  Set up to read chunks
	reader := bufio.NewReader(file)
	dataChunk := make([]byte, DATA_CHUNK_SIZE)

	//  Loop through each chunk and upload
	offset := 0
	for {
		readSize, err := reader.Read(dataChunk)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("ERROR: Read file failed:[%v]", err)
			}
			break
		}

		offset = uploadDataChunk(urlUploadMedia, offset, dataChunk[0:readSize])
	}

	fmt.Printf("SUCCESS: Uploaded file:[%s] to media url:[%s]\n", mediaPathFileName, urlUploadMedia)
}
