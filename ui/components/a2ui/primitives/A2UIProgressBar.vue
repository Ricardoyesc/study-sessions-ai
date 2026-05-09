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
  <div class="a2ui-progress-enter min-w-[12rem]">
    <div class="mb-2 flex items-center justify-between gap-3 text-xs text-base-content/60">
      <span>{{ label }}</span>
      <span>{{ percent }}%</span>
    </div>
    <progress class="a2ui-progress progress progress-primary h-2 w-full" :value="value" :max="safeMax" />
  </div>
</template>

<style scoped>
.a2ui-progress-enter {
  animation: a2ui-progress-enter 200ms ease-out both;
}

.a2ui-progress::-webkit-progress-value {
  transition: width 260ms ease-out;
}

.a2ui-progress::-moz-progress-bar {
  transition: width 260ms ease-out;
}

@keyframes a2ui-progress-enter {
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
  .a2ui-progress-enter {
    animation: none;
  }

  .a2ui-progress::-webkit-progress-value,
  .a2ui-progress::-moz-progress-bar {
    transition: none;
  }
}
</style>
