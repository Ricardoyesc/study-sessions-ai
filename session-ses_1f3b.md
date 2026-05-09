# New session - 2026-05-09T10:20:41.402Z

**Session ID:** ses_1f3bf15c5ffeQ1GjxhCQsVghXG
**Created:** 5/9/2026, 4:20:41 AM
**Updated:** 5/9/2026, 4:53:22 AM

---

## User

Actúa como un Arquitecto de Software y Desarrollador Backend experto. Necesito una propuesta arquitectónica detallada y un plan de desarrollo para implementar una Prueba de Concepto (PoC) de un sistema de estudio adaptativo multimodal basado en IA.

Infraestructura y Stack Tecnológico Base:

Arquitectura: Sistema monolítico (dockerizado) para evitar complejidad innecesaria en la PoC.

Despliegue: Kubernetes.

Lenguaje de Desarrollo Principal: Go (Golang).

Requerimientos Arquitectónicos y Funcionales a Implementar en el Monolito:

1. Motor de Interfaz Declarativa y Accesibilidad (Basado en A2UI):

El frontend debe actuar como un cliente que renderice componentes nativos basándose en el protocolo A2UI.  

El backend monolítico en Go debe actuar como el agente que envía descripciones de componentes estructuradas en mensajes JSON, utilizando el Modelo de Lista de Adyacencia.  

La arquitectura debe permitir actualizaciones de estado dinámicas (updateDataModel) para modificar en tiempo real parámetros de accesibilidad sin recargar la aplicación: transiciones a tipografías sans-serif, activación de paletas de color pastel o modo oscuro nativo, y reestructuración responsiva de contenedores para evitar el recorte de textos.  

2. Orquestador de Codificación Dual y Generación Multimodal:

El backend debe centralizar y orquestar llamadas a Modelos Fundacionales y APIs externas (generación de texto, clonación de voz, generación de imágenes/video) para crear cápsulas de microaprendizaje.

Diseña el sistema para que aplique siempre la Teoría de Codificación Dual: la arquitectura de datos de cada cápsula debe asegurar que la información se entregue en al menos dos modalidades sincronizadas (por ejemplo, texto descriptivo devuelto vía A2UI junto con un archivo de audio generado de forma simultánea).  

3. Módulo de Diagnóstico de Inicio en Frío (Cold-Start):

Diseña un servicio interno que exponga una evaluación adaptativa inicial para nuevos usuarios.

La lógica de negocio debe utilizar modelos de trazado de conocimiento o agrupamiento (clustering) de interacciones iniciales para establecer el nivel de habilidad latente. Queda estrictamente prohibido diseñar bases de datos o flujos que intenten inferir el conocimiento inicial utilizando metadatos demográficos.  

4. Bucle de Evaluación Algorítmica (La Regla del 85%):

Implementa un motor de generación de cuestionarios que fuerce la recuperación activa.

El algoritmo de ensamblaje de la sesión de estudio debe integrar repetición espaciada y práctica intercalada, mezclando diferentes tipos de problemas y conceptos.  

La variable matemática central del algoritmo de selección de preguntas debe mezclar dinámicamente material no dominado y repasos de material dominado para que la probabilidad predictiva de éxito del estudiante en esa sesión converja exactamente en un 85%.  

5. Sistema de Remediación Socrática:

Diseña el flujo del LLM (mediante ingeniería de prompts en el backend) para que, cuando el estudiante falle, genere una intervención basada en la Técnica de Feynman utilizando analogías simples.  

Integra un bucle de Interrogación Elaborativa, donde el agente solicite al usuario que explique el porqué de un concepto antes de permitirle avanzar.  

Entregables esperados de tu parte:

Un diagrama o descripción de la arquitectura general (monolito dockerizado en Kubernetes, cliente A2UI, bases de datos y conexiones a APIs de IA).

Propuesta del diseño de la API en Go para manejar el flujo bidireccional de JSONs de A2UI.

Estructura de manifiestos básicos de despliegue para Kubernetes (Deployment, Service, etc., orientados a la imagen del monolito).

Definición de la estructura de carpetas del proyecto en Go y los primeros pasos de código para iniciar la PoC.

Detalla la implementación y desarrollo por etapas de tal forma que varios agentes puedan utilizar este plan como base y desarrollar cada etapa por separado 

---

## Assistant (Plan · DeepSeek V4 Pro · 1.1s)

---

## User

Actúa como un Arquitecto de Software y Desarrollador Backend experto. Necesito una propuesta arquitectónica detallada y un plan de desarrollo para implementar una Prueba de Concepto (PoC) de un sistema de estudio adaptativo multimodal basado en IA.

Infraestructura y Stack Tecnológico Base:

Arquitectura: Sistema monolítico (dockerizado) para evitar complejidad innecesaria en la PoC.

Despliegue: Kubernetes.

Lenguaje de Desarrollo Principal: Go (Golang).

Requerimientos Arquitectónicos y Funcionales a Implementar en el Monolito:

1. Motor de Interfaz Declarativa y Accesibilidad (Basado en A2UI):

El frontend debe actuar como un cliente que renderice componentes nativos basándose en el protocolo A2UI.  

El backend monolítico en Go debe actuar como el agente que envía descripciones de componentes estructuradas en mensajes JSON, utilizando el Modelo de Lista de Adyacencia.  

La arquitectura debe permitir actualizaciones de estado dinámicas (updateDataModel) para modificar en tiempo real parámetros de accesibilidad sin recargar la aplicación: transiciones a tipografías sans-serif, activación de paletas de color pastel o modo oscuro nativo, y reestructuración responsiva de contenedores para evitar el recorte de textos.  

2. Orquestador de Codificación Dual y Generación Multimodal:

El backend debe centralizar y orquestar llamadas a Modelos Fundacionales y APIs externas (generación de texto, clonación de voz, generación de imágenes/video) para crear cápsulas de microaprendizaje.

Diseña el sistema para que aplique siempre la Teoría de Codificación Dual: la arquitectura de datos de cada cápsula debe asegurar que la información se entregue en al menos dos modalidades sincronizadas (por ejemplo, texto descriptivo devuelto vía A2UI junto con un archivo de audio generado de forma simultánea).  

3. Módulo de Diagnóstico de Inicio en Frío (Cold-Start):

Diseña un servicio interno que exponga una evaluación adaptativa inicial para nuevos usuarios.

La lógica de negocio debe utilizar modelos de trazado de conocimiento o agrupamiento (clustering) de interacciones iniciales para establecer el nivel de habilidad latente. Queda estrictamente prohibido diseñar bases de datos o flujos que intenten inferir el conocimiento inicial utilizando metadatos demográficos.  

4. Bucle de Evaluación Algorítmica (La Regla del 85%):

Implementa un motor de generación de cuestionarios que fuerce la recuperación activa.

El algoritmo de ensamblaje de la sesión de estudio debe integrar repetición espaciada y práctica intercalada, mezclando diferentes tipos de problemas y conceptos.  

La variable matemática central del algoritmo de selección de preguntas debe mezclar dinámicamente material no dominado y repasos de material dominado para que la probabilidad predictiva de éxito del estudiante en esa sesión converja exactamente en un 85%.  

5. Sistema de Remediación Socrática:

Diseña el flujo del LLM (mediante ingeniería de prompts en el backend) para que, cuando el estudiante falle, genere una intervención basada en la Técnica de Feynman utilizando analogías simples.  

Integra un bucle de Interrogación Elaborativa, donde el agente solicite al usuario que explique el porqué de un concepto antes de permitirle avanzar.  

Entregables esperados de tu parte:

Un diagrama o descripción de la arquitectura general (monolito dockerizado en Kubernetes, cliente A2UI, bases de datos y conexiones a APIs de IA).

Propuesta del diseño de la API en Go para manejar el flujo bidireccional de JSONs de A2UI.

Estructura de manifiestos básicos de despliegue para Kubernetes (Deployment, Service, etc., orientados a la imagen del monolito).

Definición de la estructura de carpetas del proyecto en Go y los primeros pasos de código para iniciar la PoC.

Detalla la implementación y desarrollo por etapas de tal forma que varios agentes puedan utilizar este plan como base y desarrollar cada etapa por separado 

---

## Assistant (Plan · DeepSeek V4 Pro · 6.9s)

**Tool: read**

**Input:**
```json
{
  "filePath": "/home/davidyesc/projects/study-sessions-with-ai"
}
```

**Output:**
```
<path>/home/davidyesc/projects/study-sessions-with-ai</path>
<type>directory</type>
<entries>


(0 entries)
</entries>
```

---

## Assistant (Plan · DeepSeek V4 Pro · 464.9s)

# Arquitectura del Sistema de Estudio Adaptativo Multimodal (PoC)

A continuación presento la propuesta completa, organizada en entregables. Como estoy en modo solo-lectura (Plan Mode), no crearé archivos — presento el diseño detallado para revisión y aprobación antes de la implementación.

---

## 1. Diagrama/Descripción de Arquitectura General

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
│                   │  │  ┌──────────┐  ┌──────────┐           │    │   │
│                   │  │  │ Session  │  │  User    │           │    │   │
│                   │  │  │ Manager  │  │ Service  │           │    │   │
│                   │  │  └──────────┘  └──────────┘           │    │   │
│                   │  └───────────────────────────────────────┘    │   │
│                   │  ┌───────────────────────────────────────┐    │   │
│                   │  │        Módulos de Dominio              │    │   │
│                   │  │  ┌────────────┐  ┌────────────────┐   │    │   │
│                   │  │  │ Dual-Code  │  │  Quiz Engine   │   │    │   │
│                   │  │  │Orchestrator│  │  (85% Rule)    │   │    │   │
│                   │  │  └────────────┘  └────────────────┘   │    │   │
│                   │  │  ┌────────────┐  ┌────────────────┐   │    │   │
│                   │  │  │ Cold-Start │  │   Socratic     │   │    │   │
│                   │  │  │ Diagnostic │  │  Remediation   │   │    │   │
│                   │  │  └────────────┘  └────────────────┘   │    │   │
│                   │  └───────────────────────────────────────┘    │   │
│                   │  ┌───────────────────────────────────────┐    │   │
│                   │  │        Ports / Adapters                │    │   │
│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │
│                   │  │  │ LLM  │ │ TTS  │ │ Image/Video  │  │    │   │
│                   │  │  │Client│ │Client│ │   Client     │  │    │   │
│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │
│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │
│                   │  │  │  DB  │ │Redis │ │   MinIO/S3   │  │    │   │
│                   │  │  │(PG)  │ │(Pub) │ │  (Assets)    │  │    │   │
│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │
│                   │  └───────────────────────────────────────┘    │   │
│                   └──────────────────────────────────────────────────┘   │
│                                                                           │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐         │
│  │   PostgreSQL 14  │  │    Redis 7       │  │   MinIO (S3)    │         │
│  │  (StatefulSet)   │  │  (StatefulSet)   │  │ (StatefulSet)   │         │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘         │
└──────────────────────────────────────────────────────────────────────────┘

                    ┌──────────────────────────────────────┐
                    │          APIs Externas                │
                    │  ┌────────┐  ┌──────────┐  ┌───────┐│
                    │  │ OpenAI │  │ElevenLabs│  │  DALL·E││
                    │  │Claude  │  │  (TTS)   │  │ /SD   ││
                    │  └────────┘  └──────────┘  └───────┘│
                    └──────────────────────────────────────┘

  ┌─────────────────────────────────────────────────────────────────┐
  │                   Cliente A2UI (Web/Mobile)                       │
  │  ┌─────────────────────────────────────────────────────────────┐ │
  │  │  A2UI Renderer Engine                                        │ │
  │  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────┐  │ │
  │  │  │Component │ │ Layout   │ │Theme/    │ │ Event        │  │ │
  │  │  │Registry  │ │ Resolver │ │A11y Mgr  │ │ Dispatcher   │  │ │
  │  │  └──────────┘ └──────────┘ └──────────┘ └──────────────┘  │ │
  │  └─────────────────────────────────────────────────────────────┘ │
  │  Native Components: Column, Row, Card, RichText, Image,         │
  │  AudioPlayer, VideoPlayer, QuizCard, SocraticDialog, ...        │
  └─────────────────────────────────────────────────────────────────┘
```

**Flujo de datos:**
- Cliente → `POST /api/a2ui/event` → Backend procesa acción → `POST /session/{id}/update` (via WS) → Cliente renderiza nuevo árbol A2UI
- Las `updateDataModel` se transmiten por WebSocket para cambios de accesibilidad en tiempo real sin recarga

---

## 2. Diseño de API REST + WebSocket (Go/Gin)

### Endpoints REST

| Método | Ruta | Descripción |
|--------|------|-------------|
| `POST` | `/api/users/register` | Registro de usuario (sin demográficos) |
| `GET`  | `/api/users/{id}/profile` | Perfil de conocimiento |
| `POST` | `/api/coldstart/start` | Inicia diagnóstico cold-start |
| `POST` | `/api/coldstart/answer` | Envía respuesta a ítem de diagnóstico |
| `GET`  | `/api/coldstart/{id}/result` | Resultado estimado de habilidad |
| `POST` | `/api/sessions` | Crea sesión de estudio |
| `GET`  | `/api/sessions/{id}` | Estado de sesión (datos A2UI) |
| `POST` | `/api/sessions/{id}/quiz/answer` | Envía respuesta a quiz |
| `POST` | `/api/sessions/{id}/socratic/response` | Envía respuesta elaborativa |
| `POST` | `/api/capsules/generate` | Solicita generación de cápsula dual |
| `GET`  | `/api/capsules/{id}` | Obtiene cápsula (A2UI + assets) |
| `POST` | `/api/sessions/{id}/a11y` | Actualiza preferencias accesibilidad |
| `GET`  | `/api/assets/{type}/{filename}` | Archivos estáticos (audio/imagen/video) |

### WebSocket `/ws/session/{id}`

Mensajes bidireccionales JSON con envelope:

```json
{
  "type": "a2ui_update | a2ui_full | data_model_update | error | ping",
  "payload": {},
  "timestamp": "ISO8601"
}
```

#### Tipos de mensaje A2UI

**a2ui_full** — renderizado inicial completo:
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
        "props": { "markdown": "La dualidad onda-partícula es...", "accessible": true }
      },
      "audio-player": {
        "id": "audio-player",
        "type": "AudioPlayer",
        "props": { "url": "/api/assets/audio/capsule-123.mp3", "autoPlay": false }
      },
      "image-diagram": {
        "id": "image-diagram",
        "type": "Image",
        "props": { "url": "/api/assets/image/diagram-123.png", "altText": "Diagrama del experimento de doble rendija..." }
      },
      "quiz-card": {
        "id": "quiz-card",
        "type": "QuizCard",
        "props": {
          "question": "¿Por qué el patrón de interferencia desaparece al observar las partículas?",
          "options": ["...", "..."],
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
      "reducedMotion": false
    }
  }
}
```

**data_model_update** — cambio de accesibilidad en caliente:
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

**a2ui_update** — parche parcial del árbol:
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

## 3. Manifiestos Kubernetes

### `deploy/k8s/namespace.yaml`
```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: sai-learning
```

### `deploy/k8s/configmap.yaml`
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
  TTS_PROVIDER: "elevenlabs"
```

### `deploy/k8s/secrets.yaml`
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
  ELEVENLABS_API_KEY: "el-..."
  STABILITY_API_KEY: "sk-..."
  JWT_SECRET: "change-me-jwt-secret"
  MINIO_ACCESS_KEY: "minioadmin"
  MINIO_SECRET_KEY: "minioadmin"
```

### `deploy/k8s/deployment-postgres.yaml`
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: postgres-pvc
  namespace: sai-learning
spec:
  accessModes: [ReadWriteOnce]
  resources:
    requests:
      storage: 10Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: postgres
  namespace: sai-learning
spec:
  replicas: 1
  selector:
    matchLabels:
      app: postgres
  template:
    metadata:
      labels:
        app: postgres
    spec:
      containers:
      - name: postgres
        image: postgres:14-alpine
        ports:
        - containerPort: 5432
        envFrom:
        - secretRef:
            name: sai-server-secrets
        env:
        - name: POSTGRES_DB
          valueFrom:
            configMapKeyRef:
              name: sai-server-config
              key: DB_NAME
        volumeMounts:
        - name: postgres-data
          mountPath: /var/lib/postgresql/data
      volumes:
      - name: postgres-data
        persistentVolumeClaim:
          claimName: postgres-pvc
---
apiVersion: v1
kind: Service
metadata:
  name: postgres-service
  namespace: sai-learning
spec:
  selector:
    app: postgres
  ports:
  - port: 5432
    targetPort: 5432
```

### `deploy/k8s/deployment-redis.yaml`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
  namespace: sai-learning
spec:
  replicas: 1
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - name: redis
        image: redis:7-alpine
        command: ["redis-server", "--appendonly", "yes"]
        ports:
        - containerPort: 6379
---
apiVersion: v1
kind: Service
metadata:
  name: redis-service
  namespace: sai-learning
spec:
  selector:
    app: redis
  ports:
  - port: 6379
    targetPort: 6379
```

### `deploy/k8s/deployment-minio.yaml`
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: minio-pvc
  namespace: sai-learning
spec:
  accessModes: [ReadWriteOnce]
  resources:
    requests:
      storage: 20Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: minio
  namespace: sai-learning
spec:
  replicas: 1
  selector:
    matchLabels:
      app: minio
  template:
    metadata:
      labels:
        app: minio
    spec:
      containers:
      - name: minio
        image: minio/minio:latest
        args: ["server", "/data"]
        ports:
        - containerPort: 9000
        - containerPort: 9001
        env:
        - name: MINIO_ROOT_USER
          valueFrom:
            secretKeyRef:
              name: sai-server-secrets
              key: MINIO_ACCESS_KEY
        - name: MINIO_ROOT_PASSWORD
          valueFrom:
            secretKeyRef:
              name: sai-server-secrets
              key: MINIO_SECRET_KEY
        volumeMounts:
        - name: minio-data
          mountPath: /data
      volumes:
      - name: minio-data
        persistentVolumeClaim:
          claimName: minio-pvc
---
apiVersion: v1
kind: Service
metadata:
  name: minio-service
  namespace: sai-learning
spec:
  selector:
    app: minio
  ports:
  - name: api
    port: 9000
    targetPort: 9000
```

### `deploy/k8s/deployment-server.yaml`
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
        env:
        - name: PORT
          value: "8080"
        - name: METRICS_PORT
          value: "8081"
        livenessProbe:
          httpGet:
            path: /healthz
            port: 8080
          initialDelaySeconds: 15
          periodSeconds: 20
        readinessProbe:
          httpGet:
            path: /readyz
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 10
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
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
  - name: ws
    port: 8081
    targetPort: 8081
```

### `deploy/k8s/ingress.yaml`
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

### `deploy/docker/Dockerfile`
```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /sai-server ./cmd/server

FROM alpine:3.19
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /sai-server /usr/local/bin/sai-server
COPY configs/ /etc/sai/
EXPOSE 8080 8081
ENTRYPOINT ["/usr/local/bin/sai-server"]
```

---

## 4. Estructura de Carpetas del Proyecto Go

```
study-sessions-with-ai/
├── cmd/
│   └── server/
│       └── main.go                    # Entry point, wire-up de dependencias
├── internal/
│   ├── api/
│   │   ├── router.go                  # Registro de rutas Gin
│   │   ├── middleware/
│   │   │   ├── auth.go                # JWT middleware
│   │   │   ├── logging.go
│   │   │   └── cors.go
│   │   └── handlers/
│   │       ├── a2ui_handler.go        # WebSocket + eventos A2UI
│   │       ├── coldstart_handler.go  # Endpoints de diagnóstico
│   │       ├── session_handler.go    # Endpoints de sesión
│   │       ├── capsule_handler.go    # Endpoints de cápsulas
│   │       ├── quiz_handler.go       # Endpoints de quiz
│   │       ├── socratic_handler.go   # Endpoints socráticos
│   │       └── user_handler.go       # Endpoints de usuario
│   ├── app/
│   │   ├── session/
│   │   │   ├── service.go            # Lógica de sesión de estudio
│   │   │   └── types.go
│   │   ├── coldstart/
│   │   │   ├── service.go            # Diagnóstico cold-start
│   │   │   ├── irt.go                # Item Response Theory
│   │   │   └── types.go
│   │   ├── dualcode/
│   │   │   ├── orchestrator.go       # Orquestador codificación dual
│   │   │   ├── capsule.go            # Generación de cápsulas
│   │   │   └── types.go
│   │   ├── quiz/
│   │   │   ├── engine.go             # Motor 85% Rule
│   │   │   ├── selector.go           # Algoritmo de selección
│   │   │   ├── evaluator.go          # Evaluación de respuestas
│   │   │   └── types.go
│   │   ├── socratic/
│   │   │   ├── remediator.go         # Remediación socrática
│   │   │   ├── prompts.go            # Plantillas de prompt
│   │   │   └── types.go
│   │   └── user/
│   │       ├── service.go
│   │       └── types.go
│   ├── domain/
│   │   ├── user.go
│   │   ├── session.go
│   │   ├── capsule.go
│   │   ├── quiz.go
│   │   ├── concept.go                # Grafo de conocimientos
│   │   └── a2ui.go                   # Tipos core A2UI
│   ├── port/
│   │   ├── llm.go                    # Interfaz LLM client
│   │   ├── tts.go                    # Interfaz TTS client
│   │   ├── imagegen.go               # Interfaz generación imágenes
│   │   ├── storage.go                # Interfaz almacenamiento
│   │   └── cache.go                  # Interfaz cache
│   ├── adapter/
│   │   ├── llm/
│   │   │   ├── openai.go             # OpenAI/Claude adapter
│   │   │   └── factory.go
│   │   ├── tts/
│   │   │   ├── elevenlabs.go
│   │   │   └── factory.go
│   │   ├── imagegen/
│   │   │   ├── dalle.go
│   │   │   ├── stability.go
│   │   │   └── factory.go
│   │   ├── storage/
│   │   │   └── minio.go
│   │   └── cache/
│   │       └── redis.go
│   └── config/
│       └── config.go                 # Carga de configuración
├── pkg/
│   ├── a2ui/
│   │   ├── protocol.go               # Definición protocolo A2UI
│   │   ├── builder.go                # Component tree builder
│   │   ├── diff.go                   # A2UI diff/patcher
│   │   └── types.go                  # Tipos de componentes
│   ├── spacedrep/
│   │   ├── sm2.go                    # SM-2 spaced repetition
│   │   └── scheduler.go
│   ├── irt/
│   │   ├── model.go                  # Modelos IRT
│   │   └── estimator.go              # Estimador de habilidad
│   ├── bkt/
│   │   └── bayesian.go               # Bayesian Knowledge Tracing
│   └── adaptive/
│       └── difficulty.go             # Calibración de dificultad
├── web/
│   └── client/                       # Frontend A2UI (en otro repo o submódulo)
├── migrations/
│   ├── 001_create_users.sql
│   ├── 002_create_concepts.sql
│   ├── 003_create_sessions.sql
│   └── ...
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
│   ├── migrate.sh
│   ├── seed_data.sh
│   └── dev.sh
├── go.mod
├── go.sum
├── Makefile
├── docker-compose.yml                 # Para desarrollo local
└── .env.example
```

---

## 5. Plan de Implementación por Etapas (Agentes Paralelos)

### Fase 0 — Fundación (Agente A)

| Tarea | Descripción | Archivos |
|-------|-------------|----------|
| 0.1 | `go mod init`, estructura de carpetas, `Makefile` | `go.mod`, `Makefile` |
| 0.2 | Sistema de configuración (`viper`) | `internal/config/config.go`, `configs/app.yaml` |
| 0.3 | Tipos de dominio (`User`, `Session`, `Concept`, `Capsule`, `Quiz`) | `internal/domain/*.go` |
| 0.4 | Puerto de interfaces (`llm`, `tts`, `imagegen`, `storage`, `cache`) | `internal/port/*.go` |
| 0.5 | Protocolo A2UI — tipos core (`Surface`, `Component`, `DataModel`) | `pkg/a2ui/types.go`, `pkg/a2ui/protocol.go` |
| 0.6 | Router Gin + middleware (auth, logging, CORS) | `internal/api/router.go`, `internal/api/middleware/*.go` |
| 0.7 | `main.go` — wire-up manual de dependencias | `cmd/server/main.go` |
| 0.8 | Dockerfile + docker-compose.yml | `deploy/docker/Dockerfile`, `docker-compose.yml` |

### Fase 1 — A2UI Engine + Core (Agente B + Agente C en paralelo)

**Agente B — A2UI Engine:**
| Tarea | Descripción |
|-------|-------------|
| 1.B.1 | A2UI `Builder`: construye árboles de componentes desde cápsulas/quizzes | `pkg/a2ui/builder.go` |
| 1.B.2 | A2UI `Diff`: calcula parches incrementales entre estados | `pkg/a2ui/diff.go` |
| 1.B.3 | WebSocket manager + A2UI handler | `internal/api/handlers/a2ui_handler.go` |
| 1.B.4 | `data_model_update` — propaga cambios de accesibilidad en caliente | en a2ui_handler.go |

**Agente C — Adapters + User Service:**
| Tarea | Descripción |
|-------|-------------|
| 1.C.1 | Adapter OpenAI (LLM) + factory | `internal/adapter/llm/openai.go` |
| 1.C.2 | Adapter ElevenLabs (TTS) + factory | `internal/adapter/tts/elevenlabs.go` |
| 1.C.3 | Adapter DALL·E/Stability (Image) + factory | `internal/adapter/imagegen/*.go` |
| 1.C.4 | Adapter MinIO (Storage) | `internal/adapter/storage/minio.go` |
| 1.C.5 | Adapter Redis (Cache) | `internal/adapter/cache/redis.go` |
| 1.C.6 | User service + handler + migración DB | `internal/app/user/`, `internal/api/handlers/user_handler.go` |

### Fase 2 — Módulos de Aprendizaje (Agentes D, E, F en paralelo)

**Agente D — Cold-Start Diagnostic:**
| Tarea | Descripción |
|-------|-------------|
| 2.D.1 | Implementar IRT models (1PL, 2PL, 3PL) | `pkg/irt/model.go` |
| 2.D.2 | Estimador de habilidad (MLE/EAP) | `pkg/irt/estimator.go` |
| 2.D.3 | Algoritmo CAT (Computerized Adaptive Testing) con Fisher Information | `internal/app/coldstart/service.go` |
| 2.D.4 | K-means clustering sobre respuestas iniciales | `internal/app/coldstart/cluster.go` |
| 2.D.5 | Handler REST para cold-start | `internal/api/handlers/coldstart_handler.go` |

**Agente E — Quiz Engine (85% Rule):**
| Tarea | Descripción |
|-------|-------------|
| 2.E.1 | BKT (Bayesian Knowledge Tracing) por concepto | `pkg/bkt/bayesian.go` |
| 2.E.2 | Algoritmo SM-2 extendido para repetición espaciada | `pkg/spacedrep/sm2.go` |
| 2.E.3 | Selector de preguntas que converge al 85% | `internal/app/quiz/selector.go` |
| 2.E.4 | Evaluador de respuestas (correct/incorrect + confianza) | `internal/app/quiz/evaluator.go` |
| 2.E.5 | Engine principal que orquesta selección + evaluación | `internal/app/quiz/engine.go` |
| 2.E.6 | Quiz handler REST | `internal/api/handlers/quiz_handler.go` |

**Agente F — Dual Coding Orchestrator:**
| Tarea | Descripción |
|-------|-------------|
| 2.F.1 | Generador de contenido textual vía LLM | `internal/app/dualcode/capsule.go` |
| 2.F.2 | Generación paralela de audio (TTS) + imagen/video | `internal/app/dualcode/orchestrator.go` |
| 2.F.3 | Ensamblador de cápsula dual (texto + audio + imagen → A2UI tree) | en orchestrator.go |
| 2.F.4 | Validación Codificación Dual: al menos 2 modalidades | en orchestrator.go |
| 2.F.5 | Capsule handler REST | `internal/api/handlers/capsule_handler.go` |

### Fase 3 — Session + Remediación Socrática (Agente G + Agente H en paralelo)

**Agente G — Session Manager:**
| Tarea | Descripción |
|-------|-------------|
| 3.G.1 | Servicio de sesión: crea, avanza, finaliza | `internal/app/session/service.go` |
| 3.G.2 | Pipeline de sesión: cold-start → cápsulas → quizzes → remediación | `internal/app/session/pipeline.go` |
| 3.G.3 | Session handler REST | `internal/api/handlers/session_handler.go` |

**Agente H — Socratic Remediation:**
| Tarea | Descripción |
|-------|-------------|
| 3.H.1 | Plantillas de prompt para Feynman Technique | `internal/app/socratic/prompts.go` |
| 3.H.2 | Remediador: clasifica error, genera intervención Feynman | `internal/app/socratic/remediator.go` |
| 3.H.3 | Bucle de interrogación elaborativa | en remediator.go |
| 3.H.4 | Socratic handler REST | `internal/api/handlers/socratic_handler.go` |

### Fase 4 — Integración, K8s, Pruebas (Agente I + Agente J)

| Agente | Tarea |
|--------|-------|
| I | Desplegar todos los manifiestos K8s, probar en minikube/kind |
| I | Configurar Ingress + TLS |
| I | Health checks + métricas Prometheus |
| J | Tests de integración (Go test + testify) |
| J | Documentación de API (OpenAPI/Swagger) |
| J | `README.md` final + instrucciones de despliegue |

---

## 6. Modelos de Datos Clave (PostgreSQL)

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
    estimated_theta FLOAT,              -- IRT ability estimate
    theta_uncertainty FLOAT,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- BKT por usuario-concepto
CREATE TABLE user_concept_mastery (
    user_id UUID REFERENCES users(id),
    concept_id UUID REFERENCES concepts(id),
    p_learned FLOAT NOT NULL DEFAULT 0.3,
    p_guess FLOAT NOT NULL DEFAULT 0.1,
    p_slip FLOAT NOT NULL DEFAULT 0.1,
    p_transit FLOAT NOT NULL DEFAULT 0.2,
    last_practiced TIMESTAMPTZ,
    easiness_factor FLOAT DEFAULT 2.5,   -- SM-2
    interval_days INT DEFAULT 1,         -- SM-2
    repetitions INT DEFAULT 0,           -- SM-2
    next_review TIMESTAMPTZ,
    PRIMARY KEY (user_id, concept_id)
);

-- Items de quiz
CREATE TABLE quiz_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    concept_id UUID REFERENCES concepts(id),
    difficulty_irt FLOAT NOT NULL,
    discrimination FLOAT DEFAULT 1.0,
    guessing FLOAT DEFAULT 0.25,
    content JSONB NOT NULL,              -- pregunta + opciones en JSON
    modality TEXT DEFAULT 'text',
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Sesiones
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
    modalities JSONB NOT NULL,          -- [{type, url, metadata}]
    a2ui_tree JSONB NOT NULL,
    session_id UUID REFERENCES sessions(id),
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Interacciones (respuestas a quizzes + respuestas socráticas)
CREATE TABLE interactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id UUID REFERENCES sessions(id),
    type TEXT NOT NULL,                  -- 'quiz_answer' | 'socratic_response'
    payload JSONB NOT NULL,
    was_correct BOOLEAN,
    response_time_ms INT,
    remediation_generated JSONB,
    created_at TIMESTAMPTZ DEFAULT now()
);
```

---

## 7. Preguntas para el Usuario

Antes de comenzar la implementación, necesito confirmar algunos puntos:

**Tool: question**

**Input:**
```json
{
  "questions": [
    {
      "question": "¿Qué API de LLM prefieres usar como principal? Esto determina el adapter inicial.",
      "header": "LLM Provider",
      "options": [
        {
          "label": "OpenAI (GPT-4o)",
          "description": "Usar OpenAI como proveedor principal, con soporte para GPT-4o y futuros modelos"
        },
        {
          "label": "Anthropic (Claude)",
          "description": "Usar Anthropic Claude como proveedor principal"
        },
        {
          "label": "Ambas (Multiprovider)",
          "description": "Implementar factory con soporte para ambas desde el inicio"
        }
      ]
    },
    {
      "question": "¿Qué motor de TTS prefieres para la generación de audio en las cápsulas duales?",
      "header": "TTS Provider",
      "options": [
        {
          "label": "ElevenLabs",
          "description": "ElevenLabs API para clonación de voz y TTS de alta calidad"
        },
        {
          "label": "OpenAI TTS",
          "description": "OpenAI Text-to-Speech API"
        },
        {
          "label": "Ambas",
          "description": "Soporte para ambas, configurable"
        }
      ]
    },
    {
      "question": "¿Para el frontend A2UI, qué enfoque prefieres?",
      "header": "Frontend Approach",
      "options": [
        {
          "label": "Incluido en este repo (SPA)",
          "description": "Single-page app servida desde el mismo monolito Go (embedded)"
        },
        {
          "label": "Repo separado",
          "description": "Frontend en repositorio independiente, solo backend aquí"
        },
        {
          "label": "MVP con templates HTML",
          "description": "Empezar con server-side templates + HTMX/Alpine para la PoC y migrar a SPA después"
        }
      ]
    },
    {
      "question": "¿Quieres que la PoC use PostgreSQL desde el inicio o es aceptable SQLite para simplificar desarrollo local?",
      "header": "Database Choice for Dev",
      "options": [
        {
          "label": "PostgreSQL (prod-like)",
          "description": "PostgreSQL desde el día 1, igual que en producción"
        },
        {
          "label": "SQLite local + PG prod",
          "description": "SQLite para desarrollo local rápido, PostgreSQL en K8s"
        }
      ]
    }
  ]
}
```

**Output:**
```
User has answered your questions: "¿Qué API de LLM prefieres usar como principal? Esto determina el adapter inicial."="OpenAI (GPT-4o)", "¿Qué motor de TTS prefieres para la generación de audio en las cápsulas duales?"="Gemini TTS", "¿Para el frontend A2UI, qué enfoque prefieres?"="Repo separado con Nuxt o equivalente. Pero la definición del front y su implementación debe estar cubierta en esta documentación.", "¿Quieres que la PoC use PostgreSQL desde el inicio o es aceptable SQLite para simplificar desarrollo local?"="SQLite local + PG prod". You can now continue with the user's answers in mind.
```

---

## Assistant (Plan · DeepSeek V4 Pro · 284.8s)

Buenas decisiones. Actualizo el plan con las respuestas:

- **LLM**: OpenAI GPT-4o (adapter principal)
- **TTS**: Gemini TTS (Google Cloud Text-to-Speech)
- **Frontend**: Repositorio separado con Nuxt/Vue 3, documentado aquí en su totalidad
- **DB**: SQLite para desarrollo local, PostgreSQL en producción/K8s

A continuación completo todos los entregables finales, incluyendo la arquitectura del frontend Nuxt + A2UI Renderer.

---

# Documento Final de Arquitectura y Plan de Desarrollo

## 1. Arquitectura General (actualizada con Gemini TTS + Nuxt)

```
┌──────────────────────────────────────────────────────────────────────────┐
│                         CLUSTER KUBERNETES                                │
│                                                                          │
│  ┌─────────────────────────────────────────────────────────────────┐    │
│  │                  Monolito Go (sai-server)                        │    │
│  │                                                                  │    │
│  │  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌───────────────┐  │    │
│  │  │  Gin     │  │  WS Hub  │  │  A2UI    │  │  Middleware   │  │    │
│  │  │ Router   │  │ (gorilla │  │  Engine  │  │ (Auth/Log/    │  │    │
│  │  │ REST     │  │  /ws)    │  │  Builder │  │  CORS)        │  │    │
│  │  └────┬─────┘  └────┬─────┘  └────┬─────┘  └───────────────┘  │    │
│  │       └──────────────┴────────────┘                             │    │
│  │                         │                                        │    │
│  │  ┌──────────────────────┴──────────────────────────────────┐   │    │
│  │  │                 Service Layer                             │   │    │
│  │  │  ┌─────────┐ ┌───────────┐ ┌───────────┐ ┌───────────┐  │   │    │
│  │  │  │ Session │ │ ColdStart │ │ DualCode  │ │ Quiz      │  │   │    │
│  │  │  │ Service │ │ Service   │ │ Orchestr. │ │ Engine    │  │   │    │
│  │  │  └─────────┘ └───────────┘ └───────────┘ └───────────┘  │   │    │
│  │  │  ┌───────────┐ ┌───────────┐                              │   │    │
│  │  │  │ Socratic  │ │  User     │                              │   │    │
│  │  │  │ Remediat. │ │  Service  │                              │   │    │
│  │  │  └───────────┘ └───────────┘                              │   │    │
│  │  └───────────────────────────────────────────────────────────┘   │    │
│  │  ┌───────────────────────────────────────────────────────────┐   │    │
│  │  │                   Adapter Layer (Ports)                     │   │    │
│  │  │  ┌────────┐ ┌──────────┐ ┌────────┐ ┌───────┐ ┌────────┐  │   │    │
│  │  │  │OpenAI  │ │GeminiTTS │ │DALL·E  │ │MinIO  │ │ SQL/   │  │   │    │
│  │  │  │GPT-4o  │ │(Google)  │ │/SD     │ │Storage│ │ GORM   │  │   │    │
│  │  │  └────────┘ └──────────┘ └────────┘ └───────┘ └────────┘  │   │    │
│  │  └───────────────────────────────────────────────────────────┘   │    │
│  └─────────────────────────────────────────────────────────────────┘    │
│                                                                          │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐                      │
│  │ PostgreSQL   │ │   Redis 7    │ │ MinIO (S3)  │                      │
│  │ (StatefulSet)│ │ (StatefulSet)│ │(StatefulSet)│                      │
│  └──────────────┘ └──────────────┘ └──────────────┘                      │
└──────────────────────────────────────────────────────────────────────────┘

APIs Externas:
┌─────────────────────────────────────────────────────────────┐
│  OpenAI GPT-4o  │  Gemini TTS  │  DALL·E / Stability AI     │
└─────────────────────────────────────────────────────────────┘

Frontend (Repositorio Separado):
┌─────────────────────────────────────────────────────────────┐
│              Nuxt 3 SPA — Cliente A2UI                       │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  A2UI Renderer (Vue 3 composables)                    │   │
│  │  ┌────────────┐ ┌──────────┐ ┌──────────────────┐   │   │
│  │  │ Component  │ │ Adjacency│ │ Accessibility    │   │   │
│  │  │ Registry   │ │List Parser│ │ Engine (a11y)   │   │   │
│  │  │(dynamic)   │ │(JSON→VDOM)│ │(dataModel patch)│   │   │
│  │  └────────────┘ └──────────┘ └──────────────────┘   │   │
│  │  ┌────────────┐ ┌──────────┐                         │   │
│  │  │ WebSocket  │ │ State    │                         │   │
│  │  │ Manager    │ │ Store    │                         │   │
│  │  └────────────┘ └──────────┘                         │   │
│  └──────────────────────────────────────────────────────┘   │
│  Componentes Nativos: Text, RichText, Image, AudioPlayer,  │
│  VideoPlayer, Card, QuizCard, SocraticDialog, ProgressBar  │
└─────────────────────────────────────────────────────────────┘
```

---

## 2. Diseño de API en Go (completo)

*(Se mantiene el diseño anterior, actualizado con Gemini TTS)*

### Estructura de Dependencias en Go (`main.go`)

```go
// cmd/server/main.go
package main

import (
    "log/slog"
    "os"

    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"  // o postgres en prod
    "gorm.io/gorm"
)

func main() {
    cfg := config.Load()

    // --- DB (SQLite local / PG prod) ---
    var dialector gorm.Dialector
    if cfg.IsProduction() {
        dialector = postgres.Open(cfg.DatabaseURL)
    } else {
        dialector = sqlite.Open("data/sai_dev.db")
    }
    db, _ := gorm.Open(dialector, &gorm.Config{})
    db.AutoMigrate(&domain.User{}, &domain.Concept{}, /* ... */)

    // --- Adapters ---
    llmClient    := llm_adapter.NewOpenAI(cfg.OpenAIKey, cfg.OpenAIModel)
    ttsClient    := tts_adapter.NewGeminiTTS(cfg.GeminiAPIKey)
    imageClient  := image_adapter.NewDalle(cfg.OpenAIKey)
    storage      := storage_adapter.NewMinIO(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)
    cache        := cache_adapter.NewRedis(cfg.RedisURL)

    // --- Services ---
    a2uiEngine   := a2ui.NewEngine()
    userSvc      := user.NewService(db)
    coldStartSvc := coldstart.NewService(db, cache)
    quizEngine   := quiz.NewEngine(db, coldStartSvc)
    dualCodeOrch := dualcode.NewOrchestrator(llmClient, ttsClient, imageClient, storage, a2uiEngine)
    socraticRem  := socratic.NewRemediator(llmClient)
    sessionSvc   := session.NewService(db, coldStartSvc, dualCodeOrch, quizEngine, socraticRem)

    // --- Router ---
    r := gin.Default()
    r.Use(middleware.CORS(), middleware.Logger())

    api.RegisterRoutes(r, &api.Handlers{
        ColdStart: coldstart.NewHandler(coldStartSvc),
        Session:   session.NewHandler(sessionSvc),
        Capsule:   capsule.NewHandler(dualCodeOrch),
        Quiz:      quiz.NewHandler(quizEngine),
        Socratic:  socratic.NewHandler(socraticRem),
        User:      user.NewHandler(userSvc),
        A2UI:      a2ui.NewWSHandler(sessionSvc, a2uiEngine),
    })

    // --- Start ---
    r.Run(":8080")
}
```

### Formato Completo del Protocolo A2UI (pkg/a2ui/protocol.go)

```go
package a2ui

// Surface representa una "pantalla" completa de A2UI
type Surface struct {
    SurfaceID     string                `json:"surfaceId"`
    RootComponent string                `json:"rootComponent"`
    Components    map[string]Component  `json:"components"`  // Adjacency List
    DataModel     DataModel             `json:"dataModel"`
}

// Component es un nodo del árbol declarativo
type Component struct {
    ID       string                 `json:"id"`
    Type     string                 `json:"type"`       // "Text", "Card", "QuizCard", etc.
    Children []string               `json:"children,omitempty"` // IDs de hijos
    Props    map[string]interface{} `json:"props"`
    Events   map[string]string      `json:"events,omitempty"` // eventName -> endpoint
}

// DataModel contiene estado global de accesibilidad/tema
type DataModel struct {
    Theme         string  `json:"theme"`          // "light" | "dark" | "system"
    FontFamily    string  `json:"fontFamily"`     // "sans-serif" | "serif" | "OpenDyslexic"
    FontScale     float64 `json:"fontScale"`      // 0.8 - 2.0
    ColorPalette  string  `json:"colorPalette"`   // "default" | "pastel" | "high-contrast"
    HighContrast  bool    `json:"highContrast"`
    ReducedMotion bool    `json:"reducedMotion"`
    Language      string  `json:"language"`       // "es" | "en"
}

// WSMessage es el envelope de WebSocket
type WSMessage struct {
    Type      string      `json:"type"`      // a2ui_full | a2ui_update | data_model_update | error
    Payload   interface{} `json:"payload"`
    Timestamp string      `json:"timestamp"`
}

// DataModelUpdate para cambios de accesibilidad en caliente
type DataModelUpdate struct {
    Path  string                 `json:"path"`  // ej: "fontFamily"
    Value interface{}            `json:"value"`
    Diff  map[string]interface{} `json:"diff"`  // cambios completos a aplicar
}

// A2UIUpdate para parches parciales del árbol
type A2UIUpdate struct {
    Updates []ComponentPatch `json:"updates"`
}

type ComponentPatch struct {
    ComponentID string                 `json:"componentId"`
    Props       map[string]interface{} `json:"props,omitempty"`
    Children    []string               `json:"children,omitempty"`
    Events      map[string]string      `json:"events,omitempty"`
}
```

### WSMessageType enum-like constants

```go
const (
    MsgTypeA2UIFull       = "a2ui_full"
    MsgTypeA2UIUpdate     = "a2ui_update"
    MsgTypeDataModelUpdate = "data_model_update"
    MsgTypeError          = "error"
    MsgTypePing           = "ping"
    MsgTypeClientEvent    = "client_event"   // Cliente → Servidor
)
```

---

## 3. Arquitectura del Frontend (Nuxt 3 — Repositorio Separado)

### Estructura de carpetas del frontend

```
sai-frontend/
├── nuxt.config.ts
├── package.json
├── tsconfig.json
├── pages/
│   ├── index.vue                    # Landing / Login
│   ├── coldstart.vue                # Diagnóstico inicial
│   └── session/
│       └── [id].vue                 # Sesión de estudio activa
├── composables/
│   ├── useA2UI.ts                   # Core: WebSocket + Surface state
│   ├── useA2UIAccessibility.ts      # DataModel observer + hot-reload a11y
│   ├── useWebSocket.ts              # WebSocket manager con reconnect
│   └── useStudySession.ts           # Wrapper de sesión (quiz, socratic, capsule)
├── registry/
│   ├── index.ts                     # ComponentRegistry: Map<string, Component>
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
│       └── SurfaceRenderer.vue      # Renderiza recursivamente el árbol A2UI
├── stores/
│   └── a2ui.ts                      # Pinia store para Surface + DataModel
├── types/
│   └── a2ui.ts                      # TypeScript types del protocolo A2UI
├── assets/
│   └── css/
│       ├── themes/
│       │   ├── light.css
│       │   ├── dark.css
│       │   └── pastel.css
│       └── accessibility.css        # Clases de a11y dinámicas
└── middleware/
    └── auth.ts
```

### Definiciones TypeScript del protocolo

```typescript
// sai-frontend/types/a2ui.ts

export interface A2UISurface {
  surfaceId: string;
  rootComponent: string;
  components: Record<string, A2UIComponent>;
  dataModel: A2UIDataModel;
}

export interface A2UIComponent {
  id: string;
  type: string;
  children?: string[];
  props: Record<string, unknown>;
  events?: Record<string, string>;
}

export interface A2UIDataModel {
  theme: 'light' | 'dark' | 'system';
  fontFamily: string;
  fontScale: number;
  colorPalette: 'default' | 'pastel' | 'high-contrast';
  highContrast: boolean;
  reducedMotion: boolean;
  language: string;
}

export interface WSMessage {
  type: 'a2ui_full' | 'a2ui_update' | 'data_model_update' | 'error' | 'ping';
  payload: unknown;
  timestamp: string;
}

export interface DataModelUpdate {
  path: string;
  value: unknown;
  diff: Record<string, unknown>;
}

export interface A2UIUpdate {
  updates: ComponentPatch[];
}

export interface ComponentPatch {
  componentId: string;
  props?: Record<string, unknown>;
  children?: string[];
  events?: Record<string, string>;
}
```

### Composables Core

**`useA2UI.ts`** — Punto central de integración:
```typescript
// sai-frontend/composables/useA2UI.ts
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
      case 'a2ui_update': {
        const update = msg.payload as A2UIUpdate;
        if (!surface.value) return;
        for (const patch of update.updates) {
          const comp = surface.value.components[patch.componentId];
          if (comp) {
            if (patch.props) Object.assign(comp.props, patch.props);
            if (patch.children) comp.children = patch.children;
            if (patch.events) Object.assign(comp.events || {}, patch.events);
          }
        }
        break;
      }
      case 'data_model_update': {
        const update = msg.payload as DataModelUpdate;
        if (dataModel.value) {
          Object.assign(dataModel.value, update.diff);
          applyAccessibilityStyles(dataModel.value);
        }
        break;
      }
    }
  }

  connect(handleMessage);

  return { surface, dataModel, send };
}
```

**`useA2UIAccessibility.ts`** — Hot-reload de accesibilidad:
```typescript
// sai-frontend/composables/useA2UIAccessibility.ts
import { watch } from 'vue';
import type { A2UIDataModel } from '~/types/a2ui';

export function useA2UIAccessibility(dataModel: Ref<A2UIDataModel>) {
  watch(dataModel, (dm) => {
    if (!dm) return;
    const root = document.documentElement;
    root.style.setProperty('--font-family', dm.fontFamily);
    root.style.setProperty('--font-scale', dm.fontScale.toString());
    root.setAttribute('data-theme', dm.theme);
    root.setAttribute('data-palette', dm.colorPalette);

    if (dm.highContrast) root.classList.add('high-contrast');
    else root.classList.remove('high-contrast');

    if (dm.reducedMotion) root.classList.add('reduced-motion');
    else root.classList.remove('reduced-motion');
  }, { immediate: true, deep: true });
}
```

### Surface Renderer (Recursive Component)

```vue
<!-- sai-frontend/components/a2ui/SurfaceRenderer.vue -->
<template>
  <div v-if="surface" :id="`a2ui-surface-${surface.surfaceId}`">
    <component
      v-if="rootComponent"
      :is="resolveComponent(rootComponent.type)"
      v-bind="rootComponent.props"
      :a2ui-children="rootComponent.children"
      :a2ui-events="rootComponent.events"
      :surface="surface"
      @a2ui-event="handleEvent"
    />
    <div v-else class="a2ui-empty-surface">
      <p>Esperando contenido...</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import type { A2UISurface } from '~/types/a2ui';
import { componentRegistry } from '~/registry';

const props = defineProps<{ surface: A2UISurface | null }>();

const rootComponent = computed(() => {
  if (!props.surface) return null;
  return props.surface.components[props.surface.rootComponent] || null;
});

function resolveComponent(type: string) {
  return componentRegistry.get(type);
}

function handleEvent(eventName: string, payload: unknown) {
  // Dispatch to session store or emit to parent
}
</script>
```

### Registry de Componentes Nativos

```typescript
// sai-frontend/registry/index.ts
import Text from './Text.vue';
import RichText from './RichText.vue';
import Image from './Image.vue';
import AudioPlayer from './AudioPlayer.vue';
import VideoPlayer from './VideoPlayer.vue';
import Card from './Card.vue';
import Column from './Column.vue';
import Row from './Row.vue';
import QuizCard from './QuizCard.vue';
import SocraticDialog from './SocraticDialog.vue';
import ProgressBar from './ProgressBar.vue';
import Button from './Button.vue';

export const componentRegistry = new Map<string, Component>([
  ['Text', Text],
  ['RichText', RichText],
  ['Image', Image],
  ['AudioPlayer', AudioPlayer],
  ['VideoPlayer', VideoPlayer],
  ['Card', Card],
  ['Column', Column],
  ['Row', Row],
  ['QuizCard', QuizCard],
  ['SocraticDialog', SocraticDialog],
  ['ProgressBar', ProgressBar],
  ['Button', Button],
]);
```

### Ejemplo de Componente Nativo — QuizCard.vue

```vue
<template>
  <div class="quiz-card" :class="{ 'quiz-card--disabled': submitted }">
    <h3 class="quiz-card__question">{{ props.question }}</h3>
    <div class="quiz-card__options">
      <button
        v-for="(opt, i) in props.options"
        :key="i"
        class="quiz-card__option"
        :class="{ 'quiz-card__option--selected': selectedIndex === i }"
        @click="selectOption(i)"
        :disabled="submitted"
      >
        {{ opt }}
      </button>
    </div>
    <div class="quiz-card__actions">
      <button
        class="quiz-card__submit"
        :disabled="selectedIndex === null || submitted"
        @click="submit"
      >
        Enviar respuesta
      </button>
    </div>
    <div v-if="feedback" class="quiz-card__feedback" :class="feedbackClass">
      {{ feedback }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const props = defineProps<{
  question: string;
  options: string[];
  mode: 'single_choice' | 'multiple_choice';
  feedback?: string;
  correct?: boolean;
}>();

const emit = defineEmits<{
  (e: 'a2ui-event', eventName: string, payload: unknown): void;
}>();

const selectedIndex = ref<number | null>(null);
const submitted = ref(false);
const feedback = ref('');
const correct = ref(false);

function selectOption(i: number) { selectedIndex.value = i; }

function submit() {
  submitted.value = true;
  emit('a2ui-event', 'onSubmit', {
    selectedIndex: selectedIndex.value,
    mode: props.mode,
  });
}
</script>
```

---

## 4. Estructura de Carpetas del Proyecto Go (Final)

```
study-sessions-with-ai/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── router.go
│   │   ├── middleware/
│   │   │   ├── auth.go
│   │   │   ├── logging.go
│   │   │   └── cors.go
│   │   └── handlers/
│   │       ├── a2ui_handler.go
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
│   │   │   ├── selector.go           # 85% convergence algorithm
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
│   ├── 001_create_users.sql
│   ├── 002_create_concepts.sql
│   ├── 003_create_sessions.sql
│   ├── 004_create_quiz_items.sql
│   ├── 005_create_capsules.sql
│   └── 006_create_interactions.sql
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
├── web/                              # Assets estáticos servidos por Go embed
│   └── static/
├── configs/
│   └── app.yaml
├── scripts/
│   ├── migrate.sh
│   ├── seed.sh
│   └── dev.sh
├── go.mod
├── go.sum
├── Makefile
├── docker-compose.yml
└── .env.example
```

---

## 5. Plan de Desarrollo por Etapas (para Agentes Paralelos)

### Fase 0: Fundación del Proyecto — Agente A (4 días)

| # | Tarea | Detalle | Archivos |
|---|-------|---------|----------|
| 0.1 | Init módulo Go + Makefile | `go mod init github.com/org/sai-server`, Makefile con targets: `build`, `run`, `test`, `docker-build`, `db-migrate`, `db-seed` | `go.mod`, `Makefile` |
| 0.2 | Configuración con viper | Carga desde `configs/app.yaml`, variables de entorno, `.env`. Estructura `Config` con campos para DB (SQLite/PG), Redis, MinIO, OpenAI, Gemini TTS | `internal/config/config.go`, `configs/app.yaml`, `.env.example` |
| 0.3 | Tipos de dominio (GORM models) | `User`, `Concept`, `Session`, `QuizItem`, `Capsule`, `Interaction`, `UserConceptMastery`. Sin metadatos demográficos en `User`. | `internal/domain/*.go` |
| 0.4 | Interfaces de puertos (ports) | `LLMClient`, `TTSClient`, `ImageGenClient`, `StorageClient`, `CacheClient` | `internal/port/*.go` |
| 0.5 | Protocolo A2UI — tipos y serialización | Structs Go + JSON marshaling para `Surface`, `Component`, `DataModel`, `WSMessage`, todos los tipos de mensaje | `pkg/a2ui/types.go`, `pkg/a2ui/protocol.go` |
| 0.6 | Router Gin + middleware | Configuración de rutas, auth JWT, CORS, logging middleware | `internal/api/router.go`, `internal/api/middleware/*.go` |
| 0.7 | `main.go` — wire-up inicial | Ensamblaje de dependencias, inicialización de DB (SQLite por defecto), migración automática, arranque del servidor | `cmd/server/main.go` |
| 0.8 | Migraciones SQL | Archivos SQL para crear tablas en PostgreSQL | `migrations/001-006_*.sql` |
| 0.9 | Dockerfile + docker-compose | Dockerfile multi-stage (Go builder → Alpine runtime). `docker-compose.yml` con: server, postgres, redis, minio | `deploy/docker/Dockerfile`, `docker-compose.yml` |

### Fase 1: Adapters + A2UI Engine — Agentes B y C (6 días, en paralelo)

**Agente B: A2UI Engine + Core Handlers**

| # | Tarea | Archivos |
|---|-------|----------|
| 1.B.1 | A2UI Builder — construye árboles desde `Capsule` y `Quiz` domains | `pkg/a2ui/builder.go` |
| 1.B.2 | A2UI Diff — calcula `A2UIUpdate` (patch) entre dos superficies | `pkg/a2ui/diff.go` |
| 1.B.3 | WebSocket Hub — maneja conexiones por sessionId, broadcast, reconexión | `internal/api/handlers/a2ui_handler.go` |
| 1.B.4 | `data_model_update` handler — procesa cambios de a11y y re-broadcastea | en a2ui_handler.go |
| 1.B.5 | User handler REST (register, login, profile) | `internal/api/handlers/user_handler.go` + `internal/app/user/*.go` |
| 1.B.6 | Session handler REST (crear sesión, obtener estado) | `internal/api/handlers/session_handler.go` + `internal/app/session/*.go` |

**Agente C: Adapters de IA + Storage**

| # | Tarea | Archivos |
|---|-------|----------|
| 1.C.1 | OpenAI LLM adapter — `ChatCompletion`, función `GenerateStructured` con JSON schema | `internal/adapter/llm/openai.go` |
| 1.C.2 | Gemini TTS adapter — Google Cloud Text-to-Speech v1, genera `.mp3` y retorna `[]byte` | `internal/adapter/tts/gemini.go` |
| 1.C.3 | DALL·E image adapter — genera imágenes desde prompt textual | `internal/adapter/imagegen/dalle.go` |
| 1.C.4 | MinIO storage adapter — upload/download de assets, URLs prefirmadas | `internal/adapter/storage/minio.go` |
| 1.C.5 | Redis cache adapter — get/set con TTL, invalidación | `internal/adapter/cache/redis.go` |
| 1.C.6 | Factory functions con switch por config | `internal/adapter/llm/factory.go`, `tts/factory.go`, `imagegen/factory.go` |

### Fase 2: Módulos de Aprendizaje — Agentes D, E, F (8 días, en paralelo)

**Agente D: Cold-Start Diagnostic**

| # | Tarea | Archivos |
|---|-------|----------|
| 2.D.1 | IRT models (1PL, 2PL, 3PL) — probabilidad de respuesta correcta | `pkg/irt/model.go` |
| 2.D.2 | Estimador MLE/EAP — cálculo de θ (habilidad) | `pkg/irt/estimator.go` |
| 2.D.3 | CAT con Maximum Fisher Information — selecciona siguiente ítem | `internal/app/coldstart/cat.go` |
| 2.D.4 | Service cold-start — orquesta el flujo: init → select item → answer → update θ → repeat N times → finalize | `internal/app/coldstart/service.go` |
| 2.D.5 | Cold-start handler REST | `internal/api/handlers/coldstart_handler.go` |
| 2.D.6 | Seed data: 50 ítems de diagnóstico calibrados para diferentes conceptos | `scripts/seed_coldstart_items.sql` |

**Agente E: Quiz Engine (85% Rule)**

| # | Tarea | Archivos |
|---|-------|----------|
| 2.E.1 | BKT por concepto — actualiza P(learned), P(guess), P(slip), P(transit) | `pkg/bkt/bayesian.go` |
| 2.E.2 | SM-2 extended — easiness factor, interval, next_review | `pkg/spacedrep/sm2.go`, `pkg/spacedrep/scheduler.go` |
| 2.E.3 | **Selector 85%** — algoritmo central: `P_target = 0.85`, mezcla ponderada α·master + (1-α)·review, greedy para converger | `internal/app/quiz/selector.go` |
| 2.E.4 | Evaluador de respuestas — compara con respuesta correcta, registra tiempo | `internal/app/quiz/evaluator.go` |
| 2.E.5 | Quiz Engine — orquesta selector + evaluador + actualiza BKT | `internal/app/quiz/engine.go` |
| 2.E.6 | Quiz handler REST | `internal/api/handlers/quiz_handler.go` |

**Agente F: Dual Coding Orchestrator**

| # | Tarea | Archivos |
|---|-------|----------|
| 2.F.1 | Generador de texto educativo vía LLM (prompt específico para cápsulas) | `internal/app/dualcode/capsule.go` |
| 2.F.2 | Orquestador: genera texto → en paralelo genera audio (Gemini TTS) + imagen (DALL·E) | `internal/app/dualcode/orchestrator.go` |
| 2.F.3 | Ensamblador de cápsula: `DualCapsule` → `a2ui.Surface` usando Builder | en orchestrator.go |
| 2.F.4 | **Validación Codificación Dual**: assert que `len(capsule.Modalities) >= 2` | en orchestrator.go |
| 2.F.5 | Capsule handler REST | `internal/api/handlers/capsule_handler.go` |

### Fase 3: Session Pipeline + Socratic Remediation — Agentes G y H (6 días, en paralelo)

**Agente G: Session Pipeline**

| # | Tarea | Archivos |
|---|-------|----------|
| 3.G.1 | Configurar pipeline de sesión: `coldstart → capsule_1 → quiz_1 → capsule_2 → quiz_2 → ... → completion` | `internal/app/session/pipeline.go` |
| 3.G.2 | Máquina de estados de sesión: `coldstart → learning → quiz → socratic_remediation → quiz → learning → completed` | `internal/app/session/service.go` |
| 3.G.3 | Event sourcing para logging de interacciones en Redis y PostgreSQL | `internal/app/session/events.go` |

**Agente H: Socratic Remediation**

| # | Tarea | Archivos |
|---|-------|----------|
| 3.H.1 | Plantillas de prompt para Feynman Technique (clasifica error, genera analogía, explica en simple) | `internal/app/socratic/prompts.go` |
| 3.H.2 | Remediador: analiza respuesta incorrecta → clasifica tipo de error → invoca LLM con prompt Feynman → devuelve intervención | `internal/app/socratic/remediator.go` |
| 3.H.3 | Bucle de Interrogación Elaborativa: genera pregunta de "por qué", espera respuesta del usuario, evalúa coherencia con LLM, decide si avanzar o repetir | en remediator.go |
| 3.H.4 | Socratic handler REST (recibe respuestas elaborativas, devuelve feedback) | `internal/api/handlers/socratic_handler.go` |

### Fase 4: Frontend + K8s + Pruebas — Agentes I y J (6 días, en paralelo)

**Agente I: Frontend Nuxt 3 (en repo separado `sai-frontend`)**

| # | Tarea |
|---|-------|
| 4.I.1 | Init proyecto Nuxt 3 + TypeScript + Pinia |
| 4.I.2 | Tipos TypeScript del protocolo A2UI |
| 4.I.3 | Implementar `useWebSocket` composable (reconexión exponencial) |
| 4.I.4 | Implementar `useA2UI` + `useA2UIAccessibility` composables |
| 4.I.5 | Implementar Component Registry completo (12 componentes) |
| 4.I.6 | Implementar `SurfaceRenderer.vue` con resolución recursiva |
| 4.I.7 | Páginas: `index.vue` (login), `coldstart.vue`, `session/[id].vue` |
| 4.I.8 | Store Pinia `a2ui.ts` |
| 4.I.9 | CSS themes: light, dark, pastel + estilos de a11y dinámicos |
| 4.I.10 | Integración end-to-end con backend |

**Agente J: K8s + CI/CD + Pruebas**

| # | Tarea |
|---|-------|
| 4.J.1 | Finalizar y validar todos los manifiestos K8s |
| 4.J.2 | Probar despliegue en minikube/kind local |
| 4.J.3 | Configurar Ingress con TLS (cert-manager) |
| 4.J.4 | Health checks: `/healthz`, `/readyz` |
| 4.J.5 | Métricas Prometheus + `/metrics` endpoint |
| 4.J.6 | Tests unitarios (Go `testing` + testify) para: IRT, BKT, SM-2, selector 85%, A2UI diff |
| 4.J.7 | Tests de integración: flujo cold-start completo, flujo quiz, flujo dual code |
| 4.J.8 | Documentación README.md completa con instrucciones de despliegue |

---

## 6. Algoritmo Central: Selector del 85% (detalle matemático)

```
P_target = 0.85

Para cada pregunta candidata q_i:
  P_success = IRT_3PL(θ_user, a_i, b_i, c_i)
  // a = discriminación, b = dificultad, c = guessing

Categorizar preguntas:
  Dominadas:     P_success >= 0.90
  Near-mastery:  0.70 <= P_success < 0.90
  No-dominadas:  P_success < 0.70

Mezcla óptima:
  α = proporción de no-dominadas en la sesión
  (1-α) = proporción de near-mastery (repaso espaciado)

  Buscar α tal que:
    α * avg(P_no_dominadas) + (1-α) * avg(P_near_mastery) = 0.85

  Si no hay suficientes preguntas en una categoría,
  rellenar con la otra y recalcular α.

Intercalado (interleaving):
  Las preguntas se ordenan por round-robin entre conceptos,
  nunca dos preguntas seguidas del mismo concepto.
  (Implementado en el scheduler de sesión)
```

---

## Resumen de Entregables

| Entregable | Estado |
|------------|--------|
| Diagrama de arquitectura (monolito + K8s + APIs IA + A2UI) | Completado (diagrama ASCII) |
| Diseño de API REST + WebSocket + protocolo A2UI JSON | Completado (estructuras Go + TS) |
| Manifiestos Kubernetes (9 archivos YAML) | Completado |
| Estructura de carpetas Go (70+ archivos) | Completado |
| Estructura de frontend Nuxt 3 (repositorio separado) | Completado |
| Plan por etapas con 6 fases para agentes paralelos | Completado |
| Algoritmo central del 85% | Completado |
| Modelos de datos PostgreSQL | Completado (6 tablas) |

El plan está listo para iniciar la implementación técnica. ¿Quieres que proceda con la Fase 0 (fundación del proyecto)?

---

## User

Guarda el plan como un archivo architecture_plan.md

---

## Assistant (Plan · DeepSeek V4 Pro · 413.1s)

---

## User

Guarda el plan como un archivo architecture_plan.md

---

## Assistant (Plan · DeepSeek V4 Pro · 10.1s)

Estoy en **Plan Mode** (solo lectura), por lo que no puedo escribir archivos en este momento. Para guardar `architecture_plan.md` necesito que me saques de este modo.

¿Procedo a escribir el archivo `architecture_plan.md` con todo el contenido del plan?

---

## User

Guarda el plan como un archivo architecture_plan.md

---

## Assistant (Build · DeepSeek V4 Pro · 166.4s)

**Tool: write**

**Input:**
```json
{}
```

**Error:**
```
Tool execution aborted
```

---

## User

Guarda el plan como un archivo architecture_plan.md

---

## Assistant (Build · DeepSeek V4 Pro · 56.2s)

**Tool: write**

**Input:**
```json
{}
```

**Error:**
```
Tool execution aborted
```

---

## User

Guarda el plan como un archivo architecture_plan.md

---

## Assistant (Build · MiniMax M2.5 Free · 76.3s)






**Tool: write**

**Input:**
```json
{
  "filePath": "/home/davidyesc/projects/study-sessions-with-ai/architecture_plan.md",
  "content": "# Arquitectura del Sistema de Estudio Adaptativo Multimodal (PoC)\n\n## Resumen Ejecutivo\n\nEste documento presenta la propuesta arquitectónica completa para una Prueba de Concepto (PoC) de un sistema de estudio adaptativo multimodal basado en IA. El sistema implementa los cinco pilares funcionales especificados:\n\n1. **Motor de Interfaz Declarativa A2UI** — Backend envía descripciones de componentes JSON, cliente renderiza componentes nativos\n2. **Orquestador de Codificación Dual** — Genera contenido en al menos dos modalidades sincronizadas (Teoría de Codificación Dual)\n3. **Módulo de Diagnóstico Cold-Start** — Evaluación adaptativa sin datos demográficos, usando IRT y clustering\n4. **Bucle de Evaluación (Regla del 85%)** — Algoritmo de selección de preguntas que converge al 85% de probabilidad de éxito\n5. **Sistema de Remediación Socrática** — Intervenciones basadas en Técnica Feynman + Interrogación Elaborativa\n\n---\n\n## 1. Arquitectura General\n\n### 1.1 Diagrama de Componentes\n\n```\n┌──────────────────────────────────────────────────────────────────────────┐\n│                         CLUSTER KUBERNETES                                │\n│                                                                           │\n│  ┌─────────────┐  ┌──────────────────────────────────────────────────┐   │\n│  │   Ingress   │  │          Monolito Go (sai-server)                │   │\n│  │   (nginx)   │  │                                                  │   │\n│  │             │  │  ┌─────────┐  ┌─────────┐  ┌──────────────┐    │   │\n│  │  /api/* ────┼─►│ API REST │  │ WebSocket│  │  A2UI Engine │    │   │\n│  │  /ws/*  ────┼─►│  (Gin)   │  │  (ws)    │  │  (JSON Gen)  │    │   │\n│  │             │  │  └────┬────┘  └────┬────┘  └──────┬───────┘    │   │\n│  └─────────────┘  │       │            │              │            │   │\n│                   │  ┌────┴────────────┴──────────────┴───────┐    │   │\n│                   │  │          Service Layer                  │    │   │\n│                   │  │  ┌──────────┐  ┌──────────┐  ┌─────────┐ │    │   │\n│                   │  │  │ Session  │  │  User    │  │ ColdStart│ │    │   │\n│                   │  │  │ Manager  │  │ Service  │  │ Service │ │    │   │\n│                   │  │  └──────────┘  └──────────┘  └─────────┘ │    │   │\n│                   │  │  ┌──────────┐  ┌──────────┐  ┌─────────┐ │    │   │\n│                   │  │  │ Quiz     │  │DualCode  │  │ Socratic│ │    │   │\n│                   │  │  │ Engine   │  │Orchest.  │  │Remediat.│ │    │   │\n│                   │  │  └──────────┘  └──────────┘  └─────────┘ │    │   │\n│                   │  └───────────────────────────────────────┘    │   │\n│                   │  ┌───────────────────────────────────────┐    │   │\n│                   │  │        Ports / Adapters                │    │   │\n│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │\n│                   │  │  │ LLM  │ │ TTS  │ │ Image/Video  │  │    │   │\n│                   │  │  │Client│ │Client│ │   Client     │  │    │   │\n│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │\n│                   │  │  ┌──────┐ ┌──────┐ ┌──────────────┐  │    │   │\n│                   │  │  │  DB  │ │Redis │ │   MinIO/S3   │  │    │   │\n│                   │  │  │(GORM)│ │(Pub) │ │  (Assets)    │  │    │   │\n│                   │  │  └──────┘ └──────┘ └──────────────┘  │    │   │\n│                   │  └───────────────────────────────────────┘    │   │\n│                   └──────────────────────────────────────────────────┘   │\n│                                                                           │\n│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐         │\n│  │ PostgreSQL 14    │  │    Redis 7       │  │   MinIO (S3)    │         │\n│  │  (StatefulSet)   │  │  (StatefulSet)   │  │ (StatefulSet)   │         │\n│  └──────────────────┘  └──────────────────┘  └──────────────────┘         │\n└──────────────────────────────────────────────────────────────────────────┘\n\n                    ┌──────────────────────────────────────┐\n                    │          APIs Externas                │\n                    │  ┌────────┐  ┌──────────┐  ┌───────┐│\n                    │  │ OpenAI │  │Gemini TTS│  │ DALL·E││\n                    │  │GPT-4o  │  │(Google)  │  │       ││\n                    │  └────────┘  └──────────┘  └───────┘│\n                    └──────────────────────────────────────┘\n\n  ┌─────────────────────────────────────────────────────────────────┐\n  │                   Cliente A2UI (Nuxt 3 - Repositorio Separado) │\n│  ┌─────────────────────────────────────────────────────────────┐ │\n│  │  A2UI Renderer Engine                                        │ │\n│  │  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────────┐  │ │\n│  │  │Component │ │ Layout   │ │Theme/    │ │ Event        │  │ │\n│  │  │Registry  │ │ Resolver │ │A11y Mgr  │ │ Dispatcher   │  │ │\n│  │  └──────────┘ └──────────┘ └──────────┘ └──────────────┘  │ │\n│  └─────────────────────────────────────────────────────────────┘ │\n│  Native Components: Text, RichText, Image, AudioPlayer,         │\n│  VideoPlayer, Card, QuizCard, SocraticDialog, ProgressBar        │\n└─────────────────────────────────────────────────────────────────┘\n```\n\n### 1.2 Stack Tecnológico\n\n| Componente | Tecnología | Versión |\n|------------|------------|---------|\n| Lenguaje Backend | Go | 1.22+ |\n| Framework HTTP | Gin | latest |\n| Base de Datos (Prod) | PostgreSQL | 14+ |\n| Base de Datos (Dev) | SQLite | 3.x |\n| Cache/Colas | Redis | 7.x |\n| Storage Objetos | MinIO | latest |\n| LLM | OpenAI | GPT-4o |\n| TTS | Google Gemini TTS | v1 |\n| Imágenes | DALL·E 3 / Stability AI | latest |\n| Contenedores | Docker | latest |\n| Orquestación | Kubernetes | 1.28+ |\n| Frontend | Nuxt 3 + Vue 3 | latest |\n\n### 1.3 Flujo de Datos Principal\n\n1. **Usuario nuevo** → `/api/coldstart/start` → IRT adaptive test → `θ` estimado → clustering\n2. **Sesión activa** → WebSocket `/ws/session/{id}` → A2UI surface inicial\n3. **Pregunta quiz** → algoritmo 85% → selecciona pregunta → `a2ui_update` → usuario responde\n4. **Respuesta incorrecta** → `/api/socratic/remediate` → LLM Feynman → Interrogación Elaborativa\n5. **Generación cápsula** → `/api/capsules/generate` → dual coding orchestrator → texto + audio + imagen → A2UI tree\n\n---\n\n## 2. Diseño de API REST + WebSocket\n\n### 2.1 Endpoints REST\n\n| Método | Ruta | Descripción |\n|--------|------|-------------|\n| `POST` | `/api/users/register` | Registro de usuario (sin datos demográficos) |\n| `POST` | `/api/users/login` | Login, retorna JWT |\n| `GET` | `/api/users/me` | Perfil de usuario con habilidad estimada |\n| `POST` | `/api/coldstart/start` | Inicia diagnóstico cold-start |\n| `POST` | `/api/coldstart/answer` | Envía respuesta a ítem de diagnóstico |\n| `GET` | `/api/coldstart/{id}/result` | Resultado estimado (θ + cluster) |\n| `POST` | `/api/sessions` | Crea nueva sesión de estudio |\n| `GET` | `/api/sessions/{id}` | Estado actual de la sesión |\n| `GET` | `/api/sessions/{id}/next` | Obtiene siguiente ítem (cápsula o quiz) |\n| `POST` | `/api/sessions/{id}/quiz/answer` | Envía respuesta a pregunta de quiz |\n| `POST` | `/api/sessions/{id}/socratic/response` | Envía respuesta a intervención socrática |\n| `POST` | `/api/capsules/generate` | Solicita generación de cápsula dual |\n| `GET` | `/api/capsules/{id}` | Obtiene胶囊 completa |\n| `POST` | `/api/sessions/{id}/a11y` | Actualiza preferencias accesibilidad |\n| `GET` | `/api/assets/{type}/{filename}` | Descarga assets (audio/imagen) |\n\n### 2.2 WebSocket\n\n**Endpoint**: `WS /ws/session/{sessionId}`\n\n#### Formato de Mensaje (Envelope JSON)\n\n```json\n{\n  \"type\": \"a2ui_full | a2ui_update | data_model_update | error | ping\",\n  \"payload\": {},\n  \"timestamp\": \"2024-01-15T10:30:00Z\"\n}\n```\n\n### 2.3 Protocolo A2UI: Tipos de Mensajes\n\n#### a2ui_full — Renderizado Inicial Completo\n\n```json\n{\n  \"type\": \"a2ui_full\",\n  \"payload\": {\n    \"surfaceId\": \"session-abc\",\n    \"rootComponent\": \"root\",\n    \"components\": {\n      \"root\": {\n        \"id\": \"root\",\n        \"type\": \"Column\",\n        \"children\": [\"header\", \"body\", \"footer\"],\n        \"props\": { \"gap\": 16, \"padding\": 24 }\n      },\n      \"header\": {\n        \"id\": \"header\",\n        \"type\": \"Row\",\n        \"children\": [\"title\", \"progress-bar\"],\n        \"props\": { \"alignment\": \"space-between\" }\n      },\n      \"title\": {\n        \"id\": \"title\",\n        \"type\": \"Text\",\n        \"props\": { \"content\": \"Física Cuántica: Dualidad Onda-Partícula\", \"variant\": \"h2\" }\n      },\n      \"progress-bar\": {\n        \"id\": \"progress-bar\",\n        \"type\": \"ProgressBar\",\n        \"props\": { \"value\": 0.35, \"max\": 1.0 }\n      },\n      \"body\": {\n        \"id\": \"body\",\n        \"type\": \"Card\",\n        \"children\": [\"text-content\", \"audio-player\", \"image-diagram\"],\n        \"props\": { \"elevation\": 2 }\n      },\n      \"text-content\": {\n        \"id\": \"text-content\",\n        \"type\": \"RichText\",\n        \"props\": { \"markdown\": \"La dualidad onda-partícula es un concepto fundamental...\", \"accessible\": true }\n      },\n      \"audio-player\": {\n        \"id\": \"audio-player\",\n        \"type\": \"AudioPlayer\",\n        \"props\": { \"url\": \"/api/assets/audio/capsule-123.mp3\", \"autoPlay\": false }\n      },\n      \"image-diagram\": {\n        \"id\": \"image-diagram\",\n        \"type\": \"Image\",\n        \"props\": { \"url\": \"/api/assets/image/diagram-123.png\", \"altText\": \"Diagrama del experimento de doble rendija\" }\n      },\n      \"quiz-card\": {\n        \"id\": \"quiz-card\",\n        \"type\": \"QuizCard\",\n        \"props\": {\n          \"question\": \"¿Por qué el patrón de interferencia desaparece al observar las partículas?\",\n          \"options\": [\"A) La luz se comporta como onda\", \"B) La medición perturba el sistema cuántico\", \"C) El detector cambia la frecuencia\", \"D) La velocidad cambia\"],\n          \"mode\": \"single_choice\"\n        },\n        \"events\": { \"onSubmit\": \"/api/sessions/abc/quiz/answer\" }\n      },\n      \"socratic-prompt\": {\n        \"id\": \"socratic-prompt\",\n        \"type\": \"SocraticDialog\",\n        \"props\": {\n          \"prompt\": \"Antes de continuar, explica con tus palabras: ¿por qué la medición colapsa la función de onda?\",\n          \"context\": \"dualidad-onda-particula\"\n        },\n        \"events\": { \"onSubmit\": \"/api/sessions/abc/socratic/response\" }\n      },\n      \"footer\": {\n        \"id\": \"footer\",\n        \"type\": \"Row\",\n        \"children\": [\"btn-prev\", \"btn-next\"],\n        \"props\": { \"alignment\": \"center\" }\n      }\n    },\n    \"dataModel\": {\n      \"theme\": \"system\",\n      \"fontFamily\": \"sans-serif\",\n      \"fontScale\": 1.0,\n      \"colorPalette\": \"default\",\n      \"highContrast\": false,\n      \"reducedMotion\": false,\n      \"language\": \"es\"\n    }\n  }\n}\n```\n\n#### data_model_update — Cambio de Accesibilidad en Caliente\n\n```json\n{\n  \"type\": \"data_model_update\",\n  \"payload\": {\n    \"path\": \"fontFamily\",\n    \"value\": \"OpenDyslexic\",\n    \"diff\": {\n      \"fontFamily\": \"OpenDyslexic\",\n      \"colorPalette\": \"pastel\",\n      \"fontScale\": 1.2\n    }\n  }\n}\n```\n\n#### a2ui_update — Parche Parcial del Árbol\n\n```json\n{\n  \"type\": \"a2ui_update\",\n  \"payload\": {\n    \"updates\": [\n      { \"componentId\": \"quiz-card\", \"props\": { \"question\": \"Nueva pregunta...\", \"options\": [\"A\", \"B\", \"C\"] }},\n      { \"componentId\": \"progress-bar\", \"props\": { \"value\": 0.70 }}\n    ]\n  }\n}\n```\n\n---\n\n## 3. Protocolo A2UI: Modelo de Datos\n\n### 3.1 Estructuras Go\n\n```go\npackage a2ui\n\n// Surface representa una \"pantalla\" completa de A2UI\ntype Surface struct {\n    SurfaceID     string                `json:\"surfaceId\"`\n    RootComponent string                `json:\"rootComponent\"`\n    Components    map[string]Component  `json:\"components\"`\n    DataModel     DataModel             `json:\"dataModel\"`\n}\n\n// Component es un nodo del árbol declarativo (Modelo de Lista de Adyacencia)\ntype Component struct {\n    ID       string                 `json:\"id\"`\n    Type     string                 `json:\"type\"` // \"Text\", \"Card\", \"QuizCard\", etc.\n    Children []string               `json:\"children,omitempty\"`\n    Props    map[string]interface{} `json:\"props\"`\n    Events   map[string]string      `json:\"events,omitempty\"`\n}\n\n// DataModel contiene estado global de accesibilidad/tema\ntype DataModel struct {\n    Theme         string  `json:\"theme\"`\n    FontFamily    string  `json:\"fontFamily\"`\n    FontScale     float64 `json:\"fontScale\"`\n    ColorPalette  string  `json:\"colorPalette\"`\n    HighContrast  bool    `json:\"highContrast\"`\n    ReducedMotion bool    `json:\"reducedMotion\"`\n    Language      string  `json:\"language\"`\n}\n\n// WSMessage es el envelope de WebSocket\ntype WSMessage struct {\n    Type      string      `json:\"type\"`\n    Payload   interface{} `json:\"payload\"`\n    Timestamp string      `json:\"timestamp\"`\n}\n```\n\n### 3.2 Componentes Soportados\n\n| Tipo | Descripción |\n|------|-------------|\n| `Text` | Texto simple con variantes (h1-h6, body, caption) |\n| `RichText` | Markdown rendering |\n| `Image` | Imagen con altText |\n| `AudioPlayer` | Reproductor de audio con controles |\n| `VideoPlayer` | Reproductor de video con subtítulos |\n| `Card` | Contenedor con elevación |\n| `Column` | Contenedor vertical |\n| `Row` | Contenedor horizontal |\n| `QuizCard` | Tarjeta de pregunta con opciones |\n| `SocraticDialog` | Diálogo socrático con input |\n| `ProgressBar` | Barra de progreso |\n| `Button` | Botón con tipos (primary, secondary) |\n\n---\n\n## 4. Módulos de Dominio\n\n### 4.1 Cold-Start Diagnostic (Diagnóstico de Inicio en Frío)\n\n**Objetivo**: Estimar el nivel de habilidad latente del usuario sin usar datos demográficos.\n\n**Algoritmo**:\n\n1. **Item Response Theory (IRT)** — Usa modelo 3PL:\n   - `P(correct) = c + (1-c) * 1/(1 + exp(-a(θ-b)))`\n   - `a` = discriminación, `b` = dificultad, `c` = adivinación\n\n2. **Computerized Adaptive Testing (CAT)**:\n   - Selecciona siguiente ítem por Maximum Fisher Information\n   - I(theta) = sum over items of (a_i^2 * P_i * (1-P_i) * ((P_i - c_i)/(1-c_i))^2)\n   - Actualiza θ con Maximum Likelihood Estimation\n\n3. **K-Means Clustering**:\n   - Después de N respuestas (ej: 10), agrupa patrones de respuesta\n   - Asigna cluster como \"nivel inicial\" (beginner, intermediate, advanced)\n\n**Restricciones**:\n- ❌ NO usar edad, género, ubicación, educación\n- ✅ Solo usar patrones de interacción y respuestas a ítems calibrados\n\n### 4.2 Quiz Engine (Regla del 85%)\n\n**Objetivo**: Mezclar material no dominado y repasos para que la probabilidad predictiva de éxito converja al 85%.\n\n**Algoritmo**:\n\n```\nP_target = 0.85\n\nPara cada pregunta candidatos:\n  P_success = IRT_3PL(θ_user, a, b, c)\n\nCategorizar:\n  Dominadas:     P_success >= 0.90\n  Near-mastery:  0.70 <= P_success < 0.90\n  No-dominadas:  P_success < 0.70\n\nMezcla óptima:\n  α = proporción de no-dominadas\n  (1-α) = proporción de near-mastery (repaso)\n  \n  Buscar α tal que:\n    α * avg(P_no_dominadas) + (1-α) * avg(P_near_mastery) ≈ 0.85\n```\n\n**Integrado con**:\n- **Bayesian Knowledge Tracing (BKT)** — Actualiza estado de dominio por concepto\n- **SM-2 Spaced Repetition** — Calcula intervalos de repaso óptimos\n- **Interleaving** — Evita cluster de mismo concepto en sucesión\n\n### 4.3 Dual Coding Orchestrator\n\n**Objetivo**: Cada cápsula de aprendizaje entrega información en al menos dos modalidades sincronizadas.\n\n**Arquitectura de datos por cápsula**:\n\n```go\ntype DualCapsule struct {\n    ID          string\n    Topic       string\n    Modalities  []Modality // Text + Audio + Image mínimo\n    A2UI_Tree   a2ui.Surface\n}\n\ntype Modality struct {\n    Type      string // \"text\", \"audio\", \"image\", \"video\"\n    Content   string // texto o URL\n    Metadata  map[string]interface{}\n}\n```\n\n**Flujo de generación**:\n1. Recibe topic → LLM genera texto educativo\n2. En paralelo: TTS genera audio, ImageGen genera imagen/diagrama\n3. Ensambla en DualCapsule\n4. Builder convierte a A2UI tree\n5. **Valida**: `len(Modalidades) >= 2`\n\n### 4.4 Socratic Remediation\n\n**Objetivo**: Cuando el estudiante falla, generar intervención basada en Técnica Feynman + Interrogación Elaborativa.\n\n**Flujo**:\n\n1. Estudiante responde incorrectamente\n2. Clasificar tipo de error (conceptual, procedimental, olvida)\n3. LLM genera explicación estilo Feynman:\n   - Identifica concepto complejo\n   - Explica en términos simples\n   - Usa analogía de la vida real\n\n4. **Interrogación Elaborativa**:\n   - \"¿Por qué crees que ocurre X?\"\n   - Evalúa respuesta del usuario con LLM\n   - Decide: más scaffolding o avanzar\n\n---\n\n## 5. Modelos de Datos (PostgreSQL)\n\n### 5.1 Esquema de Tablas\n\n```sql\n-- Conceptos (Knowledge Components)\nCREATE TABLE concepts (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    parent_id UUID REFERENCES concepts(id),\n    name TEXT NOT NULL,\n    description TEXT,\n    difficulty FLOAT NOT NULL DEFAULT 0.5,\n    created_at TIMESTAMPTZ DEFAULT now()\n);\n\n-- Usuarios (SIN metadatos demográficos)\nCREATE TABLE users (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    email TEXT UNIQUE NOT NULL,\n    password_hash TEXT NOT NULL,\n    estimated_theta FLOAT,\n    theta_uncertainty FLOAT,\n    cluster VARCHAR(50),\n    created_at TIMESTAMPTZ DEFAULT now()\n);\n\n-- Dominio por usuario-concepto (BKT state)\nCREATE TABLE user_concept_mastery (\n    user_id UUID REFERENCES users(id),\n    concept_id UUID REFERENCES concepts(id),\n    p_learned FLOAT NOT NULL DEFAULT 0.3,\n    p_guess FLOAT NOT NULL DEFAULT 0.1,\n    p_slip FLOAT NOT NULL DEFAULT 0.1,\n    p_transit FLOAT NOT NULL DEFAULT 0.2,\n    last_practiced TIMESTAMPTZ,\n    easiness_factor FLOAT DEFAULT 2.5,\n    interval_days INT DEFAULT 1,\n    repetitions INT DEFAULT 0,\n    next_review TIMESTAMPTZ,\n    PRIMARY KEY (user_id, concept_id)\n);\n\n-- Items de quiz (IRT calibrated)\nCREATE TABLE quiz_items (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    concept_id UUID REFERENCES concepts(id),\n    difficulty_irt FLOAT NOT NULL,\n    discrimination FLOAT DEFAULT 1.0,\n    guessing FLOAT DEFAULT 0.25,\n    content JSONB NOT NULL,\n    modality TEXT DEFAULT 'text',\n    created_at TIMESTAMPTZ DEFAULT now()\n);\n\n-- Sesiones de estudio\nCREATE TABLE sessions (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    user_id UUID REFERENCES users(id),\n    state TEXT NOT NULL DEFAULT 'coldstart',\n    a2ui_snapshot JSONB,\n    target_success_rate FLOAT DEFAULT 0.85,\n    started_at TIMESTAMPTZ DEFAULT now(),\n    completed_at TIMESTAMPTZ\n);\n\n-- Cápsulas de aprendizaje\nCREATE TABLE capsules (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    topic TEXT NOT NULL,\n    modalities JSONB NOT NULL,\n    a2ui_tree JSONB NOT NULL,\n    session_id UUID REFERENCES sessions(id),\n    created_at TIMESTAMPTZ DEFAULT now()\n);\n\n-- Interacciones (logs para analytics)\nCREATE TABLE interactions (\n    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n    session_id UUID REFERENCES sessions(id),\n    type TEXT NOT NULL,\n    payload JSONB NOT NULL,\n    was_correct BOOLEAN,\n    response_time_ms INT,\n    remediation_generated JSONB,\n    created_at TIMESTAMPTZ DEFAULT now()\n);\n```\n\n---\n\n## 6. Estructura de Carpetas del Proyecto Go\n\n```\nstudy-sessions-with-ai/\n├── cmd/\n│   └── server/\n│       └── main.go                    # Entry point\n├── internal/\n│   ├── api/\n│   │   ├── router.go                  # Registro de rutas Gin\n│   │   ├── middleware/\n│   │   │   ├── auth.go                # JWT middleware\n│   │   │   ├── logging.go\n│   │   │   └── cors.go\n│   │   └── handlers/\n│   │       ├── a2ui_handler.go        # WebSocket handler\n│   │       ├── coldstart_handler.go\n│   │       ├── session_handler.go\n│   │       ├── capsule_handler.go\n│   │       ├── quiz_handler.go\n│   │       ├── socratic_handler.go\n│   │       └── user_handler.go\n│   ├── app/\n│   │   ├── session/\n│   │   │   ├── service.go\n│   │   │   ├── pipeline.go\n│   │   │   └── types.go\n│   │   ├── coldstart/\n│   │   │   ├── service.go\n│   │   │   ├── irt.go\n│   │   │   ├── cat.go\n│   │   │   └── types.go\n│   │   ├── dualcode/\n│   │   │   ├── orchestrator.go\n│   │   │   ├── capsule.go\n│   │   │   └── types.go\n│   │   ├── quiz/\n│   │   │   ├── engine.go\n│   │   │   ├── selector.go\n│   │   │   ├── evaluator.go\n│   │   │   └── types.go\n│   │   ├── socratic/\n│   │   │   ├── remediator.go\n│   │   │   ├── prompts.go\n│   │   │   └── types.go\n│   │   └── user/\n│   │       ├── service.go\n│   │       └── types.go\n│   ├── domain/\n│   │   ├── user.go\n│   │   ├── concept.go\n│   │   ├── session.go\n│   │   ├── capsule.go\n│   │   ├── quiz.go\n│   │   └── a2ui.go\n│   ├── port/\n│   │   ├── llm.go\n│   │   ├── tts.go\n│   │   ├── imagegen.go\n│   │   ├── storage.go\n│   │   └── cache.go\n│   ├── adapter/\n│   │   ├── llm/\n│   │   │   ├── openai.go\n│   │   │   └── factory.go\n│   │   ├── tts/\n│   │   │   ├── gemini.go\n│   │   │   └── factory.go\n│   │   ├── imagegen/\n│   │   │   ├── dalle.go\n│   │   │   └── factory.go\n│   │   ├── storage/\n│   │   │   └── minio.go\n│   │   └── cache/\n│   │       └── redis.go\n│   └── config/\n│       └── config.go\n├── pkg/\n│   ├── a2ui/\n│   │   ├── protocol.go\n│   │   ├── builder.go\n│   │   ├── diff.go\n│   │   └── types.go\n│   ├── spacedrep/\n│   │   ├── sm2.go\n│   │   └── scheduler.go\n│   ├── irt/\n│   │   ├── model.go\n│   │   └── estimator.go\n│   ├── bkt/\n│   │   └── bayesian.go\n│   └── adaptive/\n│       └── difficulty.go\n├── migrations/\n├── deploy/\n│   ├── k8s/\n│   │   ├── namespace.yaml\n│   │   ├── configmap.yaml\n│   │   ├── secrets.yaml\n│   │   ├── deployment-postgres.yaml\n│   │   ├── deployment-redis.yaml\n│   │   ├── deployment-minio.yaml\n│   │   ├── deployment-server.yaml\n│   │   └── ingress.yaml\n│   └── docker/\n│       └── Dockerfile\n├── configs/\n│   └── app.yaml\n├── scripts/\n├── go.mod\n├── go.sum\n├── Makefile\n├── docker-compose.yml\n└── .env.example\n```\n\n---\n\n## 7. Frontend: Repositorio Separado (Nuxt 3)\n\n### 7.1 Estructura de Carpetas\n\n```\nsai-frontend/\n├── nuxt.config.ts\n├── package.json\n├── tsconfig.json\n├── pages/\n│   ├── index.vue\n│   ├── coldstart.vue\n│   └── session/\n│       └── [id].vue\n├── composables/\n│   ├── useA2UI.ts\n│   ├── useA2UIAccessibility.ts\n│   ├── useWebSocket.ts\n│   └── useStudySession.ts\n├── registry/\n│   ├── index.ts\n│   ├── Text.vue\n│   ├── RichText.vue\n│   ├── Image.vue\n│   ├── AudioPlayer.vue\n│   ├── VideoPlayer.vue\n│   ├── Card.vue\n│   ├── Column.vue\n│   ├── Row.vue\n│   ├── QuizCard.vue\n│   ├── SocraticDialog.vue\n│   ├── ProgressBar.vue\n│   └── Button.vue\n├── components/\n│   └── a2ui/\n│       └── SurfaceRenderer.vue\n├── stores/\n│   └── a2ui.ts\n├── types/\n│   └── a2ui.ts\n├── assets/\n│   └── css/\n│       ├── themes/\n│       │   ├── light.css\n│       │   ├── dark.css\n│       │   └── pastel.css\n│       └── accessibility.css\n└── middleware/\n    └── auth.ts\n```\n\n### 7.2 Composables Core\n\n```typescript\n// useA2UI.ts\nimport { ref, watch } from 'vue';\nimport type { A2UISurface, WSMessage, DataModelUpdate, A2UIUpdate } from '~/types/a2ui';\nimport { useWebSocket } from './useWebSocket';\n\nexport function useA2UI(sessionId: string) {\n  const surface = ref<A2UISurface | null>(null);\n  const dataModel = ref<A2UIDataModel | null>(null);\n  const { connect, send, close } = useWebSocket(`ws://${host}/ws/session/${sessionId}`);\n\n  function handleMessage(msg: WSMessage) {\n    switch (msg.type) {\n      case 'a2ui_full':\n        surface.value = msg.payload as A2UISurface;\n        dataModel.value = surface.value.dataModel;\n        applyAccessibilityStyles(surface.value.dataModel);\n        break;\n      case 'a2ui_update':\n        // Apply patches\n        break;\n      case 'data_model_update':\n        // Hot-reload accessibility\n        break;\n    }\n  }\n\n  connect(handleMessage);\n  return { surface, dataModel, send };\n}\n```\n\n---\n\n## 8. Manifiestos Kubernetes\n\n### 8.1 namespace.yaml\n\n```yaml\napiVersion: v1\nkind: Namespace\nmetadata:\n  name: sai-learning\n```\n\n### 8.2 configmap.yaml\n\n```yaml\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: sai-server-config\n  namespace: sai-learning\ndata:\n  DB_HOST: \"postgres-service\"\n  DB_PORT: \"5432\"\n  DB_NAME: \"sai_learning\"\n  REDIS_URL: \"redis-service:6379\"\n  MINIO_ENDPOINT: \"minio-service:9000\"\n  LLM_PROVIDER: \"openai\"\n  LLM_MODEL: \"gpt-4o\"\n  TTS_PROVIDER: \"gemini\"\n```\n\n### 8.3 secrets.yaml\n\n```yaml\napiVersion: v1\nkind: Secret\nmetadata:\n  name: sai-server-secrets\n  namespace: sai-learning\ntype: Opaque\nstringData:\n  DB_USER: \"sai_user\"\n  DB_PASSWORD: \"change-me-in-production\"\n  OPENAI_API_KEY: \"sk-...\"\n  GEMINI_API_KEY: \"ai...\"\n  MINIO_ACCESS_KEY: \"minioadmin\"\n  MINIO_SECRET_KEY: \"minioadmin\"\n  JWT_SECRET: \"change-me-jwt-secret\"\n```\n\n### 8.4 deployment-server.yaml\n\n```yaml\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: sai-server\n  namespace: sai-learning\nspec:\n  replicas: 2\n  selector:\n    matchLabels:\n      app: sai-server\n  template:\n    metadata:\n      labels:\n        app: sai-server\n    spec:\n      containers:\n      - name: sai-server\n        image: ghcr.io/your-org/sai-server:latest\n        imagePullPolicy: Always\n        ports:\n        - containerPort: 8080\n        - containerPort: 8081\n        envFrom:\n        - configMapRef:\n            name: sai-server-config\n        - secretRef:\n            name: sai-server-secrets\n        livenessProbe:\n          httpGet:\n            path: /healthz\n            port: 8080\n        readinessProbe:\n          httpGet:\n            path: /readyz\n            port: 8080\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: sai-server-service\n  namespace: sai-learning\nspec:\n  selector:\n    app: sai-server\n  ports:\n  - name: http\n    port: 80\n    targetPort: 8080\n```\n\n### 8.5 ingress.yaml\n\n```yaml\napiVersion: networking.k8s.io/v1\nkind: Ingress\nmetadata:\n  name: sai-ingress\n  namespace: sai-learning\n  annotations:\n    nginx.ingress.kubernetes.io/proxy-read-timeout: \"3600\"\n    nginx.ingress.kubernetes.io/proxy-send-timeout: \"3600\"\n    nginx.ingress.kubernetes.io/websocket-services: \"sai-server-service\"\nspec:\n  rules:\n  - host: sai-learning.local\n    http:\n      paths:\n      - path: /\n        pathType: Prefix\n        backend:\n          service:\n            name: sai-server-service\n            port:\n              number: 80\n```\n\n---\n\n## 9. Plan de Implementación por Etapas\n\n### Fase 0: Fundación (4 días)\n\n| # | Tarea |\n|---|-------|\n| 0.1 | Init módulo Go + Makefile |\n| 0.2 | Configuración con viper |\n| 0.3 | Tipos de dominio (GORM models) |\n| 0.4 | Interfaces de puertos |\n| 0.5 | Protocolo A2UI tipos Go |\n| 0.6 | Router Gin + middleware |\n| 0.7 | main.go wire-up |\n| 0.8 | Migraciones SQL |\n| 0.9 | Dockerfile + docker-compose |\n\n### Fase 1: Adapters + A2UI Engine (6 días, paralelos)\n\n**Agente B**: A2UI Engine + Handlers core\n**Agente C**: Adapters de IA (OpenAI, Gemini TTS, DALL·E, MinIO)\n\n### Fase 2: Módulos de Aprendizaje (8 días, paralelos)\n\n**Agente D**: Cold-Start Diagnostic\n**Agente E**: Quiz Engine (85% Rule)\n**Agente F**: Dual Coding Orchestrator\n\n### Fase 3: Session Pipeline + Socratic (6 días, paralelos)\n\n**Agente G**: Session Manager\n**Agente H**: Socratic Remediation\n\n### Fase 4: Frontend + K8s + Testing (6 días, paralelos)\n\n**Agente I**: Frontend Nuxt 3\n**Agente J**: K8s + CI/CD + Tests\n\n---\n\n## 10. Ejemplo de main.go\n\n```go\npackage main\n\nimport (\n    \"log/slog\"\n    \"os\"\n\n    \"github.com/gin-gonic/gin\"\n    \"gorm.io/driver/sqlite\"\n    \"gorm.io/driver/postgres\"\n    \"gorm.io/gorm\"\n    \"github.com/your-org/sai-server/internal/config\"\n    \"github.com/your-org/sai-server/internal/domain\"\n    \"github.com/your-org/sai-server/internal/port\"\n    \"github.com/your-org/sai-server/internal/adapter/llm\"\n    \"github.com/your-org/sai-server/internal/adapter/tts\"\n    \"github.com/your-org/sai-server/internal/adapter/storage\"\n    \"github.com/your-org/sai-server/internal/adapter/cache\"\n    \"github.com/your-org/sai-server/internal/app/user\"\n    \"github.com/your-org/sai-server/internal/app/session\"\n    \"github.com/your-org/sai-server/internal/app/coldstart\"\n    \"github.com/your-org/sai-server/internal/app/quiz\"\n    \"github.com/your-org/sai-server/internal/app/dualcode\"\n    \"github.com/your-org/sai-server/internal/app/socratic\"\n    \"github.com/your-org/sai-server/internal/api/handlers\"\n    a2uiEngine \"github.com/your-org/sai-server/pkg/a2ui\"\n)\n\nfunc main() {\n    cfg := config.Load()\n\n    var dialector gorm.Dialector\n    if cfg.IsProduction() {\n        dialector = postgres.Open(cfg.DatabaseURL)\n    } else {\n        dialector = sqlite.Open(\"data/sai_dev.db\")\n    }\n    db, err := gorm.Open(dialector, &gorm.Config{})\n    if err != nil {\n        slog.Error(\"failed to connect database\", \"error\", err)\n        os.Exit(1)\n    }\n\n    db.AutoMigrate(\n        &domain.User{},\n        &domain.Concept{},\n        &domain.Session{},\n        &domain.Capsule{},\n        &domain.QuizItem{},\n        &domain.Interaction{},\n    )\n\n    var llmClient port.LLMClient\n    if cfg.LLMProvider == \"openai\" {\n        llmClient = llm.NewOpenAI(cfg.OpenAIKey, cfg.LLMModel)\n    }\n\n    var ttsClient port.TTSClient\n    if cfg.TTSProvider == \"gemini\" {\n        ttsClient = tts.NewGemini(cfg.GeminiAPIKey)\n    }\n\n    storageClient := storage.NewMinIO(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey)\n    cacheClient := cache.NewRedis(cfg.RedisURL)\n\n    a2ui := a2uiEngine.NewEngine()\n\n    userSvc := user.NewService(db)\n    coldStartSvc := coldstart.NewService(db, cacheClient)\n    quizEngine := quiz.NewEngine(db, coldStartSvc)\n    dualCodeOrch := dualcode.NewOrchestrator(llmClient, ttsClient, storageClient, a2ui)\n    socraticRem := socratic.NewRemediator(llmClient)\n    sessionSvc := session.NewService(db, coldStartSvc, dualCodeOrch, quizEngine, socraticRem)\n\n    r := gin.Default()\n    r.Use(handlers.CORS(), handlers.Logger())\n\n    handlers.RegisterRoutes(r, &handlers.Handlers{\n        User:      handlers.NewUserHandler(userSvc),\n        Session:   handlers.NewSessionHandler(sessionSvc),\n        ColdStart: handlers.NewColdStartHandler(coldStartSvc),\n        Quiz:      handlers.NewQuizHandler(quizEngine),\n        Capsule:   handlers.NewCapsuleHandler(dualCodeOrch),\n        Socratic:  handlers.NewSocraticHandler(socraticRem),\n        A2UI:      handlers.NewA2UIHandler(sessionSvc, a2ui),\n    })\n\n    slog.Info(\"Starting server on :8080\")\n    r.Run(\":8080\")\n}\n```\n\n---\n\n## 11. Decisiones Arquitectónicas Clave\n\n| Decisión | Justificación |\n|----------|---------------|\n| Monolito Go en lugar de microservicios | PoC requiere velocidad, no complejidad innecesaria |\n| SQLite local + PostgreSQL prod | Desarrollo rápido, producción robusta |\n| WebSocket para A2UI | Actualizaciones en tiempo real sin polling |\n| Nuxt 3 separado del backend | Independencia de equipos, SSR/SSG opcional |\n| IRT para cold-start | Estándar académico validado, sin sesgo demográfico |\n| Algoritmo 85% para quiz | Basado en investigación de \"desirable difficulty\" |\n| MinIO para assets | S3-compatible, desplegable en K8s |\n\n---\n\n*Documento generado para la PoC del Sistema de Estudio Adaptativo Multimodal*"
}
```

**Output:**
```
Wrote file successfully.
```

---

## Assistant (Build · MiniMax M2.5 Free · 2.3s)



Archivo `architecture_plan.md` guardado exitosamente en `/home/davidyesc/projects/study-sessions-with-ai/architecture_plan.md`.

---

