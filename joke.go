package joke

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

const (
	jokeURL     = "https://icanhazdadjoke.com/"
	headerJSON  = "application/json"
	headerHTML  = "text/html"
	headerPlain = "text/plain"
)

// GetJSON returns a DadJoke.
// Possible Accept Header Values: application/json, text/html, and text/plain. Defaults to json.
func GetJSON() DadJoke {
	var dj = DadJoke{}

	res, err := makeRequest("application/json")
	if err != nil {
		return dj
	}

	json.NewDecoder(res.Body).Decode(&dj)
	return dj
}

// GetHTML returns an HTML formatted string.
func GetHTML() string {
	res, err := makeRequest("text/html")

	if err != nil {
		return ""
	}
	defer res.Body.Close()

	// log.Println(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	return string(body)
}

func makeRequest(applicationHeader string) (*http.Response, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	req, err := http.NewRequest("GET", jokeURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", applicationHeader)
	req.Header.Set("User-Agent", "https://github.com/rbo13/dad-joke")

	return client.Do(req)
}
