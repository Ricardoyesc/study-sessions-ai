<script setup>
const props = defineProps({
  prompt: {
    type: String,
    default: ''
  },
  context: {
    type: String,
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Escribe tu respuesta...'
  }
})

const emit = defineEmits(['submit'])
const response = ref('')

function submitResponse() {
  emit('submit', {
    response: response.value,
    context: props.context
  })
}
</script>

<template>
  <section class="a2ui-panel-enter rounded-lg border border-base-300 bg-base-200/70 p-4 shadow-sm transition duration-200 ease-out hover:shadow-md motion-reduce:transition-none">
    <label class="mb-2 block text-sm font-semibold text-neutral">
      {{ prompt }}
    </label>
    <textarea
      v-model="response"
      class="a2ui-textarea-contrast textarea min-h-24 w-full resize-none transition duration-200 ease-out placeholder:text-base-content/45 focus:-translate-y-0.5 focus:shadow-sm focus:outline-none motion-reduce:transform-none motion-reduce:transition-none"
      :placeholder="placeholder"
    />
    <div class="mt-3 flex justify-end">
      <button class="btn btn-secondary btn-sm shadow-sm disabled:bg-base-300 disabled:text-base-content/40" type="button" :disabled="!response.trim()" @click="submitResponse">
        Guardar reflexion
      </button>
    </div>
  </section>
</template>

<style scoped>
.a2ui-panel-enter {
  animation: a2ui-panel-enter 200ms ease-out both;
}

.a2ui-textarea-contrast {
  border: 1px solid color-mix(in oklab, var(--color-base-content) 22%, transparent) !important;
  background: var(--color-base-100) !important;
  box-shadow: inset 0 2px 4px color-mix(in oklab, var(--color-base-content) 8%, transparent);
}

.a2ui-textarea-contrast:focus {
  border-color: var(--color-primary) !important;
  box-shadow: 0 0 0 3px color-mix(in oklab, var(--color-primary) 18%, transparent), 0 1px 2px color-mix(in oklab, var(--color-base-content) 10%, transparent);
}

@keyframes a2ui-panel-enter {
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
  .a2ui-panel-enter {
    animation: none;
  }
}
</style>
