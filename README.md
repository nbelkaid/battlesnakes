# Technical Test BattleSnakes
## QUICKSTART
Install docker (and docker-compose) if it's not already installed 
Run this command in the root of the repository
```sh
make init
``` 

When the infrastructure is well started you can also run this command to start a solo game
```sh
make test_solo_game
``` 
Please use the binary join in /bin to run the game locally. I had to make a tiny change to the Official project to let my /end endpoint be called after the game is ended. I wasn't receiving the call.

## Technical Test BattleSnake
First project using gORM and Docker
For more information about battleSnake read https://play.battlesnake.com/


## What's inside
1. API with all the routes needed to play to BattleSnake 
    - GET /game 
    - POST /game/start 
    - POST /game/move 
    - POST /game/move
    
2. Database A Postgres Database used to store some informations about the games played by our API Because first of all before to improve our algorithm we need to see how efficient it is!


## TODO
A lot of things still need to be improved !
Here some ideas to implement !

1. More and more tests !
2. A Seeder for the database container after his initialization and populate with some fake history (Algorithm_version = 0)
3. Implement metrics endpoints (/api/metrics) to be able to follow efficiency of Algorithm
4. Improve algorithm
5. Make the database optional in /api  and run an other container API without the database to be able to run multiplayer game locally with 2 snakes
...
...