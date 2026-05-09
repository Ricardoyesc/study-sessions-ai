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
    class="a2ui-richtext-enter prose prose-sm max-w-none text-base-content/75"
    :class="accessible ? 'leading-6' : ''"
    v-html="html"
  />
</template>

<style scoped>
.a2ui-richtext-enter {
  animation: a2ui-richtext-enter 180ms ease-out both;
}

@keyframes a2ui-richtext-enter {
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
  .a2ui-richtext-enter {
    animation: none;
  }
}
</style>
