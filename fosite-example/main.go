package main

import (
	"fmt"
	"github.com/go-errors/errors"
	. "github.com/ory-am/fosite"
	"github.com/ory-am/fosite/client"
	"github.com/ory-am/fosite/enigma"
	"github.com/ory-am/fosite/fosite-example/internal"
	"github.com/ory-am/fosite/handler/authorize/explicit"
	"github.com/ory-am/fosite/handler/authorize/implicit"
	goauth "golang.org/x/oauth2"
	"log"
	"net/http"
	"time"
)

var store = &internal.Store{
	Clients: map[string]client.Client{
		"my-client": &client.SecureClient{
			ID:           "my-client",
			Secret:       []byte(`$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO`), // = "foobar"
			RedirectURIs: []string{"http://localhost:3846/callback"},
		},
	},
	AuthorizeCodes: map[string]internal.AuthorizeCodesRelation{},
	AccessTokens:   map[string]internal.AccessRelation{},
	RefreshTokens:  map[string]internal.AccessRelation{},
	Implicit:       map[string]internal.AuthorizeCodesRelation{},
}
var oauth2 OAuth2Provider = fositeFactory()
var clientConf = goauth.Config{
	ClientID:     "my-client",
	ClientSecret: "foobar",
	RedirectURL:  "http://localhost:3846/callback",
	Scopes:       []string{"fosite"},
	Endpoint: goauth.Endpoint{
		TokenURL: "http://localhost:3846/token",
		AuthURL:  "http://localhost:3846/auth",
	},
}

type session struct {
	User string
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/auth", authEndpoint)
	http.HandleFunc("/token", tokenEndpoint)
	http.ListenAndServe(":3846", nil)
}

func tokenEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext()
	var mySessionData session

	accessRequest, err := oauth2.NewAccessRequest(ctx, req, &mySessionData)
	if err != nil {
		oauth2.WriteAccessError(rw, accessRequest, err)
		return
	}

	response, err := oauth2.NewAccessResponse(ctx, req, accessRequest, &mySessionData)
	if err != nil {
		oauth2.WriteAccessError(rw, accessRequest, err)
		return
	}

	oauth2.WriteAccessResponse(rw, accessRequest, response)
}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(fmt.Sprintf(`
<p>You can obtain an access token using various methods</p>
<ul>
	<li>
		<a href="%s">Authorize code grant</a>
	</li>
	<li>
		<a href="%s">Implicit grant</a>
	</li>
	<li>
		<a href="%s">Make an invalid request</a>
	</li>
</ul>
`,
		clientConf.AuthCodeURL("some-random-state-foobar"),
		"http://localhost:3846/auth?client_id=my-client&redirect_uri=http%3A%2F%2Flocalhost%3A3846%2Fcallback&response_type=token&scope=fosite&state=some-random-state-foobar",
		"/auth?client_id=my-client&scope=fosite&response_type=123&redirect_uri=http://localhost:3846/callback",
	)))
}

func callbackHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.URL.Query().Get("error") != "" {
		rw.Write([]byte(fmt.Sprintf(`
<h1>Error!</h1>
Error: %s<br>
Description: %s<br>
<br>
<a href="/">Go back</a>
`,
			req.URL.Query().Get("error"),
			req.URL.Query().Get("error_description"),
		)))
		return
	}
	rw.Write([]byte(fmt.Sprintf(`
<p>Amazing! You just got an authorize code!: %s</p>
<p>Click <a href="/">here to return</a> to the front page</p>
`,
		req.URL.Query().Get("code"),
	)))

	if req.URL.Query().Get("code") == "" {
		rw.Write([]byte(fmt.Sprintf(`
<p>Could not find the authorize code. If you've used the implicit grant, check the browser location bar for the
access token <small><a href="https://en.wikipedia.org/wiki/Fragment_identifier#Basics">(the server side does not have access to url fragments)</a></small></p>
`,
		)))

		return
	}

	token, err := clientConf.Exchange(goauth.NoContext, req.URL.Query().Get("code"))
	if err != nil {
		rw.Write([]byte(fmt.Sprintf(`
<p>
	I tried to exchange the authorize code for an access token but it did not work but got error: %s
</p>
`,
			err.Error(),
		)))
	} else {
		rw.Write([]byte(fmt.Sprintf(`
<p>Cool! You are now a proud token owner.<br>
<ul>
	<li>
		Access token: %s<br>
	</li>
	<li>
		Refresh token: %s<br>
	</li>
</ul>
`,
			token.AccessToken,
			token.RefreshToken,
		)))
	}
}

func authEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := NewContext()

	ar, err := oauth2.NewAuthorizeRequest(ctx, req)
	if err != nil {
		log.Printf("Error occurred in authorize request part: %s\nStack: \n%s", err, err.(*errors.Error).ErrorStack())
		oauth2.WriteAuthorizeError(rw, ar, err)
		return
	}

	if req.Form.Get("username") != "peter" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write([]byte(`
<p>
	Howdy! This is the log in page. For this example, it is enough to supply the username.
</p>
<form method="post">
	<input type="text" name="username" />
	<input type="submit">
</form>
<em>ps: I heard that user "peter" is a valid username so why not try his name ;)</em>
`,
		))
		return
	}

	// Normally, this would be the place where you would check if the user is logged in and gives his consent.
	// For this test, let's assume that the user exists, is logged in, and gives his consent...

	sess := &session{User: "peter"}
	response, err := oauth2.NewAuthorizeResponse(ctx, req, ar, sess)
	if err != nil {
		log.Printf("Error occurred in authorize response part: %s\nStack: \n%s", err, err.(*errors.Error).ErrorStack())
		oauth2.WriteAuthorizeError(rw, ar, err)
		return
	}

	oauth2.WriteAuthorizeResponse(rw, ar, response)
}

func fositeFactory() OAuth2Provider {
	// NewMyStorageImplementation should implement all storage interfaces.

	f := NewFosite(store)
	accessTokenLifespan := time.Hour

	// Let's enable the explicit authorize code grant!
	explicitHandler := &explicit.AuthorizeExplicitEndpointHandler{
		Enigma:              &enigma.HMACSHAEnigma{GlobalSecret: []byte("some-super-cool-secret-that-nobody-knows")},
		Store:               store,
		AuthCodeLifespan:    time.Minute * 10,
		AccessTokenLifespan: accessTokenLifespan,
	}
	f.AuthorizeEndpointHandlers.Add("code", explicitHandler)
	f.TokenEndpointHandlers.Add("code", explicitHandler)

	// Implicit grant type
	implicitHandler := &implicit.AuthorizeImplicitEndpointHandler{
		Enigma:              &enigma.HMACSHAEnigma{GlobalSecret: []byte("some-super-cool-secret-that-nobody-knows")},
		Store:               store,
		AccessTokenLifespan: accessTokenLifespan,
	}
	f.AuthorizeEndpointHandlers.Add("implicit", implicitHandler)

	return f
}
