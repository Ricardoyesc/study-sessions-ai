<script setup>
const props = defineProps({
  question: {
    type: String,
    default: ''
  },
  options: {
    type: Array,
    default: () => []
  },
  mode: {
    type: String,
    default: 'single_choice'
  }
})

const emit = defineEmits(['submit'])
const selected = ref([])

const normalizedOptions = computed(() => props.options.filter((option) => typeof option === 'string'))

function toggleOption(option) {
  if (props.mode === 'multi_choice') {
    selected.value = selected.value.includes(option)
      ? selected.value.filter((item) => item !== option)
      : [...selected.value, option]
    return
  }

  selected.value = [option]
}

function submitAnswer() {
  emit('submit', {
    mode: props.mode,
    answer: props.mode === 'multi_choice' ? selected.value : selected.value[0] ?? null
  })
}
</script>

<template>
  <section class="rounded-lg border border-base-300 bg-base-100 p-4">
    <p class="mb-3 font-semibold text-neutral">{{ question }}</p>

    <div class="grid gap-2">
      <button
        v-for="option in normalizedOptions"
        :key="option"
        class="btn justify-start text-left"
        :class="selected.includes(option) ? 'btn-primary' : 'btn-outline'"
        type="button"
        @click="toggleOption(option)"
      >
        {{ option }}
      </button>
    </div>

    <div class="mt-4 flex justify-end">
      <button class="btn btn-secondary btn-sm" type="button" :disabled="selected.length === 0" @click="submitAnswer">
        Enviar respuesta
      </button>
    </div>
  </section>
</template>
