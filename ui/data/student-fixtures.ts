import type { A2UISurface } from '~/types/a2ui'
import type { StudentProfile, Subject } from '~/types/student'

export const fallbackStudent: StudentProfile = {
  id: 'student-demo-001',
  name: 'Sofia Martinez',
  email: 'sofia.martinez@example.com',
  cluster: 'Intermedio adaptable',
  estimatedTheta: 0.42,
  weeklyGoalMinutes: 240,
  studiedMinutes: 165,
  streakDays: 6
}

export const defaultDataModel = {
  theme: 'light',
  fontFamily: 'sans-serif',
  fontScale: 1,
  colorPalette: 'default',
  highContrast: false,
  reducedMotion: false,
  language: 'es'
}

export function buildImprovementSurface(evaluationId: string, title: string, focusPoint: string, reason: string, nextAction: string): A2UISurface {
  return {
    surfaceId: `agent-${evaluationId}`,
    rootComponent: 'root',
    dataModel: defaultDataModel,
    components: {
      root: {
        id: 'root',
        type: 'Column',
        children: ['header', 'focus-card', 'practice-card', 'reflection-card'],
        props: { gap: 14 }
      },
      header: {
        id: 'header',
        type: 'Row',
        children: ['title', 'confidence'],
        props: { alignment: 'space-between', gap: 12 }
      },
      title: {
        id: 'title',
        type: 'Text',
        props: { content: `Agente de mejora: ${title}`, variant: 'h3' }
      },
      confidence: {
        id: 'confidence',
        type: 'ProgressBar',
        props: { value: 0.82, max: 1, label: 'Confianza del diagnostico' }
      },
      'focus-card': {
        id: 'focus-card',
        type: 'Card',
        children: ['focus-title', 'focus-text'],
        props: { tone: 'primary' }
      },
      'focus-title': {
        id: 'focus-title',
        type: 'Text',
        props: { content: 'Punto a mejorar', variant: 'label' }
      },
      'focus-text': {
        id: 'focus-text',
        type: 'RichText',
        props: { markdown: `**${focusPoint}**\n\n${reason}` }
      },
      'practice-card': {
        id: 'practice-card',
        type: 'Card',
        children: ['practice-title', 'practice-text'],
        props: { tone: 'accent' }
      },
      'practice-title': {
        id: 'practice-title',
        type: 'Text',
        props: { content: 'Micro accion sugerida', variant: 'label' }
      },
      'practice-text': {
        id: 'practice-text',
        type: 'RichText',
        props: { markdown: nextAction }
      },
      'reflection-card': {
        id: 'reflection-card',
        type: 'SocraticDialog',
        props: {
          prompt: 'Explica en dos frases por que este error aparece y que haras distinto en el siguiente intento.',
          context: evaluationId,
          placeholder: 'Escribe tu explicacion aqui...'
        }
      }
    }
  }
}

export const subjects: Subject[] = [
  {
    id: 'math',
    name: 'Matematica aplicada',
    teacher: 'Dra. Luciana Rivas',
    progress: 68,
    accent: 'primary',
    evaluations: [
      {
        id: 'math-derivatives',
        title: 'Derivadas e interpretacion grafica',
        description: 'Lectura de pendiente, tasa de cambio y aproximacion local.',
        score: 76,
        targetScore: 85,
        status: 'completed',
        mastery: 'needs_attention',
        lastAttempt: 'Hoy, 10:20',
        insight: {
          focusPoint: 'Relacionar la pendiente con el cambio instantaneo',
          reason: 'Respondiste bien los calculos mecanicos, pero fallaste cuando la pregunta pidio interpretar el resultado en contexto.',
          nextAction: 'Resuelve 3 graficas cortas: marca dos puntos cercanos, estima la pendiente y escribe una frase que explique que significa ese numero.',
          confidence: 0.82,
          estimatedMinutes: 14,
          generatedSurface: buildImprovementSurface(
            'math-derivatives',
            'Derivadas e interpretacion grafica',
            'Relacionar la pendiente con el cambio instantaneo',
            'Respondiste bien los calculos mecanicos, pero fallaste cuando la pregunta pidio interpretar el resultado en contexto.',
            'Resuelve 3 graficas cortas: marca dos puntos cercanos, estima la pendiente y escribe una frase que explique que significa ese numero.'
          )
        }
      },
      {
        id: 'math-integrals',
        title: 'Integrales como acumulacion',
        description: 'Areas bajo la curva y acumulacion de cantidades variables.',
        score: 88,
        targetScore: 85,
        status: 'completed',
        mastery: 'strong',
        lastAttempt: 'Ayer, 18:40',
        insight: {
          focusPoint: 'Mantener precision al elegir limites',
          reason: 'Tu dominio general supera el objetivo. El unico patron debil aparece cuando los limites vienen descritos con texto.',
          nextAction: 'Antes de integrar, subraya la cantidad inicial y final del intervalo. Luego escribe los limites antes de tocar la formula.',
          confidence: 0.76,
          estimatedMinutes: 8,
          generatedSurface: buildImprovementSurface(
            'math-integrals',
            'Integrales como acumulacion',
            'Mantener precision al elegir limites',
            'Tu dominio general supera el objetivo. El unico patron debil aparece cuando los limites vienen descritos con texto.',
            'Antes de integrar, subraya la cantidad inicial y final del intervalo. Luego escribe los limites antes de tocar la formula.'
          )
        }
      }
    ]
  },
  {
    id: 'physics',
    name: 'Fisica moderna',
    teacher: 'Prof. Mateo Alvarez',
    progress: 54,
    accent: 'secondary',
    evaluations: [
      {
        id: 'physics-waves',
        title: 'Dualidad onda-particula',
        description: 'Interferencia, medicion y colapso del estado.',
        score: 63,
        targetScore: 85,
        status: 'completed',
        mastery: 'needs_attention',
        lastAttempt: 'Hoy, 09:05',
        insight: {
          focusPoint: 'Explicar por que medir cambia el patron observado',
          reason: 'Tus respuestas reconocen el fenomeno, pero no conectan medicion, informacion disponible y desaparicion de interferencia.',
          nextAction: 'Dibuja el experimento de doble rendija en dos versiones: sin detector y con detector. Debajo de cada dibujo escribe que informacion existe y que patron aparece.',
          confidence: 0.86,
          estimatedMinutes: 18,
          generatedSurface: buildImprovementSurface(
            'physics-waves',
            'Dualidad onda-particula',
            'Explicar por que medir cambia el patron observado',
            'Tus respuestas reconocen el fenomeno, pero no conectan medicion, informacion disponible y desaparicion de interferencia.',
            'Dibuja el experimento de doble rendija en dos versiones: sin detector y con detector. Debajo de cada dibujo escribe que informacion existe y que patron aparece.'
          )
        }
      },
      {
        id: 'physics-relativity',
        title: 'Relatividad especial',
        description: 'Tiempo propio, dilatacion y simultaneidad.',
        score: 71,
        targetScore: 85,
        status: 'in_progress',
        mastery: 'steady',
        lastAttempt: 'En curso',
        insight: {
          focusPoint: 'Separar tiempo propio de tiempo medido por otro observador',
          reason: 'El calculo numerico esta avanzando, pero mezclas el marco de referencia al justificar la respuesta.',
          nextAction: 'En cada ejercicio, escribe primero quien viaja con el reloj y quien observa desde fuera. Despues elige la formula.',
          confidence: 0.79,
          estimatedMinutes: 12,
          generatedSurface: buildImprovementSurface(
            'physics-relativity',
            'Relatividad especial',
            'Separar tiempo propio de tiempo medido por otro observador',
            'El calculo numerico esta avanzando, pero mezclas el marco de referencia al justificar la respuesta.',
            'En cada ejercicio, escribe primero quien viaja con el reloj y quien observa desde fuera. Despues elige la formula.'
          )
        }
      }
    ]
  },
  {
    id: 'history',
    name: 'Historia contemporanea',
    teacher: 'Lic. Valeria Sosa',
    progress: 81,
    accent: 'accent',
    evaluations: [
      {
        id: 'history-coldwar',
        title: 'Guerra fria y bloques geopoliticos',
        description: 'Causas, tensiones indirectas y consecuencias regionales.',
        score: 84,
        targetScore: 85,
        status: 'completed',
        mastery: 'steady',
        lastAttempt: 'Viernes, 16:10',
        insight: {
          focusPoint: 'Conectar causas economicas con decisiones politicas',
          reason: 'Identificas hechos clave, pero las respuestas pierden fuerza cuando deben explicar relaciones causa-consecuencia.',
          nextAction: 'Crea una cadena de 4 eslabones: presion economica, decision politica, reaccion internacional y efecto social.',
          confidence: 0.73,
          estimatedMinutes: 10,
          generatedSurface: buildImprovementSurface(
            'history-coldwar',
            'Guerra fria y bloques geopoliticos',
            'Conectar causas economicas con decisiones politicas',
            'Identificas hechos clave, pero las respuestas pierden fuerza cuando deben explicar relaciones causa-consecuencia.',
            'Crea una cadena de 4 eslabones: presion economica, decision politica, reaccion internacional y efecto social.'
          )
        }
      }
    ]
  }
]
