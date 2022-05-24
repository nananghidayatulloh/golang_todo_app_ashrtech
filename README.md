# golang_todo_app_ashrtech

todo app with golang and postgres docker

## Requirements

- Docker v20.10.14
- Go 1.18 (docker image)
- Postgres latest (docker image)
- Postman v9.19.0

### Run

- Go inside the folder
- `docker-compose up`
- Ensure gotodoashrtech and posgres docker container is running
- Open your favorite REST API Client such as [POSTMAN](https://www.postman.com/downloads/)
- Import POSTMAN Collection: asht_tech_api_collection into your POSTMAN App.

#### POSTMAN USING

- Register
- Login and copy the token return
- Choose Authorization tab and choose the Type with Bearer Token and paste the token from login return to Token.
