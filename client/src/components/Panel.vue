<template>
  <div class="panel">
    <p>{{ count }}</p>
    <div class="cursor-pointer" @click="incCount">AUGMENTER</div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
const count = ref(0)
function incCount() {
  count.value++
}
</script>
<style>
.panel {
  background-color: #484747;
  opacity: 0.7;
  width: 550px;
}
</style>
