CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh "$(CURRENT_DIR)"


exp:
	export DBURL='postgres://postgres:BEKJONS@localhost:5432/items_i?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:BEKJONS@localhost:5432/items_i?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database 'postgres://postgres:BEKJONS@localhost:5432/items_i?sslmode=disable' -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_users

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

swag-gen:
	 ~/go/bin/swag init -g ./api/routes.go -o api/docs