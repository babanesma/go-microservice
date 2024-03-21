.DEFAULT_GOAL := swagger

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models