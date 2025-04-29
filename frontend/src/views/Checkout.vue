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
            <div v-if="isEnteringNewAddress">
                <h3>Enter new address</h3>
                <input v-model="newAddress.street" type="text" placeholder="Street Address" />
                <input v-model="newAddress.city" type="text" placeholder="City" />
                <input v-model="newAddress.zip" type="text" placeholder="Zip Code" />
            </div>
        </div>
  
        <!-- Login/Register Modal -->
        <div v-if="showLoginModal" class="auth-modal" @click="closeLoginModal">
            <div class="modal-content" @click.stop>
                <button class="close-btn" @click="closeLoginModal">✖</button>
                <h2>{{ isRegistering ? 'Register' : 'Login' }}</h2>
    
                <!-- Conditional form based on Login/Register state -->
                <div v-if="!isRegistering">
                    <input v-model="email" type="email" placeholder="Enter your email" />
                    <input v-model="password" type="password" placeholder="Enter your password" />
                    <button @click="handleLogin">Login</button>
                    <p>Don't have an account? <span @click="toggleForm">Register here</span></p>
                </div>
    
                <div v-if="isRegistering">
                    <input v-model="username" type="text" placeholder="Enter your username" />
                    <input v-model="email" type="email" placeholder="Create your email" />
                    <input v-model="password" type="password" placeholder="Create a password" />
    
                    <button @click="handleRegister">Register</button>
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

// Computed property to check if an address is selected (either saved or new)
const isAddressSelected = computed(() => {
  return selectedAddress.value || (newAddress.value.street && newAddress.value.city && newAddress.value.zip)
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
  font-weight: bold;
}

.total-section {
  text-align: center;
  margin-top: 2rem;
}

.pay-btn {
  background-color: #3498db;
  color: white;
  padding: 0.8rem 1.2rem;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin-top: 1rem;
}

.pay-btn:hover {
  background-color: #2980b9;
  cursor: pointer;
  opacity: 0.8;
  transition: opacity 0.3s;
}

.pay-btn:disabled {
  background-color: #ddd;
  cursor: not-allowed;
  opacity: 0.5;
}

/* Address Section Styles */
.address-section {
  margin-top: 2rem;
}

.address-section select,
.address-section input {
  width: 100%;
  padding: 0.8rem;
  margin-bottom: 1rem;
  border-radius: 4px;
}

button {
  margin-top: 1rem;
  padding: 0.8rem;
  cursor: pointer;
  background-color: #3498db;
  color: white;
  border-radius: 4px;
  border: none;
}

button:hover {
  background-color: #2980b9;
}

.auth-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1000;
  backdrop-filter: blur(5px); /* Apply blur effect */
}

.modal-content {
  background-color: rgba(255, 255, 255, 0.8); /* Semi-transparent white background */
  padding: 2rem;
  border-radius: 8px;
  text-align: center;
  width: 300px; /* Adjust width as needed */
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  margin: 20px; /* Margin around modal content */
}

.modal-content input {
  width: 75%;
  padding: 0.8rem;
  margin-bottom: 1rem;
  border-radius: 4px;
}

.modal-content button {
  background-color: #3498db;
  color: white;
  padding: 0.8rem;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.modal-content button:hover {
  background-color: #2980b9;
}

.modal-content span {
  color: #3498db;
  cursor: pointer;
}

.auth-modal .close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 1.5rem;
  background: none;
  border: none;
  color: #aaa;
  cursor: pointer;
}

.auth-modal .close-btn:hover {
  color: #000;
}
</style>