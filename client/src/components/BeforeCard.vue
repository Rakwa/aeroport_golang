<template>
  <div class="card">
    <p class="w-full text-center beforeTitle">Précédemment</p>
    <!--<Datepicker class="mb-7" />-->
    <Tabs :tabs="['Température', 'Pression', 'Vent']" />
    <Graph />
  </div>
</template>
<script lang="ts" setup>
import Datepicker from './global/Datepicker.vue'
import Tabs from './global/Tabs.vue'
import Graph from './Graph.vue'
</script>
<style>
.card {
  background-color: #ffffff18;
  border-radius: 10px;
  padding: 15px;
}
.beforeTitle {
  padding-bottom: 15px;
  font-size: 18px;
  letter-spacing: 1.5px;
}
</style>
