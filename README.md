# go-contacts
## How to run

### ENV
Don't forget to include the following in your ```.env``` file:
```
POSTGRES_DB = gocontacts
POSTGRES_PASSWORD = secret
POSTGRES_USER = postgres
db_type = postgres
db_host = localhost
db_port = 5434
token_password = secret
```

### Setup DB
```
docker run -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=gocontacts -p 5432:5432 postgres
```

###### Reference https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
