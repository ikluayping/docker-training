# docker-training

## Instruction
Create image from dockerfile and start container with 2 go-service by <br />
go-service1 will start with port 8000 and publish port to 8888 <br />
go-service2 will start with port 8100 not publish port <br />

go-service1 required environment parameter below for startup
> PORT :the default should be 0.0.0.0:8000 <br />
> SERVICE_2_NAME :the name of the go-service2 <br />

go-service2 required environment parameter below for startup
> PORT :the default should be 0.0.0.0:8100 <br />
>and go-service2 required json file (within data_file folder) for validate user<br />
container should map volume to container path /go/src/go-service2

## Usage
go-service1 request <br />
POST: http://localhost:8888/devops/login
Request:
```json
{
  "username":"admin",
  "password":"password"
}
```
Response:
```json
{
  "status" : "00",
  "description" : "Login Success"
}
```
