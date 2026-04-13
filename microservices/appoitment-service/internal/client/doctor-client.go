package client

import (
	"context"
	"fmt"
	"net/http"
)

type DoctorClient struct {
	baseURL string
}

func NewDoctorClient(baseURL string) *DoctorClient {
	return &DoctorClient{baseURL: baseURL}
}

func (c *DoctorClient) CheckDoctorExists(ctx context.Context, doctorID string) (bool, error) {
	url := fmt.Sprintf("%s/api/v1/doctors/%s", c.baseURL, doctorID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
