# CacheStore

Cache as webservice.

This is only to be made available behind a firewall and another SSL backed webapi.

## Setup

Create `config/config.json` as a configuration file. Use `config/config.json.example` as an example config.

Then run the following command:

`python3 build_and_run_local_cache.py`

## Web API

CacheStore has an API with a single endpoint: `/`

### Requests

The following is an example of a JSON request body for a `hello world` request to CacheStore.

```JSON
// json
["INCR", "MYCOUNTER"]
```

CacheStore requests adhere to the following types:

```Golang
// Golang
type Statement = []interface{}
```

```Typescript
// Typescript
interface Statement = unknown[]
```

### Responses

The following is an example of the body of a response from CacheStore. Due to the nature of remote caches, anything can be returned.

```Golang
// Golang
type Statement = interface{}
```

```Typescript
// Typescript
type Statement = unknown
```

### Errors

The following is an example of a JSON response body for an error from CacheStore.

```JSON
[{
    "kind": "failed to exec",
    "message": "could not complete exec"
}]
```

SQL requests adhere to the following types:

```Golang
// Golang
type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}
```

```Typescript
// Typescript
interface ErrorEntity {
	kind    string
	message string
}
```

## License

Apache-2.0