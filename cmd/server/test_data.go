package main

var (
	expectedList string = `{
		"page": 1,
		"per_page": 2,
		"page_count": 2,
		"total_count": 4,
		"items": [
			{
				"id": "1",
				"name": "Locate Us",
				"description": "The location service",
				"user_id": 1,
				"version": "v1"
			},
			{
				"id": "2",
				"name": "Contact Us",
				"description": "The contact service",
				"user_id": 1,
				"version": "v2"
			}
		]
	}`

	expectedId string = `{
		"id": "2",
		"name": "Contact Us",
		"description": "The contact service",
		"user_id": 1,
		"version": "v2"
	}`

	expectedServiceVersions string = `[
		{
			"id": "7",
			"name": "v1",
			"service_id": 4,
			"enabled": false
		},
		{
			"id": "8",
			"name": "v2",
			"service_id": 4,
			"enabled": false
		},
		{
			"id": "9",
			"name": "v3",
			"service_id": 4,
			"enabled": false
		},
		{
			"id": "10",
			"name": "v4",
			"service_id": 4,
			"enabled": true
		}
	]`

	expectedFilter string = `{
		"page": 1,
		"per_page": 2,
		"page_count": 2,
		"total_count": 4,
		"items": [
			{
				"id": "4",
				"name": "Reporting",
				"description": "The reporting service",
				"user_id": 1,
				"version": "v1"
			},
			{
				"id": "4",
				"name": "Reporting",
				"description": "The reporting service",
				"user_id": 1,
				"version": "v2"
			}
		]
	}`

	expectedSort string = `{
		"page": 1,
		"per_page": 2,
		"page_count": 2,
		"total_count": 4,
		"items": [
			{
				"id": "2",
				"name": "Contact Us",
				"description": "The contact service",
				"user_id": 1,
				"version": "v2"
			},
			{
				"id": "1",
				"name": "Locate Us",
				"description": "The location service",
				"user_id": 1,
				"version": "v1"
			}
		]
	}`

	expectedPagination string = `{
		"page": 1,
		"per_page": 3,
		"page_count": 2,
		"total_count": 4,
		"items": [
			{
				"id": "1",
				"name": "Locate Us",
				"description": "The location service",
				"user_id": 1,
				"version": "v1"
			},
			{
				"id": "2",
				"name": "Contact Us",
				"description": "The contact service",
				"user_id": 1,
				"version": "v2"
			},
			{
				"id": "3",
				"name": "Notifications",
				"description": "The notifications service",
				"user_id": 1,
				"version": "v3"
			}
		]
	}`
)