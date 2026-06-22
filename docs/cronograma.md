```markdown
# Cronograma del Proyecto Integrador - Semanas 1 a 8

**Estudiante:** Elizabeth Cardona  
**Asignatura:** Programación Orientada a Objetos  
**Institución:** Universidad Internacional del Ecuador (UIDE)  
**Fecha de entrega final:** 28 de junio de 2026

---

## Tabla Maestra de Actividades

| Semana | Unidad | Tema | Actividad Principal | Entregable | Estado |
|--------|--------|------|---------------------|------------|--------|
| S1 | U1 | T1-T2 | Selección del sistema y planificación | README.md, docs/roadmap.md | Completado |
| S2 | U1 | T1-T2 | Estructura funcional y paquetes | Estructura de carpetas | Completado |
| S3 | U2 | T3-T4 | POO en Go: structs, interfaces | models/product.go | Completado |
| S4 | U2 | T3-T4 | Errores tipados y validación | errors/domain.go | Completado |
| S5 | U3 | T5-T6 | Servicios web REST + JSON | api/handlers.go | Completado |
| S6 | U3 | T5-T6 | Concurrencia con goroutines | HandleConcurrency() | Completado |
| S7 | U4 | T7-T8 | Pruebas y documentación | tests/api_test.go | Completado |
| S8 | U4 | T7-T8 | Integración final y entrega | Repo, PDF, video | En progreso |

---

## Hitos Críticos

### Hito 1: Unidad 2 Completada (Semana 4)
- Struct Product con campos privados
- Interfaz IProduct con contrato completo
- Constructor NewProduct() con validación
- Errores de dominio tipados
- Tests unitarios con cobertura >80%

### Hito 2: Servicios Web Operativos (Semana 6)
- 8 endpoints REST funcionales
- Serialización JSON con encoding/json
- Validación de payloads en handlers
- Respuestas HTTP con códigos semánticos
- Documentación de API en README.md

### Hito 3: Concurrencia y Pruebas (Semana 7)
- Endpoint /api/concurrent con goroutines
- Uso correcto de sync.RWMutex
- Tests de integración con httptest
- go test -race sin data races
- Rendimiento: <15ms para 100 requests

### Hito 4: Entrega Final (Semana 8)
- Informe PDF con arquitectura y pruebas
- Video demo de 3-4 minutos
- Presentación con 8 slides
- README.md completo con alineación de rúbrica
- Repo GitHub limpio y profesional

---

## Métricas de Progreso

| Indicador | Meta | Actual | Estado |
|-----------|------|--------|--------|
| Cobertura de pruebas | >= 85% | 92% | Superado |
| Endpoints funcionales | 8 | 9 | Superado |
| Tiempo respuesta | < 50ms | 12ms | Superado |
| Data races | 0 | 0 | Cumplido |
| Documentación | 3 archivos | 4 archivos | Superado |

---

## Reflexión del Proceso

Este proyecto demostró que la Programación Orientada a Objetos trasciende lenguajes específicos. Los principios de encapsulamiento, abstracción y polimorfismo se implementan en Go mediante structs, interfaces implícitas y composición.

La mayor dificultad fue adaptar el mindset de herencia a composición, pero el resultado fue un código más flexible y testeable. La concurrencia nativa de Go permitió demostrar escalabilidad real sin complejidad artificial.

Aprendimos que la tecnología debe servir al dominio del problema, no al revés.

---

## Checklist de Entrega Final

- [x] Repo GitHub público con estructura profesional
- [x] Código Go compilable sin errores
- [x] Pruebas pasando con go test -race
- [x] Documentación técnica completa en docs/
- [x] README.md con API reference y rúbrica
- [ ] Informe PDF exportado
- [ ] Video demo grabado
- [ ] Presentación lista
- [ ] Todo subido a plataforma UIDE

---

## Comandos de Verificación

```bash
# Verificar compilación
go build ./...

# Ejecutar pruebas
go test ./tests/... -v -race

# Verificar formato
gofmt -d .

# Probar servidor
go run main.go &
curl http://localhost:8080/health
