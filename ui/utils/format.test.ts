import { describe, test, expect } from 'vitest'
import { formatPercent, formatTheta } from '~/utils/format'

describe('formatPercent', () => {
  test('rounds and appends %', () => {
    expect(formatPercent(85.6)).toBe('86%')
  })
  test('handles 100', () => {
    expect(formatPercent(100)).toBe('100%')
  })
  test('handles 0', () => {
    expect(formatPercent(0)).toBe('0%')
  })
})

describe('formatTheta', () => {
  test('positive value gets + prefix', () => {
    expect(formatTheta(1.23)).toBe('+1.23')
  })
  test('negative value has no extra prefix', () => {
    expect(formatTheta(-0.5)).toBe('-0.50')
  })
  test('zero has no prefix', () => {
    expect(formatTheta(0)).toBe('0.00')
  })
})
