.PHONY: gen
gen: genproto genopenapi

#Create gRPC code such as Service interface and Service struct
.PHONY: genproto
genproto:
	@./scripts/genproto.sh

#Create openapi code
.PHONY: genopenapi
genopenapi:
	@./scripts/genopenapi.sh


.PHONY: run air_stock air_order air_payment


