#!/bin/sh

proto_source_dir="assets/static/proto"
grpc_web_proto_out_dir="sdk/js_grpcweb"
grpc_web_version="1.3.1"

function install_protoc {
	echo "curl -sL -o protoc-gen-grpc-web https://github.com/grpc/grpc-web/releases/download/${grpc_web_version}/protoc-gen-grpc-web-${grpc_web_version}-linux-x86_64"
	echo "cp protoc-gen-grpc-web \${BINDIR}"
}

function generate_js {
	echo "Generating JS stubs for gRPC Web..."
	protoc -I ${proto_source_dir} ${proto_source_dir}/*.proto --js_out=import_style=commonjs,binary:${grpc_web_proto_out_dir} --grpc-web_out=import_style=commonjs,mode=grpcweb:${grpc_web_proto_out_dir}
	echo "Done."
}

case "${1}" in
    install) install_protoc ;;
    generate) generate_js ;;
    *) echo "Invalid option to generate gRPC files" ;;
esac

