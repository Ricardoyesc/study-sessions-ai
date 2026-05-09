<script setup>
import SurfaceRenderer from '~/components/a2ui/SurfaceRenderer.vue'

const props = defineProps({
  surface: {
    type: Object,
    default: null
  },
  state: {
    type: String,
    default: 'idle'
  },
  topic: {
    type: String,
    default: ''
  },
  feedback: {
    type: String,
    default: null
  },
  isCorrect: {
    type: Boolean,
    default: null
  },
  loading: {
    type: Boolean,
    default: false
  },
  isActive: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['start', 'next', 'answer', 'socraticResponse', 'close'])

const inputTopic = ref('Método Científico')
const selectedOption = ref(null)
const socraticText = ref('')

const stateLabel = {
  idle: 'Sin iniciar',
  capsule: 'Contenido de estudio',
  quiz: 'Pregunta',
  remediation: 'Reflexión',
  completed: 'Completado'
}

const stateColor = {
  idle: 'badge-ghost',
  capsule: 'badge-primary',
  quiz: 'badge-accent',
  remediation: 'badge-warning',
  completed: 'badge-success'
}

function handleAnswer(index) {
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
  <div class="card border border-base-300 bg-base-200/70 shadow-lg">
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
            class="study-input-contrast input flex-1 placeholder:text-base-content/45 focus:outline-none"
            @keyup.enter="emit('start', inputTopic)"
          />
          <button class="btn btn-primary shadow-sm" :disabled="loading" @click="emit('start', inputTopic)">
            Comenzar
          </button>
        </div>
        <div class="flex flex-wrap gap-1">
          <button
            v-for="t in ['Método Científico', 'Física Cuántica', 'Álgebra Lineal', 'Biología Celular', 'Teoría de la Relatividad']"
            :key="t"
            class="study-chip-contrast btn btn-xs shadow-sm"
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
            class="btn btn-primary btn-sm shadow-sm"
            @click="handleAnswer(selectedOption ?? 0)"
          >
            Enviar respuesta ({{ (selectedOption ?? 0) + 1 }})
          </button>

          <div v-if="state === 'remediation' || (feedback && !isCorrect)" class="flex flex-col gap-2 w-full">
            <textarea
              v-model="socraticText"
              class="study-input-contrast textarea w-full placeholder:text-base-content/45 focus:outline-none"
              rows="3"
              placeholder="Escribe tu reflexión aquí..."
            />
            <button class="btn btn-secondary btn-sm shadow-sm" @click="handleSocraticSubmit">
              Enviar reflexión
            </button>
          </div>

          <button
            v-if="state === 'capsule' || (feedback && isCorrect) || state === 'completed'"
            class="btn btn-secondary btn-sm shadow-sm"
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

<style scoped>
.study-input-contrast {
  border: 1px solid color-mix(in oklab, var(--color-base-content) 22%, transparent) !important;
  background: var(--color-base-100) !important;
  box-shadow: inset 0 2px 4px color-mix(in oklab, var(--color-base-content) 8%, transparent);
}

.study-input-contrast:focus {
  border-color: var(--color-primary) !important;
  box-shadow: 0 0 0 3px color-mix(in oklab, var(--color-primary) 18%, transparent), 0 1px 2px color-mix(in oklab, var(--color-base-content) 10%, transparent);
}

.study-chip-contrast {
  border: 1px solid color-mix(in oklab, var(--color-primary) 55%, transparent) !important;
  background: color-mix(in oklab, var(--color-primary) 12%, var(--color-base-100)) !important;
  color: var(--color-primary) !important;
}

.study-chip-contrast:hover {
  background: var(--color-primary) !important;
  border-color: var(--color-primary) !important;
  color: var(--color-primary-content) !important;
}
</style>
