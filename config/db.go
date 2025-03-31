package config

import "time"

var (
	DatabaseUser     = "dev"
	DatabasePassword = "dev"
	DatabaseName     = "starter"
	DatabaseHost     = "127.0.0.1"
	DatabasePort     = "5432"

	DatabaseURI     = ""
	DatabaseURIDev  = "postgresql://dev:dev@127.0.0.1/starter"
	DatabaseURIProd = "postgresql://root@127.0.0.1:26257/starter?sslmode=disable"
	DefaultTimeout  = time.Second * 10
	SqliteDSN       = "file:ent.db?mode=memory&cache=shared&_fk=1"
)
