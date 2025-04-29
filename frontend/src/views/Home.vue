<template>
    <div> 
      <!-- Hero Section -->
      <section class="hero">
        <div class="hero-text">
          <h1>Craft Your Unique Scent</h1>
          <p>Design a personalized perfume by choosing your favorite Top, Heart, and Base notes. Express your individuality through fragrance!</p>
          <router-link to="/create">
            <button class="cta-button">Start Creating</button>
          </router-link>
        </div>
      </section>
  
      <!-- Featured Notes Section -->
      <section class="notes-preview">
        <h2>Explore Fragrance Notes</h2>
        <div class="notes-grid">
          <div v-for="note in notes" :key="note.title" class="note-card">
            <img :src="note.image" alt="Note" />
            <h3>{{ note.title }}</h3>
            <p>{{ note.description }}</p>
          </div>
        </div>
      </section>
  
      <!-- Footer -->
      <footer class="footer">
        <p>© 2025 Custom Perfume Co. All rights reserved.</p>
      </footer>
    </div>
  </template>
  
  <script setup>
  import Navbar from '../components/Navbar.vue'
  
  const notes = [
    {
      title: "Top Notes",
      image: "https://images.unsplash.com/photo-1603052875779-8de2fbb4b98e?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
      description: "The first impression, usually fresh and light scents."
    },
    {
      title: "Heart Notes",
      image: "https://images.unsplash.com/photo-1601079737394-1fa3b04fcde0?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
      description: "The essence of the fragrance, rich and balanced."
    },
    {
      title: "Base Notes",
      image: "https://images.unsplash.com/photo-1614094109475-652262e7490d?ixlib=rb-4.0.3&auto=format&fit=crop&w=800&q=80",
      description: "The long-lasting foundation with deep and musky tones."
    }
  ]
  </script>
  
  <style scoped>
  .hero {
    background: url('https://images.unsplash.com/photo-1612347202328-7dc086e4040d?ixlib=rb-4.0.3&auto=format&fit=crop&w=1200&q=80') no-repeat center center;
    background-size: cover;
    height: 90vh;
    color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
  }
  
  .hero-text {
    background-color: rgba(0, 0, 0, 0.5);
    padding: 2rem;
    border-radius: 12px;
  }
  
  .hero h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
  }
  
  .hero p {
    font-size: 1.2rem;
    margin-bottom: 2rem;
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
  
  .notes-preview {
    text-align: center;
    padding: 4rem 2rem;
  }
  
  .notes-grid {
    display: flex;
    gap: 2rem;
    flex-wrap: wrap;
    justify-content: center;
    margin-top: 2rem;
  }
  
  .note-card {
    background: #fff;
    border: 1px solid #eee;
    border-radius: 12px;
    overflow: hidden;
    width: 280px;
    transition: 0.3s;
  }
  
  .note-card:hover {
    transform: translateY(-5px);
    box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1);
  }
  
  .note-card img {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }
  
  .note-card h3 {
    margin: 1rem 0 0.5rem;
  }
  
  .note-card p {
    padding: 0 1rem 1rem;
    color: #555;
  }
  
  .footer {
    text-align: center;
    padding: 2rem;
    font-size: 0.9rem;
    color: #888;
    background: #fafafa;
    margin-top: 4rem;
  }
  </style>