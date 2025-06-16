# meli-product-api
Backend en Go para productos estilo Mercado Libre, con datos en JSON y arquitectura limpia.
# 🛒 Meli Product API

Este es un servicio backend desarrollado en **Go** que expone una API REST para consultar productos. Forma parte del reto técnico inspirado en la plataforma de Mercado Libre.

## 🚀 Funcionalidades

- Listado completo de productos: `GET /products`
- Consulta de detalle por ID: `GET /products/{id}`
- Persistencia local en archivo `.json`
- Arquitectura **hexagonal** (ports & adapters)
- Cobertura de tests por capa (`application`, `repository`)

## 🧱 Estructura del proyecto

```bash
meli-product-api/
│
├── cmd/                    # Entry point (main.go)
├── internal/
│   ├── application/        # Use cases y lógica de negocio
│   ├── config/             # Configuración (por ahora mínima)
│   ├── domain/             # Modelos y contratos (ports)
│   └── infrastructure/
│       ├── handler/        # HTTP handlers (adaptadores externos)
│       └── repository/     # Adaptadores de persistencia (archivo JSON)
├── test/                   # Pruebas externas (integration-like)
├── products.json           # "Base de datos" local
├── go.mod
├── go.sum
└── README.md
