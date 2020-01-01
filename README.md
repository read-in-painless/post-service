# Post Service
This is a post service. It just manages the posts by sending the user. Write, read, delete, publish in its scope. 


## How to build proto file
```
$ protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/post.proto
```