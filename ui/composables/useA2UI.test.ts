import { describe, test, expect, beforeEach } from 'vitest'
import type { A2UISurface, A2UIDataModel } from '~/types/a2ui'
import { useA2UI } from '~/composables/useA2UI'

const baseDataModel: A2UIDataModel = {
  theme: 'system',
  fontFamily: 'sans-serif',
  fontScale: 1.0,
  colorPalette: 'default',
  highContrast: false,
  reducedMotion: false,
  language: 'es',
}

function makeSurface(overrides: Partial<A2UISurface> = {}): A2UISurface {
  return {
    surfaceId: 'surf1',
    rootComponent: 'root',
    components: {
      root: { id: 'root', type: 'Column', props: { gap: 8 }, children: ['c1'] },
      c1: { id: 'c1', type: 'Text', props: { content: 'Hello' } },
    },
    dataModel: { ...baseDataModel },
    ...overrides,
  }
}

describe('useA2UI — applyMessage', () => {
  let composable: ReturnType<typeof useA2UI>

  beforeEach(() => {
    composable = useA2UI()
  })

  test('a2ui_full replaces surface wholesale', () => {
    const surf = makeSurface()
    composable.applyMessage({ type: 'a2ui_full', payload: surf, timestamp: '' })
    expect(composable.surface.value?.surfaceId).toBe('surf1')
    expect(composable.surface.value?.components['c1'].props['content']).toBe('Hello')
  })

  test('a2ui_update merges props on existing component', () => {
    composable.applyMessage({ type: 'a2ui_full', payload: makeSurface(), timestamp: '' })
    composable.applyMessage({
      type: 'a2ui_update',
      payload: { updates: [{ componentId: 'c1', props: { content: 'World' } }] },
      timestamp: '',
    })
    expect(composable.surface.value?.components['c1'].props['content']).toBe('World')
    expect(composable.surface.value?.components['root'].props['gap']).toBe(8)
  })

  test('a2ui_update skips unknown componentId without crash', () => {
    composable.applyMessage({ type: 'a2ui_full', payload: makeSurface(), timestamp: '' })
    expect(() => {
      composable.applyMessage({
        type: 'a2ui_update',
        payload: { updates: [{ componentId: 'does-not-exist', props: { x: 1 } }] },
        timestamp: '',
      })
    }).not.toThrow()
    expect(Object.keys(composable.surface.value?.components ?? {})).toHaveLength(2)
  })

  test('a2ui_update replaces children when update.children provided', () => {
    composable.applyMessage({ type: 'a2ui_full', payload: makeSurface(), timestamp: '' })
    composable.applyMessage({
      type: 'a2ui_update',
      payload: { updates: [{ componentId: 'root', props: {}, children: ['c1', 'c2'] }] },
      timestamp: '',
    })
    expect(composable.surface.value?.components['root'].children).toEqual(['c1', 'c2'])
  })

  test('a2ui_update preserves children when update.children omitted', () => {
    composable.applyMessage({ type: 'a2ui_full', payload: makeSurface(), timestamp: '' })
    composable.applyMessage({
      type: 'a2ui_update',
      payload: { updates: [{ componentId: 'root', props: { gap: 16 } }] },
      timestamp: '',
    })
    expect(composable.surface.value?.components['root'].children).toEqual(['c1'])
  })

  test('data_model_update shallow-merges into existing dataModel', () => {
    composable.applyMessage({ type: 'a2ui_full', payload: makeSurface(), timestamp: '' })
    composable.applyMessage({
      type: 'data_model_update',
      payload: { path: '', value: null, diff: { language: 'en' } },
      timestamp: '',
    })
    expect(composable.surface.value?.dataModel?.language).toBe('en')
    expect(composable.surface.value?.dataModel?.theme).toBe('system')
  })

  test('data_model_update ignored when surface is null', () => {
    expect(() => {
      composable.applyMessage({
        type: 'data_model_update',
        payload: { path: '', value: null, diff: { language: 'en' } },
        timestamp: '',
      })
    }).not.toThrow()
    expect(composable.surface.value).toBeNull()
  })
})
