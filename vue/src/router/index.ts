import { createRouter, createWebHistory } from 'vue-router'
import Game from '../components/Game.vue'
import Login from '../components/Login.vue'

const routes = [
  { path: '/:pathMatch(.*)*', redirect: '/' },
  {
    path: '/',
    name: 'Home',
    component: Game,
  },
  {
    path: '/game',
    name: 'Game',
    component: Game,
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
  }
]

// Create the router instance
const router = createRouter({
  history: createWebHistory(), // Use HTML5 History API
  routes,
})

export default router