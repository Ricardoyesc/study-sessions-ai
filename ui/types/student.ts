import type { A2UISurface } from './a2ui'

export type EvaluationStatus = 'completed' | 'in_progress' | 'pending'
export type MasteryLevel = 'strong' | 'steady' | 'needs_attention'

export interface StudentProfile {
  id: string
  name: string
  email: string
  cluster: string
  estimatedTheta: number
  weeklyGoalMinutes: number
  studiedMinutes: number
  streakDays: number
}

export interface EvaluationInsight {
  focusPoint: string
  reason: string
  nextAction: string
  confidence: number
  estimatedMinutes: number
  generatedSurface: A2UISurface
}

export interface Evaluation {
  id: string
  title: string
  description: string
  score: number
  targetScore: number
  status: EvaluationStatus
  mastery: MasteryLevel
  lastAttempt: string
  insight: EvaluationInsight
}

export interface Subject {
  id: string
  name: string
  teacher: string
  progress: number
  accent: string
  evaluations: Evaluation[]
}
