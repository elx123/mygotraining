package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
)

// this is for config.json
type names struct {
	Name  []string
	ALL   string
	Split string
	/*write RawQuery for all or not. 0 is off ,1 is on*/
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	config_file := flag.String("c", "./config.json", "config file")
	source_file := flag.String("s", "./url_param.log", "source file")
	output_file := flag.String("o", "./output.txt", "output file")

	flag.Parse()

	var cmd string = flag.Arg(0)

	fmt.Printf("action   : %s\n", cmd)
	fmt.Printf("config_file: %s\n", *config_file)
	fmt.Printf("source_file : %s\n", *source_file)
	fmt.Printf("output_file : %s\n", *output_file)

	fmt.Printf("-------------------------------------------------------\n")

	fmt.Printf("there are %d non-flag input param\n", flag.NArg())

	for i, param := range flag.Args() {
		fmt.Printf("#%d    :%s\n", i, param)
	}

	var key_name names
	var f *os.File
	var err4 error
	var index int
	index = -1
	var symbol string
	var find_key_name_ALL bool
	find_key_name_ALL = false

	data, err1 := ioutil.ReadFile(*config_file)
	if err1 != nil {
		log.Fatalf("Read file error %v", err1)
	}
	err2 := json.Unmarshal(data, &key_name)
	if err2 != nil {
		log.Fatalf("Unmarshal json error %v", err2)
	}
	if len(key_name.Split) == 0 {
		symbol = ","
	} else {
		symbol = key_name.Split
	}
	for i := 0; i < len(key_name.Name); i++ {
		if key_name.Name[i] == key_name.ALL {
			find_key_name_ALL = true
			break
		}
	}
	if len(key_name.ALL) > 0 {
		index = 0
		for i := 0; i < len(key_name.Name); i++ {
			if key_name.Name[i] == key_name.ALL {
				index = i
				break
			}
		}
	}
	data2, err3 := ioutil.ReadFile(*source_file)
	if err3 != nil {
		log.Fatalf("Read file error %v", err3)
	}

	f, err4 = os.OpenFile(*output_file, os.O_RDWR|os.O_CREATE, 0755)
	if err4 != nil {
		log.Fatalf("OpenFile error %v", err4)
	}
	defer f.Close()
	temp_reader := bytes.NewReader(data2)
	input := bufio.NewScanner(temp_reader)
	var temp_strings string

	for input.Scan() {
		xx := strings.Contains(input.Text(), "=")
		var temp string
		var err5 error
		if xx { //In order to exclude the queryparameters
			index_int := strings.Index(input.Text(), "=")
			temp, err5 = url.QueryUnescape((input.Text())[index_int+1:])
			if err5 != nil {
				log.Fatalf("QueryUnescape error %v", err5)
			}
		} else {
			temp, err5 = url.QueryUnescape(input.Text())
			if err5 != nil {
				log.Fatalf("QueryUnescape error %v", err5)
			}
		}
		temp2 := strings.Split(temp, "&")
		for _, it := range temp2 { //In order to exclude the bug such as abc contains bc
			for _, it2 := range key_name.Name {
				xx := strings.Contains(it, it2+"=")
				if xx {
					index_int := strings.Index(it, it2+"=")
					if index_int == 0 {
						temp_strings += it + symbol
					}
				}
			}
		}
		temp3 := strings.Split(temp_strings, symbol)
		temp_temp_temp := make([]string, len(key_name.Name))
		var temp_bool bool
		for i, it := range key_name.Name {
			for _, it2 := range temp3 {
				temp_bool = strings.Contains(it2, it+"=")
				if temp_bool {
					index_int := strings.Index(it2, it+"=")
					if index_int == 0 {
						temp_temp_temp[i] = it2
						break
					}
				}
			}
		}
		temp_strings = ""
		if index != -1 && find_key_name_ALL {
			temp_temp_temp[index] = input.Text()
		}
		for i, _ := range temp_temp_temp {
			if len(temp_temp_temp[i]) == 0 {
				temp_strings += symbol
			} else {
				temp_strings += temp_temp_temp[i] + symbol
			}
		}

		if len(temp_strings) != 0 {
			/*
				if len(key_name.Split) == 0 {
					temp_strings = strings.TrimRight(temp_strings, ",")
				} else {
					temp_strings = strings.TrimRight(temp_strings, key_name.Split)
				}
			*/
			f.WriteString(temp_strings + "\n")
		}
		temp_strings = ""
	}
}
