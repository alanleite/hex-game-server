.PHONY: dev dependencies client

dependencies:
	@dep ensure

dev:
	@realize start

client:
	@go run cmd/hex-game-client/main.go