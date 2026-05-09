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
  plugins: [
    daisyui({
      themes: [
        'light',
        'dark --prefersdark',
        'cupcake',
        'corporate  --default',
        'emerald',
        'winter',
        'night'
      ]
    })
  ]
}