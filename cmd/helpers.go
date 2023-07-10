package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func readConfigJSON(dst interface{}) error {
	val, err := os.Open("/Users/cuneytakkurt/workspace/ssh-conn/clusters.json")
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(val)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &dst)
	if err != nil {
		return err
	}

	return nil
}
