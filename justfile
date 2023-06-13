set export

PWD := `pwd`
SERVER_HOME := PWD / "server"
WEB_HOME := PWD / "web"
DATA_HOME := PWD / "data"
SQL_CONFIG := DATA_HOME / "sqlboiler.toml"
CLIENT_DIST_DIR := WEB_HOME / "dist"
CODEGEN_DIR := SERVER_HOME / "models"

setup-codegen:
	go install github.com/volatiletech/sqlboiler/v4@v4.14.2
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-sqlite3@v4.14.2
lint:
	cd ${SERVER_HOME} && staticcheck ./...
serve:
	cd ${SERVER_HOME} && go run cmd/server/main.go
migrate:
	cd ${SERVER_HOME} && go run cmd/migrations/main.go
codegen:
	cd ${DATA_HOME} && sqlboiler -o ${CODEGEN_DIR} sqlite3
	find ${CODEGEN_DIR} -name '*test.go' -delete
setup-web:
	cd ${WEB_HOME} && npm i
lint-web:
	cd ${WEB_HOME} && npm run lint
prettify-web:
	cd ${WEB_HOME} && npm run prettify
build-web:
	rm -rf ${CLIENT_DIST_DIR}
	cd ${WEB_HOME} && npm run build
ci: migrate setup-codegen codegen lint setup-web lint-web build-web
setup-to-serve: ci serve