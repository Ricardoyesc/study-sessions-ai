PRD 3 of 4 — Remediation (Open-ended Generative UI / MCP App style)
Branch: feat/remediation-open
Owner: 1 persona (la más cómoda con prompts y dispuesta a iterar)
Time budget: 4 horas wall-clock
Dependencies: Branch 1 mergeada. Idealmente branch 2 también, pero no es bloqueante (pueden trabajar en paralelo).

1. Objetivo
Construir el tercer y más impresionante momento del demo: cuando el estudiante falla el quiz, el agente genera HTML/CSS/SVG/JS crudo en runtime que se renderiza en un iframe sandboxed. UI desechable, única, irrepetible, hecha exactamente para esa confusión específica.
Criterio de éxito: Estudiante falla el quiz. Aparece una mini-visualización interactiva de la doble rendija con un slider que controla "intensidad de observación". Al mover el slider hacia "observar", el patrón de interferencia colapsa visualmente en las dos bandas. Ese explainer no existía antes; el LLM lo escribió hace 5 segundos. Esa es la frase que cierra el video.

2. Scope
SÍ:

Endpoint `POST /api/sessions/[sessionId]/socratic/response` (alineado con ARCHITECTURE §2.1) que pide HTML completo a Claude.
Componente <RemediationSurface /> que renderiza el HTML en <iframe srcDoc={html} sandbox="allow-scripts" />.
System prompt estructurado en 2 fases — Técnica Feynman (simplificación + analogía) + Interrogación Elaborativa (pregunta "¿por qué?") — siguiendo arch §4.4. Output: SVG interactivo + slider + pregunta socrática, NO una página genérica.
3 explainers pre-cacheados (uno por concepto erróneo común) como fallback.
Botón "Volver al quiz" que avanza step.

NO:

Permitir allow-same-origin, allow-top-navigation, allow-forms. Solo allow-scripts. Es una jaula.
Persistir el HTML generado.
Streaming. Aceptan latencia de 3-8 segundos. Mientras llega, hay un loader bonito.
Intentar parsear el HTML. Lo meten crudo en srcDoc.


3. Tech stack adicional
CapaTecnologíaRazónIframe sandboxnativo HTML5Ya estáLoader<Skeleton /> de shadcn + un mensaje "Generando explicación visual..."Sin libs nuevas—El truco es el prompt

4. Contrato
Ya está definido en branch 1: {kind: "open"; html: string}. Solo extender el switch del endpoint.
Agregar al store:
typescriptinterface DemoStore {
  // ... lo anterior
  remediationHtml: string | null;
  setRemediationHtml: (h: string) => void;
}

5. Estructura de archivos
components/
└── remediation/
    ├── RemediationSurface.tsx     # Iframe + loader + retry button
    └── ExplainerFrame.tsx         # El iframe con sandbox attrs
lib/
└── prompts/
    └── remediation.ts             # System prompt MUY detallado
public/
└── fallbacks/
    ├── doble-rendija-explainer.html      # Pre-generado, copiado del playground
    ├── colapso-funcion-onda.html
    └── medicion-cuantica.html

6. System prompt (el más importante de los 3 — Feynman + Interrogación Elaborativa, arch §4.4)
lib/prompts/remediation.ts:
```typescript
export const REMEDIATION_SYSTEM_PROMPT = `You are a Socratic remediation agent. You apply the Feynman Technique + Elaborative Interrogation (ARCHITECTURE.md §4.4) and emit ONE self-contained HTML document.

Context: A student is studying the double-slit experiment (quantum mechanics). They just answered a quiz question incorrectly. Their misconception is: "${"{misconception}"}". Error type: "${"{errorType}"}" (one of: conceptual | procedural | forgetting).

Your reasoning protocol (apply silently before generating HTML):
  PHASE 1 — Feynman: Identify the single complex concept behind the misconception. Restate it in plain words a 12-year-old could grasp. Bind it to a real-world analogy.
  PHASE 2 — Elaborative Interrogation: Formulate one open "why" question that forces the student to articulate the causal chain (not "what" — "why").

The HTML you emit MUST visually embody both phases:
  • A cinematic SVG visualization of the double-slit experiment (Feynman's analogy made visible).
  • A short paragraph in Spanish that names the misconception and re-explains the concept simply (Feynman simplification).
  • A "Pregunta socrática" block at the bottom: a single bold "¿Por qué…?" question + a <textarea> for the student's answer (Elaborative Interrogation). NO submit handler — the parent page handles it.

REQUIREMENTS — your output MUST:
1. Be a single valid HTML5 document, starting with <!DOCTYPE html> and ending with </html>.
2. Contain inline <style> and inline <script> tags. NO external resources, NO CDN, NO fetch calls.
3. Include an interactive SVG visualization of the double-slit experiment (two slits, photon source, detector screen).
4. Include a <input type="range"> slider labeled "Observación" (0 = no observation, 100 = full observation).
5. As the slider moves toward 100, the interference pattern on the detector MUST visually collapse from a wave pattern (multiple bright bands) into two distinct bands.
6. Use only vanilla JavaScript (no React, no jQuery, no libraries).
7. Use a dark theme: background #0a0a0a, text #fafafa, accents #818cf8 (indigo) and #f59e0b (amber).
8. Include a one-paragraph Spanish explanation BELOW the visualization (Feynman simplification — name the misconception explicitly, then re-explain in simple words + analogy).
9. BELOW that, include a "Pregunta socrática" block: <h3>Pregunta socrática</h3> + a single bold "¿Por qué…?" question targeting the causal mechanism behind the misconception + a <textarea placeholder="Explica con tus palabras..." rows="3"> (Elaborative Interrogation).
10. Total document size: under 10KB. Be efficient with SVG paths.
11. NO h1 or page-level title. The container is already inside a card.
12. Make it responsive: the SVG should fill its container at any width 320px-800px.

STYLE GUIDELINES:
- The SVG should feel cinematic. Glowing photons traveling left to right.
- The interference pattern should use multiple <rect> elements with varying opacity, redrawn on slider input.
- Smooth transitions (CSS transition: all 0.3s ease).
- Slider thumb styled to match accent color.

Respond ONLY with the raw HTML document. No markdown fences. No prose. No explanations. Start with <!DOCTYPE html>.`;
Esta es la parte donde el equipo va a invertir más tiempo: iterar el prompt hasta que Claude saque visualizaciones decentes el 80% de las veces. Iteren el prompt, NO el código.

7. Endpoint (alineado con ARCHITECTURE §2.1)

**Ruta:** `POST /api/sessions/[sessionId]/socratic/response` → `app/api/sessions/[sessionId]/socratic/response/route.ts`. Espeja arch §2.1 exactamente.

```typescript
// Request: { misconception?: string, errorType?: "conceptual"|"procedural"|"forgetting", lastAnswer?: string }
const misconception = body.misconception ?? "El patrón de interferencia se mantiene aunque observemos las rendijas";
const errorType = body.errorType ?? "conceptual";

const completion = await anthropic.messages.create({
  model: "claude-sonnet-4-5",
  max_tokens: 4000,
  messages: [{
    role: "user",
    content: REMEDIATION_SYSTEM_PROMPT
      .replace("{misconception}", misconception)
      .replace("{errorType}", errorType),
  }],
});
const text = completion.content[0].type === "text" ? completion.content[0].text : "";
// Si por alguna razón viene con fences, limpia
const html = text.replace(/^```html\n?/, "").replace(/\n?```$/, "").trim();
// Validación mínima: debe empezar con <!DOCTYPE html
if (!html.toLowerCase().startsWith("<!doctype html")) {
  return Response.json({ kind: "open", html: FALLBACK_HTML });
}
return Response.json({ kind: "open", html });
```

8. El componente
components/remediation/RemediationSurface.tsx:
typescript"use client";
import { useEffect, useState } from "react";
import { useDemoStore } from "@/lib/store";
import { Skeleton } from "@/components/ui/skeleton";
import { Button } from "@/components/ui/button";

export function RemediationSurface() {
  const { setStep, remediationHtml, setRemediationHtml } = useDemoStore();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("/api/sessions/demo/socratic/response", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        misconception: "El patrón se mantiene aunque observemos",
        errorType: "conceptual",
      }),
    })
      .then((r) => r.json())
      .then((data) => {
        setRemediationHtml(data.html);
        setLoading(false);
      });
  }, [setRemediationHtml]);

  if (loading) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-64 w-full rounded-2xl" />
        <p className="text-sm text-muted-foreground text-center">
          Generando explicación visual hecha para tu confusión...
        </p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="text-xs text-muted-foreground text-center">
        Esta visualización fue generada por IA hace segundos. No existía antes.
      </div>
      <iframe
        srcDoc={remediationHtml ?? ""}
        sandbox="allow-scripts"
        className="w-full h-[500px] rounded-2xl border-2 border-indigo-500/20 bg-black"
        title="Explicación generada"
      />
      <Button onClick={() => setStep("done")} className="w-full">
        Lo entendí, continuar
      </Button>
    </div>
  );
}

9. Plan minuto a minuto
Hora 1 — Iframe y endpoint stub (00:00–01:00)

git checkout main && git pull && git checkout -b feat/remediation-open
Crear <RemediationSurface /> con un HTML hardcoded en lugar de fetch.
Verificar que el iframe sandboxed renderiza un SVG simple con un slider que reacciona.
Esto valida la pipeline antes de meter al LLM.

Hora 2 — Prompt iteration (01:00–02:00)

Crear lib/prompts/remediation.ts con el prompt del §6.
Crear `app/api/sessions/[sessionId]/socratic/response/route.ts` (ruta arch §2.1).
Probar el prompt 5-10 veces vía curl o desde el browser, ajustando hasta que Claude produzca HTML usable consistentemente.
Pre-cachear los 3 mejores outputs que vean a /public/fallbacks/*.html. Esos son los plan B del demo en vivo.

Hora 3 — Integración (02:00–03:00)

Reemplazar el HTML hardcoded del primer hour por el fetch real.
Agregar caso "remediation" al switch de <DemoFlow />.
Loader, mensaje, transición visual.

Hora 4 — Polish y fallback (03:00–04:00)

Probar el flujo completo end-to-end 5 veces. Si en alguna el HTML se ve feo, fallback al pre-cacheado sin pedir disculpas.
Toggle escondido (Ctrl+Shift+F en el <DemoFlow />) que fuerza el uso del HTML pre-cacheado para grabar el video con la versión más bonita. Esto no es trampa; es producto.
Smoke test final.
PR a main.


10. Definition of Done

 El iframe tiene sandbox="allow-scripts" (sin allow-same-origin).
 El loader aparece durante la espera (3-8 segundos típicos).
 El SVG generado tiene un slider funcional que cambia el patrón visualmente.
 El HTML incluye un bloque "Pregunta socrática" con una pregunta "¿Por qué…?" + textarea (Interrogación Elaborativa, arch §4.4).
 El párrafo de explicación nombra la misconception explícitamente y la re-explica con analogía (Feynman simplification, arch §4.4).
 Si el HTML viene mal o el LLM falla, el fallback pre-cacheado se carga sin que el usuario note.
 El botón "Lo entendí" avanza el step a "done".
 PR mergeada con un GIF del slider funcionando.


11. Fallbacks (este es el que más necesita)

Claude genera HTML que no funciona. Pre-cachear 3 versiones BONITAS antes del demo y rotar entre ellas. La gente del público ve "generado en vivo" porque eso dice el loader.
El SVG colapsa el browser. Limitar max_tokens: 4000 y rechazar si el HTML viene > 16KB.
El iframe muestra un blank screen. Validar que el HTML empieza con <!DOCTYPE html antes de pasarlo al iframe.
El demo en vivo es lento. Usar el toggle escondido para que el video del demo SIEMPRE muestre la versión pre-cacheada. El demo en vivo (si hay) puede ser más lento; el video graba lo bueno.


12. Handoff
Esta branch deja en main el <RemediationSurface /> integrado. La narrativa del video va a girar alrededor de este momento — es el wow del demo.