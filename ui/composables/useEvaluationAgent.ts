import { subjects } from '~/data/student-fixtures'
import type { Evaluation, Subject } from '~/types/student'

export function useEvaluationAgent() {
  const { mode } = useAuth()
  const activeSubjectId = useState<string>('active-subject-id', () => subjects[0]?.id ?? '')
  const activeEvaluationId = useState<string>('active-evaluation-id', () => subjects[0]?.evaluations[0]?.id ?? '')
  const agentStatus = ref<'ready' | 'connecting' | 'remote' | 'fallback'>('ready')
  const a2ui = useA2UI()

  const subjectList = computed<Subject[]>(() => subjects)

  const activeSubject = computed(() => {
    return subjectList.value.find((subject) => subject.id === activeSubjectId.value) ?? subjectList.value[0]
  })

  const activeEvaluation = computed(() => {
    return activeSubject.value?.evaluations.find((evaluation) => evaluation.id === activeEvaluationId.value) ?? activeSubject.value?.evaluations[0]
  })

  const activeSurface = computed(() => {
    return a2ui.surface.value ?? activeEvaluation.value?.insight.generatedSurface ?? null
  })

  function selectEvaluation(subjectId: string, evaluationId: string) {
    activeSubjectId.value = subjectId
    activeEvaluationId.value = evaluationId
  }

  if (import.meta.client && mode.value === 'demo' && activeEvaluation.value) {
    agentStatus.value = 'fallback'
  }

  return {
    subjects: subjectList,
    activeSubject,
    activeEvaluation,
    activeSurface,
    agentStatus,
    websocketStatus: a2ui.status,
    selectEvaluation
  }
}
