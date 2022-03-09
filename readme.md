PostgreSQL 14.2 
go1.16

### Prerequisite
- PostgreSQL 14.2
- Go 1.16

### Usage
. ./script/env.sh

make migrate

make server

. ./script/save_record.sh
. ./scripts/gen_history.sh