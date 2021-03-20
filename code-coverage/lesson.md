## Code coverage - The command you will need

Write to a file

```sh
go test -coverprofile=coverage.out
```

Analyse the result in browser

```sh
go tool cover -html=coverage.out
```

Generate an html file

```sh
go tool cover -html=coverage.out -o coverage.html
```
