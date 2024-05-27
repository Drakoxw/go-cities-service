# Buscador de ciudades Colombia

Repositorio con el servicio de busqueda de ciudades de colombia, listando el nombre, departamento y codigo dane, la busqueda funciona con el parametro  `query` y permite otros parametros para personalizar y precisar la repuesta: `limit` define la catidad de ciudades, `page` define la paginacion de los datos, `sort` la clave por la que se desea aplicar un order y `order` define si se quiere que se ordene de manera ascendente o descendente.

<b style="color: green">Cada 24 horas actuliza la base de datos. Tambien cuenta con webHook para solicitar la actulizacion de datos</b>

<b style="color: blue">En la rama `main` esta el proyecto con una base de datos MYSQL implementada.</b>

### Herramientas
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/dl/) 1.18 o superior `REQUISITO`

## Configurar y Ejecutar en Local

### 1. Descargar Dependencias
```bash
go mod download
```

### 4. Iniciar App
```bash
go run cmd/main.go
```

## Uso del Servicio

### Valores de los perametros por defecto
* `limit` = 10
* `page` = 1
* `sort` = 'nombre'
* `order` = 'ASC'

## Endpoints:
 
### Buscar Ciudades

* Endpoint: `/cities`
* Método: `GET`
* Descripción: Busca ciudades basándose en parámetros proporcionados o por defecto
* Parámetros:
    * `query` (string): Subcadena a buscar en el nombre de las ciudades.
    * `limit` (number): cantidad de ciudades en la respuesta.
    * `page` (number): Paginación de los datos.
    * `sort` (string): Campo por el cual ordenar los resultados ('id', 'nombre', 'codigodane', 'departamento').
    * `order` (string): Orden de los resultados ('ASC', 'DESC').

> GET "http://localhost:3010/cities?query=cal&sort=nombre&order=ASC"

### Actualizar por demanda `manual`

* Endpoint: `/webhook/update-cities`
* Método: `POST`
* Descripción: Actualiza los datos de las ciudades desde el archivo JSON en línea.

> POST "http://localhost:3010/webhook/update-cities"

## Mantenimiento y Actualización

### Actualizar Datos de Ciudades
El servicio tiene una tarea periódica que actualiza los datos de las ciudades cada 24 horas. También puedes actualizar los datos manualmente utilizando el endpoint `/webhook/update-cities`.

## Generar Binarios

* Ten en cuenta modificar la conexion a la base de datos en `cmd/main.go`

### 1. Para Linux
```bash
GOOS=linux GOARCH=amd64 go build -o bin/citysearch cmd/main.go
```

### 2. Para Windows
```bash
go build -o bin/citysearch.exe cmd/main.go
```
## Ejecutar los Binarios Generados

* Error: `dial tcp: lookup mysql: no such host` -> cambia la conexion en `cmd/main.go`

### Linux
```bash
bin/citysearch
```

### Windows
```bash
bin/citysearch.exe
```

## DOCKER

* Run:
```sh
docker-compose up -d --build
```

### Logs y Monitoreo
Los logs del servidor y la base de datos se pueden revisar utilizando los comandos de Docker:

```sh
docker logs citysearch
```
```sh
docker logs mysql
```

## Licencia
Este proyecto está licenciado bajo los términos de la licencia `MIT`. Consulta el archivo LICENSE para más detalles.




