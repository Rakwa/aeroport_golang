<template>
  <div class="panel">
    <div>
      <ChooseAirport @onAirportChange="onAirportChange" />
      <template v-if="!isLoading">
        <TodayInfos :todayData="airport?.data?.today" />
        <Graph />
      </template>
      <p v-else>Chargement des datas</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useAirports } from '../composable/useAirports'
import { useAirportData } from '../composable/useAirportData'
import ChooseAirport from './ChooseAirport.vue'
import TodayInfos from './TodayInfos.vue'
import Graph from './Graph.vue'

const { airport, isLoading, fetchData } = useAirportData()
const onAirportChange = async (id: string) => {
  airport.value = { acronym: id }
  await fetchData()
}
</script>
<style>
.panel {
  background-color: #484747b2;
  width: 550px;
  height: 100%;
}
</style>
