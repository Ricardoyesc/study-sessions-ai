const THEME_STORAGE_KEY = 'study-sessions-theme'
const DEFAULT_THEME = 'corporate'

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
  const currentTheme = useState('current-theme', () => DEFAULT_THEME)
  const initialized = useState('theme-initialized', () => false)

  function normalizeTheme(theme: string | null) {
    return theme && availableThemes.includes(theme) ? theme : DEFAULT_THEME
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