# Gokedex - a pokédex written in Go
This is an sample project for learning Go.  
You can interact with a PostgreSQL database via RestAPI.  
This is written in Microservice-Arch, so you could easily add a Frontend e.g. with React.  

## The Stack
**Language:** Go  
**Frameworks:** Gin (RestAPI), pgx (driver), Swagger (docs)     
**Database:** PostgreSQL  
**Containerization**: Docker  
**IDE**: VSCode  

## How to run
Get a postgresql-Container:
```
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -v postgres_data:/var/lib/postgresql/data -d postgres
```