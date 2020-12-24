import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
import Vuelidate from 'vuelidate';

import axios from 'axios';
import Vueaxios from 'vue-axios';

import App from './App.vue';
import router from './router';
import store from './store';
// scss style
import './assets/scss/index.scss';

Vue.config.productionTip = false;
// Install BootstrapVue
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
// Install Vuelidate
Vue.use(Vuelidate);
// Install axios,vue-axios
Vue.use(Vueaxios, axios);
new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
