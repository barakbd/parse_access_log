// Run for debugging puposes:
// > go run parse_log.go access.log
// Build and run:
// > go build parse_log.go
// > ./parse_log access.log
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"encoding/json"
	"regexp"
) 

func main() { 

	var input_file_path string  = os.Args[1]
	log.Println(input_file_path)
	status_code_count_map, err := parse_log_file_to_map(input_file_path)
	if err != nil {
		log.Fatal(err)
	}
	status_code_count_json, err := json.MarshalIndent(status_code_count_map, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(status_code_count_json))
}

func parse_log_file_to_map(input_file_path string) (map[string]int, error) {	
	// Open the file
	input_file, err := os.Open(input_file_path) 

	if err != nil { 
		log.Println("failed to open", input_file_path)
		return map[string]int{}, err
	}
	 
	// Pass the input_file (os.File object)  to bufio.NewScanner()
	scanner := bufio.NewScanner(input_file) 

	// Split the lines supplied by bufio.ScanLines
	scanner.Split(bufio.ScanLines)

	// Iterate over each line and append to a list
	var text_list []string 
	for scanner.Scan() { 
		text_list = append(text_list, scanner.Text()) 
	} 

	// close the file
	defer input_file.Close() 

	// Create a map to store  {status_code: count}
	var status_code_count_map map[string]int = make(map[string]int)

	// Create a regex which will be uised to split each line
	tab := regexp.MustCompile(`\t`)

	// Loop the text_list
	for _, each_ln := range text_list {
		// Split each line by tab, without a limit of splits
		split_line := tab.Split(each_ln, -1)
		// Capture the status_code in the 6 column
		var status_code string = split_line[5]
		if status_code_count_map[status_code] == 0 {
			status_code_count_map[status_code] = 1
		} else {
			status_code_count_map[status_code]++
		}
	} // end for

	// Debug - print key: value
	// for key, value := range status_code_count_map {
    //     fmt.Printf("%s:\t%v\n", key, value)
    // }
	return status_code_count_map, nil
}