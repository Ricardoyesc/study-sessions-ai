<script setup>

const props = defineProps({
  student: {
    type: Object,
    required: true
  },
  authMode: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['logout'])

const initials = computed(() => {
  return props.student.name
    .split(' ')
    .map((part) => part.charAt(0))
    .join('')
    .slice(0, 2)
    .toUpperCase()
})
</script>

<template>
  <div data-theme="light" class="min-h-screen bg-base-200 text-base-content">
    <header class="sticky top-0 z-30 border-b border-base-300 bg-base-100/95 backdrop-blur">
      <div class="mx-auto flex min-h-16 max-w-[1440px] items-center justify-between gap-4 px-4 md:px-6">
        <div class="flex min-w-0 items-center gap-3">
          <div class="grid h-10 w-10 shrink-0 place-items-center rounded-lg bg-primary text-sm font-bold text-primary-content">
            SAI
          </div>
          <div class="min-w-0">
            <p class="truncate text-sm font-semibold text-neutral">Study Sessions AI</p>
            <p class="truncate text-xs text-base-content/60">Panel adaptativo del estudiante</p>
          </div>
        </div>

        <div class="flex items-center gap-3">
          <span class="badge hidden sm:inline-flex" :class="authMode === 'api' ? 'badge-success' : 'badge-warning'">
            {{ authMode === 'api' ? 'API conectada' : 'Modo demo' }}
          </span>
          <div class="avatar placeholder">
            <div class="w-10 rounded-lg bg-neutral text-neutral-content">
              <span class="text-sm">{{ initials }}</span>
            </div>
          </div>
          <button class="btn btn-ghost btn-sm" type="button" @click="emit('logout')">Salir</button>
        </div>
      </div>
    </header>

    <div class="mx-auto grid max-w-[1440px] grid-cols-1 gap-4 px-4 py-4 md:grid-cols-[20rem_1fr] md:px-6">
      <aside class="md:sticky md:top-20 md:h-[calc(100vh-6rem)]">
        <slot name="sidebar" />
      </aside>

      <main class="grid min-w-0 gap-4">
        <slot />
      </main>
    </div>
  </div>
</template>
