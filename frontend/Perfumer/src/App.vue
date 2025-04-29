<template>
  <Navbar :user="user" />
  <router-view />
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import Navbar from './components/Navbar.vue'

const user = ref(null)  // Track logged-in user info

// Check if the user is already logged in on mount
onMounted(() => {
  const savedUser = localStorage.getItem('userToken')
  if (savedUser) {
    user.value = { username: localStorage.getItem('username') }
  }
})

// Watch for changes to localStorage and update user state
watch(() => localStorage.getItem('userToken'), (newToken) => {
  if (newToken) {
    user.value = { username: localStorage.getItem('username') } // Update user info
  } else {
    user.value = null // Reset user info on logout
  }
})
</script>

<style>
/* Your styles */
</style>