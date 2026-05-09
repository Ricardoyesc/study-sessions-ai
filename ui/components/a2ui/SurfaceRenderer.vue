<script setup>

defineOptions({ name: 'SurfaceRenderer' })

const props = defineProps({
  surface: {
    type: Object,
    required: true
  },
  componentId: {
    type: String,
    default: undefined
  }
})

const component = computed(() => {
  return props.surface.components[props.componentId ?? props.surface.rootComponent] ?? null
})

const children = computed(() => component.value?.children ?? [])

function propString(name, fallback = '') {
  const value = component.value?.props[name]
  return typeof value === 'string' ? value : fallback
}

function propNumber(name, fallback = 0) {
  const value = component.value?.props[name]
  return typeof value === 'number' ? value : fallback
}

function propBoolean(name, fallback = false) {
  const value = component.value?.props[name]
  return typeof value === 'boolean' ? value : fallback
}

function propStringArray(name) {
  const value = component.value?.props[name]
  return Array.isArray(value) ? value.filter((item) => typeof item === 'string') : []
}

const columnStyle = computed(() => ({ gap: `${propNumber('gap', 12)}px` }))

const rowClass = computed(() => {
  const alignment = propString('alignment', 'start')
  return [
    'flex flex-col gap-3 sm:flex-row sm:items-center',
    alignment === 'space-between' ? 'sm:justify-between' : '',
    alignment === 'center' ? 'sm:justify-center' : ''
  ]
})

const textClass = computed(() => {
  const variant = propString('variant', 'body')
  const classes = {
    h1: 'text-3xl font-bold text-neutral',
    h2: 'text-2xl font-bold text-neutral',
    h3: 'text-xl font-bold text-neutral',
    label: 'text-xs font-semibold uppercase text-primary',
    caption: 'text-xs text-base-content/60',
    body: 'text-sm leading-6 text-base-content/75'
  }

  return classes[variant] ?? classes.body
})

const cardClass = computed(() => {
  const tone = propString('tone', 'base')
  const tones = {
    primary: 'border-primary/20 bg-primary/5',
    accent: 'border-accent/25 bg-accent/5',
    base: 'border-base-300 bg-base-100'
  }

  return ['rounded-lg border p-4', tones[tone] ?? tones.base]
})

function markdownToHtml(markdown) {
  return markdown
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\n/g, '<br>')
}
</script>

<template>
  <div v-if="component">
    <div v-if="component.type === 'Column'" class="flex flex-col" :style="columnStyle">
      <SurfaceRenderer
        v-for="childId in children"
        :key="childId"
        :surface="surface"
        :component-id="childId"
      />
    </div>

    <div v-else-if="component.type === 'Row'" :class="rowClass">
      <SurfaceRenderer
        v-for="childId in children"
        :key="childId"
        :surface="surface"
        :component-id="childId"
      />
    </div>

    <article v-else-if="component.type === 'Card'" :class="cardClass">
      <div class="grid gap-3">
        <SurfaceRenderer
          v-for="childId in children"
          :key="childId"
          :surface="surface"
          :component-id="childId"
        />
      </div>
    </article>

    <p v-else-if="component.type === 'Text'" :class="textClass">
      {{ propString('content') }}
    </p>

    <div
      v-else-if="component.type === 'RichText'"
      class="prose prose-sm max-w-none text-base-content/75"
      v-html="markdownToHtml(propString('markdown'))"
    />

    <div v-else-if="component.type === 'ProgressBar'" class="min-w-[12rem]">
      <div class="mb-2 flex items-center justify-between gap-3 text-xs text-base-content/60">
        <span>{{ propString('label', 'Progreso') }}</span>
        <span>{{ Math.round((propNumber('value') / Math.max(propNumber('max', 1), 1)) * 100) }}%</span>
      </div>
      <progress class="progress progress-primary h-2 w-full" :value="propNumber('value')" :max="propNumber('max', 1)" />
    </div>

    <div v-else-if="component.type === 'Button'" class="pt-1">
      <button class="btn btn-primary btn-sm" type="button">
        {{ propString('label', 'Continuar') }}
      </button>
    </div>

    <div v-else-if="component.type === 'QuizCard'" class="rounded-lg border border-base-300 bg-base-100 p-4">
      <p class="mb-3 font-semibold text-neutral">{{ propString('question') }}</p>
      <div class="grid gap-2">
        <button
          v-for="option in propStringArray('options')"
          :key="option"
          class="btn btn-outline justify-start"
          type="button"
        >
          {{ option }}
        </button>
      </div>
    </div>

    <div v-else-if="component.type === 'SocraticDialog'" class="rounded-lg border border-base-300 bg-base-100 p-4">
      <label :for="component.id" class="mb-2 block text-sm font-semibold text-neutral">
        {{ propString('prompt') }}
      </label>
      <textarea
        :id="component.id"
        class="textarea textarea-bordered min-h-24 w-full resize-none"
        :placeholder="propString('placeholder')"
        :disabled="propBoolean('disabled')"
      />
      <div class="mt-3 flex justify-end">
        <button class="btn btn-secondary btn-sm" type="button">Guardar reflexion</button>
      </div>
    </div>

    <img
      v-else-if="component.type === 'Image'"
      class="w-full rounded-lg border border-base-300 object-cover"
      :src="propString('url')"
      :alt="propString('altText')"
    >

    <audio v-else-if="component.type === 'AudioPlayer'" class="w-full" controls :src="propString('url')" />

    <video v-else-if="component.type === 'VideoPlayer'" class="w-full rounded-lg border border-base-300" controls :src="propString('url')" />

    <div v-else class="rounded-lg border border-warning/30 bg-warning/10 p-3 text-sm text-warning-content">
      Componente A2UI no soportado: {{ component.type }}
    </div>
  </div>
</template>
