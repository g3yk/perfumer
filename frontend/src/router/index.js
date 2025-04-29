import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Create from '../views/Create.vue'
import Cart from '../views/Cart.vue'
import Checkout from '../views/Checkout.vue'
import Payment from '../views/Payment.vue'
import OrderPlaced from '../views/OrderPlaced.vue' 
import Profile from '../views/Profile.vue' 
import SignUp from '../views/SignUp.vue'
import Login from '../views/Login.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/create', component: Create },
  { path: '/cart', component: Cart },
  { path: '/checkout', component: Checkout },
  { path: '/payment', component: Payment },
  { path: '/order-placed', component: OrderPlaced }, 
  { path: '/profile', component: Profile } ,
  { path: '/signup', component: SignUp },
  { path: '/login', component: Login },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router