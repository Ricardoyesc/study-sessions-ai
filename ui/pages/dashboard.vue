<script setup>
import AppShell from '~/components/layout/AppShell.vue'
import EvaluationDetail from '~/components/dashboard/EvaluationDetail.vue'
import StudentSummary from '~/components/dashboard/StudentSummary.vue'
import SubjectSidebar from '~/components/dashboard/SubjectSidebar.vue'
import StudyPanel from '~/components/dashboard/StudyPanel.vue'

const { student, mode, isAuthenticated, logout } = useAuth()
const {
  subjects,
  activeSubject,
  activeEvaluation,
  activeSurface,
  agentStatus,
  websocketStatus,
  selectEvaluation
} = useEvaluationAgent()

const study = useStudySession()

onMounted(() => {
  if (!isAuthenticated.value) {
    navigateTo('/')
  }
})

async function handleLogout() {
  study.reset()
  logout()
  await navigateTo('/')
}

async function handleStartSession(topic: string) {
  await study.startSession(topic)
  await study.nextItem()
}

async function handleNext() {
  await study.nextItem()
}

async function handleAnswer(index: number) {
  await study.submitAnswer(index)
}

async function handleSocraticResponse(response: string) {
  await study.submitSocraticResponse(response)
}

function handleCloseSession() {
  study.reset()
}

const showStudyPanel = computed(() => study.sessionState.value !== 'idle')
</script>

<template>
  <AppShell :student="student" :auth-mode="mode" @logout="handleLogout">
    <template #sidebar>
      <SubjectSidebar
        v-if="activeSubject && activeEvaluation"
        :subjects="subjects"
        :active-subject-id="activeSubject.id"
        :active-evaluation-id="activeEvaluation.id"
        @select="selectEvaluation"
      />
    </template>

    <div class="flex flex-col gap-4">
      <StudentSummary :student="student" />

      <EvaluationDetail
        v-if="activeSubject && activeEvaluation && !showStudyPanel"
        :subject="activeSubject"
        :evaluation="activeEvaluation"
        :surface="activeSurface"
        :agent-status="agentStatus"
        :websocket-status="websocketStatus"
      />

      <StudyPanel
        v-if="showStudyPanel"
        :surface="study.currentSurface.value"
        :state="study.sessionState.value"
        :topic="study.sessionTopic.value"
        :feedback="study.feedback.value"
        :is-correct="study.isCorrect.value"
        :loading="study.isLoading.value"
        :is-active="showStudyPanel"
        @start="handleStartSession"
        @next="handleNext"
        @answer="handleAnswer"
        @socratic-response="handleSocraticResponse"
        @close="handleCloseSession"
      />

      <StudyPanel
        v-if="!showStudyPanel"
        :surface="null"
        :state="'idle'"
        :topic="''"
        :feedback="null"
        :is-correct="null"
        :loading="false"
        :is-active="false"
        @start="handleStartSession"
        @next="handleNext"
        @answer="handleAnswer"
        @socratic-response="handleSocraticResponse"
        @close="handleCloseSession"
      />
    </div>
  </AppShell>
</template>
