export async function getAirportsTowns(): Promise<[Airport]> {
  const sleep = (ms: number) =>
    new Promise((resolve) => setTimeout(resolve, ms))
  await sleep(2000)
  return [{ label: 'Nantes', acronym: 'NTE' }]
}
