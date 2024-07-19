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
	migrate create -ext sql -dir migrations -seq insert_table

mig-force:
	migrate -path migrations -database 'postgres://postgres:BEKJONS@localhost:5432/items_i?sslmode=disable' force 1
