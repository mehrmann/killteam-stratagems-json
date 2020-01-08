package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Stratagem struct {
	Title             string `json:"title"`
	Subtitle          string `json:"sub"`
	Description       string `json:"desc"`
	Cost              int    `json:"cp"`
	Phase             string `json:"phase"`
	FactionKeyword    string `json:"faction,omitempty"`
	SpecialistKeyword string `json:"specialist,omitempty"`
	SpecialistLevel   int    `json:"level,omitempty"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parse(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func main() {
	flag.Parse()

	csvFile, _ := os.Open(flag.Arg(0))
	r := csv.NewReader(bufio.NewReader(csvFile))

	var stratagems []Stratagem
	r.Read()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		subtitle := "Tactic"

		if len(record[3]) > 0 {
			subtitle = fmt.Sprintf("%s Tactic", record[3])
		} else if len(record[4]) > 0 {
			subtitle = fmt.Sprintf("%s Level %s Tactic", record[4], record[5])
		}

		stratagems = append(stratagems, Stratagem{
			Title:             record[1],
			Subtitle:          subtitle,
			FactionKeyword:    record[3],
			SpecialistKeyword: record[4],
			SpecialistLevel:   parse(record[5]),
			Description:       record[6],
			Cost:              parse(record[7]),
			Phase:             record[8],
		})
	}
	//stratgemJson, _ := json.MarshalIndent(stratagems, "", "\t")
	stratgemJson, _ := json.Marshal(stratagems)
	fmt.Println(string(stratgemJson))
}
