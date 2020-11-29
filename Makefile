oapi-codegen -generate spec -package api openapi.yaml > api/spec.go
oapi-codegen -generate types -package api openapi.yaml > api/types.go
oapi-codegen -generate chi-server -package api openapi.yaml > api/server.go
