<template>
  <div class="form-container">
    <h2>{{ isEdit ? 'Edit Expense' : 'Add New Expense' }}</h2>

    <form @submit.prevent="handleSubmit">
      <input
        v-model.number="form.amount"
        type="number"
        placeholder="Amount"
        required
      />


      <select v-model="form.category" required>
        <option disabled value="">Select Category</option>
        <option v-for="cat in categories" :key="cat" :value="cat">
          {{ cat }}
        </option>
      </select>

      <input
        v-model="form.description"
        type="text"
        placeholder="Description"
        required
      />
      <input
        v-model="form.date"
        type="date"
        required
      />

      <button type="submit">{{ isEdit ? 'Update' : 'Add' }} Expense</button>

      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'

const ADD_EXPENSE = gql`
  mutation AddExpense($input: NewExpense!) {
    addExpense(input: $input) {
      id
    }
  }
`

const UPDATE_EXPENSE = gql`
  mutation UpdateExpense($input: UpdateExpenseInput!) {
    updateExpense(input: $input) {
      id
    }
  }
`

const GET_EXPENSE_BY_ID = gql`
  query GetExpenseById($id: ID!) {
    getExpenseById(id: $id) {
      id
      amount
      category
      description
      date
    }
  }
`

const router = useRouter()
const route = useRoute()
const isEdit = ref(route.name === 'EditExpense')
const errorMessage = ref(null)


const categories = [
  'Food & Dining',
  'Transportation',
  'Housing & Utilities',
  'Healthcare',
  'Entertainment & Leisure',
  'Education & Personal Development',
  'Clothing & Personal Care',
  'Savings & Investments',
  'Travel',
  'Miscellaneous',
]

const form = ref({
  id: '',
  amount: '',
  category: '',
  description: '',
  date: '',
})

const { mutate: addExpense } = useMutation(ADD_EXPENSE)
const { mutate: updateExpense } = useMutation(UPDATE_EXPENSE)

if (isEdit.value) {
  const expenseId = route.params.id
  const { result, loading, error } = useQuery(GET_EXPENSE_BY_ID, { id: expenseId })

  watchEffect(() => {
    if (result.value?.getExpenseById) {
      form.value = { ...result.value.getExpenseById }
    }
    if (error.value) {
      errorMessage.value = 'Failed to load expense'
    }
  })
} else {
  form.value.date = new Date().toISOString().split('T')[0]
  // Optionally set default category
  form.value.category = categories[0]
}

async function handleSubmit() {
  errorMessage.value = null
  try {
    if (isEdit.value) {
      const { data } = await updateExpense({ input: form.value })
      if (data?.updateExpense?.id) router.push('/home')
      else errorMessage.value = 'Update failed.'
    } else {
      
      const { id, ...inputWithoutId } = form.value
      const { data } = await addExpense({ input: inputWithoutId })
      if (data?.addExpense?.id) router.push('/home')
      else errorMessage.value = 'Adding failed.'
    }
  } catch (err) {
    errorMessage.value = err.message || 'Something went wrong.'
  }
}
</script>

<style scoped>
.form-container {
  width: 60%;
  margin: auto;
  padding: 2rem;
  background: #ffffff;
  border-radius: 12px;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.08);
  display: flex;
  flex-direction: column;
  text-align: center;
}

h2 {
  margin-bottom: 1.5rem;
  color: #2c3e50;
}

form {
  display: flex;
  flex-direction: column;
  width: 80%;
  justify-content: center;
  align-items: center;
}

input, select {
  padding: 0.75rem 1rem;
  margin-bottom: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-size: 1rem;
  width: 100%;
  box-sizing: border-box;
  transition: border 0.2s ease;
}

input:focus, select:focus {
  border-color: #2c3e50;
  outline: none;
}

button {
  background-color: #2c3e50;
  color: white;
  border: none;
  padding: 0.75rem;
  border-radius: 8px;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.3s ease;
  margin-top: 0.5rem;
}

button:hover {
  background-color: #1a252f;
}

.error {
  margin-top: 1rem;
  color: red;
  font-weight: 500;
}
</style>
