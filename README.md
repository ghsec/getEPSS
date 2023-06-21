# Exploit Prediction Scoring System (EPSS) 

```
Usage of epss:
  -c string
    	Sort data by CVE
  -d string
    	Sort data by date
  -l int
    	Number of results to limit
  -md
    	Sort data by most dangerous
```

Output of command: `go run epss.go -l 3`


```
Total: 205273
Offset: 0
Limit: 3
CVE ID: CVE-2023-35866
EPSS: 0.000420000
Percentile: 0.004860000
Date: 2023-06-19
NIST: https://nvd.nist.gov/vuln/detail/CVE-2023-35866
MITRE: https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-35866

CVE ID: CVE-2023-35862
EPSS: 0.000450000
Percentile: 0.122720000
Date: 2023-06-19
NIST: https://nvd.nist.gov/vuln/detail/CVE-2023-35862
MITRE: https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-35862

CVE ID: CVE-2023-35857
EPSS: 0.000430000
Percentile: 0.070020000
Date: 2023-06-19
NIST: https://nvd.nist.gov/vuln/detail/CVE-2023-35857
MITRE: https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-35857
```
