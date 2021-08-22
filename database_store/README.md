# SuperDB

Deployable SQL Db as webservice.

## Setup

Create `config/config.json` as a configuration file. Use `config/config.json.example` as an example config.

Then run the following command:

`python3 build_and_run_superdb.py`

## Web API

SupderDB has a simple API.

There is only one endpoint: `/`

### Requests

The following is an example of a JSON request body for a `hello world` request to SuperDB.

```JSON
// json
{
    "sql": "SELECT $1;",
    "values": ["hello world!"]
}
```

SQL requests adhere to the following types:

```Golang
// Golang
type Statement struct {
	Sql    string
	Values []interface{}
}
```

```Typescript
// Typescript
interface Statement {
	sql    string
	values unknown[]
}
```

### Responses

The following is an example of the body of a response from SuperDB.

```
// psuedocode
[["row", "values"], ...]
```

```JSON
// json
[["hello", "world!"]]
```

SQL requests adhere to the following types:

```Golang
// Golang
type Statement = [][]interface{}
```

```Typescript
// Typescript
type Statement = [][]unknown,
```

### Errors

The following is an example of a JSON response body for an error from SuperDB.

```JSON
[{
    "kind": "failed to exec",
    "message": "could not complete query"
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