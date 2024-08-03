#!make

generate-model:
	@echo "generate model"
	@sqlboiler psql

test-model:
	@echo "test model"
	@go test -v ./internal/models
