<script setup>
import { resolveA2UIComponent } from './registry'

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

const emit = defineEmits(['a2ui-event'])

const component = computed(() => {
  return props.surface.components[props.componentId ?? props.surface.rootComponent] ?? null
})

const componentView = computed(() => {
  return component.value ? resolveA2UIComponent(component.value.type) : null
})

const componentProps = computed(() => component.value?.props ?? {})
const children = computed(() => component.value?.children ?? [])

const componentListeners = computed(() => {
  if (!component.value) {
    return {}
  }

  if (component.value.type === 'Button') {
    return { click: (payload) => forwardEvent('click', payload) }
  }

  if (component.value.type === 'QuizCard' || component.value.type === 'SocraticDialog') {
    return { submit: (payload) => forwardEvent('submit', payload) }
  }

  return {}
})

function eventEndpoint(eventName) {
  if (!component.value?.events) {
    return undefined
  }

  const normalizedName = eventName.charAt(0).toUpperCase() + eventName.slice(1)
  return component.value.events[eventName] ?? component.value.events[`on${normalizedName}`]
}

function forwardEvent(eventName, payload = {}) {
  if (!component.value) {
    return
  }

  emit('a2ui-event', {
    componentId: component.value.id,
    componentType: component.value.type,
    eventName,
    endpoint: eventEndpoint(eventName),
    payload
  })
}
</script>

<template>
  <component
    :is="componentView"
    v-if="component && componentView"
    v-bind="componentProps"
    v-on="componentListeners"
  >
    <SurfaceRenderer
      v-for="childId in children"
      :key="childId"
      :surface="surface"
      :component-id="childId"
      @a2ui-event="emit('a2ui-event', $event)"
    />
  </component>

  <div v-else-if="component" class="rounded-lg border border-warning/30 bg-warning/10 p-3 text-sm text-warning-content">
    Componente A2UI no soportado: {{ component.type }}
  </div>
</template>
