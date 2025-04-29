<template>
  <div class="builder">
    <h1>Create Your Perfume</h1>

    <div class="selection">
      <label>Top Notes (at least 1, up to 3):</label>
      <div class="options">
        <div v-for="note in topNotes" :key="note">
          <input
            type="checkbox"
            :value="note"
            v-model="topNotesSelected"
            @change="checkLimit('top')"
            :id="'top-' + note"
          />
          <label :for="'top-' + note">{{ note }}</label>
        </div>
      </div>
      <p v-if="topNotesSelected.length === 0" class="error">Please select at least one Top Note.</p>
    </div>

    <div class="selection">
      <label>Heart Notes (at least 1, up to 3):</label>
      <div class="options">
        <div v-for="note in heartNotes" :key="note">
          <input
            type="checkbox"
            :value="note"
            v-model="heartNotesSelected"
            @change="checkLimit('heart')"
            :id="'heart-' + note"
          />
          <label :for="'heart-' + note">{{ note }}</label>
        </div>
      </div>
      <p v-if="heartNotesSelected.length === 0" class="error">Please select at least one Heart Note.</p>
    </div>

    <div class="selection">
      <label>Base Notes (at least 1, up to 3):</label>
      <div class="options">
        <div v-for="note in baseNotes" :key="note">
          <input
            type="checkbox"
            :value="note"
            v-model="baseNotesSelected"
            @change="checkLimit('base')"
            :id="'base-' + note"
          />
          <label :for="'base-' + note">{{ note }}</label>
        </div>
      </div>
      <p v-if="baseNotesSelected.length === 0" class="error">Please select at least one Base Note.</p>
    </div>

    <button @click="createPerfume">Add to Cart</button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { useCartStore } from '@/stores/cart'

const router = useRouter()
const toast = useToast()
const cartStore = useCartStore()

// Ref for selected notes
const topNotesSelected = ref([])
const heartNotesSelected = ref([])
const baseNotesSelected = ref([])

// Available options for each note category
const topNotes = ["Bergamot", "Lemon", "Mandarin", "Peppermint"]
const heartNotes = ["Rose", "Jasmine", "Lavender", "Cinnamon"]
const baseNotes = ["Vanilla", "Sandalwood", "Patchouli", "Amber"]

// Limit check function for each category
function checkLimit(type) {
  if (type === 'top' && topNotesSelected.value.length > 3) {
    topNotesSelected.value.pop()
    toast.error('You can select up to 3 Top Notes.', { timeout: 2000 })
  }
  if (type === 'heart' && heartNotesSelected.value.length > 3) {
    heartNotesSelected.value.pop()
    toast.error('You can select up to 3 Heart Notes.', { timeout: 2000 })
  }
  if (type === 'base' && baseNotesSelected.value.length > 3) {
    baseNotesSelected.value.pop()
    toast.error('You can select up to 3 Base Notes.', { timeout: 2000 })
  }
}

// Create custom perfume and add to cart
function createPerfume() {
  // Ensure at least 1 note is selected in each category
  if (topNotesSelected.value.length === 0 || heartNotesSelected.value.length === 0 || baseNotesSelected.value.length === 0) {
    toast.error('Please select at least one note from each category.', { timeout: 2000 })
    return
  }

  const customPerfume = {
    id: Date.now(),
    topNotes: topNotesSelected.value,
    heartNotes: heartNotesSelected.value,
    baseNotes: baseNotesSelected.value
  }

  cartStore.addToCart(customPerfume)

  toast.success('Perfume added to cart!', { timeout: 2000 })
  cartStore.isNavbarCartOpen = true
  router.push('/')
}
</script>

<style scoped>
.builder {
  max-width: 600px;
  margin: auto;
  padding: 2rem;
  text-align: center;
}

.selection {
  margin: 2rem 0;
}

.options {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 1rem;
  margin-top: 0.5rem;
}

.options div {
  display: flex;
  align-items: center;
}

input[type="checkbox"] {
  margin-right: 0.5rem;
  transform: scale(1.2);
}

button {
  margin-top: 2rem;
  padding: 0.8rem 1.5rem;
  border: none;
  background-color: #2ecc71;
  color: white;
  border-radius: 8px;
  cursor: pointer;
}

button:hover {
  background-color: #27ae60;
}

.error {
  color: red;
  font-size: 0.9rem;
  margin-top: 0.5rem;
}
</style>