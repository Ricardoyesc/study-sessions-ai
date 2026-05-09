<script setup>
const props = defineProps({
  prompt: {
    type: String,
    default: 'Pregunta o concepto'
  },
  answer: {
    type: String,
    default: 'Respuesta correcta'
  },
  promptLabel: {
    type: String,
    default: 'Frente'
  },
  answerLabel: {
    type: String,
    default: 'Respuesta correcta'
  },
  initiallyRevealed: {
    type: Boolean,
    default: false
  },
  tone: {
    type: String,
    default: 'base'
  }
})

const emit = defineEmits(['flip'])
const revealed = ref(props.initiallyRevealed)

const cardClass = computed(() => {
  const tones = {
    primary: 'border-primary/25 bg-primary/5',
    secondary: 'border-secondary/25 bg-secondary/5',
    accent: 'border-accent/25 bg-accent/5',
    warning: 'border-warning/30 bg-warning/10',
    base: 'border-base-300 bg-base-100'
  }

  return [
    'min-h-52 rounded-lg border p-5 text-left shadow-sm transition duration-200 hover:shadow-md focus:outline-none focus:ring-2 focus:ring-primary/40',
    tones[props.tone] ?? tones.base
  ]
})

const currentLabel = computed(() => revealed.value ? props.answerLabel : props.promptLabel)
const currentContent = computed(() => revealed.value ? props.answer : props.prompt)
const actionLabel = computed(() => revealed.value ? 'Ver pregunta' : 'Ver respuesta')

function flipCard() {
  revealed.value = !revealed.value
  emit('flip', {
    revealed: revealed.value,
    visibleSide: revealed.value ? 'answer' : 'prompt'
  })
}
</script>

<template>
  <button :class="cardClass" type="button" :aria-pressed="revealed" @click="flipCard">
    <span class="text-xs font-semibold uppercase text-primary">{{ currentLabel }}</span>
    <span class="mt-4 block text-xl font-bold leading-snug text-neutral">{{ currentContent }}</span>
    <span class="mt-6 inline-flex rounded-full bg-base-200 px-3 py-1 text-xs font-semibold text-base-content/70">
      {{ actionLabel }}
    </span>
  </button>
</template>