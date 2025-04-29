// cart.js (Pinia Store)
import { defineStore } from 'pinia'

export const useCartStore = defineStore('cart', {
  state: () => ({
    cart: [],                 // Cart items
    isCartOpen: false,        // For aside cart visibility
    isNavbarCartOpen: false,  // For navbar cart visibility
  }),
  actions: {
    // Add item to the cart
    addToCart(item) {
      this.cart.push(item)
      // Optionally open both carts
      this.isCartOpen = true
      this.isNavbarCartOpen = true  // Ensure navbar cart is opened too
    },
    // Remove item from the cart
    removeFromCart(id) {
      this.cart = this.cart.filter(item => item.id !== id)
    },
    // Toggle visibility for the aside cart
    toggleCart() {
      this.isCartOpen = !this.isCartOpen
    },
    // Toggle visibility for the navbar cart
    toggleNavbarCart() {
      this.isNavbarCartOpen = !this.isNavbarCartOpen
    }
  }
})