# API Hello World con Go Gin y Swagger

Proyecto simple que muestra cómo integrar Swagger/OpenAPI en una API Go con Gin Framework.

## Características

- Framework: Gin Web Framework
- Documentación: Swagger/OpenAPI 2.0
- Endpoint simple: Hello World

## Estructura del Proyecto

```
google-ai-service/
├── docs/              # Documentación Swagger (generada automáticamente)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── main.go            # Código principal con endpoint y anotaciones Swagger
├── go.mod             # Dependencias
└── README.md
```

## Requisitos

- Go 1.21 o superior

## Instalación

1. Instalar dependencias:
```bash
go mod download
```

2. Instalar la herramienta `swag` CLI (necesaria para generar la documentación):
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

## Cómo Generar la Documentación Swagger

La documentación Swagger se genera a partir de comentarios especiales en el código.

### Paso 1: Agregar anotaciones en el código

En el archivo [main.go](main.go) encontrarás dos tipos de comentarios:

**Comentarios generales (al inicio del archivo):**
```go
// @title           API Hello World
// @version         1.0
// @description     API simple de ejemplo con Gin y Swagger
// @host            localhost:8080
// @BasePath        /api/v1
```

**Comentarios del endpoint:**
```go
// HelloWorld godoc
// @Summary      Hello World
// @Description  Retorna un mensaje de saludo
// @Tags         hello
// @Accept       json
// @Produce      json
// @Success      200  {object}  HelloResponse
// @Router       /hello [get]
func HelloWorld(c *gin.Context) {
    // ...
}
```

### Paso 2: Generar los archivos de documentación

Ejecuta el comando `swag init` en la raíz del proyecto:

```bash
swag init
```

O si instalaste swag en `~/go/bin`:
```bash
~/go/bin/swag init
```

Esto generará automáticamente el directorio `docs/` con los archivos:
- `docs.go` - Código Go con la documentación
- `swagger.json` - Especificación OpenAPI en JSON
- `swagger.yaml` - Especificación OpenAPI en YAML

### Paso 3: Ejecutar el servidor

```bash
go run main.go
```

### Paso 4: Ver la documentación

Abre tu navegador en:
```
http://localhost:8080/swagger/index.html
```

Verás la interfaz interactiva de Swagger donde puedes probar el endpoint.

## Probar el Endpoint

### Usando curl:
```bash
curl http://localhost:8080/api/v1/hello
```

Respuesta:
```json
{
  "message": "Hello World!"
}
```

### Usando Swagger UI:
1. Ve a http://localhost:8080/swagger/index.html
2. Haz clic en el endpoint `GET /api/v1/hello`
3. Haz clic en "Try it out"
4. Haz clic en "Execute"

## Formato de Anotaciones Swagger

### Anotaciones Generales (nivel API)
```go
// @title           Título de la API
// @version         Versión
// @description     Descripción de la API
// @host            hostname:puerto
// @BasePath        /ruta/base
```

### Anotaciones de Endpoint
```go
// NombreFuncion godoc
// @Summary      Resumen breve del endpoint
// @Description  Descripción detallada
// @Tags         nombre-tag
// @Accept       json
// @Produce      json
// @Param        nombreParam  tipo  tipoData  requerido  "descripción"
// @Success      código  {tipo}  NombreTipo
// @Failure      código  {tipo}  NombreTipo
// @Router       /ruta [método]
```

**Tipos de parámetros:**
- `path` - Parámetro en la URL (ej: `/users/:id`)
- `query` - Parámetro query string (ej: `?name=value`)
- `body` - Parámetro en el cuerpo de la petición

**Métodos HTTP:**
- `[get]`, `[post]`, `[put]`, `[delete]`, `[patch]`

## Ejemplo de Nuevo Endpoint

Para agregar un nuevo endpoint con documentación:

```go
type ByeResponse struct {
    Message string `json:"message" example:"Goodbye!"`
}

// SayBye godoc
// @Summary      Despedida
// @Description  Retorna un mensaje de despedida
// @Tags         hello
// @Produce      json
// @Success      200  {object}  ByeResponse
// @Router       /bye [get]
func SayBye(c *gin.Context) {
    c.JSON(http.StatusOK, ByeResponse{
        Message: "Goodbye!",
    })
}
```

Luego en `main()`:
```go
v1.GET("/bye", SayBye)
```

Y regenera la documentación:
```bash
swag init
```

## Dependencias

- **gin-gonic/gin**: Framework web HTTP
- **swaggo/swag**: Generador de documentación Swagger
- **swaggo/gin-swagger**: Middleware Swagger para Gin
- **swaggo/files**: Archivos estáticos para Swagger UI
