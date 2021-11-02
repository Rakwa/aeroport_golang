<script setup lang="ts">
import TodayMeteo from './components/TodayMeteo.vue'
import Panel from './components/Panel.vue'
import LoadingPage from './components/LoadingPage.vue'
</script>

<template>
  <div class="mainBackground">
    <TodayMeteo />
    <Suspense>
      <template #default>
        <Panel />
      </template>
      <template #fallback>
        <LoadingPage />
      </template>
    </Suspense>
  </div>
</template>

<style>
.mainBackground {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  background: url('./assets/bg-sunset.jpg') no-repeat center center fixed;
  background-size: cover;
  height: 100%;
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
* {
  @apply text-white;
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}
</style>
