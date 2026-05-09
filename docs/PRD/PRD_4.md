PRD 4 of 4 — Demo Narrative & Submission
Branch: feat/demo-narrative
Owner: 1 persona (perfil 100% producto/diseño/storytelling)
Time budget: 4 horas wall-clock
Dependencies: Cero al principio. Las últimas 2 horas dependen de que branches 1, 2, 3 estén mergeadas.

1. Objetivo
Producir todo lo no-código que decide si ganan: README estelar, video de 3 minutos, slide deck del proyecto, formulario de submission impecable, narrativa coherente. Sin esto, el código de los otros 3 vale poco. Esta es la branch que más se subestima y más mueve la aguja.
Criterio de éxito: Un juez que solo ve el video + lee el README (sin correr el código) entiende qué hicieron, por qué importa, y dónde están los 3 niveles de Generative UI. Un juez técnico que clona el repo encuentra todo funcionando en 2 minutos.

2. Scope
SÍ:

Nombre y branding del producto.
README con estructura específica (ver §5).
Video de 2:30-3:00 con guion ya escrito.
Hero image del proyecto.
Slide deck de 5 slides (problema, solución, los 3 niveles, stack, futuro).
Submission form lleno con copy revisado.
Hardcodear los 3 prompts ganadores en el código (PR final que coordina con las otras branches).
Demo script de 90 segundos para presentación en vivo (si aplica).

NO:

Logo profesional. Un emoji + tipografía buena alcanza.
Landing page separada del README. El README ES la landing.
Múltiples versiones del video. Una sola, bien.
Texto largo. Los jueces escanean.


3. Tech stack
CapaHerramientaRazónVideoLoom o ScreenStudio (Mac)Loom gratis, 5 min máx, suficienteHero imageCanva o Figma1 sola imagen 1200x630SlidesGamma o Canva5 slides maxAudio (opcional)ElevenLabs o voz propiaVoz propia es más auténticaREADMEMarkdown puroGitHub renderiza

4. Naming y branding
Sugerencia (ajustable): "Adaptive Mind" o "Tutor Mosaic" o "Capilla" (cápsula + capilla, lugar de aprendizaje). Algo corto, googleable, sin "AI" en el nombre (todos lo tienen).
Tagline: "Un tutor que diseña la lección mientras la enseña."
Color primario: indigo (#818cf8) — alineado con el accent de la remediación.

5. README structure (texto exacto, ajusten datos)
markdown# [Nombre del producto]

> Un tutor que diseña la lección mientras la enseña.

[![Demo video](./public/hero.png)](LINK_AL_VIDEO)

## El problema

Cada estudiante aprende distinto, pero las apps de estudio tratan a todos
igual: misma UI, mismo flujo, mismo formato de pregunta. Eso es lo opuesto
a cómo funciona el aprendizaje real.

## La solución

[Nombre] es un tutor adaptativo que **diseña la interfaz de cada momento de
estudio según el estudiante**. No elige entre layouts pre-hechos: compone,
genera y a veces inventa la UI en tiempo real.

Demostramos los **tres niveles del espectro de Generative UI** en una sola
sesión de 3 minutos:

### 1. Controlada — La cápsula de estudio (Dual Coding)
El agente recibe un tema y rellena un componente fijo (`<StudyCapsule />`)
con texto, imagen y audio sincronizados — **Teoría de Codificación Dual**
de Paivio: cada cápsula entrega ≥ 2 modalidades. Pixel-perfect, on-brand,
repetible. Endpoint: `POST /api/capsules/generate`.

### 2. Declarativa — El quiz adaptativo (A2UI + Regla del 85%)
El agente emite un **A2UI Surface** (rootComponent + components map +
dataModel) que un mini-renderer pinta con primitivas. La **forma** del
quiz cambia según el perfil del estudiante: multiple choice visual,
drag-and-drop kinestésico, o pregunta abierta. El selector de dificultad
usa **IRT-3PL** y converge a P(éxito) = 0.85 (desirable difficulty).
Endpoint: `POST /api/sessions/{id}/next`.

### 3. Abierta — La remediación socrática (Feynman + Interrogación Elaborativa)
Cuando el estudiante falla, el agente aplica la **Técnica Feynman**
(simplificación + analogía) y genera **HTML+SVG+JS crudo en runtime**
que se renderiza en un iframe sandboxed, cerrando con una pregunta de
**Interrogación Elaborativa** ("¿por qué…?"). Una mini-visualización
interactiva única, hecha para esa confusión específica, que no existía
hace 5 segundos. Endpoint: `POST /api/sessions/{id}/socratic/response`.

## Demo

[GIF o screenshot de cada uno de los 3 niveles, en orden]

## Stack

**Frontend `sai-web/`:**
- Nuxt 3 + Vue 3 + TypeScript
- Tailwind CSS via `@nuxtjs/tailwindcss` + shadcn-vue
- Pinia para estado
- `vue-draggable-plus` para drag-and-drop en quiz kinestésico
- Mini-renderer A2UI propio (~150 líneas Vue, shape Surface alineado con `docs/ARCHITECTURE.md` §3.1)

**Backend `sai-server/`:**
- Go 1.22+ + Gin
- `go-openai` (GPT-4o, arch §1.2)
- `log/slog` stdlib
- SQLite o in-memory (Postgres/Redis/Docker en roadmap)
- Rutas exactas a ARCHITECTURE.md §2.1

## Cómo correrlo

\`\`\`bash
git clone https://github.com/Ricardoyesc/study-sessions-ai
cd study-sessions-ai

# Backend
cd sai-server
echo "OPENAI_API_KEY=sk-..." > .env
go run ./cmd/server  # localhost:8080

# Frontend (otra terminal)
cd ../sai-web
echo "NUXT_PUBLIC_API_BASE=http://localhost:8080" > .env
npm install
npm run dev  # localhost:3000
\`\`\`

Abrir http://localhost:3000.

## Decisiones de diseño

[Tres bullets de qué cortaron y por qué]

## Roadmap (si tuviéramos más de 4 horas)

Ya está documentada la arquitectura completa de producción en `docs/ARCHITECTURE.md`:

- **Cold-start diagnostic** con IRT-3PL + CAT (Maximum Fisher Information) + K-Means clustering, sin datos demográficos (arch §4.1)
- **Bayesian Knowledge Tracing** por concepto, con state {p_learned, p_guess, p_slip, p_transit} (arch §4.2 + tabla `user_concept_mastery`)
- **Spaced repetition SM-2** para retención a largo plazo (arch §4.2)
- **Backend Go monolito** (Gin + GORM + Postgres + Redis + MinIO) con Adapters Hexagonales para LLM/TTS/ImageGen (arch §6)
- **WebSocket A2UI** con `a2ui_full | a2ui_update | data_model_update` para hot-reload de accesibilidad (arch §2.3)
- **Despliegue Kubernetes** (arch §8) — manifiestos listos
- MCP App para que la sesión corra dentro de Claude/ChatGPT

## El equipo

[4 nombres + linkedin]

## Licencia

MIT.

6. Video script (3 minutos exactos)
Estructura por escena:
TiempoEscenaVoz en off (Spanish)Visual0:00–0:15Hook"Imagina un tutor que diseña la lección mientras la enseña, no que la elige de un menú."Hero shot del producto, cursor moviendo, indigo glow0:15–0:35Problema"Las apps de estudio actuales son rígidas: misma UI para todos, mismos formatos, misma forma de preguntar. Pero el aprendizaje real no es así."Screenshots de Duolingo/Khan Academy lado a lado, mostrando uniformidad0:35–1:05Nivel 1 — Controlada"Empezamos con una cápsula de estudio. El agente recibe el tema, elige un componente pre-diseñado y lo rellena con texto, imagen y audio — codificación dual, dos canales sensoriales sincronizados. Control total, marca consistente. Esto es Generative UI controlada."Screen recording: input "doble rendija" → cápsula aparece con imagen + texto + audio1:05–1:50Nivel 2 — Declarativa"Pero un quiz no debería ser igual para todos. Cambiamos el perfil del estudiante a 'visual', y el agente emite un schema JSON con cuatro imágenes. Lo cambiamos a 'kinestésico' y el quiz se vuelve drag-and-drop. A 'lector', y aparece una pregunta abierta. El agente no eligió de un menú: compuso la UI desde primitivas. Esto es Generative UI declarativa, vía A2UI."Screen recording: toggle entre los 3 perfiles, cada uno renderiza algo distinto1:50–2:35Nivel 3 — Abierta"Y cuando el estudiante falla — fallamos a propósito — el agente aplica la técnica Feynman: re-explica con una analogía simple, y cierra con una pregunta socrática 'por qué'. Para hacerlo, genera HTML, SVG y JavaScript en runtime, los mete en un iframe sandboxed, y produce esto: una visualización interactiva única, hecha exactamente para esta confusión, que no existía hace cinco segundos. Esto es Generative UI abierta."Screen recording: fallar quiz → loader → SVG aparece → mover slider → ver patrón colapsar2:35–3:00Cierre"Tres niveles de Generative UI en una sola sesión de estudio. Control donde importa la marca. Composición donde importa la adaptación. Generación abierta donde importa la creatividad. Eso es la pirámide completa, aplicada a aprendizaje."Logo + tagline + url del repo

7. Plan minuto a minuto
Hora 1 (00:00–01:00) — Independiente, no espera nada

Decidir nombre + tagline + color primario.
Crear hero image en Canva (1200x630, dark, indigo glow, nombre del producto centrado).
Empezar el README con copy del §5, dejando blanks para screenshots.
Crear el slide deck de 5 slides en Gamma con prompt: "5-slide investor deck for an adaptive learning tutor that demonstrates 3 levels of Generative UI. Audience: hackathon judges. Dark indigo theme."

Hora 2 (01:00–02:00) — Branding y form

Llenar el formulario de submission en borrador.
Pulir el README.
Escribir las 3 frases hook que el equipo va a usar en el chat con jueces o redes.
Definir los hashtags / canales donde van a postear el demo.

Hora 3 (02:00–03:00) — Grabación (las otras branches deben estar mergeadas)

Hacer rebase de main → tener el demo funcionando local.
Ensayar el guion 2-3 veces sin grabar.
Grabar el video en Loom o ScreenStudio. Si el primer take es decente, no rehacer. Subir a YouTube unlisted o Loom.
Capturar 3 GIFs (uno por nivel) para el README.

Hora 4 (03:00–04:00) — Submission

Insertar GIFs y link del video al README.
Final pass del README (typos, links rotos).
Coordinar con las otras 3 personas: hardcodear los prompts ganadores que descubrieron iterando.
Llenar el form de submission, copy revisado, links verificados (clic en cada uno).
Submit.
Cerrar laptops. No tocar nada.


8. Definition of Done

 README en main con todas las secciones del §5 llenas.
 Video subido y linkeado en README.
 Hero image en /public/hero.png.
 3 GIFs en README (uno por nivel).
 Slide deck en URL pública linkeada.
 Formulario de submission enviado.
 Tweet o post de LinkedIn del equipo con el video (opcional pero alto-leverage).


9. Riesgos

Las otras branches no llegan a tiempo a la hora 3. Plan B: grabar con la branch que esté más completa + voiceover que diga "imaginen aquí el quiz" mientras se ven los mockups del slide deck. No ideal pero salvable.
Loom corta a 5 min en plan free. El video debe quedar bajo 3 min. Si se pasa, recortar el problema (no la demo).
Submission form pide cosas que no tenemos (ej. URL de deploy en producción). Plan B: deploy a Vercel en 5 min con vercel --prod. La key de Anthropic se mete como env var en Vercel UI.


10. Coordinación con las otras 3 branches

Hora 1: Esta branch puede arrancar independiente. Las otras 3 ni se enteran.
Hora 2 (01:00): Sync de equipo de 5 min. Ver cómo va cada branch. Esta persona toma notas para el README.
Hora 3 (02:30): Las otras 3 branches DEBEN estar mergeadas a main. Si no, esta persona graba con lo que haya.
Hora 4 (03:30): Stand-up final de 5 min antes del submit.


Resumen de las 4 branches
BranchNivel GenUIOutputRiesgofeat/capsule-controlledControladaStudyCapsule + scaffold del proyectoBajofeat/quiz-a2uiDeclarativaA2UI mini-renderer + 3 perfilesMedio (drag-drop)feat/remediation-openAbiertaIframe sandboxed con HTML generadoAlto (calidad LLM)feat/demo-narrative—README + video + submissionBajo si arranca a tiempo
Regla de oro: si en la hora 3 algo no funciona, cortar features, no calidad. Mejor 2 niveles bien hechos + un mock del tercero, que 3 niveles a medias.