# jwtserver
This is a small "dockerized" repo to provide auth via firebase, along with basic postgres migrations using goose. 


## Start Using
1. spin up postgres container
```
docker compose up -d postgres
```

2. run the migrations
```
make migrate
```
