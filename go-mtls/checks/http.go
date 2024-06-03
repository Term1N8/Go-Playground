package checks

var HTTPIPCurl = Check{
	Name:          "IPCurl.net",
	Description:   "Can access IPCurl.net",
	Endpoint:      "http://ipcurl.net/n",
	SuccessStatus: 200,
	SuccessString: "",
}

var HTTPS2j = Check{
	Name:          "s2j.cc",
	Description:   "Can access s2j.cc",
	Endpoint:      "https://s2j.cc/ip",
	SuccessStatus: 200,
	SuccessString: "",
}

var HTTPSAzureFP = Check{
	Name:          "Azure Function Proxies",
	Description:   "Can access Azure function proxies",
	Endpoint:      "https://sync-ad.azurewebsites.net/dont/care/about/me/dom/index.html",
	SuccessStatus: 200,
	SuccessString: "hello",
}
var NateAWS = Check{
	Name:          "AWS HTTP",
	Description:   "Can access N8s sketchy HTTP server?",
	Endpoint:      "http://18.221.246.49/index.html",
	SuccessStatus: 200,
	SuccessString: "ls -al",
}

