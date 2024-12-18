include .env

inventory_db_create:
	echo "Creating db container.."
	docker run -d \
		--name ${INVENTORY_DB_CONTAINER} \
		-p 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=${INVENTORY_DB_ROOT_PASSWORD} \
		-e MYSQL_DATABASE=${INVENTORY_DB_NAME} \
		-e MYSQL_USER=${INVENTORY_DB_USER} \
		-e MYSQL_PASSWORD=${INVENTORY_DB_PASSWORD} \
		mysql:9.1.0

inventory_db_start:
	if docker inspect "${INVENTORY_DB_CONTAINER}" > /dev/null 2>&1; then \
		echo "Starting db container.."; \
		docker container start ${INVENTORY_DB_CONTAINER}; \
	else \
		echo "No docker containers found"; \
		$(MAKE) inventory_db_create; \
	fi

inventory_db_stop:
	if docker inspect "${INVENTORY_DB_CONTAINER}" > /dev/null 2>&1; then \
		echo "Stopping db container..."; \
		docker container stop ${INVENTORY_DB_CONTAINER}; \
	else \
		echo "No docker containers found"; \
	fi

inventory_db_delete:
	if docker inspect "${INVENTORY_DB_CONTAINER}" > /dev/null 2>&1; then \
		echo "Deleting db container..."; \
		docker container rm -f ${INVENTORY_DB_CONTAINER}; \
	else \
		echo "No docker containers found"; \
	fi

inventory_db_up:
	migrate -path ./services/inventory/db/migrations \
	-database "mysql://${INVENTORY_DB_USER}:${INVENTORY_DB_PASSWORD}@tcp(localhost:${INVENTORY_DB_PORT})/${INVENTORY_DB_NAME}" \
	up

inventory_db_down:
	migrate  -path ./services/inventory/db/migrations \
	-database "mysql://${INVENTORY_DB_USER}:${INVENTORY_DB_PASSWORD}@tcp(localhost:${INVENTORY_DB_PORT})/${INVENTORY_DB_NAME}" \
	down 

proto_def:
	protoc --proto_path="common/pb" \
		--go_out="common/pb" --go_opt=paths=source_relative \
		--go-grpc_out="common/pb" --go-grpc_opt=paths=source_relative \
		common/pb/ecomm.proto

build_rest_gateway:
	if [ -f "./bin/${REST_GATEWAY_BINARY}" ]; then \
		rm "./bin/${REST_GATEWAY_BINARY}"; \
		echo "Deleted ${REST_GATEWAY_BINARY}"; \
	fi

	if [ ! -d "./bin" ]; then \
		mkdir bin; \
	fi

	echo "Building ${REST_GATEWAY_BINARY} binary"
	go build -o ./bin/${REST_GATEWAY_BINARY} ./gateways/rest/cmd/rest/main.go

run_rest_gateway: build_rest_gateway
	PORT=${PORT} ./bin/${REST_GATEWAY_BINARY}


build_inventory_service:
	if [ -f "./bin/${INVENTORY_SERVICE_BINARY}" ]; then \
		rm "./bin/${INVENTORY_SERVICE_BINARY}"; \
		echo "Deleted ${INVENTORY_SERVICE_BINARY}"; \
	fi

	if [ ! -d "./bin" ]; then \
		mkdir bin; \
	fi

	echo "Building ${INVENTORY_SERVICE_BINARY} binary"
	go build -o ./bin/${INVENTORY_SERVICE_BINARY} \
	./services/inventory/cmd/inventory/main.go


run_inventory_service: build_inventory_service
	DATABASE_URL=${INVENTORY_DATABASE_URL} \
	SERVICE_URL=${INVENTORY_SERVICE_URL} \
	./bin/${INVENTORY_SERVICE_BINARY}

