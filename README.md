# Metasearch API

A REST API that collects metasearch data.

- duckduckgo
- google
- wikipedia

### Deploy to heroku

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Commands

```
runs a metasearch rest api engine

Usage:
metasearch_api [command]

Available Commands:
api         start the api
complete    autocomplete for a query
completion  Generate the autocompletion script for the specified shell
help        Help about any command
query       search for a query

Flags:
-h, --help   help for metasearch_api
```

## API

A gin rest API controls the http requests

```
GET /health       | Checks the health of the API.
POST /search/:query     | Runs a search against the specified query
POST /autocomplete/:keyword       | Get's top 10 autocomplete possibilities for keyword
```
