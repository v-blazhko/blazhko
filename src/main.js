import Vue from 'vue'
import BootstrapVue from "bootstrap-vue"
import App from '../src/App.vue'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import Aboutme from '../src/components/Aboutme.vue'
import Navbar from '../src/components/Navbar.vue'
import Myfooter from '../src/components/Myfooter.vue'
import Contactform from '../src/components/Contactform.vue'
import Game from '../src/components/Game.vue'
import Skill from '../src/components/Skill.vue'

var Lang = require('vuejs-localization');
Vue.use(BootstrapVue)
Vue.http.options.root = 'https://v-blazhko.github.io/blazhko/';
Vue.config.productionTip = false
Vue.component('Aboutme', Aboutme)
Vue.component('Navbar', Navbar)
Vue.component('Contactform', Contactform)
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
