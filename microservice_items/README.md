# Microservice_items  
Microservice Items of bookstore project  

## Description

This microservice has the main responsabilities:  
  
<!-- * Create a user
* Update a user
* Delete a user 
* Get user by ID
* Get users by status (active, inactive)
-->

This microservice also use [oauth lib](https://github.com/diegoclair/go_oauth-lib) to authenticate the request.</b></b> 

## Start application
### Requirements:
* To run this project you need to have the <b>docker</b> installed in your computer.  
  - To install docker, [click here](https://docs.docker.com/get-docker/)
* To run this microservice, you need first to run the <b>[microservice_oauth](https://github.com/diegoclair/bookstore_microservices/tree/master/microservice_oauth)</b>

### Permissions first:  

* For <b>Unix</b> enviroment, run the comand:  
<b>```chmod +x .docker/entrypoint.sh```</b>  

* For <b>Windows</b> enviroment, run the comand:   
<b>```dos2unix +x .docker/entrypoint.sh```</b>  
  
### Creating a shared network:
* Create a network for the comunication between the microservices:  
<b>```docker network create shared_net``` </b>  

### Start:
* Now you can run:  <br>
<b>```docker-compose up```</b>
<br><br>
