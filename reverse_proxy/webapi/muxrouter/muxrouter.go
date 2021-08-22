package muxrouter

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"webapi/details"
)

const (
	homeRoute      = "/"
	httpsScheme    = "https"
	XForwardedFor  = "X-Forwarded-For"
	XForwardedHost = "X-Forwarded-Host"
	commaDelimiter = ","
	emptyString    = ""
)

var (
	ReverseProxyMux, errReverseProxyMux = createReverseProxyMux(&details.Details.Routes)
	RedirectMux                         = createRedirectMux()
)

type ProxyMux map[string]http.Handler

func (proxyMux ProxyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	redactedURL, errRedactedURL := redactURL(r, nil)
	if redactedURL == nil || errRedactedURL != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	proxyKey := redactedURL.String()
	mux, muxFound := proxyMux[proxyKey]
	if !muxFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	setXForwardedHost(r)
	setXForwardedFor(r)

	mux.ServeHTTP(w, r)
}

func redactURLFromString(fullURL string, err error) (*url.URL, error) {
	redactedURL, errRedactedURL := url.Parse(fullURL)
	if redactedURL != nil {
		redactedURL.RawQuery = emptyString
		redactedURL.Fragment = emptyString
	}

	return redactedURL, errRedactedURL
}

func redactURL(r *http.Request, err error) (*url.URL, error) {
	if err != nil {
		return nil, err
	}

	reqURL := r.Header.Get(XForwardedHost)
	if reqURL == emptyString {
		reqURL = r.URL.String()
	}

	return redactURLFromString(reqURL, nil)
}

func setXForwardedFor(r *http.Request) {
	forwardedFor := r.Header.Get(XForwardedFor)
	if forwardedFor == emptyString {
		r.Header.Set(XForwardedFor, r.RemoteAddr)
	} else {
		r.Header.Set(
			XForwardedFor,
			fmt.Sprint(
				forwardedFor,
				commaDelimiter,
				r.RemoteAddr,
			),
		)
	}
}

func setXForwardedHost(r *http.Request) {
	forwardedHost := r.Header.Get(XForwardedHost)
	if forwardedHost == emptyString {
		r.Header.Set(XForwardedHost, r.Host)
	}
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	dest, errDest := url.Parse(r.URL.String())
	if errDest != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	dest.Scheme = httpsScheme
	destStr := dest.String()

	http.Redirect(
		w,
		r,
		destStr,
		http.StatusMovedPermanently,
	)
}

func createReverseProxyMux(routes *map[string]string) (*ProxyMux, error) {
	proxyMux := make(ProxyMux)

	for dest, target := range *routes {
		destURL, errDestURL := redactURLFromString(dest, nil)
		if errDestURL != nil {
			return nil, errDestURL
		}
		targetURL, errTargetURL := redactURLFromString(target, nil)
		if errTargetURL != nil {
			return nil, errTargetURL
		}

		destRedacted := destURL.String()

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		if details.Details.Server.SkipSSLVerificationOnForwards {
			transportPolicy := http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			proxy.Transport = &transportPolicy
		}

		proxyMux[destRedacted] = httputil.NewSingleHostReverseProxy(targetURL)
	}

	return &proxyMux, nil
}

func createRedirectMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(homeRoute, redirectToHTTPS)

	return mux
}
