# Arquitectura del Sistema - Sistema de Gestión de E-commerce (Golang)

**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Institución:** Universidad Internacional del Ecuador (UIDE)  
**Fecha:** 28 de junio de 2026

---

## Diagrama de Arquitectura General

```mermaid
flowchart LR
    A[Cliente HTTP] --> B[Servidor Go :8080]
    B --> C[api.Handlers]
    C --> D[services.OrderService]
    D --> E[models.Product]
    E --> F[interfaces.IProduct]
    
    subgraph Capas
    C
    D
    E
    F
    end
    
    style A fill:#e1f5fe
    style B fill:#fff9c4
    style E fill:#e8f5e9
