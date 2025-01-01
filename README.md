### React Go Government Data Viewer
===============================

This project demonstrates a React frontend and Go backend application that fetches and displays government data using Docker containers.

* * * * *

## Prerequisites
-------------

-   Docker installed on your machine.

* * * * *

## Setup
-----

Clone the repository:

```
git clone https://github.com/ethirocks/react-go-gov-search-data
```

* * * * *

## Build and Run the Application
-----------------------------

## Using Docker Compose:

```
docker-compose up --build
```

This will:

-   Build the Go backend and React frontend Docker images.

-   Start the backend and frontend containers.

* * * * *

## Access the Application
----------------------

1.  Open your browser and navigate to: <http://localhost:3000>

2.  The frontend should display a data viewer with filtering capabilities.

* * * * *

## Stopping the Application
------------------------

To stop the running containers:

```
docker-compose down
```

* * * * *

## Rebuilding the Containers
-------------------------

If you make changes to the code and need to rebuild the containers, use:

```
docker-compose up --build
```

* * * * *

## Troubleshooting
---------------

## Common Issues

## 1\. Failed to Fetch Data

-   Ensure the backend container is running.

-   Verify the backend API is accessible:

    ```
    curl http://localhost:8080/api/data
    ```

-   Check Docker logs for errors:

    ```
    docker logs react-go-gov-search-data
    ```

## 2\. Frontend Not Loading

-   Verify the frontend container is running.

-   Check the logs:

    ```
    docker logs react-frontend
    ```

## 3\. CORS Errors

-   Ensure the backend is correctly configured to allow requests from the frontend.

-   Verify the `AllowedOrigins` setting in the backend CORS configuration matches the frontend URL.

* * * * *

Additional Commands
-------------------

## View Running Containers

To see all running Docker containers:

```
docker ps
```

## Stop and Remove Containers

To stop and remove all containers:

```
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
```

## Check Logs

To view logs for a specific container:

```
docker logs <container_name>
```

Replace `<container_name>` with `react-frontend` or `react-go-gov-search-data`.

* * * * *

Environment Variables
---------------------

The application uses the following environment variables:

## Backend

-   `PORT`: Port on which the backend listens (default: `8080`).

## Frontend

-   `PORT`: Port on which the frontend runs (default: `3000`).

-   `REACT_APP_BACKEND_URL`: URL for the backend API (default: `http://localhost:8080`).

* * * * *

API Endpoints
-------------

## Backend

## `GET /api/data`

Fetches government data from a public API and serves it to the frontend.

**Example Response:**

```
{
  "success": true,
  "data": [
    {
      "state": "California",
      "population": "39346023",
      "state_fips": "06"
    },
    {
      "state": "Texas",
      "population": "28635442",
      "state_fips": "48"
    }
  ],
  "message": "Government data retrieved successfully"
}
```

* * * * *

Project Structure
-----------------

```
.
├── backend/            # Go backend code
│   ├── govData/v1/     # Backend API for data fetching
│   └── main.go         # Main backend entry point
├── frontend/           # React frontend code
│   ├── src/            # Source code for the React app
│   └── Dockerfile      # Dockerfile for frontend
├── docker-compose.yml  # Docker Compose configuration
└── README.md           # Instructions for running the application
```

* * * * *

Contact
-------

For questions or support, please contact: [your-email@example.com].