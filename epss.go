package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type CVE struct {
	CVE        string `json:"cve"`
	EPSS       string `json:"epss"`
	Percentile string `json:"percentile"`
	Date       string `json:"date"`
}

type APIResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status-code"`
	Version    string `json:"version"`
	Access     string `json:"access"`
	Total      int    `json:"total"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	Data       []CVE  `json:"data"`
}

func main() {
	date := flag.String("d", "", "Sort data by date")
	cve := flag.String("c", "", "Sort data by CVE")
	limit := flag.Int("l", 0, "Number of results to limit")
	mostDangerous := flag.Bool("md", false, "Sort data by most dangerous")
	flag.Parse()

	url := "https://api.first.org/data/v1/epss"

	if *date != "" {
		parsedDate, err := time.Parse("2006-01-02", *date)
		if err != nil {
			fmt.Println("Invalid date format:", err)
			return
		}
		url += fmt.Sprintf("?date=%s", parsedDate.Format("2006-01-02"))
	}

	if *cve != "" {
		if *date == "" {
			url += "?"
		} else {
			url += "&"
		}
		url += fmt.Sprintf("cve=%s", *cve)
	}

	if *limit > 0 {
		if *date == "" && *cve == "" {
			url += "?"
		} else {
			url += "&"
		}
		url += fmt.Sprintf("limit=%d", *limit)
	}

	if *mostDangerous {
		if *date == "" && *cve == "" && *limit == 0 {
			url += "?"
		} else {
			url += "&"
		}
		url += "percentile-gt=0.95"
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if apiResponse.Status != "OK" {
		fmt.Println("API response status is not OK")
		return
	}

	fmt.Println("Total:", apiResponse.Total)
	fmt.Println("Offset:", apiResponse.Offset)
	fmt.Println("Limit:", apiResponse.Limit)

	for _, cve := range apiResponse.Data {
		fmt.Println("CVE ID:", cve.CVE)
		fmt.Println("EPSS:", cve.EPSS)
		fmt.Println("Percentile:", cve.Percentile)
		fmt.Println("Date:", cve.Date)
		fmt.Println("NIST:", "https://nvd.nist.gov/vuln/detail/"+cve.CVE)
		fmt.Println("MITRE:", "https://cve.mitre.org/cgi-bin/cvename.cgi?name="+cve.CVE)
		fmt.Println()
	}
}

