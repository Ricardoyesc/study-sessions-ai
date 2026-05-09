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
  <audio class="w-full" controls :autoplay="autoPlay" :src="url" />
</template>
