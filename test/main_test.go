package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

/*
	Setup test sever

	TODO: .env debug: {true} or {false}

	1. function debug (bool)
	@function debug: {true} or {false}

	2. os.Setenv("debug", ...)
	@debug: {true} or {1}
	os.Setenv("debug", "true") or,
	os.Setenv("debug", "1")

	@debug: {false} or {0}
	os.Setenv("debug", "false") or,
	os.Setenv("debug", "0")

*/
func setupTestServer(t *testing.T, debug ...bool) (noAuth *httpexpect.Expect) {
	os.Setenv("debug", "0")

	handler := setupTestHandler()

	server := httptest.NewServer(handler)
	defer server.Close()

	newConfig := httpexpect.Config{
		BaseURL: server.URL,
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCompactPrinter(t),
		},
	}

	if (len(debug) == 1 && debug[0] == true) || (os.Getenv("debug") == "1" || os.Getenv("debug") == "true") {
		newConfig.Printers = []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		}
	} else if len(debug) > 1 {
		panic("func setupTestServer: (debug [1]: true or false) or no debug")
	}

	noAuth = httpexpect.WithConfig(newConfig)
	return
}

// Setup test sever authentication
// request with cookie session
func setupTestServerAuth(e *httpexpect.Expect) (auth *httpexpect.Expect) {
	auth = e.Builder(func(req *httpexpect.Request) {
		// TODO: if (isAdmin or isUser: bool) {...}
		req.WithCookie("session", session)
	})
	return
}

/*
	HTTP(s)
	----
	[+] Request Headers
	Cookie: session=...

	or,

	[+] Request Cookies
	session: ...

	-------
	Cookie:
	MaxAge=0 means no Max-Age attribute specified and the cookie will be
	deleted after the browser session ends.

	sessions.Options{.., MaxAge: 0,..}

	-------
	func. SetSession:

	session_gorilla, err = session.Get("session", ...)
	...
	session_gorilla.Values["username"] = user.Username
	session_gorilla.Values["is_auth_type"] = 2 // admin: 1 and user: 2
	---
*/
const session = "MTY0MDA4MzU1MnxEdi1CQkFFQ180SUFBUkFCRUFBQVNfLUNBQUlHYzNSeWFXNW" +
	"5EQW9BQ0hWelpYSnVZVzFsQm5OMGNtbHVad3dNQUFwdlkydHBZbUZuZFhOd0JuTjBjbWx1Wnd3" +
	"T0FBeHBjMTloZFhSb1gzUjVjR1VEYVc1MEJBSUFCQT09fIlgmThOxd1Xxc_uh6jeRFkCwwHLW7" +
	"rA_0tH8qPT9t41"

func TestServer(t *testing.T) {
	//
}