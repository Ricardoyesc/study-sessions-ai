<script setup>
const props = defineProps({
  gap: {
    type: Number,
    default: 12
  },
  padding: {
    type: Number,
    default: 0
  },
  alignment: {
    type: String,
    default: 'stretch'
  }
})

const columnClass = computed(() => {
  const alignmentClasses = {
    stretch: 'items-stretch',
    start: 'items-start',
    center: 'items-center',
    end: 'items-end'
  }

  return ['flex flex-col', alignmentClasses[props.alignment] ?? alignmentClasses.stretch]
})

const columnStyle = computed(() => ({
  gap: `${props.gap}px`,
  padding: `${props.padding}px`
}))
</script>

<template>
  <div :class="columnClass" :style="columnStyle">
    <slot />
  </div>
</template>
