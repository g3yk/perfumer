<template>
  <div class="checkout">
    <h1>Checkout</h1>

    <!-- Show if cart is empty -->
    <div v-if="cart.length === 0">
      Your cart is empty.
    </div>

    <!-- Show cart items if cart is not empty -->
    <div v-else>
      <div v-for="item in cart" :key="item.id" class="checkout-item">
        <h3>Custom Perfume</h3>
        <p><strong>Top:</strong> {{ item.topNotes?.join(', ') || 'None' }}</p>
        <p><strong>Heart:</strong> {{ item.heartNotes?.join(', ') || 'None' }}</p>
        <p><strong>Base:</strong> {{ item.baseNotes?.join(', ') || 'None' }}</p>

        <!-- Quantity Control -->
        <div class="quantity-control">
          <button @click="decreaseQuantity(item)">➖</button>
          <span>{{ item.quantity }}</span>
          <button @click="increaseQuantity(item)">➕</button>
        </div>

        <p><strong>Price:</strong> €{{ (item.quantity * 49.99).toFixed(2) }}</p>
      </div>

      <!-- Total Amount -->
      <div class="total-section">
        <h2>Total: €{{ totalAmount }}</h2>
        <button
          v-if="!isUserLoggedIn"
          class="pay-btn"
          @click="openLoginModal"
        >
          Sign Up / Login
        </button>

        <!-- Show proceed to payment if user is logged in and address is selected -->
        <button
          v-else
          class="pay-btn"
          @click="goToPayment"
          :disabled="!isAddressSelected"
        >
          Proceed to Payment
        </button>
      </div>
    </div>

    <!-- Address Section -->
    <div v-if="isUserLoggedIn" class="address-section">
      <h2>Select or Enter Address</h2>

      <!-- Select from saved addresses -->
      <div v-if="savedAddresses.length > 0 && !isEnteringNewAddress">
        <label for="address">Select an address:</label>
        <select v-model="selectedAddress" id="address">
          <option v-for="address in savedAddresses" :key="address.id" :value="address">
            {{ address.street }}, {{ address.city }}, {{ address.zip }}
          </option>
        </select>
      </div>

      <!-- Toggle for entering a new address -->
      <button @click="toggleEnterNewAddress">
        {{ isEnteringNewAddress ? 'Cancel' : 'Enter New Address' }}
      </button>

      <!-- Enter new address if toggle is active -->
      <div v-if="isEnteringNewAddress" class="new-address-form">
        <h3>Enter New Address</h3>
        <div class="form-group">
          <label for="street">Street</label>
          <input
            v-model="newAddress.street"
            id="street"
            type="text"
            placeholder="123 Main St"
            required
            pattern=".{3,}"
          />
        </div>
        <div class="form-group">
          <label for="city">City</label>
          <input
            v-model="newAddress.city"
            id="city"
            type="text"
            placeholder="e.g., Paris"
            required
            pattern="[A-Za-z\s]{2,}"
          />
        </div>
        <div class="form-group">
          <label for="zip">Zip Code</label>
          <input
            v-model="newAddress.zip"
            id="zip"
            type="text"
            placeholder="75001"
            required
            @input="filterZipInput"
            maxlength="10"
            inputmode="numeric"
          />
        </div>
      </div>
    </div>

    <!-- Login/Register Modal -->
    <div v-if="showLoginModal" class="auth-modal" @click="closeLoginModal">
      <div class="modal-content" @click.stop>
        <button class="close-btn" @click="closeLoginModal">✖</button>
        <h2>{{ isRegistering ? 'Register' : 'Login' }}</h2>

        <!-- Conditional form based on Login/Register state -->
        <div v-if="!isRegistering">
          <div class="form-group">
            <input v-model="email" type="email" placeholder="Enter your email" required />
            <span v-if="email && !isValidEmail(email)" class="error-message">Please enter a valid email.</span>
          </div>
          <div class="form-group">
            <input v-model="password" type="password" placeholder="Enter your password" required />
            <span v-if="password && password.length < 6" class="error-message">Password must be at least 6 characters.</span>
          </div>
          <button @click="handleLogin" :disabled="!isValidLogin">Login</button>
          <p>Don't have an account? <span @click="toggleForm">Register here</span></p>
        </div>

        <div v-if="isRegistering">
          <div class="form-group">
            <input v-model="username" type="text" placeholder="Enter your username" required />
            <span v-if="username && username.length < 3" class="error-message">Username must be at least 3 characters.</span>
          </div>
          <div class="form-group">
            <input v-model="email" type="email" placeholder="Create your email" required />
            <span v-if="email && !isValidEmail(email)" class="error-message">Please enter a valid email.</span>
          </div>
          <div class="form-group">
            <input v-model="password" type="password" placeholder="Create a password" required />
            <span v-if="password && password.length < 6" class="error-message">Password must be at least 6 characters.</span>
          </div>
          <button @click="handleRegister" :disabled="!isValidRegister">Register</button>
          <p>Already have an account? <span @click="toggleForm">Login here</span></p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const cart = ref([])
const router = useRouter()
const showLoginModal = ref(false)
const isRegistering = ref(false)
const email = ref('')
const password = ref('')
const username = ref('')
const isUserLoggedIn = ref(false)
const savedAddresses = ref([])
const selectedAddress = ref(null)  // The selected address will be an object, not just an ID
const newAddress = ref({
  street: '',
  city: '',
  zip: ''
})
const isEnteringNewAddress = ref(false)  // Track if the user wants to enter a new address

onMounted(() => {
  const savedCart = JSON.parse(localStorage.getItem('cart')) || []
  cart.value = savedCart.map(item => ({
    ...item,
    quantity: item.quantity ?? 1
  }))

  const userToken = localStorage.getItem('userToken')
  isUserLoggedIn.value = !!userToken

  const currentUserEmail = localStorage.getItem('email')

  if (currentUserEmail) {
    savedAddresses.value = JSON.parse(localStorage.getItem(`addresses_${currentUserEmail}`)) || []
  }

  // Load the username from localStorage
  const savedUsername = localStorage.getItem('username')
  if (savedUsername) {
    username.value = savedUsername
  }
})

function increaseQuantity(item) {
  item.quantity++
  saveCart()
}

function decreaseQuantity(item) {
  if (item.quantity > 1) {
    item.quantity--
  } else {
    cart.value = cart.value.filter(p => p.id !== item.id)
  }
  saveCart()
}

function saveCart() {
  localStorage.setItem('cart', JSON.stringify(cart.value))
}

const totalAmount = computed(() => {
  return cart.value
    .reduce((total, item) => total + item.quantity * 49.99, 0)
    .toFixed(2)
})

function openLoginModal() {
  showLoginModal.value = true
}

function closeLoginModal() {
  showLoginModal.value = false
}

function goToPayment() {
  if (cart.value.length === 0) {
    alert('Your cart is empty.')
    return
  }

  const userToken = localStorage.getItem('userToken')

  if (!userToken) {
    openLoginModal()
    return
  }

  if (!isAddressSelected.value) {
    alert('Please select or enter an address before proceeding.')
    return
  }

  router.push('/payment')
}

function handleLogin() {
  const storedEmail = localStorage.getItem('email')
  const storedPassword = localStorage.getItem('password')

  if (storedEmail === email.value && storedPassword === password.value) {
    localStorage.setItem('userToken', 'fake-token')
    closeLoginModal()
    isUserLoggedIn.value = true
    router.push('/checkout')
  } else {
    alert('Invalid email or password! Please try again or register.')
  }
}

function handleRegister() {
  if (!username.value || !email.value || !password.value) {
    alert('Please provide a username, email, and password.')
    return
  }

  // Save username along with email and password
  localStorage.setItem('username', username.value)
  localStorage.setItem('email', email.value)
  localStorage.setItem('password', password.value)
  localStorage.setItem('userToken', 'fake-token')

  closeLoginModal()
  isUserLoggedIn.value = true
  router.push('/checkout')
}

function toggleForm() {
  isRegistering.value = !isRegistering.value
}

function toggleEnterNewAddress() {
  isEnteringNewAddress.value = !isEnteringNewAddress.value
}

function filterZipInput(e) {
  newAddress.value.zip = e.target.value.replace(/\D/g, '')
}

// Computed property to check if an address is selected (either saved or new)
const isAddressSelected = computed(() => {
  return selectedAddress.value || (newAddress.value.street && newAddress.value.city && newAddress.value.zip)
})

function isValidEmail(value) {
  const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
  return emailPattern.test(value)
}

const isValidLogin = computed(() => {
  return email.value && password.value.length >= 6 && isValidEmail(email.value)
})

const isValidRegister = computed(() => {
  return username.value.length >= 3 && email.value && password.value.length >= 6 && isValidEmail(email.value)
})
</script>

<style scoped>
/* Existing Styles */
.checkout {
  max-width: 600px;
  margin: auto;
  padding: 2rem;
}

.checkout-item {
  border: 1px solid #ddd;
  padding: 1rem;
  margin-bottom: 1rem;
  border-radius: 8px;
}

.quantity-control {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 1rem 0;
}

.quantity-control button {
  background-color: #eee;
  border: none;
  padding: 0.3rem 0.7rem;
  font-size: 1.2rem;
  cursor: pointer;
  border-radius: 5px;
}

.quantity-control span {
  font-size: 1.2rem;
}

.total-section {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 2rem;
}

.pay-btn {
  background-color: #4CAF50;
  color: white;
  border: none;
  padding: 0.7rem 1.5rem;
  font-size: 1.2rem;
  cursor: pointer;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.pay-btn:hover {
  background-color: #45a049;
}

.address-section {
  margin-top: 2rem;
}

.auth-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: #fff;
  padding: 2rem;
  border-radius: 8px;
  width: 100%;
  max-width: 400px;
  position: relative;
  box-shadow: 0px 10px 15px rgba(0, 0, 0, 0.1);
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  background-color: transparent;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
}

.modal-content input {
  width: 92%;
  padding: 0.8rem;
  margin: 0.5rem 0;
  border-radius: 5px;
  border: 1px solid #ddd;
}

button {
  width: 100%;
  padding: 0.8rem;
  border-radius: 5px;
  border: none;
  background-color: #4CAF50;
  color: white;
  font-size: 1rem;
  margin-top: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #45a049;
}

button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

form p {
  text-align: center;
  margin-top: 1rem;
}

form p span {
  color: #4CAF50;
  cursor: pointer;
  text-decoration: underline;
}

.error-message {
  color: red;
  font-size: 0.9rem;
  margin-top: 0.3rem;
}

/* Address Section Styles */
.address-section {
  margin-top: 2rem;
  padding: 1.5rem;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.address-section h2 {
  font-size: 1.8rem;
  margin-bottom: 1.2rem;
  color: #333;
}

.address-section label {
  display: block;
  font-size: 1rem;
  margin-bottom: 0.5rem;
  color: #666;
}

.address-section select,
.address-section input {
  width: 100%;
  padding: 0.8rem;
  margin-bottom: 1rem;
  border-radius: 8px;
  border: 1px solid #ddd;
  font-size: 1rem;
  transition: border-color 0.3s ease;
}

.address-section select:focus,
.address-section input:focus {
  border-color: #4CAF50;
  outline: none;
}

.address-section .form-group {
  margin-bottom: 1rem;
}

.address-section button {
  background-color: #4CAF50;
  color: white;
  padding: 0.8rem 1.5rem;
  font-size: 1rem;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  display: inline-block;
  margin-top: 1rem;
}

.address-section button:hover {
  background-color: #45a049;
}

.address-section button:focus {
  outline: none;
}

.address-section .new-address-form {
  margin-top: 2rem;
  padding: 1.5rem;
  background-color: #f1f1f1;
  border-radius: 8px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}

.address-section .new-address-form h3 {
  font-size: 1.5rem;
  margin-bottom: 1rem;
}

.address-section .new-address-form .form-group {
  margin-bottom: 1rem;
}

.address-section .new-address-form input {
  font-size: 1rem;
}

.address-section .new-address-form button {
  background-color: #4CAF50;
  color: white;
  padding: 0.8rem 1.5rem;
  font-size: 1rem;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  display: inline-block;
}

.address-section .new-address-form button:hover {
  background-color: #45a049;
}

.address-section button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}
</style>