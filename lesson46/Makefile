CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:pass@localhost:5432/go10?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:pass@localhost:5432/go10?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_courier

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table
#   rm -r db/migrations