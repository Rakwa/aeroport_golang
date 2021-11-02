<template>
  <div class="panel">
    <div class="cardsContainer">
      <ChooseAirport @onAirportChange="onAirportChange" />
      <template v-if="!isLoading">
        <TodayInfos class="mb-5" :todayData="airport?.data?.today" />
        <BeforeCard :measures="airport?.data?.measures" />
      </template>
      <p v-else>Chargement des datas</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAirportData } from '../composable/useAirportData'
import ChooseAirport from './ChooseAirport.vue'
import TodayInfos from './TodayInfos.vue'
import BeforeCard from './BeforeCard.vue'

const { airport, isLoading, fetchData } = useAirportData()
const onAirportChange = async (id: string) => {
  airport.value = { acronym: id }
  await fetchData()
}
</script>
<style>
.panel {
  overflow: scroll;
  background-color: #4847477c;
  -webkit-backdrop-filter: blur(7px);
  backdrop-filter: blur(7px);
  width: 510px;
  height: 100%;
}
.cardsContainer {
  padding: 20px 35px;
}
</style>
