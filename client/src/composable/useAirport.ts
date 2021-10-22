import { ref, watch } from 'vue'
import { getAirportsTowns } from '../services/airports'
import createAsyncProcess from '../utils/createAsyncProcess'

export function useAirport() {
  const airports = ref<[Airport] | null>(null)
  const airportSelected = ref<Airport | null>(null)

  async function fetchProfile(): Promise<void> {
    const result = await getAirportsTowns()
    airports.value = result
  }

  watch(
    airportSelected,
    () => console.log('airport changed, data loading ...'),
    { immediate: true }
  )
  const { active: isLoading, run: fetchAirports } =
    createAsyncProcess(fetchProfile)

  return {
    airports,
    isLoading,
    fetchAirports,
  }
}
