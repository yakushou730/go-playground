tag-service 使用 gRPC 去打 blog-service

tag-service gRPC port 50051

blog-service http port 8000

```
$ brew install grpcui
$ grpcui -plaintext localhost:50051
```

swagger-ui

https://github.com/swagger-api/swagger-ui

put all /dist file into local project third_party/swagger-ui


```
go get -u github.com/jteeuwen/go-bindata/... 
go-bindata --nocompress -pkg swagger -o pkg/ui/data/swagger/data.go third_party/swagger-ui/...

go get -u github.com/elazarl/go-bindata-assetfs/...
```

http://localhost:8004/swagger-ui/

在輸入框中輸入
http://localhost:8004/swagger/tag.swagger.json
