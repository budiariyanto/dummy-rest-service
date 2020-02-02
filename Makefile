build-linux:
	env GOOS=linux GOARCH=amd64 go build -o dummy-rest-service-linux
build-docker: build-linux
	docker build . -t dummy-rest-service
run-docker: 
	docker run --rm --name dummy-rest-service -d -p 3456:3456 dummy-rest-service
stop-docker:
	docker stop dummy-rest-service
delete-docker:
	docker rmi -f dummy-rest-service
