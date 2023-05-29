#!/bin/sh

proto_source_dir="assets/static/proto"

install_protoc()
{
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	echo
	echo "Download protoc-gen-grpc and protoc. Put them on PATH. Linux example:"
	echo "export BINDIR=${HOME}/bin"
	echo "mkdir \${BINDIR}"
	echo "curl -sL -o protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip"
	echo "unzip protoc.zip"
	echo "cp protoc/bin/* \${BINDIR}"
}

generate_go()
{
    go_proto_out_dir="pb"
	echo "Generating Go stubs..."
	mkdir -p ${go_proto_out_dir} || echo
	protoc -I ${proto_source_dir} ${proto_source_dir}/*.proto --proto_path=${go_proto_out_dir} --go_out=${go_proto_out_dir} --go-grpc_out=${go_proto_out_dir}
{{ if .HasGoGrpcSdk }}
	echo "Copying Go stubs to SDK Go dir..."
	cp -r ${go_proto_out_dir} sdk/go_grpc/proto
{{ end }}
}

case "${1}" in
    install) install_protoc ;;
    generate) generate_go ;;
    *) echo "Invalid option to generate gRPC files" ;;
esac

