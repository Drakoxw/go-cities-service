# Etapa 1: Construcción del binario
FROM golang:1.18 AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Descargar las dependencias
RUN go mod download

# Construir el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o citysearch cmd/main.go

# Etapa 2: Construcción de la imagen final
FROM alpine:latest

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/citysearch .

# Copiar el archivo de datos
COPY data/cities.json ./data/cities.json



# Exponer el puerto 3010
EXPOSE 3010

# Ejecutar el binario
CMD ["./citysearch"]
