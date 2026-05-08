


export function cleanInput(input: string): string[] {
  const trimmed = input.trim()
  return trimmed ? trimmed.split(/\s+/) : []
}
