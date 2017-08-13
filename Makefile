APEX=$(shell which apex)
target=hello

setup:
	cp .env.sample .env

install:
	curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sh

list/function:
	$(APEX) list

deploy:
	$(APEX) deploy $(target)

logs:
	$(APEX) logs $(target)

metrics:
	$(APEX) metrics $(target)

run:
	$(APEX) invoke $(target) < params.json

