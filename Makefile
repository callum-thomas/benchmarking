.PHONY: run-go

run-go:
	docker build -t go-service ./go && docker run -p 8080:8080 go-service

run-nestjs:
	docker build -t nestjs-service ./nestjs && docker run -p 3000:3000 nestjs-service
