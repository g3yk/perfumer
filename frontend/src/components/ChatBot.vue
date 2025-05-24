<template>
  <div>
    <!-- Chatbot toggle button -->
    <button class="chatbot-button" @click="toggleChat">
      <img src="../assets/chat.png" alt="Chatbot Icon" class="chat-icon" />
      
    </button>

    <!-- Animated chat window -->
    <transition name="expand">
      <div v-if="isChatOpen" class="chat-window">
        <div class="chat-header">
          <span>Chatbot</span>
          <button class="close-button" @click="toggleChat">X</button>
        </div>

        <transition-group name="message" tag="div" class="chat-body">
          <div v-for="(msg, index) in messages" :key="index" :class="msg.isUser ? 'user-message' : 'bot-message'" class="message-item">

            {{ msg.text }}
          </div>
        </transition-group>

        <div class="chat-footer">
          <input v-model="messageInput" @keyup.enter="sendMessage" placeholder="Type a message..." />
          <button @click="sendMessage">Send</button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isChatOpen: false,
      messageInput: "",
      messages: [],
    };
  },
  methods: {
    toggleChat() {
      this.isChatOpen = !this.isChatOpen;
    },
    addBotMessage(message) {
      this.messages.push({ text: message, isUser: false });
    },
    sendMessage() {
      const text = this.messageInput.trim();
      if (text === "") return;

      // Add user message
      this.messages.push({ text, isUser: true });

      // Clear input field
      this.messageInput = "";

      // TODO: Add chatbot response logic here if needed
      // fetch request to localhost to get response for bot response

      let currentText = '';
      let index = this.messages.push({ text: '', isUser: false }) - 1; 
      
      fetch('http://localhost:11434/api/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          model: 'perfume-assistant',
          // prompt: 'What suggest you to me?'
          prompt: text
        })
      }).then(response => {
        const reader = response.body.getReader();
        const decoder = new TextDecoder();
        let result = '';

        const read = () =>{
          return reader.read().then(({ done, value }) => {
            if (done) {
              console.log('Stream complete');
              return;
            }
              
            let response = decoder.decode(value, { stream: true });
            let data = JSON.parse(response);
              
            currentText += data['response'] ;
            console.log(currentText)
            this.messages[index].text = currentText;
              
            return read();
          });
        }
        return read();
      });
      
    },
  },

  mounted() {
      this.addBotMessage("Hello! How can I help you today?");
  }
};
</script>

<style scoped>

.chat-icon {
  width: 30px;
  height: 30px
}

/* Chatbot toggle button */
.chatbot-button {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 1000;
  cursor: pointer;
  padding: 10px 15px;
  border: none;
  background-color: #ff9a9e;
  border-radius: 50%;
  font-size: 16px;
}

.chatbot-button:hover {
  background-color: #ff6b6b;
  transform: translateY(-2px);
  transition: all 0.3s ease;
}

/* Chat window */
.chat-window {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 300px;
  height: 400px;
  background: white;
  border: 1px solid #ccc;
  border-radius: 10px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  transform-origin: bottom right;
  overflow: hidden;
  z-index: 1000;
}

/* Chat header */
.chat-header {
  background: #ff6b6b;
  color: white;
  padding: 10px;
  display: flex;
  justify-content: space-between;
  border-top-left-radius: 10px;
  border-top-right-radius: 10px;
}

/* Chat body */
.chat-body {
  flex: 1;
  padding: 10px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
}

/* Chat footer */
.chat-footer {
  padding: 10px;
  display: flex;
}

.chat-footer input {
  flex: 1;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.chat-footer button {
  margin-left: 5px;
  padding: 5px 10px;
  background: #ff6b6b;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

/* Base style for messages */
.message-item {
  margin: 5px 0;
  padding: 8px;
  border-radius: 8px;
  max-width: 70%;
  word-wrap: break-word;
}

/* User and bot message styles */
.user-message {
  text-align: right;
  background-color: #dcf8c6;
  margin: 5px 0;
  padding: 8px;
  border-radius: 8px;
  align-self: flex-end;
}

.bot-message {
  text-align: left;
  background-color: #f1f0f0;
  margin: 5px 0;
  padding: 8px;
  border-radius: 8px;
  align-self: flex-start;
}

/* TransitionGroup enter animation */
.message-enter-active {
  transition: transform 0.4s ease, opacity 0.4s ease;
  display: flex;
}

/* User message enters: slide from left to right */
.bot-message.message-enter-from {
  transform: translateX(-100%);
  opacity: 0;
}

.bot-message.message-enter-to {
  transform: translateX(0);
  opacity: 1;
}

/* Bot message enters: slide from right to left */
.user-message.message-enter-from {
  transform: translateX(100%);
  opacity: 0;
}

.user-message.message-enter-to {
  transform: translateX(0);
  opacity: 1;
}

/* Transition animation */
.expand-enter-from {
  transform: scale(0) translate(100%, 100%);
  opacity: 0;
}

.expand-enter-to {
  transform: scale(1) translate(0, 0);
  opacity: 1;
}

.expand-leave-from {
  transform: scale(1) translate(0, 0);
  opacity: 1;
}

.expand-leave-to {
  transform: scale(0) translate(100%, 100%);
  opacity: 0;
}

.expand-enter-active,
.expand-leave-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  cursor: pointer;
  background: none;
  border: none;
  font-size: 20px;
}
</style>
