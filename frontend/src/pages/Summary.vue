<template>
  <div class="summary-page">
    <nav class="navbar">
      <div class="logo">WalletLog</div>
      <router-link to="/home" class="back-btn">← Back</router-link>
    </nav>

    <div class="summary-container">
      <h2>Expense Summary</h2>

      <div class="date-range">
        <input type="date" v-model="startDate" @change="refetch" />
        <input type="date" v-model="endDate" @change="refetch" />
      </div>

      <div v-if="loading">Loading summary...</div>

      <div v-else-if="summary">
        <p><strong>Total Expenses:</strong> ₹{{ summary.totalAmount.toFixed(2) }}</p>
        <p><strong>Total Count:</strong> {{ summary.totalCount }}</p>

        <PieChart v-if="chartData" :chart-data="chartData" />
      </div>

      <div v-else>No data found for selected range.</div>

      <div v-if="error" style="color: red; margin-top: 1rem;">
        Error: {{ error.message }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import gql from 'graphql-tag'
import { useQuery } from '@vue/apollo-composable'
import PieChart from './PieChart.vue'
import { categoryColors, fallbackColor } from '../utils/categoryColors.js' 

const today = new Date()
const formatDate = (date) => date.toISOString().slice(0, 10)

const startDate = ref(formatDate(today))
const endDate = ref(formatDate(today))

const EXPENSE_SUMMARY = gql`
  query ExpenseSummary($startDate: String!, $endDate: String!) {
    expenseSummary(startDate: $startDate, endDate: $endDate) {
      totalAmount
      totalCount
      byCategory {
        category
        total
      }
    }
  }
`

const variables = ref({
  startDate: startDate.value,
  endDate: endDate.value,
})


const { result, loading, error, refetch } = useQuery(EXPENSE_SUMMARY, variables, {
  fetchPolicy: 'network-only',
})

watch([startDate, endDate], ([newStart, newEnd]) => {
  variables.value = { startDate: newStart, endDate: newEnd }
})

const summary = ref(null)
const chartData = ref(null)

watch(result, (res) => {
  if (res?.expenseSummary) {
    summary.value = res.expenseSummary

    const labels = res.expenseSummary.byCategory.map(c => c.category)
    const data = res.expenseSummary.byCategory.map(c => c.total)

    const backgroundColors = labels.map(label => categoryColors[label] || fallbackColor)

    chartData.value = {
      labels,
      datasets: [
        {
          label: 'Expenses by Category',
          backgroundColor: backgroundColors,
          data,
        },
      ],
    }
  } else {
    summary.value = null
    chartData.value = null
  }
})


onMounted(() => {
  refetch()
})
</script>

<style scoped>
.summary-page {
  padding: 1rem;
  font-family: 'Segoe UI', sans-serif;
  background-color: #f9fafb;
  min-height: 100vh;
}

.navbar {
  display: flex;
  justify-content: space-between;
  padding: 1rem;
  background: #2c3e50;
  color: white;
}

.back-btn {
  text-decoration: none;
  color: white;
}

.summary-container {
  max-width: 700px;
  margin: auto;
  padding: 2rem;
  background: white;
  border-radius: 12px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  text-align: center;
}

.date-range {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 1rem;
}

.date-range input,
.date-range button {
  padding: 0.5rem;
  font-size: 1rem;
  border-radius: 6px;
  border: 1px solid #ccc;
}

.date-range button {
  background-color: #2c3e50;
  color: white;
  cursor: pointer;
  border: none;
  transition: background-color 0.3s ease;
}

.date-range button:hover {
  background-color: #1a252f;
}
</style>
