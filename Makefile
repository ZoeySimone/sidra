sidra:
	go build -o ./bin/protoc-gen-$@ .

.PHONY: sidra

install: sidra
	mv ./bin/protoc-gen-sidra ~/go/bin/protoc-gen-sidra

generate-new: sidra install
	protoc ./example-api/example.proto --sidra_out=. --sidra_opt=no-grpc  