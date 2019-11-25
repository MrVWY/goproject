This file is create when I just started learning Micro.  
You must have these foundations before you can start watching :  
* what is Micro API 
- How to wirte and compile the files with .proto
* The structure of Micro 
- How to use after 
  
How compile the files with .proto
====
* use the command "proto --proto_path=. --go_out=. --micro_out=. Your's proto file path
```  
proto --proto_path=. --go_out=. --micro_out=. ./proto/(Your proto file name)
```  
  
How to run
====  
Vsersion 1  
* run micro API in your ternimal  
```
#micro api --handler=api
```
- run ./service/main.go  
```
go run main.go
```
* run ./api/main.go
```
go run main.go
```
  
Below is the rendering:  
* ./api/main.go  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573873017(1).png)  
  
- ./service/main.go  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/service.png)  
  
* Micro API  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573872840(1).png)  
  
___
So,you can see run when ./api/main.go and ./service/main.go running, they `run on different ports`.  
If you want to konw clearly which ports they are running on. you can run this command "micro web" in you ternimal and entry http://localhost:8082/ to see your services.  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573875469(1).png)  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573876096(1).png)  
  
How to use
====
In here, I use Postman, you also use the curl
* postman  
  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573873197(1).jpg)  
  
So this is the base usage!
  
Version 2
====
Version 2 is only on version one, adding a broker message subscription  
  
* /message/message.go  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573987194(1).png)  
  
- V2.go  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/1573986249(1).png)
  
* run micro API in your ternimal  
```
#micro api --handler=api
```
- run ./service/V2.go  
```
go run V2.go
```
* run /service2/main.go
```
go run main.go
```  
- run ./api/main.go
```
go run main.go
```
  
The structure:  
![alt](https://github.com/MrVWY/goproject/blob/master/microservives/emamples/service/img/A.png)
  
Now shop! shop! Do you konw this how to run and work? How the Micro can `find the service` which running on `different ports`?Please see the web that you just run and see the [emamples/readme.md](https://github.com/MrVWY/go/tree/master/shippy/emamples)! I will update soon!
  
