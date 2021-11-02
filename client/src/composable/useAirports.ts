import { ref } from 'vue'
import { getAirportsTowns } from '../services/airports'
import createAsyncProcess from '../utils/createAsyncProcess'

export function useAirports() {
  const airports = ref<Airport[]>()
  const airportSelected = ref<Airport | null>(null)

  async function getAirports(): Promise<Airport[]> {
    const result = await getAirportsTowns()
      console.log(result)
    airports.value = result
    airportSelected.value = airports.value[0]
    return result
  }

  const { active: isLoading, run: fetchAirports } =
    createAsyncProcess(getAirports)

  return {
    airports,
    isLoading,
    fetchAirports,
    airportSelected,
  }
}
