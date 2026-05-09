<script setup>
import SurfaceRenderer from '~/components/a2ui/SurfaceRenderer.vue'

const topics = [
  'Dualidad onda-particula',
  'Derivadas como cambio instantaneo',
  'Guerra fria y bloques geopoliticos',
  'Integrales como acumulacion',
  'Relatividad especial'
]

const tones = ['base', 'primary', 'secondary', 'accent', 'warning']
const variants = ['h1', 'h2', 'h3', 'body', 'label', 'caption']
const eventLog = ref([])
const seed = ref(Date.now())

function pick(list) {
  return list[Math.floor(Math.random() * list.length)]
}

function numberBetween(min, max) {
  return Math.round(min + Math.random() * (max - min))
}

function decimalBetween(min, max) {
  return Number((min + Math.random() * (max - min)).toFixed(2))
}

function makePlaygroundSurface() {
  const topic = pick(topics)
  const score = numberBetween(48, 92)
  const target = 85
  const confidence = decimalBetween(0.62, 0.93)
  const tone = pick(tones)
  const textVariant = pick(variants)

  return {
    surfaceId: `playground-${seed.value}`,
    rootComponent: 'root',
    dataModel: {
      theme: 'light',
      fontFamily: 'sans-serif',
      fontScale: 1,
      colorPalette: 'default',
      highContrast: false,
      reducedMotion: false,
      language: 'es'
    },
    components: {
      root: {
        id: 'root',
        type: 'Column',
        children: ['intro-card', 'layout-row', 'media-card', 'interactive-row', 'cta-row'],
        props: { gap: 18, padding: 0, alignment: 'stretch' }
      },
      'intro-card': {
        id: 'intro-card',
        type: 'Card',
        children: ['eyebrow', 'title', 'summary', 'progress'],
        props: { elevation: 1, tone, padding: 18 }
      },
      eyebrow: {
        id: 'eyebrow',
        type: 'Text',
        props: { content: 'Playground A2UI', variant: 'label' }
      },
      title: {
        id: 'title',
        type: 'Text',
        props: { content: `Superficie generada para ${topic}`, variant: textVariant }
      },
      summary: {
        id: 'summary',
        type: 'RichText',
        props: {
          markdown: `**Tema:** ${topic}\n\nEsta tarjeta prueba Text, RichText, Card, Column y ProgressBar con datos variables.`,
          accessible: true
        }
      },
      progress: {
        id: 'progress',
        type: 'ProgressBar',
        props: { value: confidence, max: 1, label: 'Confianza del agente' }
      },
      'layout-row': {
        id: 'layout-row',
        type: 'Row',
        children: ['score-card', 'target-card', 'delta-card'],
        props: { gap: 12, alignment: 'space-between', wrap: true }
      },
      'score-card': {
        id: 'score-card',
        type: 'Card',
        children: ['score-label', 'score-value'],
        props: { elevation: 0, tone: 'base', padding: 14 }
      },
      'score-label': {
        id: 'score-label',
        type: 'Text',
        props: { content: 'Resultado actual', variant: 'caption' }
      },
      'score-value': {
        id: 'score-value',
        type: 'Text',
        props: { content: `${score}%`, variant: 'h2' }
      },
      'target-card': {
        id: 'target-card',
        type: 'Card',
        children: ['target-label', 'target-value'],
        props: { elevation: 0, tone: 'primary', padding: 14 }
      },
      'target-label': {
        id: 'target-label',
        type: 'Text',
        props: { content: 'Objetivo 85%', variant: 'caption' }
      },
      'target-value': {
        id: 'target-value',
        type: 'Text',
        props: { content: `${target}%`, variant: 'h2' }
      },
      'delta-card': {
        id: 'delta-card',
        type: 'Card',
        children: ['delta-label', 'delta-value'],
        props: { elevation: 0, tone: score >= target ? 'secondary' : 'warning', padding: 14 }
      },
      'delta-label': {
        id: 'delta-label',
        type: 'Text',
        props: { content: 'Brecha', variant: 'caption' }
      },
      'delta-value': {
        id: 'delta-value',
        type: 'Text',
        props: { content: `${Math.max(target - score, 0)} pts`, variant: 'h2' }
      },
      'media-card': {
        id: 'media-card',
        type: 'Card',
        children: ['media-title', 'image', 'audio', 'video'],
        props: { elevation: 1, tone: 'base', padding: 18 }
      },
      'media-title': {
        id: 'media-title',
        type: 'Text',
        props: { content: 'Media primitives', variant: 'h3' }
      },
      image: {
        id: 'image',
        type: 'Image',
        props: {
          url: `https://picsum.photos/seed/a2ui-${seed.value}/960/540`,
          altText: `Imagen generada para ${topic}`
        }
      },
      audio: {
        id: 'audio',
        type: 'AudioPlayer',
        props: {
          url: 'https://interactive-examples.mdn.mozilla.net/media/cc0-audio/t-rex-roar.mp3',
          autoPlay: false
        }
      },
      video: {
        id: 'video',
        type: 'VideoPlayer',
        props: {
          url: 'https://interactive-examples.mdn.mozilla.net/media/cc0-videos/flower.mp4',
          poster: `https://picsum.photos/seed/poster-${seed.value}/960/540`,
          captionsUrl: '',
          autoPlay: false
        }
      },
      'interactive-row': {
        id: 'interactive-row',
        type: 'Row',
        children: ['quiz', 'socratic'],
        props: { gap: 12, alignment: 'start', wrap: true }
      },
      quiz: {
        id: 'quiz',
        type: 'QuizCard',
        props: {
          question: `Que conviene reforzar primero en ${topic}?`,
          options: ['Definicion base', 'Interpretacion en contexto', 'Calculo mecanico', 'Resumen final'],
          mode: Math.random() > 0.5 ? 'single_choice' : 'multi_choice'
        },
        events: { onSubmit: '/api/playground/quiz' }
      },
      socratic: {
        id: 'socratic',
        type: 'SocraticDialog',
        props: {
          prompt: 'Explica con tus palabras por que este punto podria mejorar tu siguiente evaluacion.',
          context: topic,
          placeholder: 'Escribe una explicacion breve...'
        },
        events: { onSubmit: '/api/playground/socratic' }
      },
      'cta-row': {
        id: 'cta-row',
        type: 'Row',
        children: ['primary-action', 'secondary-action'],
        props: { gap: 10, alignment: 'end', wrap: true }
      },
      'primary-action': {
        id: 'primary-action',
        type: 'Button',
        props: { label: 'Generar practica', variant: 'primary', disabled: false },
        events: { onClick: '/api/playground/practice' }
      },
      'secondary-action': {
        id: 'secondary-action',
        type: 'Button',
        props: { label: 'Guardar superficie', variant: 'outline', disabled: false },
        events: { onClick: '/api/playground/save' }
      }
    }
  }
}

const surface = ref(makePlaygroundSurface())

function regenerate() {
  seed.value = Date.now()
  surface.value = makePlaygroundSurface()
  eventLog.value = []
}

function handleA2UIEvent(event) {
  eventLog.value = [
    {
      ...event,
      at: new Date().toLocaleTimeString()
    },
    ...eventLog.value
  ].slice(0, 8)
}
</script>

<template>
  <div data-theme="light" class="min-h-screen bg-base-200 text-base-content">
    <header class="sticky top-0 z-20 border-b border-base-300 bg-base-100/95 backdrop-blur">
      <div class="mx-auto flex max-w-6xl flex-col gap-3 px-4 py-4 sm:flex-row sm:items-center sm:justify-between md:px-6">
        <div>
          <p class="text-sm font-semibold uppercase text-primary">A2UI Playground</p>
          <h1 class="text-2xl font-bold text-neutral">Primitives con props validos</h1>
        </div>
        <div class="flex flex-wrap gap-2">
          <NuxtLink class="btn btn-ghost btn-sm" to="/dashboard">Dashboard</NuxtLink>
          <button class="btn btn-primary btn-sm" type="button" @click="regenerate">Randomizar datos</button>
        </div>
      </div>
    </header>

    <main class="mx-auto grid max-w-6xl gap-4 px-4 py-6 md:px-6 lg:grid-cols-[1fr_20rem]">
      <section class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
        <SurfaceRenderer :surface="surface" @a2ui-event="handleA2UIEvent" />
      </section>

      <aside class="grid content-start gap-4">
        <section class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
          <p class="text-sm font-semibold uppercase text-primary">Surface</p>
          <p class="mt-2 break-all text-sm text-base-content/70">{{ surface.surfaceId }}</p>
          <div class="mt-4 grid grid-cols-2 gap-2 text-center">
            <div class="rounded-lg bg-base-200 p-3">
              <p class="text-xs text-base-content/60">Root</p>
              <p class="font-semibold text-neutral">{{ surface.rootComponent }}</p>
            </div>
            <div class="rounded-lg bg-base-200 p-3">
              <p class="text-xs text-base-content/60">Nodes</p>
              <p class="font-semibold text-neutral">{{ Object.keys(surface.components).length }}</p>
            </div>
          </div>
        </section>

        <section class="rounded-lg border border-base-300 bg-base-100 p-4 shadow-sm">
          <p class="text-sm font-semibold uppercase text-primary">Eventos</p>
          <div v-if="eventLog.length" class="mt-3 grid gap-2">
            <article v-for="entry in eventLog" :key="`${entry.componentId}-${entry.at}`" class="rounded-lg bg-base-200 p-3 text-xs">
              <div class="flex items-center justify-between gap-2">
                <span class="font-semibold text-neutral">{{ entry.componentType }}</span>
                <span class="text-base-content/60">{{ entry.at }}</span>
              </div>
              <p class="mt-1 text-base-content/70">{{ entry.eventName }} -> {{ entry.endpoint ?? 'sin endpoint' }}</p>
            </article>
          </div>
          <p v-else class="mt-3 text-sm text-base-content/60">Interactua con Button, QuizCard o SocraticDialog.</p>
        </section>
      </aside>
    </main>
  </div>
</template>
