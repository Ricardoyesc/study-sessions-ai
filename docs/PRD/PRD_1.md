

# PRD 1 of 4 — Capsule (Controlled Generative UI)

**Branch:** `feat/capsule-controlled`
**Owner:** 1 persona (la más cómoda con Vue/Nuxt + Go del equipo)
**Time budget:** 4 horas wall-clock, dividido en 4 bloques de 1 hora
**Dependencies:** ninguna. Esta es la branch que arranca primero porque las otras 3 dependen del scaffolding que esta deja en main.

---

## 1. Objetivo

Construir el primer momento del demo: una **cápsula de estudio multimodal** que demuestra Generative UI **Controlada** (`useComponent` pattern). El agente recibe un tema de estudio, decide usar el componente `<StudyCapsule />`, y le pasa props (título, texto, imagen, audio). La UI nunca cambia de forma; solo cambia de contenido.

**Criterio de éxito:** En el video del demo, el narrador dice "el agente eligió la cápsula y la rellenó con datos sobre la doble rendija", y en pantalla se ve una tarjeta lindísima con texto + imagen generada + botón de play de audio. La tarjeta se renderiza idéntica en estructura cada vez, con contenido distinto según el tema.

---

## 2. Scope (qué SÍ y qué NO)

**SÍ:**
- Scaffold de **dos repos hermanos**: `sai-web/` (Nuxt 3 + Vue 3 + TypeScript) y `sai-server/` (Go monolito + Gin) que serán la base para las otras 3 branches.
- Endpoint `POST /api/capsules/generate` (Go handler) que recibe `{topic, sessionId}` y devuelve `{kind: "controlled", component: "StudyCapsule", props: {...}, modalities: [...]}`.
- Componente Vue `<StudyCapsule.vue />` con diseño production-grade (Tailwind via `@nuxtjs/tailwindcss` + `shadcn-vue`).
- Una sola cápsula hardcoded: tema "Doble Rendija" (Quantum Mechanics).
- Imagen pre-generada con DALL·E 3 o Gemini, guardada en `sai-server/assets/capsule-doble-rendija.png`.
- Audio pre-generado con Gemini TTS o Web Speech API en vivo, guardada en `sai-server/assets/capsule-doble-rendija.mp3`.
- Estado global mínimo con Pinia (un solo store en `sai-web/stores/demo.ts`) que las otras 3 branches van a importar.

**NO:**
- Persistencia (sin localStorage, sin DB).
- Auth.
- Múltiples cápsulas. Una sola, hardcoded, perfecta.
- Generar imagen/audio en runtime. Pre-generadas, cacheadas, listas.
- Streaming de tokens. Respuesta JSON completa de un solo shot.
- IRT, BKT, SM-2, ningún algoritmo pedagógico real. El agente solo decide "usar StudyCapsule con este tema".

---

## 3. Tech stack exacto (Nuxt 3 frontend + Go backend, alineado con ARCHITECTURE §6 §7)

### Frontend — `sai-web/`

| Capa | Tecnología | Versión | Razón |
|---|---|---|---|
| Framework | Nuxt 3 | 3.x | SSR/SSG opcional, file-based routing, alineado con arch §7 |
| Lenguaje | TypeScript | 5.x | Strict mode |
| UI Layer | Vue 3 (Composition API + `<script setup>`) | 3.x | SFC + composables |
| Estilos | Tailwind CSS | 3.x via `@nuxtjs/tailwindcss` | Tokens consistentes |
| Componentes | shadcn-vue | latest | `Card`, `Button`, `Skeleton`, `Badge` |
| Estado | Pinia | 2.x | Store oficial de Nuxt 3 |
| Iconos | `lucide-vue-next` | latest | |
| Audio | HTML5 `<audio>` nativo | — | Sin player custom |

### Backend — `sai-server/` (Go monolito)

| Capa | Tecnología | Versión | Razón |
|---|---|---|---|
| Lenguaje | Go | 1.22+ | Arch §1.2 |
| HTTP | Gin | latest | Router + middleware ligero |
| ORM | GORM | latest | Postgres prod / SQLite dev (arch §10) |
| LLM SDK | `github.com/sashabaranov/go-openai` | latest | GPT-4o (arch §1.2) — **no** Anthropic en arch oficial |
| LLM modelo | `gpt-4o` | — | Arch §8.2 ConfigMap |
| Imágenes | DALL·E 3 vía OpenAI client | — | Arch §1.2 |
| Audio TTS | Google Gemini TTS | v1 | Arch §1.2 (en runtime o pre-gen para PoC) |
| Storage assets | filesystem local (dev) / MinIO (prod) | — | Arch §1.2 |
| Logging | `log/slog` | stdlib | |

### Imagen / audio pre-generados
Para PoC ambos quedan en `sai-server/assets/` (servidos vía `GET /api/assets/{type}/{filename}`, arch §2.1). Se generan una vez con scripts en `sai-server/scripts/gen-assets.go` o playground externo y se commitean.

**No usar:** Next.js, CopilotKit, A2UI library externa, LangChain, LangGraph, Vercel AI SDK, Anthropic SDK. **Postgres/Redis/Docker quedan opcionales en PoC** (SQLite + sin cache + `go run` directo).

---

## 4. Variables de entorno

**`sai-server/.env`** (cargado vía `godotenv` o export shell):
```
OPENAI_API_KEY=sk-...
GEMINI_API_KEY=...           # opcional, solo si TTS en runtime
PORT=8080
ASSETS_DIR=./assets
```

**`sai-web/.env`**:
```
NUXT_PUBLIC_API_BASE=http://localhost:8080
```

Eso es todo. Sin Postgres, sin Redis, sin nada más en la branch 1. SQLite/in-memory OK si hace falta state.

---

## 5. Estructura de archivos a crear

### `sai-web/` (Nuxt 3)

```
sai-web/
├── app.vue                           # Root, layout + <NuxtPage />
├── pages/
│   └── index.vue                     # Landing del demo, monta <DemoFlow />
├── components/
│   ├── DemoFlow.vue                  # Orquestador (placeholder para branches 2 y 3)
│   ├── StudyCapsule.vue              # ⭐ Componente controlled
│   └── ui/                           # shadcn-vue (Card, Button, Skeleton, Badge)
├── stores/
│   └── demo.ts                       # Pinia store compartido
├── types/
│   └── genui.ts                      # Mirror de structs Go (GenUIResponse union)
├── composables/
│   └── useApi.ts                     # `$fetch` wrapper, base = NUXT_PUBLIC_API_BASE
├── assets/css/tailwind.css
├── nuxt.config.ts
├── tailwind.config.ts
├── tsconfig.json
├── package.json
└── .env
```

### `sai-server/` (Go + Gin)

```
sai-server/
├── cmd/
│   └── server/
│       └── main.go                   # Bootstrap Gin, registra rutas
├── internal/
│   ├── handlers/
│   │   └── capsules.go               # POST /api/capsules/generate (arch §2.1)
│   ├── llm/
│   │   └── openai.go                 # Cliente go-openai singleton
│   ├── prompts/
│   │   └── capsule.go                # System prompt cápsula
│   ├── types/
│   │   └── genui.go                  # GenUIResponse, CapsuleProps, Modality
│   └── assets/
│       └── handler.go                # GET /api/assets/:type/:filename
├── assets/
│   ├── capsule-doble-rendija.png     # Pre-generada
│   └── capsule-doble-rendija.mp3     # Pre-generada
├── scripts/
│   └── gen-assets.go                 # One-off generator
├── go.mod
├── go.sum
└── .env
```

---

## 6. Contratos exactos (las 3 personas siguientes dependen de esto)

### 6.1 Tipos compartidos

**Backend Go canonical** — `sai-server/internal/types/genui.go`:

```go
package types

type GenUIResponse struct {
    Kind       string         `json:"kind"` // "controlled" | "declarative" | "open"
    Component  string         `json:"component,omitempty"`
    CapsuleID  string         `json:"capsuleId,omitempty"`
    Topic      string         `json:"topic,omitempty"`
    Modalities []Modality     `json:"modalities,omitempty"`
    Props      *CapsuleProps  `json:"props,omitempty"`
    Surface    *A2UISurface   `json:"surface,omitempty"` // Branch 2 lo llena
    HTML       string         `json:"html,omitempty"`    // Branch 3 lo llena
}

type CapsuleProps struct {
    Topic                string `json:"topic"`
    Title                string `json:"title"`               // max 60 chars
    Body                 string `json:"body"`                // markdown OK, max 400 chars
    ImageURL             string `json:"imageUrl"`
    AudioURL             string `json:"audioUrl"`
    EstimatedReadTimeSec int    `json:"estimatedReadTimeSec"`
}

// Dual coding (ARCHITECTURE §4.3)
type Modality struct {
    Type     string                 `json:"type"`    // "text"|"audio"|"image"|"video"
    Content  string                 `json:"content"` // texto plano o URL del asset
    Metadata map[string]interface{} `json:"metadata,omitempty"`
}
```

**Frontend TS mirror** — `sai-web/types/genui.ts`:

```typescript
export type GenUIResponse =
  | { kind: "controlled"; component: "StudyCapsule"; capsuleId: string; topic: string; modalities: Modality[]; props: CapsuleProps }
  | { kind: "declarative"; surface: A2UISurface }   // Branch 2
  | { kind: "open"; html: string };                  // Branch 3

export type CapsuleProps = {
  topic: string;
  title: string;
  body: string;
  imageUrl: string;
  audioUrl: string;
  estimatedReadTimeSec: number;
};

export type Modality = {
  type: "text" | "audio" | "image" | "video";
  content: string;
  metadata?: Record<string, unknown>;
};

export type DemoStep = "capsule" | "quiz" | "remediation" | "done";
```

**Pinia store** — `sai-web/stores/demo.ts`:

```typescript
import { defineStore } from "pinia";
import type { CapsuleProps, DemoStep } from "~/types/genui";

export const useDemoStore = defineStore("demo", {
  state: () => ({
    step: "capsule" as DemoStep,
    capsule: null as CapsuleProps | null,
  }),
  actions: {
    setStep(s: DemoStep) { this.step = s; },
    setCapsule(c: CapsuleProps) { this.capsule = c; },
  },
});
```

### 6.2 Endpoint `POST /api/capsules/generate`

> Ruta alineada con ARCHITECTURE.md §2.1. Handler Go en `sai-server/internal/handlers/capsules.go`, registrado vía `r.POST("/api/capsules/generate", handlers.GenerateCapsule)` en `cmd/server/main.go`. Frontend Nuxt llama vía `$fetch` contra `NUXT_PUBLIC_API_BASE`. Branches 2 y 3 agregan handlers hermanos.

**Request:**
```json
{ "topic": "doble-rendija", "sessionId": "demo" }
```

**Response (controlled, dual coding §4.3):**
```json
{
  "kind": "controlled",
  "component": "StudyCapsule",
  "capsuleId": "cap-doble-rendija",
  "topic": "doble-rendija",
  "modalities": [
    { "type": "text",  "content": "Cuando disparamos electrones uno por uno...", "metadata": { "lang": "es" } },
    { "type": "image", "content": "/capsule-doble-rendija.png", "metadata": { "altText": "Diagrama doble rendija" } },
    { "type": "audio", "content": "/capsule-doble-rendija.mp3", "metadata": { "durationSec": 45 } }
  ],
  "props": {
    "topic": "doble-rendija",
    "title": "El experimento de la doble rendija",
    "body": "Cuando disparamos electrones uno por uno...",
    "imageUrl": "/capsule-doble-rendija.png",
    "audioUrl": "/capsule-doble-rendija.mp3",
    "estimatedReadTimeSec": 45
  }
}
```

`modalities[]` espeja la struct `DualCapsule.Modalities` de ARCHITECTURE §4.3. Validación: `modalities.length >= 2` (Teoría de Codificación Dual). En la PoC los 3 vienen siempre; el LLM solo rellena `props.title/body/estimatedReadTimeSec`. El sistema añade `imageUrl/audioUrl/topic` y arma `modalities[]`.

**No cambien la forma del envelope** (`kind`, `component`, `props`); solo agreguen variantes al union type. Branches 2/3 extienden con `kind: "declarative" | "open"`.

---

## 7. System prompt exacto para LLM (cápsula)

`sai-server/internal/prompts/capsule.go`:

```go
package prompts

const CapsuleSystemPrompt = `You are an adaptive study agent. Your job is to choose a UI component and fill it with educational content.

For this request, you MUST:
1. Choose component "StudyCapsule" (controlled UI).
2. Generate props with these exact fields: title, body, estimatedReadTimeSec.
3. The other fields (imageUrl, audioUrl, topic) will be filled by the system.

Constraints:
- title: max 60 chars, no emojis, Spanish.
- body: 200-400 chars, markdown allowed, conversational, second person ("imagina que..."), Spanish.
- estimatedReadTimeSec: realistic estimate based on body length, ~3 chars/second.

Topic: "%s"

Respond ONLY with valid JSON matching this shape:
{
  "title": string,
  "body": string,
  "estimatedReadTimeSec": number
}

No prose, no markdown fences, no explanations. JSON only.`
```

El handler Go (`capsules.go`) inyecta `imageUrl`, `audioUrl`, `topic` después porque son hardcoded. Si la respuesta del LLM trae fences markdown, hacer strip antes de `json.Unmarshal`.

---

## 8. Diseño visual del componente `StudyCapsule`

- Card centrada, max-width 640px, rounded-2xl, shadow-xl.
- Header: `topic` chip (badge) arriba a la izquierda, "Cápsula 1 de 1" arriba a la derecha.
- Imagen full-width, aspect 16:9, rounded top corners.
- Título en h2, 24px, font-semibold.
- Body en prose, 16px, leading-relaxed, color foreground/80.
- Footer: botón ghost "Escuchar" (lucide `Volume2`), tiempo estimado a la derecha "~45s".
- Botón primary "Continuar al quiz" abajo, fullwidth — **dispara `setStep("quiz")` en el Pinia store**.
- Skeleton loader mientras `POST /api/capsules/generate` resuelve.

Inspiración: el card de Linear que muestra un issue. Limpio, denso, no decorativo.

---

## 9. Plan minuto a minuto

**Hora 1 — Scaffold dual repo (00:00–01:00)**
- Frontend: `npx nuxi@latest init sai-web` → `cd sai-web` → `npm i` → `npm i -D @nuxtjs/tailwindcss @pinia/nuxt` → `npm i lucide-vue-next` → `npx shadcn-vue@latest init` → add `card button skeleton badge`.
- Backend: `mkdir sai-server && cd sai-server && go mod init github.com/<org>/sai-server` → `go get github.com/gin-gonic/gin github.com/sashabaranov/go-openai github.com/joho/godotenv`.
- Crear `sai-server/internal/types/genui.go` y `sai-web/types/genui.ts` con los tipos del §6.1.
- Push inicial a `main`. **Las branches 2, 3, 4 deben hacer rebase sobre este commit antes de empezar.**

**Hora 2 — Endpoint Go y prompt (01:00–02:00)**
- Crear `sai-server/cmd/server/main.go` con bootstrap Gin + CORS para `localhost:3000`.
- Crear `sai-server/internal/handlers/capsules.go` (handler `POST /api/capsules/generate`, ruta espeja arch §2.1). Branches 2/3 crean handlers hermanos (`/api/sessions/:id/next`, `/api/sessions/:id/quiz/answer`, `/api/sessions/:id/socratic/response`) — cero conflictos de merge.
- Crear `sai-server/internal/prompts/capsule.go`.
- `go run ./cmd/server` → probar con `curl -X POST localhost:8080/api/capsules/generate -d '{"topic":"doble-rendija","sessionId":"demo"}' -H "Content-Type: application/json"` → debe devolver JSON válido con `modalities.length >= 2`.

**Hora 3 — Componentes Vue y home (02:00–03:00)**
- Crear `sai-web/components/StudyCapsule.vue` siguiendo el §8.
- Crear `sai-web/components/DemoFlow.vue` que lea `step` del Pinia store y renderice `<StudyCapsule />` cuando `step === "capsule"`. Cuando branches 2 y 3 mergeen, agregarán sus componentes al switch.
- En `sai-web/pages/index.vue`: hero + `<DemoFlow />`. Llamar `$fetch` contra `NUXT_PUBLIC_API_BASE + "/api/capsules/generate"` en `onMounted`.
- Verificar visual en navegador, ajustar Tailwind.

**Hora 4 — Pre-generación de assets y polish (03:00–04:00)**
- Generar imagen en OpenAI Images Playground o Gemini AI Studio. Prompt sugerido: "Cinematic illustration of the double-slit experiment, photons hitting a screen forming an interference pattern, dark background, glowing particles, scientific style, no text, 16:9". Descargar como `capsule-doble-rendija.png`, copiar a `sai-server/assets/`.
- Audio: lo más rápido es **Web Speech API en runtime en el browser** (`new SpeechSynthesisUtterance(body)` desde Vue). Si hay tiempo, generar mp3 con Gemini TTS y guardar en `sai-server/assets/`.
- Servir assets vía `r.Static("/api/assets", "./assets")` (arch §2.1).
- Smoke test: `go run ./cmd/server` + `npm run dev` (sai-web), reload 5 veces, verificar que la cápsula siempre carga bien.
- Push final a `main`.

---

## 10. Definition of Done

- [ ] `go run ./cmd/server` (sai-server) y `npm run dev` (sai-web) levantan sin errores.
- [ ] `localhost:3000` muestra la cápsula con imagen, texto, audio button funcional.
- [ ] Tipos en `sai-server/internal/types/genui.go` y `sai-web/types/genui.ts` exportados y son los del §6.1 exactos (mismos field names, mismo JSON shape).
- [ ] El endpoint `POST /api/capsules/generate` (Go handler) responde el envelope del §6.2 exacto, con `modalities >= 2` (texto + imagen + audio).
- [ ] El botón "Continuar al quiz" cambia el step en el Pinia store a `"quiz"` (aunque renderice un placeholder "Quiz coming from branch 2").
- [ ] `main` tiene un commit `chore: scaffold sai-web + sai-server + capsule (controlled UI)` que las otras 3 branches pueden rebasear.
- [ ] El README tiene una sección "Branch 1 — Capsule" con un screenshot.

---

## 11. Riesgos y fallbacks

- **LLM devuelve JSON inválido.** Fallback: hardcodear `body` y `title` en `capsules.go` y solo llamar al LLM si `os.Getenv("OPENAI_API_KEY") != ""`. El demo nunca debe romperse por una llamada al LLM.
- **shadcn-vue falla en init.** Fallback: HTML + Tailwind crudo. Es feo pero funciona.
- **La imagen pre-generada tarda en aparecer.** Fallback: imagen de Unsplash con search term "physics quantum" cacheada localmente.
- **El audio no suena en el demo.** Fallback: omitir audio del demo, no es bloqueante. La narrativa puede decir "imagine podríamos también incluir audio TTS" y dejarlo como future work.

---

## 12. Handoff a las otras branches

Al terminar, esta branch deja en `main`:

1. El scaffold completo de ambos repos (las otras 3 nunca tocan `package.json`, `nuxt.config.ts`, `tailwind.config.ts`, `go.mod`, `cmd/server/main.go`).
2. `sai-server/internal/types/genui.go` + `sai-web/types/genui.ts` con los tipos del demo. Las otras 3 branches solo extienden el union, no lo modifican.
3. Handlers Go alineados con ARCHITECTURE §2.1: `internal/handlers/capsules.go` (esta branch). Branch 2 agrega `/api/sessions/:id/next` + `/api/sessions/:id/quiz/answer`. Branch 3 agrega `/api/sessions/:id/socratic/response`.
4. `sai-web/components/DemoFlow.vue` con el switch sobre `step`. Branches 2 y 3 agregan sus componentes al switch.
5. Un mensaje en Slack/Discord al equipo: "main listo, hagan rebase y arranquen branches 2-4".
