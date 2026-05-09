<script setup>
import { useTheme } from '~/composables/useTheme'

const { availableThemes, currentTheme, setTheme } = useTheme()

const labels = {
  light: 'Light',
  dark: 'Dark',
  cupcake: 'Cupcake',
  corporate: 'Corporate',
  emerald: 'Emerald',
  winter: 'Winter',
  night: 'Night'
}

function handleChange(event) {
  setTheme(event.target.value)
}
</script>

<template>
  <label class="flex items-center gap-2 rounded-lg border border-base-300 bg-base-200/70 px-2 py-1 text-sm shadow-sm">
    <span class="hidden text-xs font-semibold text-base-content/60 sm:inline">Tema</span>
    <select class="select select-bordered select-sm min-h-8 w-32 border-base-content/20 bg-base-100" :value="currentTheme" @change="handleChange">
      <option v-for="theme in availableThemes" :key="theme" :value="theme">
        {{ labels[theme] ?? theme }}
      </option>
    </select>
  </label>
</template>