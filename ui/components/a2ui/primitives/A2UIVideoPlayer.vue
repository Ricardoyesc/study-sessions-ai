<script setup>
defineProps({
  url: {
    type: String,
    default: ''
  },
  poster: {
    type: String,
    default: ''
  },
  captionsUrl: {
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
  <video class="w-full rounded-lg border border-base-300 bg-neutral" controls :autoplay="autoPlay" :poster="poster" :src="url">
    <track v-if="captionsUrl" kind="captions" srclang="es" label="Espanol" :src="captionsUrl">
  </video>
</template>
