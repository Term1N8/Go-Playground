package checks

import (
	"io"
	"log"
	"net/http"
	"strings"
	"github.com/miekg/dns"
)

type Check struct {
	Name          string
	Description   string
	Endpoint      string
	RecordType    string
	Server        string
	SuccessStatus int
	SuccessString string
	
}

func (c *Check) Run() bool {
	if strings.HasPrefix(strings.ToLower(c.Endpoint), "http") {
		// http
		res, err := http.Get(c.Endpoint)
		if err != nil {
			return false
		}

		if c.SuccessString != "" {
			bodyBytes, err := io.ReadAll(res.Body)
			body := string(bodyBytes)
			//log.Println(body)
			if err != nil {
				log.Println(err)
			} else if strings.Contains(body, c.SuccessString) {
				return true
			} else{
				return false;
			}
		}

		if c.SuccessStatus > 0 && res.StatusCode == c.SuccessStatus {
			return true
		}
	}
	if strings.Contains(strings.ToLower(c.Name), "dns") {
		m := new(dns.Msg)
		if c.RecordType == "A"{
		m.SetQuestion(c.Endpoint, dns.TypeA)}
		if c.RecordType == "TXT"{
			m.SetQuestion(c.Endpoint, dns.TypeTXT)}
		b := new(dns.Client)
		in, _, err := b.Exchange(m, c.Server)
		_ = in
		if err != nil {
			return false
		}

		return true
	}
	return false
}

var All = []Check{
	HTTPIPCurl,
	HTTPS2j,
	HTTPSAzureFP,
	DNSOpenDNS,
	DNSQuad9,
	NateAWS,

}
