# Simple and extended search

## Description

Execute a simple search. You may recognize that the results can contain many companies, that are similar to the searched name.
To narrow the search result to the company you are looking for, you can use an extended search.
These can contain more informated queries.

## Steps

Use atreugo to implement an API with two routes.
The first should execute the simple search, using the name from the path as query.
The second should execute the extended search. The query comes from the body of the POST request.

1. Get familiar with atreugo. You can find the documentation [here](https://pkg.go.dev/github.com/savsgio/atreugo/v11)
2. The functions, that are called by atreugo should include these steps:
    1. Read the parameter or body from the request.
    2. Check the response for failures
    3. Return the result from the request as `SearchResponse`
3. Log all incoming requests to console
4. Log all errors to console
