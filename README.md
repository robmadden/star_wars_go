# Star Wars Go Example 
An example Golang server that uses GraphQL and Relay that serves up Star Wars data.

Taken from: [Playground](https://github.com/graphql-go/playground/blob/master/main.go)

<hr />

###  Features

- [graphql-go](https://github.com/graphql-go/graphql): Golang GraphQL library
- [graphql-hander](https://github.com/graphql-go/hander): Golang HTTP.Handler for graphl-go
- [Object88's Relay](https://github.com/object88/relay/)

### Useful Links

- [How to Write Go Code](https://golang.org/doc/code.html#Workspaces)
- [Deploying Go servers with Docker](https://blog.golang.org/docker)

### How to run Hello Go locally:

1. Make sure you have go installed.
    - [macOS Install Instructions](https://golang.org/doc/install) 

2. Make sure you have `GOPATH` set to the directory where this project is installed:

```bash
$ export GOPATH=`pwd`
```

3. Run the app: 

```bash
$ go run main.go
```

### cURL Examples

#### Rebel Alliance Queries

Fetch the Rebel faction name:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { rebels { name } }", 
    "variables": null
}'
```

Fetch the first Rebel ship:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { rebels { name ships(first: 1) { edges { node { name } } } } }",
    "variables": null
}'
```

Fetch all Rebel ships:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { rebels { name ships { edges { node { name } } } } }",
    "variables": null
}'
```

#### Empire Queries

Fetch the Empire faction name:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { empire { name } }", 
    "variables": null
}'
```

Fetch the first Empire ship:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { empire { name ships(first: 1) { edges { node { name } } } } }",
    "variables": null
}'
```

Fetch all Empire ships:

```bash
$ curl -X POST http://localhost:3000/graphql -H 'Content-Type: application/json' -d '
{
    "query": "query { empire { name ships { edges { node { name } } } } }",
    "variables": null
}'
```

### Docker How To

How to do stuff with Docker.

#### Docker Build

```bash
$ docker build -t star_wars_go .
```

#### Docker Run

```bash
$ docker run --publish 3000:3000 --name star_wars_go --rm star_wars_go
```

#### Docker Stop

```bash
$ docker stop star_wars_go
```
