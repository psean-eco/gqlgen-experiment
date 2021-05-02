- Setup using https://gqlgen.com/

- to run server `go run ./server.go`

- connect to http://localhost:8080/ for GraphQL playground

---

### sample query:

```gql
query {
  teams {
    id
    name
    location
    players {
      id
      name
      position
    }
  }
  players {
    id
    name
  }
}
```
