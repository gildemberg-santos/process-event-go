Start:
	@echo "Starting the server"
	rm -rf db/*.db
	go run cmd/main.go
