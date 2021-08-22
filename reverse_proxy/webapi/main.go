package main

import (
	"fmt"
	"log"
	"net/http"

	"webapi/details"
	"webapi/muxrouter"
)

var (
	httpPort     = fmt.Sprint(":", details.Details.Server.HTTPPort)
	httpsPort    = fmt.Sprint(":", details.Details.Server.HTTPSPort)
	certFilepath = details.Details.CertPaths.Cert
	keyFilepath  = details.Details.CertPaths.PrivateKey
)

func main() {
	if details.Details.Server.RedirectFromHTTPToHTTPS {
		go http.ListenAndServe(
			httpPort,
			muxrouter.RedirectMux,
		)
	}

	errServer := http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		muxrouter.ReverseProxyMux,
	)

	if errServer != nil {
		log.Println(errServer.Error())
	}
}
