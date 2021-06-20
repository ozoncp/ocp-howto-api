PHONY: .generate
.generate:
	protoc -I ./vendor.protogen -I ./api/ocp-howto-api \
			--go_out=pkg/ocp-howto-api --go_opt=paths=import \
			--go-grpc_out=pkg/ocp-howto-api --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg/ocp-howto-api \
			--grpc-gateway_opt=logtostderr=true \
			--grpc-gateway_opt=paths=import \
			--validate_out lang=go:pkg/ocp-howto-api \
			--swagger_out=allow_merge=true,merge_file_name=pkg/ocp-howto-api/api:. \
			./api/ocp-howto-api/ocp-howto-api.proto &&\
	mv ./pkg/ocp-howto-api/github.com/ozoncp/ocp-howto-api/pkg/ocp-howto-api/* ./pkg/ocp-howto-api/ &&\
	rm -r "./pkg/ocp-howto-api/github.com"
	
.PHONY: .install-deps
.install-deps:
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/envoyproxy/protoc-gen-validate
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install github.com/envoyproxy/protoc-gen-validate

.PHONY: .vendor-proto
.vendor-proto:
	@if [ ! -d vendor.protogen ]; then \
		mkdir -p vendor.protogen ;\
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/github.com/envoyproxy/ &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate &&\
		mv vendor.protogen/github.com/envoyproxy/protoc-gen-validate/validate/ vendor.protogen/ &&\
		rm -rf vendor.protogen/github.com;\
	fi
