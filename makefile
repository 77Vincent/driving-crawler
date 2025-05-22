build:
	docker build -t myapp .
go:
	go run .
run:
	docker run -d --name myapp myapp
log:
	docker logs -f myapp
stop:
	docker stop myapp
