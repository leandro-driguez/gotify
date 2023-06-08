package utils

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func GetJsonMetadataKeyHash(key string) string {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(key), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
	}

	jsonBytes, _ := json.Marshal(data)
	jsonStrFromMap := string(jsonBytes)
	if err != nil {
		fmt.Printf("could not marshal map: %s\n", err)
	}
	jsonHash := sha1.Sum([]byte(jsonStrFromMap))
	resultHash := jsonHash[:]
	hash := base64.RawStdEncoding.EncodeToString(resultHash)
	return hash
}
