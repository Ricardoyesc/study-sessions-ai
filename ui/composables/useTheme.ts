const THEME_STORAGE_KEY = 'study-sessions-theme'

export const availableThemes = [
  'light',
  'dark',
  'cupcake',
  'corporate',
  'emerald',
  'winter',
  'night'
]

export function useTheme() {
  const currentTheme = useState('current-theme', () => 'light')
  const initialized = useState('theme-initialized', () => false)

  function normalizeTheme(theme: string | null) {
    return theme && availableThemes.includes(theme) ? theme : 'light'
  }

  function applyTheme(theme: string) {
    const nextTheme = normalizeTheme(theme)
    currentTheme.value = nextTheme

    if (import.meta.client) {
      document.documentElement.setAttribute('data-theme', nextTheme)
      localStorage.setItem(THEME_STORAGE_KEY, nextTheme)
    }
  }

  onMounted(() => {
    if (initialized.value) {
      applyTheme(currentTheme.value)
      return
    }

    initialized.value = true
    applyTheme(localStorage.getItem(THEME_STORAGE_KEY))
  })

  return {
    availableThemes,
    currentTheme,
    setTheme: applyTheme
  }
}