package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/alexeyco/simpletable"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Domain"},
			{Align: simpletable.AlignCenter, Text: "Has MX"},
			{Align: simpletable.AlignCenter, Text: "Has SPF"},
			{Align: simpletable.AlignCenter, Text: "SPF Record"},
			{Align: simpletable.AlignCenter, Text: "Has DMARC"},
			{Align: simpletable.AlignCenter, Text: "DMARC Record"},
		},
	}

	for scanner.Scan() {
		domain := scanner.Text()
		row := checkDomain(domain)
		table.Body.Cells = append(table.Body.Cells, row)
		table.SetStyle(simpletable.StyleCompactLite) 
		fmt.Println(table.String())     
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
       
}

func checkDomain(domain string) []*simpletable.Cell {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// MX Records Check
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else if len(mxRecords) > 0 {
		hasMx = true
	}

	// SPF Records Check
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range txtRecords {
		if len(record) >= 6 && record[:6] == "v=spf1" {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// DMARC Records Check
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if len(record) >= 8 && record[:8] == "v=DMARC1" {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	return []*simpletable.Cell{
		{Text: domain},
		{Text: fmt.Sprintf("%t", hasMx)},
		{Text: fmt.Sprintf("%t", hasSPF)},
		{Text: spfRecord},
		{Text: fmt.Sprintf("%t", hasDMARC)},
		{Text: dmarcRecord},
	}
}
