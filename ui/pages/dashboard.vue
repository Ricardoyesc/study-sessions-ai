<script setup>
import AppShell from '~/components/layout/AppShell.vue'
import EvaluationDetail from '~/components/dashboard/EvaluationDetail.vue'
import StudentSummary from '~/components/dashboard/StudentSummary.vue'
import SubjectSidebar from '~/components/dashboard/SubjectSidebar.vue'

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

onMounted(() => {
  if (!isAuthenticated.value) {
    navigateTo('/')
  }
})

async function handleLogout() {
  logout()
  await navigateTo('/')
}
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

    <StudentSummary :student="student" />

    <EvaluationDetail
      v-if="activeSubject && activeEvaluation"
      :subject="activeSubject"
      :evaluation="activeEvaluation"
      :surface="activeSurface"
      :agent-status="agentStatus"
      :websocket-status="websocketStatus"
    />
  </AppShell>
</template>
