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
  <section class="a2ui-panel-enter rounded-lg border border-base-300 bg-base-100 p-4 transition duration-200 ease-out hover:shadow-md motion-reduce:transition-none">
    <label class="mb-2 block text-sm font-semibold text-neutral">
      {{ prompt }}
    </label>
    <textarea
      v-model="response"
      class="textarea textarea-bordered min-h-24 w-full resize-none transition duration-200 ease-out focus:-translate-y-0.5 focus:shadow-sm motion-reduce:transform-none motion-reduce:transition-none"
      :placeholder="placeholder"
    />
    <div class="mt-3 flex justify-end">
      <button class="btn btn-secondary btn-sm" type="button" :disabled="!response.trim()" @click="submitResponse">
        Guardar reflexion
      </button>
    </div>
  </section>
</template>

<style scoped>
.a2ui-panel-enter {
  animation: a2ui-panel-enter 200ms ease-out both;
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
