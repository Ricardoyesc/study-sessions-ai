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

  return ['btn btn-sm', variants[props.variant] ?? variants.primary]
})
</script>

<template>
  <button :class="buttonClass" type="button" :disabled="disabled" @click="emit('click')">
    {{ label }}
  </button>
</template>
