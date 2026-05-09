

# PRD 1 of 4 — Capsule (Controlled Generative UI)

**Branch:** `feat/capsule-controlled`
**Owner:** 1 persona (la más cómoda con React/Next.js del equipo)
**Time budget:** 4 horas wall-clock, dividido en 4 bloques de 1 hora
**Dependencies:** ninguna. Esta es la branch que arranca primero porque las otras 3 dependen del scaffolding que esta deja en main.

---

## 1. Objetivo

Construir el primer momento del demo: una **cápsula de estudio multimodal** que demuestra Generative UI **Controlada** (`useComponent` pattern). El agente recibe un tema de estudio, decide usar el componente `<StudyCapsule />`, y le pasa props (título, texto, imagen, audio). La UI nunca cambia de forma; solo cambia de contenido.

**Criterio de éxito:** En el video del demo, el narrador dice "el agente eligió la cápsula y la rellenó con datos sobre la doble rendija", y en pantalla se ve una tarjeta lindísima con texto + imagen generada + botón de play de audio. La tarjeta se renderiza idéntica en estructura cada vez, con contenido distinto según el tema.

---

## 2. Scope (qué SÍ y qué NO)

**SÍ:**
- Scaffold del proyecto Next.js 15 (App Router) que será la base para las otras 3 branches.
- Endpoint `/api/agent` que recibe `{topic: string}` y devuelve `{component: "StudyCapsule", props: {...}}`.
- Componente `<StudyCapsule />` con diseño production-grade (Tailwind + shadcn/ui).
- Una sola cápsula hardcoded: tema "Doble Rendija" (Quantum Mechanics).
- Imagen pre-generada con DALL-E o Gemini, guardada en `/public/capsule-doble-rendija.png`.
- Audio pre-generado con ElevenLabs free tier o Web Speech API en vivo, guardado en `/public/capsule-doble-rendija.mp3`.
- Estado global mínimo con Zustand (un solo store) que las otras 3 branches van a importar.

**NO:**
- Persistencia (sin localStorage, sin DB).
- Auth.
- Múltiples cápsulas. Una sola, hardcoded, perfecta.
- Generar imagen/audio en runtime. Pre-generadas, cacheadas, listas.
- Streaming de tokens. Respuesta JSON completa de un solo shot.
- IRT, BKT, SM-2, ningún algoritmo pedagógico real. El agente solo decide "usar StudyCapsule con este tema".

---

## 3. Tech stack exacto

| Capa | Tecnología | Versión | Razón |
|---|---|---|---|
| Framework | Next.js | 15.x | App Router, Vercel deploy en 1 comando |
| Lenguaje | TypeScript | 5.x | Strict mode |
| LLM SDK | `@anthropic-ai/sdk` | latest | Claude Sonnet 4.5 vía API key directo. Sin Vercel AI SDK para evitar bugs de versión. |
| LLM modelo | `claude-sonnet-4-5` | — | Más barato que Opus, suficiente para "elegir componente y props" |
| Estilos | Tailwind CSS | 4.x | Ya viene con `create-next-app` |
| Componentes | shadcn/ui | latest | `Card`, `Button`, `Skeleton` |
| Estado | Zustand | 5.x | Un solo store global |
| Iconos | `lucide-react` | latest | |
| Audio | HTML5 `<audio>` nativo | — | No reproductor custom |
| Imagen pre-gen | DALL-E 3 vía playground o Gemini Image | — | UNA imagen, fuera de runtime |

**No usar:** Vercel AI SDK, CopilotKit, A2UI library, LangChain, LangGraph, Postgres, Redis, Docker. Todo eso suma minutos de setup que no tienen.

---

## 4. Variables de entorno

`.env.local`:
```
ANTHROPIC_API_KEY=sk-ant-...
```

Eso es todo. Sin Postgres, sin Redis, sin nada más en la branch 1.

---

## 5. Estructura de archivos a crear

```
study-sessions-ai/
├── app/
│   ├── layout.tsx                    # Layout raíz, fonts, metadata
│   ├── page.tsx                      # Landing del demo, llama a <DemoFlow />
│   ├── globals.css                   # Tailwind base + variables de tema
│   └── api/
│       └── capsules/
│           └── generate/
│               └── route.ts          # POST /api/capsules/generate (arch §2.1)
├── components/
│   ├── DemoFlow.tsx                  # Orquestador del demo (placeholder para branches 2 y 3)
│   ├── StudyCapsule.tsx              # ⭐ El componente controlled
│   └── ui/                           # shadcn (card, button, skeleton)
├── lib/
│   ├── anthropic.ts                  # Cliente Claude singleton
│   ├── store.ts                      # Zustand store compartido
│   └── prompts/
│       └── capsule.ts                # System prompt para la cápsula
├── public/
│   ├── capsule-doble-rendija.png     # Pre-generada
│   └── capsule-doble-rendija.mp3     # Pre-generada
├── package.json
├── tsconfig.json
├── tailwind.config.ts
├── next.config.ts
└── .env.local
```

---

## 6. Contratos exactos (las 3 personas siguientes dependen de esto)

### 6.1 Tipo compartido en `lib/store.ts`

```typescript
export type GenUIResponse =
  | { kind: "controlled"; component: "StudyCapsule"; capsuleId: string; topic: string; modalities: Modality[]; props: CapsuleProps }
  | { kind: "declarative"; surface: A2UISurface }   // Branch 2 lo llena (ver §3.1 ARCHITECTURE)
  | { kind: "open"; html: string };                  // Branch 3 lo llena

export type CapsuleProps = {
  topic: string;
  title: string;
  body: string;             // Markdown OK, max 400 chars
  imageUrl: string;
  audioUrl: string;
  estimatedReadTimeSec: number;
};

// Dual coding (ARCHITECTURE §4.3 — DualCapsule.Modalities)
export type Modality = {
  type: "text" | "audio" | "image" | "video";
  content: string;          // texto plano o URL del asset
  metadata?: Record<string, unknown>;
};

export type DemoStep = "capsule" | "quiz" | "remediation" | "done";

interface DemoStore {
  step: DemoStep;
  capsule: CapsuleProps | null;
  setStep: (s: DemoStep) => void;
  setCapsule: (c: CapsuleProps) => void;
}
```

### 6.2 Endpoint `POST /api/capsules/generate`

> Ruta alineada con ARCHITECTURE.md §2.1. El backend Go futuro expone exactamente este path. Para la PoC en Next.js implementamos el handler en `app/api/capsules/generate/route.ts`. Branches 2 y 3 usan rutas hermanas (ver sus PRDs).

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

## 7. System prompt exacto para Claude (cápsula)

`lib/prompts/capsule.ts`:

```typescript
export const CAPSULE_SYSTEM_PROMPT = `You are an adaptive study agent. Your job is to choose a UI component and fill it with educational content.

For this request, you MUST:
1. Choose component "StudyCapsule" (controlled UI).
2. Generate props with these exact fields: title, body, estimatedReadTimeSec.
3. The other fields (imageUrl, audioUrl, topic) will be filled by the system.

Constraints:
- title: max 60 chars, no emojis, Spanish.
- body: 200-400 chars, markdown allowed, conversational, second person ("imagina que..."), Spanish.
- estimatedReadTimeSec: realistic estimate based on body length, ~3 chars/second.

Topic: "${"{topic}"}"

Respond ONLY with valid JSON matching this shape:
{
  "title": string,
  "body": string,
  "estimatedReadTimeSec": number
}

No prose, no markdown fences, no explanations. JSON only.`;
```

El handler en `route.ts` mete los `imageUrl`, `audioUrl`, `topic` después porque son hardcoded.

---

## 8. Diseño visual del componente `StudyCapsule`

- Card centrada, max-width 640px, rounded-2xl, shadow-xl.
- Header: `topic` chip (badge) arriba a la izquierda, "Cápsula 1 de 1" arriba a la derecha.
- Imagen full-width, aspect 16:9, rounded top corners.
- Título en h2, 24px, font-semibold.
- Body en prose, 16px, leading-relaxed, color foreground/80.
- Footer: botón ghost "Escuchar" (lucide `Volume2`), tiempo estimado a la derecha "~45s".
- Botón primary "Continuar al quiz" abajo, fullwidth — **dispara `setStep("quiz")` en el Zustand**.
- Skeleton loader mientras `/api/agent` resuelve.

Inspiración: el card de Linear que muestra un issue. Limpio, denso, no decorativo.

---

## 9. Plan minuto a minuto

**Hora 1 — Scaffold (00:00–01:00)**
- `npx create-next-app@latest study-sessions-ai --typescript --tailwind --app --src-dir=false --import-alias="@/*"`
- `npm install @anthropic-ai/sdk zustand lucide-react`
- `npx shadcn@latest init` → `npx shadcn@latest add card button skeleton badge`
- Crear `lib/anthropic.ts`, `lib/store.ts` con los tipos del §6.1.
- Push inicial a `main`. **Las branches 2, 3, 4 deben hacer rebase sobre este commit antes de empezar.**

**Hora 2 — Endpoint y prompt (01:00–02:00)**
- Crear `app/api/capsules/generate/route.ts` (handler dedicado, ruta espeja arch §2.1). Branches 2/3 crean rutas hermanas (`/api/sessions/[id]/quiz/answer`, `/api/sessions/[id]/socratic/response`) — cero conflictos de merge.
- Crear `lib/prompts/capsule.ts`.
- Probar con `curl -X POST localhost:3000/api/capsules/generate -d '{"topic":"doble-rendija","sessionId":"demo"}' -H "Content-Type: application/json"` → debe devolver JSON válido con `modalities.length >= 2`.
- Si Claude devuelve markdown fences, hacer strip antes de `JSON.parse`.

**Hora 3 — Componente y home (02:00–03:00)**
- Crear `components/StudyCapsule.tsx` siguiendo el §8.
- Crear `components/DemoFlow.tsx` que lea `step` del store y renderice `<StudyCapsule />` cuando `step === "capsule"`. Cuando branches 2 y 3 mergeen, agregarán sus componentes al switch.
- En `app/page.tsx`: hero con título del producto + `<DemoFlow />`.
- Verificar visual en el navegador, ajustar Tailwind.

**Hora 4 — Pre-generación de assets y polish (03:00–04:00)**
- Ir a https://platform.openai.com/playground/images o Gemini AI Studio, generar la imagen del experimento de doble rendija. Prompt sugerido: "Cinematic illustration of the double-slit experiment, photons hitting a screen forming an interference pattern, dark background, glowing particles, scientific style, no text, 16:9". Descargar como `capsule-doble-rendija.png`, copiar a `/public`.
- Audio: la opción más rápida es **Web Speech API en runtime** (`new SpeechSynthesisUtterance(body)`) en lugar de pre-grabar. Si tienen ElevenLabs, generar mp3 con voz "Bella" o similar y guardarlo en `/public`.
- Smoke test: `npm run dev`, abrir el browser, hit reload 5 veces, verificar que la cápsula siempre carga bien.
- Push final a `main`.

---

## 10. Definition of Done

- [ ] `npm run dev` levanta sin errores.
- [ ] `localhost:3000` muestra la cápsula con imagen, texto, audio button funcional.
- [ ] Tipos en `lib/store.ts` están exportados y son los del §6.1 exactos.
- [ ] El endpoint `POST /api/capsules/generate` responde el envelope del §6.2 exacto, con `modalities.length >= 2` (texto + imagen + audio).
- [ ] El botón "Continuar al quiz" cambia el step en el Zustand a `"quiz"` (aunque renderice un placeholder "Quiz coming from branch 2").
- [ ] `main` tiene un commit `chore: scaffold + capsule (controlled UI)` que las otras 3 branches pueden rebasear.
- [ ] El README tiene una sección "Branch 1 — Capsule" con un screenshot.

---

## 11. Riesgos y fallbacks

- **Claude devuelve JSON inválido.** Fallback: hardcodear el `body` y el `title` en `route.ts` y solo llamar a Claude si `process.env.ANTHROPIC_API_KEY` existe. El demo nunca debe romperse por una llamada al LLM.
- **shadcn/ui falla en init.** Fallback: HTML + Tailwind crudo. Es feo pero funciona.
- **La imagen pre-generada tarda en aparecer.** Fallback: imagen de Unsplash con search term "physics quantum" cacheada localmente.
- **El audio no suena en el demo.** Fallback: omitir audio del demo, no es bloqueante. La narrativa puede decir "imagine podríamos también incluir audio TTS" y dejarlo como future work.

---

## 12. Handoff a las otras branches

Al terminar, esta branch deja en `main`:

1. El scaffold completo (las otras 3 nunca tocan `package.json`, `next.config.ts`, `tailwind.config.ts`, `globals.css`).
2. `lib/store.ts` con los tipos del demo. Las otras 3 branches solo extienden el union type, no lo modifican.
3. Rutas alineadas con ARCHITECTURE §2.1: `app/api/capsules/generate/route.ts` (esta branch). Branch 2 agrega `/api/sessions/[id]/next` + `/api/sessions/[id]/quiz/answer`. Branch 3 agrega `/api/sessions/[id]/socratic/response`. Cliente puede swap a backend Go cambiando solo el base URL.
4. `components/DemoFlow.tsx` con el switch sobre `step`. Branches 2 y 3 agregan sus componentes al switch.
5. Un mensaje en Slack/Discord al equipo: "main listo, hagan rebase y arranquen branches 2-4".
