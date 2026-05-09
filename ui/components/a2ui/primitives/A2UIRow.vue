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
    'a2ui-layout-enter flex flex-col sm:flex-row sm:items-center',
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

<style scoped>
.a2ui-layout-enter {
  animation: a2ui-soft-enter 180ms ease-out both;
}

@keyframes a2ui-soft-enter {
  from {
    opacity: 0;
    transform: translateY(4px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (prefers-reduced-motion: reduce) {
  .a2ui-layout-enter {
    animation: none;
  }
}
</style>
