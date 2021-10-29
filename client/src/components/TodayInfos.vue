<template>
  <!-- TODO : Ici c'est pour toi loïs -->
  <div class="testLois"><p>Today Infos</p></div>
  {{}}
  <p>pression : {{ todayData?.pressure }}</p>
  <p>température : {{ todayData?.temp }}</p>
  <p>vent : {{ todayData?.wind }}</p>
</template>
<script lang="ts">
import { PropType, defineComponent } from 'vue'
export default defineComponent({
  props: {
    todayData: {
      type: Object as PropType<Measure>,
    },
  },
})
</script>
<style lang="scss" scoped>
.testLois {
  p {
    color: red;
  }
}
</style>
