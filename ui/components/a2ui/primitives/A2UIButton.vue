<script setup>
const props = defineProps({
  label: {
    type: String,
    default: 'Continuar'
  },
  variant: {
    type: String,
    default: 'primary'
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const buttonClass = computed(() => {
  const variants = {
    primary: 'btn-primary shadow-sm',
    secondary: 'btn-secondary shadow-sm',
    accent: 'btn-accent shadow-sm',
    outline: 'a2ui-button-outline',
    ghost: 'btn-ghost bg-base-200/80 hover:bg-base-300'
  }

  return [
    'btn btn-sm transition duration-200 ease-out hover:-translate-y-0.5 active:translate-y-0 active:scale-[0.98] motion-reduce:transform-none motion-reduce:transition-none',
    variants[props.variant] ?? variants.primary
  ]
})
</script>

<template>
  <button :class="buttonClass" type="button" :disabled="disabled" @click="emit('click')">
    {{ label }}
  </button>
</template>

<style scoped>
.a2ui-button-outline {
  border: 1px solid color-mix(in oklab, var(--color-primary) 55%, transparent) !important;
  background: color-mix(in oklab, var(--color-primary) 12%, var(--color-base-100)) !important;
  color: var(--color-primary) !important;
  box-shadow: 0 1px 2px color-mix(in oklab, var(--color-base-content) 12%, transparent);
}

.a2ui-button-outline:hover {
  background: var(--color-primary) !important;
  border-color: var(--color-primary) !important;
  color: var(--color-primary-content) !important;
}
</style>
