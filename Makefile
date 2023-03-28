version := 0.3.0
containername=tcr

SERVICE_BUILD_PATH=tmp/main
HTTP_LISTEN_PORT=8080

.PHONY: run migrate paginate post_random_payload

staging := 34.89.169.189/api/v1
local := localhost:8080/api/v1
baseURL:=$(local)

ifneq (,$(wildcard ./.local.env))
    include .local.env
    export
endif

get:
	go get -d -v ./...

run:
	rm -f $(SERVICE_BUILD_PATH)
	go build -o $(SERVICE_BUILD_PATH) cmd/service/main.go
	chmod +x $(SERVICE_BUILD_PATH)

	HTTP_LISTEN_PORT=${HTTP_LISTEN_PORT} $(SERVICE_BUILD_PATH)

vet:
	go mod tidy
	go vet ./...
	go fmt ./...

docker:
	docker build -f Dockerfile -t  $(containername):$(version) .

docker_upload: docker
	docker tag $(containername):$(version) leondroid/$(containername):$(version)
	docker push leondroid/$(containername):$(version)

check-swagger:
	which swagger || go get github.com/go-swagger/go-swagger/cmd/swagger

swagger: 
	swagger generate spec -o swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger swagger.yaml

healthcheck:
	curl  $(baseURL)/healthcheck

post:
	curl -X POST $(baseURL)/tweet \
   		-H 'Content-Type: application/json' \
   		-d '{"payload":"my_logins"}' | jq .

paginate:
	curl -X GET $(baseURL)/tweet?offset=50\&limit=10 | jq .

migrate:
	cd migrate; goose  postgres "postgres://$(postgres_user):$(postgres_pwd)@$(postgres_url)/postgres?sslmode=disable" up

down:
	cd migrate; goose  postgres "postgres://$(postgres_user):$(postgres_pwd)@$(postgres_url)/postgres?sslmode=disable" down

status:
	cd migrate; goose postgres "postgres://$(postgres_user):$(postgres_pwd)@$(postgres_url)/postgres?sslmode=disable" status

recreate: down migrate jet

restart: recreate jet run

jet:
	jet -dsn=postgres://$(postgres_user):$(postgres_pwd)@$(postgres_url)/postgres?sslmode=disable -schema=tcr -path=./.gen

db_up:
	POSTGRES_PORT=$(postgres_port) POSTGRES_HOST=$(postgres_host) POSTGRES_DB=$(postgres_db) POSTGRES_PWD=$(postgres_pwd) POSTGRES_USER=$(postgres_user) docker-compose up -d

nuke:
	cd migrate
	docker-compose stop
	rm -rf pgdata
	$(MAKE) db_up
	
init_db:
	cd migrate
	mkdir pgdate
	$(MAKE) db_up
	$(MAKE) migrate

post_random_payload:
	curl -X POST $(baseURL)/tweet \
		-H 'Content-Type: application/json' \
		-d '{"payload":"$(shell openssl rand -hex 5)"}' | jq .

seed_tweets:
	@seq 20 | xargs -Iz $(MAKE) post_random_payload
