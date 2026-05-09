<script setup>
const props = defineProps({
  markdown: {
    type: String,
    default: ''
  },
  accessible: {
    type: Boolean,
    default: true
  }
})

const html = computed(() => {
  return props.markdown
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\n/g, '<br>')
})
</script>

<template>
  <div
    class="prose prose-sm max-w-none text-base-content/75"
    :class="accessible ? 'leading-6' : ''"
    v-html="html"
  />
</template>
