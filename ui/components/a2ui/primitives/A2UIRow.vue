<script setup>
const props = defineProps({
  gap: {
    type: Number,
    default: 12
  },
  alignment: {
    type: String,
    default: 'start'
  },
  wrap: {
    type: Boolean,
    default: true
  }
})

const rowClass = computed(() => {
  const alignmentClasses = {
    start: 'justify-start',
    center: 'justify-center',
    end: 'justify-end',
    'space-between': 'justify-between'
  }

  return [
    'flex flex-col sm:flex-row sm:items-center',
    props.wrap ? 'flex-wrap' : 'flex-nowrap',
    alignmentClasses[props.alignment] ?? alignmentClasses.start
  ]
})

const rowStyle = computed(() => ({ gap: `${props.gap}px` }))
</script>

<template>
  <div :class="rowClass" :style="rowStyle">
    <slot />
  </div>
</template>
