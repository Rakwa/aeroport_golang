<template>
  <div class="flex mb-5 dropdownContainer">
    <div class="flex justify-center w-full">
      <p class="airport">Nantes</p>
    </div>
  </div>
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
<style>
.dropdownContainer {
  border-bottom: 1px solid #b3b3b399;
  padding-bottom: 7px;
  padding-top: 5px;
}
.airport {
  font-size: 35px;
  font-weight: 400;
  color: #f9fafa;
  letter-spacing: 3px;
}
</style>
