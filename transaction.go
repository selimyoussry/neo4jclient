package neo4jclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Request calls the database and returns the response
func (n *Neo) Request(payload *Payload) (*Response, error) {

	serializedPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(
		REQUEST_METHOD,
		n.endpoint,
		bytes.NewBuffer(serializedPayload),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_AUTHORIZATION, n.authHeader)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	deserializedBody := &Response{}
	err = json.Unmarshal(body, deserializedBody)
	if err != nil {
		return nil, err
	}

	return deserializedBody, nil

}
