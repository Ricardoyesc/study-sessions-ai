import { fallbackStudent } from '~/data/student-fixtures'
import type { StudentProfile } from '~/types/student'

interface LoginResponse {
  token: string
  user?: {
    id?: string
    email?: string
    name?: string
    cluster?: string
    estimated_theta?: number
  }
}

interface MeResponse {
  id?: string
  email?: string
  name?: string
  cluster?: string
  estimated_theta?: number
}

function buildProfileFromEmail(email: string, partial?: MeResponse | LoginResponse['user']): StudentProfile {
  const localPart = email.split('@')[0] ?? ''
  const emailName = localPart
    .split(/[._-]/)
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(' ')

  return {
    ...fallbackStudent,
    id: partial?.id ?? fallbackStudent.id,
    name: partial?.name ?? emailName ?? fallbackStudent.name,
    email: partial?.email ?? email,
    cluster: partial?.cluster ?? fallbackStudent.cluster,
    estimatedTheta: partial?.estimated_theta ?? fallbackStudent.estimatedTheta
  }
}

export function useAuth() {
  const config = useRuntimeConfig()
  const token = useState<string | null>('auth-token', () => null)
  const student = useState<StudentProfile>('auth-student', () => fallbackStudent)
  const loading = useState('auth-loading', () => false)
  const errorMessage = useState<string | null>('auth-error', () => null)
  const mode = useState<'api' | 'demo'>('auth-mode', () => 'demo')

  const isAuthenticated = computed(() => Boolean(token.value))

  function hydrate() {
    if (!import.meta.client) {
      return
    }

    const storedToken = localStorage.getItem('sai-token')
    const storedStudent = localStorage.getItem('sai-student')

    if (storedToken) {
      token.value = storedToken
    }

    if (storedStudent) {
      try {
        student.value = JSON.parse(storedStudent) as StudentProfile
      } catch {
        student.value = fallbackStudent
      }
    }
  }

  async function fetchProfile(authToken: string, email: string) {
    try {
      const profile = await $fetch<MeResponse>(`${config.public.apiBase}/api/users/me`, {
        timeout: 5000,
        headers: {
          Authorization: `Bearer ${authToken}`
        }
      })

      student.value = buildProfileFromEmail(profile.email ?? email, profile)
    } catch (err) {
      console.warn('Profile fetch failed, using email-based profile:', err)
      student.value = buildProfileFromEmail(email)
    }

    if (import.meta.client) {
      localStorage.setItem('sai-student', JSON.stringify(student.value))
    }
  }

  async function login(email: string, password: string) {
    loading.value = true
    errorMessage.value = null

    try {
      const response = await $fetch<LoginResponse>(`${config.public.apiBase}/api/users/login`, {
        method: 'POST',
        timeout: 5000,
        body: { email, password }
      })

      token.value = response.token
      mode.value = 'api'
      student.value = buildProfileFromEmail(email, response.user)

      if (import.meta.client) {
        localStorage.setItem('sai-token', response.token)
        localStorage.setItem('sai-student', JSON.stringify(student.value))
      }

      await fetchProfile(response.token, email)
      return { ok: true, mode: mode.value }
    } catch (err) {
      console.warn('Login API failed, falling back to demo mode:', err)
      token.value = 'demo-token'
      mode.value = 'demo'
      student.value = buildProfileFromEmail(email)
      errorMessage.value = 'Backend no disponible. Entraste en modo demo con datos locales.'

      if (import.meta.client) {
        localStorage.setItem('sai-token', token.value)
        localStorage.setItem('sai-student', JSON.stringify(student.value))
      }

      return { ok: true, mode: mode.value }
    } finally {
      loading.value = false
    }
  }

  function logout() {
    token.value = null
    student.value = fallbackStudent
    mode.value = 'demo'
    errorMessage.value = null

    if (import.meta.client) {
      localStorage.removeItem('sai-token')
      localStorage.removeItem('sai-student')
    }
  }

  hydrate()

  return {
    token,
    student,
    loading,
    errorMessage,
    mode,
    isAuthenticated,
    login,
    logout,
    hydrate
  }
}
