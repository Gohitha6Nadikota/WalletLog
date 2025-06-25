<template>
  <div class="form-container">
    <h2>Add New Expense</h2>

    <form @submit.prevent="handleSubmit">
      <input
        v-model.number="form.amount"
        type="number"
        placeholder="Amount (e.g. 500)"
        required
      />
      <input
        v-model="form.category"
        type="text"
        placeholder="Category (e.g. Food)"
        required
      />
      <input
        v-model="form.description"
        type="text"
        placeholder="Description (e.g. Lunch at cafe)"
        required
      />
      <input
        v-model="form.date"
        type="date"
        required
      />

      <button type="submit">Add Expense</button>

      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useMutation } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { useRouter } from 'vue-router'

const router = useRouter()

const ADD_EXPENSE = gql`
  mutation AddExpense($input: NewExpense!) {
    addExpense(input: $input) {
      id
      amount
      category
      description
      date
    }
  }
`

const form = ref({
  amount: '',
  category: '',
  description: '',
  date: ''
})

onMounted(() => {
  form.value.date = new Date().toISOString().split('T')[0]
})

const errorMessage = ref(null)
const { mutate: addExpense } = useMutation(ADD_EXPENSE)

async function handleSubmit() {
  errorMessage.value = null
  try {
    const { data } = await addExpense({ input: form.value })
    if (data?.addExpense?.id) {
      router.push('/home')
    } else {
      errorMessage.value = 'Failed to add expense.'
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

input {
  padding: 0.75rem 1rem;
  margin-bottom: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  font-size: 1rem;
  width: 100%;
  box-sizing: border-box;
  transition: border 0.2s ease;
}

input:focus {
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
