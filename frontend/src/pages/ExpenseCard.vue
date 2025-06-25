<template>
  <div class="expense-card">
    <p class="category" :style="{ backgroundColor: categoryColor }">{{ expense.category }}</p>

    <div class="expense-header">
      <p class="amount">₹ {{ expense.amount.toFixed(2) }}</p>
      <div class="actions">
        <button class="edit-btn" @click="goToEdit" title="Edit">✏️</button>
        <button class="delete-btn" @click="onDelete" title="Delete">🗑️</button>
      </div>
    </div>

    <p class="description">{{ expense.description }}</p>
  </div>
</template>

<script setup>
import { defineProps, defineEmits, computed } from 'vue'
import { useRouter } from 'vue-router'

const categoryColors = {
  'Food & Dining': '#3498db',
  'Transportation': '#2ecc71',
  'Housing & Utilities': '#e74c3c',
  'Healthcare': '#f1c40f',
  'Entertainment & Leisure': '#9b59b6',
  'Education & Personal Development': '#1abc9c',
  'Clothing & Personal Care': '#e67e22',
  'Savings & Investments': '#34495e',
  'Travel': '#9b59b6',
  'Miscellaneous': '#95a5a6',
}

const fallbackColor = '#dfe6e9'

const props = defineProps({
  expense: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['delete'])

const router = useRouter()

function goToEdit() {
  router.push({
    name: 'EditExpense',
    params: {
      id: props.expense.id,
      expense: JSON.stringify(props.expense), 
    },
  })
}

function onDelete() {
  emit('delete', props.expense.id)
}


const categoryColor = computed(() => {
  return categoryColors[props.expense.category] || fallbackColor
})
</script>

<style scoped>
.expense-card {
  display: flex;
  flex-direction: column;
  background: #ffffff;
  border-radius: 10px;
  box-sizing: border-box;
  padding: 1.5rem;
  margin: 1rem 0;
  width: 100%;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s ease;
}

.expense-card:hover {
  transform: translateY(-2px);
}

.expense-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: nowrap;
}

.category {
  display: inline-block;
  padding: 4px 10px;
  color: #2c3e50;
  font-size: 0.85rem;
  font-weight: 600;
  text-align: center;
  text-transform: capitalize;
  border-radius: 6px;
  max-width: fit-content;
  margin-bottom: 0.5rem;
}

.amount {
  font-weight: bold;
  color: #2c3e50;
  margin: 0;
}

.description {
  margin-top: 0.5rem;
  color: #555;
  word-wrap: break-word;
  line-height: 1.4;
}

.actions {
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;
}

.actions button {
  font-size: 1.2rem;
  background: none;
  border: none;
  cursor: pointer;
}

.edit-btn {
  color: #3498db;
}
.edit-btn:hover {
  color: #2980b9;
}

.delete-btn {
  color: #e74c3c;
}
.delete-btn:hover {
  color: #c0392b;
}
</style>
