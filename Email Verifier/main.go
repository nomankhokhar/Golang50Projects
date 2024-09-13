package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func checkDomain (domain string){
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Fatal("Error: %v\n", err)
	}


	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Println("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if record == "v=spf1" {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords , err :=  net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Println("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if record == "v=DMARC1" {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	} 

	fmt.Println(domain + "," + fmt.Sprintf("%t", hasMx) + "," + fmt.Sprintf("%t", hasSPF) + "," + spfRecord + "," + fmt.Sprintf("%t", hasDMARC) + "," + dmarcRecord)
}