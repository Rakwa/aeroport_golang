declare interface Airport {
  label?: string
  acronym: string
  data?: AirportData
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
