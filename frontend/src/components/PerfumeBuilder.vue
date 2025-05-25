<template>
  <div class="builder">
    <h1>Create Your Perfume</h1>

    <!-- Top Notes Section -->
    <div class="selection">
      <label>Top Notes (at least 1, up to 3):</label>
      <div class="search-container">
        <input
          v-model="topNoteSearch"
          placeholder="Search Top Notes..."
          class="note-search"
        />
        <button @click="clearSearch('top')" class="clear-btn">Clear Search</button>
      </div>

      <div class="options" v-if="topNoteSearch && filteredTopNotes.length > 0">
        <div v-for="note in filteredTopNotes" :key="note" class="option">
          <input
            type="checkbox"
            :value="note"
            v-model="topNotesSelected"
            @change="checkLimit('top')"
            :id="'top-' + note"
            class="checkbox"
          />
          <label :for="'top-' + note" class="note-label">{{ note }}</label>
        </div>
      </div>
      <p v-if="topNotesSelected.length === 0" class="error">Please select at least one Top Note.</p>
      <p v-if="filteredTopNotes.length === 0 && topNoteSearch" class="error">No results found.</p>
      <div class="selected-notes">
        <strong>Selected Top Notes:</strong> {{ topNotesSelected.join(', ') }}
      </div>
    </div>

    <!-- Heart Notes Section -->
    <div class="selection">
      <label>Heart Notes (at least 1, up to 3):</label>
      <div class="search-container">
        <input
          v-model="heartNoteSearch"
          placeholder="Search Heart Notes..."
          class="note-search"
        />
        <button @click="clearSearch('heart')" class="clear-btn">Clear Search</button>
      </div>

      <div class="options" v-if="heartNoteSearch && filteredHeartNotes.length > 0">
        <div v-for="note in filteredHeartNotes" :key="note" class="option">
          <input
            type="checkbox"
            :value="note"
            v-model="heartNotesSelected"
            @change="checkLimit('heart')"
            :id="'heart-' + note"
            class="checkbox"
          />
          <label :for="'heart-' + note" class="note-label">{{ note }}</label>
        </div>
      </div>
      <p v-if="heartNotesSelected.length === 0" class="error">Please select at least one Heart Note.</p>
      <p v-if="filteredHeartNotes.length === 0 && heartNoteSearch" class="error">No results found.</p>
      <div class="selected-notes">
        <strong>Selected Heart Notes:</strong> {{ heartNotesSelected.join(', ') }}
      </div>
    </div>

    <!-- Base Notes Section -->
    <div class="selection">
      <label>Base Notes (at least 1, up to 3):</label>
      <div class="search-container">
        <input
          v-model="baseNoteSearch"
          placeholder="Search Base Notes..."
          class="note-search"
        />
        <button @click="clearSearch('base')" class="clear-btn">Clear Search</button>
      </div>

      <div class="options" v-if="baseNoteSearch && filteredBaseNotes.length > 0">
        <div v-for="note in filteredBaseNotes" :key="note" class="option">
          <input
            type="checkbox"
            :value="note"
            v-model="baseNotesSelected"
            @change="checkLimit('base')"
            :id="'base-' + note"
            class="checkbox"
          />
          <label :for="'base-' + note" class="note-label">{{ note }}</label>
        </div>
      </div>
      <p v-if="baseNotesSelected.length === 0" class="error">Please select at least one Base Note.</p>
      <p v-if="filteredBaseNotes.length === 0 && baseNoteSearch" class="error">No results found.</p>
      <div class="selected-notes">
        <strong>Selected Base Notes:</strong> {{ baseNotesSelected.join(', ') }}
      </div>
    </div>

    <button @click="createPerfume" class="add-to-cart-btn">Add to Cart</button>

    <div v-if="similarPerfumes.length > 0" class="similar-slider">
  <h2 class="slider-title">
    Similar Perfumes ({{ similarPerfumes.length }} match<span v-if="similarPerfumes.length > 1">es</span>)
  </h2>

  <swiper
    v-if="similarPerfumes.length > 0"
    :slides-per-view="1"
    :space-between="20"
    :breakpoints="{
      640: { slidesPerView: 1.3 },
      768: { slidesPerView: 2 },
      1024: { slidesPerView: 3 }
    }"
    :pagination="{ clickable: true }"
    class="perfume-swiper"
  >
    <swiper-slide
      v-for="perfume in similarPerfumes"
      :key="perfume.name"
      class="swiper-slide"
    >
      <div class="perfume-card">
        <img :src="perfume.image" alt="Perfume Image" class="perfume-image" />
        <div class="perfume-info">
          <h3>{{ perfume.name }}</h3>
          <p class="brand">{{ perfume.brand }}</p>
          <p class="match-score">
            Match Score: <br>{{ perfume.matchCount }} Notes
            <span v-if="perfume.isExactMatch" class="exact-match"><br>Exactly Matched</span>
          </p>
        </div>
      </div>
    </swiper-slide>
  </swiper>
</div>
  </div>
</template>
<script setup>
import { computed } from 'vue' 
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useToast } from 'vue-toastification'
import { useCartStore } from '@/stores/cart'
import { Swiper, SwiperSlide } from 'swiper/vue'
import { perfumeDataset } from '@/data/perfumeDataset'
import 'swiper/css'
import 'swiper/css/navigation'
import 'swiper/css/pagination'


const router = useRouter()
const toast = useToast()
const cartStore = useCartStore()

const topNoteSearch = ref("")
const heartNoteSearch = ref("")
const baseNoteSearch = ref("")

// Ref for selected notes
const topNotesSelected = ref([])
const heartNotesSelected = ref([])
const baseNotesSelected = ref([])

// Available options for each note category (Added 50 new notes for each)
const topNotes = [
  "Bergamot", "Lemon", "Mandarin", "Peppermint", "Grapefruit", "Orange", "Lime", "Cardamom", "Saffron", "Cedarwood",
  "Basil", "Apple", "Blackcurrant", "Pineapple", "Lavender", "Rosemary", "Violet", "Mint", "Eucalyptus", "Lemon Verbena",
  "Tangerine", "Yuzu", "Pink Pepper", "Rose", "Pear", "Cucumber", "Almond", "Cedar", "Anise", "Jasmine", "Litchi",
  "Neroli", "Ginger", "Fennel", "Tarragon", "Lemon Blossom", "Pine", "Vanilla", "Coconut", "Honey", "Geranium", "Mango",
  "Cinnamon", "Plum", "Lemon Balm", "Pomegranate", "Chili Pepper", "Freesia", "Juniper", "Clove", "Papaya"
];
const heartNotes = [
  "Rose", "Jasmine", "Lavender", "Cinnamon", "Clove", "Geranium", "Ylang Ylang", "Tuberose", "Violet", "Freesia",
  "Orchid", "Honeysuckle", "Magnolia", "Gardenia", "Lily", "Iris", "Coriander", "Apple Blossom", "Pepper", "Patchouli",
  "Marigold", "Lilac", "Peony", "Champa", "Anise", "Muguet", "Benzoin", "Sandalwood", "Cherry Blossom", "Pineapple",
  "Cardamom", "Ginger", "Saffron", "Bergamot", "Bitter Orange", "Cedarwood", "Mandarin Blossom", "Carnation", "Cistus",
  "Sage", "Clary Sage", "Basil", "Orange Blossom", "Coconut", "Lemon Blossom", "Oud", "Sea Notes", "Neroli"
];
const baseNotes = [
  "Vanilla", "Sandalwood", "Amber", "Musk", "Tonka Bean", "Leather", "Cedarwood", "Benzoin", "Oud", "Moss",
  "Labdanum", "Myrrh", "Civet", "Cashmere", "Tobacco", "Tonka", "Cocoa", "Cashmere Wood", "Saffron", "Frankincense",
  "Balsam Fir", "Vetiver", "Ambergris", "Almond", "Pine", "Gaiac Wood", "Oakmoss", "Patchouli", "Agarwood", "Tolu Balsam",
  "Galbanum", "Balsam", "Orris", "Tobacco Leaf", "Balsamic", "Sage", "Honey", "Cypress", "Caramel", "Cedar"
];
const allPerfumes = [...perfumeDataset]
// Computed filtered results for each category based on search input
const filteredTopNotes = computed(() =>
  topNotes.filter(note =>
    note.toLowerCase().startsWith(topNoteSearch.value.toLowerCase())
  )
)

const filteredHeartNotes = computed(() =>
  heartNotes.filter(note =>
    note.toLowerCase().startsWith(heartNoteSearch.value.toLowerCase())
  )
)

const filteredBaseNotes = computed(() =>
  baseNotes.filter(note =>
    note.toLowerCase().startsWith(baseNoteSearch.value.toLowerCase())
  )
)

const similarPerfumes = ref([])

watch([topNotesSelected, heartNotesSelected, baseNotesSelected], () => {
  const selectedNotes = new Set([
    ...topNotesSelected.value,
    ...heartNotesSelected.value,
    ...baseNotesSelected.value
  ])

  if (selectedNotes.size === 0) {
    similarPerfumes.value = []
    return
  }

  const scored = allPerfumes.map(perfume => {
    const perfumeNotes = new Set([
      ...perfume.topNotes,
      ...perfume.heartNotes,
      ...perfume.baseNotes
    ])
    let matchCount = 0
    for (const note of selectedNotes) {
      if (perfumeNotes.has(note)) {
        matchCount++
      }
    }
    
    // Check for exact match (all selected notes match exactly)
    const isExactMatch = selectedNotes.size === perfumeNotes.size && [...selectedNotes].every(note => perfumeNotes.has(note))

    return { ...perfume, matchCount, isExactMatch }
  })

  // Filter perfumes with at least 1 match and sort by match count descending
  similarPerfumes.value = scored
    .filter(p => p.matchCount > 0)
    .sort((a, b) => b.matchCount - a.matchCount)
})


function clearSearch(category) {
  // Clear the search field
  if (category === 'top') {
    topNoteSearch.value = ""
    topNotesSelected.value = []  // Clear the selected top notes
  }
  if (category === 'heart') {
    heartNoteSearch.value = ""
    heartNotesSelected.value = []  // Clear the selected heart notes
  }
  if (category === 'base') {
    baseNoteSearch.value = ""
    baseNotesSelected.value = []  // Clear the selected base notes
  }
}

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
  max-width: 700px;
  margin: auto;
  padding: 2rem;
  text-align: center;
  background: transparent;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

h1 {
  font-size: 2.2rem;
  color: #2c3e50;
  font-weight: bold;
  margin-bottom: 2rem;
}

.selection {
  margin: 2rem 0;
  background: #fff;
  border-radius: 8px;
  padding: 1rem;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
}

label {
  font-size: 1.1rem;
  font-weight: bold;
  color: #333;
  margin-bottom: 0.75rem;
  display: block;
}

.options {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  gap: 1rem;
  margin-top: 1rem;
  padding: 1rem;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.1);
  max-height: 200px;
  overflow-y: auto;
}

.options div {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1rem;
  color: #555;
}

input[type="checkbox"] {
  margin-right: 0.5rem;
  transform: scale(1.3);
  accent-color: #2ecc71; /* Checkbox accent color */
}

button {
  margin-top: 2rem;
  padding: 1rem 2rem;
  background-color: #2ecc71;
  color: white;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  font-size: 1.1rem;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  transition: background-color 0.3s ease, transform 0.2s ease;
}

button:hover {
  background-color: #27ae60;
  transform: translateY(-2px);
}

button:active {
  transform: translateY(2px);
}

.clear-btn {
  padding: 0.6rem 1.2rem;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease;
  font-size: 0.9rem;
}

.clear-btn:hover {
  background-color: #c0392b;
}

.error {
  color: red;
  font-size: 0.95rem;
  margin-top: 1rem;
  animation: fadeIn 1s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.similar-slider {
  padding: 1.5rem;
  text-align: left;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  margin-top: 2rem;
}

.slider-title {
  text-align: center;
  font-size: 1.7rem;
  margin-bottom: 1.5rem;
  color: #2c3e50;
  font-weight: bold;
}

.perfume-swiper {
  padding-bottom: 2rem;
}

.perfume-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  padding: 1.5rem;
  text-align: center;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.perfume-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
}

.perfume-image {
  width: 100%;
  height: 180px;
  object-fit: contain;
  border-radius: 12px;
  margin-bottom: 1rem;
}

.perfume-info h3 {
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
  color: #2c3e50;
}

.perfume-info .brand {
  color: #7f8c8d;
  font-size: 1rem;
}

.match-score {
  margin-top: 1rem;
  font-weight: bold;
  color: #2ecc71;
}

.note-search {
  width: 95%;
  padding: 0.8rem;
  margin: 1rem 0;
  border: 2px solid #ddd;
  border-radius: 8px;
  font-size: 1rem;
  background-color: #fafafa;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.note-search:focus {
  border-color: #2ecc71;
  outline: none;
}

.selected-notes {
  margin-top: 1.5rem;
  font-size: 1.2rem;
  font-weight: bold;
  color: #2c3e50;
}

.selected-notes span {
  color: #7f8c8d;
}

.selected-notes a {
  color: #2ecc71;
  font-weight: bold;
  text-decoration: none;
}

.selected-notes a:hover {
  text-decoration: underline;
}
.exact-match {
  color: green;
  font-weight: bold;
  margin-left: 10px;
}
</style>