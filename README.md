# 创建rpcgo 目录和初始化
    mkdir -p rpcgo
    go mod init github.com/shangguanairen/rpcgo

# 生成代码
    make proto_gen

# run server/client
    go run server/main.go
    go run client/main.go
 
 