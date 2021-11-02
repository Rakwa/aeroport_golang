<template>
  <div class="flex justify-between items-center tabsContainer">
    <div
      v-for="(tab, index) in tabs"
      class="tabItem h-full flex items-center justify-center"
      :class="{ selected: index == selected }"
      @click="onClick(index)"
    >
      {{ tab }}
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
const emit = defineEmits(['onChange'])
const { tabs } = defineProps({
  tabs: {
    type: Array,
    required: true,
  },
})
const selected = ref<number>(0)
const onClick = (index: number) => {
  selected.value = index
  emit('onChange', tabs[index])
}
</script>
<style>
.tabItem {
  font-size: 14px;
  color: white;
  width: 100%;
  box-sizing: content-box;
  cursor: pointer;
}
.tabsContainer {
  border: 1px solid #ffffff;
  border-radius: 6px;
  height: 35px;
}
.selected {
  border-radius: 6px;
  border: 1px solid #ffffff;

  background-color: #ffffff33;
}
</style>
