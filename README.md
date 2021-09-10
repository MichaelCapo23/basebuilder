# jwtserver
This is a small "dockerized" repo to provide auth via JWT and JWKS, along with basic postgres migrations using goose. 


## Start Using
1. spin up postgres container
```
docker compose up -d postgres
```

2. run the migrations
```
make migrate

or 

cd migration
go run . up
```