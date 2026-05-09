import daisyui from 'daisyui'

export default {
  content: [
    './app.vue',
    './components/**/*.{vue,js,ts}',
    './layouts/**/*.vue',
    './pages/**/*.vue',
    './plugins/**/*.{js,ts}',
    './nuxt.config.{js,ts}'
  ],
  theme: {
    extend: {}
  },
  plugins: [daisyui],
  daisyui: {
    themes: [
      {
        study: {
          primary: '#2563eb',
          secondary: '#16a34a',
          accent: '#f59e0b',
          neutral: '#172554',
          'base-100': '#f8fafc',
          'base-200': '#e2e8f0',
          'base-300': '#cbd5e1',
          info: '#0891b2',
          success: '#16a34a',
          warning: '#d97706',
          error: '#dc2626'
        }
      }
    ]
  }
}