package main

import (
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"
)

// DbHealthCheck API to check database health
type DbHealthCheck struct {
	DB *gorm.DB
}

// HealthCheckResponse is definition of API response
type HealthCheckResponse struct {
	Healty bool
}

// Check is public API to inspect database health
func (t DbHealthCheck) Check(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HealthCheckResponse, error) {
	r := HealthCheckResponse{false}

	if t.DB.Exec("select 1").Error == nil {
		r.Healty = true
	}

	return http.StatusOK, nil, &r, nil
}
