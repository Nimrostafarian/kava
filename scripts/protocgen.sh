  #!/usr/bin/env bash

# Adapted from https://github.com/cosmos/cosmos-sdk/blob/master/scripts/protocgen.sh

set -eo pipefail

protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the root kava folder."
    return 1
  fi

  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
}

protoc_gen_doc() {
  go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc 2>/dev/null
  go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
}

cp go.mod go.mod.bak

protoc_gen_gocosmos
protoc_gen_doc

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  buf alpha protoc \
    -I "proto" \
    -I "third_party/proto" \
    --gocosmos_out=plugins=grpc,\
Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. \
    --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done

mkdir -p ./docs/core

# command to generate docs using protoc-gen-doc
buf alpha protoc \
  -I "proto" \
  -I "third_party/proto" \
  --plugin=/go/bin/protoc-gen-doc \
  --doc_out=./docs/core \
  --doc_opt=./docs/protodoc-markdown.tmpl,proto-docs.md \
  $(find "$(pwd)/proto" -maxdepth 5 -name '*.proto')

mv go.mod.bak go.mod
go mod tidy

# generate codec/testdata proto code
# buf protoc -I "proto" -I "third_party/proto" -I "testutil/testdata" --gocosmos_out=plugins=grpc,\
# Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types:. ./testutil/testdata/*.proto

# move proto files to the right places
cp -r github.com/kava-labs/kava/* ./
rm -rf github.com
