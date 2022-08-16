dc := docker-compose -f ./docker-compose.yml

up:
	$(dc) up -d --build

down:
	$(dc) down

restart:
	$(dc) restart

reup:
	@make down
	@make up

rm:
	$(dc) down --rmi all

logs:
	$(dc) logs -f

app:
	$(dc) exec app /bin/sh

db:
	$(dc) exec db /bin/sh

.PHONY:	setup up down restart reup rm logs app db