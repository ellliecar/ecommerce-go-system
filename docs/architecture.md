# Arquitectura del Sistema

## Sistema de Gestión E-commerce en Golang

**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Lenguaje:** Golang  
**Fecha:** Junio 2026  

---

# 1. Descripción General

El proyecto consiste en un Sistema de Gestión E-commerce desarrollado en Golang. Su objetivo es administrar productos, pedidos y operaciones comerciales mediante servicios web REST utilizando serialización JSON.

La arquitectura fue diseñada siguiendo principios de programación orientada a objetos adaptados a Go mediante el uso de structs, interfaces, encapsulación, manejo de errores y concurrencia.

---

# 2. Estructura del Proyecto

```text
ecommerce-go-system/
├── main.go
├── go.mod
├── README.md
├── internal/
│   ├── models/
│   │   └── product.go
│   ├── interfaces/
│   │   └── contracts.go
│   ├── errors/
│   │   └── domain.go
│   ├── services/
│   │   └── order.go
│   └── api/
│       └── handlers.go
├── tests/
│   ├── product_test.go
│   ├── service_test.go
│   └── api_test.go
└── docs/
    ├── architecture.md
    └── cronograma.md

# 3. Diagrama de Arquitectura General
flowchart TD
    Client[🖥️ Cliente HTTP] -->|Solicitud POST/GET + JSON| API[📡 internal/api/ handlers.go]
    API -->|Decodifica & Valida| SVC[⚙️ internal/services/ order.go]
    SVC -->|Invoca contrato| IF[🔌 internal/interfaces/ contracts.go]
    IF -->|Implementa entidad| MOD[📦 internal/models/ product.go]
    SVC -->|Sincroniza acceso| MUT[🔒 sync.RWMutex]
    SVC -->|Captura fallos| ERR[⚠️ internal/errors/ domain.go]
    API -->|Serializa respuesta| JSON[📋 encoding/json]
    JSON -->|Retorna HTTP 200/400/404| Client

    subgraph 🔐 Capa de Lógica y Datos
    MOD
    IF
    SVC
    MUT
    ERR
    end

    style Client fill:#e1f5fe,stroke:#01579b
    style API fill:#fff9c4,stroke:#fbc02d
    style SVC fill:#e8f5e9,stroke:#2e7d32
    style JSON fill:#e1f5fe,stroke:#01579b
Explicación del flujo:
El cliente envía una solicitud HTTP con payload JSON.
El handler (api/handlers.go) decodifica la petición y la dirige al servicio.
OrderService orquesta la lógica de negocio, validando reglas y gestionando el estado.
Se aplica el contrato IProduct sobre la entidad Product, garantizando encapsulamiento.
sync.RWMutex protege el acceso concurrente a recursos compartidos sin bloquear lecturas innecesarias.
Los errores de dominio se capturan y retornan de forma tipada.
La respuesta se serializa a JSON y se devuelve al cliente con el código HTTP semántico correspondiente.
4. Componentes Principales
Product
Representa un producto disponible dentro del sistema.
Responsabilidades:
Almacenar información del producto.
Controlar el stock.
Validar disponibilidad.
Aplicar encapsulación mediante atributos privados.
Propiedades:
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
Define el contrato que deben cumplir los productos.
Funciones:
GetID()
GetName()
GetPrice()
GetStock()
ReduceStock()
IsAvailable()
Beneficios:
Abstracción.
Bajo acoplamiento.
Facilidad para futuras extensiones.
OrderService
Gestiona el procesamiento de pedidos.
Responsabilidades:
Validar productos.
Verificar stock.
Registrar pedidos.
Mantener cola de pedidos.
Gestionar concurrencia.
Funciones principales:
AddProduct()
ProcessOrder()
GetQueueStatus()
API REST
Expone las funcionalidades del sistema mediante servicios web.
Funciones:
Recepción de solicitudes HTTP.
Procesamiento de datos JSON.
Respuesta estructurada al cliente.
5. Servicios Web Implementados
Endpoint
Método
Descripción
/health
GET
Estado general del sistema
/api/products
GET
Consulta de productos
/api/products/{id}
GET
Consulta individual
/api/orders
POST
Registro de pedidos
/api/payments
POST
Procesamiento de pagos
/api/users/history
GET
Historial de compras
/api/inventory
PUT
Actualización de inventario
/api/analytics
GET
Estadísticas del sistema
/api/concurrent
POST
Demostración de concurrencia
6. Aplicación de Programación Orientada a Objetos
Encapsulación
Los atributos internos de Product permanecen privados y sólo pueden consultarse mediante métodos públicos.
Abstracción
Las interfaces permiten definir comportamientos sin exponer detalles de implementación.
Modularidad
Cada componente tiene responsabilidades claramente definidas.
Reutilización
Los contratos mediante interfaces facilitan la reutilización del código.
7. Manejo de Errores
El sistema implementa errores personalizados para controlar situaciones como:
Producto inexistente.
Stock insuficiente.
Solicitudes duplicadas.
Errores de procesamiento.
Esto mejora la robustez y mantenibilidad del sistema.
8. Concurrencia
La concurrencia se implementa utilizando:
Goroutines
sync.WaitGroup
sync.RWMutex
Beneficios:
Procesamiento simultáneo de solicitudes.
Protección de recursos compartidos.
Mayor escalabilidad.
9. Pruebas de Software
Se desarrollaron pruebas para:
Validación de productos.
Procesamiento de pedidos.
Verificación de errores.
Integración de servicios.
Objetivos:
Garantizar estabilidad.
Detectar errores tempranamente.
Validar el correcto funcionamiento del sistema.
10. Integración de las Cuatro Unidades
Unidad 1
Análisis y diseño del sistema.
Unidad 2
Implementación de estructuras, interfaces y encapsulación.
Unidad 3
Servicios web REST y serialización JSON.
Unidad 4
Concurrencia, pruebas de software e integración final.
11. Conclusión
El Sistema de Gestión E-commerce desarrollado en Golang integra los conceptos fundamentales estudiados durante las cuatro unidades de la asignatura. La solución implementa programación orientada a objetos, servicios web, serialización JSON, manejo de errores y concurrencia, permitiendo construir una aplicación escalable, modular y mantenible.
