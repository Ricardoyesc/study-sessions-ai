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

function optionClass(option) {
  return selected.value.includes(option) ? 'a2ui-option-selected' : 'a2ui-option-idle'
}

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
  <section class="a2ui-panel-enter rounded-lg border border-base-300 bg-base-200/70 p-4 shadow-sm transition duration-200 ease-out hover:shadow-md motion-reduce:transition-none">
    <p class="mb-3 font-semibold text-neutral">{{ question }}</p>

    <div class="grid gap-2">
      <button
        v-for="option in normalizedOptions"
        :key="option"
        class="btn justify-start border text-left transition duration-200 ease-out hover:-translate-y-0.5 active:translate-y-0 active:scale-[0.99] motion-reduce:transform-none motion-reduce:transition-none"
        :class="optionClass(option)"
        type="button"
        @click="toggleOption(option)"
      >
        {{ option }}
      </button>
    </div>

    <div class="mt-4 flex justify-end">
      <button class="btn btn-secondary btn-sm shadow-sm disabled:bg-base-300 disabled:text-base-content/40" type="button" :disabled="selected.length === 0" @click="submitAnswer">
        Enviar respuesta
      </button>
    </div>
  </section>
</template>

<style scoped>
.a2ui-panel-enter {
  animation: a2ui-panel-enter 200ms ease-out both;
}

.a2ui-option-idle {
  border-color: color-mix(in oklab, var(--color-base-content) 20%, transparent) !important;
  background: color-mix(in oklab, var(--color-base-100) 88%, var(--color-base-content) 6%) !important;
  color: var(--color-neutral) !important;
  box-shadow: 0 1px 2px color-mix(in oklab, var(--color-base-content) 10%, transparent);
}

.a2ui-option-idle:hover {
  border-color: color-mix(in oklab, var(--color-primary) 55%, transparent) !important;
  background: color-mix(in oklab, var(--color-primary) 12%, var(--color-base-100)) !important;
}

.a2ui-option-selected {
  border-color: var(--color-primary) !important;
  background: var(--color-primary) !important;
  color: var(--color-primary-content) !important;
  box-shadow: 0 0 0 3px color-mix(in oklab, var(--color-primary) 22%, transparent), 0 8px 18px color-mix(in oklab, var(--color-base-content) 14%, transparent);
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
