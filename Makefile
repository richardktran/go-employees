build:
	@go build -o employee cmd/main.go
	@./employee
sync-data:
	@go run cmd/syncdata/main.go