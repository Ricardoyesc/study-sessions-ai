# Testing Design — study-sessions-ai

**Date:** 2026-05-09  
**Goal:** Regression safety during 4-hour hackathon sprints  
**Scope:** Go backend (pure logic + HTTP contracts) + Nuxt 3 frontend (composables + utils)  
**Constraint:** Zero mock infra — no external deps, no testcontainers, no Playwright

---

## Decision Record

- **Why tests now:** Four parallel branches (`feat/capsule-controlled`, `feat/quiz-a2ui`, `feat/remediation-open`, `feat/demo-narrative`) risk API contract drift. Tests catch this before demo-day merges.
- **Why skip adapters:** LLM/TTS/Redis/MinIO adapters all hit external services. Mocking them buys little regression safety relative to setup cost in a hackathon.
- **Why not IRT math yet:** `pSuccess(θ,a,b,c)` and `pickDifficulty(θ)` live in handler stubs. Tests get added when real logic lands (PRD 2 sprint).

---

## Tooling

| Layer | Runner | New deps |
|---|---|---|
| Go | stdlib `testing` + `net/http/httptest` + Gin test mode | none |
| Frontend | Vitest | `vitest`, `happy-dom` (devDeps) |

---

## Go Tests

### `src/pkg/a2ui/builder_test.go`

Pure constructor functions — assert `ID`, `Type`, and key `Props` fields.

```
TestNewText_setsTypeAndProps
TestNewRichText_setsMarkdownAndAccessible
TestNewImage_setsUrlAndAltText
TestNewAudioPlayer_setsUrlAndAutoPlayFalse
TestNewColumn_nilProps_setsDefaults
TestNewColumn_customProps_usesProvided
TestNewRow_nilProps_setsDefaults
TestNewCard_nilProps_setsDefaults
TestNewQuizCard_setsQuestionOptionsModeAndOnSubmitEvent
TestNewSocraticDialog_setsPromptContextAndOnSubmitEvent
TestNewProgressBar_setsValueAndMax
TestNewButton_setsLabelAndVariant
TestDefaultDataModel_setsLanguageEsAndThemeSystem
```

Pattern per test:
```go
got := NewText("id1", "Hello", "h1")
assert got.ID == "id1"
assert got.Type == ComponentTypeText
assert got.Props["content"] == "Hello"
assert got.Props["variant"] == "h1"
```

### `src/pkg/a2ui/diff_test.go`

```
TestNewA2UIUpdate_wrapsUpdatesSlice
TestDiffProps_setsComponentIDAndProps
TestDiffProps_emptyProps_stillValid
```

### Handler tests — one file per handler

Each test: spin up bare `gin.New()`, register only the route under test, fire `httptest.NewRecorder`.

**`user_handler_test.go`**
```
TestRegister_validPayload_returns201AndToken
TestRegister_missingEmail_returns400
TestRegister_invalidEmail_returns400
TestRegister_shortPassword_returns400
TestLogin_validPayload_returns200AndToken
TestLogin_missingPassword_returns400
TestMe_returns200WithIdAndEmail
```

**`session_handler_test.go`**
```
TestCreate_returns201
TestGet_returns200
TestNext_returns200
TestUpdateAccessibility_returns200
```

**`capsule_handler_test.go`**
```
TestGenerate_returns201
TestGet_returns200
TestServeAsset_returns200
```

**`quiz_handler_test.go`**
```
TestAnswer_returns200
```

**`socratic_handler_test.go`**
```
TestResponse_returns200
```

Handler test contract: assert **status code** + **top-level JSON keys present** (e.g. `token` for auth routes, `message` for stubs). When handlers get real logic, assertions extend — no rewrite needed.

---

## Frontend Tests

### Vitest config — `ui/vitest.config.ts`

```ts
import { defineConfig } from 'vitest/config'

export default defineConfig({
  test: {
    environment: 'happy-dom',
    globals: true,
  },
})
```

### `package.json` script additions

```json
"test": "vitest run",
"test:watch": "vitest"
```

### `ui/composables/useA2UI.test.ts`

`applyMessage` is exported directly — import and test without mounting Nuxt or opening a WebSocket.

```
test: a2ui_full replaces surface wholesale
test: a2ui_update merges props on existing component
test: a2ui_update skips unknown componentId (no crash)
test: a2ui_update replaces children when update.children provided
test: a2ui_update preserves children when update.children omitted
test: data_model_update shallow-merges into existing dataModel
test: data_model_update ignored when surface is null
```

### `ui/utils/format.test.ts`

```
test: formatPercent(85.6) → "86%"
test: formatPercent(100) → "100%"
test: formatPercent(0) → "0%"
test: formatTheta(1.23) → "+1.23"
test: formatTheta(-0.5) → "-0.50"
test: formatTheta(0) → "0.00"
```

---

## File Structure

```
src/
  pkg/a2ui/
    builder_test.go
    diff_test.go
  internal/api/handlers/
    user_handler_test.go
    session_handler_test.go
    capsule_handler_test.go
    quiz_handler_test.go
    socratic_handler_test.go

ui/
  vitest.config.ts
  composables/
    useA2UI.test.ts
  utils/
    format.test.ts
```

---

## Run Commands

```bash
# Go — from repo root
cd src && go test ./...

# Frontend — from repo root
cd ui && npx vitest run
```

Expected: Go <2s, Vitest <3s. Both run with no network, no DB, no external services.

---

## Extension Points (post-hackathon / when real logic lands)

- `pkg/a2ui/builder_test.go` — add remaining 7 PoC primitives as constructors are added
- `handlers/session_handler_test.go` — extend assertions when session stores real state
- `useA2UI.test.ts` — add WebSocket reconnect logic tests when connection handling firms up
- New file: `composables/useAuth.test.ts` — token parse + expiry logic
- New file: IRT math tests (`pSuccess`, `pickDifficulty`) when PRD 2 handler gets real implementation
