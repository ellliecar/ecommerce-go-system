# 🏗️ Arquitectura del Sistema & Justificación Técnica
**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Institución:** Universidad Internacional del Ecuador (UIDE)  
**Proyecto:** Sistema de Gestión de E-commerce (Golang)  
**Fecha:** 28 de junio de 2026

---

## 📐 Diagrama de Arquitectura General

```mermaid
graph TD
    A[Cliente HTTP] --> B[net/http Server]
    B --> C[api.Handlers]
    C --> D[services.OrderService]
    
    D --> E{Validación Concurrente}
    E -->|Lectura| F[sync.RWMutex]
    E -->|Escritura| G[sync.Mutex]
    
    F --> H[map[string]IProduct]
    F --> I[map[string]bool sessions]
    
    D --> J[models.Product]
    J --> K[Encapsulamiento: campos minúsculos]
    J --> L[Invariantes: constructor NewProduct]
    
    D --> M[Manejo de Errores]
    M --> N[errors.ErrInsufficientStock]
    M --> O[errors.ErrDuplicateRequest]
    
    C --> P[encoding/json Serialización]
    P --> Q[Respuesta HTTP 200/400/500]
    
    style A fill:#e1f5fe
    style D fill:#fff9c4
    style J fill:#e8f5e9
    style M fill:#ffebee
    style P fill:#f3e5f5
