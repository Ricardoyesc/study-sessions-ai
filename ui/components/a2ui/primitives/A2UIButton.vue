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
    primary: 'btn-primary',
    secondary: 'btn-secondary',
    accent: 'btn-accent',
    outline: 'btn-outline',
    ghost: 'btn-ghost'
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
