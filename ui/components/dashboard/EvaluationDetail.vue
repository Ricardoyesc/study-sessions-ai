<script setup>
import SurfaceRenderer from '~/components/a2ui/SurfaceRenderer.vue'

const props = defineProps({
  subject: {
    type: Object,
    required: true
  },
  evaluation: {
    type: Object,
    required: true
  },
  surface: {
    type: Object,
    default: null
  },
  agentStatus: {
    type: String,
    required: true
  },
  websocketStatus: {
    type: String,
    required: true
  }
})

const delta = computed(() => props.evaluation.targetScore - props.evaluation.score)

const masteryBadge = computed(() => {
  const classes = {
    strong: 'badge-success',
    steady: 'badge-info',
    needs_attention: 'badge-warning'
  }

  return classes[props.evaluation.mastery] ?? 'badge-ghost'
})

const masteryLabel = computed(() => {
  const labels = {
    strong: 'Dominio alto',
    steady: 'Estable',
    needs_attention: 'Requiere foco'
  }

  return labels[props.evaluation.mastery] ?? props.evaluation.mastery
})
</script>

<template>
  <section class="grid gap-4">
    <div class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
      <div class="flex flex-col gap-4 lg:flex-row lg:items-start lg:justify-between">
        <div class="min-w-0">
          <p class="text-sm font-semibold uppercase text-primary">{{ subject.name }}</p>
          <h2 class="mt-1 text-2xl font-bold text-neutral">{{ evaluation.title }}</h2>
          <p class="mt-2 max-w-3xl text-sm leading-6 text-base-content/70">{{ evaluation.description }}</p>
        </div>

        <div class="grid min-w-[14rem] grid-cols-2 gap-2 text-center">
          <div class="rounded-lg bg-base-200 p-3">
            <p class="text-xs text-base-content/60">Resultado</p>
            <p class="text-xl font-bold text-neutral">{{ evaluation.score }}%</p>
          </div>
          <div class="rounded-lg bg-base-200 p-3">
            <p class="text-xs text-base-content/60">Objetivo</p>
            <p class="text-xl font-bold text-neutral">{{ evaluation.targetScore }}%</p>
          </div>
        </div>
      </div>

      <div class="mt-4 flex flex-wrap items-center gap-2">
        <span class="badge" :class="masteryBadge">{{ masteryLabel }}</span>
        <span class="badge badge-outline">Brecha {{ Math.max(delta, 0) }} pts</span>
        <span class="badge badge-ghost">{{ evaluation.insight.estimatedMinutes }} min sugeridos</span>
      </div>
    </div>

    <div class="grid gap-4 lg:grid-cols-[0.85fr_1.15fr]">
      <article class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
        <p class="text-sm font-semibold uppercase text-primary">Diagnostico</p>
        <h3 class="mt-2 text-xl font-bold text-neutral">{{ evaluation.insight.focusPoint }}</h3>
        <p class="mt-3 text-sm leading-6 text-base-content/70">{{ evaluation.insight.reason }}</p>

        <div class="mt-4 rounded-lg bg-base-200 p-4">
          <p class="text-xs font-semibold uppercase text-base-content/60">Siguiente accion</p>
          <p class="mt-2 text-sm leading-6 text-base-content/75">{{ evaluation.insight.nextAction }}</p>
        </div>
      </article>

      <article class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
        <div class="mb-4 flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <p class="text-sm font-semibold uppercase text-primary">UI del agente</p>
            <h3 class="text-xl font-bold text-neutral">Intervencion adaptativa</h3>
          </div>
          <div class="flex flex-wrap gap-2">
            <span class="badge" :class="agentStatus === 'remote' ? 'badge-success' : 'badge-warning'">
              {{ agentStatus === 'remote' ? 'A2UI remoto' : 'Fallback TS' }}
            </span>
            <span class="badge badge-ghost">WS {{ websocketStatus }}</span>
          </div>
        </div>

        <SurfaceRenderer v-if="surface" :surface="surface" />
        <div v-else class="rounded-lg border border-base-300 bg-base-200 p-4 text-sm text-base-content/70">
          Preparando la superficie A2UI de esta evaluacion.
        </div>
      </article>
    </div>
  </section>
</template>
