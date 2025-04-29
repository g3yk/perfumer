<script setup>
import { useCartStore } from '@/stores/cart'

const cartStore = useCartStore()
</script>

<template>
  <div class="cart">
    <h1>Your Custom Perfumes</h1>

    <div v-if="cartStore.cart.length === 0">
      Your cart is empty.
    </div>

    <div v-else>
      <div v-for="item in cartStore.cart" :key="item.id" class="cart-item">
        <h3>Custom Perfume</h3>
        <p><strong>Top:</strong> {{ item.topNote }}</p>
        <p><strong>Heart:</strong> {{ item.heartNote }}</p>
        <p><strong>Base:</strong> {{ item.baseNote }}</p>
      </div>
    </div>
  </div>
</template>