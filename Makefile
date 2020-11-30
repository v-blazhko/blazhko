api-generate:
	oapi-codegen -generate spec -package api ./backend/openapi.yaml > ./backend/api/spec.go
	oapi-codegen -generate types -package api ./backend/openapi.yaml > ./backend/api/types.go
	oapi-codegen -generate chi-server -package api ./backend/openapi.yaml > ./backend/api/server.go


