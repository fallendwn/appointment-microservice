package client

import (
	"context"
	"fmt"
	"net/http"
)

func CheckDoctorExists(ctx context.Context, doctorID string) (bool, error) {
	url := fmt.Sprintf("http://localhost:8081/api/v1/doctors/%s", doctorID)

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK, nil
}
