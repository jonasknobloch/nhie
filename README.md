# Never Have I Ever

[![Go Report Card](https://goreportcard.com/badge/github.com/nhie-io/api)](https://goreportcard.com/report/github.com/nhie-io/api)

## API

### Endpoints

#### v1

```http request
GET https://api.nhie.io/v1/statements/random
```

**This endpoint is deprecated and might be removed at any time.**

Since a surprising amount of other projects depend on the original API endpoint I kept it around for now.
Note that the `history_id` query parameter is not supported anymore. See [duplicate statements](#duplicate-statements)
for a similar functionality.

#### v2

```http request
GET https://api.nhie.io/v2/statements/next
```

```json
{
  "ID":"e9ed69c3-b1ee-4d20-8e8c-9b766577a49a",
  "statement":"Never have I ever played a drinking game.",
  "category":"harmless"
}
```

### Query Parameters

| Key          | Value                           | Endpoint    |
|--------------|---------------------------------|-------------|
| category     | harmless, delicate or offensive | v1 & v2     |
| language     | IETF BCP 47 language tag        | v1 & v2     |
| statement_id | UUID of previous statement      | **v2 only** |

### Multiple Categories

Multiple categories can be queried by adding multiple `category` parameters.
A random category is used if no `category` parameter is set.

```http request
GET https://api.nhie.io/v2/statements/next?category=delicate&category=offensive
```

### Supported Languages

The currently supported languages are listed below. More languages might be added in the future.

| Language          | BCP 47 |
|-------------------|--------|
| English (default) | en     |
| Deutsch           | de     |

```http request
GET https://api.nhie.io/v2/statements/next?language=de
```

### Duplicate Statements

The `statement_id` parameter can be used to avoid duplicate statements during a game session.
All available statements are internally ordered in a random fashion. With `statement_id` set,
the returned statement is guaranteed be different from previous statements.

```http request
GET https://api.nhie.io/v2/statements/next?statement_id=e9ed69c3-b1ee-4d20-8e8c-9b766577a49a
```

## Contributing

Pull requests are welcome. Please open an issue first to discuss what you would like to change.
