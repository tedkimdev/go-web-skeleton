# go-web-skeleton

### Libraries

- Chi (http router)
- oapi-codegen (open api tool)
- Sqlx (database)
- Goose (postgres db migration tool)
- slog (standard library slog)

### DB Migration

- goose command
  - create a new sql migration
    - goose -s create "filename" sql
  - migration to up
    - goose -dir db/postgres/migrations postgres "db_connection_string" up
  - migration to down
    - goose -dir db/postgres/migrations postgres "db_connection_string" down

### oapi-codeegen

- `make openapi`

### Ref

- project layout
  - https://github.com/golang-standards/project-layout

- openapi yaml example
  - https://github.com/OAI/OpenAPI-Specification/blob/main/examples/v3.0/petstore-expanded.yaml
  - https://github.com/deepmap/oapi-codegen

- api documentation
  - https://redocly.com/

- sqlx
  - https://jmoiron.github.io/sqlx/
  - https://jmoiron.github.io/sqlx/#advancedScanning
  - http://go-database-sql.org/modifying.html

- xid
  - https://supabase.com/blog/choosing-a-postgres-primary-key
  - https://github.com/modfin/pg-xid

