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
  <section class="rounded-lg border border-base-300 bg-base-100 p-4">
    <label class="mb-2 block text-sm font-semibold text-neutral">
      {{ prompt }}
    </label>
    <textarea
      v-model="response"
      class="textarea textarea-bordered min-h-24 w-full resize-none"
      :placeholder="placeholder"
    />
    <div class="mt-3 flex justify-end">
      <button class="btn btn-secondary btn-sm" type="button" :disabled="!response.trim()" @click="submitResponse">
        Guardar reflexion
      </button>
    </div>
  </section>
</template>
