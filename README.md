# Sistema de Gestión de E-commerce (Golang)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

> Implementación académica de un sistema empresarial en Go, integrando Programación Orientada a Objetos, encapsulamiento, manejo explícito de errores, concurrencia nativa y 8 servicios web con serialización JSON.

---

## Datos del Proyecto

| Campo | Valor |
|-------|-------|
| Estudiante | Elizabeth Cardona |
| Asignatura | Programación Orientada a Objetos |
| Institución | Universidad Internacional del Ecuador (UIDE) |
| Fecha de entrega | 28 de junio de 2026 |
| Lenguaje | Go 1.22+ |

---

## Objetivo del Sistema

Desarrollar una capa de dominio robusta para un e-commerce aplicando:
- Encapsulamiento: Campos privados + getters públicos
- Interfaces: Contratos implícitos para desacoplamiento
- Manejo de errores: Tipos error personalizados para fallos de dominio
- Concurrencia segura: sync.RWMutex, goroutines y WaitGroup
- Servicios web: 8 endpoints REST con serialización JSON nativa

---

## Ejecución Rápida

```bash
# 1. Clonar el proyecto
git clone https://github.com/ellliecar/ecommerce-go-system.git
cd ecommerce-go-system

# 2. Descargar dependencias
go mod tidy

# 3. Ejecutar pruebas
go test ./tests/... -v

# 4. Iniciar servidor
go run main.go

# Servidor activo en: http://localhost:8080
