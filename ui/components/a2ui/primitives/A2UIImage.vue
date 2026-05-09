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
  <figure class="overflow-hidden rounded-lg border border-base-300 bg-base-100">
    <img class="aspect-video w-full object-cover" :src="url" :alt="altText">
  </figure>
</template>
