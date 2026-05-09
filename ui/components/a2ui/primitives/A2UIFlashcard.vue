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

const frameClass = computed(() => {
  const tones = {
    primary: 'border-primary/25',
    secondary: 'border-secondary/25',
    accent: 'border-accent/25',
    warning: 'border-warning/30',
    base: 'border-base-300'
  }

  return [
    'group h-60 w-full min-w-0 rounded-lg border text-left perspective-normal focus:outline-none focus:ring-2 focus:ring-primary/40',
    tones[props.tone] ?? tones.base
  ]
})

const faceClass = computed(() => {
  const tones = {
    primary: 'bg-primary/5',
    secondary: 'bg-secondary/5',
    accent: 'bg-accent/5',
    warning: 'bg-warning/10',
    base: 'bg-base-100'
  }

  return [
    'flashcard-face grid content-between p-5 shadow-sm',
    tones[props.tone] ?? tones.base
  ]
})

const flipStyle = computed(() => ({
  transform: revealed.value ? 'rotateY(180deg)' : 'rotateY(0deg)'
}))

function flipCard() {
  revealed.value = !revealed.value
  emit('flip', {
    revealed: revealed.value,
    visibleSide: revealed.value ? 'answer' : 'prompt'
  })
}
</script>

<template>
  <button :class="frameClass" type="button" :aria-pressed="revealed" @click="flipCard">
    <span class="sr-only">{{ revealed ? 'Ver pregunta' : 'Ver respuesta' }}</span>
    <span class="flashcard-inner" :style="flipStyle">
      <span :class="[faceClass, 'flashcard-face-front']" :aria-hidden="revealed">
        <span>
          <span class="text-xs font-semibold uppercase text-primary">{{ promptLabel }}</span>
          <span class="mt-4 block text-xl font-bold leading-snug text-neutral break-words">{{ prompt }}</span>
        </span>
        <span class="inline-flex w-fit rounded-full bg-base-200 px-3 py-1 text-xs font-semibold text-base-content/70">
          Ver respuesta
        </span>
      </span>

      <span :class="[faceClass, 'flashcard-face-back']" :aria-hidden="!revealed">
        <span>
          <span class="text-xs font-semibold uppercase text-primary">{{ answerLabel }}</span>
          <span class="mt-4 block text-xl font-bold leading-snug text-neutral break-words">{{ answer }}</span>
        </span>
        <span class="inline-flex w-fit rounded-full bg-base-200 px-3 py-1 text-xs font-semibold text-base-content/70">
          Ver pregunta
        </span>
      </span>
    </span>
  </button>
</template>

<style scoped>
.perspective-normal {
  perspective: 1000px;
  transform-style: preserve-3d;
}

.flashcard-inner {
  position: relative;
  display: block;
  width: 100%;
  height: 100%;
  transition: transform 500ms ease-out;
  transform-style: preserve-3d;
  -webkit-transform-style: preserve-3d;
  will-change: transform;
}

.flashcard-face {
  position: absolute;
  inset: 0;
  overflow: hidden;
  border-radius: 0.5rem;
  backface-visibility: hidden;
  -webkit-backface-visibility: hidden;
  transform-style: preserve-3d;
  -webkit-transform-style: preserve-3d;
}

.flashcard-face-front {
  transform: rotateY(0deg) translateZ(1px);
}

.flashcard-face-back {
  transform: rotateY(180deg) translateZ(1px);
}

@media (prefers-reduced-motion: reduce) {
  .flashcard-inner {
    transition-duration: 1ms;
  }
}
</style>