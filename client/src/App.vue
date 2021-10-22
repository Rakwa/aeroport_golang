<script setup lang="ts">
import TodayMeteo from './components/TodayMeteo.vue'
import Panel from './components/Panel.vue'
</script>

<template>
  <div class="mainBackground">
    <TodayMeteo />
    <Panel />
  </div>
</template>

<style>
.mainBackground {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  background: url('./assets/background-storm.jpg') no-repeat center center fixed;
  background-size: cover;
  height: 100%;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
* {
  @apply text-white;
}
</style>
