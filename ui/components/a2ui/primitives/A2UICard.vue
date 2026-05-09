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
    primary: 'border-primary/35 bg-primary/10',
    secondary: 'border-secondary/35 bg-secondary/10',
    accent: 'border-accent/40 bg-accent/10',
    warning: 'border-warning/45 bg-warning/15',
    base: 'border-base-300 bg-base-200/70'
  }

  return [
    'rounded-lg border transition duration-200 ease-out hover:-translate-y-0.5 hover:shadow-md motion-reduce:transform-none motion-reduce:transition-none',
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
