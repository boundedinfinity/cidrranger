makefile_dir 	:= $(abspath $(shell pwd))

project_name	:= cidrranger

docker_service 	:= bi-$(project_name)
docker_tag		:= $(shell cat docker-compose.yml | grep image | sed 's/.*image://' | tr -d ' ')

go_package		:= github.com/boundedinfinity/$(project_name)

.PHONY: list

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

bootstrap:
	@make bower-bootstrap
	@make go-bootstrap

purge:
	@make bower-purge
	@make go-purge

clean:
	@make go-clean

docker-tag:
	@echo $(docker_tag)

docker-up:
	docker-compose up $(docker_service)

docker-stop:
	docker-compose stop $(docker_service)

docker-build:
	docker-compose build $(docker_service)

docker-bash:
	docker-compose run --rm $(docker_service) bash

docker-dev:
	docker-compose -f $(makefile_dir)/docker-compose-dev.yml run --rm $(docker_service) bash

docker-push:
	docker push $(docker_tag)

go-bootstrap:
	glide install

go-clean:
	go clean

go-purge:
	@make go-clean
	rm -rf $(makefile_dir)/vendor
	rm -rf $(makefile_dir)/glide.lock

go-test:
#	go test $$(go list $(go_package)/... | grep -v vendor)
	go test -v $$(go list $(go_package)/... | grep -v vendor)

go-test-watch:
	CompileDaemon -command "make go-test"

go-build:
	go build $(go_package)

serve:
	cd $(makefile_dir)/static && python -m SimpleHTTPServer

bower-bootstrap:
	bower install

bower-purge:
	rm -rf $(makefile_dir)/$(shell cat $(makefile_dir)/.bowerrc | jq -r .directory)

compile-watch:
	CompileDaemon -command "$(makefile_dir)/cidrranger"
