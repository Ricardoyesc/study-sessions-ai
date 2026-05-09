<script setup>
import ThemeSelector from '~/components/layout/ThemeSelector.vue'
import { useTheme } from '~/composables/useTheme'

const email = ref('sofia.martinez@example.com')
const password = ref('demo-password')
const formError = ref(null)

const { isAuthenticated, loading, errorMessage, login } = useAuth()
const { currentTheme } = useTheme()

onMounted(() => {
  if (isAuthenticated.value) {
    navigateTo('/dashboard')
  }
})

async function submitLogin() {
  formError.value = null

  if (!email.value.includes('@')) {
    formError.value = 'Ingresa un email valido.'
    return
  }

  if (password.value.length < 6) {
    formError.value = 'La contrasena debe tener al menos 6 caracteres.'
    return
  }

  await login(email.value, password.value)
  await navigateTo('/dashboard')
}
</script>

<template>
  <div :data-theme="currentTheme" class="min-h-screen bg-base-200 text-base-content">
    <header class="absolute right-4 top-4 z-20 sm:right-6 sm:top-6">
      <ThemeSelector />
    </header>

    <main class="grid min-h-screen lg:grid-cols-[0.95fr_1.05fr]">
      <section class="flex items-center border-b border-base-300 bg-base-100 px-4 py-10 lg:border-b-0 lg:border-r lg:px-10">
        <div class="mx-auto w-full max-w-md">
          <div class="mb-8 flex items-center gap-3">
            <div class="grid h-11 w-11 place-items-center rounded-lg bg-primary text-sm font-bold text-primary-content">
              SAI
            </div>
            <div>
              <p class="font-bold text-neutral">Study Sessions AI</p>
              <p class="text-sm text-base-content/60">Acceso estudiante</p>
            </div>
          </div>

          <h1 class="text-3xl font-bold text-neutral">Inicia sesion para continuar tu plan adaptativo</h1>
          <p class="mt-3 text-sm leading-6 text-base-content/70">
            El panel carga tu perfil, materias y evaluaciones para que el agente proponga una intervencion enfocada.
          </p>

          <form class="mt-8 grid gap-4" @submit.prevent="submitLogin">
            <label class="form-control">
              <span class="label-text font-medium">Email</span>
              <input
                v-model="email"
                class="input input-bordered mt-1"
                type="email"
                autocomplete="email"
                required
              >
            </label>

            <label class="form-control">
              <span class="label-text font-medium">Contrasena</span>
              <input
                v-model="password"
                class="input input-bordered mt-1"
                type="password"
                autocomplete="current-password"
                required
              >
            </label>

            <div v-if="formError || errorMessage" class="alert alert-warning py-3 text-sm">
              <span>{{ formError ?? errorMessage }}</span>
            </div>

            <button class="btn btn-primary mt-2" type="submit" :disabled="loading">
              <span v-if="loading" class="loading loading-spinner loading-sm" />
              Entrar al panel
            </button>
          </form>
        </div>
      </section>

      <section class="flex items-center px-4 py-10 lg:px-10">
        <div class="mx-auto grid w-full max-w-2xl gap-4">
          <div class="rounded-lg border border-base-300 bg-base-100 p-5 shadow-sm">
            <p class="text-sm font-semibold uppercase text-primary">Flujo activo</p>
            <div class="mt-4 grid gap-3">
              <div class="flex items-center justify-between rounded-lg bg-base-200 px-4 py-3">
                <span class="font-medium text-neutral">Perfil del estudiante</span>
                <span class="badge badge-success badge-outline">Listo</span>
              </div>
              <div class="flex items-center justify-between rounded-lg bg-base-200 px-4 py-3">
                <span class="font-medium text-neutral">Materias y evaluaciones</span>
                <span class="badge badge-primary">Dashboard</span>
              </div>
              <div class="flex items-center justify-between rounded-lg bg-base-200 px-4 py-3">
                <span class="font-medium text-neutral">Intervencion A2UI</span>
                <span class="badge badge-warning badge-outline">Agente</span>
              </div>
            </div>
          </div>

          <div class="grid gap-4 sm:grid-cols-3">
            <div class="rounded-lg border border-base-300 bg-base-100 p-4 text-center shadow-sm">
              <p class="text-xs text-base-content/60">Objetivo</p>
              <p class="text-2xl font-bold text-neutral">85%</p>
            </div>
            <div class="rounded-lg border border-base-300 bg-base-100 p-4 text-center shadow-sm">
              <p class="text-xs text-base-content/60">Materias</p>
              <p class="text-2xl font-bold text-neutral">3</p>
            </div>
            <div class="rounded-lg border border-base-300 bg-base-100 p-4 text-center shadow-sm">
              <p class="text-xs text-base-content/60">Modo</p>
              <p class="text-2xl font-bold text-neutral">A2UI</p>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>