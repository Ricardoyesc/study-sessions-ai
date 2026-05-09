# PRD 2 of 4 — Quiz (Declarative Generative UI / A2UI)

**Branch:** `feat/quiz-a2ui`
**Owner:** 1 persona (la que tenga más tolerancia a JSON-driven UI)
**Time budget:** 4 horas wall-clock
**Dependencies:** Branch 1 mergeada en `main`. `git checkout main && git pull && git checkout -b feat/quiz-a2ui` antes de empezar.

---

## 1. Objetivo

Construir el segundo momento del demo: un quiz adaptativo donde el agente no escoge entre componentes pre-hechos, sino que emite un **schema JSON** que un mini-renderer pinta con primitivas Vue. El agente decide la **forma** del quiz según el "perfil del estudiante", no solo el contenido.

**Criterio de éxito:** En el demo, hay un toggle visible "Perfil del estudiante: [Visual | Kinestésico | Lector]". Al cambiar el toggle, el agente regenera el quiz con un schema estructuralmente distinto:

- Visual → multiple choice con imágenes en grid 2x2.
- Kinestésico → drag-and-drop para ordenar pasos.
- Lector → input de texto libre con pregunta abierta.

Los jueces ven que un mismo concepto se evalúa con UI distinta y que la UI es composición, no template-switching.

---

## 2. Scope

**SÍ:**
- Mini-renderer A2UI casero (~150 líneas TS) que mapea schema JSON → componentes Vue en `sai-web/`.
- 7 primitivas: `Stack`, `Row`, `Text`, `Image`, `Button`, `TextInput`, `DragItem`.
- Handler Go `POST /api/sessions/:id/next` que devuelve `{kind: "declarative", surface: A2UISurface}`.
- Handler Go `POST /api/sessions/:id/quiz/answer` que actualiza `thetaUser`.
- 3 prompts distintos (uno por perfil) en `sai-server/internal/prompts/quiz.go`.
- Componente `<QuizSurface.vue :surface="..." />` que renderiza el árbol y dispara eventos.
- Toggle de perfil arriba del quiz, conectado al Pinia store.
- Integración con el flujo: cuando el quiz termina (correcto o incorrecto), avanza el `step`.

**NO:**
- A2UI library oficial (a2ui.org). Implementan su propia versión mini.
- Validación exhaustiva del schema. Confían en el LLM. Try/catch alrededor del render y fallback a un quiz hardcoded si el schema viene mal.
- Más de 7 primitivas. Cualquier cosa que no esté en la lista, no se renderiza.
- Animaciones complejas, drag-drop con física. El drag es funcional, no bonito.
- Persistencia de respuestas. Solo importa el `did they answer?` boolean.

---

## 3. Tech stack adicional (sobre lo que branch 1 dejó)

| Capa | Tecnología | Versión | Razón |
|---|---|---|---|
| Drag & drop (Vue) | `vue-draggable-plus` (o nativo HTML5 drag) | latest | Sin equivalente Vue de @dnd-kit estable. Alternativa: HTML5 drag API + listas. |
| Validación schema | nada (try/catch) | — | No vale la pena Zod en 4h |
| Renderer | Componente Vue recursivo propio | — | ~150 líneas |

**Frontend:** `cd sai-web && npm i vue-draggable-plus`.
**Backend:** sin deps nuevas (sigue con `gin` + `go-openai`).

---

## 4. Contrato del schema A2UI (alineado con ARCHITECTURE §3.1)

> **Decisión:** el shape sigue la struct Go `Surface` de ARCHITECTURE §3.1 (`surfaceId`, `rootComponent`, `components` map, `dataModel`). Backend Go canonical → frontend TS mirror.

### 4.1 Backend Go — extender `sai-server/internal/types/genui.go`

```go
type A2UISurface struct {
    SurfaceID     string                   `json:"surfaceId"`
    RootComponent string                   `json:"rootComponent"`
    Components    map[string]A2UIComponent `json:"components"`
    DataModel     *A2UIDataModel           `json:"dataModel,omitempty"`
    Meta          A2UIMeta                 `json:"meta"`
}

type A2UIComponent struct {
    ID       string                 `json:"id"`
    Type     string                 `json:"type"` // Stack|Row|Text|Image|Button|TextInput|DragItem
    Children []string               `json:"children,omitempty"` // adjacency
    Props    map[string]interface{} `json:"props,omitempty"`
    Events   map[string]string      `json:"events,omitempty"`
}

type A2UIDataModel struct {
    Theme         string `json:"theme,omitempty"`         // "system"|"light"|"dark"
    FontFamily    string `json:"fontFamily,omitempty"`
    FontScale     float64 `json:"fontScale,omitempty"`
    ColorPalette  string `json:"colorPalette,omitempty"`
    HighContrast  bool   `json:"highContrast,omitempty"`
    ReducedMotion bool   `json:"reducedMotion,omitempty"`
    Language      string `json:"language,omitempty"`
}

type A2UIMeta struct {
    Intent             string         `json:"intent"` // "quiz"
    Profile            string         `json:"profile"` // visual|kinesthetic|reader
    CorrectAnswerEvent string         `json:"correctAnswerEvent"`
    IRT                *IRTTrace      `json:"irt,omitempty"`
}

type IRTTrace struct {
    ThetaUser         float64 `json:"thetaUser"`
    DifficultyIRT     float64 `json:"difficultyIRT"`
    Discrimination    float64 `json:"discrimination"`
    PSuccessPredicted float64 `json:"pSuccessPredicted"`
}
```

### 4.2 Frontend TS mirror — extender `sai-web/types/genui.ts`

```typescript
export type A2UIComponent =
  | { id: string; type: "Stack"; children: string[]; props?: { gap?: number; align?: "start"|"center"|"end" } }
  | { id: string; type: "Row"; children: string[]; props?: { gap?: number; wrap?: boolean } }
  | { id: string; type: "Text"; props: { content: string; variant?: "h2"|"h3"|"body"|"caption" } }
  | { id: string; type: "Image"; props: { url: string; alt: string; aspectRatio?: "1:1"|"16:9"|"4:3" } }
  | { id: string; type: "Button"; props: { label: string; variant?: "primary"|"secondary"|"ghost" }; events?: { onClick?: string }; eventPayload?: Record<string, unknown> }
  | { id: string; type: "TextInput"; props: { placeholder: string }; events?: { onSubmit?: string } }
  | { id: string; type: "DragItem"; props: { label: string; correctOrder: number } };

export type A2UIDataModel = {
  theme?: "system" | "light" | "dark";
  fontFamily?: string;
  fontScale?: number;
  colorPalette?: string;
  highContrast?: boolean;
  reducedMotion?: boolean;
  language?: string;
};

export type A2UISurface = {
  surfaceId: string;
  rootComponent: string;
  components: Record<string, A2UIComponent>;
  dataModel?: A2UIDataModel;
  meta: {
    intent: "quiz";
    profile: StudentProfile;
    correctAnswerEvent: string;
    irt?: { thetaUser: number; difficultyIRT: number; discrimination: number; pSuccessPredicted: number };
  };
};

export type StudentProfile = "visual" | "kinesthetic" | "reader";

export type GenUIResponse =
  | { kind: "controlled"; component: "StudyCapsule"; capsuleId: string; topic: string; modalities: Modality[]; props: CapsuleProps }
  | { kind: "declarative"; surface: A2UISurface }   // ⭐ ESTE PRD
  | { kind: "open"; html: string };
```

### 4.3 Pinia store — extender `sai-web/stores/demo.ts`

```typescript
import { defineStore } from "pinia";
import type { CapsuleProps, DemoStep, A2UISurface, StudentProfile } from "~/types/genui";

export const useDemoStore = defineStore("demo", {
  state: () => ({
    step: "capsule" as DemoStep,
    capsule: null as CapsuleProps | null,
    quizSurface: null as A2UISurface | null,
    studentProfile: "visual" as StudentProfile,
    thetaUser: 0.0,
    lastAnswerCorrect: null as boolean | null,
  }),
  actions: {
    setStep(s: DemoStep) { this.step = s; },
    setCapsule(c: CapsuleProps) { this.capsule = c; },
    setQuizSurface(s: A2UISurface) { this.quizSurface = s; },
    setProfile(p: StudentProfile) { this.studentProfile = p; },
    setAnswerResult(correct: boolean) {
      this.lastAnswerCorrect = correct;
      this.thetaUser += correct ? 0.3 : -0.3;
    },
  },
});
```

---

## 5. Estructura de archivos a crear

### Frontend `sai-web/`

```
components/quiz/
├── QuizSurface.vue          # Wrapper, hace $fetch, maneja toggle de perfil
├── A2UIRenderer.vue         # ⭐ Renderer recursivo
├── ProfileToggle.vue        # Tabs visuales
└── primitives/
    ├── A2UIStack.vue
    ├── A2UIRow.vue
    ├── A2UIText.vue
    ├── A2UIImage.vue
    ├── A2UIButton.vue
    ├── A2UITextInput.vue
    └── DragArea.vue         # Wrapper de vue-draggable-plus que recibe DragItems

composables/
└── useA2UIEventBus.ts       # window CustomEvent helpers

```

### Backend `sai-server/`

```
internal/
├── handlers/
│   ├── sessions_next.go     # POST /api/sessions/:id/next
│   └── sessions_answer.go   # POST /api/sessions/:id/quiz/answer
├── prompts/
│   └── quiz.go              # 3 system prompts (uno por perfil)
└── irt/
    └── selector.go          # pSuccess + pickDifficulty (85% rule)
assets/
├── quiz-doble-rendija-1.png # Servidos vía GET /api/assets/quiz-doble-rendija-1.png
├── quiz-doble-rendija-2.png
├── quiz-doble-rendija-3.png
└── quiz-doble-rendija-4.png
```

---

## 6. El renderer (corazón de esta branch)

`sai-web/components/quiz/A2UIRenderer.vue`:

```vue
<script setup lang="ts">
import type { A2UISurface } from "~/types/genui";
const props = defineProps<{ surface: A2UISurface }>();
</script>

<template>
  <RenderComponent :id="surface.rootComponent" :surface="surface" />
</template>
```

`RenderComponent.vue` (componente recursivo, mismo dir):

```vue
<script setup lang="ts">
import type { A2UISurface, A2UIComponent } from "~/types/genui";
import A2UIStack from "./primitives/A2UIStack.vue";
import A2UIRow from "./primitives/A2UIRow.vue";
import A2UIText from "./primitives/A2UIText.vue";
import A2UIImage from "./primitives/A2UIImage.vue";
import A2UIButton from "./primitives/A2UIButton.vue";
import A2UITextInput from "./primitives/A2UITextInput.vue";
import DragArea from "./primitives/DragArea.vue";

const props = defineProps<{ id: string; surface: A2UISurface }>();
const node = computed<A2UIComponent | undefined>(() => props.surface.components[props.id]);

const allDraggable = computed(() => {
  const n = node.value;
  if (!n || !("children" in n)) return false;
  const kids = n.children.map((c) => props.surface.components[c]).filter(Boolean);
  return kids.length > 0 && kids.every((c) => c.type === "DragItem");
});
</script>

<template>
  <template v-if="!node" />
  <DragArea v-else-if="allDraggable" :items="node.children.map((c) => surface.components[c])" :surface="surface" />
  <A2UIStack v-else-if="node.type === 'Stack'" v-bind="node.props">
    <RenderComponent v-for="cid in node.children" :key="cid" :id="cid" :surface="surface" />
  </A2UIStack>
  <A2UIRow v-else-if="node.type === 'Row'" v-bind="node.props">
    <RenderComponent v-for="cid in node.children" :key="cid" :id="cid" :surface="surface" />
  </A2UIRow>
  <A2UIText v-else-if="node.type === 'Text'" v-bind="node.props" />
  <A2UIImage v-else-if="node.type === 'Image'" v-bind="node.props" />
  <A2UIButton v-else-if="node.type === 'Button'" :id="node.id" :event-name="node.events?.onClick" v-bind="node.props" />
  <A2UITextInput v-else-if="node.type === 'TextInput'" :id="node.id" :event-name="node.events?.onSubmit" v-bind="node.props" />
</template>
```

`sai-web/composables/useA2UIEventBus.ts`:

```typescript
export function emit(eventName: string, payload?: Record<string, unknown>) {
  window.dispatchEvent(new CustomEvent(`a2ui:${eventName}`, { detail: payload }));
}

export function on(eventName: string, handler: (payload: any) => void) {
  const wrapped = (e: Event) => handler((e as CustomEvent).detail);
  window.addEventListener(`a2ui:${eventName}`, wrapped);
  return () => window.removeEventListener(`a2ui:${eventName}`, wrapped);
}
```

`<QuizSurface.vue />` escucha el `correctAnswerEvent` del meta vía `on(...)` en `onMounted`, decide correct/incorrect, llama a `setAnswerResult` y avanza el step.

---

## 7. System prompts (los 3 perfiles)

`sai-server/internal/prompts/quiz.go`:

```go
package prompts

const quizCommon = `You are a Generative UI agent. You output a JSON schema describing UI to render.

Allowed component types ONLY: Stack, Row, Text, Image, Button, TextInput, DragItem. Any other type will be ignored.

Schema format (Surface, ARCHITECTURE §3.1):
{
  "surfaceId": "quiz-doble-rendija",
  "rootComponent": "root",
  "components": {
    "root": { "id": "root", "type": "Stack", "props": {"gap": 16}, "children": ["q-text", "options"] },
    ...
  },
  "dataModel": { "language": "es", "theme": "system" },
  "meta": { "intent": "quiz", "profile": "...", "correctAnswerEvent": "..." }
}

children is an array of component IDs (string[]), NOT inline nodes (adjacency map).
Buttons emit events via events.onClick (string). TextInputs use events.onSubmit.
Keep tree depth <= 4. Keep total nodes <= 15.
Topic: "doble rendija" (double-slit experiment, quantum mechanics).
Question concept: "What happens to the interference pattern when we observe which slit each photon passes through?"
Correct answer: the pattern disappears (collapses into two bands).
All text in Spanish.

Respond ONLY with valid JSON. No prose, no fences.`

const QuizVisual = quizCommon + `

Profile: VISUAL learner.
Generate a multiple choice quiz with 4 options, EACH option being an Image (urls /quiz-doble-rendija-1.png through /quiz-doble-rendija-4.png) inside a Button with a short label.
Layout: Stack with question Text on top, then a Row of 4 Buttons (wrap: true).
The Button for the correct image must set events.onClick to "answer:correct". The other three set events.onClick to "answer:wrong".
Set meta.correctAnswerEvent to "answer:correct".`

const QuizKinesthetic = quizCommon + `

Profile: KINESTHETIC learner.
Generate a drag-and-drop ordering quiz. The student must order 4 DragItems representing the steps of the experiment in causal sequence.
The 4 DragItems with their correctOrder values:
1. "Disparamos un electrón hacia la doble rendija"
2. "Decidimos si observamos por cuál rendija pasa"
3. "El electrón impacta la pantalla detectora"
4. "Tras muchos electrones, observamos el patrón resultante"
Layout: Stack with instruction Text, then the 4 DragItem nodes as children of the Stack (renderer auto-wraps them in DragArea), then a Button "Verificar orden" with events.onClick = "order:check".
Set meta.correctAnswerEvent to "order:correct".`

const QuizReader = quizCommon + `

Profile: READER learner.
Layout: Stack with paragraph Text introducing context, h3 Text with the question, TextInput (events.onSubmit = "answer:submitted"), Button "Enviar respuesta" with events.onClick = "answer:submitted".
TextInput placeholder hints 1-2 sentences expected.
Set meta.correctAnswerEvent to "answer:submitted".`

func GetPrompt(profile string) string {
    switch profile {
    case "visual":      return QuizVisual
    case "kinesthetic": return QuizKinesthetic
    case "reader":      return QuizReader
    default:            return QuizVisual
    }
}
```

---

## 8. Endpoints (alineado con ARCHITECTURE §2.1)

### 8.1 `POST /api/sessions/:id/next`

Handler Go en `sai-server/internal/handlers/sessions_next.go`. Espeja arch §2.1 `GET /api/sessions/{id}/next` (PoC usa POST porque el body trae `profile` y `thetaUser`).

**Request:**
```json
{ "profile": "visual", "thetaUser": 0.0, "topic": "doble-rendija" }
```

**Response:** `{ "kind": "declarative", "surface": A2UISurface }`.

**Selector de dificultad (Regla del 85%, arch §4.2 simplificado para PoC)** — `sai-server/internal/irt/selector.go`:

```go
package irt

import "math"

// IRT-3PL: P_success = c + (1-c) / (1 + exp(-a*(theta-b)))
func PSuccess(theta, a, b, c float64) float64 {
    return c + (1-c)/(1+math.Exp(-a*(theta-b)))
}

type Item struct{ A, B float64 }

func PickDifficulty(theta float64) (Item, float64) {
    items := []Item{{1.0, -0.5}, {1.0, 0.0}, {1.2, 0.8}}
    best, bestP := items[0], PSuccess(theta, items[0].A, items[0].B, 0.25)
    for _, it := range items[1:] {
        p := PSuccess(theta, it.A, it.B, 0.25)
        if math.Abs(p-0.85) < math.Abs(bestP-0.85) {
            best, bestP = it, p
        }
    }
    return best, bestP
}
```

Handler pseudocode:

```go
func GenerateNext(c *gin.Context) {
    var body struct {
        Profile   string  `json:"profile"`
        ThetaUser float64 `json:"thetaUser"`
        Topic     string  `json:"topic"`
    }
    c.BindJSON(&body)
    item, p := irt.PickDifficulty(body.ThetaUser)
    prompt := prompts.GetPrompt(body.Profile) + fmt.Sprintf("\n\nDifficulty target (b): %.2f. Calibrate vocabulary/cognitive load.", item.B)

    raw, err := llm.Complete(prompt) // go-openai client
    var surface types.A2UISurface
    if err != nil || json.Unmarshal(stripFences(raw), &surface) != nil {
        surface = fallbackSurfaces[body.Profile]
    }
    surface.Meta.IRT = &types.IRTTrace{
        ThetaUser: body.ThetaUser, DifficultyIRT: item.B,
        Discrimination: item.A, PSuccessPredicted: p,
    }
    c.JSON(200, gin.H{"kind": "declarative", "surface": surface})
}
```

### 8.2 `POST /api/sessions/:id/quiz/answer`

Body: `{ "correct": boolean }`. Heurística PoC: `thetaUser += correct ? 0.3 : -0.3`. Backend Go futuro hace MLE real.

---

## 9. Plan minuto a minuto

**Hora 1 — Primitivas Vue (00:00–01:00)**
- `git checkout main && git pull && git checkout -b feat/quiz-a2ui`
- `cd sai-web && npm i vue-draggable-plus`
- Crear las 7 primitivas en `components/quiz/primitives/`. Cada una ~15 líneas, Tailwind, sin lógica de negocio.
- Crear `composables/useA2UIEventBus.ts`.

**Hora 2 — Renderer y schema types (01:00–02:00)**
- Extender `sai-server/internal/types/genui.go` y `sai-web/types/genui.ts` con `A2UIComponent`, `A2UISurface`, etc.
- Extender `sai-web/stores/demo.ts` (Pinia) con `quizSurface`, `studentProfile`, `thetaUser`, `lastAnswerCorrect`.
- Crear `components/quiz/A2UIRenderer.vue` + `RenderComponent.vue`.
- Test con un schema hardcoded en `pages/test-renderer.vue` para verificar render.

**Hora 3 — Prompts y handlers Go (02:00–03:00)**
- Crear `sai-server/internal/prompts/quiz.go` con los 3 prompts.
- Crear `sai-server/internal/irt/selector.go`.
- Crear `internal/handlers/sessions_next.go` (selector 85% + LLM) y `internal/handlers/sessions_answer.go` (actualiza thetaUser).
- Registrar rutas en `cmd/server/main.go`.
- `curl` con cada perfil, validar que los JSONs son parseables y renderizables.
- Si el LLM se equivoca de estructura, ajustar el prompt (no el código). El prompt es la API.

**Hora 4 — Toggle, integración y assets (03:00–04:00)**
- Crear `components/quiz/ProfileToggle.vue` con 3 botones tipo tabs. `@change` → `setProfile()` + re-fetch.
- Crear `components/quiz/QuizSurface.vue`: hace `$fetch` al cambiar perfil, muestra skeleton, renderiza `<A2UIRenderer />`, escucha el evento del meta, llama a `setAnswerResult`.
- Agregar caso `"quiz"` al switch de `<DemoFlow.vue />`.
- Generar las 4 imágenes pre-cacheadas para perfil visual.
- Smoke test: cambiar perfil 3 veces, verificar que la UI cambia estructuralmente.
- PR a `main`.

---

## 10. Definition of Done

- [ ] Cambiar el toggle regenera el quiz con un schema estructuralmente distinto (no solo distinto contenido).
- [ ] Los 3 perfiles renderizan algo usable.
- [ ] El `correctAnswerEvent` del meta dispara `setAnswerResult` y el step avanza a `"remediation"` cuando incorrecto, o `"done"` cuando correcto.
- [ ] El renderer ignora silenciosamente cualquier `type` que no esté en la lista de 7 primitivas (no crashea).
- [ ] El response trae `meta.irt.pSuccessPredicted` cercano a 0.85 para `thetaUser=0.0` (verificable en devtools).
- [ ] El shape del Surface (`surfaceId`, `rootComponent`, `components` map, `dataModel`, `meta`) coincide con ARCHITECTURE §3.1.
- [ ] Si el LLM devuelve JSON inválido, el fallback hardcoded se usa.
- [ ] PR mergeada a `main` con un screenshot por cada perfil en la descripción.

---

## 11. Fallbacks

- **LLM devuelve surfaces inconsistentes.** `fallbackSurfaces` hardcoded en `internal/handlers/sessions_next.go` (var package-level), uno por perfil. El demo nunca se queda sin quiz.
- **DragArea es un dolor en Vue.** Fallback: el perfil kinesthetic muestra los 4 DragItems como botones numerables (clic en orden) en lugar de drag real.
- **Genera más de 15 nodos.** Truncar el rendering a depth 4 en `RenderComponent.vue` (prop `depth`).
- **El evento no se dispara.** Botón fullscreen "Avanzar" siempre visible para forzar el flujo en vivo.

---

## 12. Handoff

Esta branch deja en `main`:

- 7 primitivas Vue reusables (branch 3 podría reusarlas).
- Tipos `A2UIComponent`, `A2UISurface` exportados en Go y TS.
- Patrón `$fetch` + render + event handling que branch 3 imita.
- Selector IRT-3PL en Go reusable por branch 3 si decide adaptar dificultad de la remediación.
