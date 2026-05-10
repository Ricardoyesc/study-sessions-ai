# Study Sessions AI

**Sistema de estudio adaptativo multimodal impulsado por IA.** Una plataforma que genera experiencias de aprendizaje personalizadas en tiempo real, combinando texto, imágenes, audio y evaluación adaptativa basada en el rendimiento del estudiante.

---

## Los 5 pilares

| Pilar | Descripción |
|-------|-------------|
| **A2UI Declarativo** | Backend envía árboles de componentes JSON. El frontend los renderiza nativamente. Sin código UI en el servidor. |
| **Codificación Dual** | Cada cápsula de aprendizaje entrega ≥2 modalidades sincronizadas: texto, audio, imagen. Teoría de Codificación Dual de Paivio. |
| **Diagnóstico Cold-Start** | Estima el nivel del estudiante sin datos demográficos usando Item Response Theory (IRT-3PL). |
| **Regla del 85%** | Algoritmo de selección de preguntas que converge al 85% de probabilidad de éxito — el punto óptimo de "dificultad deseable" (Wilson et al., 2019). |
| **Remediación Socrática** | Cuando el estudiante falla, el sistema genera explicaciones estilo Feynman + Interrogación Elaborativa. |

---

## Demo rápida

### Requisitos previos

- **Docker** + Docker Compose
- **Node.js 22+** (solo para el frontend en desarrollo)
- Claves API de OpenAI y Gemini (opcionales — la demo funciona sin ellas en modo fallback)

### Iniciar con Docker (recomendado)

```bash
# 1. Clonar el repositorio
git clone https://github.com/Ricardoyesc/study-sessions-ai.git
cd study-sessions-ai

# 2. Configurar variables de entorno
cp .env.example .env
# Editar .env con tus claves API (opcional)

# 3. Levantar servicios
docker compose up -d --build

# 4. Verificar que todo funciona
curl http://127.0.0.1:8080/healthz
# → {"status":"ok"}
```

### Iniciar el frontend

```bash
cd ui
npm install
npm run dev
# Abrir http://localhost:3000
```

### Credenciales demo

| Campo | Valor |
|-------|-------|
| Email | `demo@demo.com` |
| Password | `demo1234` |

El usuario demo se crea automáticamente al iniciar el servidor. Tiene perfil preconfigurado con nivel "intermediate".

### Flujo de la demo

```
Login → Dashboard → Elegir tema → Comenzar sesión
  → 📚 Cápsula (texto + imagen + audio)
  → ❓ Quiz (4 opciones generadas por IA)
  → ✅ Correcto → Siguiente tema (dificultad sube)
  → ❌ Incorrecto → Remediation socrática (Feynman + reflexión)
  → 🔁 Repetir hasta 8 ítems → Sesión completada
```

### Temas de prueba precargados

Método Científico · Física Cuántica · Teoría de la Relatividad · Álgebra Lineal · Biología Celular

---

## Arquitectura

```
┌──────────────┐     REST + WebSocket     ┌──────────────────────┐
│   Nuxt 3     │ ◄──────────────────────► │   Go (Gin + GORM)    │
│   Frontend   │     A2UI Surfaces        │   Backend            │
│  :3000       │                          │  :8080               │
└──────────────┘                          └──────────┬───────────┘
                                                     │
                                          ┌──────────┴───────────┐
                                          │   PostgreSQL / SQLite │
                                          │   Redis / MinIO       │
                                          └──────────────────────┘
```

### Componentes

| Capa | Tecnología | Propósito |
|------|-----------|-----------|
| **Frontend** | Nuxt 3 + Vue 3 + Tailwind + DaisyUI | Renderiza superficies A2UI declarativas. Dashboard con panel de estudio interactivo. |
| **Backend** | Go 1.24 + Gin + GORM | API REST + WebSocket. Orquesta LLM, TTS, generación de imágenes. |
| **LLM** | OpenAI GPT-4o | Genera contenido educativo, preguntas de quiz, explicaciones socráticas. |
| **Imagen** | Gemini Imagen + DALL·E 3 fallback | Diagramas e ilustraciones educativas. |
| **Audio** | OpenAI TTS | Narración de las cápsulas de estudio. |
| **DB** | SQLite (dev) / PostgreSQL (prod) | Usuarios, sesiones, interacciones, quizzes. |
| **Cache** | Redis | Sesiones A2UI en tiempo real, pub/sub. |
| **Storage** | MinIO / S3 | Assets multimedia (imágenes, audio). |

---

## Endpoints principales

| Método | Ruta | Auth | Descripción |
|--------|------|------|-------------|
| `POST` | `/api/users/register` | No | Registro de usuario |
| `POST` | `/api/users/login` | No | Login → JWT |
| `GET` | `/api/users/me` | Bearer | Perfil del usuario |
| `POST` | `/api/sessions` | Bearer | Crear sesión de estudio |
| `GET` | `/api/sessions/{id}/next` | Bearer | Siguiente ítem (cápsula/quiz) |
| `POST` | `/api/sessions/{id}/quiz/answer` | Bearer | Enviar respuesta de quiz |
| `POST` | `/api/sessions/{id}/socratic/response` | Bearer | Enviar reflexión socrática |
| `POST` | `/api/capsules/generate` | Bearer | Generar cápsula (texto+imagen+audio) |
| `WS` | `/ws/session/{id}` | No | Canal A2UI en tiempo real |

---

## Estructura del proyecto

```
study-sessions-ai/
├── src/                          # Backend Go
│   ├── cmd/server/main.go        # Entry point
│   ├── internal/
│   │   ├── api/                  # Handlers, router, middleware
│   │   ├── app/                  # Lógica de negocio
│   │   │   ├── dualcode/         # Orquestador de codificación dual
│   │   │   ├── quiz/             # Motor de preguntas + IRT
│   │   │   ├── socratic/         # Remediación Feynman + EI
│   │   │   ├── session/          # Pipeline de sesiones
│   │   │   └── user/             # Gestión de usuarios
│   │   ├── domain/               # Modelos GORM
│   │   ├── port/                 # Interfaces (LLM, TTS, Storage, Cache)
│   │   ├── adapter/              # Implementaciones concretas
│   │   │   ├── llm/openai.go     # Cliente OpenAI
│   │   │   ├── tts/openai.go     # TTS via OpenAI
│   │   │   ├── imagegen/         # Gemini + DALL·E fallback
│   │   │   ├── storage/minio.go  # Cliente MinIO/S3
│   │   │   └── cache/redis.go    # Cliente Redis
│   │   └── config/               # Configuración (viper)
│   ├── pkg/                      # Librerías compartidas
│   │   └── a2ui/                 # Motor A2UI (engine, builder, diff)
│   ├── migrations/               # Esquemas SQL
│   ├── deploy/docker/            # Dockerfile
│   └── configs/app.yaml          # Config YAML
├── ui/                           # Frontend Nuxt 3
│   ├── pages/                    # Login, Dashboard
│   ├── components/
│   │   ├── a2ui/                 # SurfaceRenderer + 13 primitivos
│   │   └── dashboard/            # StudyPanel, StudentSummary, SubjectSidebar
│   ├── composables/              # useAuth, useStudySession, useA2UI
│   └── types/                    # Tipos A2UI + Student
├── docker-compose.yml            # Servicios (Go + Postgres + Redis + MinIO)
├── architecture_plan.md          # Plan arquitectónico completo
└── CLAUDE.md                     # Guía para agentes
```

---

## Potencial de personalización

Esta PoC demuestra una arquitectura que escala a un sistema de aprendizaje **100% personalizado**.

### Lo que ya hace
- Genera contenido educativo **en tiempo real** sobre cualquier tema
- Adapta la **dificultad** según el rendimiento del estudiante
- Combina **múltiples modalidades** (texto, imagen, audio) por cada concepto
- Evalúa con **preguntas generadas por IA** (no banco fijo)
- Aplica **remediación socrática** personalizada cuando hay errores

### Lo que se puede construir sobre esta base

**Personalización por estudiante:**
- Modelado de conocimiento con Bayesian Knowledge Tracing (BKT)
- Espaciado de repasos con SM-2/SM-18
- Cold-start IRT sin datos demográficos (ya diseñado, no implementado en PoC)
- Clustering por patrones de aprendizaje
- Rutas de aprendizaje adaptativas por objetivo

**Contenido multimodal:**
- Videos generados por IA (Sora, Runway)
- Simulaciones interactivas (PhET-style)
- Ejercicios con drag & drop (ya diseñados en A2UI)
- Narración multilingüe (el sistema ya soporta español/inglés)

**Escala empresarial:**
- Kubernetes (manifiestos ya diseñados en `deploy/k8s/`)
- PostgreSQL + Redis + MinIO para producción
- WebSocket A2UI para actualizaciones en tiempo real
- Analytics de interacciones (ya se loguean en `interactions`)
- White-label con temas y accesibilidad configurables (A2UI dataModel)

**Despliegue como SaaS:**
- Multi-tenant con namespaces por organización
- Catálogo de temas curado por expertos + generación libre
- Dashboard de progreso para estudiantes y docentes
- Exportación de reportes de aprendizaje

---

## Comandos útiles

```bash
# Backend
cd src
make build          # Compilar
make run            # Ejecutar (desarrollo)
make dev            # Ejecutar con SQLite
go test ./...       # Tests

# Frontend
cd ui
npm run dev         # Dev server (localhost:3000)
npm run build       # Build producción
npm run preview     # Previsualizar build

# Docker
docker compose up -d --build   # Levantar todo
docker compose logs server     # Logs del backend
docker compose down -v         # Bajar y limpiar volúmenes
```

---

## Variables de entorno

Ver `.env.example` para la lista completa. Las más importantes:

| Variable | Default | Descripción |
|----------|---------|-------------|
| `OPENAI_API_KEY` | — | Para LLM, TTS y DALL·E |
| `GEMINI_API_KEY` | — | Para generación de imágenes (primario) |
| `SERVER_PORT` | `8080` | Puerto del backend |
| `ENVIRONMENT` | `development` | `development` usa SQLite, `production` PostgreSQL |
| `NUXT_PUBLIC_API_BASE` | `http://127.0.0.1:8080` | URL del backend (frontend) |
| `NUXT_PUBLIC_WS_BASE` | `ws://127.0.0.1:8080` | WebSocket URL (frontend) |

---

## Licencia

MIT
