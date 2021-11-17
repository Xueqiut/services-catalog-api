package main

import (
	"net/http"
	"testing"
	"log"

	integration "github.com/Xueqiut/services-catalog-api/pkg/test"
	helper "github.com/Xueqiut/services-catalog-api/internal/test"
)

var (
	serviceData [][]interface{} = [][]interface{}{
		{"Locate Us", "The location service", 1},
		{"Contact Us", "The contact service", 1},
		{"Notifications", "The notifications service", 1},
		{"Reporting", "The reporting service", 1},
	}

	VersionData [][]interface{} = [][]interface{}{
		{"v1", 1, true},
		{"v1", 2, false},
		{"v2", 2, true},
		{"v1", 3, false},
		{"v2", 3, false},
		{"v3", 3, true},
		{"v1", 4, false},
		{"v2", 4, false},
		{"v3", 4, false},
		{"v4", 4, true},
	}
)

func TestServiceRoute(t *testing.T) {
	cfg := config.Init()

	testHelper, err := helper.NewTest(cfg.ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	db := testHelper.GetDb()
	router := setupRouter(db, log.Default())

	testHelper.ProvisionService(serviceData)
	testHelper.ProvisionVersion(VersionData)
	defer testHelper.Truncate()

	tests := []integration.TestCase{
		{"list all services", "GET", "/api/v1/services", "", nil, http.StatusOK, expectedList},
		{"get one specific service", "GET", "/api/v1/services/2", "", nil, http.StatusOK, expectedId},
		{"get all versions of one service", "GET", "/api/v1/services/4/versions", "", nil, http.StatusOK, expectedServiceVersions},
		{"filter services", "GET", "/api/v1/services?search=reporting", "", nil, http.StatusOK, expectedFilter},
		{"sort services", "GET", "/api/v1/services?sort=name", "", nil, http.StatusOK, expectedSort},
		{"pagination", "GET", "/api/v1/services?page=1&per_page=3", "", nil, http.StatusOK, expectedPagination},
		{"pagination", "GET", "/api/v1/services?page=1&per_page=3", "", nil, http.StatusOK, expectedPagination},
	}

	for _, tc := range tests {
		integration.Run(t, router, tc)
	}
}