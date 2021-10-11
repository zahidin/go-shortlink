dev:
	./bin/air -c .air.toml

build:
	go build -o ./dist/main ./src
	
start:
	chmod +x ./dist/main
	./dist/main