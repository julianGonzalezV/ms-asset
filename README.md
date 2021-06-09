# ms-asset

Microservice for  data managing related to assets

# Instalation
In your local host you have to create the next folders structrure:
xxworkspace
    bin
    src
    pkg

>Then, create the GOPATH environment variable, point to xxworkspace folder--ok

>Additionally, you  have to create  GOPATH/bin into your PATH env variable (This is a pending step, please read https://golang.org/doc/gopath_code.html before )
    

# Commands to execute if you want to run this project
-To install gorilla mux
$ go get -u github.com/gorilla/mux

# Install the MongoDB Go Driver
https://blog.friendsofgo.tech/posts/driver-oficial-mongodb-golang/
go get -u go.mongodb.org/mongo-driver

## Linux
```bash
$ source .env && go run main.go
```

## Windows
```bash
$ start.bat
```

# Get zip to AWS lambda: 
Doc
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

Process
- Into your root project folder execute:
1) $ go get github.com/aws/aws-lambda-go/lambda   (if required)
2) $ GOOS=linux go build main.go --> this command create an executable file called main as the .go name file
3) $ zip ms-asset.zip main
4) $ upload zip to S3 via aws cli or manually
- Nota: Para proyectos desde cero es mejor hacerlo con la lib de aws 
pero en existentes se pueden usar wrappers como: (para este proyecto se usó apexGateway, ya que contabamos con una arquitectura ya hecha)
https://github.com/apex/gateway (https://www.ocelotconsulting.com/2019/02/25/the-right-abstraction-for-lambdas.html)

## AWS configuration:
- See "aws_conf_folder" where all process is documented via images


## Consumos:
Crear POST http://localhost:8000/assets
{
	"code":"AP0001", 
	"RegistrationNumber":"INE0001224234", 
	"country":"COLOMBIA", 
	"province":"VALLE DEL CAUCA", 
	"city":"TULUÁ", 
	"description":"Apartemento cerca al metro, con parqueaderos comunes pero siempre en la unidad te aseguran un espacio (no el mismo siempre) para los vehiculos de propietarios o inquilins. Cerca a hospital XYZ , a centro comercal ABC , varias tiendas de barrio alrededor , dos parques cercano el de la familia y el de las chimeneas , super central ubicado sobre todo el centro de la moda ", 
	"category":"", 
	"state":"ACTIVO", 
	"rentingPrice":1200000, 
	"area":63, 
	"rooms":3, 
	"bathRooms":1, 
	"parkings":1, 
	"furnished":false, 
	"type":"APARTAMENTO", 
	"images":[{"url":"url","description":"sala comedor muy agradable con balcon","type":"PRINCIPAL","state":"ACTIVO"}]
}