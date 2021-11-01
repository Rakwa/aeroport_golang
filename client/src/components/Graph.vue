<template>
  <LineChart class="graph" v-bind="lineChartProps" />
</template>

<script lang="ts">
import { LineChart, useLineChart } from 'vue-chart-3'
import { Chart, ChartData, ChartOptions, registerables } from 'chart.js'
import { computed, onMounted, ref } from 'vue'

Chart.register(...registerables)

export default {
  components: { LineChart },
  mounted() {},
  setup() {
    const dataValues = ref([30, 40, 60, 70, 58, 30, 40, 60, 70, 5])
    const toggleLegend = ref(true)
    const linechart = ref<any>(null)
    const root = ref(null)
    const gradientFill = ref<any>(null)

    onMounted(() => {
      const a: any = document.querySelector('.graph > canvas')
      if (!a) return
      gradientFill.value = a.getContext('2d').createLinearGradient(0, 0, 0, 500)
      gradientFill.value.addColorStop(0, '#3c91ad')
      gradientFill.value.addColorStop(1, '#3c91ad00')
    })

    const chartData = computed<ChartData<'line'>>(() => ({
      labels: [
        'Paris',
        'Nîmes',
        'Toulon',
        'Perpignan',
        'Autre',
        'Paris',
        'Nîmes',
        'Toulon',
        'Perpignan',
        'Autre',
      ],
      datasets: [
        {
          data: dataValues.value,
          backgroundColor: gradientFill.value,
          fill: true,
          tension: 0.5,
          pointRadius: 0,
        },
      ],
    }))
    const options = computed<ChartOptions<'line'>>(() => ({
      responsive: true,
      maintainAspectRatio: false,
      bezierCurve: true,
      elements: {
        point: {
          radius: 0,
        },
      },
    }))

    const { lineChartProps, lineChartRef } = useLineChart({
      chartData,
    })

    function switchLegend() {
      toggleLegend.value = !toggleLegend.value
    }

    return {
      switchLegend,
      chartData,
      options,
      lineChartRef,
      lineChartProps,
      linechart,
      root,
    }
  },
}
</script>
<style>
.graph {
  height: 300px;
}
</style>
