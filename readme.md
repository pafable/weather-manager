# Weather Manager

The purpose of this project is to explore creating commandline tools using Go.

Create an account on [WeatherAPI](https://weatherapi.com) and set your api key as an environment variable.

| Env Var |    Value     | Required |
|:-------:|:------------:|:--------:|
| APIKEY  | YOUR-API-KEY |   True   |

```
go run main.go temperature --location <zip_code> 
```

```
go run main.go condition --location <zip_code>
```

### Building a binary

MacOS arm64

```
GOOS=darwin GOARCH=arm64 go build -o wm
```

MacOS x86_64

```
GOOS=darwin GOARCH=amd64 go build -o wm
```

Linux x86_64

```
GOOS=linux GOARCH=amd64 go build -o wm
```

Windows x86_64

```
GOOS=windows GOARCH=amd64 go build -o wm
```
