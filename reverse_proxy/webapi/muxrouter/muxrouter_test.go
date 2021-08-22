package muxrouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	// "net/url"
	"os"
	"testing"

	"webapi/details"
)

const (
	testHTTPStr        = "http://awesome.sauce.com/yoyoyo?=hello=world#framented-fragement"
	testURLStr         = "https://awesome.sauce.com/yoyoyo?=hello=world#framented-fragement"
	expectedTestURLStr = "https://awesome.sauce.com/yoyoyo"
	superawesome       = "https://superawesome.com"
	localAddress       = "https://127.0.0.16:6000"
	expectedAddress    = "https://127.0.0.1:5000"
)

var (
	exampleDetailsPath = os.Getenv("CONFIG_FILEPATH_TEST")
)

func TestRedactURL(t *testing.T) {
	request, errRequest := http.NewRequest("GET", testURLStr, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	redactedURL, errRedactedURL := redactURL(request, nil)
	if errRedactedURL != nil {
		t.Fail()
		t.Logf(errRedactedURL.Error())
	}

	redactedURLStr := redactedURL.String()
	if redactedURLStr != expectedTestURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("example detail cert path: ", expectedTestURLStr, "\nfound:", redactedURLStr))
	}
}

func TestRedactURLWithXForwardedHost(t *testing.T) {
	request, errRequest := http.NewRequest("GET", localAddress, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	request.Header.Set(XForwardedHost, testURLStr)

	redactedURL, errRedactedURL := redactURL(request, nil)
	if errRedactedURL != nil {
		t.Fail()
		t.Logf(errRedactedURL.Error())
	}

	redactedURLStr := redactedURL.String()
	if redactedURLStr != expectedTestURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("example detail cert path: ", expectedTestURLStr, "\nfound:", redactedURLStr))
	}
}

func TestRedactURLFromString(t *testing.T) {
	redactedURL, errRedactedURL := redactURLFromString(testURLStr, nil)
	if errRedactedURL != nil {
		t.Fail()
		t.Logf(errRedactedURL.Error())
		return
	}

	redactedURLStr := redactedURL.String()
	if redactedURLStr != expectedTestURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("example detail cert path: ", expectedTestURLStr, "\nfound:", redactedURLStr))
	}
}

func TestRedirectToHTTPS(t *testing.T) {
	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	testRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(redirectToHTTPS)
	handler.ServeHTTP(testRecorder, request)

	locations := testRecorder.HeaderMap["Location"]
	if locations[0] != testURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
	}
}

func TestCreateRedirectMux(t *testing.T) {
	request, errRequest := http.NewRequest("GET", testHTTPStr, nil)
	if errRequest != nil {
		t.Fail()
		t.Logf(errRequest.Error())
	}

	testRecorder := httptest.NewRecorder()

	mux := createRedirectMux()
	mux.ServeHTTP(testRecorder, request)

	locations := testRecorder.HeaderMap["Location"]
	if locations[0] != testURLStr {
		t.Fail()
		t.Logf(fmt.Sprint("redirected http request should be: ", testURLStr, "\nfound:", locations[0]))
	}
}

func TestCreateReverseProxyMux(t *testing.T) {
	exampleDetails, errExampleDetails := details.ReadDetailsFromFile(exampleDetailsPath)
	if errExampleDetails != nil {
		t.Fail()
		t.Logf(errExampleDetails.Error())
	}

	proxyMux, errProxyMux := createReverseProxyMux(&exampleDetails.Routes)
	if errProxyMux != nil {
		t.Fail()
		t.Logf(errProxyMux.Error())
	}

	if proxyMux == nil {
		t.Fail()
		t.Logf("proxyMux was not created")
	}
}
