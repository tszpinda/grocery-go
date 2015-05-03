package main

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"net/url"
)

type DbHealthCheck struct {
	DB *gorm.DB
}

type HealthCheckResponse struct {
	Healty bool
}

func (t DbHealthCheck) Check(u *url.URL, h http.Header, _ interface{}) (int, http.Header, *HealthCheckResponse, error) {
	r := HealthCheckResponse{false}

	if t.DB.Exec("select 1").Error == nil {
		r.Healty = true
	}

	return http.StatusOK, nil, &r, nil
}
