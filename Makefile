.PHONY: openapi
openapi:
	oapi-codegen -config api/openapi/restapi.cfg.yaml api/openapi/restapi.yaml
	
.PHONY: lint
lint:
	golangci-lint runs
	