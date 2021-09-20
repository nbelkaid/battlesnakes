# battlesnakes
Technical Test BattleSnake

Wanted to discover some concept or libraries !
So first project using gORM and also Docker

QUICKSTART
    Install docker if it's not already installed
    Run Make init in the root of the repository

What's inside
    1. API
        An API with all the routes needed to play to BattleSnakes
        Endpoints:
            GET /game
            POST /game/start
            POST /game/move
            POST /game/move
    2. Database
        A Postgres Database used to store some informations about the games played by our API
        Because first of all before to improve our algorithm we need to see how efficient it is !
