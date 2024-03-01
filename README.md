Instalar GO
https://go.dev/doc/install

Instalar Dependencias : 

1. JWT para autenticación:
   go get github.com/dgrijalva/jwt-go@v3.2.0

2. Conexión a PostgreSQL con pgx:
   go get github.com/jackc/pgx/v4@v4.11.0

3. Gorilla Mux para el enrutamiento:
   go get github.com/gorilla/mux@v1.8.1

4. Swaggo para documentación de la API (si decides utilizarlo):
   go get github.com/swaggo/swag@v1.16.3

Para instalar todas las dependencias indirectas y asegurarte de que tu proyecto esté actualizado, puedes ejecutar:
 
 go mod tidy

 Para ejecutar tu proyecto en Go

 go run .

 go run cmd/miapp/main.go

Base de datos : 

la base de datos fue creada de RDS (AWS) em PostgreSQL
les dejo las credenciales para que se conecten con el IDE PgAdmin

host : database-1.ct0o62yqab33.us-east-1.rds.amazonaws.com

puerto : 5432

data base : postgres

username : postgres

clave : calasiki

en postman : usar las siguiente url de manera local :

1) http://localhost:8080/login

	 {
       "Username": "pruebas",
       "Password": "})VV}40Nl#0n"
     }

     para logear o validar las url que estan protegidas jwt.


 2    registro de customers : 

     http://localhost:8080/customers


     {
    "customer_id": 1,
    "name": "Juan Pérez",
    "email": "juan.perez@example.com",
    "created_at": "0001-01-01T00:00:00Z"
}


3) registro de  merchants
http://localhost:8080/merchants

{
  "name": "SAM",
  "merchant_code": "777",
  "created_at": "2023-04-01T12:00:00Z"
}

4) para hacer el pago: payments : 

http://localhost:8080/process-payment

{
  "merchant_id": "8",
  "customer_id": "1",
  "amount": 100.50,
  "status": "pending",
  "created_at": "2023-04-12T15:04:05Z"
}

5) para hacer reembolo : refunds

http://localhost:8080/refunds

     {
       "payment_id": 8,
       "amount": 50.00
     }

NOTA : en todas la urls, tomar en cuenta que deben de ingresar el token que se genera, en la pestaña de headers : Authorization.

de todas maneras les adjunto el la colleccion del postman.


de los adicionales, se hizo la proteccion con jwt para la url
de los adicionales la auditoria, lo estube desarrollando por temas famliares tuve que salir, pero lo puedo culminar.

el IDE que se uso para desarrollar fue Visual Studio Code

para correr el proyecto ir al directorio cmd/api 
escribir en el terminal go run .

Tecnologías en la Nube: AWS RDS para PostgreSQL

Especificación:
Para la base de datos de este proyecto, se ha elegido utilizar Amazon RDS (Relational Database Service) con PostgreSQL. Amazon RDS es un servicio de base de datos relacional distribuido por AWS que facilita la configuración, operación y escalado de bases de datos en la nube.

Justificación de la Elección:

1. Gestión Simplificada: Amazon RDS maneja tareas de administración de bases de datos como el aprovisionamiento de hardware, la configuración de la base de datos, el parcheo de software y los respaldos. Esto permite a los desarrolladores centrarse en el diseño y la implementación de la aplicación sin preocuparse por la infraestructura subyacente.

2. Escalabilidad: Con RDS, es fácil escalar la capacidad de almacenamiento y el rendimiento de la base de datos según las necesidades del proyecto. Esto incluye tanto la escalabilidad vertical (aumento del tamaño de la instancia) como la horizontal (aumento del rendimiento a través de réplicas de lectura).

3. Disponibilidad y Durabilidad: RDS ofrece alta disponibilidad y durabilidad con su opción de implementaciones Multi-AZ. En caso de fallos en la infraestructura, RDS promueve automáticamente una réplica de lectura a ser la nueva instancia principal, minimizando así el tiempo de inactividad.

4. Seguridad: La seguridad es una prioridad en RDS, ofreciendo cifrado en reposo y en tránsito, control de acceso basado en roles y aislamiento de red mediante Amazon VPC. Esto asegura que los datos sensibles estén protegidos adecuadamente.

5. Compatibilidad: RDS para PostgreSQL es totalmente compatible con las aplicaciones que utilizan PostgreSQL, lo que significa que no hay necesidad de cambiar el código de la aplicación para migrar a RDS.

6. Costo-Efectividad: Con RDS, solo pagas por los recursos que utilizas. El modelo de precios a demanda permite a las empresas ajustar los costos según el uso real, lo que puede resultar en ahorros significativos en comparación con la gestión de una base de datos en servidores propios o dedicados.

7. Integración con AWS: RDS se integra bien con otros servicios de AWS, como AWS Lambda, Amazon EC2 y Amazon S3, proporcionando una solución completa para aplicaciones en la nube y facilitando la implementación de arquitecturas de aplicaciones modernas.

En resumen, la elección de AWS RDS para PostgreSQL se basa en su capacidad para proporcionar una solución de base de datos robusta, escalable y de alta disponibilidad que puede adaptarse a las necesidades cambiantes del proyecto, al tiempo que reduce la carga de gestión y operación para el equipo de desarrollo.



------------- IN ENGLISH ---------------


## Well done code translation:

**Installing Go**

* [https://go.dev/doc/install](https://go.dev/doc/install)

**Installing Dependencies**

1. JWT for authentication:
```
go get github.com/dgrijalva/jwt-go@v3.2.0
```

2. PostgreSQL connection with pgx:
```
go get github.com/jackc/pgx/v4@v4.11.0
```

3. Gorilla Mux for routing:
```
go get github.com/gorilla/mux@v1.8.1
```

4. Swaggo for API documentation (optional):
```
go get github.com/swaggo/swag@v1.16.3
```

To install all indirect dependencies and ensure your project is up to date, you can run:
```
go mod tidy
```

**Running your Go project**

* `go run .`
* `go run cmd/miapp/main.go`

**Database**

* The database was created using RDS (AWS) in PostgreSQL.
* Credentials for connecting with PgAdmin IDE:
    * Host: `database-1.ct0o62yqab33.us-east-1.rds.amazonaws.com`
    * Port: `5432`
    * Database: `postgres`
    * Username: `postgres`
    * Password: `calasiki`

**Postman**

* Use the following URLs locally:

1. Login: `http://localhost:8080/login`

```json
{
  "Username": "pruebas",
  "Password": "})VV}40Nl#0n"
}
```

* Use the generated token in the Authorization header for protected URLs.

2. Customer registration: `http://localhost:8080/customers`

```json
{
  "customer_id": 1,
  "name": "Juan Pérez",
  "email": "juan.perez@example.com",
  "created_at": "0001-01-01T00:00:00Z"
}
```

3. Merchant registration: `http://localhost:8080/merchants`

```json
{
  "name": "SAM",
  "merchant_code": "777",
  "created_at": "2023-04-01T12:00:00Z"
}
```

4. Payment processing: `http://localhost:8080/process-payment`

```json
{
  "merchant_id": "8",
  "customer_id": "1",
  "amount": 100.50,
  "status": "pending",
  "created_at": "2023-04-12T15:04:05Z"
}
```

5. Refunds: `http://localhost:8080/refunds`

```json
{
  "payment_id": 8,
  "amount": 50.00
}
```

**Notes:**

* Include the token in the Authorization header for all URLs.
* The Postman collection is attached.
* JWT protection is implemented for the URL.
* Audit functionality was being developed but was paused due to family reasons. It can be completed later.
* Visual Studio Code was used as the IDE.
* To run the project, navigate to the `cmd/api` directory and type `go run .` in the terminal.

**Cloud Technologies:** AWS RDS for PostgreSQL

**Specification:**

* Amazon RDS (Relational Database Service) with PostgreSQL was chosen for the project's database.
* AWS RDS is a managed relational database service that simplifies setup, operation, and scaling of databases in the cloud.

**Justification for the choice:**

1. **Simplified Management:** AWS RDS handles database administration tasks such as hardware provisioning, database configuration, software patching, and backups. This allows developers to focus on application design and implementation without worrying about the underlying infrastructure.

2. **Scalability:** With RDS, it is easy to scale the database storage and performance according to the project's needs. This includes both vertical scaling (increasing instance size) and horizontal scaling (increasing performance through read replicas).

3. **Availability and Durability:** RDS offers high availability and durability with its Multi-AZ deployment option. In case of infrastructure failures, RDS automatically promotes a read replica to be the new primary instance, minimizing downtime.

4. **Security:** Security is a priority in RDS, offering encryption at rest and in
