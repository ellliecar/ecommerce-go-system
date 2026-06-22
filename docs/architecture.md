Arquitectura del Sistema
Sistema de Gestión E-commerce en Golang

Estudiante: Elizabeth Cardona
Asignatura: Programación Orientada a Objetos
Lenguaje: Golang
Fecha: Junio 2026

1. Descripción General

El Sistema de Gestión E-commerce es una aplicación desarrollada en Golang que permite administrar productos, procesar pedidos y exponer funcionalidades mediante servicios web REST.

La solución integra los conceptos estudiados durante las cuatro unidades de la asignatura, incluyendo programación orientada a objetos, interfaces, encapsulación, manejo de errores, serialización JSON, concurrencia y pruebas de software.

El objetivo principal es demostrar la aplicación práctica de los fundamentos de desarrollo de software mediante una arquitectura modular, escalable y mantenible.

2. Objetivo del Sistema

Desarrollar un sistema de gestión E-commerce que permita administrar productos y pedidos utilizando servicios web REST, aplicando buenas prácticas de programación en Golang, principios de orientación a objetos, manejo de errores y procesamiento concurrente.

3. Arquitectura General del Sistema
Descripción de la Arquitectura

La arquitectura implementa una separación clara de responsabilidades mediante capas lógicas que facilitan el mantenimiento y la escalabilidad del sistema.

Capa de Presentación (API REST)

Responsable de recibir solicitudes HTTP y devolver respuestas estructuradas.

Funciones:

Recepción de solicitudes GET, POST y PUT.
Validación básica de datos.
Comunicación con la lógica de negocio.
Serialización de respuestas JSON.
Capa de Lógica de Negocio

Implementa las reglas principales del sistema.

Responsabilidades:

Procesamiento de pedidos.
Validación de inventario.
Gestión de productos.
Coordinación de operaciones concurrentes.
Capa de Modelos e Interfaces

Representa las entidades principales y los contratos del sistema.

Componentes:

Product
IProduct

Beneficios:

Encapsulación.
Abstracción.
Bajo acoplamiento.
Reutilización.
Capa de Manejo de Errores

Controla situaciones excepcionales durante la ejecución.

Ejemplos:

Producto inexistente.
Stock insuficiente.
Solicitudes inválidas.
Errores de procesamiento.
Capa de Concurrencia

Implementada utilizando herramientas nativas de Golang:

Goroutines
sync.WaitGroup
sync.RWMutex

Permite procesar múltiples solicitudes simultáneamente de forma segura.

Serialización JSON

Toda la comunicación entre cliente y servidor utiliza formato JSON para garantizar interoperabilidad y compatibilidad con aplicaciones modernas.

4. Estructura del Proyecto
ecommerce-go-system/
├── docs/
│   ├── architecture.md
│   └── cronograma.md
│
├── internal/
│   ├── api/
│   │   └── handlers.go
│   │
│   ├── errors/
│   │   └── domain.go
│   │
│   ├── interfaces/
│   │   └── contracts.go
│   │
│   ├── models/
│   │   └── product.go
│   │
│   └── services/
│       └── order.go
│
├── tests/
│   ├── api_test.go
│   ├── product_test.go
│   └── service_test.go
│
├── .gitignore
├── LICENSE
├── README.md
├── go.mod
└── main.go
Descripción de la estructura
docs/: documentación técnica y planificación del proyecto.
docs/architecture.md: documento de arquitectura del sistema.
docs/cronograma.md: planificación y cronograma de desarrollo.
internal/api/: controladores y endpoints de los servicios web REST.
internal/api/handlers.go: implementación de rutas y manejo de solicitudes HTTP.
internal/errors/: definición de errores personalizados del dominio.
internal/errors/domain.go: catálogo de errores de negocio utilizados por el sistema.
internal/interfaces/: contratos e interfaces utilizadas por la aplicación.
internal/interfaces/contracts.go: definición de interfaces para desacoplar componentes.
internal/models/: entidades y modelos de negocio.
internal/models/product.go: estructura y comportamiento de los productos.
internal/services/: implementación de la lógica de negocio.
internal/services/order.go: procesamiento de pedidos, validaciones y control de inventario.
tests/: pruebas unitarias e integración del sistema.
tests/api_test.go: pruebas de los servicios web.
tests/product_test.go: pruebas de la entidad Product.
tests/service_test.go: pruebas de la lógica de negocio.
main.go: punto de entrada principal de la aplicación.
go.mod: gestión de dependencias y configuración del módulo Go.
README.md: documentación principal del proyecto.
LICENSE: licencia de distribución del software.
.gitignore: configuración de archivos excluidos del control de versiones.
5. Componentes Principales
Product

Representa un producto disponible para la venta.

Responsabilidades:

Almacenar información del producto.
Gestionar el inventario.
Verificar disponibilidad.
Aplicar encapsulación.

Atributos:

id
name
price
stock

Métodos:

GetID()
GetName()
GetPrice()
GetStock()
IsAvailable()
ReduceStock()
IProduct

Define el contrato que deben cumplir las entidades de producto.

Métodos:

GetID()
GetName()
GetPrice()
GetStock()
ReduceStock()
IsAvailable()

Beneficios:

Desacoplamiento.
Reutilización.
Extensibilidad.
OrderService

Responsabilidades:

Validación de productos.
Procesamiento de pedidos.
Control de stock.
Gestión de concurrencia.

Métodos:

AddProduct()
ProcessOrder()
GetQueueStatus()
API REST

Funciones:

Recepción de solicitudes.
Procesamiento de datos JSON.
Respuestas HTTP.
Comunicación con servicios internos.
6. Servicios Web Implementados
Endpoint	Método	Descripción
/health	GET	Estado general del sistema
/api/products	GET	Consulta de productos
/api/products/{id}	GET	Consulta individual
/api/orders	POST	Registro de pedidos
/api/payments	POST	Procesamiento de pagos
/api/users/history	GET	Historial de compras
/api/inventory	PUT	Actualización de inventario
/api/analytics	GET	Estadísticas del sistema
/api/concurrent	POST	Demostración de concurrencia
7. Programación Orientada a Objetos

Encapsulación: protege los datos internos mediante métodos.
Abstracción: uso de interfaces.
Modularidad: separación por capas.
Reutilización: contratos reutilizables.

8. Manejo de Errores

Casos:

Producto inexistente.
Stock insuficiente.
Solicitudes inválidas.
Errores de procesamiento.
9. Concurrencia

Tecnologías:

Goroutines
sync.WaitGroup
sync.RWMutex

Beneficios:

Paralelismo.
Mayor rendimiento.
Seguridad en recursos compartidos.
10. Pruebas de Software

Pruebas realizadas:

Productos.
Pedidos.
Errores.
Servicios.
Concurrencia.
11. Integración de Unidades

Unidad 1: análisis.
Unidad 2: modelos e interfaces.
Unidad 3: API REST.
Unidad 4: concurrencia y pruebas.

12. Conclusión

El sistema integra programación orientada a objetos, APIs REST, JSON, manejo de errores, pruebas y concurrencia, logrando una arquitectura modular, escalable y mantenible en Golang.
