.PHONY: dev packages client echo

packages:
	@dep ensure

dev:
	@realize start --name="hex-game-server"

echo:
	@realize start --name="hex-game-client"