# 创建rpcgo 目录和初始化
    mkdir -p rpcgo
    go mod init github.com/AIntelligenceGame/rpcgo

# 生成代码
    make proto_gen

# run server/client
    go run server/main.go
    go run client/main.go
 
# bufbuild 管理protobuf文件
    go install github.com/bufbuild/buf/cmd/buf@latest
    buf --version
    cd pb
    buf mod init

