.PHONY: default
default: run

App :=
Token :=

client_id :=
client_secret :=
refresh :=
auth_code :=

list:
	@echo "Available commands"
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^.PHONY' | grep -v '^default:' | grep -v '^docker-' | awk '/^docker:/ {print "docker:"; next} {print}'

doc:
	godoc -http=:1337 -goroot=.

docker-build:
	docker build --no-cache -t shikimori-docs -f docker/Dockerfile .

docker-start:
	docker run -d -p 1337:1337 shikimori-docs

docker: docker-build docker-start

test:
	go test -v
	go test -v ./internal/request
	go test -v ./internal/concatination
	go test -v ./neko
	go test -v ./genres
	go vet .

bench:
	cd bench && go test -bench . -benchmem

fmt:
	gofmt -w .

graphql-request:
	curl -X POST https://shikimori.io/api/graphql \
	-H "User-Agent: $(App)" \
	-H "Authorization: Bearer $(Token)" \
	-H 'Content-Type: application/json' \
	-d '{"query": "{ animes(search: \"initial d first stage\", limit: 1) { id name russian english japanese score status episodes description } }"}'

request:
	curl -X GET https://shikimori.io/api/animes?search=death+note&genre \
	-H "User-Agent: $(App)" \
	-H "Authorization: Bearer $(Token)"

token:
	curl -X POST "https://shikimori.io/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="refresh_token" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F refresh_token="$(refresh)"

auth:
	curl -X POST "https://shikimori.io/oauth/token" \
	-H "User-Agent: $(App)" \
	-F grant_type="authorization_code" \
	-F client_id="$(client_id)" \
	-F client_secret="$(client_secret)" \
	-F code="$(auth_code)" \
	-F redirect_uri="urn:ietf:wg:oauth:2.0:oob"
