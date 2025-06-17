# meli-product-api
Backend en Go para productos estilo Mercado Libre, con datos en JSON y arquitectura limpia.
# 🛒 Meli Product API

Este es un servicio backend desarrollado en **Go** que expone una API REST para consultar productos. Forma parte del reto técnico inspirado en la plataforma de Mercado Libre.

## 🚀 Funcionalidades

- Consulta de detalle por ID: `GET /products/{id}` TODO falta generar version a la API
- Persistencia local en archivo `.json`
- Arquitectura **hexagonal** (ports & adapters)
- Cobertura de tests por capa (`application`, `repository`)

## 🧱 Estructura del proyecto

```bash

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

meli-product-api/
├── main.go                # Punto de entrada
├── internal/              # Código privado de la aplicación
│   ├── application/       # Casos de uso y lógica de negocio
│   ├── domain/           # Entidades y puertos
│   └── config/           # Configuración y logger TODO solo debe existir un punto de cableado
├── src/                   # Código fuente
│   └── infrastructure/    # Implementaciones concretas
│       ├── handler/      # Handlers HTTP
│       ├── repository/   # Repositorios TODO los repositorios no son de la misma capa que los handlers
│       └── config/       # Configuración de infraestructura TODO no me suena que este dentro de infra
├── data/                 # Datos y recursos estáticos
│   └── products.json
├── tests/                # Tests
├── .circleci/            # Configuración de CI/CD
├── go.mod
├── go.sum
└── README.md
