<template>
  <div>
    <!-- Hero Section -->
    <div class="hero">
  <!-- Make Swiper the background -->
  <Swiper
    class="hero-swiper"
    :modules="[EffectCoverflow, Autoplay, Pagination]"
    :effect="'coverflow'"
    :grabCursor="true"
    :centeredSlides="true"
    :slidesPerView="1"  
    :loop="true"
    :autoplay="{ delay: 2500, disableOnInteraction: false }"
    :coverflowEffect="{
      rotate: 30,
      stretch: 0,
      depth: 100,
      modifier: 1,
      slideShadows: true
    }"
    :pagination="{ clickable: true }"
  >
    <SwiperSlide v-for="(image, index) in images" :key="index">
      <img :src="image" :alt="'Image ' + (index + 1)" class="carousel-img" />
    </SwiperSlide>
  </Swiper>

  <!-- Overlay -->
  <div class="hero-overlay"></div>

  <!-- Centered Text -->
  <div class="hero-text">
    <h1>Craft Your Unique Scent</h1>
    <p>
      Design a personalized perfume by choosing your favorite Top, Heart, and Base notes.
      Express your individuality through fragrance!
    </p>
    <router-link to="/create">
      <button class="cta-button">Start Creating</button>
    </router-link>
  </div>
</div>

    <!-- Notes Section -->
    <section class="notes-preview">
      <h2>Explore Fragrance Notes</h2>
      <div class="notes-grid">
        <div
          v-for="note in notes"
          :key="note.title"
          class="note-card"
          @click="openModal(note)"
        >
          <img :src="note.image" alt="Note" />
          <div class="note-content">
            <h3>{{ note.title }}</h3>
            <p>{{ note.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Modal -->
    <Modal v-if="activeNote" :note="activeNote" @close="activeNote = null" />

 
  </div>
</template>

<script setup>
import { ref } from 'vue'
import Modal from '../components/Modal.vue'

// ✅ Import Swiper and modules correctly
import { Swiper, SwiperSlide } from 'swiper/vue'
import { EffectCoverflow, Autoplay, Pagination } from 'swiper/modules'

// ✅ Import Swiper styles
import 'swiper/css'
import 'swiper/css/effect-coverflow'
import 'swiper/css/pagination'

const activeNote = ref(null)

const notes = [
  {
    title: "Top Notes",
    image: new URL('@/assets/image1.jpeg', import.meta.url).href,
    description: "The first impression, usually fresh and light scents.",
    detail: "Top notes are the first scents you smell when you apply a fragrance. They are light and refreshing, often citrusy or herbal. They evaporate quickly and are designed to grab attention."
  },
  {
    title: "Heart Notes",
    image: new URL('@/assets/image2.jpg', import.meta.url).href,
    description: "The essence of the fragrance, rich and balanced.",
    detail: "Heart notes form the core of the fragrance, emerging after the top notes fade. These are typically floral, spicy, or fruity and last longer, providing harmony and depth."
  },
  {
    title: "Base Notes",
    image: new URL('@/assets/image3.jpg', import.meta.url).href,
    description: "The long-lasting foundation with deep and musky tones.",
    detail: "Base notes appear once the top and heart notes evaporate. They are rich, heavy, and long-lasting—usually including musk, vanilla, amber, and woods. They give the fragrance its staying power."
  }
]

const images = [
  new URL('@/assets/image1.jpeg', import.meta.url).href,
  new URL('@/assets/image2.jpg', import.meta.url).href,
  new URL('@/assets/image3.jpg', import.meta.url).href,
  new URL('@/assets/image4.jpg', import.meta.url).href,
  new URL('@/assets/image5.webp', import.meta.url).href
]

function openModal(note) {
  activeNote.value = note
}
</script>

<style scoped>
/* Hero Section */
.hero {
  position: relative;
  width: 100%;
  height: 100vh; /* Use full viewport height */
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
}

.hero-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(12px);
  z-index: 0;
}

.hero-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  color: #fff;
  max-width: 900px;
  width: 100%;
  z-index: 2;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
  text-align: center;
}

.hero-carousel {
  width: 100%;
  max-width: 1000px;
  height: 250px;
}

.mySwiper .swiper-slide {
  display: flex;
  justify-content: center;
}

.carousel-img {
  width: 100%;
  height: 100vh;
  object-fit: cover;
}

/* Hero text */
.hero h1 {
  font-size: 3.5rem;
  font-weight: 700;
  margin-bottom: 1rem;
  letter-spacing: 1px;
}

.hero p {
  font-size: 1.3rem;
  margin-bottom: 2rem;
  line-height: 1.6;
}

.cta-button {
  padding: 0.9rem 2.2rem;
  background: linear-gradient(to right, #ff6b6b, #ff9a9e);
  border: none;
  border-radius: 30px;
  color: white;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 2rem;
}

.cta-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 15px rgba(0, 0, 0, 0.2);
}

/* Notes Section */
.notes-preview {
  text-align: center;
  padding: 4rem 2rem;
  margin-top: 4rem;
  background-color: transparent;
}

.notes-preview h2 {
  font-size: 2.5rem;
  margin-bottom: 2rem;
  font-weight: 700;
  color: #222;
}

.notes-grid {
  display: flex;
  gap: 2rem;
  flex-wrap: wrap;
  justify-content: center;
}

.note-card {
  background: #fff;
  border-radius: 16px;
  overflow: hidden;
  width: 300px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.06);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  cursor: pointer;
}

.note-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 16px 32px rgba(0, 0, 0, 0.12);
}

.note-card img {
  width: 100%;
  height: 200px;
  object-fit: cover;
}

.note-content {
  padding: 1rem;
}

.note-card h3 {
  font-size: 1.3rem;
  margin-bottom: 0.5rem;
  color: #333;
}

.note-card p {
  color: #666;
  font-size: 0.95rem;
  line-height: 1.4;
}

</style>