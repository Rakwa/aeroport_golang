<template>{{ airports }} {{ isLoading }}</template>
<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import { useAirport } from '../composable/useAirport'

const { airports, fetchAirports, isLoading } = useAirport()

await fetchAirports()
</script>
