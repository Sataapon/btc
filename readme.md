### Prerequisite
- PostgreSQL 14.2
- Go 1.16

### Usage
```
. ./script/env.sh

make migrate

make server

. ./script/save_record.sh

. ./scripts/get_history.sh
```