# Snippetbox

## Description
This repository contains the the code that each file in the project directory of building a fast, secure and maintainable web snippetbox application using the fantastic programming language Go. The idea behind this project is using dependency injection to implement efficient functions, taking into consideration best practices and writing efficient unit and integration tests, scallable project structure, session management, authenticating users, securing the server and uses a MySQL database but structured in a way that makes integration with other databases very easy. The project is designed as intended for usage in a production environment and a docker documentation and image link is provided below.


## Features
* Home-Page that shows snippets made by all users
* About Page
* User Registration
* User Authentication/Login
* Logout
* 12hour Session Enabled
* Create a New Snippet for Authenticated Users
* User Profile For Authenticated Users
* Get a Specific Snippet 
* Change Password
* Request and Error Logging to logfiles
* Unit and Integration tests
* Encrypting session tokens
* Protection against CSRF attacks
* Using a self signed TLS certificate

## Docker Image
 <a href="https://hub.docker.com/repository/docker/ifedayoawe/snippetbox/general" target="_blank"> Snippetbox-docker-image </a>
To pull the docker image **docker pull ifedayoawe/snippetbox:1** which will automatically pull the image.
To run the container DBPASS, DBUSER and DBNAME enviromental variables representing a database password, user and name needs to be set on the container, the default application port is **4000** and the database conection **@tcp(snippetbox_db:3306)** here snippetbox_db refers to a running mysql container with same configurations set, an easy way would be to use my docker-compose.yml file to automate the process, of course a .env file containing the enviromental variables would be needed.


```
version: '3.8'

services:
  snippetbox:
    image: snippetbox:1
    container_name: snippetbox
    restart: unless-stopped
    ports: 
      - '4000:4000'
    environment:
      - DBUSER=${DBUSER}
      - DBPASS=${DBPASS}
      - DBNAME=${DBNAME}

  snippetbox_db:
    image: mysql:5.7
    container_name: snippetbox_db
    restart: unless-stopped
    ports: 
      - '3306:3306'
    environment:
      MYSQL_DATABASE: snippetbox
      MYSQL_ROOT_PASSWORD: ${DBPASS}
    volumes: 
      - db-data:/var/lib/mysql

volumes: 
  db-data:
    driver: local
```

## Contribution
Pull requests are and new features suggestions are welcomed.
I also plan on adding more features and API's to this project.