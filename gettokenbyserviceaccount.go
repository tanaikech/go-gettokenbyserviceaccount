// Package gettokenbyserviceaccount (gettokenbyserviceaccount.go) :
// This is a Golang library to retrieve access token from Service Account of Google without using Google's OAuth2 package.
package gettokenbyserviceaccount

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	url = "https://www.googleapis.com/oauth2/v3/token"
)

// AccessToken :
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Start       int64  `json:"start_time"`
	End         int64  `json:"end_time"`
}

// fetch : Fetch access token.
func fetch(jwt string) ([]byte, error) {
	payload := struct {
		Assertion string `json:"assertion"`
		GrantType string `json:"grant_type"`
	}{
		jwt,
		"urn:ietf:params:oauth:grant-type:jwt-bearer",
	}
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(p),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	res, err := client.Do(req)
	body, errc := ioutil.ReadAll(res.Body)
	if err != nil || res.StatusCode-300 >= 0 {
		return nil, fmt.Errorf("status code is %d. error message is %s", res.StatusCode, body)
	}
	if errc != nil {
		return nil, errc
	}
	defer res.Body.Close()
	return body, nil
}

// getPrivateKey : Retrieve private key.
func getPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	pkb, _ := pem.Decode([]byte(privateKey))
	if pkb == nil || pkb.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("this is invalid private key value")
	}
	pk, err := x509.ParsePKCS8PrivateKey(pkb.Bytes)
	if err != nil {
		return nil, err
	}
	key, ok := pk.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("this is not RSA private key")
	}
	return key, nil
}

// createSignature : Create signature.
func (a *AccessToken) createSignature(clientEmail, impersonateEmail, scopes string) (string, error) {
	h := struct {
		Alg string `json:"alg"`
		Typ string `json:"typ"`
	}{
		"RS256",
		"JWT",
	}
	c := struct {
		Iss   string `json:"iss"`
		Sub   string `json:"sub"`
		Scope string `json:"scope"`
		Aud   string `json:"aud"`
		Exp   string `json:"exp"`
		Iat   string `json:"iat"`
	}{
		Iss:   clientEmail,
		Scope: scopes,
		Aud:   url,
		Exp:   strconv.FormatInt(a.End, 10),
		Iat:   strconv.FormatInt(a.Start, 10),
	}

	// add impersonateEmail to the struct
	if impersonateEmail != "" {
		c.Sub = impersonateEmail
	}

	he, err := json.Marshal(h)
	if err != nil {
		return "", err
	}
	cl, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(he) + "." + base64.StdEncoding.EncodeToString(cl)
	return signature, nil
}

// Do : Retrieve access token from Google Service Account
func Do(privateKey, clientEmail, impersonateEmail, scopes string) (*AccessToken, error) {
	if privateKey == "" || clientEmail == "" || scopes == "" {
		return nil, fmt.Errorf("invalid argument. please confirm the private key, client email and scopes")
	}
	nt := time.Now().Unix()
	a := &AccessToken{
		Start: nt,
		End:   nt + 3600,
	}
	signature, err := a.createSignature(clientEmail, impersonateEmail, scopes)
	if err != nil {
		return nil, err
	}
	key, err := getPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	ha := sha256.Sum256([]byte(signature))
	signatureSuffix, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, ha[:])
	if err != nil {
		return nil, err
	}
	jwt := signature + "." + base64.StdEncoding.EncodeToString(signatureSuffix)
	res, err := fetch(jwt)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(res, &a)
	return a, nil
}
