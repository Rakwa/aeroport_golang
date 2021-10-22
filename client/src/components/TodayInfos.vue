<template><br />Today Infos {{ todaysInfos }}</template>
<script setup lang="ts">
import { ref } from 'vue'
import { useAirport } from '../composable/useAirport'
const todaysInfos = ref('')
const sleep = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms))
await sleep(5000)
todaysInfos.value = 'temperature : 39, vent: 40, pression : 40'
</script>
