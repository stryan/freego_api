# freego_api
freego implementation as a simple REST API.

Currently stores no state and has no user authentication, so using it should just be
```
  go build
  ./freego_api
```
# Implemented routes
```
  POST /game
  GET /game/{id}
  GET /game/{id}/status
  POST /game/{id}/move
  GET /game/{id}/move/{movenum}
```

## Authorization headers
  Requests should contain header "Player-id"="red|blue"
