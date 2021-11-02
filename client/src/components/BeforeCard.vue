<template>
  <div class="card">
    <p class="w-full text-center beforeTitle">3 derniers jours</p>
    <!--<Datepicker class="mb-7" />-->
    <Tabs :tabs="['Température', 'Pression', 'Vent']" @onChange="onTabChange" />
    <Graph class="mt-4" :measures="filteredMeasures()" />
  </div>
</template>
<script lang="ts" setup>
import { watch, ref } from 'vue'
import Datepicker from './global/Datepicker.vue'
import Tabs from './global/Tabs.vue'
import Graph from './Graph.vue'
const tabSelected = ref('Température')
const { measures } = defineProps({
  measures: {
    type: Array,
  },
})

const onTabChange = (e: string) => {
  console.log(e)
  tabSelected.value = e
}
const filteredMeasures = () => {
  if (tabSelected.value == 'Température') {
    return measures?.filter((e: any) => e.type == 'temperature')
  } else if (tabSelected.value == 'Pression') {
    return measures?.filter((e: any) => e.type == 'pressure')
  } else {
    return measures?.filter((e: any) => e.type == 'wind')
  }
}

watch(tabSelected, (newValue) => {})
</script>
<style scoped>
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
