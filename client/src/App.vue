<script setup lang="ts">
import Panel from './components/Panel.vue'
</script>

<template>
  <div class="mainBackground"><Panel /></div>
</template>

<style>
.mainBackground {
  display: flex;
  flex-direction: row-reverse;
  background: url('./assets/background-storm.jpg') no-repeat center center fixed;
  background-size: cover;
  height: 100%;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
</style>
