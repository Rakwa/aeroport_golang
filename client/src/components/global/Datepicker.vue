<template>
  <v-date-picker
    is-range
    class="dateRange mt-1"
    :columns="1"
    v-model="periodRange"
  >
    <template v-slot="{ inputValue, inputEvents }">
      {{ inputValue }}
      <input
        class="bg-white border px-2 py-1 rounded"
        :value="inputValue.start"
        v-on="inputEvents"
      />
    </template>
  </v-date-picker>

  <!-- <vc-date-picker
    v-model="periodRange"
    :columns="2"
    is-range
    class="dateRange mt-1"
    @input="changePeriodRange"
  >
    <template v-slot="{ inputValue, inputEvents }">
      <div class="flex justify-center items-center">
        <input
          :value="inputValue.start"
          v-on="inputEvents.start"
          class="
            border
            px-2
            py-1
            w-32
            rounded
            focus:outline-none focus:border-indigo-300
          "
        />
        <svg
          class="w-4 h-4 mx-2"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M14 5l7 7m0 0l-7 7m7-7H3"
          />
        </svg>
        <input
          :value="inputValue.end"
          v-on="inputEvents.end"
          class="
            border
            px-2
            py-1
            w-32
            rounded
            focus:outline-none focus:border-indigo-300
          "
        />
      </div>
    </template>
  </vc-date-picker> -->
</template>
<script lang="ts" setup>
import { ref } from 'vue'
const periodRange = ref([new Date()])
const emit = defineEmits(['changePeriodRange'])

const changePeriodRange = () => {
  emit('changePeriodRange', periodRange)
}
</script>
