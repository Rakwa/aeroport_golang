<template>
  <div class="flex mb-5 dropdownContainer">
    <div class="flex justify-center items-center w-full selectContainer">
      <select v-model="airportSelected" name="" id="">
        <option v-for="airport in airports" :value="airport.acronym">
          <div class="flex justify-center">
            <p>{{ airport.name }}</p>
          </div>
        </option>
      </select>
      <svg
        class="cursor-pointer ml-5"
        width="18"
        height="18"
        viewBox="0 0 32 16"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M16.1022 15.9956C16.4261 15.9755 16.734 15.8475 16.9767 15.6322L31.5138 2.54741C31.6569 2.41932 31.7734 2.26423 31.8564 2.09106C31.9395 1.91788 31.9876 1.73004 31.9979 1.53825C32.0082 1.34646 31.9806 1.15451 31.9166 0.973417C31.8526 0.792322 31.7535 0.625626 31.625 0.482902C31.4965 0.340177 31.3411 0.224229 31.1677 0.141686C30.9943 0.059143 30.8063 0.0116275 30.6145 0.00188103C30.4227 -0.00786546 30.2309 0.0203327 30.05 0.0848722C29.8691 0.149412 29.7027 0.24903 29.5604 0.377992L16 12.5881L2.43959 0.377992C2.29727 0.249029 2.1309 0.149413 1.95002 0.0848722C1.76914 0.0203318 1.57731 -0.00786497 1.38551 0.00188103C1.1937 0.011627 1.00571 0.0591208 0.832305 0.141664C0.658897 0.224207 0.503481 0.340176 0.374966 0.482902C0.246451 0.625627 0.147363 0.792321 0.0833816 0.973417C0.0194001 1.15451 -0.00821703 1.34644 0.00211308 1.53823C0.0124432 1.73002 0.0605175 1.91788 0.14358 2.09106C0.226643 2.26423 0.343062 2.41932 0.486164 2.54741L15.0233 15.6322C15.1694 15.7622 15.3403 15.8615 15.5257 15.9239C15.7111 15.9864 15.9072 16.0108 16.1022 15.9956Z"
          fill="#F9FAFA"
        />
      </svg>
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
    const airportSelected = ref<string>()
    watch(airportSelected, () => {
      emit('onAirportChange', airportSelected.value)
    })
    const result = await fetchAirports()
    airportSelected.value = result[0].acronym

    return {
      airportSelected,
      airports,
    }
  },
})
</script>
<style scoped>
.selectContainer {
  position: relative;
  display: flex;
  height: 3em;
  border-radius: 0.25em;
  overflow: hidden;
  border-bottom: 1px solid #b3b3b399;
  padding-bottom: 7px;
  padding-top: 5px;
}
.chevron {
  font-size: 100px;
}
select {
  font-size: 20px;
  font-weight: 400;
  color: #f9fafa;
  letter-spacing: 1.3px;
  text-align: center;
}

select {
  appearance: none;
  outline: 0;
  border: 0;
  box-shadow: none;
  /* Personalize */
  flex: 1;
  padding: 0 1em;
  color: #fff;
  background-color: transparent;
  background-image: none;
  cursor: pointer;
}

select option {
  color: black;
}
</style>
