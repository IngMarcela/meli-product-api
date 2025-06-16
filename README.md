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
├── main.go                # Punto de entrada
├── internal/              # Código privado de la aplicación
│   ├── application/       # Casos de uso y lógica de negocio
│   ├── domain/           # Entidades y puertos
│   └── config/           # Configuración y logger
├── src/                   # Código fuente
│   └── infrastructure/    # Implementaciones concretas
│       ├── handler/      # Handlers HTTP
│       ├── repository/   # Repositorios
│       └── config/       # Configuración de infraestructura
├── data/                 # Datos y recursos estáticos
│   └── products.json
├── tests/                # Tests
├── .circleci/            # Configuración de CI/CD
├── go.mod
├── go.sum
└── README.md
