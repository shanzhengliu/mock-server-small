format of the file you mount is 
```json
[
  {
    "path": "/api/hello",
    "method": "GET",
    "responseBody": "Hello, World!",
    "contextType": "text/plain",
    "status": 200
  },
  {
    "path": "/api/data",
    "method": "POST",
    "responseBody": "{\"message\": \"Data received\"}",
    "contextType": "application/json",
    "status": 401
  }
]
```
run the command with
```shell
docker run -d -v {your route path}:/app/routes.json -p 8080:8080 samliu960522/mock-server-small:latest
```