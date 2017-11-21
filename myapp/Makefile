# Go parameters


.PHONY: package
package: build_cli create

build_cli: build_myapp
	go build ./echo/echo_client/myapp-cli.go
	
build_myapp: 
	go build ./myapp.go

create:
	docker build -t myapp .

