<template>
  <div class="landing-container">
    <img src="/logo.png" alt="WalletLog Logo" class="logo" />

    <div class="banner">
      <h1>Welcome to WalletLog</h1>
      <p>Your personal expense tracker. Stay on top of your finances.</p>

      <div class="buttons">
        <RouterLink to="/login" class="btn">Login</RouterLink>
        <RouterLink to="/register" class="btn btn-outline">Register</RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup>

</script>

<style scoped>
.landing-container {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  overflow: hidden;
  justify-content: center;
  align-items: center;
  padding: 2rem;
  height: 100vh;
  text-align: center;
}

.logo {
  width: 120px;
  margin-bottom: 1.5rem;
}

.banner {
  max-width: 500px;
}

.banner h1 {
  font-size: 2rem;
  margin-bottom: 0.5rem;
}

.banner p {
  color: #555;
  margin-bottom: 1.5rem;
}

.buttons {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

.btn {
  padding: 0.5rem 1.5rem;
  background-color: #2c3e50;
  color: white;
  border-radius: 5px;
  text-decoration: none;
}

.btn-outline {
  background-color: transparent;
  color: #2c3e50;
  border: 1px solid #2c3e50;
}
</style>
