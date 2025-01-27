This repo provides to develop simple REST api calls for a Books store with GET, POST, DELETE, PUT, PATCH API calls.

Run the service inside the GORESTapi/ with command: "go run main.go"

A successful starting of the service shows the logs on screen like- "[GIN-debug] Listening and serving HTTP on localhost:8080"

While the service is running and listening on the localhost:8080, a REST API resquest of type GET/POST/DELETE can be fired from another terminal using curl like this:


Get list of all books:
curl http://localhost:8080/books \
> --header "Content-Type: application/json" \
> --request "GET"

Get a specific by its ID:
curl http://localhost:8080/books/3 \
> --header "Content-Type: application/json" \
> --request "GET"

Add a new book:
curl http://localhost:8080/books \
> --include \
> --header "Content-Type: application/json" \
> --request "POST" \
> --data '{"id": "8", "title": "Dynamic Programming", "author": "Meenakshi Rawat"}'

Delete a book:
curl http://localhost:8080/books/2 \
> --header "Content-Type: application/json" \
> --request "DELETE"