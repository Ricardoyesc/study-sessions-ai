<script setup>
import { formatTheta } from '~/utils/format'

const props = defineProps({
  student: {
    type: Object,
    required: true
  }
})

const weeklyProgress = computed(() => Math.min(Math.round((props.student.studiedMinutes / props.student.weeklyGoalMinutes) * 100), 100))
</script>

<template>
  <section class="grid gap-4 rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm lg:grid-cols-[1.2fr_0.8fr]">
    <div class="min-w-0">
      <p class="text-sm font-semibold uppercase text-primary">Estudiante</p>
      <h1 class="mt-1 truncate text-2xl font-bold text-neutral">{{ student.name }}</h1>
      <p class="mt-1 truncate text-sm text-base-content/60">{{ student.email }}</p>
      <div class="mt-4 flex flex-wrap gap-2">
        <span class="badge badge-primary badge-outline">{{ student.cluster }}</span>
        <span class="badge badge-ghost">theta {{ formatTheta(student.estimatedTheta) }}</span>
      </div>
    </div>

    <div class="grid grid-cols-3 gap-2 text-center">
      <div class="rounded-lg bg-base-200 p-3">
        <p class="text-xs text-base-content/60">Meta</p>
        <p class="text-lg font-bold text-neutral">{{ weeklyProgress }}%</p>
      </div>
      <div class="rounded-lg bg-base-200 p-3">
        <p class="text-xs text-base-content/60">Minutos</p>
        <p class="text-lg font-bold text-neutral">{{ student.studiedMinutes }}</p>
      </div>
      <div class="rounded-lg bg-base-200 p-3">
        <p class="text-xs text-base-content/60">Racha</p>
        <p class="text-lg font-bold text-neutral">{{ student.streakDays }}</p>
      </div>
    </div>
  </section>
</template>
