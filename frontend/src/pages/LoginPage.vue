<template>
  <div class="wrapper">
    <div class="form-container">
      <h2>Login</h2>
      <form @submit.prevent="handleLogin">
        <input
          v-model="form.email"
          type="email"
          placeholder="Email"
          required
        />
        <input
          v-model="form.password"
          type="password"
          placeholder="Password"
          required
        />
        <button type="submit">Login</button>
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>

      <p class="register-link">
        Don’t have an account?
        <RouterLink to="/register">Register here</RouterLink>
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import gql from 'graphql-tag'
import { useMutation } from '@vue/apollo-composable'
import { useRouter } from 'vue-router'

const router = useRouter()

const LOGIN_MUTATION = gql`
  mutation Login($input: LoginInput!) {
    login(input: $input) {
      token
      user {
        id
        email
        name
      }
    }
  }
`

const form = ref({
  email: '',
  password: '',
})

const errorMessage = ref(null)
const { mutate: login } = useMutation(LOGIN_MUTATION)

async function handleLogin() {
  errorMessage.value = null

  try {
    const { data } = await login({
      input: {
        email: form.value.email,
        password: form.value.password,
      },
    })

    if (data?.login?.token) {
      localStorage.setItem('token', data.login.token)
      localStorage.setItem('user', JSON.stringify(data.login.user))
      router.push('/home')
    } else {
      errorMessage.value = 'Login failed. Please check your credentials.'
    }
  } catch (err) {
    errorMessage.value = err.message || 'An error occurred during login.'
  }
}
</script>

<style scoped>
.wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 1rem;
  background-color: #f3f4f6;
}

.form-container {
  max-width: 400px;
  width: 100%;
  background-color: #ffffff;
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
  box-sizing: border-box;
  text-align: center;
}

h2 {
  margin-bottom: 1.5rem;
  font-size: 1.8rem;
  color: #2c3e50;
}

input {
  width: 100%;
  margin-bottom: 1rem;
  padding: 0.75rem 1rem;
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  font-size: 1rem;
  box-sizing: border-box;
}

input:focus {
  border-color: #2c3e50;
  outline: none;
}

button {
  width: 100%;
  margin-top: 0.5rem;
  padding: 0.75rem;
  background-color: #2c3e50;
  color: white;
  border: none;
  border-radius: 0.5rem;
  font-size: 1rem;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #1a242f;
}

.error {
  color: red;
  margin-top: 0.5rem;
  font-size: 0.9rem;
}

.register-link {
  margin-top: 1.25rem;
  font-size: 0.95rem;
  color: #444;
}

.register-link a {
  color: #2c3e50;
  text-decoration: underline;
  font-weight: 600;
  transition: color 0.2s ease;
}

.register-link a:hover {
  color: #1a242f;
}
</style>
