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

#### ----------------------- Run command configuration -----------------------  ####

# - configPath: path to the configuration environment file
# - migrate: true or false (run database migration)
prod:
	@GIN_MODE=release go run cmd/trekkstay/main.go -conf=${configPath} -migrate=${migrate}

dev:
	@GIN_MODE=debug go run cmd/trekkstay/main.go -conf=./env/dev.env -migrate=true

#### ----------------------- Swagger command configuration ----------------------- ####
gen_swagger:
	swag init -g cmd/trekkstay/main.go
