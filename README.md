#  bruh-boys-net



This is a simple blog that I have been making just for fun.
Im trying to save prices in the server so the server will only render the posts that i publish.
The client will work as a database.
You can use it if you want but i dont really care.


## how to use it



For using it you dont need too much, first you need to change the password for security reasons if you dont change it i dont care but whatever.
Then you need to change a variable called `host` located in `client/src/vars.go` to the hostname of your server.

Then you only need to run the server and the client, you need to keep your computer on so keep that in mind.
Also, if you want to change the port you only need to create  a new variable in your .env file fo
```sh
#mira leo tienes que ver esto
go run client/main.go <contraseña># la contraseña por default es 123
go run server/main.go 
# tienes que transpilar el codigo ts eso si lo sabes hacer y ya
``` 

If you want to edit a post you can do it but for seeing the changes you need to restart the service or clear the cache.

Everyone with the password can post whatever he want but he will need to mantain his computer on for hosting the posts





![img](https://media.discordapp.net/attachments/907631182240436305/1078119013013524590/imagen.png?width=1920&height=400)
