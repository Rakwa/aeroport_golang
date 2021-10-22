import { ref } from 'vue'
export function useAirport() {
    const airportName = ref('Nantes')
    return { airportName }
}