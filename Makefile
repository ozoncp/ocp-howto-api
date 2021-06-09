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