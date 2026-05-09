import type { A2UIMessage, A2UISurface } from '~/types/a2ui'

export function useA2UI() {
  const config = useRuntimeConfig()
  const surface = ref<A2UISurface | null>(null)
  const status = ref<'idle' | 'connecting' | 'connected' | 'fallback' | 'error'>('idle')
  const lastError = ref<string | null>(null)
  let socket: WebSocket | null = null

  function applyMessage(message: A2UIMessage) {
    if (message.type === 'a2ui_full') {
      surface.value = message.payload
      return
    }

    if (message.type === 'a2ui_update' && surface.value) {
      for (const update of message.payload.updates) {
        const component = surface.value.components[update.componentId]
        if (!component) {
          continue
        }

        surface.value.components[update.componentId] = {
          ...component,
          props: {
            ...component.props,
            ...update.props
          },
          children: update.children ?? component.children
        }
      }
      return
    }

    if (message.type === 'data_model_update' && surface.value) {
      surface.value.dataModel = {
        ...surface.value.dataModel,
        ...message.payload.diff
      }
    }
  }

  function connect(sessionId: string) {
    if (!import.meta.client) {
      return
    }

    close()
    status.value = 'connecting'
    lastError.value = null

    try {
      socket = new WebSocket(`${config.public.wsBase}/ws/session/${sessionId}`)

      socket.onopen = () => {
        status.value = 'connected'
        socket?.send(JSON.stringify({ type: 'ping', timestamp: new Date().toISOString() }))
      }

      socket.onmessage = (event) => {
        try {
          applyMessage(JSON.parse(event.data) as A2UIMessage)
        } catch {
          lastError.value = 'No se pudo interpretar el mensaje A2UI del servidor.'
        }
      }

      socket.onerror = () => {
        status.value = 'error'
        lastError.value = 'No se pudo conectar el WebSocket A2UI.'
      }

      socket.onclose = () => {
        if (status.value === 'connecting') {
          status.value = 'fallback'
        }
      }
    } catch {
      status.value = 'error'
      lastError.value = 'El navegador no pudo crear la conexion WebSocket.'
    }
  }

  function close() {
    if (socket) {
      socket.close()
      socket = null
    }
  }

  onBeforeUnmount(close)

  return {
    surface,
    status,
    lastError,
    connect,
    close,
    applyMessage
  }
}
