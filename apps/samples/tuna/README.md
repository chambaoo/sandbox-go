# tuna

## tree

.
└── tuna
    ├── app
    │   ├── go.mod       # module "example"
    │   └── hello.go     # main Go file
    └── lib
        └── menu.go      # "example/lib" package

## package

### lib
```sh
# cd ./lib
go mod init example/lib
```

### app

```sh
# cd ./app
go mod init example/app
go mod edit -replace example/lib=../lib
go get example/lib
```

## execute

```sh
# cd ./app
go run hello.go
```

