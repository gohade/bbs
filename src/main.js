import Vue from 'vue'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import store from './store'
// import VueRouter from 'vue-router'
import router from './router/index.js'

Vue.use(ElementUI)

new Vue({
  el: '#app',
  router: router,
  store,
  render: h => h(App)
})
