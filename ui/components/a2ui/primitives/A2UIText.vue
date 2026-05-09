<script setup>
const props = defineProps({
  content: {
    type: String,
    default: ''
  },
  variant: {
    type: String,
    default: 'body'
  }
})

const tag = computed(() => {
  const tags = {
    h1: 'h1',
    h2: 'h2',
    h3: 'h3',
    h4: 'h4',
    h5: 'h5',
    h6: 'h6',
    label: 'p',
    caption: 'p',
    body: 'p'
  }

  return tags[props.variant] ?? 'p'
})

const textClass = computed(() => {
  const classes = {
    h1: 'text-3xl font-bold text-neutral',
    h2: 'text-2xl font-bold text-neutral',
    h3: 'text-xl font-bold text-neutral',
    h4: 'text-lg font-bold text-neutral',
    h5: 'text-base font-bold text-neutral',
    h6: 'text-sm font-bold text-neutral',
    label: 'text-xs font-semibold uppercase text-primary',
    caption: 'text-xs text-base-content/60',
    body: 'text-sm leading-6 text-base-content/75'
  }

  return classes[props.variant] ?? classes.body
})
</script>

<template>
  <component :is="tag" :class="textClass">
    {{ content }}
  </component>
</template>
