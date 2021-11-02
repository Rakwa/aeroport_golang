declare interface Airport {
  label?: string
  acronym: string
  data?: AirportData
  name?: string
}
declare interface AirportData {
  today: Measure
  measures: Array<{
    date: string
    measure: Measure
  }>
}
declare interface Measure {
  temp: number
  wind: number
  pressure: number
}
