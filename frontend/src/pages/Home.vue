<script setup>
import { ref, computed, onMounted } from 'vue'
import gql from 'graphql-tag'
import axios from 'axios'
import { useQuery } from '@vue/apollo-composable'
import ExpenseCard from './ExpenseCard.vue'

const API_BASE = 'https://walletlog-ju32.onrender.com/",' 

const selectedDate = ref('')
const dateInput = ref(null)
const expenses = ref([])

const GET_EXPENSES_QUERY = gql`
  query GetExpenses($date: String) {
    getExpenses(date: $date) {
      id
      amount
      category
      description
      date
    }
  }
`

const variables = ref({ date: selectedDate.value })
const { result, loading, error, refetch } = useQuery(GET_EXPENSES_QUERY, variables)

onMounted(() => {
  const today = new Date().toISOString().split('T')[0]
  selectedDate.value = today
  fetchExpenses()
})

function openCalendar() {
  dateInput.value?.showPicker?.() || dateInput.value?.click()
}

async function fetchExpenses() {
  variables.value.date = selectedDate.value
  const { data } = await refetch(variables.value)
  expenses.value = data.getExpenses
}

function logout() {
  localStorage.removeItem('token')
  localStorage.removeItem('user')
  window.location.href = '/login'
}

const formattedDate = computed(() => {
  if (!selectedDate.value) return ''
  const date = new Date(selectedDate.value)
  return date.toLocaleDateString('en-GB', {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
  })
})

async function handleDelete(id) {
  try {
    const token = localStorage.getItem('token')
    const mutation = `
      mutation {
        deleteExpense(id: "${id}")
      }
    `
    const response = await axios.post(API_BASE, { query: mutation }, {
      headers: {
        Authorization: `Bearer ${token}`
      }
    })

    if (response.data?.data?.deleteExpense) {
      expenses.value = expenses.value.filter(e => e.id !== id)
    } else {
      console.error('Failed to delete expense.')
    }
  } catch (err) {
    console.error('Error deleting expense:', err)
  }
}
</script>

<template>
  <div class="home-page">
    <nav class="navbar">
      <div class="logo">WalletLog</div>
      <div class="nav-links">
        <router-link to="/summary" class="nav-link">Summary</router-link>
        <button @click="logout" class="logout-btn">Logout</button>
      </div>
    </nav>

    <div class="content">
      <div class="date-picker">
        <span class="calendar-icon" @click="openCalendar">📅</span>
        <input
          ref="dateInput"
          type="date"
          v-model="selectedDate"
          @change="fetchExpenses"
          class="visually-hidden-input"
        />
        <span class="formatted-date">{{ formattedDate }}</span>
      </div>

      <div v-if="loading" class="expenses-list">Loading expenses...</div>
      <div v-else-if="expenses.length === 0" class="expenses-list">No expenses found for this day.</div>

      <div class="expenses-list">
        <ExpenseCard
          v-for="expense in expenses"
          :key="expense.id"
          :expense="expense"
          @delete="handleDelete"
        />
      </div>
    </div>

    <router-link :to="{ name: 'AddExpense' }" class="add-expense-btn">+</router-link>
  </div>
</template>

<style scoped>
.home-page {
  font-family: 'Segoe UI', sans-serif;
  box-sizing: border-box;

  background-color: #f9fafb;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background: #2c3e50;
  color: white;
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.nav-link {
  color: white;
  text-decoration: none;
}

.logout-btn {
  background: #ff4d4d;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 6px;
  cursor: pointer;
}

.content {
  padding: 2rem;
}

.date-picker {
  display: flex;
  align-items: center;
  gap: 1rem;
  overflow: hidden;
  box-sizing: border-box;
  background: white;
  padding: 1rem;
  justify-content: space-between;
  border-radius: 8px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.05);
  width: 100%;
  position: relative;
}

.calendar-icon {
  font-size: 1.5rem;
  color: #2c3e50;
  cursor: pointer;
}

.visually-hidden-input {
  position: absolute;
  opacity: 0;
  pointer-events: none;
  height: 0;
  width: 0;
  border: none;
}

.formatted-date {
  font-weight: 600;
  color: #2c3e50;
}

.expenses-list {
  margin-top: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 1rem;
  overflow-y: auto;
  max-height: 70vh;
  box-sizing: border-box;
}


.add-expense-btn {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  text-decoration: none;
  background-color: #2c3e50;
  color: white;
  font-size: 2rem;
  border-radius: 20%;
  width: 56px;
  height: 56px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  transition: transform 0.2s ease;
  z-index: 1000;
}

.add-expense-btn:hover {
  transform: scale(1.1);
  background-color: #1a252f;
}
</style>
