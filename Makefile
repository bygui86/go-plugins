
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

build :		## Build application and plugins
	go build
	go build -buildmode=plugin -o ./plugins/fedex.so fedex/fedex.go
	go build -buildmode=plugin -o ./plugins/royalmail.so royalmail/royalmail.go
	go build -buildmode=plugin -o ./plugins/mydelivery.so mydelivery/mydelivery.go

# command: go run main.go {shipping method} {weight}

run-fedex : build		## Run calculation using fedex
	SHIPPING_METHOD=fedex \
	PACKAGE_WEIGHT=5 \
	go run main.go

run-royalmail : build		## Run calculation using royalmail
	SHIPPING_METHOD=royalmail \
	PACKAGE_WEIGHT=5 \
	go run main.go

run-mydelivery : build		## Run calculation using mydelivery
	SHIPPING_METHOD=mydelivery \
	PACKAGE_WEIGHT=5 \
	go run main.go

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
