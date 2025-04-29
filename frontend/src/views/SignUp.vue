<template>
    <div class="auth-page">
      <div class="auth-card">
        <h1>Create Account</h1>
  
        <form @submit.prevent="handleSignUp">
          <div class="input-group">
            <input
              v-model="username"
              type="text"
              placeholder="Username"
              class="auth-input"
              required
            />
          </div>
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
          <button type="submit" class="auth-btn">Sign Up</button>
        </form>
  
        <p class="auth-toggle">
          Already have an account? 
          <router-link to="/login">Login here</router-link>
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useRouter } from 'vue-router'
  
  const username = ref('')
  const email = ref('')
  const password = ref('')
  const router = useRouter()
  
  // Handle Sign Up
  function handleSignUp() {
    if (!username.value || !email.value || !password.value) {
      alert('Please fill out all fields.')
      return
    }
  
    // Save user details in localStorage
    localStorage.setItem('username', username.value)
    localStorage.setItem('email', email.value)
    localStorage.setItem('password', password.value)
    localStorage.setItem('userToken', 'fake-token') // You can replace this with a real token
  
    alert('Account created successfully!')
    router.push('/checkout') // Redirect to a protected route (e.g., Checkout)
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
    background: #fff;
    padding: 2rem;
    border-radius: 15px;
    width: 100%;
    max-width: 400px;
    box-shadow: 0 8px 24px rgba(0, 0, 0, 0.1);
    text-align: center;
  }
  
  .auth-card h1 {
    margin-bottom: 1.5rem;
    color: #333;
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