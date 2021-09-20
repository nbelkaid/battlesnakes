## start environment from scratch
init: destroy build up sleep test_solo_game

## stop
down:
	docker-compose down

## start
up:
	docker-compose up -d

## stop and remove docker images
destroy:
	docker-compose down -v --remove-orphans

## create builder image
build:
	docker-compose build

## sleep 10 seconds
sleep:
	echo "Sleep 10 seconds to be sure DB init" && sleep 10

##Test Solo Game BattleSnake
test_solo_game:
	./bin/battlesnake play -W 50 -H 30 --name "HelloWorld" --url "http://0.0.0.0:8080/game/" -g solo  -v