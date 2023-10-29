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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// Create an analyse job and display results
func analyseJob(pathFileName string) {

	//  Create upload media url and upload file in data chunks
	urlUploadMedia := createUploadMediaUrl(pathFileName)
	uploadMedia(urlUploadMedia, pathFileName)

	//  Create an analyse job url
	urlJobStatus := createAnalyseJob(urlUploadMedia)

	//  Track the job status, wait until analyse job is completed
	jsonJobStatus := &JobStatus{}
	for {
		jsonJobStatus = getJobStatus(urlJobStatus)

		//  Wait for the job to be completed
		if (jsonJobStatus.Status == JOB_STATUS_WAITING) || (jsonJobStatus.Status == JOB_STATUS_PROCESSING) {
			time.Sleep(5 * time.Second)
			fmt.Println("INFO: Sleeping for a few seconds...")

			//  If completed we break out
		} else if jsonJobStatus.Status == JOB_STATUS_COMPLETED {
			break

			//  Some unknown state?
		} else {
			log.Fatalf("ERROR: Analyse job failed with status:[%s], message:[%s]", jsonJobStatus.Status, jsonJobStatus.Message)
		}
	}

	//  Get and display results
	dataResults := getJobResults(jsonJobStatus.ResultUrl)
	var mapJsonResults map[string]interface{}
	json.Unmarshal(dataResults, &mapJsonResults)

	//  NOTE: You can call getJobResults() for image url resouces contained within the JSON result
	printPretty(mapJsonResults)
}

// Create an analyse job
func createAnalyseJob(urlUploadMedia string) string {
	fmt.Println("CALL: createAnalyseJob()")

	//  Build up Json Analyser Job
	jobAnalyse := &JobAnalyse{
		Key:                 filepath.Base(urlUploadMedia),
		EnableTips:          true,
		EnableExtraAnalysis: true,
	}
	jsonJobAnalyse, _ := json.Marshal(jobAnalyse)

	//  Call API to create an analyse job
	req, err := http.NewRequest(
		http.MethodPost,
		URL_SNAPWONDERS_API+URL_JOB_CREATE_ANALYSE,
		bytes.NewBuffer(jsonJobAnalyse),
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
	if res.StatusCode != http.StatusOK {
		dumpApiError("Create analyse job failed", res)
	}

	//  Success - Extract the media url
	urlJobAnalyse := res.Header.Get("Location")
	fmt.Printf("SUCCESS: Created analyse job located at url:[%s]\n", urlJobAnalyse)
	return urlJobAnalyse
}

// Gets the job status
func getJobStatus(urlJobStatus string) *JobStatus {
	fmt.Println("CALL: getJobStatus()")

	//  Call API to get job status
	req, err := http.NewRequest(
		http.MethodPost,
		urlJobStatus,
		nil)
	if err != nil {
		log.Fatalf("ERROR: Create 'http.NewRequest' failed:[%v]", err)
	}
	addApiHeaders(req, HTTP_CONTENT_TYPE_JSON)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("ERROR: Send POST request failed:[%v]", err)
	}

	//  Check POST status for errors
	if res.StatusCode != http.StatusOK {
		dumpApiError("Get job status failed", res)
	}

	//  Success - Extract job status
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR: Reading response body failed:[%v], status:[%s]", err, res.Status)
	}

	var jsonJobStatus = &JobStatus{}
	json.Unmarshal([]byte(resBody), &jsonJobStatus)
	fmt.Printf("SUCCESS: Have job status:[%s]\n", jsonJobStatus.Status)
	return jsonJobStatus
}

// Gets the job results (this can be a JSON or image content)
func getJobResults(urlJobResults string) []byte {
	fmt.Println("CALL: getJobStatus()")

	//  Call API to get job status
	req, err := http.NewRequest(
		http.MethodGet,
		urlJobResults,
		nil)
	if err != nil {
		log.Fatalf("ERROR: Create 'http.NewRequest' failed:[%v]", err)
	}
	addApiHeaders(req, HTTP_CONTENT_TYPE_JSON)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("ERROR: Send GET request failed:[%v]", err)
	}

	//  Check GET status for errors
	if res.StatusCode != http.StatusOK {
		dumpApiError("Get job status failed", res)
	}

	//  Success - Extract the job results
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ERROR: Reading response body failed:[%v], status:[%s]", err, res.Status)
	}

	fmt.Printf("SUCCESS: Have results with data size:[%d]\n", len(resBody))
	return resBody
}
