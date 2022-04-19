# Spear-Backend
### Notes of using
- Import postman-collection from [here](https://github.com/spear-app/spear-go/blob/authen-feature/postman_collection/spear-backend.postman_collection.json).
- In update endpoint, gender have to take only one of these values [MALE, FEMALE].
-  Database is seeded, you can look at the data from [here](https://github.com/spear-app/spear-go/blob/authen-feature/pkg/driver/seed.go)
-  After using signup or login, copy the token and paste it in authorization section in postman. Choose bearer token. This is for using any profile and notification endpoints.
-  Notification date of creation has format like this '2022-04-19 20:01:09.385531+02'
### How to run?
- install [docker](https://docs.docker.com/engine/install/) and [docker compose](https://docs.docker.com/compose/install/)
```
docker-compose up
```
To stop the server, ctrl+c and run this command
```
docker-compose down
 ```
 
 ### What's New
Application has sound alarm feature. Mobile sensor can detect sounds around the user. These sounds need to be stored as notifications. We store them in database for 24 hours. After that application uses cron jobs to delete notifications. 
