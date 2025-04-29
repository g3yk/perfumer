<template>
  <nav class="navbar">
    <div class="logo">
      <router-link to="/">Perfumer</router-link>
    </div>

    <!-- Profile Section (Display if user is logged in) -->
<div v-if="userToken" class="profile-section">
  <div class="hover-group">
    <span class="username">Hello, {{ username }}!</span>

    <div class="hover-actions">
      <router-link to="/profile">
        <button class="cta-button">My Profile</button>
      </router-link>

      <button class="logout-button" @click="signOut">Sign Out</button>
    </div>
  </div>
</div>

    <!-- Guest Section (Display if user is not logged in) -->
    <div v-else class="profile-section">
      <span class="username">Welcome, Guest!</span>
      <router-link to="/login">
        <button class="cta-button">Login</button>
      </router-link>
    </div>

    <div class="cart-icon" @click="toggleCart">
      🛒
      <transition name="pop">
        <span v-if="cartStore.cart.length > 0" class="cart-count">{{ cartStore.cart.length }}</span>
      </transition>
    </div>


    <!-- Cart Sidebar -->
    <div class="cart-sidebar" :class="{ open: cartStore.isNavbarCartOpen }">
      <h2>Your Cart</h2>

      <div v-if="cartStore.cart.length === 0" class="empty-cart">
        Cart is empty.
      </div>

      <div v-else>
        <div v-for="item in cartStore.cart" :key="item.id" class="cart-item">
          <p><strong>Top Notes:</strong> {{ item.topNotes?.join(', ') || 'None' }}</p>
          <p><strong>Heart Notes:</strong> {{ item.heartNotes?.join(', ') || 'None' }}</p>
          <p><strong>Base Notes:</strong> {{ item.baseNotes?.join(', ') || 'None' }}</p>

          <div class="quantity-controls">
            <button @click="decreaseQuantity(item)" class="qty-btn">-</button>
            <span>{{ item.quantity || 1 }}</span>
            <button @click="increaseQuantity(item)" class="qty-btn">+</button>
          </div>
        </div>
        <router-link to="/checkout">
          <button class="checkout-btn">Go to Checkout</button>
        </router-link>
      </div>
    </div>
    
  </nav>
  <div class="marquee-container">
  <div class="marquee">
    ✨ Welcome to Perfumer! Enjoy your scent journey ✨ Free shipping on orders over $50 ✨ All Customized Perfumes only 49,99.- ✨
  </div>
</div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, defineExpose, watch } from 'vue'
import { useToast } from 'vue-toastification'
import { useRouter } from 'vue-router'
import { useCartStore } from '@/stores/cart'

// Toast notification and router
const toast = useToast()
const router = useRouter()
const cartStore = useCartStore()

// Reactive states for user and cart
const userToken = ref(localStorage.getItem('userToken'))
const username = ref(localStorage.getItem('username') || 'Guest')

// Load cart from localStorage when Navbar loads
const savedCart = JSON.parse(localStorage.getItem('cart'))
if (savedCart && Array.isArray(savedCart)) {
  cartStore.cart = savedCart
}

// Watch cart changes and save to localStorage
watch(
  () => cartStore.cart,
  (newCart) => {
    localStorage.setItem('cart', JSON.stringify(newCart))
  },
  { deep: true }
)

// Toggle the navbar cart visibility based on cartStore's state
function toggleCart() {
  cartStore.toggleNavbarCart()
}

// Remove perfume from cart
function removeFromCart(id) {
  cartStore.removeFromCart(id)
  toast.error('Perfume removed from cart!', { timeout: 2000 })
}

// Decrease quantity of an item
function decreaseQuantity(item) {
  if (!item.quantity) {
    item.quantity = 1
  }
  item.quantity--

  if (item.quantity <= 0) {
    cartStore.removeFromCart(item.id)
    toast.error('Perfume removed from cart!', { timeout: 2000 })
  }
}

// Increase quantity of an item
function increaseQuantity(item) {
  if (!item.quantity) {
    item.quantity = 1
  }
  item.quantity++
}

// Handle sign out logic
function signOut() {
  localStorage.removeItem('userToken')
  localStorage.removeItem('username')
  
  userToken.value = null
  username.value = 'Guest'

  toast.success('You have been signed out.', { timeout: 2000 })

  router.push('/') // Redirect to homepage after logout
}

// Handle click outside to close the cart
function handleClickOutside(event) {
  const sidebar = document.querySelector('.cart-sidebar')
  const cartIcon = document.querySelector('.cart-icon')
  if (
    cartStore.isNavbarCartOpen &&
    sidebar &&
    !sidebar.contains(event.target) &&
    !cartIcon.contains(event.target)
  ) {
    cartStore.toggleNavbarCart() // Close the cart if clicking outside
  }
}

// Watch for changes to userToken and update the UI accordingly
watch(userToken, (newValue) => {
  if (newValue) {
    username.value = localStorage.getItem('username') || 'Guest' // Update username after login
  } else {
    username.value = 'Guest' // Reset username after logout
  }
})

// Watch for changes to localStorage to synchronize with the reactive state
onMounted(() => {
  document.addEventListener('click', handleClickOutside)
  router.beforeEach(() => {
    cartStore.isNavbarCartOpen = false // Close cart on navigation
    return true
  })
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

// Expose openCart function so outside components can call it
defineExpose({
  openCart: () => cartStore.toggleNavbarCart()
})
</script>

<style scoped>
/* Navbar Styles */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background-color: transparent;
  position: relative;
}

.logo a {
  font-size: 1.8rem;
  font-weight: bold;
  color: #333;
  text-decoration: none;
}

/* Profile Section Styles */
.profile-section {
  display: flex;
  align-items: center;
  margin-right: 20px;
}

.username {
  margin-right: 10px;
  font-size: 1rem;
  color: #333;
}

.cta-button {
  padding: 0.75rem 2rem;
  background-color: #ff6b6b;
  border: none;
  border-radius: 50px;
  color: white;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s ease;
}

.cta-button:hover {
  background-color: #ff4b4b;
}

/* Cart Icon Styles */
.cart-icon {
  position: relative;
  font-size: 2rem;
  cursor: pointer;
}

.cart-count {
  position: absolute;
  top: -8px;
  right: -12px;
  background-color: red;
  color: white;
  font-size: 0.75rem;
  padding: 2px 6px;
  border-radius: 50%;
  font-weight: bold;
}

/* Cart Sidebar Styles */
.cart-sidebar {
  position: fixed;
  top: 0;
  right: -400px;
  width: 300px;
  height: 100vh;
  background: white;
  box-shadow: -2px 0 5px rgba(0,0,0,0.2);
  padding: 1rem;
  transition: right 0.3s ease;
  overflow-y: auto;
  z-index: 1000;
}

.cart-sidebar.open {
  right: 0;
}

.cart-item {
  border-bottom: 1px solid #ddd;
  padding: 1rem 0;
}

.remove-btn {
  background-color: #dc3545;
  color: white;
  border: none;
  padding: 5px 10px;
  margin-top: 5px;
  border-radius: 4px;
  cursor: pointer;
}

.remove-btn:hover {
  background-color: #c82333;
}

.checkout-btn {
  width: 100%;
  padding: 10px;
  background-color: #28a745;
  color: white;
  margin-top: 1rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
}

.checkout-btn:hover {
  background-color: #218838;
}

.empty-cart {
  text-align: center;
  margin-top: 2rem;
  color: #888;
}

/* Animation for cart count */
.pop-enter-active, .pop-leave-active {
  transition: all 0.3s ease;
}

.pop-enter-from, .pop-leave-to {
  transform: scale(0);
  opacity: 0;
}

.logout-button {
  margin-left: 10px;
  padding: 0.75rem 1.5rem;
  background-color: #6c757d;
  border: none;
  border-radius: 50px;
  color: white;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s ease;
}

.logout-button:hover {
  background-color: #5a6268;
}

.hover-group {
  position: relative;
  display: flex;
  align-items: center;
}

.username {
  font-size: 1rem;
  color: #333;
  cursor: pointer;
  padding: 0.5rem;
  transition: color 0.3s;
}

.username:hover {
  color: #ff4b4b;
}

.hover-actions {
  display: none;
  align-items: center;
  gap: 10px;
  margin-left: 10px;
}

.hover-group:hover .hover-actions {
  display: flex;
}

.marquee-container {
  width: 100%;
  overflow: hidden;
  white-space: nowrap;
  color: #333;
  padding: 8px 0;
  border-bottom: 1px solid #ddd;
  font-weight: 500;
}

.marquee {
  display: inline-block;
  padding-left: 100%;
  animation: scroll-left 15s linear infinite;
}

@keyframes scroll-left {
  0% {
    transform: translateX(0%);
  }
  100% {
    transform: translateX(-100%);
  }
}
</style>