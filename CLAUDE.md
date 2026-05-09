# CLAUDE.md

Operating brief for Claude Code (and any agent following the `AGENTS.md` convention) on this repo. Read before any non-trivial change.

---

## Project context

**Repository:** AI Tinkerers Hack — Adaptive Multimodal Study System (PoC).

**What it is:** A hackathon-scope adaptive tutor that demonstrates the **three tiers of Generative UI** (Controlled / Declarative-A2UI / Open-ended) inside a single 3-minute study session. Pedagogically grounded in Dual Coding Theory, Item Response Theory (IRT-3PL), the 85% Rule, the Feynman Technique, and Elaborative Interrogation.

**My role:** Hackathon competitor. Optimize for *demo impact* and *judging criteria* over polish. Ship working slices end-to-end before deepening any one of them.

---

## Authoritative docs (read in this order)

1. **`docs/ARCHITECTURE.md`** — full production design (Go monolith + Nuxt 3 + K8s + Postgres + Redis + MinIO + WebSocket A2UI + IRT/BKT/SM-2). This is the **target** architecture. Hackathon code does NOT implement all of this.
2. **`docs/PRD/PRD_1.md`** — Capsule (Controlled UI). 4hr block.
3. **`docs/PRD/PRD_2.md`** — Quiz (Declarative A2UI). 4hr block.
4. **`docs/PRD/PRD_3.md`** — Remediation (Open-ended HTML). 4hr block.
5. **`docs/PRD/PRD_4.md`** — Demo narrative + submission. 4hr block.

If a PRD contradicts ARCHITECTURE.md on stack/scope, **the PRD wins** for the hackathon build (4-hour, Next.js). ARCHITECTURE.md is the post-hackathon roadmap and the **contract source of truth** (route names, schema shapes).

---

## Hybrid stack decision (key recent decision)

PoC keeps a lightweight Next.js stack but **contracts mirror ARCHITECTURE.md** so the frontend can swap to the Go backend by changing only the base URL.

| Layer | PoC (PRDs) | Production (ARCHITECTURE.md) |
|---|---|---|
| Frontend | Next.js 15 (App Router) + TypeScript | Nuxt 3 + Vue 3 (separate repo) |
| State | Zustand single store | composables + WebSocket |
| LLM | `@anthropic-ai/sdk` direct (Claude Sonnet 4.5) | OpenAI GPT-4o via port/adapter |
| Backend | Next.js route handlers (no DB) | Go + Gin + GORM + Postgres + Redis + MinIO |
| Transport | REST | REST + WebSocket (`/ws/session/{id}`) |
| Schema | A2UI Surface (PoC subset, 7 primitives) | A2UI Surface (full, 12+ components) |
| Drag/drop | `@dnd-kit/core` + `@dnd-kit/sortable` | — |
| UI kit | Tailwind 4 + shadcn/ui | — |

**Forbidden in PoC:** Vercel AI SDK, CopilotKit, LangChain, LangGraph, Postgres, Redis, Docker, K8s. (See PRD §3 "No usar".)

---

## Endpoint contracts (aligned with ARCHITECTURE §2.1)

PoC routes mirror arch §2.1 paths exactly. Frontend swap to Go backend = change base URL only.

| Method | Route | Owner branch | Purpose |
|---|---|---|---|
| `POST` | `/api/capsules/generate` | `feat/capsule-controlled` (PRD 1) | Dual-coded capsule (text + image + audio) |
| `POST` | `/api/sessions/[id]/next` | `feat/quiz-a2ui` (PRD 2) | Next quiz item, IRT-3PL selector targeting P(success)=0.85 |
| `POST` | `/api/sessions/[id]/quiz/answer` | `feat/quiz-a2ui` (PRD 2) | Answer submit, updates `thetaUser` |
| `POST` | `/api/sessions/[id]/socratic/response` | `feat/remediation-open` (PRD 3) | Feynman simplification + Elaborative Interrogation |

Old `/api/agent` intent-switch route is **deprecated**. Each branch creates its own dedicated route under `app/api/...` — zero merge conflicts.

---

## The five pillars (where each lives)

| Pillar | Arch ref | PoC home | Status |
|---|---|---|---|
| **Dual Coding** | §4.3 | PRD 1 — capsule emits `Modality[]`, `length >= 2` | wired |
| **A2UI Surface schema** | §3.1 | PRD 2 — `surfaceId`, `rootComponent`, `components` map (adjacency), `dataModel`, `meta` | wired |
| **85% Rule + IRT-3PL** | §4.2 | PRD 2 — `pSuccess(theta,a,b,c)` + `pickDifficulty(theta)` in `/api/sessions/[id]/next` | wired (simplified) |
| **Socratic (Feynman + EI)** | §4.4 | PRD 3 — prompt forces 2 phases: simplify+analogy, then "¿Por qué…?" + textarea | wired |
| **Cold-start IRT diagnostic** | §4.1 | NOT in PoC — roadmap only | deferred |

Whenever I add UI or an LLM prompt, **name the pillar** it serves (or say "none — pure scaffold").

---

## Generative UI tier selection

When adding a new UI element, ask in this order:

1. Core repeatable entity (capsule, lead, deal)? → **Controlled** component, props-only.
2. Variation / long-tail layout (quiz forms, profile-driven UI)? → **Declarative A2UI** Surface.
3. One-off, throwaway, data-grounded view? → **Open-ended** raw HTML in sandboxed iframe (`sandbox="allow-scripts"` only — no `allow-same-origin`).

---

## Shared types (lib/store.ts) — single source of truth

PRD 1 owns the file. PRD 2 and PRD 3 **extend** the union, never sustitute it.

```typescript
export type GenUIResponse =
  | { kind: "controlled"; component: "StudyCapsule"; capsuleId: string; topic: string; modalities: Modality[]; props: CapsuleProps }
  | { kind: "declarative"; surface: A2UISurface }
  | { kind: "open"; html: string };

export type Modality = { type: "text"|"audio"|"image"|"video"; content: string; metadata?: Record<string, unknown> };

export type A2UISurface = {
  surfaceId: string;
  rootComponent: string;
  components: Record<string, A2UIComponent>;  // adjacency map; children are string[] of IDs
  dataModel?: A2UIDataModel;
  meta: { intent: "quiz"; profile: StudentProfile; correctAnswerEvent: string;
          irt?: { thetaUser: number; difficultyIRT: number; discrimination: number; pSuccessPredicted: number } };
};
```

7 PoC A2UI primitives: `Stack`, `Row`, `Text`, `Image`, `Button`, `TextInput`, `DragItem`. Anything else → silently ignored by renderer.

---

## Hackathon priorities

When in doubt how to spend time, weight in this order:

1. **End-to-end demoable slice** — capsule → quiz → fail → remediation → done. One slice working before deepening any.
2. **Three GenUI tiers visible in one session** — that is the single biggest judging hook.
3. **Pre-cache fallbacks** — every LLM call has a hardcoded fallback. The demo never crashes on a bad JSON.parse or empty HTML.
4. **Narrative** — README + 3-min video matter as much as code. PRD 4 owns this.
5. **Polish last** — only after the slice runs reliably twice in a row.

**Anti-priorities:** auth, persistence, IRT cold-start, BKT, SM-2, Postgres, Docker, K8s, Vercel AI SDK, custom drag-drop physics, multi-language, perfecting the agent prompt before the UI exists.

---

## Branch coordination

| Branch | PRD | Owns | Depends on |
|---|---|---|---|
| `feat/capsule-controlled` | 1 | scaffold, `lib/store.ts`, `<StudyCapsule />`, `/api/capsules/generate` | nothing — runs first |
| `feat/quiz-a2ui` | 2 | A2UI renderer, primitives, `<QuizSurface />`, `/api/sessions/[id]/next`, `/api/sessions/[id]/quiz/answer` | PRD 1 merged to main |
| `feat/remediation-open` | 3 | `<RemediationSurface />`, `/api/sessions/[id]/socratic/response`, sandboxed iframe | PRD 1 merged (PRD 2 ideal but not blocking) |
| `feat/demo-narrative` | 4 | README, hero image, video, slides, submission form | PRDs 1–3 merged by hour 3 |

**Rule of gold:** if at hour 3 something doesn't work, cut features, not quality. Better 2 tiers polished + a mock of the 3rd than 3 half-done.

---

## Working agreement with Claude Code

- Default to **concise output**. Long explanations only on request.
- When changing across PRDs, state the plan in 3–6 bullets first, then execute.
- Don't introduce new dependencies without flagging the tradeoff (size, license, lock-in).
- If a PRD contradicts ARCHITECTURE.md on **stack** → PRD wins (PoC). On **contract shape** (route name, schema field) → ARCHITECTURE.md wins.
- Bias toward **small, reversible commits** with clear messages so rollback during hackathon stays fast.
- Never paste API keys (Anthropic, OpenAI, Gemini) into chat, commits, or logs.
- When proposing UI, name the generative UI tier (controlled / A2UI / open-ended) AND the pillar (dual coding / 85% / Feynman+EI) so we stay deliberate.
- Caveman mode active by default in this repo — terse output, fragments OK, code/commits/security normal.

---

## Legacy starter kit

Original CopilotKit + LangGraph starter lives at `Generative-UI-Global-Hackathon-Starter-Kit/`. NOT used by current build. Keep for reference only — do not edit unless explicitly asked.

---

## Useful references

- A2UI spec: https://a2ui.org/
- Anthropic SDK: https://docs.anthropic.com/
- shadcn/ui: https://ui.shadcn.com/
- @dnd-kit: https://docs.dndkit.com/
- IRT-3PL background: https://en.wikipedia.org/wiki/Item_response_theory
- 85% Rule paper (Wilson et al., 2019): https://www.nature.com/articles/s41467-019-12552-4
- Feynman Technique: https://fs.blog/feynman-technique/

In-repo:
- `docs/ARCHITECTURE.md` — production target
- `docs/PRD/PRD_1.md`–`PRD_4.md` — 4hr execution plan
