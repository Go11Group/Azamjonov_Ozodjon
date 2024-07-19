CURRENT_DIR=$(shell pwd)
DBURL="postgres://postgres:BEKJONS@localhost:5432/user_i?sslmode=disable"
exp:
	export DBURL="postgres://postgres:BEKJONS@localhost:5432/user_i?sslmode=disable"

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down

mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}