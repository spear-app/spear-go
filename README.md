# Spear-Backend
### Notes of using
- Import postman-collection from [here](https://github.com/spear-app/spear-go/blob/authen-feature/postman_collection/spear-backend.postman_collection.json).
- In update endpoint, gender have to take only one of these values [MALE, FEMALE].
-  Database is seeded, you can look at the data from [here](https://github.com/spear-app/spear-go/blob/authen-feature/pkg/driver/seed.go)
-  After using signup or login, copy the token and paste it in authorization section in postman. Choose bearer token. This is for using any profile and notification endpoints.
### How to run?
- install [docker](https://docs.docker.com/engine/install/) and [docker compose](https://docs.docker.com/compose/install/)
```
docker-compose up
```
To stop the server, ctrl+c and run this command
```
docker-compose down
 ```
