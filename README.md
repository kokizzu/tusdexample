# tusdexample

1. clone tusd `git clone git@github.com:tus/tusd.git`
2. build the executable:
```shell
cd tusd
CGO_ENABLED=0 GOOS=linux go build -o tusd cmd/tusd/main.go
cd ..
```
3. run container:
```shell
APPNAME=tusdserver
VERSION=$(ruby -e 't = Time.now; print "v1.#{t.month+(t.year-2021)*12}%02d.#{t.hour}%02d" % [t.day, t.min]')
COMMIT=$(git rev-parse --verify HEAD)
docker image build -f 'Dockerfile' . \
  --build-arg "app_name=$APPNAME" \
  -t "${APPNAME}:latest" \
  -t "$APPNAME:$COMMIT" \
  -t "$APPNAME:$VERSION"
 
docker run --network="host" tusdserver 
```
4. run httpmain.go
```shell
go run httpmain.go
```
5. upload using tusjsserver
```shell
# first time only
npm install --save tus-js-client

node upload-to-tus-server.js
```

it should receive something like this:
```go
 map[string]interface {}{
    "HTTPRequest": map[string]interface {}{
        "Header": map[string]interface {}{
            "Connection": []interface {}{
                "close",
            },
            "Content-Length": []interface {}{
                "695",
            },
            "Content-Type": []interface {}{
                "application/offset+octet-stream",
            },
            "Tus-Resumable": []interface {}{
                "1.0.0",
            },
            "Upload-Offset": []interface {}{
                "0",
            },
        },
        "Method":     "PATCH",
        "RemoteAddr": "127.0.0.1:47716",
        "URI":        "/files/47db1645c5aac87dc05fe05dcf9c5e41",
    },
    "Upload": map[string]interface {}{
        "ID":        "47db1645c5aac87dc05fe05dcf9c5e41",
        "IsFinal":   bool(false),
        "IsPartial": bool(false),
        "MetaData":  map[string]interface {}{
            "filename": "README.md",
            "filetype": "text/plain",
        },
        "Offset":         float64(695),
        "PartialUploads": nil,
        "Size":           float64(695),
        "SizeIsDeferred": bool(false),
        "Storage":        map[string]interface {}{
            "Path": "/data/47db1645c5aac87dc05fe05dcf9c5e41",
            "Type": "filestore",
        },
    },
}
```

6. uncomment grpc section on Dockerfile, rebuild (see number 3)
7. compile proto
```shell
protoc tusdhooks/tusdhooks.proto \
  --go_out=.. \
  --proto_path=. \
  --go-grpc_out=..
```

8. run the grpcmain:
```shell
go run grpcmain.go
```

9. repeat number 5 to resend

it should show something like this:
```go
 &tusdhooks.Hook{
    state:         impl.MessageState{},
    sizeCache:     0,
    unknownFields: nil,
    Upload:        &tusdhooks.Upload{
        state:          impl.MessageState{},
        sizeCache:      0,
        unknownFields:  nil,
        Id:             "7a648ca230da5c28e48182a4d775cbcf",
        Size:           2355,
        SizeIsDeferred: false,
        Offset:         2355,
        MetaData:       {"filename":"README.md", "filetype":"text/plain"},
        IsPartial:      false,
        IsFinal:        false,
        PartialUploads: nil,
        Storage:        {"Path":"/data/7a648ca230da5c28e48182a4d775cbcf", "Type":"filestore"},
    },
    HttpRequest: &tusdhooks.HTTPRequest{
        state:         impl.MessageState{},
        sizeCache:     0,
        unknownFields: nil,
        Method:        "PATCH",
        Uri:           "/files/7a648ca230da5c28e48182a4d775cbcf",
        RemoteAddr:    "127.0.0.1:47810",
    },
    Name: "post-receive",
}
```


Flow:
```
upload-to-tus-server.js --> localhost:1080 (Docker) --> callback to 8081 (http) or 8083 (grpc) --> httpmain.go or grpcmain.go

```
