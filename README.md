### Metasearch API

The Metasearch API is a robust REST API that enables users to gather metasearch data from various sources, including DuckDuckGo, Google, and Wikipedia. This API can be easily deployed to Heroku using the provided "Deploy to Heroku" button.

### Deploy to heroku

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Usage

The following commands are available:

```
metasearch_api [command]
```

#### Available Commands:

```
api         start the api
complete    autocomplete for a query
completion  Generate the autocompletion script for the specified shell
help        Help about any command
query       search for a query
```

#### Flags:

-h, --help help for metasearch_api

### API Endpoints

The Metasearch API is powered by Gin, a high-performance REST API framework. The following endpoints are available:

```
GET /health                       | Checks the health of the API.
POST /search/:query               | Runs a search against the specified query
POST /autocomplete/:keyword       | Get's top 10 autocomplete possibilities for keyword
```

Users can use these endpoints to control the HTTP requests, retrieve search results, and get top autocomplete possibilities for a particular keyword.

Please note that this API is designed for use with Python.

### Deployment

To deploy the Metasearch API, simply click on the "Deploy to Heroku" button above. This will create an instance of the API on Heroku's cloud platform.

### Contributing

Contributions to the Metasearch API are always welcome. To get started, simply fork the repository and make your changes. Once your changes are complete, submit a pull request with a detailed description of the changes you've made.

Please note that all contributors are expected to follow the project's code of conduct, which can be found in the CODE_OF_CONDUCT.md file.

#### License

The Metasearch API is open-source software licensed under the MIT license. For more information, please see the LICENSE file.
