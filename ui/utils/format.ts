export function formatPercent(value: number) {
  return `${Math.round(value)}%`
}

export function formatTheta(value: number) {
  const sign = value > 0 ? '+' : ''
  return `${sign}${value.toFixed(2)}`
}
