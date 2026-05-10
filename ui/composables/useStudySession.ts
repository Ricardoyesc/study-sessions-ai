import type { A2UISurface } from '~/types/a2ui'

interface SessionResponse {
  id: string
  state: string
  topic: string
  ws_url: string
}

interface NextItemResponse {
  state: string
  topic: string
  message: string
  surface: A2UISurface | null
}

interface QuizAnswerResponse {
  state: string
  is_correct: boolean
  correct_answer: string
  feedback: string
  message: string
}

function authHeaders(token: string | null) {
  return {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
}

export function useStudySession() {
  const config = useRuntimeConfig()
  const { token } = useAuth()

  const sessionId = useState<string | null>('session-id', () => null)
  const sessionState = useState<string>('session-state', () => 'idle')
  const sessionTopic = useState<string>('session-topic', () => '')
  const currentSurface = useState<A2UISurface | null>('study-surface', () => null)
  const feedback = useState<string | null>('study-feedback', () => null)
  const isCorrect = useState<boolean | null>('study-last-correct', () => null)
  const isLoading = useState('study-loading', () => false)

  async function startSession(topic: string) {
    isLoading.value = true
    try {
      const resp = await $fetch<SessionResponse>(`${config.public.apiBase}/api/sessions`, {
        method: 'POST',
        timeout: 10000,
        headers: authHeaders(token.value),
        body: { topic }
      })
      sessionId.value = resp.id
      sessionState.value = resp.state
      sessionTopic.value = topic
      return resp
    } catch (e) {
      console.error('Failed to start session:', e)
      return null
    } finally {
      isLoading.value = false
    }
  }

  async function nextItem() {
    if (!sessionId.value) return
    isLoading.value = true
    try {
      const resp = await $fetch<NextItemResponse>(
        `${config.public.apiBase}/api/sessions/${sessionId.value}/next?topic=${encodeURIComponent(sessionTopic.value)}`,
        { timeout: 60000, headers: authHeaders(token.value) }
      )
      sessionState.value = resp.state
      currentSurface.value = resp.surface
      feedback.value = null
      isCorrect.value = null
      return resp
    } catch (e) {
      console.error('Failed to get next item:', e)
      return null
    } finally {
      isLoading.value = false
    }
  }

  async function submitAnswer(selectedIndex: number) {
    if (!sessionId.value) return
    isLoading.value = true
    try {
      const resp = await $fetch<QuizAnswerResponse>(
        `${config.public.apiBase}/api/sessions/${sessionId.value}/quiz/answer`,
        {
          method: 'POST',
          timeout: 10000,
          headers: authHeaders(token.value),
          body: { selected_index: selectedIndex }
        }
      )
      isCorrect.value = resp.is_correct
      feedback.value = resp.feedback
      sessionState.value = resp.state
      return resp
    } catch (e) {
      console.error('Failed to submit answer:', e)
      return null
    } finally {
      isLoading.value = false
    }
  }

  async function submitSocraticResponse(studentResponse: string) {
    if (!sessionId.value) return
    isLoading.value = true
    try {
      const resp = await $fetch<{ state: string; message: string }>(
        `${config.public.apiBase}/api/sessions/${sessionId.value}/socratic/response`,
        {
          method: 'POST',
          timeout: 10000,
          headers: authHeaders(token.value),
          body: { student_response: studentResponse }
        }
      )
      sessionState.value = resp.state
      feedback.value = null
      return resp
    } catch (e) {
      console.error('Failed to submit socratic response:', e)
      return null
    } finally {
      isLoading.value = false
    }
  }

  function reset() {
    sessionId.value = null
    sessionState.value = 'idle'
    sessionTopic.value = ''
    currentSurface.value = null
    feedback.value = null
    isCorrect.value = null
  }

  return {
    sessionId,
    sessionState,
    sessionTopic,
    currentSurface,
    feedback,
    isCorrect,
    isLoading,
    startSession,
    nextItem,
    submitAnswer,
    submitSocraticResponse,
    reset
  }
}
