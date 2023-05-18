#!/bin/sh

proto_source_dir="assets/static/proto"

function install_protoc {
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	echo
	echo "Download protoc-gen-grpc and protoc. Put them on PATH. Linux example:"
	echo "export BINDIR=${HOME}/bin"
	echo "mkdir ${BINDIR}"
	echo "curl -sL -o protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/${grpc_web_version}/protoc-gen-grpc-web-${grpc_web_version}-linux-x86_64"
	echo "cp protoc-gen-grpc-web ${BINDIR}"
	echo "curl -sL -o protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.20.1/protoc-3.20.1-linux-x86_64.zip"
	echo "unzip protoc.zip"
	echo "cp protoc/bin/* ${BINDIR}"
}

function generate_go {
    go_proto_out_dir="pb"
	echo "Generating Go stubs..."
	mkdir -p ${go_proto_out_dir} || echo
	protoc -I ${proto_source_dir} ${proto_source_dir}/*.proto --proto_path=${go_proto_out_dir} --go_out=${go_proto_out_dir} --go-grpc_out=${go_proto_out_dir}
	echo "Copying Go stubs to SDK Go dir..."
	cp -r ${go_proto_out_dir} sdk/go_grpc/proto
}

function generate_js {
    grpc_web_proto_out_dir="sdk/js_grpcweb"
    grpc_web_version="1.3.1"
	echo "Generating JS stubs for gRPC Web..."
	protoc -I ${proto_source_dir} ${proto_source_dir}/*.proto --js_out=import_style=commonjs,binary:${grpc_web_proto_out_dir} --grpc-web_out=import_style=commonjs,mode=grpcweb:${grpc_web_proto_out_dir}
	echo "Done."
}

case "${1}" in
    install) install_protoc ;;
    generate_go) generate_go ;;
    generate_js) generate_js ;;
    *) echo "Invalid option to generate gRPC files" ;;
esac

