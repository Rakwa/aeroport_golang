import { ComputedRef, ref, watch } from 'vue'
import { fetchAirportData } from '../services/airports'
import createAsyncProcess from '../utils/createAsyncProcess'

interface useAirportDataProps {
  airportId: ComputedRef<string | null>
}

export function useAirportData() {
  const airport = ref<Airport | null>(null)

  const setAirportId = (id: string) => (airport.value = { acronym: id })

  async function getAirportData(): Promise<void> {
    if (!airport.value?.acronym) return
    const result = await fetchAirportData(airport.value?.acronym)
    airport.value.data = result
  }

  const { active: isLoading, run: fetchData } =
    createAsyncProcess(getAirportData)

  return {
    isLoading,
    airport,
    fetchData,
    setAirportId,
  }
}
