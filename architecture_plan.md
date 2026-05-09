# Arquitectura del Sistema de Estudio Adaptativo Multimodal (PoC)

## Resumen Ejecutivo

Este documento presenta la propuesta arquitectónica completa para una Prueba de Concepto (PoC) de un sistema de estudio adaptativo multimodal basado en IA. El sistema implementa los cinco pilares funcionales especificados:

1. **Motor de Interfaz Declarativa A2UI** — Backend envía descripciones de componentes JSON, cliente renderiza componentes nativos
2. **Orquestador de Codificación Dual** — Genera contenido en al menos dos modalidades sincronizadas (Teoría de Codificación Dual)
3. **Módulo de Diagnóstico Cold-Start** — Evaluación adaptativa sin datos demográficos, usando IRT y clustering
4. **Bucle de Evaluación (Regla del 85%)** — Algoritmo de selección de preguntas que converge al 85% de probabilidad de éxito
5. **Sistema de Remediación Socrática** — Intervenciones basadas en Técnica Feynman + Interrogación Elaborativa

---

## 1. Arquitectura General

### 1.1 Diagrama de Componentes

```
┌──────────────────────────────────────────────────────────────────────────┐
│                         CLUSTER KUBERNETES                                │
│                                                                           │
│  ┌─────────────┐  ┌──────────────────────────────────────────────────┐   │
│  │   Ingress   │  │          Monolito Go (sai-server)                │   │
│  │   (nginx)   │  │                                                  │   │
│  │             │  │  ┌─────────┐  ┌─────────┐  ┌──────────────┐    │   │
│  │  /api/* ────┼─►│ API REST │  │ WebSocket│  │  A2UI Engine │    │   │
│  │  /ws/*  ────┼─►│  (Gin)   │  │  (ws)    │  │  (JSON Gen)  │    │   │
│  │             │  │  └────┬────┘  └────┬────┘  └──────┬───────┘    │   │
│  └─────────────┘  │       │            │              │            │   │
│                   │  ┌────┴────────────┴──────────────┴───────┐    │   │
│                   │  │          Service Layer                  │    │   │
│                   │  │  ┌──────────┐  ┌──────────┐  ┌─────────┐ │    │   │
│                   │  │  │ Session  │  │  User    │  │ ColdStart│ │    │   │
│                   │  │  │ Manager  │  │ Service  │  │ Service │ │    │   │
│                   │  │  └──────────┘  └──────────┘  └─────────┘ │    │   │
│                   │  │  ┌──────────┐  ┌──────────┐  ┌─────────┐ │    │   │
│                   │  │  │ Quiz     │  │DualCode  │  │ Socratic│ │    │   │
│                   │  │  │ Engine   │  │Orchest.  │  │Remediat.│ │    │   │
│                   │  │  └──────────┘  └──────────┘  └─────────┘ │    │   │
│                   │  └───────────────────────────────────────┘    │   │
│                   │  ┌───────────────────────────────────────┐    │   │
│                   │  │        Ports / Adapters                │    │   │
│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │
│                   │  │  │ LLM  │ │ TTS  │ │ Image/Video  │  │    │   │
│                   │  │  │Client│ │Client│ │   Client     │  │    │   │
│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │
│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │
│                   │  │  │  DB  │ │Redis │ │   MinIO/S3   │  │    │   │
│                   │  │  │(GORM)│ │(Pub) │ │  (Assets)    │  │    │   │
│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │
│                   │  └───────────────────────────────────────┘    │   │
│                   └──────────────────────────────────────────────────┘   │
│                                                                           │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐         │
│  │ PostgreSQL 14    │  │    Redis 7       │  │   MinIO (S3)    │         │
│  │  (StatefulSet)   │  │  (StatefulSet)   │  │ (StatefulSet)   │         │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘         │
└──────────────────────────────────────────────────────────────────────────┘

                    ┌──────────────────────────────────────┐
                    │          APIs Externas                │
                    │  ┌────────┐  ┌──────────┐  ┌───────┐│
                    │  │ OpenAI │  │Gemini TTS│  │ DALL·E││
                    │  │GPT-4o  │  │(Google)  │  │       ││
                    │  └────────┘  └──────────┘  └───────┘│
                    └──────────────────────────────────────┘

  ┌─────────────────────────────────────────────────────────────────┐
  │                   Cliente A2UI (Nuxt 3 - Repositorio Separado) │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │  A2UI Renderer Engine                                        │ │
│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────┐  │ │
│  │  │Component │ │ Layout   │ │Theme/    │ │ Event        │  │ │
│  │  │Registry  │ │ Resolver │ │A11y Mgr  │ │ Dispatcher   │  │ │
│  │  └──────────┘ └──────────┘ └──────────┘ └──────────────┘  │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  Native Components: Text, RichText, Image, AudioPlayer,         │
│  VideoPlayer, Card, QuizCard, SocraticDialog, ProgressBar        │
└─────────────────────────────────────────────────────────────────┘
```

### 1.2 Stack Tecnológico

| Componente | Tecnología | Versión |
|------------|------------|---------|
| Lenguaje Backend | Go | 1.22+ |
| Framework HTTP | Gin | latest |
| Base de Datos (Prod) | PostgreSQL | 14+ |
| Base de Datos (Dev) | SQLite | 3.x |
| Cache/Colas | Redis | 7.x |
| Storage Objetos | MinIO | latest |
| LLM | OpenAI | GPT-4o |
| TTS | Google Gemini TTS | v1 |
| Imágenes | DALL·E 3 / Stability AI | latest |
| Contenedores | Docker | latest |
| Orquestación | Kubernetes | 1.28+ |
| Frontend | Nuxt 3 + Vue 3 | latest |

### 1.3 Flujo de Datos Principal

1. **Usuario nuevo** → `/api/coldstart/start` → IRT adaptive test → `θ` estimado → clustering
2. **Sesión activa** → WebSocket `/ws/session/{id}` → A2UI surface inicial
3. **Pregunta quiz** → algoritmo 85% → selecciona pregunta → `a2ui_update` → usuario responde
4. **Respuesta incorrecta** → `/api/socratic/remediate` → LLM Feynman → Interrogación Elaborativa
5. **Generación cápsula** → `/api/capsules/generate` → dual coding orchestrator → texto + audio + imagen → A2UI tree

---

## 2. Diseño de API REST + WebSocket

### 2.1 Endpoints REST

| Método | Ruta | Descripción |
|--------|------|-------------|
| `POST` | `/api/users/register` | Registro de usuario (sin datos demográficos) |
| `POST` | `/api/users/login` | Login, retorna JWT |
| `GET` | `/api/users/me` | Perfil de usuario con habilidad estimada |
| `POST` | `/api/coldstart/start` | Inicia diagnóstico cold-start |
| `POST` | `/api/coldstart/answer` | Envía respuesta a ítem de diagnóstico |
| `GET` | `/api/coldstart/{id}/result` | Resultado estimado (θ + cluster) |
| `POST` | `/api/sessions` | Crea nueva sesión de estudio |
| `GET` | `/api/sessions/{id}` | Estado actual de la sesión |
| `GET` | `/api/sessions/{id}/next` | Obtiene siguiente ítem (cápsula o quiz) |
| `POST` | `/api/sessions/{id}/quiz/answer` | Envía respuesta a pregunta de quiz |
| `POST` | `/api/sessions/{id}/socratic/response` | Envía respuesta a intervención socrática |
| `POST` | `/api/capsules/generate` | Solicita generación de cápsula dual |
| `GET` | `/api/capsules/{id}` | Obtiene胶囊 completa |
| `POST` | `/api/sessions/{id}/a11y` | Actualiza preferencias accesibilidad |
| `GET` | `/api/assets/{type}/{filename}` | Descarga assets (audio/imagen) |

### 2.2 WebSocket

**Endpoint**: `WS /ws/session/{sessionId}`

#### Formato de Mensaje (Envelope JSON)

```json
{
  "type": "a2ui_full | a2ui_update | data_model_update | error | ping",
  "payload": {},
  "timestamp": "2024-01-15T10:30:00Z"
}
```

### 2.3 Protocolo A2UI: Tipos de Mensajes

#### a2ui_full — Renderizado Inicial Completo

```json
{
  "type": "a2ui_full",
  "payload": {
    "surfaceId": "session-abc",
    "rootComponent": "root",
    "components": {
      "root": {
        "id": "root",
        "type": "Column",
        "children": ["header", "body", "footer"],
        "props": { "gap": 16, "padding": 24 }
      },
      "header": {
        "id": "header",
        "type": "Row",
        "children": ["title", "progress-bar"],
        "props": { "alignment": "space-between" }
      },
      "title": {
        "id": "title",
        "type": "Text",
        "props": { "content": "Física Cuántica: Dualidad Onda-Partícula", "variant": "h2" }
      },
      "progress-bar": {
        "id": "progress-bar",
        "type": "ProgressBar",
        "props": { "value": 0.35, "max": 1.0 }
      },
      "body": {
        "id": "body",
        "type": "Card",
        "children": ["text-content", "audio-player", "image-diagram"],
        "props": { "elevation": 2 }
      },
      "text-content": {
        "id": "text-content",
        "type": "RichText",
        "props": { "markdown": "La dualidad onda-partícula es un concepto fundamental...", "accessible": true }
      },
      "audio-player": {
        "id": "audio-player",
        "type": "AudioPlayer",
        "props": { "url": "/api/assets/audio/capsule-123.mp3", "autoPlay": false }
      },
      "image-diagram": {
        "id": "image-diagram",
        "type": "Image",
        "props": { "url": "/api/assets/image/diagram-123.png", "altText": "Diagrama del experimento de doble rendija" }
      },
      "quiz-card": {
        "id": "quiz-card",
        "type": "QuizCard",
        "props": {
          "question": "¿Por qué el patrón de interferencia desaparece al observar las partículas?",
          "options": ["A) La luz se comporta como onda", "B) La medición perturba el sistema cuántico", "C) El detector cambia la frecuencia", "D) La velocidad cambia"],
          "mode": "single_choice"
        },
        "events": { "onSubmit": "/api/sessions/abc/quiz/answer" }
      },
      "socratic-prompt": {
        "id": "socratic-prompt",
        "type": "SocraticDialog",
        "props": {
          "prompt": "Antes de continuar, explica con tus palabras: ¿por qué la medición colapsa la función de onda?",
          "context": "dualidad-onda-particula"
        },
        "events": { "onSubmit": "/api/sessions/abc/socratic/response" }
      },
      "footer": {
        "id": "footer",
        "type": "Row",
        "children": ["btn-prev", "btn-next"],
        "props": { "alignment": "center" }
      }
    },
    "dataModel": {
      "theme": "system",
      "fontFamily": "sans-serif",
      "fontScale": 1.0,
      "colorPalette": "default",
      "highContrast": false,
      "reducedMotion": false,
      "language": "es"
    }
  }
}
```

#### data_model_update — Cambio de Accesibilidad en Caliente

```json
{
  "type": "data_model_update",
  "payload": {
    "path": "fontFamily",
    "value": "OpenDyslexic",
    "diff": {
      "fontFamily": "OpenDyslexic",
      "colorPalette": "pastel",
      "fontScale": 1.2
    }
  }
}
```

#### a2ui_update — Parche Parcial del Árbol

```json
{
  "type": "a2ui_update",
  "payload": {
    "updates": [
      { "componentId": "quiz-card", "props": { "question": "Nueva pregunta...", "options": ["A", "B", "C"] }},
      { "componentId": "progress-bar", "props": { "value": 0.70 }}
    ]
  }
}
```

---

## 3. Protocolo A2UI: Modelo de Datos

### 3.1 Estructuras Go

```go
package a2ui

// Surface representa una "pantalla" completa de A2UI
type Surface struct {
    SurfaceID     string                `json:"surfaceId"`
    RootComponent string                `json:"rootComponent"`
    Components    map[string]Component  `json:"components"`
    DataModel     DataModel             `json:"dataModel"`
}

// Component es un nodo del árbol declarativo (Modelo de Lista de Adyacencia)
type Component struct {
    ID       string                 `json:"id"`
    Type     string                 `json:"type"` // "Text", "Card", "QuizCard", etc.
    Children []string               `json:"children,omitempty"`
    Props    map[string]interface{} `json:"props"`
    Events   map[string]string      `json:"events,omitempty"`
}

// DataModel contiene estado global de accesibilidad/tema
type DataModel struct {
    Theme         string  `json:"theme"`
    FontFamily    string  `json:"fontFamily"`
    FontScale     float64 `json:"fontScale"`
    ColorPalette  string  `json:"colorPalette"`
    HighContrast  bool    `json:"highContrast"`
    ReducedMotion bool    `json:"reducedMotion"`
    Language      string  `json:"language"`
}

// WSMessage es el envelope de WebSocket
type WSMessage struct {
    Type      string      `json:"type"`
    Payload   interface{} `json:"payload"`
    Timestamp string      `json:"timestamp"`
}
```

### 3.2 Componentes Soportados

| Tipo | Descripción |
|------|-------------|
| `Text` | Texto simple con variantes (h1-h6, body, caption) |
| `RichText` | Markdown rendering |
| `Image` | Imagen con altText |
| `AudioPlayer` | Reproductor de audio con controles |
| `VideoPlayer` | Reproductor de video con subtítulos |
| `Card` | Contenedor con elevación |
| `Column` | Contenedor vertical |
| `Row` | Contenedor horizontal |
| `QuizCard` | Tarjeta de pregunta con opciones |
| `SocraticDialog` | Diálogo socrático con input |
| `ProgressBar` | Barra de progreso |
| `Button` | Botón con tipos (primary, secondary) |

---

## 4. Módulos de Dominio

### 4.1 Cold-Start Diagnostic (Diagnóstico de Inicio en Frío)

**Objetivo**: Estimar el nivel de habilidad latente del usuario sin usar datos demográficos.

**Algoritmo**:

1. **Item Response Theory (IRT)** — Usa modelo 3PL:
   - `P(correct) = c + (1-c) * 1/(1 + exp(-a(θ-b)))`
   - `a` = discriminación, `b` = dificultad, `c` = adivinación

2. **Computerized Adaptive Testing (CAT)**:
   - Selecciona siguiente ítem por Maximum Fisher Information
   - I(theta) = sum over items of (a_i^2 * P_i * (1-P_i) * ((P_i - c_i)/(1-c_i))^2)
   - Actualiza θ con Maximum Likelihood Estimation

3. **K-Means Clustering**:
   - Después de N respuestas (ej: 10), agrupa patrones de respuesta
   - Asigna cluster como "nivel inicial" (beginner, intermediate, advanced)

**Restricciones**:
- ❌ NO usar edad, género, ubicación, educación
- ✅ Solo usar patrones de interacción y respuestas a ítems calibrados

### 4.2 Quiz Engine (Regla del 85%)

**Objetivo**: Mezclar material no dominado y repasos para que la probabilidad predictiva de éxito converja al 85%.

**Algoritmo**:

```
P_target = 0.85

Para cada pregunta candidatos:
  P_success = IRT_3PL(θ_user, a, b, c)

Categorizar:
  Dominadas:     P_success >= 0.90
  Near-mastery:  0.70 <= P_success < 0.90
  No-dominadas:  P_success < 0.70

Mezcla óptima:
  α = proporción de no-dominadas
  (1-α) = proporción de near-mastery (repaso)
  
  Buscar α tal que:
    α * avg(P_no_dominadas) + (1-α) * avg(P_near_mastery) ≈ 0.85
```

**Integrado con**:
- **Bayesian Knowledge Tracing (BKT)** — Actualiza estado de dominio por concepto
- **SM-2 Spaced Repetition** — Calcula intervalos de repaso óptimos
- **Interleaving** — Evita cluster de mismo concepto en sucesión

### 4.3 Dual Coding Orchestrator

**Objetivo**: Cada cápsula de aprendizaje entrega información en al menos dos modalidades sincronizadas.

**Arquitectura de datos por cápsula**:

```go
type DualCapsule struct {
    ID          string
    Topic       string
    Modalities  []Modality // Text + Audio + Image mínimo
    A2UI_Tree   a2ui.Surface
}

type Modality struct {
    Type      string // "text", "audio", "image", "video"
    Content   string // texto o URL
    Metadata  map[string]interface{}
}
```

**Flujo de generación**:
1. Recibe topic → LLM genera texto educativo
2. En paralelo: TTS genera audio, ImageGen genera imagen/diagrama
3. Ensambla en DualCapsule
4. Builder convierte a A2UI tree
5. **Valida**: `len(Modalidades) >= 2`

### 4.4 Socratic Remediation

**Objetivo**: Cuando el estudiante falla, generar intervención basada en Técnica Feynman + Interrogación Elaborativa.

**Flujo**:

1. Estudiante responde incorrectamente
2. Clasificar tipo de error (conceptual, procedimental, olvida)
3. LLM genera explicación estilo Feynman:
   - Identifica concepto complejo
   - Explica en términos simples
   - Usa analogía de la vida real

4. **Interrogación Elaborativa**:
   - "¿Por qué crees que ocurre X?"
   - Evalúa respuesta del usuario con LLM
   - Decide: más scaffolding o avanzar

---

## 5. Modelos de Datos (PostgreSQL)

### 5.1 Esquema de Tablas

```sql
-- Conceptos (Knowledge Components)
CREATE TABLE concepts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    parent_id UUID REFERENCES concepts(id),
    name TEXT NOT NULL,
    description TEXT,
    difficulty FLOAT NOT NULL DEFAULT 0.5,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Usuarios (SIN metadatos demográficos)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    estimated_theta FLOAT,
    theta_uncertainty FLOAT,
    cluster VARCHAR(50),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Dominio por usuario-concepto (BKT state)
CREATE TABLE user_concept_mastery (
    user_id UUID REFERENCES users(id),
    concept_id UUID REFERENCES concepts(id),
    p_learned FLOAT NOT NULL DEFAULT 0.3,
    p_guess FLOAT NOT NULL DEFAULT 0.1,
    p_slip FLOAT NOT NULL DEFAULT 0.1,
    p_transit FLOAT NOT NULL DEFAULT 0.2,
    last_practiced TIMESTAMPTZ,
    easiness_factor FLOAT DEFAULT 2.5,
    interval_days INT DEFAULT 1,
    repetitions INT DEFAULT 0,
    next_review TIMESTAMPTZ,
    PRIMARY KEY (user_id, concept_id)
);

-- Items de quiz (IRT calibrated)
CREATE TABLE quiz_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    concept_id UUID REFERENCES concepts(id),
    difficulty_irt FLOAT NOT NULL,
    discrimination FLOAT DEFAULT 1.0,
    guessing FLOAT DEFAULT 0.25,
    content JSONB NOT NULL,
    modality TEXT DEFAULT 'text',
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Sesiones de estudio
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    state TEXT NOT NULL DEFAULT 'coldstart',
    a2ui_snapshot JSONB,
    target_success_rate FLOAT DEFAULT 0.85,
    started_at TIMESTAMPTZ DEFAULT now(),
    completed_at TIMESTAMPTZ
);

-- Cápsulas de aprendizaje
CREATE TABLE capsules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    topic TEXT NOT NULL,
    modalities JSONB NOT NULL,
    a2ui_tree JSONB NOT NULL,
    session_id UUID REFERENCES sessions(id),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Interacciones (logs para analytics)
CREATE TABLE interactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES sessions(id),
    type TEXT NOT NULL,
    payload JSONB NOT NULL,
    was_correct BOOLEAN,
    response_time_ms INT,
    remediation_generated JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);
```

---

## 6. Estructura de Carpetas del Proyecto Go

```
study-sessions-with-ai/
├── cmd/
│   └── server/
│       └── main.go                    # Entry point
├── internal/
│   ├── api/
│   │   ├── router.go                  # Registro de rutas Gin
│   │   ├── middleware/
│   │   │   ├── auth.go                # JWT middleware
│   │   │   ├── logging.go
│   │   │   └── cors.go
│   │   └── handlers/
│   │       ├── a2ui_handler.go        # WebSocket handler
│   │       ├── coldstart_handler.go
│   │       ├── session_handler.go
│   │       ├── capsule_handler.go
│   │       ├── quiz_handler.go
│   │       ├── socratic_handler.go
│   │       └── user_handler.go
│   ├── app/
│   │   ├── session/
│   │   │   ├── service.go
│   │   │   ├── pipeline.go
│   │   │   └── types.go
│   │   ├── coldstart/
│   │   │   ├── service.go
│   │   │   ├── irt.go
│   │   │   ├── cat.go
│   │   │   └── types.go
│   │   ├── dualcode/
│   │   │   ├── orchestrator.go
│   │   │   ├── capsule.go
│   │   │   └── types.go
│   │   ├── quiz/
│   │   │   ├── engine.go
│   │   │   ├── selector.go
│   │   │   ├── evaluator.go
│   │   │   └── types.go
│   │   ├── socratic/
│   │   │   ├── remediator.go
│   │   │   ├── prompts.go
│   │   │   └── types.go
│   │   └── user/
│   │       ├── service.go
│   │       └── types.go
│   ├── domain/
│   │   ├── user.go
│   │   ├── concept.go
│   │   ├── session.go
│   │   ├── capsule.go
│   │   ├── quiz.go
│   │   └── a2ui.go
│   ├── port/
│   │   ├── llm.go
│   │   ├── tts.go
│   │   ├── imagegen.go
│   │   ├── storage.go
│   │   └── cache.go
│   ├── adapter/
│   │   ├── llm/
│   │   │   ├── openai.go
│   │   │   └── factory.go
│   │   ├── tts/
│   │   │   ├── gemini.go
│   │   │   └── factory.go
│   │   ├── imagegen/
│   │   │   ├── dalle.go
│   │   │   └── factory.go
│   │   ├── storage/
│   │   │   └── minio.go
│   │   └── cache/
│   │       └── redis.go
│   └── config/
│       └── config.go
├── pkg/
│   ├── a2ui/
│   │   ├── protocol.go
│   │   ├── builder.go
│   │   ├── diff.go
│   │   └── types.go
│   ├── spacedrep/
│   │   ├── sm2.go
│   │   └── scheduler.go
│   ├── irt/
│   │   ├── model.go
│   │   └── estimator.go
│   ├── bkt/
│   │   └── bayesian.go
│   └── adaptive/
│       └── difficulty.go
├── migrations/
├── deploy/
│   ├── k8s/
│   │   ├── namespace.yaml
│   │   ├── configmap.yaml
│   │   ├── secrets.yaml
│   │   ├── deployment-postgres.yaml
│   │   ├── deployment-redis.yaml
│   │   ├── deployment-minio.yaml
│   │   ├── deployment-server.yaml
│   │   └── ingress.yaml
│   └── docker/
│       └── Dockerfile
├── configs/
│   └── app.yaml
├── scripts/
├── go.mod
├── go.sum
├── Makefile
├── docker-compose.yml
└── .env.example
```

---

## 7. Frontend: Repositorio Separado (Nuxt 3)

### 7.1 Estructura de Carpetas

```
sai-frontend/
├── nuxt.config.ts
├── package.json
├── tsconfig.json
├── pages/
│   ├── index.vue
│   ├── coldstart.vue
│   └── session/
│       └── [id].vue
├── composables/
│   ├── useA2UI.ts
│   ├── useA2UIAccessibility.ts
│   ├── useWebSocket.ts
│   └── useStudySession.ts
├── registry/
│   ├── index.ts
│   ├── Text.vue
│   ├── RichText.vue
│   ├── Image.vue
│   ├── AudioPlayer.vue
│   ├── VideoPlayer.vue
│   ├── Card.vue
│   ├── Column.vue
│   ├── Row.vue
│   ├── QuizCard.vue
│   ├── SocraticDialog.vue
│   ├── ProgressBar.vue
│   └── Button.vue
├── components/
│   └── a2ui/
│       └── SurfaceRenderer.vue
├── stores/
│   └── a2ui.ts
├── types/
│   └── a2ui.ts
├── assets/
│   └── css/
│       ├── themes/
│       │   ├── light.css
│       │   ├── dark.css
│       │   └── pastel.css
│       └── accessibility.css
└── middleware/
    └── auth.ts
```

### 7.2 Composables Core

```typescript
// useA2UI.ts
import { ref, watch } from 'vue';
import type { A2UISurface, WSMessage, DataModelUpdate, A2UIUpdate } from '~/types/a2ui';
import { useWebSocket } from './useWebSocket';

export function useA2UI(sessionId: string) {
  const surface = ref<A2UISurface | null>(null);
  const dataModel = ref<A2UIDataModel | null>(null);
  const { connect, send, close } = useWebSocket(`ws://${host}/ws/session/${sessionId}`);

  function handleMessage(msg: WSMessage) {
    switch (msg.type) {
      case 'a2ui_full':
        surface.value = msg.payload as A2UISurface;
        dataModel.value = surface.value.dataModel;
        applyAccessibilityStyles(surface.value.dataModel);
        break;
      case 'a2ui_update':
        // Apply patches
        break;
      case 'data_model_update':
        // Hot-reload accessibility
        break;
    }
  }

  connect(handleMessage);
  return { surface, dataModel, send };
}
```

---

## 8. Manifiestos Kubernetes

### 8.1 namespace.yaml

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: sai-learning
```

### 8.2 configmap.yaml

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: sai-server-config
  namespace: sai-learning
data:
  DB_HOST: "postgres-service"
  DB_PORT: "5432"
  DB_NAME: "sai_learning"
  REDIS_URL: "redis-service:6379"
  MINIO_ENDPOINT: "minio-service:9000"
  LLM_PROVIDER: "openai"
  LLM_MODEL: "gpt-4o"
  TTS_PROVIDER: "gemini"
```

### 8.3 secrets.yaml

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: sai-server-secrets
  namespace: sai-learning
type: Opaque
stringData:
  DB_USER: "sai_user"
  DB_PASSWORD: "change-me-in-production"
  OPENAI_API_KEY: "sk-..."
  GEMINI_API_KEY: "ai..."
  MINIO_ACCESS_KEY: "minioadmin"
  MINIO_SECRET_KEY: "minioadmin"
  JWT_SECRET: "change-me-jwt-secret"
```

### 8.4 deployment-server.yaml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: sai-server
  namespace: sai-learning
spec:
  replicas: 2
  selector:
    matchLabels:
      app: sai-server
  template:
    metadata:
      labels:
        app: sai-server
    spec:
      containers:
      - name: sai-server
        image: ghcr.io/your-org/sai-server:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 8080
        - containerPort: 8081
        envFrom:
        - configMapRef:
            name: sai-server-config
        - secretRef:
            name: sai-server-secrets
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: sai-server-service
  namespace: sai-learning
spec:
  selector:
    app: sai-server
  ports:
  - name: http
    port: 80
    targetPort: 8080
```

### 8.5 ingress.yaml

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: sai-ingress
  namespace: sai-learning
  annotations:
    nginx.ingress.kubernetes.io/proxy-read-timeout: "3600"
    nginx.ingress.kubernetes.io/proxy-send-timeout: "3600"
    nginx.ingress.kubernetes.io/websocket-services: "sai-server-service"
spec:
  rules:
  - host: sai-learning.local
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: sai-server-service
            port:
              number: 80
```

---

## 9. Plan de Implementación por Etapas

### Fase 0: Fundación (4 días)

| # | Tarea |
|---|-------|
| 0.1 | Init módulo Go + Makefile |
| 0.2 | Configuración con viper |
| 0.3 | Tipos de dominio (GORM models) |
| 0.4 | Interfaces de puertos |
| 0.5 | Protocolo A2UI tipos Go |
| 0.6 | Router Gin + middleware |
| 0.7 | main.go wire-up |
| 0.8 | Migraciones SQL |
| 0.9 | Dockerfile + docker-compose |

### Fase 1: Adapters + A2UI Engine (6 días, paralelos)

**Agente B**: A2UI Engine + Handlers core
**Agente C**: Adapters de IA (OpenAI, Gemini TTS, DALL·E, MinIO)

### Fase 2: Módulos de Aprendizaje (8 días, paralelos)

**Agente D**: Cold-Start Diagnostic
**Agente E**: Quiz Engine (85% Rule)
**Agente F**: Dual Coding Orchestrator

### Fase 3: Session Pipeline + Socratic (6 días, paralelos)

**Agente G**: Session Manager
**Agente H**: Socratic Remediation

### Fase 4: Frontend + K8s + Testing (6 días, paralelos)

**Agente I**: Frontend Nuxt 3
**Agente J**: K8s + CI/CD + Tests

---

## 10. Ejemplo de main.go

```go
package main

import (
    "log/slog"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/your-org/sai-server/internal/config"
    "github.com/your-org/sai-server/internal/domain"
    "github.com/your-org/sai-server/internal/port"
    "github.com/your-org/sai-server/internal/adapter/llm"
    "github.com/your-org/sai-server/internal/adapter/tts"
    "github.com/your-org/sai-server/internal/adapter/storage"
    "github.com/your-org/sai-server/internal/adapter/cache"
    "github.com/your-org/sai-server/internal/app/user"
    "github.com/your-org/sai-server/internal/app/session"
    "github.com/your-org/sai-server/internal/app/coldstart"
    "github.com/your-org/sai-server/internal/app/quiz"
    "github.com/your-org/sai-server/internal/app/dualcode"
    "github.com/your-org/sai-server/internal/app/socratic"
    "github.com/your-org/sai-server/internal/api/handlers"
    a2uiEngine "github.com/your-org/sai-server/pkg/a2ui"
)

func main() {
    cfg := config.Load()

    var dialector gorm.Dialector
    if cfg.IsProduction() {
        dialector = postgres.Open(cfg.DatabaseURL)
    } else {
        dialector = sqlite.Open("data/sai_dev.db")
    }
    db, err := gorm.Open(dialector, &gorm.Config{})
    if err != nil {
        slog.Error("failed to connect database", "error", err)
        os.Exit(1)
    }

    db.AutoMigrate(
        &domain.User{},
        &domain.Concept{},
        &domain.Session{},
        &domain.Capsule{},
        &domain.QuizItem{},
        &domain.Interaction{},
    )

    var llmClient port.LLMClient
    if cfg.LLMProvider == "openai" {
        llmClient = llm.NewOpenAI(cfg.OpenAIKey, cfg.LLMModel)
    }

    var ttsClient port.TTSClient
    if cfg.TTSProvider == "gemini" {
        ttsClient = tts.NewGemini(cfg.GeminiAPIKey)
    }

    storageClient := storage.NewMinIO(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)
    cacheClient := cache.NewRedis(cfg.RedisURL)

    a2ui := a2uiEngine.NewEngine()

    userSvc := user.NewService(db)
    coldStartSvc := coldstart.NewService(db, cacheClient)
    quizEngine := quiz.NewEngine(db, coldStartSvc)
    dualCodeOrch := dualcode.NewOrchestrator(llmClient, ttsClient, storageClient, a2ui)
    socraticRem := socratic.NewRemediator(llmClient)
    sessionSvc := session.NewService(db, coldStartSvc, dualCodeOrch, quizEngine, socraticRem)

    r := gin.Default()
    r.Use(handlers.CORS(), handlers.Logger())

    handlers.RegisterRoutes(r, &handlers.Handlers{
        User:      handlers.NewUserHandler(userSvc),
        Session:   handlers.NewSessionHandler(sessionSvc),
        ColdStart: handlers.NewColdStartHandler(coldStartSvc),
        Quiz:      handlers.NewQuizHandler(quizEngine),
        Capsule:   handlers.NewCapsuleHandler(dualCodeOrch),
        Socratic:  handlers.NewSocraticHandler(socraticRem),
        A2UI:      handlers.NewA2UIHandler(sessionSvc, a2ui),
    })

    slog.Info("Starting server on :8080")
    r.Run(":8080")
}
```

---

## 11. Decisiones Arquitectónicas Clave

| Decisión | Justificación |
|----------|---------------|
| Monolito Go en lugar de microservicios | PoC requiere velocidad, no complejidad innecesaria |
| SQLite local + PostgreSQL prod | Desarrollo rápido, producción robusta |
| WebSocket para A2UI | Actualizaciones en tiempo real sin polling |
| Nuxt 3 separado del backend | Independencia de equipos, SSR/SSG opcional |
| IRT para cold-start | Estándar académico validado, sin sesgo demográfico |
| Algoritmo 85% para quiz | Basado en investigación de "desirable difficulty" |
| MinIO para assets | S3-compatible, desplegable en K8s |

---

*Documento generado para la PoC del Sistema de Estudio Adaptativo Multimodal*