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
go run main.go
```

La API quedará corriendo en:

```
http://localhost:3000
```

---

## 📚 Endpoints disponibles

### GET /products/:id
Retorna un producto por su ID.

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
├── main.go
├── internal/
│   ├── application/
│   │   └── queries/
│   ├── domain/
│   │   ├── model/
│   │   └── ports/
│   └── config/
│       └── logger.go
├── infrastructure/
│   ├── handler/
│   ├── repository/
│   └── config/
│       ├── routes.go
│       └── wiring.go
├── data/
│   └── products.json
├── tests/
├── .circleci/
├── go.mod
├── go.sum
└── README.md
```

---

## 📌 Notas

- Los productos se leen desde `data/products.json`
- No se persisten cambios (solo lectura)
- El servicio incluye logs básicos para trazabilidad

---

¡Listo para ejecutar y testear! 🙌
