export DB_USERNAME=root
export DB_PASSWORD=root
export DB_PORT=3306
export DB_HOST=192.168.0.103
export DB_NAME=project_aka
export MYSQL_URI="mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

create_migration:
	migrate create -ext sql -dir db/migrations $(name)

migrate:
	migrate -database ${MYSQL_URI} -source file://db/migrations up