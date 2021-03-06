package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("abnAmroExport.TAB")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[ \t]+")
	scanner := bufio.NewScanner(file)
	spent := 0.0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		replace := re.ReplaceAllString(text, " ")
		values := strings.Split(replace, " ")
		spendingStr := values[6]
		spendingStr = strings.Replace(spendingStr, ",", ".", 1)
		spending, err := strconv.ParseFloat(spendingStr, 64)
		if err != nil {
			log.Fatal(err)
		}
		if (spending < 0) {
			spent += -spending
		}
		fmt.Printf("%.2f\n", spent)
	}
	fmt.Printf("Total spent: %.2f", spent)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}