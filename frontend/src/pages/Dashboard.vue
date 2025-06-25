<template>
  <div class="dashboard">
    <h1>Welcome, {{ user.name }}</h1>
    <p>Your dashboard summary will appear here.</p>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const user = ref({ name: '', email: '' })

onMounted(() => {
  const storedUser = localStorage.getItem('user')
  if (storedUser) {
    user.value = JSON.parse(storedUser)
  }
})
</script>

<style scoped>
.dashboard {
  max-width: 600px;
  margin: auto;
  text-align: center;
}
</style>
