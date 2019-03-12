package metawrap

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPClient struct {
	URL    string
	client http.Client
}

func (c HTTPClient) Candidates(text string) (candidates []MappingCandidate, err error) {
	req, err := http.NewRequest("POST", c.URL+"/mm/candidates", bytes.NewBufferString(text))
	if err != nil {
		return
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	if resp.ContentLength == 0 {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&candidates)
	if err != nil {
		return
	}
	err = resp.Body.Close()
	if err != nil {
		return
	}
	return
}
