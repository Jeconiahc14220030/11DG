package models

import "log"

// InsertLogError mencatat log error atau informasi eksekusi fungsi
func InsertLogError(ip string, functionName string, err error) {
	if err != nil {
		log.Printf("ERROR: [%s] Function: %s, Error: %s", ip, functionName, err.Error())
	} else {
		log.Printf("INFO: [%s] Function: %s executed successfully", ip, functionName)
	}
}
