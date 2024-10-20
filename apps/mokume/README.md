# mokume

## package

### src

```sh
# cd ./src/lib
go mod init mokume/lib
```

```sh
# cd ./src
go mod init mokume
go mod edit -replace mokume/lib=./lib
go get mokume/lib
```

## execute

```sh
# cd ./src
go run app.go
```

