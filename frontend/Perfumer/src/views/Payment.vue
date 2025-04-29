<template>
    <div class="payment">
      <h1>Payment</h1>
  
      <form @submit.prevent="submitPayment" class="payment-form">
        <div class="form-group">
          <label>Cardholder Name</label>
          <input v-model="cardName" required />
        </div>
  
        <div class="form-group">
          <label>Card Number</label>
          <input v-model="cardNumber" required maxlength="16" />
        </div>
  
        <div class="form-group">
          <label>Expiry Date</label>
          <input v-model="expiry" required placeholder="MM/YY" />
        </div>
  
        <div class="form-group">
          <label>CVV</label>
          <input v-model="cvv" required maxlength="3" />
        </div>
  
        <button class="pay-btn" type="submit">Pay Now</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  import { useToast } from 'vue-toastification'
  import { useRouter } from 'vue-router'
  import { useCartStore } from '@/stores/cart'
  
  const cardName = ref('')
  const cardNumber = ref('')
  const expiry = ref('')
  const cvv = ref('')
  
  const toast = useToast()
  const router = useRouter()
  const cartStore = useCartStore() 
  
  function submitPayment() {
    // Check if all fields are filled
    if (!cardName.value || !cardNumber.value || !expiry.value || !cvv.value) {
      toast.error('Please fill all payment details correctly.')
      return
    }
  
    // Show success toast
    toast.success('Transaction has been successful! 🎉')
  
    // Create order object
    const order = {
      orderId: Math.floor(Math.random() * 90000) + 10000, // Random 5-digit order ID
      date: new Date().toISOString().split('T')[0],       // YYYY-MM-DD
      status: 'Processing',
      total: cartStore.cart.reduce((sum, item) => sum + item.price * item.quantity, 0)                     // Get total price from cart
    }
  
    console.log('Order created:', order) // Debugging log
  
    // Get user email from localStorage
    const userEmail = localStorage.getItem('email')
    if (!userEmail) {
      toast.error('User not logged in. Please log in first.')
      return
    }
  
    // Save order to localStorage
    const orderKey = `orders_${userEmail}`
    const existingOrders = JSON.parse(localStorage.getItem(orderKey)) || []
    existingOrders.unshift(order)
    localStorage.setItem(orderKey, JSON.stringify(existingOrders))
  
    console.log('Orders saved to localStorage:', existingOrders) // Debugging log
  
    // Clear cart
    cartStore.cart = []
    localStorage.removeItem('cart')
  
    // Redirect after 2 seconds
    setTimeout(() => {
      router.push('/order-placed')
    }, 2000)
  }
  </script>
  
  <style scoped>
  .payment {
    max-width: 500px;
    margin: auto;
    padding: 2rem;
  }
  
  .payment-form {
    display: flex;
    flex-direction: column;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  input {
    padding: 0.6rem;
    border: 1px solid #ddd;
    border-radius: 6px;
    width: 100%;
  }
  
  .pay-btn {
    background-color: #2ecc71;
    color: white;
    padding: 0.8rem 1.2rem;
    border: none;
    border-radius: 8px;
    cursor: pointer;
  }
  
  .pay-btn:hover {
    background-color: #27ae60;
  }
  </style>