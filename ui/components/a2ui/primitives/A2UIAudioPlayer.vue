<script setup>
defineProps({
  url: {
    type: String,
    default: ''
  },
  autoPlay: {
    type: Boolean,
    default: false
  }
})
</script>

<template>
  <div class="a2ui-media-enter rounded-lg border border-base-300 bg-base-200 p-2 shadow-sm transition duration-200 ease-out hover:-translate-y-0.5 hover:shadow-md motion-reduce:transform-none motion-reduce:transition-none">
    <audio class="w-full" controls :autoplay="autoPlay" :src="url" />
  </div>
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
