all:
	go get -d ./src/main/
	go build -o ./bin/isup ./src/main

install:
	cp ./bin/isup /usr/bin
	
uninstall:
	rm /usr/bin/isup