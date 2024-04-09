#### ----------------------- Database command configuration ----------------------- ####
init_db:
	@docker run \
		--name trekkstay-db \
		--restart always \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=0000 \
		-e POSTGRES_DB=trekkstay_db\
		-v ./.volumes/trekkstay-volume:/var/lib/postgresql/data \
		-p 5432:5432 -it \
		-d postgres:latest
rm_db:
	@docker rm -f trekkstay-db

install_pgadmin:
	@docker run -p 8000:80 \
		--name pg_admin \
		--restart always \
        -e 'PGADMIN_DEFAULT_EMAIL=admin@trekkstay.com' \
        -e 'PGADMIN_DEFAULT_PASSWORD=000000' \
        -v 	./.volumes/pg_admin:/var/lib/pgadmin/data \
        -d dpage/pgadmin4

install_redis:
	@docker run -d --name redis -p 6379:6379 -p 8001:8001 --network trekkstay_network redis/redis-stack:latest

#### ----------------------- Database command configuration ----------------------- ####
create_network:
	@docker network create trekkstay_network

#### ----------------------- Run command configuration -----------------------  ####

# - configPath: path to the configuration environment file
# - migrate: true or false (run database migration)
prod:
	@GIN_MODE=release go run cmd/trekkstay/main.go -conf=./env/prod.env -migrate=false

dev:
	@GIN_MODE=debug go run cmd/trekkstay/main.go -conf=./env/dev.env -migrate=true

compose-prod-up:
	@docker compose -f docker-compose.prod.yml up -d

compose-dev-up:
	@docker compose -f docker-compose.dev.yml up -d

compose-prod-down:
	@docker rmi -f trekkstay-backend
	@docker compose -f docker-compose.prod.yml down

compose-dev-down:
	@docker rmi -f trekkstay-backend
	@docker compose -f docker-compose.dev.yml down

push-image:
	# Remove old container and image
	@docker rm -f trekkstay-backend
	@docker rmi -f thanhanphan17/trekkstay-backend
	# Build and push new image
	@docker build -t thanhanphan17/trekkstay-backend .
	@docker push thanhanphan17/trekkstay-backend

run-image:
	@docker rm -f trekkstay-backend
	@docker rmi -f thanhanphan17/trekkstay-backend
	@docker container run \
		--restart unless-stopped \
		--env CONFIG_PATH=./env/prod.env \
		--env MIGRATE=false \
		--name trekkstay-backend \
		--network trekkstay_network \
		-dp 8888:8888 \
		thanhanphan17/trekkstay-backend

#### ----------------------- Swagger command configuration ----------------------- ####
gen_swagger:
	swag init -g cmd/trekkstay/main.go
