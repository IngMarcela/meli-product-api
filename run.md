# 🛠️ meli-product-api

Backend del detalle de producto inspirado en Mercado Libre, construido con Go utilizando arquitectura hexagonal.

---

## 🚀 Requisitos

- Go 1.20+
- Git
- (Opcional) `make` para tareas automatizadas

---

## 📦 Instalación

1. Clona el repositorio:

```bash
git clone https://github.com/IngMarcela/meli-product-api.git
cd meli-product-api
```

2. Descarga las dependencias:

```bash
go mod tidy
```

---

## ▶️ Ejecución

Ejecutá el proyecto con:

```bash
go run cmd/main.go
```

La API quedará corriendo en:

```
http://localhost:3000
```

---

## 📚 Endpoints disponibles

### GET /products
Retorna todos los productos.

### GET /products/:id
Retorna un producto por su ID.

### GET /products/search?q=texto
Busca productos cuyo título o descripción contengan `texto`.

---

## 🧪 Tests

Para ejecutar todos los tests y ver la cobertura:

```bash
go test -cover ./...
```

Para ver el reporte de cobertura en navegador:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 📂 Estructura del proyecto

```
meli-product-api/
├── cmd/                     # Punto de entrada (main.go)
├── internal/
│   ├── application/         # Casos de uso y queries
│   ├── config/              # Configuración y logger
│   ├── domain/              # Entidades y puertos
│   └── infrastructure/      # Handlers HTTP y repositorios
├── data/products.json       # Fuente de datos local
├── run.md                   # Instrucciones de ejecución 📄
└── go.mod                   # Dependencias
```

---

## 📌 Notas

- Los productos se leen desde `data/products.json`
- No se persisten cambios (solo lectura)
- El servicio incluye logs básicos para trazabilidad

---

¡Listo para ejecutar y testear! 🙌
