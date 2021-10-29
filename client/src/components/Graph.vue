<!-- <template>
  {{ doughnutChartProps }}
  <DoughnutChart v-bind="doughnutChartProps" />
</template>

<script lang="ts">
import { LineChart, useLineChart } from 'vue-chart-3'
import { Chart, ChartData, ChartOptions, registerables } from 'chart.js'
import { computed, ref } from 'vue'

Chart.register(...registerables)

export default {
  components: { LineChart },
  setup() {
    const dataValues = ref([30, 40, 60, 70, 5])
    const toggleLegend = ref(true)

    const chartData = computed<ChartData<'line'>>(() => ({
      labels: ['Paris', 'Nîmes', 'Toulon', 'Perpignan', 'Autre'],
      datasets: [
        {
          data: dataValues.value,
          backgroundColor: ['#77CEFF', '#0079AF', '#123E6B', '#97B0C4', '#A5C8ED'],
        },
      ],
    }));
    const options = computed<ChartOptions<'line'>>(() => ({
      scales: {
        myScale: {
          type: 'logarithmic',
          position: toggleLegend.value ? 'left' : 'right',
        },
      },
      plugins: {
        legend: {
          position: toggleLegend.value ? 'top' : 'bottom',
        },
        title: {
          display: true,
          text: 'Chart.js Doughnut Chart',
        },
      },
    }))

    const { lineChartProps, lineChartRef } = useLineChart({
      chartData: chartData,
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
    }
  },
}
</script> -->
