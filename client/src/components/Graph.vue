<template>
  <LineChart class="graph" v-bind="lineChartProps" />
</template>

<script lang="ts">
import { LineChart, useLineChart } from 'vue-chart-3'
import { Chart, ChartData, ChartOptions, registerables } from 'chart.js'
import {
  computed,
  onMounted,
  ref,
  toRefs,
  defineComponent,
  ComputedRef,
} from 'vue'

Chart.register(...registerables)

export default defineComponent({
  components: { LineChart },
  mounted() {},
  props: {
    measures: {
      type: Array,
    },
  },
  computed: {
    // need annotation
    chartValues(): number[] {
      if (!this.measures) return []
      return this.measures?.map<number>((v: any) => v.value)
    },
  },
  setup(props) {
    const { measures } = toRefs(props)
    const chartLabel: ComputedRef<string[] | undefined> = computed(() =>
      measures.value?.map((v: any) => '')
    )
    const chartValues: any = computed(() =>
      measures.value?.map<number>((v: any) => v.value)
    )

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
      labels: chartLabel.value,
      datasets: [
        {
          data: chartValues.value,
          backgroundColor: gradientFill.value,
          fill: true,
          tension: 0.3,
          pointRadius: 0,
          spanGaps: true,
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
      plugins: {
        legend: {
          display: false,
        },
      },
      scales: {
        y: {
          ticks: {
            color: 'white',
          },
          beginAtZero: false,
        },
      },
    }))

    const { lineChartProps, lineChartRef } = useLineChart({
      chartData,
      options,
    })

    return {
      lineChartProps,
      options,
    }
  },
})
</script>
<style>
.graph {
  height: 300px;
}
</style>
