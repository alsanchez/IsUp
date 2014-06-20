all:
	go get -d ./src/main/
	go build -o ./bin/IsUp ./src/main