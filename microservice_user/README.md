# Microservice_user  
Microservice User of bookstore project  
  
## Description

This microservice has the main responsabilities:  
  
* Create a user
* Update a user
* Delete a user 
* Get user by ID
* Get users by status (active, inactive)

This microservice also use oauth lib to authenticate a user: github.com/diegoclair/bookstore_oauth-go/oauth</b></b> 

## Start application
### Permissions first:  
For <b>Unix</b> enviroment run the comand:  
* <b>```chmod +x .docker/entrypoint.sh```</b>  

For <b>Windows</b> enviroment run the comand:   
* <b>```dos2unix +x .docker/entrypoint.sh```</b>  

Now you can run:  <br>
* <b>```docker-compose up```</b>
<br><br>
