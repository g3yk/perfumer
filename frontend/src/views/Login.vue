<template>
    <div class="auth-page">
      <div class="auth-card">
        <h1>Login</h1>
  
        <form @submit.prevent="handleLogin">
          <div class="input-group">
            <input
              v-model="email"
              type="email"
              placeholder="Email"
              class="auth-input"
              required
            />
          </div>
          <div class="input-group">
            <input
              v-model="password"
              type="password"
              placeholder="Password"
              class="auth-input"
              required
            />
          </div>
          <button type="submit" class="auth-btn">Login</button>
        </form>
  
        <p class="auth-toggle">
          Don't have an account? 
          <router-link to="/signup">Sign Up here</router-link>
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  import { useToast } from 'vue-toastification'
  
  const email = ref('')
  const password = ref('')
  const router = useRouter()
  const toast = useToast() // Accessing the toast instance
  
  // Handle Login
  function handleLogin() {
    const storedEmail = localStorage.getItem('email')
    const storedPassword = localStorage.getItem('password')
  
    if (storedEmail === email.value && storedPassword === password.value) {
      toast.success('Login successful!', {
        position: 'top-right',
        timeout: 3000,
        hideProgressBar: false,
      })
      localStorage.setItem('userToken', 'fake-token') // You can replace this with a real token
      router.push('/home') // Redirect to a protected route (e.g., Checkout)
    } else {
      toast.error('Invalid credentials. Please try again.', {
        position: 'top-right',
        timeout: 3000,
        hideProgressBar: false,
      })
    }
  }
  </script>
  
  <style scoped>
  /* Main container */
  .auth-page {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: linear-gradient(45deg, #3498db, #8e44ad);
    font-family: 'Roboto', sans-serif;
  }
  
  /* The card where forms are displayed */
  .auth-card {
    background: blur(12px);
    padding: 2rem;
    border-radius: 15px;
    width: 100%;
    max-width: 400px;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    text-align: center;
  }
  
  .auth-card h1 {
    margin-bottom: 1.5rem;
    color: white;
    font-size: 2rem;
    font-weight: 500;
  }
  
  /* Input fields styling */
  .input-group {
    margin-bottom: 1.2rem;
    width: 92%;
  }
  
  .auth-input {
    width: 100%;
    padding: 0.8rem;
    border: 2px solid #ddd;
    border-radius: 8px;
    font-size: 1rem;
    transition: border-color 0.3s ease-in-out;
  }
  
  .auth-input:focus {
    outline: none;
    border-color: #3498db;
  }
  
  /* Button styling */
  .auth-btn {
    width: 100%;
    padding: 1rem;
    background-color: #3498db;
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 1.1rem;
    cursor: pointer;
    transition: background 0.3s ease-in-out;
  }
  
  .auth-btn:hover {
    background-color: #2980b9;
  }
  
  /* Link and toggle message */
  .auth-toggle {
    margin-top: 1rem;
    font-size: 0.9rem;
    color: #555;
  }
  
  .auth-toggle a {
    color: #3498db;
    text-decoration: none;
    transition: color 0.3s ease-in-out;
  }
  
  .auth-toggle a:hover {
    color: #2980b9;
  }
  </style>