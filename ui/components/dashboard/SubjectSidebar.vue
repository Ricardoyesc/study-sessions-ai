<script setup>

const props = defineProps({
  subjects: {
    type: Array,
    required: true
  },
  activeSubjectId: {
    type: String,
    required: true
  },
  activeEvaluationId: {
    type: String,
    required: true
  }
})

const emit = defineEmits(['select'])

function statusLabel(status) {
  const labels = {
    completed: 'Completada',
    in_progress: 'En curso',
    pending: 'Pendiente'
  }

  return labels[status] ?? status
}

function badgeClass(status) {
  const classes = {
    completed: 'badge-success badge-outline',
    in_progress: 'badge-primary',
    pending: 'badge-ghost'
  }

  return classes[status] ?? 'badge-ghost'
}
</script>

<template>
  <nav class="flex h-full flex-col rounded-lg border border-base-300 bg-base-100 shadow-sm">
    <div class="border-b border-base-300 p-4">
      <p class="text-sm font-semibold uppercase text-primary">Materias</p>
      <p class="mt-1 text-sm text-base-content/60">Evaluaciones y focos activos</p>
    </div>

    <div class="min-h-0 flex-1 overflow-auto p-3">
      <section v-for="subject in props.subjects" :key="subject.id" class="mb-3 rounded-lg border border-base-300 bg-base-100">
        <div class="border-b border-base-300 p-3">
          <div class="flex items-start justify-between gap-3">
            <div class="min-w-0">
              <h2 class="truncate text-sm font-bold text-neutral">{{ subject.name }}</h2>
              <p class="truncate text-xs text-base-content/60">{{ subject.teacher }}</p>
            </div>
            <span class="badge badge-sm">{{ subject.progress }}%</span>
          </div>
          <progress class="progress progress-primary mt-3 h-1.5 w-full" :value="subject.progress" max="100" />
        </div>

        <div class="grid gap-1 p-2">
          <button
            v-for="evaluation in subject.evaluations"
            :key="evaluation.id"
            class="rounded-md p-3 text-left transition hover:bg-base-200"
            :class="evaluation.id === activeEvaluationId ? 'bg-primary/10 ring-1 ring-primary/30' : ''"
            type="button"
            @click="emit('select', subject.id, evaluation.id)"
          >
            <div class="flex items-start justify-between gap-2">
              <span class="text-sm font-semibold text-neutral">{{ evaluation.title }}</span>
              <span class="badge badge-xs whitespace-nowrap" :class="badgeClass(evaluation.status)">
                {{ statusLabel(evaluation.status) }}
              </span>
            </div>
            <div class="mt-2 flex items-center justify-between text-xs text-base-content/60">
              <span>{{ evaluation.lastAttempt }}</span>
              <span>{{ evaluation.score }}/100</span>
            </div>
          </button>
        </div>
      </section>
    </div>
  </nav>
</template>
