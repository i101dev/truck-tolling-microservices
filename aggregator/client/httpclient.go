package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/microservices/types"
)

type HTTPClient struct {
	Endpoint string
}

func NewHTTPClient(endpoint string) *HTTPClient {
	return &HTTPClient{
		Endpoint: endpoint,
	}
}

// Give clients the name of the respective endpoint
func (c *HTTPClient) AggregateInvoice(distance types.Distance) error {

	b, err := json.Marshal(distance)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(b))

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the service responded with non-200 status code: %d", resp.StatusCode)
	}

	return nil
}