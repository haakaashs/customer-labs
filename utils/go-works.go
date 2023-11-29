package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
)

func goWorker(req map[string]string, channel chan Responce) {

	defer close(channel)

	funcDesc := "goWorker"
	log.Println("entered " + funcDesc)

	var wg sync.WaitGroup

	// get high level field mapped
	response := work1(req)

	// loop through different types of attributes
	prefixArray := [2]string{"atr", "uatr"}
	var atrList []map[string]map[string]string
	for _, prefix := range prefixArray {
		// get atr atributes map
		resultMap := work2(req, prefix)
		atrList = append(atrList, resultMap)
	}

	atrChannel := make(chan map[string]map[string]string)
	uatrChannel := make(chan map[string]map[string]string)

	wg.Wait()
	go work3(atrList[0], atrChannel, &wg)
	go work3(atrList[1], uatrChannel, &wg)
	wg.Add(2)
	response.Atks = <-atrChannel
	response.Uatks = <-uatrChannel
	wg.Wait()
	channel <- response
	log.Println("exit " + funcDesc)
}

// get modified response key for attributes
func work3(atrMaps map[string]map[string]string, channel chan map[string]map[string]string, wg *sync.WaitGroup) {
	defer func() {
		close(channel)
		wg.Done()
	}()

	funcDesc := "work3"
	log.Println("entered " + funcDesc)

	outputMap := make(map[string]map[string]string)
	for _, atribute := range atrMaps {
		k := ""
		for key, value := range atribute {
			if strings.Contains(key, "atrk") {
				k = value
				outputMap[k] = make(map[string]string)
			}
		}
		for key, value := range atribute {
			if strings.Contains(key, "atrt") {
				outputMap[k]["type"] = value
			} else if strings.Contains(key, "atrv") {
				outputMap[k]["value"] = value
			}
		}
	}
	channel <- outputMap
	log.Println("exit " + funcDesc)
}

// group attributes by prefix and last digits in the key
func work2(in map[string]string, prefix string) map[string]map[string]string {

	funcDesc := "work2"
	log.Println("entered " + funcDesc)

	mapGroups := make(map[string]map[string]string)

	pattern := fmt.Sprintf(`^%s[kvt](\d+)$`, prefix)

	regexObj, err := regexp.Compile(pattern)
	if err != nil {
		err := errors.New(fmt.Errorf("error in regex compile: %s", err).Error())
		log.Println(err.Error())
		return mapGroups
	}

	for key, value := range in {
		matchSlice := regexObj.FindStringSubmatch(key)
		if len(matchSlice) > 0 {
			groupId := matchSlice[1]
			if _, ok := mapGroups[groupId]; !ok {
				mapGroups[groupId] = make(map[string]string)
			}
			mapGroups[groupId][key] = value
		}
	}
	log.Println("exit " + funcDesc)
	return mapGroups
}

// high level fields mapping
func work1(in map[string]string) Responce {

	funcDesc := "work1"
	log.Println("entered " + funcDesc)

	out := Responce{
		Ev:  in["ev"],
		Et:  in["et"],
		ID:  in["id"],
		UID: in["uid"],
		MID: in["mid"],
		T:   in["t"],
		P:   in["p"],
		L:   in["l"],
		SC:  in["sc"],
	}
	log.Println("exit " + funcDesc)
	return out
}
