<template>
  <div>
    <canvas ref="canvas"></canvas>
  </div>
</template>

<script setup>
import { onMounted, watch, ref } from 'vue'
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js'

Chart.register(PieController, ArcElement, Tooltip, Legend)

const props = defineProps({
  chartData: {
    type: Object,
    required: true,
  },
})

const canvas = ref(null)
let chartInstance = null

onMounted(() => {
  if (!canvas.value) return

  chartInstance = new Chart(canvas.value.getContext('2d'), {
    type: 'pie',
    data: props.chartData,
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'bottom',
        },
        tooltip: {
          callbacks: {
            label: function(context) {
              return context.label + ': ₹' + context.parsed.toFixed(2)
            }
          }
        }
      }
    }
  })
})

watch(() => props.chartData, (newData) => {
  if (chartInstance) {
    chartInstance.data = newData
    chartInstance.update()
  }
})
</script>

<style scoped>
canvas {
  max-width: 100%;
  max-height: 400px;
  margin: auto;
  display: block;
}
</style>
