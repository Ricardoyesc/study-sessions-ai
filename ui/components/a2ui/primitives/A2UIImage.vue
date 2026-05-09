<script setup>
defineProps({
  url: {
    type: String,
    default: ''
  },
  altText: {
    type: String,
    default: ''
  }
})
</script>

<template>
  <figure class="a2ui-media-enter group overflow-hidden rounded-lg border border-base-300 bg-base-200 shadow-sm transition duration-200 ease-out hover:shadow-md motion-reduce:transition-none">
    <img class="aspect-video w-full object-cover transition duration-300 ease-out group-hover:scale-[1.02] motion-reduce:transform-none motion-reduce:transition-none" :src="url" :alt="altText">
  </figure>
</template>

<style scoped>
.a2ui-media-enter {
  animation: a2ui-media-enter 220ms ease-out both;
}

@keyframes a2ui-media-enter {
  from {
    opacity: 0;
    transform: translateY(5px);
  }

  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@media (prefers-reduced-motion: reduce) {
  .a2ui-media-enter {
    animation: none;
  }
}
</style>
