<script setup>
import type { A2UISurface } from '~/types/a2ui'

const props = defineProps<{
  surface: A2UISurface | null
  state: string
  topic: string
  feedback: string | null
  isCorrect: boolean | null
  loading: boolean
  isActive: boolean
}>()

const emit = defineEmits<{
  start: [topic: string]
  next: []
  answer: [index: number]
  socraticResponse: [response: string]
  close: []
}>()

const inputTopic = ref('Método Científico')
const selectedOption = ref<number | null>(null)
const socraticText = ref('')

const stateLabel: Record<string, string> = {
  idle: 'Sin iniciar',
  capsule: 'Contenido de estudio',
  quiz: 'Pregunta',
  remediation: 'Reflexión',
  completed: 'Completado'
}

const stateColor: Record<string, string> = {
  idle: 'badge-ghost',
  capsule: 'badge-primary',
  quiz: 'badge-accent',
  remediation: 'badge-warning',
  completed: 'badge-success'
}

function handleAnswer(index: number) {
  selectedOption.value = index
  emit('answer', index)
}

async function handleSocraticSubmit() {
  if (!socraticText.value.trim()) return
  emit('socraticResponse', socraticText.value)
  socraticText.value = ''
}

function resetAndNext() {
  selectedOption.value = null
  emit('next')
}
</script>

<template>
  <div class="card bg-base-100 border border-base-300 shadow-lg">
    <div class="card-body p-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-2">
          <h3 class="card-title text-lg">Sesión de Estudio</h3>
          <span v-if="isActive" class="badge" :class="stateColor[state] ?? 'badge-ghost'">
            {{ stateLabel[state] ?? state }}
          </span>
          <span v-if="loading" class="loading loading-spinner loading-sm" />
        </div>
        <button v-if="isActive" class="btn btn-ghost btn-sm" @click="emit('close')">Cerrar</button>
      </div>

      <div v-if="!isActive" class="flex flex-col gap-3 py-4">
        <p class="text-base-content/70">Elige un tema para comenzar una sesión de estudio adaptativa.</p>
        <div class="flex gap-2">
          <input
            v-model="inputTopic"
            type="text"
            placeholder="Tema de estudio..."
            class="input input-bordered flex-1"
            @keyup.enter="emit('start', inputTopic)"
          />
          <button class="btn btn-primary" :disabled="loading" @click="emit('start', inputTopic)">
            Comenzar
          </button>
        </div>
        <div class="flex flex-wrap gap-1">
          <button
            v-for="t in ['Método Científico', 'Física Cuántica', 'Álgebra Lineal', 'Biología Celular', 'Teoría de la Relatividad']"
            :key="t"
            class="btn btn-outline btn-xs"
            @click="inputTopic = t; emit('start', t)"
          >
            {{ t }}
          </button>
        </div>
      </div>

      <div v-else class="flex flex-col gap-3 py-2">
        <div class="text-sm text-base-content/60 flex items-center gap-2">
          <span>Tema: <strong>{{ topic }}</strong></span>
        </div>

        <SurfaceRenderer v-if="surface" :surface="surface" :component-id="surface.rootComponent" />

        <div v-if="feedback" class="alert" :class="isCorrect ? 'alert-success' : 'alert-error'">
          <span>{{ feedback }}</span>
        </div>

        <div class="flex flex-wrap gap-2 pt-2">
          <button
            v-if="state === 'quiz'"
            class="btn btn-primary btn-sm"
            @click="handleAnswer(selectedOption ?? 0)"
          >
            Enviar respuesta ({{ (selectedOption ?? 0) + 1 }})
          </button>

          <div v-if="state === 'remediation' || (feedback && !isCorrect)" class="flex flex-col gap-2 w-full">
            <textarea
              v-model="socraticText"
              class="textarea textarea-bordered w-full"
              rows="3"
              placeholder="Escribe tu reflexión aquí..."
            />
            <button class="btn btn-secondary btn-sm" @click="handleSocraticSubmit">
              Enviar reflexión
            </button>
          </div>

          <button
            v-if="state === 'capsule' || (feedback && isCorrect) || state === 'completed'"
            class="btn btn-secondary btn-sm"
            :disabled="loading"
            @click="resetAndNext"
          >
            {{ state === 'completed' ? 'Reiniciar' : 'Siguiente' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
