import Vue from 'vue' // 引入vue项目
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue' //引入App组件
import store from './store'
import router from './router/index.js'

Vue.use(ElementUI)

// 创建vue对象
new Vue({
  el: '#app',  // 绑定id为app的元素
  router: router,
  store: store,
  render: h => h(App) // 将App组建渲染在这个元素中
})
