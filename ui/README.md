# Study Sessions AI UI

Nuxt UI for the student-facing adaptive study flow. The first screen is login; after authentication, the student sees a dashboard with profile information, subjects, evaluations, and an A2UI-rendered improvement panel.

## Setup

```bash
npm install
```

## Development

```bash
npm run dev -- --host 127.0.0.1
```

Open `http://127.0.0.1:3000/`.

The UI tries to use the Go backend from the repository `src` folder:

```bash
NUXT_PUBLIC_API_BASE=http://127.0.0.1:8080
NUXT_PUBLIC_WS_BASE=ws://127.0.0.1:8080
```

Current backend integration:

- `POST /api/users/login` is used for login when the backend is available.
- `GET /api/users/me` is used to hydrate profile details when available.
- `WS /ws/session/:sessionId` is prepared for A2UI messages.

Subjects, evaluations, and generated A2UI surfaces currently use TypeScript fixtures in `data/student-fixtures.ts` until the Go backend exposes those contracts.

## Build

```bash
npm run build
```
