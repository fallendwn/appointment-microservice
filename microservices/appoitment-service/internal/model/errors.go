package model

import "errors"

var ErrNotFound = errors.New("not found")
var ErrDoctorNotFound = errors.New("doctor not found")
var ErrDoctorUnavailable = errors.New("network error: cannot reach doctor-service")
