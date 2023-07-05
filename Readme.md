# Gokedex - a pokédex written in Go
This is an sample project for learning Go.  
You can perform CRUD operations on a PostgreSQL database via RestAPI.  
This is written in Microservice-Arch, so you could easily add a Frontend e.g. with React.  

## The Stack
**Language:** Go  
**Frameworks:** Gin (RestAPI), pgx (driver), Swagger (docs)     
**Database:** PostgreSQL  
**Containerization**: Docker  
**IDE**: VSCode  

## How to run
**Get a postgresql-Container:**  
Basic command for running a postgres container:
```
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -v postgres_data:/var/lib/postgresql/data -d postgres
```

We can use this container to run the app locally with ```go run main.go```.  
If we want to containerize the app and run it, we need to make some adjustments. 

**Build the app container:**
Use the Dockerfile
```
docker build -t gokedocker .
```
***Note how this uses multi-stage build to create a minimal image.***

**Create Network:**  
To make the containers be able to talk to each other, we need to create a dedicated network  
```
docker network create gokedex
```

**Run the containers in the network:**
For the postgres:
```
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -v postgres_data:/var/lib/postgresql/data -d --network=gokedex postgres
```
For the app:
```
docker run -dp 8080:8080 -e DATABASE_URL=postgres://postgres:mysecretpassword@some-postgres:5432/ --network=gokedex gokedocker
```
***Note how we pass the hostname in the DATABASE_URL, using the container name!***

We could aswell use docker-compose to build the stack.  

### Learnings
- lowercase function definitions are considered as unexported, so if creating packages you have to write them Uppercase to use them in other packages.  
- container networking
- Structure of Go apps
- SQL