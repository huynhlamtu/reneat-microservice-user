build-app:
	docker-compose build app
start:
	docker-compose up -d
logs:
	docker logs -f reneat-microservice-user
ssh-app:
	docker exec -it reneat-microservice-user sh
ps:
	docker-compose ps
swagger:
	rm -rf docs/ && swag fmt && swag init ./controllers/*
proto-user:
	protoc -I grpc/proto/user/ \
		-I /usr/include \
		--go_out=paths=source_relative,plugins=grpc:grpc/proto/user/ \
		grpc/proto/user/user.proto
proto-auth:
	protoc -I grpc/proto/auth/ \
		-I /usr/include \
		--go_out=paths=source_relative,plugins=grpc:grpc/proto/auth/ \
		grpc/proto/auth/auth.proto
stop: #down
	@echo "=============cleaning up============="
	docker-compose down
	docker system prune -f
	docker volume prune -f