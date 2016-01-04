# ![Fosite](fosite.png)

Simple and extensible OAuth2 server-side helpers with enterprise security and zero suck.
This library implements [rfc6749](https://tools.ietf.org/html/rfc6749) and enforces countermeasures suggested in [rfc6819](https://tools.ietf.org/html/rfc6819).

[![Build Status](https://travis-ci.org/ory-am/fosite.svg?branch=master)](https://travis-ci.org/ory-am/fosite?branch=master)
[![Coverage Status](https://coveralls.io/repos/ory-am/fosite/badge.svg?branch=master&service=github)](https://coveralls.io/github/ory-am/fosite?branch=master)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Motivation](#motivation)
- [Good to know](#good-to-know)
- [Security](#security)
  - [Encourage security by enforcing it!](#encourage-security-by-enforcing-it)
    - [Secure Tokens](#secure-tokens)
    - [No state, no token](#no-state-no-token)
    - [Opaque tokens](#opaque-tokens)
    - [Advanced Token Validation](#advanced-token-validation)
    - [Encrypt credentials at rest](#encrypt-credentials-at-rest)
    - [Implement peer reviewed IETF Standards](#implement-peer-reviewed-ietf-standards)
  - [Provide extensibility and interoperability](#provide-extensibility-and-interoperability)
  - [Tokens](#tokens)
- [Usage](#usage)
  - [Store](#store)
  - [Authorize Endpoint](#authorize-endpoint)
  - [Token Endpoint](#token-endpoint)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Motivation

Why write another OAuth2 server side library for Go Lang?

Other libraries are perfect for a non-critical set ups, but [fail](https://github.com/RangelReale/osin/issues/107) to comply with enterprise security standards.
This is unfortunately not an issue exclusive to Go's eco system but to many other eco systems as well.

OpenID Connect on top of OAuth2? Not possible with popular OAuth2 libraries. Current libraries do not support capture
the extensibility of OAuth2 and instead bind you to a pattern-enforcing framework with almost no possibilities for extension.

Fosite was written because [Hydra](https://github.com/ory-am/hydra) required a more secure and extensible OAuth2 library
then the one it was using.

## Good to know

Fosite is in early development. We will use gopkg for releasing new versions of the API.
Be aware that "go get github.com/ory-am/fosite" will give you the master branch, which is and always will be *unstable*.
Once releases roll out, you will be able to fetch a specific fosite API version through gopkg.in.

## Security

Fosite has two commandments.

### Encourage security by enforcing it!

#### Secure Tokens

Tokens are generated with a minimum entropy of 256 bit. You can use more, if you want.

#### No state, no token

Without a random-looking state, *GET /oauth2/auth* will fail.

#### Opaque tokens

Token generators should know nothing about the request or context.

#### Advanced Token Validation

Tokens are layouted as `<key>.<signature>`. The following workflow requires an attacker to gain

a. access to the database
b. write permission in the persistent store,
c. get hold of the token encryption secret.

A database leak or (exclusively) the knowledge of the token encrpytion secret are not enough to maliciously obtain or create a valid token. Tokens and credentials can
however still be stolen by man-in-the-middle attacks, by malicious or vulnerable clients and other attack vectors.

**Issuance**

1. The key is hashed using BCrypt (variable) and used as `<signature>`.
2. The client is presented and entrusted with `<key>.<signature>` which can act for example as an access token or an authorize code.
3. The signature is encrypted and stored in the database using AES (variable).

**Validation**

1. The client presents `<key>.<signature>`.
2. It is validated if <key> matches <signature> using BCrypt (variable).
3. The signature is encrypted using AES (variable).
4. The encrypted signature is looked up in the database, failing validation if no such row exists.
5. They key is considered valid and is now subject for other validations, like audience, redirect or state matching.

A token generated by `generator.CryptoGenerator` looks like:

```
GUULhK6Od/7UAlnKvMau8APHSKXSRwm9aoOk56SHBns.JDJhJDEwJDdwVmpCQmJLYzM2VDg1VHJ5aEdVOE81NVdRSkt6bHBHR1QwOC9pbTNFWmpQRXliTWRPeDQy
```

#### Encrypt credentials at rest

Credentials (token signatures, passwords and secrets) are always encrypted at rest.

#### Implement peer reviewed IETF Standards

Fosite implements [rfc6749](https://tools.ietf.org/html/rfc6749) and enforces countermeasures suggested in [rfc6819](https://tools.ietf.org/html/rfc6819).

### Provide extensibility and interoperability

... because OAuth2 is an extensible and flexible **framework**. Fosite let's you register new response types, new grant
types and new response key value pares. This is useful, if you want to provide OpenID Connect on top of your
OAuth2 stack. Or custom assertions, what ever you like and as long as it is secure. ;)

## Usage

This section is WIP and we welcome discussions via PRs or in the issues.

### Store

To use fosite, you need to implement `fosite.Storage`. Example implementations (e.g. postgres) of `fosite.Storage`
will be added in the close future.

### Authorize Endpoint

```go
package main

import "github.com/ory-am/fosite"
import "github.com/ory-am/fosite/session"
import "github.com/ory-am/fosite/storage"

// Let's assume that we're in a http handler
func handleAuth(rw http.ResponseWriter, req *http.Request) {
    store := fosite.NewPostgreSQLStore()
    oauth2 := fosite.NewDefaultOAuth2(store)

    authorizeRequest, err := oauth2.NewAuthorizeRequest(r, store)
    if err != nil {
        // ** This part of the API is not finalized yet **
        // oauth2.RedirectError(rw, error)
        // oauth2.WriteError(rw, error)
        // oauth2.handleError...
        // ****
        return
    }

    // you have now access to authorizeRequest, Code ResponseTypes, Scopes ...
    // and can show the user agent a login or consent page.

    // it would also be possible to redirect the user to an identity provider (google, microsoft live, ...) here
    // and do fancy stuff like OpenID Connect amongst others

    // Once you have confirmed the users identity and consent that he indeed wants to give app XYZ authorization,
    // you will use the user's id to create an authorize session
    user := "12345"
    session := fosite.NewAuthorizeSession(authorizeRequest, user)

    // You can store additional metadata, for example:
    session.SetExtra(map[string]interface{}{
         "userEmail": "foo@bar",
         "lastSeen": new Date(),
         "usingIdentityProvider": "google",
    })

    // or
    session.SetExtra(&struct{
        Name string
    } {
        Name: "foo"
    })

    // Now is the time to handle the response types

    // ** This part of the API is not finalized yet **
    // response = &AuthorizeResponse{}
    // err = oauth2.HandleResponseTypes(authorizeRequest, response, session)
    // err = alsoHandleMyCustomResponseType(authorizeRequest, response, "fancyArguments", 1234)
    //
    // or
    //
    // this approach would make it possible to check if all response types could be served or not
    // additionally, a callback for FinishAccessRequest could be provided
    //
    // response = &AuthorizeResponse{}
    // oauth2.RegisterResponseTypeHandler("custom_type", alsoHandleMyCustomResponseType)
    // err = oauth2.HandleResponseTypes(authorizeRequest, response, session)
    // ****

    // Almost done! The next step is going to persist the session in the database for later use.
    // It is additionally going to output a result based on response_type.

    // ** API not finalized yet **
    // err := oauth2.FinishAuthorizeRequest(rw, response, session)
    // ****
}
```

### Token Endpoint