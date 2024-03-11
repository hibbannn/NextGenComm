.PHONY: stub run

stub:
ifeq ($(do),generate)
	@echo "Running generate-stub.sh"
	./script/generate-stub.sh
else ifeq ($(do),import)
	@echo "Running import-stubs.sh"
	./script/proto/import-stubs.sh
else
	@echo "Invalid command. Use 'make run stub=generate' or 'make run stub=import'"
endif

run:
	go run ./cmd/main.go

