export type A2UIComponentType =
  | 'Text'
  | 'RichText'
  | 'Image'
  | 'AudioPlayer'
  | 'VideoPlayer'
  | 'Card'
  | 'Column'
  | 'Row'
  | 'QuizCard'
  | 'SocraticDialog'
  | 'ProgressBar'
  | 'Button'

export interface A2UIComponent {
  id: string
  type: A2UIComponentType
  children?: string[]
  props: Record<string, unknown>
  events?: Record<string, string>
}

export interface A2UIDataModel {
  theme: string
  fontFamily: string
  fontScale: number
  colorPalette: string
  highContrast: boolean
  reducedMotion: boolean
  language: string
}

export interface A2UISurface {
  surfaceId: string
  rootComponent: string
  components: Record<string, A2UIComponent>
  dataModel: A2UIDataModel
}

export type A2UIMessageType = 'a2ui_full' | 'a2ui_update' | 'data_model_update' | 'error' | 'ping' | 'pong'

export interface A2UIFullMessage {
  type: 'a2ui_full'
  payload: A2UISurface
  timestamp: string
}

export interface A2UIUpdateMessage {
  type: 'a2ui_update'
  payload: {
    updates: Array<{
      componentId: string
      props?: Record<string, unknown>
      children?: string[]
    }>
  }
  timestamp: string
}

export interface A2UIDataModelUpdateMessage {
  type: 'data_model_update'
  payload: {
    path: string
    value: unknown
    diff: Partial<A2UIDataModel>
  }
  timestamp: string
}

export interface A2UIErrorMessage {
  type: 'error'
  payload: {
    code: string
    message: string
  }
  timestamp: string
}

export type A2UIMessage = A2UIFullMessage | A2UIUpdateMessage | A2UIDataModelUpdateMessage | A2UIErrorMessage
