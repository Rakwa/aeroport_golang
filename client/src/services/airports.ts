const randomMeasure = {
  wind: Math.floor(Math.random() * (30 - 10 + 1) + 10),
  temp: Math.floor(Math.random() * (30 - 10 + 1) + 10),
  pressure: Math.floor(Math.random() * (30 - 10 + 1) + 10),
}
const fakedData = [
  {
    acronym: 'NTE',
    today: randomMeasure,
    measures: [
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
    ],
  },
  {
    acronym: 'RNE',
    today: randomMeasure,
    measures: [
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
    ],
  },
  {
    acronym: 'VGN',
    today: randomMeasure,
    measures: [
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
      {
        date: '22/10/2022',
        measure: randomMeasure,
      },
    ],
  },
]

export async function getAirportsTowns(): Promise<Airport[]> {
  const sleep = (ms: number) =>
    new Promise((resolve) => setTimeout(resolve, ms))
  await sleep(2000)
  return [
    { label: 'Nantes', acronym: 'NTE' },
    { label: 'Vingeux', acronym: 'VGN' },
  ]
}

export async function fetchAirportData(
  airportID: string | null
): Promise<AirportData> {
  console.log(airportID)
  const sleep = (ms: number) =>
    new Promise((resolve) => setTimeout(resolve, ms))
  await sleep(2000)
  return fakedData.filter((item) => item.acronym == airportID)[0]
}
