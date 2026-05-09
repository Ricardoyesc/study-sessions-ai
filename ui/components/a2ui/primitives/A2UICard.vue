<script setup>
const props = defineProps({
  elevation: {
    type: Number,
    default: 0
  },
  tone: {
    type: String,
    default: 'base'
  },
  padding: {
    type: Number,
    default: 16
  }
})

const cardClass = computed(() => {
  const tones = {
    primary: 'border-primary/20 bg-primary/5',
    secondary: 'border-secondary/20 bg-secondary/5',
    accent: 'border-accent/25 bg-accent/5',
    warning: 'border-warning/30 bg-warning/10',
    base: 'border-base-300 bg-base-100'
  }

  return [
    'rounded-lg border',
    tones[props.tone] ?? tones.base,
    props.elevation > 0 ? 'shadow-sm' : ''
  ]
})

const cardStyle = computed(() => ({ padding: `${props.padding}px` }))
</script>

<template>
  <article :class="cardClass" :style="cardStyle">
    <div class="grid gap-3">
      <slot />
    </div>
  </article>
</template>
