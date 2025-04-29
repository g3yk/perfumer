<template>
  <div class="profile-container">
    <h1>User Profile</h1>

    <div v-if="userToken">
      <!-- Profile Details -->
      <div class="profile-details">
        <label for="username">Username:</label>
        <input type="text" id="username" v-model="username" placeholder="Enter your username" />
      </div>

      <div class="profile-details">
        <label for="email">Email:</label>
        <input type="email" id="email" v-model="email" placeholder="Enter your email" disabled />
      </div>

      <!-- Address Management -->
      <div class="addresses-container">
        <h2>Your Addresses</h2>

        <div v-for="(address, index) in addresses" :key="index" class="address-item">
          <p><strong>{{ address.name }}</strong> - {{ address.street }}, {{ address.city }}, {{ address.zip }}</p>
          <button @click="startEditingAddress(index)">Edit</button>
          <button @click="deleteAddress(index)" style="background-color: red;">Delete</button>
        </div>

        <button @click="startAddingAddress" class="save-button" style="margin-top: 1rem;">
          + Add New Address
        </button>

        <!-- Address Form -->
        <div v-if="showAddressForm" style="margin-top: 2rem;">
          <h3>{{ isEditing ? 'Edit Address' : 'Add a New Address' }}</h3>

          <div class="profile-details">
            <label for="name">Address Name:</label>
            <input type="text" id="name" v-model="newName" placeholder="Enter address name (e.g., Home, Office)" />
          </div>

          <div class="profile-details">
            <label for="street">Street Address:</label>
            <input type="text" id="street" v-model="newStreet" placeholder="Enter street address" />
          </div>

          <div class="profile-details">
            <label for="city">City:</label>
            <input type="text" id="city" v-model="newCity" placeholder="Enter city" />
          </div>

          <div class="profile-details">
            <label for="zip">Zip Code:</label>
            <input type="text" id="zip" v-model="newZip" placeholder="Enter zip code" />
          </div>

          <button @click="saveAddress" class="save-button">
            {{ isEditing ? 'Update Address' : 'Save Address' }}
          </button>
        </div>
      </div>

      <!-- Orders Section -->
      <div class="orders-container" style="margin-top: 3rem;">
        <h2>Your Orders</h2>
        <div v-if="orders.length === 0">
          <p>You have no orders yet.</p>
        </div>
        <div v-for="(order, index) in orders" :key="index" class="order-item">
          <p><strong>Order ID:</strong> {{ order.orderId }}</p>
          <p><strong>Date:</strong> {{ order.date }}</p>
          <p><strong>Status:</strong> {{ order.status }}</p>
          <p><strong>Total:</strong> ${{ order.total && !isNaN(order.total) ? order.total.toFixed(2) : '0.00' }}</p>
          <button @click="viewOrderDetails(order)">View Details</button>
        </div>
      </div>

      <button @click="saveProfile" class="save-button" style="margin-top: 2rem;">Save Profile</button>
    </div>

    <div v-else>
      <p>Please log in to view and edit your profile.</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const userToken = ref(localStorage.getItem('userToken'))
const username = ref(localStorage.getItem('username') || '')
const email = ref(localStorage.getItem('email') || '')
const addresses = ref([])
const orders = ref([])

const newName = ref('')
const newStreet = ref('')
const newCity = ref('')
const newZip = ref('')

const isEditing = ref(false)
const editingIndex = ref(null)
const showAddressForm = ref(false)

function getAddressKey() {
  return `addresses_${email.value}`
}

function getOrdersKey() {
  return `orders_${email.value}`
}

function loadAddresses() {
  const key = getAddressKey()
  const stored = localStorage.getItem(key)
  addresses.value = stored ? JSON.parse(stored) : []
}

function loadOrders() {
  const key = getOrdersKey()
  const storedOrders = localStorage.getItem('orders') || '[]'
  const stored = localStorage.getItem(key)
  orders.value = stored ? JSON.parse(stored) : []
  console.log(orders.value)  // Debugging line
}

onMounted(() => {
  loadOrders()
})

function saveAddresses() {
  const key = getAddressKey()
  localStorage.setItem(key, JSON.stringify(addresses.value))
}

function saveOrders() {
  const key = getOrdersKey()
  localStorage.setItem(key, JSON.stringify(orders.value))
}

function startAddingAddress() {
  if (showAddressForm.value) {
    showAddressForm.value = false
    clearAddressForm()
    isEditing.value = false
    editingIndex.value = null
  } else {
    isEditing.value = false
    editingIndex.value = null
    clearAddressForm()
    showAddressForm.value = true
  }
}

function startEditingAddress(index) {
  const address = addresses.value[index]
  newName.value = address.name
  newStreet.value = address.street
  newCity.value = address.city
  newZip.value = address.zip

  isEditing.value = true
  editingIndex.value = index
  showAddressForm.value = true
}

function saveAddress() {
  if (newName.value && newStreet.value && newCity.value && newZip.value) {
    const updatedAddress = {
      name: newName.value,
      street: newStreet.value,
      city: newCity.value,
      zip: newZip.value
    }

    if (isEditing.value && editingIndex.value !== null) {
      addresses.value[editingIndex.value] = updatedAddress
    } else {
      addresses.value.push(updatedAddress)
    }

    saveAddresses()
    clearAddressForm()
    isEditing.value = false
    editingIndex.value = null
    showAddressForm.value = false
  } else {
    alert("Please fill in all fields.")
  }
}

function deleteAddress(index) {
  addresses.value.splice(index, 1)
  saveAddresses()
}

function clearAddressForm() {
  newName.value = ''
  newStreet.value = ''
  newCity.value = ''
  newZip.value = ''
}

function saveProfile() {
  localStorage.setItem('username', username.value)
  localStorage.setItem('email', email.value)
  alert('Profile updated successfully!')
}

function viewOrderDetails(order) {
  const total = order.total && !isNaN(order.total) ? order.total : 0; // Fallback to 0 if total is invalid
  alert(`Order ID: ${order.orderId}\nDate: ${order.date}\nStatus: ${order.status}\nTotal: $${total.toFixed(2)}`)
}

onMounted(() => {
  if (!userToken.value) {
    alert('You must be logged in to view or edit your profile.')
  } else {
    loadAddresses()
    loadOrders()  // Load orders from localStorage
  }
})
</script>

<style scoped>
.profile-container {
  padding: 2rem;
  max-width: 600px;
  margin: auto;
}

h1 {
  text-align: center;
  margin-bottom: 2rem;
}

.profile-details {
  margin-bottom: 1.5rem;
}

label {
  font-weight: bold;
  display: block;
  margin-bottom: 0.5rem;
}

input {
  width: 100%;
  padding: 0.8rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.save-button {
  width: 100%;
  padding: 1rem;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
}

.save-button:hover {
  background-color: #218838;
}

.address-item {
  margin-bottom: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 8px;
}

.order-item {
  margin-bottom: 1rem;
  padding: 1rem;
  background: #e9ecef;
  border-radius: 8px;
}

button {
  margin-top: 0.5rem;
  margin-right: 0.5rem;
  padding: 0.5rem 1rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>