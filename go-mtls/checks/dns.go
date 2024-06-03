package checks



var DNSOpenDNS = Check{
	Name:          "DNS OpenDNS s2j.cc A record",
	Description:   "Can resolves the s2j.cc via OpenDNS",
	Endpoint:      "s2j.cc.",
	Server:			"208.67.222.222:53",
	RecordType:		"A",
}
var DNSQuad9 = Check{
	Name:          "DNS Quad9 uvcyber.com TXT record",
	Description:   "Can resolves the s2j.cc via Quad9",
	Endpoint:      "uvcyber.com.",
	RecordType:	   "TXT",
	Server:			"9.9.9.9:53",
}