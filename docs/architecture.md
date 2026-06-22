# Arquitectura del Sistema - Sistema de Gestión de E-commerce (Golang)

**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Institución:** Universidad Internacional del Ecuador (UIDE)  
**Fecha:** 28 de junio de 2026

---

## Diagrama de Arquitectura General

```mermaid
flowchart TB
    subgraph Cliente["🖥️ Cliente HTTP"]
        C1[Navegador / Postman / curl]
    end

    subgraph Servidor["🚀 Servidor Go :8080"]
        direction TB
        
        subgraph API["📡 internal/api/ - Handlers"]
            H1[HealthCheck]
            H2[ListProducts]
            H3[GetProduct]
            H4[CreateOrder]
            H5[ProcessPayment]
            H6[GetUserHistory]
            H7[UpdateInventory]
            H8[GetAnalytics]
            H9[HandleConcurrency]
        end

        subgraph Services["⚙️ internal/services/ - OrderService"]
            OS[OrderService struct]
            OS1[catalog: map[string]IProduct]
            OS2[orderQueue: []string FIFO]
            OS3[sessions: map[string]bool]
            OS4[mu: sync.RWMutex]
        end

        subgraph Models["📦 internal/models/ - Product"]
            P[Product struct]
            P1[campos privados: id,name,price,stock]
            P2[NewProduct() con validación]
            P3[getters públicos]
            P4[ReduceStock() controlado]
        end

        subgraph Interfaces["🔌 internal/interfaces/"]
            I1[IProduct interface]
            I2[IOrderProcessor interface]
        end

        subgraph Errors["⚠️ internal/errors/"]
            E1[ErrProductNotFound]
            E2[ErrInsufficientStock]
            E3[ErrDuplicateRequest]
            E4[ErrPaymentFailed]
        end

        subgraph JSON["📋 encoding/json"]
            J1[Serialización]
            J2[Deserialización]
        end
    end

    %% Flujo principal
    C1 -->|Request HTTP + JSON| API
    API -->|Decode JSON| J2
    J2 -->|Llama interfaz| I2
    I2 -->|Implementa| OS
    OS -->|Usa interfaz| I1
    I1 -->|Implementa| P
    P -->|Valida invariantes| P2
    OS -->|Maneja errores| Errors
    OS -->|Respuesta struct| J1
    J1 -->|Encode JSON| API
    API -->|Response HTTP + JSON| C1

    %% Concurrencia
    OS4 -.->|RLock para lecturas| OS1
    OS4 -.->|Lock para escrituras| OS2
    OS4 -.->|Protege estado compartido| OS3

    %% Estilos
    classDef cliente fill:#e1f5fe,stroke:#01579b
    classDef api fill:#fff9c4,stroke:#fbc02d
    classDef services fill:#e8f5e9,stroke:#2e7d32
    classDef models fill:#f3e5f5,stroke:#7b1fa2
    classDef interfaces fill:#ffebee,stroke:#c62828
    classDef errors fill:#fff3e0,stroke:#ef6c00
    classDef json fill:#e0f7fa,stroke:#00838f

    class C1 cliente
    class H1,H2,H3,H4,H5,H6,H7,H8,H9 api
    class OS,OS1,OS2,OS3,OS4 services
    class P,P1,P2,P3,P4 models
    class I1,I2 interfaces
    class E1,E2,E3,E4 errors
    class J1,J2 json
