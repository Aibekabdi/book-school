compose-build:
	go mod vendor
	docker-compose build

compose-up:
	docker-compose up

swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/api/main.go

ip:
	docker inspect book-school_web_1 --format='{{ json .NetworkSettings }}' | jq .Networks

open:
	docker exec -it book-school_db_1 psql -U postgres