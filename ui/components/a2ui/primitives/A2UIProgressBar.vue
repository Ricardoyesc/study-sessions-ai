<script setup>
const props = defineProps({
  value: {
    type: Number,
    default: 0
  },
  max: {
    type: Number,
    default: 1
  },
  label: {
    type: String,
    default: 'Progreso'
  }
})

const safeMax = computed(() => Math.max(props.max, 1))
const percent = computed(() => Math.round((props.value / safeMax.value) * 100))
</script>

<template>
  <div class="min-w-[12rem]">
    <div class="mb-2 flex items-center justify-between gap-3 text-xs text-base-content/60">
      <span>{{ label }}</span>
      <span>{{ percent }}%</span>
    </div>
    <progress class="progress progress-primary h-2 w-full" :value="value" :max="safeMax" />
  </div>
</template>
