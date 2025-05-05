<template>
  <Navbar :user="user" />
  <router-view />
     <!-- Footer -->
     <footer class="footer">
      <p>© 2025 Perfumer Co. All rights reserved.</p>
    </footer>
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
/* Global Reset for margin and padding */
* {
  margin: 0px;
  padding: 0px;
}

/* Global Base Font Setup */
html, body {
  margin: 0;
  font-family: 'Poppins', sans-serif; /* Use Poppins globally */
  background: linear-gradient(to bottom right, #ffe6e6, #fff0f5, #ffe0b3); /* Apply gradient */
  height: 100vh; /* Ensure body takes full height */
}

/* Headings with a fancy serif */
h1, h2, h3 {
  font-family: 'Playfair Display', serif; /* Serif for headings */
}

/* Footer */
.footer {
  text-align: center;
  padding: 2rem;
  font-size: 0.9rem;
  color: #888;
  background: #f5f5f5;
  margin-top: 4rem;
}
</style>