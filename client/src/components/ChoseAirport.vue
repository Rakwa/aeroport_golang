<template>
  <div class="panel">
    <p class="text-white">{{ airportName }}</p>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useAirport } from '../composable/useAirport'

const { airportName }: any = useAirport()
</script>

<style>
.panel {
  background-color: #484747;
  opacity: 0.7;
  width: 550px;
  height: 100%;
}
</style>
