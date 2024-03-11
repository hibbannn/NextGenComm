# RUN FROM HERE ONLY FOR THE FIRST TIME OR IF YOU DON'T HAVE THE BIN FILES
#go install \
#    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
#    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest \
#    google.golang.org/protobuf/cmd/protoc-gen-go@latest \
#    google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
#    github.com/infobloxopen/protoc-gen-gorm@latest
#    github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc


# RUN HERE WHEN YOU MAKE CHANGES TO THE PROTO FILES
protoc  --proto_path=./proto \
        --proto_path=./plugin \
        --plugin=protoc-gen-go="$(go env GOPATH)"/bin/protoc-gen-go \
        --plugin=protoc-gen-go-grpc="$(go env GOPATH)"/bin/protoc-gen-go-grpc \
        --plugin=protoc-gen-grpc-gateway="$(go env GOPATH)"/bin/protoc-gen-grpc-gateway \
        --plugin=protoc-gen-openapiv2="$(go env GOPATH)"/bin/protoc-gen-openapiv2 \
        --go_out=paths=source_relative:./stub/serviceUser \
        --go-grpc_out=paths=source_relative:./stub/serviceUser \
        --grpc-gateway_out=paths=source_relative:./stub/serviceUser \
        ./proto/user_route.proto


protoc  --proto_path=./proto \
        --proto_path=./plugin \
        --plugin=protoc-gen-go="$(go env GOPATH)"/bin/protoc-gen-go \
        --plugin=protoc-gen-go-grpc="$(go env GOPATH)"/bin/protoc-gen-go-grpc \
        --plugin=protoc-gen-gorm="$(go env GOPATH)"/bin/protoc-gen-gorm \
        --plugin=protoc-gen-grpc-gateway="$(go env GOPATH)"/bin/protoc-gen-grpc-gateway \
        --plugin=protoc-gen-openapiv2="$(go env GOPATH)"/bin/protoc-gen-openapiv2 \
        --go_out=paths=source_relative:./stub/serviceUser \
        --gorm_out=paths=source_relative:./stub/serviceUser \
        ./proto/user_model.proto


protoc  --proto_path=./proto \
        --proto_path=./plugin \
        --plugin=protoc-gen-go="$(go env GOPATH)"/bin/protoc-gen-go \
        --plugin=protoc-gen-go-grpc="$(go env GOPATH)"/bin/protoc-gen-go-grpc \
        --plugin=protoc-gen-grpc-gateway="$(go env GOPATH)"/bin/protoc-gen-grpc-gateway \
        --plugin=protoc-gen-openapiv2="$(go env GOPATH)"/bin/protoc-gen-openapiv2 \
        --go_out=paths=source_relative:./stub/serviceUser \
        --go-grpc_out=paths=source_relative:./stub/serviceUser \
        --grpc-gateway_out=paths=source_relative:./stub/serviceUser \
        ./proto/user_request.proto \
        ./proto/user_response.proto

protoc   --proto_path=./proto \
         --proto_path=./plugin \
         --plugin="$(go env GOPATH)"/bin/protoc-gen-openapiv2 \
         --openapiv2_out ./doc \
         --openapiv2_opt logtostderr=true,repeated_path_param_separator=ssv \
         ./proto/user_route.proto

mv ./doc/user_route.swagger.json ./doc/openapi/swagger.json

protoc    --proto_path=./proto \
          --proto_path=./plugin \
          --plugin=protoc-gen-doc="$(go env GOPATH)"/bin/protoc-gen-doc \
          --doc_out=./doc/api --doc_opt=html,index.html \
          ./proto/user_route.proto \
          ./proto/user_request.proto \
          ./proto/user_model.proto \
          ./proto/user_response.proto