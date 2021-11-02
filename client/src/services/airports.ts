export async function getAirportsTowns(): Promise<Airport[]> {
  await sleep(3000)

  const airports = await fetch('http://localhost:3333/api/airports')
    .then((res) => {
      return res.json()
    })
    .catch((err) => console.log(err))
  return airports.map((airport: any) => ({
    acronym: airport._id,
    label: airport.city,
    name: decodeURIComponent(escape(airport.name)),
  }))
}
function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms))
}
export async function fetchAirportData(
  airportID: string | null
): Promise<AirportData> {
  const averages = await fetch(
    `http://localhost:3333/api/airports/${airportID}/averages`
  )
    .then((res) => {
      return res.json()
    })
    .catch((err) => console.log(err))
  const startDate = new Date()
  startDate.setDate(startDate.getDate() - 3)
  const endDate = new Date()
  const measures = await fetch(
    `http://localhost:3333/api/airports/${airportID}/measures?startDate=${startDate.toISOString()}&endDate=${endDate.toISOString()}`
  )
    .then((res) => {
      return res.json()
    })
    .catch((err) => console.log(err))
  const m = measures.map((e: any) => ({
    ...e,
    date: new Date(e.date * 1000),
  }))
  return {
    today: {
      wind:
        Math.floor(averages.filter((e: any) => e._id == 'wind')[0].average) ||
        0,
      temp: Math.floor(
        averages.filter((e: any) => e._id == 'temperature')[0].average || 0
      ),
      pressure:
        Math.floor(
          averages.filter((e: any) => e._id == 'pressure')[0]?.average
        ) || 0,
    },
    measures: m,
  }
}
