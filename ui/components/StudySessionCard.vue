<script setup lang="ts">
defineProps<{
  title: string
  description: string
  duration: string
  status: string
}>()
</script>

<template>
  <article class="card h-full border border-base-300 bg-base-100 shadow-sm">
    <div class="card-body gap-4 p-5">
      <div class="flex items-start justify-between gap-4">
        <h3 class="card-title text-xl text-neutral">{{ title }}</h3>
        <span class="badge badge-accent whitespace-nowrap">{{ duration }}</span>
      </div>
      <p class="text-sm leading-6 text-base-content/70">{{ description }}</p>
      <div class="card-actions justify-end">
        <span class="badge badge-outline">{{ status }}</span>
      </div>
    </div>
  </article>
</template>