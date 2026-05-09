PRD 2 of 4 — Quiz (Declarative Generative UI / A2UI)
Branch: feat/quiz-a2ui
Owner: 1 persona (la que tenga más tolerancia a JSON-driven UI)
Time budget: 4 horas wall-clock
Dependencies: Branch 1 mergeada en main. Hacer git checkout main && git pull && git checkout -b feat/quiz-a2ui antes de empezar.

1. Objetivo
Construir el segundo momento del demo: un quiz adaptativo donde el agente no escoge entre componentes pre-hechos, sino que emite un schema JSON que un mini-renderer paint con primitivas. El agente decide la forma del quiz según el "perfil del estudiante", no solo el contenido.
Criterio de éxito: En el demo, hay un toggle visible "Perfil del estudiante: [Visual | Kinestésico | Lector]". Al cambiar el toggle, el agente regenera el quiz con un schema estructuralmente distinto:

Visual → multiple choice con imágenes en grid 2x2.
Kinestésico → drag-and-drop para ordenar pasos.
Lector → input de texto libre con pregunta abierta.

Los jueces ven que un mismo concepto se evalúa con UI distinta y que la UI es composición, no template-switching.

2. Scope
SÍ:

Mini-renderer A2UI casero (~150 líneas TS) que mapea schema JSON → componentes React.
6 primitivas: Stack (Column), Row, Text, Image, Button, TextInput, DragItem.
Endpoint POST /api/agent extendido con intent: "generate_quiz" que devuelve {kind: "declarative", schema: A2UISurface}.
3 prompts distintos (uno por perfil) que generan schemas distintos para el mismo concepto.
Componente <QuizSurface schema={...} /> que renderiza el árbol y dispara eventos.
Toggle de perfil arriba del quiz, conectado al Zustand.
Integración con el flujo: cuando el quiz termina (correcto o incorrecto), avanza el step.

NO:

A2UI library oficial (a2ui.org). Implementan su propia versión mini. La oficial requiere setup que no tienen.
Validación exhaustiva del schema. Confían en el LLM. Try/catch alrededor del render y fallback a un quiz hardcoded si el schema viene mal.
Más de 6 primitivas. Cualquier cosa que no esté en la lista, no se renderiza.
Animaciones complejas, drag-drop con física. El drag es funcional, no bonito.
Persistencia de respuestas. Solo importa el "did they answer?" boolean.


3. Tech stack adicional (sobre lo que branch 1 dejó)
CapaTecnologíaVersiónRazónDrag & drop@dnd-kit/core + @dnd-kit/sortablelatestMucho más simple que react-beautiful-dndValidación schemanada (try/catch)—No vale la pena Zod en 4hRendererFunción recursiva propia—~150 líneas
npm install @dnd-kit/core @dnd-kit/sortable

4. Contrato del schema A2UI (alineado con ARCHITECTURE §3.1)

> **Decisión:** el shape sigue la struct Go `Surface` de ARCHITECTURE §3.1 (`surfaceId`, `rootComponent`, `components` map, `dataModel`). Así el mismo cliente puede consumir el backend Go futuro sin tocar el renderer. Las 7 primitivas TS son un subset — el backend Go añadirá `RichText`, `AudioPlayer`, `VideoPlayer`, `Card`, `QuizCard`, `SocraticDialog`, `ProgressBar`, `Column` (alias de Stack vertical).

Agregar a lib/store.ts (extender, no sustituir):
```typescript
// Primitivas soportadas en la PoC. CUALQUIER otro type se ignora silenciosamente.
// Mantenemos el shape Component de arch §3.1: { id, type, children?: string[], props, events? }.
export type A2UIComponent =
  | { id: string; type: "Stack"; children: string[]; props?: { gap?: number; align?: "start"|"center"|"end" } }
  | { id: string; type: "Row"; children: string[]; props?: { gap?: number; wrap?: boolean } }
  | { id: string; type: "Text"; props: { content: string; variant?: "h2"|"h3"|"body"|"caption" } }
  | { id: string; type: "Image"; props: { url: string; alt: string; aspectRatio?: "1:1"|"16:9"|"4:3" } }
  | { id: string; type: "Button"; props: { label: string; variant?: "primary"|"secondary"|"ghost" }; events?: { onClick?: string }; eventPayload?: Record<string, unknown> }
  | { id: string; type: "TextInput"; props: { placeholder: string }; events?: { onSubmit?: string } }
  | { id: string; type: "DragItem"; props: { label: string; correctOrder: number } };

// DataModel global (arch §3.1) — accesibilidad/tema. PoC lo usa con defaults.
export type A2UIDataModel = {
  theme?: "system" | "light" | "dark";
  fontFamily?: string;
  fontScale?: number;
  colorPalette?: string;
  highContrast?: boolean;
  reducedMotion?: boolean;
  language?: string;
};

// Surface = "pantalla" completa (arch §3.1)
export type A2UISurface = {
  surfaceId: string;
  rootComponent: string;                     // id del componente raíz (era "rootId")
  components: Record<string, A2UIComponent>; // adjacency map (era "nodes")
  dataModel?: A2UIDataModel;
  meta: {
    intent: "quiz";
    profile: StudentProfile;
    correctAnswerEvent: string;              // ej: "answer:correct" o "order:check"
    // Trazabilidad del 85% rule + IRT (arch §4.2). PoC: el LLM rellena difficultyIRT y discrimination heurísticamente; el frontend solo muestra/loguea.
    irt?: { thetaUser: number; difficultyIRT: number; discrimination: number; pSuccessPredicted: number };
  };
};

export type StudentProfile = "visual" | "kinesthetic" | "reader";

// Extender el union type que branch 1 dejó:
export type GenUIResponse =
  | { kind: "controlled"; component: "StudyCapsule"; capsuleId: string; topic: string; modalities: Modality[]; props: CapsuleProps }
  | { kind: "declarative"; surface: A2UISurface }     // ⭐ ESTE PRD (arch §3.1)
  | { kind: "open"; html: string };
```

Y al store:
```typescript
interface DemoStore {
  step: DemoStep;
  capsule: CapsuleProps | null;
  quizSurface: A2UISurface | null;
  studentProfile: StudentProfile;
  thetaUser: number;                  // habilidad latente (arch §4.1, default 0.0)
  lastAnswerCorrect: boolean | null;
  setStep: (s: DemoStep) => void;
  setCapsule: (c: CapsuleProps) => void;
  setQuizSurface: (s: A2UISurface) => void;
  setProfile: (p: StudentProfile) => void;
  setAnswerResult: (correct: boolean) => void;
}
```

5. Estructura de archivos a crear
components/
├── quiz/
│   ├── QuizSurface.tsx          # Wrapper, hace fetch, maneja toggle de perfil
│   ├── A2UIRenderer.tsx         # ⭐ El renderer recursivo
│   ├── ProfileToggle.tsx        # Tabs visuales
│   └── primitives/
│       ├── Stack.tsx
│       ├── Row.tsx
│       ├── A2UIText.tsx
│       ├── A2UIImage.tsx
│       ├── A2UIButton.tsx
│       ├── A2UITextInput.tsx
│       └── DragArea.tsx         # Wrapper de @dnd-kit que detecta DragItems en children
lib/
├── prompts/
│   └── quiz.ts                  # 3 system prompts (uno por perfil)
└── a2ui/
    └── eventBus.ts              # window CustomEvent helpers para que primitivas emitan
public/
├── quiz-doble-rendija-1.png     # Pre-generadas para perfil visual
├── quiz-doble-rendija-2.png
├── quiz-doble-rendija-3.png
└── quiz-doble-rendija-4.png

6. El renderer (corazón de esta branch)
components/quiz/A2UIRenderer.tsx:
```typescript
"use client";
import { A2UIComponent, A2UISurface } from "@/lib/store";
import { Stack } from "./primitives/Stack";
import { Row } from "./primitives/Row";
import { A2UIText } from "./primitives/A2UIText";
import { A2UIImage } from "./primitives/A2UIImage";
import { A2UIButton } from "./primitives/A2UIButton";
import { A2UITextInput } from "./primitives/A2UITextInput";
import { DragArea } from "./primitives/DragArea";

// children son IDs (string[]) — adjacency map, arch §3.1.
export function A2UIRenderer({ surface }: { surface: A2UISurface }) {
  return <RenderComponent id={surface.rootComponent} surface={surface} />;
}

function RenderComponent({ id, surface }: { id: string; surface: A2UISurface }) {
  const node = surface.components[id];
  if (!node) return null;

  // Detectar si los children son DragItems → wrap en DragArea
  if ("children" in node) {
    const childNodes = node.children.map((cid) => surface.components[cid]).filter(Boolean);
    const allDraggable = childNodes.length > 0 && childNodes.every((c) => c.type === "DragItem");
    if (allDraggable) return <DragArea items={childNodes as any} surface={surface} />;
  }

  switch (node.type) {
    case "Stack":
      return (
        <Stack {...(node.props ?? {})}>
          {node.children.map((cid) => <RenderComponent key={cid} id={cid} surface={surface} />)}
        </Stack>
      );
    case "Row":
      return (
        <Row {...(node.props ?? {})}>
          {node.children.map((cid) => <RenderComponent key={cid} id={cid} surface={surface} />)}
        </Row>
      );
    case "Text":      return <A2UIText {...node.props} />;
    case "Image":     return <A2UIImage {...node.props} />;
    case "Button":    return <A2UIButton id={node.id} eventName={node.events?.onClick} eventPayload={(node as any).eventPayload} {...node.props} />;
    case "TextInput": return <A2UITextInput id={node.id} eventName={node.events?.onSubmit} {...node.props} />;
    case "DragItem":  return null; // Solo dentro de DragArea
    default:          return null;
  }
}
```
lib/a2ui/eventBus.ts:
typescriptexport function emit(eventName: string, payload?: Record<string, unknown>) {
  window.dispatchEvent(new CustomEvent(`a2ui:${eventName}`, { detail: payload }));
}

export function on(eventName: string, handler: (payload: any) => void) {
  const wrapped = (e: Event) => handler((e as CustomEvent).detail);
  window.addEventListener(`a2ui:${eventName}`, wrapped);
  return () => window.removeEventListener(`a2ui:${eventName}`, wrapped);
}
<QuizSurface /> escucha el correctAnswerEvent del meta, decide correct/incorrect, llama a setAnswerResult y avanza el step.

7. System prompts (los 3 perfiles)
lib/prompts/quiz.ts:
typescriptconst COMMON = `You are a Generative UI agent. You output a JSON schema describing UI to render.

Allowed component types ONLY: Stack, Row, Text, Image, Button, TextInput, DragItem. Any other type will be ignored.

Schema format (Surface, alineado con ARCHITECTURE §3.1):
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
Buttons emit events via `events.onClick` (string), not `emitsEvent`. TextInputs use `events.onSubmit`.
Keep tree depth <= 4. Keep total nodes <= 15.
Topic: "doble rendija" (double-slit experiment, quantum mechanics).
Question concept: "What happens to the interference pattern when we observe which slit each photon passes through?"
Correct answer: the pattern disappears (collapses into two bands).
All text in Spanish.

Respond ONLY with valid JSON. No prose, no fences.`;

export const QUIZ_VISUAL = `${COMMON}

Profile: VISUAL learner.
Generate a multiple choice quiz with 4 options, EACH option being an Image (use urls /quiz-doble-rendija-1.png through /quiz-doble-rendija-4.png) inside a Button with a short label.
Layout: Stack with question Text on top, then a Row of 4 Buttons (wrap: true).
The Button for the correct image must set events.onClick to "answer:correct". The other three set events.onClick to "answer:wrong".
Set meta.correctAnswerEvent to "answer:correct".`;

export const QUIZ_KINESTHETIC = `${COMMON}

Profile: KINESTHETIC learner.
Generate a drag-and-drop ordering quiz. The student must order 4 DragItems representing the steps of the experiment in causal sequence.
The 4 DragItems with their correctOrder values:
1. "Disparamos un electrón hacia la doble rendija"
2. "Decidimos si observamos por cuál rendija pasa"
3. "El electrón impacta la pantalla detectora"
4. "Tras muchos electrones, observamos el patrón resultante"
Layout: Stack with instruction Text, then the 4 DragItem nodes as children of the Stack (the renderer auto-wraps them in a DragArea), then a Button "Verificar orden" with events.onClick = "order:check".
Set meta.correctAnswerEvent to "order:correct" (the renderer will check correctOrder).`;

export const QUIZ_READER = `${COMMON}

Profile: READER learner.
Generate an open-ended question quiz. Layout: Stack with a paragraph-style Text introducing context, an h3 Text with the question itself, a TextInput (events.onSubmit = "answer:submitted"), and a Button "Enviar respuesta" with events.onClick = "answer:submitted".
The TextInput placeholder should hint that a 1-2 sentence answer is expected.
Set meta.correctAnswerEvent to "answer:submitted" (this profile accepts any non-empty answer as correct in the demo).`;

export function getPrompt(profile: "visual" | "kinesthetic" | "reader"): string {
  return { visual: QUIZ_VISUAL, kinesthetic: QUIZ_KINESTHETIC, reader: QUIZ_READER }[profile];
}

8. Endpoint (alineado con ARCHITECTURE §2.1)

**Ruta:** `POST /api/sessions/[sessionId]/next` → `app/api/sessions/[sessionId]/next/route.ts`. Espeja arch §2.1 `GET /api/sessions/{id}/next`. Para la PoC usamos POST porque el body trae `profile` y `thetaUser`.

**Request:**
```json
{ "profile": "visual", "thetaUser": 0.0, "topic": "doble-rendija" }
```

**Response:** `{ "kind": "declarative", "surface": A2UISurface }` — envelope del PRD 1 §6.1 extendido.

**Selector de dificultad (Regla del 85%, arch §4.2 simplificado para PoC):**
```typescript
// IRT-3PL: P_success = c + (1-c) * 1/(1 + exp(-a*(theta-b)))
function pSuccess(theta: number, a: number, b: number, c = 0.25): number {
  return c + (1 - c) / (1 + Math.exp(-a * (theta - b)));
}

// Picker: para topic "doble-rendija" tenemos 3 dificultades calibradas (b in {-0.5, 0.0, 0.8}).
// Elige la b cuya P_success_predicha esté más cerca de 0.85.
function pickDifficulty(theta: number) {
  const items = [{b:-0.5,a:1.0},{b:0.0,a:1.0},{b:0.8,a:1.2}];
  return items.reduce((best, it) => {
    const p = pSuccess(theta, it.a, it.b);
    return Math.abs(p - 0.85) < Math.abs(best.p - 0.85) ? { ...it, p } : best;
  }, { ...items[0], p: pSuccess(theta, items[0].a, items[0].b) });
}
```

```typescript
// app/api/sessions/[sessionId]/next/route.ts
const profile = body.profile as StudentProfile;
const theta = typeof body.thetaUser === "number" ? body.thetaUser : 0.0;
const irt = pickDifficulty(theta);
const prompt = getPrompt(profile) + `\n\nDifficulty target (b): ${irt.b}. Calibrate vocabulary/cognitive load accordingly.`;

const completion = await anthropic.messages.create({
  model: "claude-sonnet-4-5",
  max_tokens: 2000,
  messages: [{ role: "user", content: prompt }],
});
const text = completion.content[0].type === "text" ? completion.content[0].text : "";
const cleaned = text.replace(/```json|```/g, "").trim();
let surface: A2UISurface;
try {
  surface = JSON.parse(cleaned) as A2UISurface;
} catch {
  surface = FALLBACK_SURFACES[profile]; // hardcoded, ver §11
}
// Inyectar trazabilidad IRT (no la rellena el LLM)
surface.meta.irt = { thetaUser: theta, difficultyIRT: irt.b, discrimination: irt.a, pSuccessPredicted: irt.p };
return Response.json({ kind: "declarative", surface });
```

**Respuesta del usuario:** `POST /api/sessions/[sessionId]/quiz/answer` (arch §2.1) con `{ correct: boolean }`. PoC: actualiza `thetaUser` con un step heurístico (`+0.3` si correct, `-0.3` si no). Backend Go futuro hace MLE real.

9. Plan minuto a minuto
Hora 1 — Primitivas (00:00–01:00)

git checkout main && git pull && git checkout -b feat/quiz-a2ui
Crear las 7 primitivas en components/quiz/primitives/. Cada una ~15 líneas, Tailwind, sin lógica de negocio.
Crear lib/a2ui/eventBus.ts.

Hora 2 — Renderer y schema types (01:00–02:00)

Extender lib/store.ts con A2UIComponent, A2UISurface, A2UIDataModel, StudentProfile, y los nuevos campos del store (incluido thetaUser).
Crear components/quiz/A2UIRenderer.tsx.
Test con un schema hardcoded en una página /test-renderer para verificar que renderiza correctamente.

Hora 3 — Prompts y endpoint (02:00–03:00)

Crear lib/prompts/quiz.ts con los 3 prompts.
Crear `app/api/sessions/[sessionId]/next/route.ts` (selector 85% + LLM) y `app/api/sessions/[sessionId]/quiz/answer/route.ts` (actualiza thetaUser).
Probar curl con cada perfil, validar que los JSONs son parseables y renderizables.
Si Claude se equivoca de estructura, ajustar el prompt (no el código). El prompt es la API.

Hora 4 — Toggle, integración y assets (03:00–04:00)

Crear components/quiz/ProfileToggle.tsx con 3 botones tipo tabs. onChange → setProfile() + re-fetch del quiz.
Crear components/quiz/QuizSurface.tsx: hace fetch al cambiar perfil, muestra skeleton, renderiza <A2UIRenderer />, escucha el evento del meta, llama a setAnswerResult.
Agregar caso "quiz" al switch de <DemoFlow />.
Generar las 4 imágenes pre-cacheadas para el perfil visual (DALL-E o Gemini, prompt: 4 variaciones del experimento de doble rendija mostrando distintos resultados — patrón de interferencia, dos bandas, ondas, partículas).
Smoke test: cambiar perfil 3 veces, verificar que la UI cambia estructuralmente, no solo de contenido.
PR a main.


10. Definition of Done

 Cambiar el toggle regenera el quiz con un schema estructuralmente distinto (no solo distinto contenido).
 Los 3 perfiles renderizan algo usable (no perfecto, usable).
 El evento correctAnswerEvent del meta dispara setAnswerResult y el step avanza a "remediation" cuando la respuesta es incorrecta, o a "done" cuando es correcta.
 El renderer ignora silenciosamente cualquier type que no esté en la lista de 7 primitivas (no crashea).
 El response trae `meta.irt.pSuccessPredicted` cercano a 0.85 para `thetaUser=0.0` (verificable en devtools).
 El shape del Surface (surfaceId, rootComponent, components map, dataModel, meta) coincide con ARCHITECTURE §3.1 — un swap a backend Go solo cambia la base URL del fetch.
 Si Claude devuelve JSON inválido, el fallback hardcoded se usa.
 PR mergeada a main con un screenshot por cada perfil en la descripción.


11. Fallbacks

Claude devuelve surfaces inconsistentes. Tener FALLBACK_SURFACES hardcoded en lib/prompts/quiz.ts, uno por perfil, listos para usar si el JSON.parse falla. El demo nunca se queda sin quiz.
DragArea es un dolor. Fallback: el perfil kinesthetic muestra los 4 DragItems como botones numerables (clic en orden) en lugar de drag real. La narrativa lo cubre igual.
Genera más de 15 nodos. Truncar el rendering a depth 4 en el renderer. Mejor visual roto que crash.
El evento no se dispara. Botón fullscreen "Avanzar" siempre visible para forzar el flujo en vivo.


12. Handoff
Esta branch deja en main:

6 primitivas reusables (branch 3 podría reusarlas si quisiera).
Tipos A2UINode, A2UISurface exportados.
Patrón de fetch + render + event handling que branch 3 imita.