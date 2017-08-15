APEX=$(shell which apex)
target=slackbot

setup:
	cp .env.sample .env

install:
	curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sh

list/function:
	$(APEX) list

deploy:
	source .env && $(APEX) deploy $(target)

logs:
	source .env && $(APEX) logs $(target)

metrics:
	source .env && $(APEX) metrics $(target)

run:
	source .env && $(APEX) invoke $(target) < params.json

run/local:
	cat params.json | go run functions/slackbot/main.go
