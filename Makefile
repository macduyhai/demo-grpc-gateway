gen-proto:
	protoc gatewayproto/gatewaypb.proto --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.
gen-gateway:
	protoc gatewayproto/gatewaypb.proto --grpc-gateway_out=logtostderr=true,paths=source_relative:.
gen-readme:
	protoc --doc_out=. --doc_opt=markdown,README.md gatewayproto/gatewaypb.proto