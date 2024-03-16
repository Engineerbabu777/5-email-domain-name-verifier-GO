package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("domain, hasMX, hasSPF, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("err : could not read from input:", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecordss string
	mx, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mx) > 0 {
		hasMX = true
	}

	textRecord, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range textRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
		}
	}

	dmacRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, dmarcRecord := range dmacRecords {
		if strings.HasPrefix(dmarcRecord, "v=DMARC1") {
			hasDMARC = true
			dmarcRecordss = dmarcRecord
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v,", domain, hasMX, hasSPF, hasDMARC, spfRecord, dmarcRecordss)

}
