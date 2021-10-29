<template>
  <div class="flex">
    <div
      v-for="airport in airports"
      class="bg-blue-400 rounded-lg px-3 py-1 mr-6 cursor-pointer"
      @click="airportSelected = airport"
    >
      {{ airport.label }}
    </div>
  </div>

  aéroport séléctionné : {{ airportSelected?.label }}
</template>
<script lang="ts">
import { defineComponent, ref, watch } from 'vue'
import { useAirports } from '../composable/useAirports'
export default defineComponent({
  emits: ['onAirportChange'],
  async setup(_, { emit }) {
    const { airports, fetchAirports } = useAirports()
    const airportSelected = ref<Airport>()
    watch(airportSelected, () => {
      emit('onAirportChange', airportSelected.value?.acronym)
    })
    const result = await fetchAirports()
    airportSelected.value = result[0]
    return {
      airportSelected,
      airports,
    }
  },
})
</script>
