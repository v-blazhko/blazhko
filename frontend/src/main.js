import Vue from 'vue/dist/vue.js'
import BootstrapVue from "bootstrap-vue"
import Meta from 'vue-meta'
import App from '../src/App.vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Aboutme from '../src/components/Aboutme.vue'
import Navbar from '../src/components/Navbar.vue'
import Myfooter from '../src/components/Myfooter.vue'
import Contactform from '../src/components/Contactform.vue'
import Game from '../src/components/Game.vue'
import Skill from '../src/components/Skill.vue'
import axios from 'axios'
import VueYaMetrica from 'vue-ya-metrica'
import VueAxios from 'vue-axios'
import VueYandexMetrika from 'vue-yandex-metrika'
import VueRecaptcha from 'vue-recaptcha'
Vue.prototype.$axios = axios
window.axios = require('axios');

var Lang = require('vuejs-localization');

Vue.use(VueYandexMetrika, {
    id: 48711056,
    env: process.env.NODE_ENV
})
Vue.use(BootstrapVue)
Vue.use(Meta)
Vue.use(VueAxios, axios)
Vue.config.productionTip = false
Vue.component('Aboutme', Aboutme)
Vue.component('Navbar', Navbar)
Vue.component('VueYaMetrica', VueYaMetrica)
Vue.component('Contactform', Contactform)
Vue.component('VueRecaptcha', VueRecaptcha)
Vue.component('Game', Game)
Vue.component('Skill', Skill)
Vue.component('Myfooter', Myfooter)
Lang.requireAll(require.context('../lang/', true, /\.js$/));
Vue.use(Lang);
/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  components: { App }
})
