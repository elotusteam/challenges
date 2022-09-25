You need to have postgres server with info:
 - localhost: 5432
 - user: postgres
 - password: secret

And already have database name "hackathonDB" for project to connection and migration.



I built 3 api for user to interact
 -localhost:8080/login (you have to login with google account)

 -localhost:8080/register 
	body:	{
    		"user_name":"hackathon",
    		"password":"pw"
		}

 -localhost:8080/upload
	- Authorization (must have)
	- form-data: key: file