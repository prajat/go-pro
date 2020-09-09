package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	//example1
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	//example2
	var err error
	log.WithFields(log.Fields{
		"error":       err,
		"api handler": "FetchFollowersList",
	}).Error("while converting string to int for user_id")

}
