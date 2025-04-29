<template>
    <div class="modal-overlay" @click.self="$emit('close')">
      <div class="modal-box">
        <h3>{{ note.title }}</h3>
        <p>{{ note.detail }}</p>
        <button @click="$emit('close')" class="close-button">Close</button>
      </div>
    </div>
</template>
  
  <script setup>
  defineProps(['note'])
  defineEmits(['close'])
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0; left: 0;
    width: 100vw; height: 100vh;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }
  
  .modal-box {
    background: white;
    padding: 2rem;
    border-radius: 16px;
    max-width: 500px;
    width: 90%;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
    animation: fadeIn 0.3s ease;
    text-align: center;
  }
  
  .modal-box h3 {
    font-size: 1.8rem;
    margin-bottom: 1rem;
  }
  
  .modal-box p {
    font-size: 1rem;
    color: #444;
    margin-bottom: 2rem;
    line-height: 1.6;
  }
  
  .close-button {
    background-color: #ff6b6b;
    color: white;
    padding: 0.6rem 1.5rem;
    border: none;
    border-radius: 30px;
    font-weight: 600;
    cursor: pointer;
    transition: background 0.3s ease;
  }
  
  .close-button:hover {
    background-color: #ff4d4d;
  }
  
  @keyframes fadeIn {
    from {
      transform: scale(0.95);
      opacity: 0;
    }
    to {
      transform: scale(1);
      opacity: 1;
    }
  }
  </style>