# services-catalog-api
An api services to manage the services catalog

### Supported Endpoints
* *GET /api/v1/services*
    * List all services.
    * Please note that services will not be searchable if there is no enabled version exists
* *GET /api/v1/services/:id*
    * Get a single service
* *GET /api/v1/services/:id/versions*
    * Get all available versions belonging to a service
* *GET /api/v1/services?search=reporting*
    * Filter the services 
    * Search applied to both service name and description
* *GET /api/v1/services?sort=name*
    * Sort the services
* *GET /api/v1/services?page=1&per_page=2*
    * Pagination Sample response
        ```
        {
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
        }
        ```
    * page is the 1-based index of current page
    * per_page is the number of item returned in each response
    * page_count is the total available pages
    * total_count is the total available items
* *POST /api/v1/services* 
    * Create a new service
    * Sample curl command
        ```shell
        curl http://localhost:8080/api/v1/services \
        --header "Content-Type: application/json" \
        --request "POST" \
        --data '{"name": "Security","description": "The security service","user_id": 1}'
        ```

### Getting Started
Docker is needed to run the application in the local
``` shell
# start the api services with seeded data
make run

# run integration test
make test
```

### Future enhancement
* Add support for more endpoints
* Better error handling
* Better test coverage
* Query String and Json Body validation
* Better Pagination strategy
    * Return pre-generated query parameter in the response, which could be easily used in the following query
    * Cache the results in centralized cache to avoid querying DB every time

