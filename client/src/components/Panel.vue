<template>
  <div class="panel">
    <Suspense>
      <template #default>
        <div>
          <ChooseAirport />
          <TodayInfos />
        </div>
      </template>
      <template #fallback> Chargement Suspense </template>
    </Suspense>
  </div>
</template>

<script setup lang="ts">
import ChooseAirport from './ChooseAirport.vue'
import TodayInfos from './TodayInfos.vue'
</script>
<style>
.panel {
  background-color: #484747;
  opacity: 0.7;
  width: 550px;
  height: 100%;
}
</style>
