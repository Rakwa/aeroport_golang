<template>
  <div class="panel">
    <p class="text-white">{{ count }}</p>
    <div class="cursor-pointer" @click="incCount">AUGMENTER</div>
    <h1>{{ fortnite }}</h1>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
const count = ref(0)
const fortnite = ref('fortnite')
function incCount() {
  count.value++
}
</script>
<style>
.panel {
  background-color: #484747;
  opacity: 0.7;
  width: 550px;
  height: 100%;
}
</style>
