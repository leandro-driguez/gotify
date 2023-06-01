package kademlia

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"

	b58 "github.com/jbenet/go-base58"
)

type Tracker struct {
	dht *DHT
}

func NewTracker(ip string, port int, bootIp string, bootPort int) (*Tracker, error) {
	tracker := &Tracker{}
	id, _ := getNewID()
	var err error = nil
	tracker.dht, err = NewDHT(getNewInMemoryStore(), &Options{
		ID:   id,
		IP:   ip,
		Port: strconv.Itoa(port),
		BootstrapNodes: []*NetworkNode{
			NewNetworkNode(bootIp, strconv.Itoa(bootPort)),
		},
	})
	return tracker, err
}

func (t *Tracker) GetSongList(key string) []string {
	songList := []string{}

	flatArray, found, err := t.dht.Get(key)
	if err != nil {
		fmt.Println("Error when retrieving data:", err)
	}

	if found {
		formatedArray := getFormatedArray(flatArray)
		songList = getStringSliceFromByteArray(formatedArray)
	}

	return songList
}

func (t *Tracker) StoreSongMetadata(jsonSongMetadata string, songData []byte) []string {
	hashesPowerSet := GetHashesPowerSet(jsonSongMetadata)
	songDataHash := sha1.Sum(songData)
	value := songDataHash[:]

	for _, hash := range hashesPowerSet {
		key := []byte(hash)
		id, err := t.dht.Store(key, value)
		if err != nil {
			fmt.Println("Error when storing key:", id, err)
		}
	}

	return hashesPowerSet
}

func getNewInMemoryStore() *MemoryStore {
	memStore := &MemoryStore{}
	return memStore
}

// newID generates a new random ID
func getNewID() ([]byte, error) {
	result := make([]byte, 20)
	_, err := rand.Read(result)
	return result, err
}

func getFormatedArray(flatArray []byte) [][20]byte {
	result := [][20]byte{}
	lenght := len(flatArray)
	currentElem := [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < lenght; i++ {
		pos := i % 10
		if pos == 0 {
			result = append(result, currentElem)
			currentElem := [20]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
			currentElem[pos] = flatArray[i]
		} else if pos < lenght-1 {
			currentElem[pos] = flatArray[i]
		} else {
			currentElem[pos] = flatArray[i]
			result = append(result, currentElem)
		}
	}

	return result
}

func getStringSliceFromByteArray(array [][20]byte) []string {
	result := []string{}
	for _, byteArray := range array {
		str := string(byteArray[:])
		result = append(result, str)
	}
	return result
}

func GetStringKeyFromRawJson(jsonSongMetadata string) string {
	subJson := make(map[string]interface{})
	jsonFromMap, err := json.Marshal(subJson)
	if err != nil {
		fmt.Printf("could not marshal map: %s\n", err)
	}
	jsonHash := sha1.Sum(jsonFromMap)
	resultHash := jsonHash[:]
	hash := b58.Encode(resultHash)
	return hash
}